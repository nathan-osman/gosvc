package gosvc

import (
	"os"
	"os/signal"
	"syscall"
)

// Run executes the application according to the current platform. This
// function will not return until the application needs to be closed.
//
// On Windows, this means checking if it is being run as a Windows service. If
// so, the service control manager is used. If not, signals are used.
//
// On all other platforms, signals (SIGTERM, SIGINT) are used to determine
// when to close.
func Run(cfg *Config) error {
	return runPlatform(cfg)
}

func runSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
