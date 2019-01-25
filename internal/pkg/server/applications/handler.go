/*
 * Copyright (C)  2019 Nalej - All Rights Reserved
 */

package applications

import (
	"context"
	"github.com/nalej/authx/pkg/interceptor"
	"github.com/nalej/derrors"
	"github.com/nalej/device-api/internal/pkg/entities"
	"github.com/nalej/grpc-application-manager-go"
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
func (h*Handler) RetrieveTargetApplications(ctx context.Context, filter *grpc_application_manager_go.ApplicationFilter) (*grpc_application_manager_go.TargetApplications, error){
	rm, err := interceptor.GetDeviceRequestMetadata(ctx)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	if filter.OrganizationId != rm.OrganizationID {
		return nil, derrors.NewPermissionDeniedError("cannot access requested OrganizationID")
	}
	if filter.DeviceGroupId != rm.DeviceGroupID {
		return nil, derrors.NewPermissionDeniedError("cannot access requested DeviceGroupID")
	}
	err = entities.ValidApplicationFilter(filter)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return h.Manager.RetrieveTargetApplications(filter)
}

// RetrieveTargetApplications retrieves the list of target applications that accept data from the device.
func (h*Handler) RetrieveEndpoints(ctx context.Context, request *grpc_application_manager_go.RetrieveEndpointsRequest) (*grpc_device_api_go.ApplicationEndpoints, error){
	rm, err := interceptor.GetDeviceRequestMetadata(ctx)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	if request.OrganizationId != rm.OrganizationID {
		return nil, derrors.NewPermissionDeniedError("cannot access requested OrganizationID")
	}
	err = entities.ValidRetrieveEndpointsRequest(request)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return h.Manager.RetrieveEndpoints(request)
}
