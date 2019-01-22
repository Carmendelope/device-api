/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package server

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/nalej/authx/pkg/interceptor"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-utils/pkg/tools"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"strings"
)

type Service struct {
	Configuration Config
	Server        *tools.GenericGRPCServer
}

// NewService creates a new system model service.
func NewService(conf Config) *Service {
	return &Service{
		conf,
		tools.NewGenericGRPCServer(uint32(conf.Port)),
	}
}

type Clients struct {
}

func (s *Service) GetClients() (*Clients, derrors.Error) {
	dmConn, err := grpc.Dial(s.Configuration.DeviceManagerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with the device manager")
	}
	log.Debug().Interface("dm", dmConn).Msg("dmConn")

	return &Clients{}, nil
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
	// TODO Use clientAddr and opts
	addr := fmt.Sprintf(":%d", s.Configuration.HTTPPort)
	//clientAddr := fmt.Sprintf(":%d", s.Configuration.Port)
	//opts := []grpc.DialOption{grpc.WithInsecure()}
	mux := runtime.NewServeMux()

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
	log.Debug().Interface("clients", clients).Msg("clients")

	grpcServer := grpc.NewServer(interceptor.WithServerAuthxInterceptor(
		interceptor.NewConfig(authConfig, s.Configuration.AuthSecret, s.Configuration.AuthHeader)))

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	log.Info().Int("port", s.Configuration.Port).Msg("Launching gRPC server")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Errs("failed to serve: %v", []error{err})
	}
	return nil
}
