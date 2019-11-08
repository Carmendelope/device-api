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

package entities

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-application-manager-go"
	"github.com/nalej/grpc-device-api-go"
	"github.com/nalej/grpc-device-controller-go"
	"github.com/nalej/grpc-device-go"
)

const emptyOrganizationId = "organization_id cannot be empty"
const emptyDeviceGroupId = "device_group_id cannot be empty"
const emptyDeviceGroupName = "device_group_name cannot be empty"
const emptyDeviceId = "device_id cannot be empty"
const emptyAppInstanceId = "app_instance_id cannot be empty"
const invalidMeasure = "Measure cannot be zero or less than zero"

func ValidDeviceId(deviceId *grpc_device_go.DeviceId) derrors.Error {
	if deviceId.OrganizationId == "" {
		return derrors.NewInvalidArgumentError(emptyOrganizationId)
	}
	if deviceId.DeviceGroupId == "" {
		return derrors.NewInvalidArgumentError(emptyDeviceGroupId)
	}
	if deviceId.DeviceId == "" {
		return derrors.NewInvalidArgumentError(emptyDeviceId)
	}
	return nil
}

func ValidApplicationFilter(filter *grpc_device_api_go.ApplicationFilter) derrors.Error {

	if filter.OrganizationId == "" {
		return derrors.NewInvalidArgumentError(emptyOrganizationId)
	}
	if filter.DeviceGroupName == "" {
		return derrors.NewInvalidArgumentError(emptyDeviceGroupName)
	}
	return nil
}

func ValidRetrieveEndpointsRequest(request *grpc_application_manager_go.RetrieveEndpointsRequest) derrors.Error {
	if request.OrganizationId == "" {
		return derrors.NewInvalidArgumentError(emptyOrganizationId)
	}
	if request.AppInstanceId == "" {
		return derrors.NewInvalidArgumentError(emptyAppInstanceId)
	}
	return nil
}

func ValidRegisterLatencyRequest(latency *grpc_device_controller_go.RegisterLatencyRequest) derrors.Error {
	if latency.OrganizationId == "" {
		return derrors.NewInvalidArgumentError(emptyOrganizationId)
	}
	if latency.DeviceGroupId == "" {
		return derrors.NewInvalidArgumentError(emptyDeviceGroupId)
	}
	if latency.DeviceId == "" {
		return derrors.NewInvalidArgumentError(emptyDeviceId)
	}
	if latency.Latency <= 0 {
		return derrors.NewInvalidArgumentError(invalidMeasure)
	}
	return nil
}
