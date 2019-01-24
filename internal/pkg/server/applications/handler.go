/*
 * Copyright (C)  2019 Nalej - All Rights Reserved
 */

package applications

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

// RetrieveTargetApplications retrieves the list of target applications that accept data from the device.
func (h*Handler) RetrieveTargetApplications(ctx context.Context, filter *grpc_device_api_go.ApplicationFilter) (*grpc_device_api_go.TargetApplications, error){
	return nil, conversions.ToGRPCError(derrors.NewUnimplementedError("retrieve target applications not implemented"))
}

// RetrieveTargetApplications retrieves the list of target applications that accept data from the device.
func (h*Handler) RetrieveEndpoints(ctx context.Context, request *grpc_device_api_go.RetrieveEndpointsRequest) (*grpc_device_api_go.ApplicationEndpoints, error){
	return nil, conversions.ToGRPCError(derrors.NewUnimplementedError("retrieve endpoints not implemented"))
}
