//go:build linux

package gosvc

import (
	"fmt"
	"os"
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

// SystemdInstaller implements the Installer interface.
type SystemdInstaller struct {
	Name         string
	Description  string
	Args         []string
	Dependencies []string
}

func (s *SystemdInstaller) unitFileName() string {
	return fmt.Sprintf(
		"/lib/systemd/system/%s.service",
		s.Name,
	)
}

func (s *SystemdInstaller) Install() error {

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

	// Attempt to open the unit file
	f, err := os.Create(s.unitFileName())
	if err != nil {
		return err
	}
	defer f.Close()

	// Write the template
	return t.Execute(f, s)
}

func (s *SystemdInstaller) Remove() error {
	return os.Remove(s.unitFileName())
}
