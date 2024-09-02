//go:build linux

package gosvc

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

const (
	systemdUnitFile = `[Unit]
Description={{.Description}}{{if .Dependencies}}
Requires={{join .Dependencies}}
After={{join .Dependencies}}{{end}}

[Service]
ExecStart={{exec}}{{if .Args}} {{join .Args}}{{end}}

[Install]
WantedBy=default.target
`
)

// SystemdService implements the Installer interface.
type SystemdService struct {
	Name         string
	Description  string
	Args         []string
	Dependencies []string
}

func (s *SystemdService) unitFileName() string {
	return fmt.Sprintf(
		"/lib/systemd/system/%s.service",
		s.Name,
	)
}

func runSystemdCommand(args ...string) error {
	cmd := exec.Command("systemctl", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (s *SystemdService) Install() error {

	// Determine the full path to the executable
	p, err := os.Executable()
	if err != nil {
		return err
	}

	// Compile the template
	t, err := template.New("").Funcs(template.FuncMap{
		"exec": func() string { return p },
		"join": func(v []string) string { return strings.Join(v, " ") },
	}).Parse(systemdUnitFile)
	if err != nil {
		return err
	}

	// Attempt to create the unit file
	f, err := os.Create(s.unitFileName())
	if err != nil {
		return err
	}
	defer f.Close()

	// Write the template
	if err := t.Execute(f, s); err != nil {
		return err
	}

	// Enable the service
	return runSystemdCommand("enable", s.Name)
}

func (s *SystemdService) Remove() error {

	// Disable the service
	if err := runSystemdCommand("disable", s.Name); err != nil {
		return err
	}

	// Remove the unit file
	return os.Remove(s.unitFileName())
}

func (s *SystemdService) Start() error {
	return runSystemdCommand("start", s.Name)
}

func (s *SystemdService) Stop() error {
	return runSystemdCommand("stop", s.Name)
}
