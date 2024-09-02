package gosvc

import (
	"os"
	"os/signal"
	"syscall"
)

// SignalRunner implements the Runner interface and waits until the SIGINT or
// SIGTERM signals are received before quitting.
type SignalRunner struct{}

func (s *SignalRunner) Run() error {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	return nil
}
