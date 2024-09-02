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
	Name            string
	Description     string
	Args            []string
	RequiresNetwork bool
}

// Platform selects an appropriate Runner, Installer, and Starter for the current
// platform and returns a Platform for it.
func (a *Application) Platform() Platform {
	return newPlatform(a)
}
