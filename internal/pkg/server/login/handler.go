/*
 * Copyright (C)  2019 Nalej - All Rights Reserved
 */

package login

import (
	"context"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-authx-go"
	"github.com/nalej/grpc-utils/pkg/conversions"
)

// Handler structure for the node requests.
type Handler struct {
	Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler {
	return &Handler{manager}
}

// DeviceLogin checks the device API Key and whether the device is enabled to create a JWT token.
func (h*Handler) DeviceLogin(ctx context.Context, request *grpc_authx_go.DeviceLoginRequest) (*grpc_authx_go.LoginResponse, error){
	return nil, conversions.ToGRPCError(derrors.NewUnimplementedError("login not implemented"))
}
