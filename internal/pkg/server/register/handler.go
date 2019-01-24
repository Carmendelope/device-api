/*
 * Copyright (C)  2019 Nalej - All Rights Reserved
 */

package register

import (
	"context"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-device-api-go"
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

// RegisterDevice registers a new device in a given device group.
func (h*Handler) RegisterDevice(ctx context.Context, request *grpc_device_api_go.RegisterDeviceRequest) (*grpc_device_api_go.RegisterResponse, error){
	return nil, conversions.ToGRPCError(derrors.NewUnimplementedError("register not implemented"))
}