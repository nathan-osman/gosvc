package urfavecliv2

import (
	"github.com/nathan-osman/gosvc"
	"github.com/urfave/cli/v2"
)

// InstallCommand returns an "install" command suitable for use with
// github.com/urfave/cli.
func InstallCommand(i gosvc.Installer) *cli.Command {
	return &cli.Command{
		Name:  "install",
		Usage: "install the application",
		Action: func(*cli.Context) error {
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

// Commands returns a list of commands suitable for use with
// github.com/urfave/cli.
func Commands(p gosvc.Platform) []*cli.Command {
	return []*cli.Command{
		InstallCommand(p),
		RemoveCommand(p),
		StartCommand(p),
		StopCommand(p),
	}
}
