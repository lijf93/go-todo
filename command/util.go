package command

import (
	"database/sql"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	// IconGood is success symbol
	IconGood = "✔"
	// IconBad is failed symbol
	IconBad = "✗"
)

var SuccessEmoji = []string{"🎉", "🍻", "😊", "🤡"}
var FailedEmoji = []string{"😈", "💥", "😭", "💣"}

const (
	CreateTable = `
		CREATE TABLE IF NOT EXISTS todo_list(
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            content TEXT NULL,
            is_done TINYINT(1) NULL,
            is_deleted TINYINT(1) NULL
        );`

	AddToTable = `INSERT into todo_list(content, is_done, is_deleted) values(?,?,?)`

	QueryNotDelete = `SELECT id, content, is_done FROM todo_list WHERE is_deleted != 1 ORDER BY id DESC, is_done ASC`

	FindById = `SELECT id, content, is_done FROM todo_list WHERE id = ?`

	UpdateDoneStatusById = `UPDATE todo_list SET is_done = ? WHERE id = ?`

	DeleteById = `UPDATE todo_list SET is_deleted = ? WHERE id = ?`

	EditById = `UPDATE todo_list SET content = ? WHERE id = ?`

	CountNotDelete = `SELECT count(*) FROM todo_list WHERE is_deleted != 1`
)

func red(str string) string {
	return color.RedString(str)
}

func green(str string) string {
	return color.GreenString(str)
}

func getDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./gotodo.db")
	checkDbErr(err)
	_, err = db.Exec(CreateTable)
	checkDbErr(err)

	return db
}

func addDB(content string, db *sql.DB) (bool, error) {
	stmt, err := db.Prepare(AddToTable)
	checkDbErr(err)
	res, err := stmt.Exec(content, 0, 0)
	id, err := res.LastInsertId()
	return id > 0, err
}

func findById(id int, db *sql.DB) (bool, error) {
	rows, err := db.Query(FindById, id)
	checkDbErr(err)
	res := rows.Next()
	_ = rows.Close()
	if res {
		return true, err
	} else {
		return false, err
	}
}

func doneById(id int, db *sql.DB) (bool, error) {
	stmt, err := db.Prepare(UpdateDoneStatusById)
	checkDbErr(err)
	res, err := stmt.Exec(1, id)
	rows, err := res.RowsAffected()
	return rows > 0, err
}

func undoneById(id int, db *sql.DB) (bool, error) {
	stmt, err := db.Prepare(UpdateDoneStatusById)
	checkDbErr(err)
	res, err := stmt.Exec(0, id)
	rows, err := res.RowsAffected()
	return rows > 0, err
}

func deleteById(id int, db *sql.DB) (bool, error) {
	stmt, err := db.Prepare(DeleteById)
	checkDbErr(err)
	res, err := stmt.Exec(1, id)
	rows, err := res.RowsAffected()
	return rows > 0, err
}

func editById(id int, content string, db *sql.DB) (bool, error) {
	stmt, err := db.Prepare(EditById)
	checkDbErr(err)
	res, err := stmt.Exec(content, id)
	rows, err := res.RowsAffected()
	return rows > 0, err
}

func printAllTodo(db *sql.DB) error {
	var count int64
	err := db.QueryRow(CountNotDelete).Scan(&count)
	if count == 0 {
		fmt.Printf("%s %s %s\n", green(IconGood), "You have already done all your todos", randomSuccessEmoji())
		if err != nil {
			return err
		}
	}
	rows, err := db.Query(QueryNotDelete)
	checkDbErr(err)
	var buf []string
	for rows.Next() {
		var prefix string
		err = rows.Scan(&Id, &Content, &IsDone)
		checkErr(err)
		if IsDone == 1 {
			prefix = IconGood
		} else {
			prefix = " "
		}
		buf = append(buf, fmt.Sprintf("%s[%s] %s", color.GreenString(prefix), color.CyanString(strconv.Itoa(Id)), Content))
	}
	err = rows.Close()
	fmt.Printf("%v\n", strings.Join(buf, "\n"))
	if err != nil {
		return err
	}

	return nil
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("%s %s %s\n", red(IconBad), "SYSTEM ERROR", randomFailedEmoji())
	}
}

func checkDbErr(err error) {
	if err != nil {
		fmt.Printf("%s %s %s\n", red(IconBad), "DB ERROR", randomFailedEmoji())
	}
}

func randomFailedEmoji() string {
	rand.Seed(time.Now().UnixNano())
	return FailedEmoji[rand.Intn(len(FailedEmoji))]
}

func randomSuccessEmoji() string {
	rand.Seed(time.Now().UnixNano())
	return SuccessEmoji[rand.Intn(len(SuccessEmoji))]
}
