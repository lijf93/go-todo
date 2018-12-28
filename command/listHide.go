package command

import (
	"github.com/urfave/cli"
)

var ListHide = cli.Command{
	Name:      "listhide",
	Usage:     "List hide todos",
	UsageText: "go-todo listhide / go-todo lh",
	ShortName: "lh",
	Action:    listHide,
}

func listHide(c *cli.Context) error {
	if c.NArg() > 0 {
		cli.ShowCommandHelp(c, "lh")
		return nil
	}

	doListHide()
	return nil
}

func doListHide() {
	db := getDB()
	_ = printAllHideTodo(db)
}
