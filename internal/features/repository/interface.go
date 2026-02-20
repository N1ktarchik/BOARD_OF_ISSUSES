package repository

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUserEmail(ctx context.Context, user *User) error
	UpdateUserName(ctx context.Context, user *User) error
	UpdateUserPassword(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id int) (*User, error)
	GetUserByLoginOrEmail(ctx context.Context, login, email string) (*User, error)
	CheckUserByEmailAndLogin(ctx context.Context, login, email string) (bool, error)

	GetUserDesks(ctx context.Context, user *User) ([]int, error)
	CreateUserDesk(ctx context.Context, user *User, desk *Desk) error
	DeleteUserDesk(ctx context.Context, user *User, desk *Desk) error

	CreateDesk(ctx context.Context, desk *Desk) error
	UpdateDeskName(ctx context.Context, desk *Desk) error
	UpdateDesksPassword(ctx context.Context, desk *Desk) error
	UpdateDeskOwner(ctx context.Context, desk *Desk) error
	DeleteDesk(ctx context.Context, desk *Desk) error

	CreateTask(ctx context.Context, task *Task) error
	UpdateTaskDecription(ctx context.Context, task *Task) error
	UpdateTaskTime(ctx context.Context, task *Task) error
	UpdateTaskDone(ctx context.Context, task *Task) error
	DeleteTask(ctx context.Context, task *Task) error
	GetAllTasksFromOneDesk(ctx context.Context, desk *Desk) ([]Task, error)
	GetDoneTasksFromOneDesk(ctx context.Context, desk *Desk) ([]Task, error)
	GetNotDoneTasksFromOneDesk(ctx context.Context, desk *Desk) ([]Task, error)
	GetOverdueTasksFromOneDesk(ctx context.Context, desk *Desk) ([]Task, error)
}
