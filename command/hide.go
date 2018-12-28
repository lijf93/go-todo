package command

import (
	"fmt"
	"github.com/urfave/cli"
	"strconv"
)

var Hide = cli.Command{
	Name:      "hide",
	Usage:     "Hide a todo",
	UsageText: "go-todo hide [id] / go-todo hd [id]",
	ShortName: "hd",
	Action:    hide,
}

func hide(c *cli.Context) error {
	if c.NArg() < 1 {
		err := cli.ShowCommandHelp(c, "hide")
		if err != nil {
			return err
		}

		return nil
	}

	err := doHide(c)

	if err != nil {
		checkDbErr(err)
		return err
	}

	return nil
}

func doHide(c *cli.Context) error {
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
		res, err := hideById(intId, db)
		if res {
			fmt.Printf("%s %s %s\n", green(IconGood), fmt.Sprintf("Go-Todo hide %d success", intId), randomSuccessEmoji())
			_ = printAllTodo(db)
		}

		return err
	}

	return err
}
