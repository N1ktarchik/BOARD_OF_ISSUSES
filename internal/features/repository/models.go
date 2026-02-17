package repository

import "time"

type User struct {
	Id         int
	Login      string
	Password   string
	Email      string
	Name       string
	Created_at time.Time
}

type Task struct {
	Id          int
	UserId      int
	DeskId      int
	Name        string
	Description string
	Done        bool
	Time        time.Time
	Created_at  time.Time
}

type Desk struct {
	Id         int
	Name       string
	Password   string
	OwnerId    int
	Created_at time.Time
}
