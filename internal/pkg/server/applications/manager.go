/*
 * Copyright (C)  2019 Nalej - All Rights Reserved
 */

package applications

import (
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