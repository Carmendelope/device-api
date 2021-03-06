/*
 * Copyright 2019 Nalej
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package server

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/nalej/authx/pkg/interceptor"
	"github.com/nalej/authx/pkg/interceptor/devinterceptor"
	"github.com/nalej/derrors"
	"github.com/nalej/device-api/internal/pkg/server/applications"
	"github.com/nalej/device-api/internal/pkg/server/device"
	"github.com/nalej/grpc-application-manager-go"
	"github.com/nalej/grpc-authx-go"
	"github.com/nalej/grpc-device-api-go"
	"github.com/nalej/grpc-device-manager-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"strings"
)

type Service struct {
	Configuration Config
}

// NewService creates a new system model service.
func NewService(conf Config) *Service {
	return &Service{
		conf,
	}
}

type Clients struct {
	deviceClient  grpc_device_manager_go.DevicesClient
	appClient     grpc_application_manager_go.ApplicationManagerClient
	authxClient   grpc_authx_go.AuthxClient
	latencyClient grpc_device_manager_go.LatencyClient
}

func (s *Service) GetClients() (*Clients, derrors.Error) {
	dmConn, err := grpc.Dial(s.Configuration.DeviceManagerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with the device manager")
	}
	deviceClient := grpc_device_manager_go.NewDevicesClient(dmConn)
	dClient := grpc_device_manager_go.NewLatencyClient(dmConn)

	appConn, err := grpc.Dial(s.Configuration.ApplicationsManagerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with the applications manager")
	}
	appClient := grpc_application_manager_go.NewApplicationManagerClient(appConn)

	authxConn, err := grpc.Dial(s.Configuration.AuthxAddress, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with the authx component")
	}
	authxClient := grpc_authx_go.NewAuthxClient(authxConn)

	return &Clients{
		deviceClient:  deviceClient,
		appClient:     appClient,
		authxClient:   authxClient,
		latencyClient: dClient}, nil
}

// Run the service, launch the REST service handler.
func (s *Service) Run() error {
	vErr := s.Configuration.Validate()
	if vErr != nil {
		log.Fatal().Str("err", vErr.DebugReport()).Msg("invalid configuration")
	}

	s.Configuration.Print()

	authConfig, authErr := s.Configuration.LoadAuthConfig()
	if authErr != nil {
		log.Fatal().Str("err", authErr.DebugReport()).Msg("cannot load authx config")
	}

	log.Info().Bool("AllowsAll", authConfig.AllowsAll).Int("permissions", len(authConfig.Permissions)).Msg("Auth config")

	go s.LaunchGRPC(authConfig)
	return s.LaunchHTTP()
}

// allowCORS allows Cross Origin Resource Sharing from any origin.
// Don't do this without consideration in production systems.
func (s *Service) allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}

func (s *Service) LaunchHTTP() error {
	addr := fmt.Sprintf(":%d", s.Configuration.HTTPPort)
	clientAddr := fmt.Sprintf(":%d", s.Configuration.Port)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	mux := runtime.NewServeMux()

	if err := grpc_device_api_go.RegisterDeviceHandlerFromEndpoint(context.Background(), mux, clientAddr, opts); err != nil {
		log.Fatal().Err(err).Msg("failed to start device handler")
	}
	if err := grpc_device_api_go.RegisterApplicationsHandlerFromEndpoint(context.Background(), mux, clientAddr, opts); err != nil {
		log.Fatal().Err(err).Msg("failed to start applications handler")
	}

	server := &http.Server{
		Addr:    addr,
		Handler: s.allowCORS(mux),
	}
	log.Info().Str("address", addr).Msg("HTTP Listening")
	return server.ListenAndServe()
}

func (s *Service) LaunchGRPC(authConfig *interceptor.AuthorizationConfig) error {
	clients, cErr := s.GetClients()
	if cErr != nil {
		log.Fatal().Str("err", cErr.DebugReport()).Msg("cannot generate clients")
		return cErr
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Configuration.Port))
	if err != nil {
		log.Fatal().Errs("failed to listen: %v", []error{err})
	}

	// Create handlers
	deviceManager := device.NewManager(s.Configuration.Threshold, clients.deviceClient, clients.latencyClient)
	deviceHandler := device.NewHandler(deviceManager)

	applicationsManager := applications.NewManager(clients.appClient)
	applicationsHandler := applications.NewHandler(applicationsManager)

	accessManager, aErr := devinterceptor.NewMngtSecretAccessWithClient(clients.authxClient, devinterceptor.DefaultCacheEntries)
	if err != nil {
		log.Fatal().Str("trace", aErr.DebugReport()).Msg("cannot create management secret access")
	}

	authxConfig := interceptor.NewConfig(authConfig, "", s.Configuration.AuthHeader)
	grpcServer := grpc.NewServer(interceptor.WithDeviceAuthxInterceptor(accessManager, authxConfig))
	grpc_device_api_go.RegisterDeviceServer(grpcServer, deviceHandler)
	grpc_device_api_go.RegisterApplicationsServer(grpcServer, applicationsHandler)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	log.Info().Int("port", s.Configuration.Port).Msg("Launching gRPC server")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Errs("failed to serve: %v", []error{err})
	}
	return nil
}
