package gosvc

// Runner is the interface that wraps the Run method.
type Runner interface {
	Run() error
}

// Installer is the interface that wraps the Install and Remove methods.
type Installer interface {
	Install() error
	Remove() error
}

// Starter is the interface that wraps the Start and Stop methods.
type Starter interface {
	Start() error
	Stop() error
}

// Platform is the interface that combines Runner, Installer, and Starter.
type Platform interface {
	Runner
	Installer
	Starter
}
