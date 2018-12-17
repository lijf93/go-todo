package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = logo
	app.Version = "v0.1.0"
	app.Usage = "A cli todo list app in Golang."
	app.Author = "lijf93"
	app.Email = "lijf93.hz@outlook.com"
	app.Commands = []cli.Command{}
	app.Run(os.Args)
}
