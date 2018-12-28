package command

import (
	"fmt"
	"github.com/urfave/cli"
	"strconv"
)

var Unhide = cli.Command{
	Name:      "unhide",
	Usage:     "Unhide a todo",
	UsageText: "go-todo unhide [id] / go-todo ud [id]",
	ShortName: "uh",
	Action:    unhide,
}

func unhide(c *cli.Context) error {
	if c.NArg() < 1 {
		err := cli.ShowCommandHelp(c, "unhide")
		if err != nil {
			return err
		}

		return nil
	}

	err := doUnhide(c)

	if err != nil {
		checkDbErr(err)
		return err
	}

	return nil
}

func doUnhide(c *cli.Context) error {
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
		res, err := unHideById(intId, db)
		if res {
			fmt.Printf("%s %s %s\n", green(IconGood), fmt.Sprintf("Go-Todo unhide %d success", intId), randomSuccessEmoji())
			_ = printAllHideTodo(db)
		}

		return err
	}

	return err
}
