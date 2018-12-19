package command

import (
	"fmt"
	"github.com/urfave/cli"
	"strings"
)

var Add = cli.Command{
	Name:      "add",
	Usage:     "Add a todo",
	UsageText: "go-todo add [content] / go-todo a [content]",
	ShortName: "a",
	Action:    add,
}

func add(c *cli.Context) error {
	if c.NArg() < 1 {
		err := cli.ShowCommandHelp(c, "add")
		if err != nil {
			return err
		}

		return nil
	}

	err := doAdd(c)

	if err != nil {
		checkDbErr(err)
		return err
	}

	return nil
}

func doAdd(c *cli.Context) error {
	db := getDB()
	content := strings.Join(c.Args()[0:], " ")
	res, err := addDB(content, db)
	if res {
		fmt.Printf("%s %s\n", green(IconGood), "Go-Todo add success 🍻")
		_ = printAllTodo(db)
	}

	if err != nil {
		return err
	}

	return nil
}
