//go:build windows

package gosvc

import (
	"os"

	"golang.org/x/sys/windows/svc/eventlog"
	"golang.org/x/sys/windows/svc/mgr"
)

func installPlatform(cfg *Config) error {

	// Connect to the service control manager
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	// Get the path to the executable
	p, err := os.Executable()
	if err != nil {
		return err
	}

	// Determine dependencies
	dependencies := []string{}
	if cfg.RequiresNetwork {
		dependencies = append(dependencies, "nsi")
	}

	// Create the service
	s, err := m.CreateService(
		cfg.Name,
		p,
		mgr.Config{
			StartType:    mgr.StartAutomatic,
			Dependencies: dependencies,
			DisplayName:  cfg.DisplayName,
			Description:  cfg.Description,
		},
		cfg.Args...,
	)
	if err != nil {
		return err
	}
	defer s.Close()

	// Log the installation
	return eventlog.InstallAsEventCreate(
		cfg.Name,
		eventlog.Error|eventlog.Warning|eventlog.Info,
	)
}
