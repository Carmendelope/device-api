/*
 * Copyright (C)  2019 Nalej - All Rights Reserved
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

func (m * Manager) RetrieveTargetApplications(filter *grpc_application_manager_go.ApplicationFilter) (*grpc_application_manager_go.TargetApplicationList, error) {
	return m.appsClient.RetrieveTargetApplications(context.Background(), filter)
}

func ( m * Manager) RetrieveEndpoints(request *grpc_application_manager_go.RetrieveEndpointsRequest) (*grpc_application_manager_go.ApplicationEndpoints, error){
	return m.appsClient.RetrieveEndpoints(context.Background(), request)
}
