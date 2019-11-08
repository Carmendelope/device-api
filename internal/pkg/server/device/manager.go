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
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-device-controller-go"
	"github.com/nalej/grpc-device-go"
	"github.com/nalej/grpc-device-manager-go"
	"time"
)

const DeviceClientTimeout = time.Second * 5

// Manager structure with the required clients for node operations.
type Manager struct {
	Threshold int
	deviceClient grpc_device_manager_go.DevicesClient
	latencyClient grpc_device_manager_go.LatencyClient
}

// NewManager creates a Manager using a set of clients.
func NewManager(threshold int, deviceClient grpc_device_manager_go.DevicesClient, latencyClient grpc_device_manager_go.LatencyClient) Manager {
	return Manager{
		Threshold:threshold,
		deviceClient: deviceClient,
		latencyClient:latencyClient,
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

func (m * Manager) Ping () (*grpc_common_go.Success, error) {
	return &grpc_common_go.Success{}, nil
}


func (m *Manager) RegisterLatency(latency *grpc_device_controller_go.RegisterLatencyRequest) (*grpc_device_controller_go.RegisterLatencyResult, error){
	result := grpc_device_controller_go.RegisterResult_OK
	if int(latency.Latency) > m.Threshold {
		result = grpc_device_controller_go.RegisterResult_LATENCY_CHECK_REQUIRED
	}

	request :=  &grpc_device_controller_go.RegisterLatencyRequest{
		OrganizationId: latency.OrganizationId,
		DeviceGroupId: latency.DeviceGroupId,
		DeviceId: latency.DeviceId,
		Latency: latency.Latency,
	}

	_, err := m.latencyClient.RegisterLatency(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return &grpc_device_controller_go.RegisterLatencyResult{
		Result: result,
	}, nil


}
