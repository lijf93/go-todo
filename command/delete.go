package command

import (
	"fmt"
	"github.com/urfave/cli"
	"strconv"
)

var Delete = cli.Command{
	Name:      "deleteTodo",
	Usage:     "Delete a todo",
	UsageText: "go-todo deleteTodo [id] / go-todo del [id]",
	ShortName: "del",
	Action:    deleteTodo,
}

func deleteTodo(c *cli.Context) error {
	if c.NArg() < 1 {
		err := cli.ShowCommandHelp(c, "deleteTodo")
		if err != nil {
			return err
		}

		return nil
	}

	err := doDelete(c)

	if err != nil {
		checkDbErr(err)
		return err
	}

	return nil
}

func doDelete(c *cli.Context) error {
	db := getDB()
	id := c.Args()[0]
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("%s %s\n", red(IconBad), "Id should be a integer 😈")
		return err
	}
	res, err := findById(intId, db)
	if !res {
		fmt.Printf("%s %s\n", red(IconBad), fmt.Sprintf("Go-Todo id=%d not exist 😈", intId))
		_ = printAllTodo(db)
	} else {
		res, err := deleteById(intId, db)
		if res {
			fmt.Printf("%s %s\n", green(IconGood), fmt.Sprintf("Go-Todo deleteTodo %d success 🍻", intId))
			_ = printAllTodo(db)
		}

		if err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}

	return nil
}
