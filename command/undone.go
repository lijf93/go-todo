package command

import (
	"fmt"
	"github.com/urfave/cli"
	"strconv"
)

var Undone = cli.Command{
	Name:      "undone",
	Usage:     "Undone a todo",
	UsageText: "go-todo undone [id] / go-todo un [id]",
	ShortName: "un",
	Action:    undone,
}

func undone(c *cli.Context) error {
	if c.NArg() < 1 {
		err := cli.ShowCommandHelp(c, "undone")
		if err != nil {
			return err
		}

		return nil
	}

	err := doUndone(c)

	if err != nil {
		checkDbErr(err)
		return err
	}

	return nil
}

func doUndone(c *cli.Context) error {
	db := getDB()
	id := c.Args()[0]
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("%s %s %s\n", red(IconBad), "Id must be a integer", randomFailedEmoji())
		return nil
	}
	res, err := findById(intId, db)
	if !res {
		fmt.Printf("%s %s %s\n", red(IconBad), fmt.Sprintf("Go-Todo id=%d not exist", intId), randomFailedEmoji())
		_ = printAllTodo(db)
	} else {
		res, err := undoneById(intId, db)
		if res {
			fmt.Printf("%s %s %s\n", green(IconGood), fmt.Sprintf("Go-Todo undone %d success", intId), randomSuccessEmoji())
			_ = printAllTodo(db)
		}

		return err
	}

	return err
}
