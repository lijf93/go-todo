package main

type Todo struct {
	Id        int64  `json:"id"`
	Content   string `json:"content"`
	IsDone    bool   `json:"isDone"`
	IsDeleted bool   `json:"isDeleted"`
}
