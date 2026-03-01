package dto

import (
	dn "Board_of_issuses/internal/core/domains"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	DeskId      int       `json:"desk_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Done        bool      `json:"status"`
	Time        time.Time `json:"time"`
	Created_at  time.Time `json:"created_at"`
}

func (t *Task) ToServicenTask() *dn.Task {
	return &dn.Task{
		Id:          t.Id,
		UserId:      t.UserId,
		DeskId:      t.DeskId,
		Name:        t.Name,
		Description: t.Description,
		Done:        t.Done,
		Time:        t.Time,
		Created_at:  t.Created_at,
	}
}

type UpdateTaskTimeRequest struct {
	Hours int `json:"hours"`
}

type UpdateTaskDescriptionRequest struct {
	Description string `json:"description"`
}
