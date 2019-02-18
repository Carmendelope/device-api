/*
 * Copyright (C)  2019 Nalej - All Rights Reserved
 */

package device

import (
	"context"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-device-go"
	"github.com/nalej/grpc-device-manager-go"
	"time"
)

const DeviceClientTimeout = time.Second * 5

// Manager structure with the required clients for node operations.
type Manager struct {
	deviceClient grpc_device_manager_go.DevicesClient
}

// NewManager creates a Manager using a set of clients.
func NewManager(deviceClient grpc_device_manager_go.DevicesClient) Manager {
	return Manager{
		deviceClient: deviceClient,
	}
}

// RetrieveDeviceLabels retrieves the list of labels associated with the current device.
func (m*Manager) RetrieveDeviceLabels(deviceId *grpc_device_go.DeviceId) (*grpc_common_go.Labels, error){
	ctx, cancel := context.WithTimeout(context.Background(), DeviceClientTimeout)
	defer cancel()
	retrieved, err := m.deviceClient.GetDevice(ctx, deviceId)
	if err != nil{
		return nil, err
	}
	return &grpc_common_go.Labels{
		Labels: retrieved.Labels,
	}, nil
}