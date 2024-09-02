//go:build !windows

package gosvc

func runPlatform(*Config) error {
	runSignals()
	return nil
}
