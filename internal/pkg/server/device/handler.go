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

package device

import (
	"context"
	"github.com/nalej/authx/pkg/interceptor"
	"github.com/nalej/derrors"
	"github.com/nalej/device-api/internal/pkg/entities"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-device-controller-go"
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
func (h *Handler) RetrieveDeviceLabels(ctx context.Context, deviceID *grpc_device_go.DeviceId) (*grpc_common_go.Labels, error){
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

// Ping is an operation triggered by the SDK
func (h *Handler) Ping(ctx context.Context, in *grpc_common_go.Empty) (*grpc_common_go.Success, error){
	return h.Manager.Ping()
}
// RegisterLatency Operation that is called by the SDK to inform the target cluster of the last latency measurement
func (h *Handler) RegisterLatency(ctx context.Context, latency *grpc_device_controller_go.RegisterLatencyRequest) (*grpc_device_controller_go.RegisterLatencyResult, error){

	rm, err := interceptor.GetDeviceRequestMetadata(ctx)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	if latency.OrganizationId != rm.OrganizationID {
		return nil, derrors.NewPermissionDeniedError("cannot access requested OrganizationID")
	}
	if latency.DeviceGroupId != rm.DeviceGroupID {
		return nil, derrors.NewPermissionDeniedError("cannot access requested DeviceGroupID")
	}
	if latency.DeviceId != rm.DeviceID {
		return nil, derrors.NewPermissionDeniedError("cannot access requested DeviceID")
	}
	vErr := entities.ValidRegisterLatencyRequest(latency)
	if vErr != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return h.Manager.RegisterLatency(latency)
}