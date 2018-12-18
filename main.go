package main

import (
	"github.com/urfave/cli"
	"go-todo/command"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = logo
	app.Version = "v0.1.0"
	app.Usage = "A cli todo list app in Golang."
	app.Author = "lijf93"
	app.Email = "lijf93.hz@outlook.com"
	app.Commands = []cli.Command{
		command.Add,
		command.List,
		command.Done,
		command.Delete,
	}
	app.Run(os.Args)
}
