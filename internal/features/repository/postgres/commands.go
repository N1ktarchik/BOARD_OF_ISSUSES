package postgres

import (
	"context"
	"time"

	repo "Board_of_issuses/internal/features/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type connect struct {
	db *pgxpool.Pool
}

func (c *connect) CreateUser(ctx context.Context, user *repo.User) error {
	query := `INSERT INTO users (login,password,email,name,created_at) VALUES ($1,$2,$3,$4,$5) `

	if _, err := c.db.Exec(ctx, query, user.Login, user.Password, user.Email, user.Name, user.Created_at); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateUserEmail(ctx context.Context, user *repo.User) error {
	query := `UPDATE users SET email = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, user.Email, user.Id); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateUserName(ctx context.Context, user *repo.User) error {
	query := `UPDATE users SET name = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, user.Name, user.Id); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateUserPassword(ctx context.Context, user *repo.User) error {
	query := `UPDATE users SET password = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, user.Password, user.Id); err != nil {
		return err
	}

	return nil
}

func (c *connect) DeleteUser(ctx context.Context, user *repo.User) error {
	query := `DELETE FROM users WHERE id = $1`

	if _, err := c.db.Exec(ctx, query, user.Id); err != nil {
		return err
	}

	query = `DELETE FROM tasksusers WHERE userid = $1`
	if _, err := c.db.Exec(ctx, query, user.Id); err != nil {
		return err
	}

	return nil
}

func (c *connect) GetUserByID(ctx context.Context, id int) (*repo.User, error) {
	query := `SELECT id,login,password,email,name,created_at FROM users WHERE id = $1`

	user := &repo.User{}

	err := c.db.QueryRow(ctx, query, id).Scan(
		&user.Id,
		&user.Login,
		&user.Password,
		&user.Email,
		&user.Name,
		&user.Created_at,
	)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (c *connect) GetUserByLogin(ctx context.Context, login string) (*repo.User, error) {
	query := `SELECT id,login,password,email,name,created_at FROM users WHERE login = $1`

	user := &repo.User{}

	err := c.db.QueryRow(ctx, query, login).Scan(
		&user.Id,
		&user.Login,
		&user.Password,
		&user.Email,
		&user.Name,
		&user.Created_at,
	)

	if err != nil {
		return nil, err
	}

	return user, nil

}

/////////////////////////

func (c *connect) GetUserDesks(ctx context.Context, user *repo.User) ([]int, error) {
	query := `SELECT deskid FROM desksuser WHERE userid = $1`

	deskArr := make([]int, 0)

	rows, err := c.db.Query(ctx, query, user.Id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var deskId int

		if err := rows.Scan(&deskId); err != nil {
			return nil, err
		}

		deskArr = append(deskArr, deskId)

		if err := rows.Err(); err != nil {
			return nil, err
		}

	}

	return deskArr, nil
}

func (c *connect) CreateUserDesk(ctx context.Context, user *repo.User, desk *repo.Desk) error {
	query := `INSERT INTO desksusers (userid,deskid) VALUES ($1,$2) `

	if _, err := c.db.Exec(ctx, query, user.Id, desk.Id); err != nil {
		return err
	}

	return nil
}

func (c *connect) DeleteUserDesk(ctx context.Context, user *repo.User, desk *repo.Desk) error {
	query := `DELETE FROM desksusers WHERE userid = $1 AND deskid = $2`

	if _, err := c.db.Exec(ctx, query, user.Id, desk.Id); err != nil {
		return err
	}

	return nil
}

///////////////////////

func (c *connect) CreateDesk(ctx context.Context, desk *repo.Desk) error {
	query := `INSERT INTO desks (name,password,created_at) VALUES ($1,$2,$3) `

	if _, err := c.db.Exec(ctx, query, desk.Name, desk.Password, desk.Created_at); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateDeskName(ctx context.Context, desk *repo.Desk) error {
	query := `UPDATE desks SET name = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, desk.Name, desk.Id); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateDesksPassword(ctx context.Context, desk *repo.Desk) error {
	query := `UPDATE desks SET password = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, desk.Password, desk.Id); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateDeskOwner(ctx context.Context, desk *repo.Desk) error {
	query := `UPDATE desks SET ownerid = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, desk.OwnerId, desk.Id); err != nil {
		return err
	}

	return nil
}

func (c *connect) DeleteDesk(ctx context.Context, desk *repo.Desk) error {
	query := `DELETE FROM desks WHERE id = $1`

	if _, err := c.db.Exec(ctx, query, desk.Id); err != nil {
		return err
	}

	query = `DELETE FROM tasksusers WHERE deskid = $1`
	if _, err := c.db.Exec(ctx, query, desk.Id); err != nil {
		return err
	}

	return nil
}

