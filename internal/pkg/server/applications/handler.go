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
func (h *Handler) RetrieveTargetApplications(ctx context.Context, filter *grpc_device_api_go.ApplicationFilter) (*grpc_application_manager_go.TargetApplicationList, error) {
	rm, err := interceptor.GetDeviceRequestMetadata(ctx)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	if filter.OrganizationId != rm.OrganizationID {
		return nil, derrors.NewPermissionDeniedError("cannot access requested OrganizationID")
	}
	err = entities.ValidApplicationFilter(filter)

	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}

	return h.Manager.RetrieveTargetApplications(&grpc_application_manager_go.ApplicationFilter{
		OrganizationId:  filter.OrganizationId,
		DeviceGroupId:   rm.DeviceGroupID,
		DeviceGroupName: filter.DeviceGroupName,
		MatchLabels:     filter.MatchLabels,
	})
}

// RetrieveTargetApplications retrieves the list of target applications that accept data from the device.
func (h *Handler) RetrieveEndpoints(ctx context.Context, request *grpc_application_manager_go.RetrieveEndpointsRequest) (*grpc_application_manager_go.ApplicationEndpoints, error) {
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
