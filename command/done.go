package command

import (
	"fmt"
	"github.com/urfave/cli"
	"strconv"
)

var Done = cli.Command{
	Name:      "done",
	Usage:     "Done a todo",
	UsageText: "go-todo done [id] / go-todo do [id]",
	ShortName: "do",
	Action:    done,
}

func done(c *cli.Context) error {
	if c.NArg() < 1 {
		err := cli.ShowCommandHelp(c, "done")
		if err != nil {
			return err
		}

		return nil
	}

	err := doDone(c)

	if err != nil {
		checkDbErr(err)
		return err
	}

	return nil
}

func doDone(c *cli.Context) error {
	db := getDB()
	id := c.Args()[0]
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("%s %s\n", red(IconBad), "Id must be a integer 😈")
		return nil
	}
	res, err := findById(intId, db)
	if !res {
		fmt.Printf("%s %s\n", red(IconBad), fmt.Sprintf("Go-Todo id=%d not exist 😈", intId))
		_ = printAllTodo(db)
	} else {
		res, err := doneById(intId, db)
		if res {
			fmt.Printf("%s %s\n", green(IconGood), fmt.Sprintf("Go-Todo done %d success 🍻", intId))
			_ = printAllTodo(db)
		}

		return err
	}

	return err
}
