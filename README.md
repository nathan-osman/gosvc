## gosvc

[![Go Reference](https://pkg.go.dev/badge/github.com/nathan-osman/gosvc.svg)](https://pkg.go.dev/github.com/nathan-osman/gosvc)
[![MIT License](https://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](https://opensource.org/licenses/MIT)

This package helps you focus on writing your application, not the boilerplate necessary to run it on each platform. By using this package, you get functions that:

- Install and run the application as a Windows service
- Create a systemd unit file for easily running the application on Linux

And it's really easy to use!

### Usage

Begin by importing the package:

```golang
import "github.com/nathan-osman/gosvc"
```

#### Windows

On Windows, create an instance of the `WindowsService` type:

```golang
s := &gosvc.WindowsService{
    Name:         "myservice",
    DisplayName:  "My Service",
    Description:  "Does the things",
    Args:         []string{"-c", "config.json"},
    Dependencies: []string{"nsi"},
}
```

You can then call the `Run()`, `Install()`, etc. methods directly.

#### Linux

On Linux, create an instance of the `SignalRunner` and `SystemdInstaller` types:

```golang
r := &gosvc.SignalRunner{}
i := &gosvc.SystemdInstaller{
    Name:         "myservice",
    Description:  "My Service",
    Args:         []string{"-c", "config.json"},
    Dependencies: []string{"network.target"},
}
```
