package postgres

import (
	"context"
	"database/sql"
	"time"

	repo "Board_of_issuses/internal/features/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type connect struct {
	db *pgxpool.Pool
}

func (c *connect) CreateUser(ctx context.Context, user *repo.User) error {
	query := `INSERT INTO users (login,password,email,name) VALUES ($1,$2,$3,$4) `

	if _, err := c.db.Exec(ctx, query, user.Login, user.Password, user.Email, user.Name); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateUserEmail(ctx context.Context, email string, userId int) error {
	query := `UPDATE users SET email = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, email, userId); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateUserName(ctx context.Context, name string, userId int) error {
	query := `UPDATE users SET name = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, name, userId); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateUserPassword(ctx context.Context, password string, userId int) error {
	query := `UPDATE users SET password = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, password, userId); err != nil {
		return err
	}

	return nil
}

func (c *connect) DeleteUser(ctx context.Context, userId string) error {
	tx, err := c.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	query := `DELETE FROM deskusers WHERE userid = $1`
	result, err := tx.Exec(ctx, query, userId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return sql.ErrNoRows
	}

	query = `DELETE FROM users WHERE id = $1`
	result, err = tx.Exec(ctx, query, userId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return sql.ErrNoRows
	}

	tx.Commit(ctx)

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

func (c *connect) GetUserByLoginOrEmail(ctx context.Context, login, email string) (*repo.User, error) {
	query := `SELECT id,login,password,email,name,created_at FROM users WHERE login = $1 OR email = $2`

	user := &repo.User{}

	err := c.db.QueryRow(ctx, query, login, email).Scan(
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

func (c *connect) CheckUserByEmailAndLogin(ctx context.Context, login, email string) (bool, error) {
	query := `SELECT EXISTS(
		SELECT 1 FROM users WHERE login = $1 AND email = $2)`

	var exists bool

	err := c.db.QueryRow(ctx, query, login, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

/////////////////////////

func (c *connect) GetUserDesks(ctx context.Context, userId int) ([]int, error) {
	query := `SELECT deskid FROM desksuser WHERE userid = $1`

	deskArr := make([]int, 0)

	rows, err := c.db.Query(ctx, query, userId)
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

// ?
// func (c *connect) CreateUserDesk(ctx context.Context, user *repo.User, desk *repo.Desk) error {
// 	query := `INSERT INTO desksusers (userid,deskid) VALUES ($1,$2) `
// 	if _, err := c.db.Exec(ctx, query, user.Id, desk.Id); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (c *connect) DeleteUserDesk(ctx context.Context, userId, deskId int) error {
	query := `DELETE FROM desksusers WHERE userid = $1 AND deskid = $2`

	if _, err := c.db.Exec(ctx, query, userId, deskId); err != nil {
		return err
	}

	return nil
}

///////////////////////

func (c *connect) CreateDesk(ctx context.Context, desk *repo.Desk) error {
	tx, err := c.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	query := `INSERT INTO desks (name,password,ownerid) VALUES ($1,$2,$3,) RETURNING id`
	err = tx.QueryRow(ctx, query, desk.Name, desk.Password, desk.OwnerId).Scan(&desk.Id)
	if err != nil {
		return err
	}

	query = `INSERT INTO desksusers (userid,deskid) VALUES ($1,$2)`
	result, err := tx.Exec(ctx, query, desk.OwnerId, desk.Id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return sql.ErrNoRows
	}

	tx.Commit(ctx)
	return nil
}

func (c *connect) UpdateDeskName(ctx context.Context, deskId int, name string) error {
	query := `UPDATE desks SET name = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, name, deskId); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateDesksPassword(ctx context.Context, deskId int, password string) error {
	query := `UPDATE desks SET password = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, password, deskId); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateDeskOwner(ctx context.Context, ownerid, deskid int) error {
	query := `UPDATE desks SET ownerid = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, ownerid, deskid); err != nil {
		return err
	}

	return nil
}

func (c *connect) DeleteDesk(ctx context.Context, deskId int) error {
	tx, err := c.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM tasksusers WHERE deskid = $1`
	result, err := tx.Exec(ctx, query, deskId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return sql.ErrNoRows
	}

	query = `DELETE FROM desks WHERE id = $1`
	result, err = c.db.Exec(ctx, query, deskId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return sql.ErrNoRows
	}

	tx.Commit(ctx)

	return nil
}

// ////////////////////////////
func (c *connect) CreateTask(ctx context.Context, task *repo.Task) error {
	query := `INSERT INTO tasks (userid,deskid,name,description,time) VALUES ($1,$2,$3,$4,$5) `

	if _, err := c.db.Exec(ctx, query, task.UserId, task.DeskId, task.Name, task.Description, task.Time); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateTaskDecription(ctx context.Context, id int, description string) error {
	query := `UPDATE taks SET description = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, description, id); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateTaskTime(ctx context.Context, id int, time time.Time) error {
	query := `UPDATE taks SET time = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, time, id); err != nil {
		return err
	}

	return nil
}

func (c *connect) UpdateTaskDone(ctx context.Context, id int) error {
	query := `UPDATE taks SET done = $1 WHERE id = $2 `

	if _, err := c.db.Exec(ctx, query, true, id); err != nil {
		return err
	}

	return nil
}

func (c *connect) DeleteTask(ctx context.Context, id int) error {
	query := `DELETE FROM tasks WHERE id = $1 `

	if _, err := c.db.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}

func (c *connect) GetAllTasksFromOneDesk(ctx context.Context, deskId int) ([]repo.Task, error) {
	query := `SELECT id,userid,deskid,name,description,done,time,created_at FROM tasks WHERE deskid = $1`

	rows, err := c.db.Query(ctx, query, deskId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasksArr := make([]repo.Task, 0)

	for rows.Next() {
		task := repo.Task{}
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

func (c *connect) GetDoneTasksFromOneDesk(ctx context.Context, deskId int) ([]repo.Task, error) {
	query := `SELECT id,userid,deskid,name,description,done,time,created_at FROM tasks WHERE deskid = $1 AND done = true`

	rows, err := c.db.Query(ctx, query, deskId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasksArr := make([]repo.Task, 0)

	for rows.Next() {
		task := repo.Task{}
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

func (c *connect) GetNotDoneTasksFromOneDesk(ctx context.Context, deskId int) ([]repo.Task, error) {
	query := `SELECT id,userid,deskid,name,description,done,time,created_at FROM tasks WHERE deskid = $1 AND done = false`

	rows, err := c.db.Query(ctx, query, deskId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasksArr := make([]repo.Task, 0)

	for rows.Next() {
		task := repo.Task{}
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

func (c *connect) GetOverdueTasksFromOneDesk(ctx context.Context, deskId int) ([]repo.Task, error) {
	query := `SELECT id,userid,deskid,name,description,done,time,created_at FROM tasks WHERE deskid = $1 AND time >= %2`

	rows, err := c.db.Query(ctx, query, deskId, time.Now())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasksArr := make([]repo.Task, 0)

	for rows.Next() {
		task := repo.Task{}
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
