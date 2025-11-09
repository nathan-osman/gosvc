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

Fill in an `Application` struct:

```golang
a := &gosvc.Application{
    Name:            "myservice",
    Description:     "My Service",
    Args:            []string{"-c", "config.json"},
    RequiresNetwork: true,
}
```

Convert this to a `Platform`:

```golang
p := a.Platform()
```

You can now use `p.Run()` in the main body of your application.

You also have access to `p.Install()`, `p.Remove()`, `p.Start()`, and `p.Stop()` for controlling the service. If you are using [github.com/urfave/cli/v2](https://github.com/urfave/cli), you can easily initialize your application with:

```golang
import "github.com/nathan-osman/gosvc/cli/urfavecliv2"

app, err := urfavecliv2.App(a)
if err != nil {
    panic(err)
}

if err := app.Run(os.Args); err != nil {
    panic(err)
}
```

Users of your application will be able to install the service with:

    yourapp.exe install
