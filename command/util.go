package command

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

const (
	// IconGood is success symbol
	IconGood = "✔"
	// IconBad is failed symbol
	IconBad = "✗"
)

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

	DoneById = `UPDATE todo_list SET is_done = ? WHERE id = ?`

	DeleteById = `UPDATE todo_list SET is_deleted = ? WHERE id = ?`
)

func red(str string) string {
	return fmt.Sprintf("\x1b[0;31m%s\x1b[0m", str)
}

func green(str string) string {
	return fmt.Sprintf("\x1b[0;32m%s\x1b[0m", str)
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
	stmt, err := db.Prepare(FindById)
	checkDbErr(err)
	_, err = stmt.Exec(id)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func doneById(id int, db *sql.DB) (bool, error) {
	stmt, err := db.Prepare(DoneById)
	checkDbErr(err)
	res, err := stmt.Exec(1, id)
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

func printAllTodo(db *sql.DB) {
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
		buf = append(buf, fmt.Sprintf("[%s][%d] %s", green(prefix), Id, Content))
	}
	fmt.Printf("%v\n", strings.Join(buf, "\n"))
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("%s %s\n", red(IconBad), "SYSTEM ERROR")
		panic(err)
	}
}

func checkDbErr(err error) {
	if err != nil {
		fmt.Printf("%s %s\n", red(IconBad), "DB ERROR")
		panic(err)
	}
}
