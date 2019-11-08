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

 /*
  The test are pending,
  RetrieveTargetApplications and RetrieveEndpoints have been tested in the application manager service
  */
package applications

import (
	"context"
	"github.com/nalej/device-api/internal/pkg/utils"
	"github.com/nalej/grpc-application-manager-go"
	"github.com/nalej/grpc-device-api-go"
	"github.com/nalej/grpc-utils/pkg/test"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"os"
)

var _ = ginkgo.Describe("Device API service", func() {

	if ! utils.RunIntegrationTests() {
		log.Warn().Msg("Integration tests are skipped")
		return
	}

	var (
		appMngAddress = os.Getenv("IT_APPMNG_ADDRESS")
	)

	if appMngAddress == "" {
		ginkgo.Fail("missing environment variables")
	}

	// gRPC server
	var server *grpc.Server
	// grpc test listener
	var listener *bufconn.Listener
	// client

	var appMngClient grpc_application_manager_go.ApplicationManagerClient
	var client grpc_device_api_go.ApplicationsClient

	ginkgo.BeforeSuite(func() {
		listener = test.GetDefaultListener()
		server = grpc.NewServer()

		appConn := utils.GetConnection(appMngAddress)
		appMngClient = grpc_application_manager_go.NewApplicationManagerClient(appConn)

		test.LaunchServer(server, listener)

		// Register the service
		manager := NewManager(appMngClient)
		handler := NewHandler(manager)
		grpc_device_api_go.RegisterApplicationsServer(server, handler)

		conn, err := test.GetConn(*listener)
		gomega.Expect(err).Should(gomega.Succeed())

		client = grpc_device_api_go.NewApplicationsClient(conn)
	})

	ginkgo.AfterSuite(func() {
		server.Stop()
		listener.Close()
	})

	ginkgo.PIt("Should be able to retrieve target applications", func(){
		// TODO: prepare the system and check the results
		client.RetrieveTargetApplications(context.Background(), nil)
	})
	ginkgo.PIt("Should be able to retrieve target applications without labels filering", func(){
	})
	ginkgo.PIt("Should not be able to retrieve target applications of a non existing organization", func(){
	})
	ginkgo.PIt("Should be able to retrieve an empty list (no match deviceGroupId)", func(){
	})
	ginkgo.PIt("Should be able to retrieve an empty list (no match labels)", func(){
	})
	ginkgo.PIt("Should be able to retrieve endpoints", func(){
	})
	ginkgo.PIt("Should be able to retrieve an empty endpoints (service is waiting)", func(){
	})
})