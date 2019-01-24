/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package commands

import (
	"github.com/nalej/device-api/internal/pkg/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var config = server.Config{}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Launch the server API",
	Long:  `Launch the server API`,
	Run: func(cmd *cobra.Command, args []string) {
		SetupLogging()
		log.Info().Msg("Launching API!")
		server := server.NewService(config)
		server.Run()
	},
}

func init() {
	runCmd.Flags().IntVar(&config.Port, "port", 6000, "Port to launch the Device gRPC API")
	runCmd.Flags().IntVar(&config.HTTPPort, "httpPort", 6001, "Port to launch the Device HTTP API")
	runCmd.PersistentFlags().StringVar(&config.DeviceManagerAddress, "deviceManager", "localhost:6010",
		"Device Manager address (host:port)")
	runCmd.PersistentFlags().StringVar(&config.AuthxAddress, "authxAddress", "localhost:8810",
		"Authx address (host:port)")
	runCmd.PersistentFlags().StringVar(&config.ApplicationsManagerAddress, "applicationsManagerAddress", "localhost:8910",
		"Applications Manager address (host:port)")
	runCmd.PersistentFlags().StringVar(&config.AuthHeader, "authHeader", "", "Authorization Header")
	runCmd.PersistentFlags().StringVar(&config.AuthSecret, "authSecret", "", "Authorization secret")
	runCmd.PersistentFlags().StringVar(&config.AuthConfigPath, "authConfigPath", "", "Authorization config path")
	rootCmd.AddCommand(runCmd)
}
