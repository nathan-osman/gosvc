//go:build linux

package gosvc

type linuxPlatform struct {
	Runner
	SystemdService
}

func newPlatform(a *Application) Platform {
	l := &linuxPlatform{
		Runner: &SignalRunner{},
		SystemdService: SystemdService{
			Name:        a.Name,
			Description: a.Description,
			Args:        a.Args,
		},
	}
	if a.RequiresNetwork {
		l.Dependencies = append(l.Dependencies, "network.target")
	}
	return l
}
