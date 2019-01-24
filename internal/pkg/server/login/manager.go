/*
 * Copyright (C)  2019 Nalej - All Rights Reserved
 */

package login

import (
	"github.com/nalej/grpc-authx-go"
)

// Manager structure with the required clients for node operations.
type Manager struct {
	authxClient grpc_authx_go.AuthxClient
}

// NewManager creates a Manager using a set of clients.
func NewManager(authxClient grpc_authx_go.AuthxClient) Manager {
	return Manager{
		authxClient: authxClient,
	}
}