/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package server

import (
	"github.com/nalej/authx/pkg/interceptor"
	"github.com/nalej/derrors"
	"github.com/nalej/device-api/version"
	"github.com/rs/zerolog/log"
)

type Config struct {
	// Port where the gRPC API service will listen requests.
	Port int
	// HTTPPort where the HTTP gRPC gateway will be listening.
	HTTPPort int
	// DeviceManagerAddress with the host:port to connect to Device Manager
	DeviceManagerAddress string
	// ApplicationsManagerAddress with the host:port to connect to the Applications manager.
	ApplicationsManagerAddress string
	// AuthHeader contains the name of the target header.
	AuthHeader string
	// AuthConfigPath contains the path of the file with the authentication configuration.
	AuthConfigPath string
	// AuthxAddress with the host:port to connect to the Authx manager.
	AuthxAddress string
	// Threshold in milliseconds by which it will be considered if a latency is acceptable or not
	Threshold int
}


func (conf *Config) Validate() derrors.Error {

	if conf.Port <= 0 || conf.HTTPPort <= 0 {
		return derrors.NewInvalidArgumentError("ports must be valid")
	}

	if conf.DeviceManagerAddress == "" {
		return derrors.NewInvalidArgumentError("deviceManager must be set")
	}

	if conf.ApplicationsManagerAddress == "" {
		return derrors.NewInvalidArgumentError("applicationsManagerAddress must be set")
	}

	if conf.AuthHeader == "" {
		return derrors.NewInvalidArgumentError("Authorization header must be set")
	}

	if conf.AuthConfigPath == "" {
		return derrors.NewInvalidArgumentError("authConfigPath must be set")
	}

	if conf.AuthxAddress == "" {
		return derrors.NewInvalidArgumentError("authxAddress must be set")
	}

	if conf.Threshold <= 0 {
		return derrors.NewInvalidArgumentError("Threshold must be valid")
	}

	return nil
}

// LoadAuthConfig loads the security configuration.
func (conf *Config) LoadAuthConfig() (*interceptor.AuthorizationConfig, derrors.Error) {
	return interceptor.LoadAuthorizationConfig(conf.AuthConfigPath)
}

func (conf *Config) Print() {
	log.Info().Str("app", version.AppVersion).Str("commit", version.Commit).Msg("Version")
	log.Info().Int("port", conf.Port).Msg("gRPC port")
	log.Info().Int("port", conf.HTTPPort).Msg("HTTP port")
	log.Info().Str("URL", conf.DeviceManagerAddress).Msg("Device Manager")
	log.Info().Str("URL", conf.ApplicationsManagerAddress).Msg("Applications Manager")
	log.Info().Str("URL", conf.AuthxAddress).Msg("Authx")
	log.Info().Str("header", conf.AuthHeader).Msg("Authorization")
	log.Info().Str("path", conf.AuthConfigPath).Msg("Permissions file")
	log.Info().Int("Threshold", conf.Threshold).Msg("Threshold in milliseconds")

}
