package command

import (
	"fmt"
	"github.com/urfave/cli"
	"strconv"
)

var Delete = cli.Command{
	Name:      "delete",
	Usage:     "Delete a todo",
	ShortName: "del",
	Action:    delete,
}

func delete(c *cli.Context) error {
	if c.NArg() < 1 {
		err := cli.ShowCommandHelp(c, "delete")
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
		fmt.Printf("%s %s\n", red(IconBad), "ID SHOULD BE A INTEGER")
		return err
	}
	res, err := findById(intId, db)
	if !res {
		fmt.Printf("%s %s\n", red(IconBad), fmt.Sprintf("Go-Todo id = %d not exist", intId))
		printAllTodo(db)
	} else {
		res, err := deleteById(intId, db)
		if res {
			fmt.Printf("%s %s\n", green(IconGood), fmt.Sprintf("Go-Todo delete %d success", intId))
			printAllTodo(db)
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
