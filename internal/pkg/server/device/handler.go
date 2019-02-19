/*
 * Copyright (C)  2019 Nalej - All Rights Reserved
 */

package device

import (
	"context"
	"github.com/nalej/authx/pkg/interceptor"
	"github.com/nalej/derrors"
	"github.com/nalej/device-api/internal/pkg/entities"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-device-go"
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

// RetrieveDeviceLabels retrieves the list of labels associated with the current device.
func (h*Handler) RetrieveDeviceLabels(ctx context.Context, deviceID *grpc_device_go.DeviceId) (*grpc_common_go.Labels, error){
	rm, err := interceptor.GetDeviceRequestMetadata(ctx)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	if deviceID.OrganizationId != rm.OrganizationID {
		return nil, derrors.NewPermissionDeniedError("cannot access requested OrganizationID")
	}
	if deviceID.DeviceGroupId != rm.DeviceGroupID {
		return nil, derrors.NewPermissionDeniedError("cannot access requested DeviceGroupID")
	}
	if deviceID.DeviceId != rm.DeviceID {
		return nil, derrors.NewPermissionDeniedError("cannot access requested DeviceID")
	}
	err = entities.ValidDeviceId(deviceID)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return h.Manager.RetrieveDeviceLabels(deviceID)
}