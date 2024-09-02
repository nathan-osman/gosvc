package gosvc

// Config is used to provide the runtime with information about the
// application and customize behavior if desired.
type Config struct {

	// Name is the name of the service as it should be displayed to the
	// system. On Windows, for example, this will be the name displayed in the
	// service control manager.
	//
	// Do not change this value in future updates for your application or
	// future actions (such as removal) will stop working.
	Name string

	// DisplayName is the human-friendly name for the service.
	DisplayName string

	// Description is a brief summary of the application's function / purpose.
	Description string

	// Args is a list of arguments that should be passed to the application
	// when it is started.
	Args []string

	// RequiresNetwork indicates that the application should only be started
	// once the network is up and running.
	RequiresNetwork bool

	// RunInteractivelyOnWindows indicates whether the application can be run
	// interactively or if it must be run as a Windows service. True will
	// cause an error to be returned if the application is run interactively.
	RunInteractivelyOnWindows bool
}
