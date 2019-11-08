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

package applications

import (
	"context"
	"github.com/nalej/grpc-application-manager-go"
)

// Manager structure with the required clients for node operations.
type Manager struct {
	appsClient grpc_application_manager_go.ApplicationManagerClient
}

// NewManager creates a Manager using a set of clients.
func NewManager(appsClient grpc_application_manager_go.ApplicationManagerClient) Manager {
	return Manager{
		appsClient: appsClient,
	}
}

func (m *Manager) RetrieveTargetApplications(filter *grpc_application_manager_go.ApplicationFilter) (*grpc_application_manager_go.TargetApplicationList, error) {
	return m.appsClient.RetrieveTargetApplications(context.Background(), filter)
}

func (m *Manager) RetrieveEndpoints(request *grpc_application_manager_go.RetrieveEndpointsRequest) (*grpc_application_manager_go.ApplicationEndpoints, error) {
	return m.appsClient.RetrieveEndpoints(context.Background(), request)
}
