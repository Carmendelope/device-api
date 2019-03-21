/*
 * Copyright (C)  2019 Nalej - All Rights Reserved
 */

package entities

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-application-manager-go"
	"github.com/nalej/grpc-device-go"
)

const emptyOrganizationId = "organization_id cannot be empty"
const emptyDeviceGroupId = "device_group_id cannot be empty"
const emptyDeviceGroupName = "device_group_name cannot be empty"
const emptyDeviceId = "device_id cannot be empty"
const emptyAppInstanceId = "app_instance_id cannot be empty"

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

func ValidApplicationFilter(filter *grpc_application_manager_go.ApplicationFilter) derrors.Error {

	if filter.OrganizationId == "" {
		return derrors.NewInvalidArgumentError(emptyOrganizationId)
	}
	if filter.DeviceGroupName == "" {
		return derrors.NewInvalidArgumentError(emptyDeviceGroupName)
	}
	return nil
}

func ValidRetrieveEndpointsRequest (request *grpc_application_manager_go.RetrieveEndpointsRequest) derrors.Error {
	if request.OrganizationId == "" {
		return derrors.NewInvalidArgumentError(emptyOrganizationId)
	}
	if request.AppInstanceId == "" {
		return derrors.NewInvalidArgumentError(emptyAppInstanceId)
	}
	return nil
}