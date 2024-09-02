//go:build windows

package gosvc

import (
	"os"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/eventlog"
	"golang.org/x/sys/windows/svc/mgr"
)

const (
	cmdSvcInstall = "install"
	cmdSvcRemove  = "remove"
	cmdSvcStart   = "start"
	cmdSvcStop    = "stop"
)

// WindowsService implements the Runner, Installer, and Starter interfaces for
// Windows services.
type WindowsService struct {
	Name         string
	DisplayName  string
	Description  string
	Args         []string
	Dependencies []string
}

// svc.Run() expects a Handler but we want to be able to simply pass a
// function; thankfully we can create a special type to achieve this

type serviceHandlerFunc func([]string, <-chan svc.ChangeRequest, chan<- svc.Status) (bool, uint32)

func (f serviceHandlerFunc) Execute(
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

func (w *WindowsService) serviceCommand(cmd string) error {

	// Connect to the service control manager
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	// Handle the install command separately
	if cmd == cmdSvcInstall {

		// Get the full path to the executable
		p, err := os.Executable()
		if err != nil {
			return err
		}

		// Create the service
		s, err := m.CreateService(
			w.Name,
			p,
			mgr.Config{
				StartType:    mgr.StartAutomatic,
				Dependencies: w.Dependencies,
				DisplayName:  w.DisplayName,
				Description:  w.Description,
			},
			w.Args...,
		)
		if err != nil {
			return err
		}
		defer s.Close()

		// Create the event log for the service
		return eventlog.InstallAsEventCreate(
			w.Name,
			eventlog.Error|eventlog.Warning|eventlog.Info,
		)
	}

	// Open the service
	s, err := m.OpenService(w.Name)
	if err != nil {
		return err
	}
	defer s.Close()

	// Run the appropriate command
	switch cmd {
	case cmdSvcRemove:
		if err := eventlog.Remove(w.Name); err != nil {
			return err
		}
		return s.Delete()
	case cmdSvcStart:
		return s.Start()
	case cmdSvcStop:
		_, err := s.Control(svc.Stop)
		return err
	}
	return nil
}

func (w *WindowsService) Run() error {
	return svc.Run(w.Name, serviceHandlerFunc(runService))
}

func (w *WindowsService) Install() error {
	return w.serviceCommand(cmdSvcInstall)
}

func (w *WindowsService) Remove() error {
	return w.serviceCommand(cmdSvcRemove)
}

func (w *WindowsService) Start() error {
	return w.serviceCommand(cmdSvcStart)
}

func (w *WindowsService) Stop() error {
	return w.serviceCommand(cmdSvcStop)
}
