//go:build windows

package gosvc

import (
	"fmt"

	"golang.org/x/sys/windows/svc"
)

// svc.Run() expects a Handler but we want to be able to simply pass a
// function; thankfully we can create a special type to achieve this

type handlerFunc func([]string, <-chan svc.ChangeRequest, chan<- svc.Status) (bool, uint32)

func (f handlerFunc) Execute(
	args []string,
	chChan <-chan svc.ChangeRequest,
	stChan chan<- svc.Status,
) (bool, uint32) {
	return f(args, chChan, stChan)
}

func runService(
	args []string,
	chChan <-chan svc.ChangeRequest,
	stChan chan<- svc.Status,
) (bool, uint32) {

	// Indicate that the service has been started
	stChan <- svc.Status{
		State:   svc.Running,
		Accepts: svc.AcceptStop | svc.AcceptShutdown,
	}

	// Respond to service requests
	for c := range chChan {
		switch c.Cmd {
		case svc.Interrogate:
			stChan <- c.CurrentStatus
		case svc.Stop, svc.Shutdown:
			stChan <- svc.Status{
				State: svc.StopPending,
			}
			return false, 0
		}
	}

	// This line should never be reached
	return false, 0
}

func runPlatform(cfg *Config) error {

	// Check if the application is being run as a Windows service; if not,
	// see if we are allowed to run it interactively
	if i, err := svc.IsWindowsService(); err != nil {
		return err
	} else if !i {
		if !cfg.RunInteractivelyOnWindows {
			return fmt.Errorf(
				"%s must be run as a Windows service",
				cfg.Name,
			)
		}
		runSignals()
		return nil
	}

	return svc.Run(cfg.Name, handlerFunc(runService))
}