// ////////////////////////////
func (c *connect) CreateTask(ctx context.Context, task *repo.Task) error {
	query := `INSERT INTO tasks (userid,deskid,name,description,time,created_at) VALUES ($1,$2,$3,$4,$5,$6) `

	if _, err := c.db.Exec(ctx, query, task.UserId, task.DeskId, task.Name, task.Description, task.Time, task.Created_at); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateTaskDecription(ctx context.Context, task *repo.Task) error {
	query := `UPDATE taks SET description = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, task.Description, task.Id); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateTaskTime(ctx context.Context, task *repo.Task) error {
	query := `UPDATE taks SET time = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, task.Time, task.Id); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateTaskDone(ctx context.Context, task *repo.Task) error {
	query := `UPDATE taks SET done = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, true, task.Id); err != nil {
		return err
	}

	return nil
}

func (c *connect) DeleteTask(ctx context.Context, task *repo.Task) error {
	query := `DELETE FROM tasks WHERE id = $1 `

	if _, err := c.db.Exec(ctx, query, task.Id); err != nil {
		return err
	}

	return nil
}

func (c *connect) GetAllTasksFromOneDesk(ctx context.Context, desk *repo.Desk) ([]repo.Task, error) {
	query := `SELECT id,userid,deskid,name,description,done,time,created_at FROM tasks WHERE deskid = $1`

	rows, err := c.db.Query(ctx, query, desk.Id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasksArr := make([]repo.Task, 0)

	for rows.Next() {
		var task repo.Task
		err := rows.Scan(
			&task.Id,
			&task.UserId,
			&task.DeskId,
			&task.Name,
			&task.Description,
			&task.Done,
			&task.Time,
			&task.Created_at,
		)

		if err != nil {
			return nil, err
		}

		tasksArr = append(tasksArr, task)

		if err := rows.Err(); err != nil {
			return nil, err
		}
	}

	return tasksArr, nil
}

func (c *connect) GetDoneTasksFromOneDesk(ctx context.Context, desk *repo.Desk) ([]repo.Task, error) {
	query := `SELECT id,userid,deskid,name,description,done,time,created_at FROM tasks WHERE deskid = $1 AND done = true`

	rows, err := c.db.Query(ctx, query, desk.Id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasksArr := make([]repo.Task, 0)

	for rows.Next() {
		var task repo.Task
		err := rows.Scan(
			&task.Id,
			&task.UserId,
			&task.DeskId,
			&task.Name,
			&task.Description,
			&task.Done,
			&task.Time,
			&task.Created_at,
		)

		if err != nil {
			return nil, err
		}

		tasksArr = append(tasksArr, task)

		if err := rows.Err(); err != nil {
			return nil, err
		}
	}

	return tasksArr, nil
}

func (c *connect) GetNotDoneTasksFromOneDesk(ctx context.Context, desk *repo.Desk) ([]repo.Task, error) {
	query := `SELECT id,userid,deskid,name,description,done,time,created_at FROM tasks WHERE deskid = $1 AND done = false`

	rows, err := c.db.Query(ctx, query, desk.Id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasksArr := make([]repo.Task, 0)

	for rows.Next() {
		var task repo.Task
		err := rows.Scan(
			&task.Id,
			&task.UserId,
			&task.DeskId,
			&task.Name,
			&task.Description,
			&task.Done,
			&task.Time,
			&task.Created_at,
		)

		if err != nil {
			return nil, err
		}

		tasksArr = append(tasksArr, task)

		if err := rows.Err(); err != nil {
			return nil, err
		}
	}

	return tasksArr, nil
}

func (c *connect) GetOverdueTasksFromOneDesk(ctx context.Context, desk *repo.Desk) ([]repo.Task, error) {
	query := `SELECT id,userid,deskid,name,description,done,time,created_at FROM tasks WHERE deskid = $1 AND time <= %2`

	rows, err := c.db.Query(ctx, query, desk.Id, time.Now())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasksArr := make([]repo.Task, 0)

	for rows.Next() {
		var task repo.Task
		err := rows.Scan(
			&task.Id,
			&task.UserId,
			&task.DeskId,
			&task.Name,
			&task.Description,
			&task.Done,
			&task.Time,
			&task.Created_at,
		)

		if err != nil {
			return nil, err
		}

		tasksArr = append(tasksArr, task)

		if err := rows.Err(); err != nil {
			return nil, err
		}
	}

	return tasksArr, nil
}
