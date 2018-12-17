package main

const (
	CREATE_TABLE = `CREATE TABLE IF NOT EXISTS go_todo(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        content VARCHAR(4096) NULL,
        status VARCHAR(64) NULL,
        todo_date DATE NULL
    );`
)
