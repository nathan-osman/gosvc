package gosvc

// Platform is the interface that combines Runner, Installer, and Starter.
type Platform interface {
	Runner
	Installer
	Starter
}

// Application provides the information necessary to run the application on
// the current platform.
type Application struct {

	// Name is the human-readable name of the service.
	Name string

	// Description is a brief description of the service and its purpose.
	Description string

	// Args indicates the command-line arguments needed to launch the
	// application.
	Args []string

	// RequiresNetwork indicates that this application requires network access
	// and should only be started after network services are available.
	RequiresNetwork bool
}

// Platform selects an appropriate Runner, Installer, and Starter for the current
// platform and returns a Platform for it.
func (a *Application) Platform() Platform {
	return newPlatform(a)
}
