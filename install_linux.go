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
Description={{.description}}{{if .network}}
After=network.target{{end}}

[Service]
ExecStart={{.exec}}

[Install]
WantedBy=default.target
`
)

func installPlatform(cfg *Config) error {

	// Compile the template
	t, err := template.New("").Parse(systemdUnitFile)
	if err != nil {
		return err
	}

	// Determine the full path to the executable
	p, err := os.Executable()
	if err != nil {
		return err
	}

	// Attempt to open the unit file
	f, err := os.Create(
		fmt.Sprintf(
			"/lib/systemd/system/%s.service",
			cfg.Name,
		),
	)
	if err != nil {
		return err
	}
	defer f.Close()

	// If there were arguments, append them to the path
	if len(cfg.Args) != 0 {
		p += " " + strings.Join(cfg.Args, " ")
	}

	// Write the template
	t.Execute(f, map[string]interface{}{
		"description": cfg.Description,
		"network":     cfg.RequiresNetwork,
		"exec":        p,
	})

	return nil
}
