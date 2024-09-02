## gosvc

[![Go Reference](https://pkg.go.dev/badge/github.com/nathan-osman/gosvc.svg)](https://pkg.go.dev/github.com/nathan-osman/gosvc)
[![MIT License](https://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](https://opensource.org/licenses/MIT)

This package helps you focus on writing your application, not the boilerplate necessary to run it on each platform. By using this package, you get functions that:

- Install and run the application as a Windows service
- Create a systemd unit file for easily running the application on Linux

And it's really easy to use!

### Usage

Begin by importing the package and creating a simple config:

```golang
import "github.com/nathan-osman/gosvc"

cfg := gosvc.Config{
    Name:            "myservice",
    DisplayName:     "My Service",
    Description:     "Does various things"
    RequiresNetwork: true,
}
```

To install the application (perhaps in response to a "myapp install" invocation), simply run:

```golang
gosvc.Install(&cfg)
```

When you are ready to run the application, use:

```golang
gosvc.Run(&cfg)
```

That's it!
