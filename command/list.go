package command

import (
	"github.com/urfave/cli"
)

var List = cli.Command{
	Name:      "list",
	Usage:     "List todos",
	ShortName: "l",
	Action:    list,
}

func list(c *cli.Context) error {
	if c.NArg() > 0 {
		cli.ShowCommandHelp(c, "l")
		return nil
	}

	doList()
	return nil
}

func doList() {
	db := getDB()
	_ = printAllTodo(db)
}
