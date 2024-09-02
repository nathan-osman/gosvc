package gosvc

// Install prepares the application for running on the current platform. Note
// that this does not start the application.
//
// On Windows, this means registering the service with the service control
// manager.
//
// On Linux, this means creating a systemd unit file.
func Install(cfg *Config) error {
	return installPlatform(cfg)
}
