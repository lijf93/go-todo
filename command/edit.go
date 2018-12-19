package command

import (
	"fmt"
	"github.com/urfave/cli"
	"strconv"
	"strings"
)

var Edit = cli.Command{
	Name:      "edit",
	Usage:     "Edit a todo",
	UsageText: "go-todo edit [id] [content] / go-todo e [id] [content]",
	ShortName: "e",
	Action:    edit,
}

func edit(c *cli.Context) error {
	if c.NArg() < 2 {
		err := cli.ShowCommandHelp(c, "edit")
		if err != nil {
			return err
		}

		return nil
	}

	err := doEdit(c)

	if err != nil {
		checkDbErr(err)
		return err
	}

	return nil
}

func doEdit(c *cli.Context) error {
	db := getDB()
	id := c.Args()[0]
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("%s %s %s\n", red(IconBad), "Id must be a integer", randomFailedEmoji())
		return err
	}
	res, err := findById(intId, db)
	if !res {
		fmt.Printf("%s %s %s\n", red(IconBad), fmt.Sprintf("Go-Todo id=%d not exist", intId), randomFailedEmoji())
		_ = printAllTodo(db)
	} else {
		newContent := strings.Trim(strings.Join(c.Args()[1:], " "), " ")
		if newContent == "" {
			fmt.Printf("%s %s %s\n", red(IconBad), "Go-Todo content is empty", randomFailedEmoji())
			return nil
		}
		res, err := editById(intId, newContent, db)
		if res {
			fmt.Printf("%s %s %s\n", green(IconGood), fmt.Sprintf("Go-Todo edit %d success", intId), randomSuccessEmoji())
			_ = printAllTodo(db)
		}

		return err
	}

	return err
}
