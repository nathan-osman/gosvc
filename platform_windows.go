//go:build windows

package gosvc

func newPlatform(a *Application) Platform {
	w := &WindowsService{
		Name:        a.Name,
		Description: a.Description,
		Args:        a.Args,
	}
	if a.RequiresNetwork {
		w.Dependencies = []string{"nsi"}
	}
	return w
}
