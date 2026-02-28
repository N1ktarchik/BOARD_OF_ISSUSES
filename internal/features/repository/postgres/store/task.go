package store

import (
	"context"
	"time"

	repo "Board_of_issuses/internal/features/repository"
)

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

func (c *connect) GetTaskOwner(ctx context.Context, taskID int) (int, error) {
	query := `SELECT userid FROM tasks WHERE id =$1`

	var owner int

	if err := c.db.QueryRow(ctx, query, taskID).Scan(&owner); err != nil {
		return 0, err
	}

	return owner, nil
}

func (c *connect) GetDeskIDByTask(ctx context.Context, taskID int) (int, error) {
	query := `SELECT dekid FROM tasks WHERE id =$1`

	var deskid int

	if err := c.db.QueryRow(ctx, query, taskID).Scan(&deskid); err != nil {
		return 0, err
	}

	return deskid, nil
}
