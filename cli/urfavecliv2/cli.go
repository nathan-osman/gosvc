package urfavecliv2

import (
	"os"
	"path/filepath"

	"github.com/nathan-osman/gosvc"
	"github.com/urfave/cli/v2"
)

// Config provides extra configuration information specific to urfave/cli.
type Config struct {

	// InstallFlags provides flags for the install command.
	InstallFlags []cli.Flag

	// InstallCustomizer provides a function for running custom code before
	// the install command. This might be used for adding command-line
	// arguments, for example.
	InstallCustomizer func(ctx *cli.Context)
}

// InstallCommand returns an "install" command suitable for use with
// github.com/urfave/cli.
func InstallCommand(i gosvc.Installer, cfg *Config) *cli.Command {
	if cfg == nil {
		cfg = &Config{}
	}
	return &cli.Command{
		Name:  "install",
		Usage: "install the application",
		Flags: cfg.InstallFlags,
		Action: func(ctx *cli.Context) error {
			if cfg.InstallCustomizer != nil {
				cfg.InstallCustomizer(ctx)
			}
			return i.Install()
		},
	}
}

// RemoveCommand returns a "remove" command suitable for use with
// github.com/urfave/cli.
func RemoveCommand(i gosvc.Installer) *cli.Command {
	return &cli.Command{
		Name:  "remove",
		Usage: "remove the application",
		Action: func(*cli.Context) error {
			return i.Remove()
		},
	}
}

// StartCommand returns a "start" command suitable for use with
// github.com/urfave/cli.
func StartCommand(s gosvc.Starter) *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "starts the application",
		Action: func(*cli.Context) error {
			return s.Start()
		},
	}
}

// StopCommand returns a "stop" command suitable for use with
// github.com/urfave/cli.
func StopCommand(s gosvc.Starter) *cli.Command {
	return &cli.Command{
		Name:  "stop",
		Usage: "stops the application",
		Action: func(*cli.Context) error {
			return s.Stop()
		},
	}
}

// App returns a *cli.App initialized with the commands from Commands().
func App(app *gosvc.Application, cfg *Config) (*cli.App, error) {
	e, err := os.Executable()
	if err != nil {
		return nil, err
	}
	p := app.Platform()
	return &cli.App{
		Name:  filepath.Base(e),
		Usage: app.Description,
		Commands: []*cli.Command{
			InstallCommand(p, cfg),
			RemoveCommand(p),
			StartCommand(p),
			StopCommand(p),
		},
		Action: func(ctx *cli.Context) error {
			return p.Run()
		},
	}, nil
}
