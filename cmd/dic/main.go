package main

import (
	"os"

	"github.com/ryoma123/dic"
	"github.com/urfave/cli"
)

var version = "0.2.0"

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "dic"
	app.Usage = "CLI tool for collecting domain information from multiple DNS servers."
	app.Version = version
	app.Author = "ryoma123"
	app.Email = "ryoma.ono.2661@gmail.com"
	app.Flags = dic.Flags
	app.Commands = dic.Commands
	app.Action = func(c *cli.Context) error {
		dic.RunCLI(c)
		return nil
	}
	return app
}
