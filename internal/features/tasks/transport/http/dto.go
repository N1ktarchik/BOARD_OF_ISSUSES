package http

import (
	"N1ktarchik/Board_of_issues/internal/core/domain"
	"time"

	"github.com/google/uuid"
)

type TaskRequestDTO struct {
	DeskId      uuid.UUID `json:"desk_id" example:"832t758-a12g-47y9-i999-123456789098"`
	Name        string    `json:"name" example:"Task name"`
	Description string    `json:"description" example:"Task description"`
	Done        bool      `json:"status" example:"false"`
	Deadline    time.Time `json:"deadline" example:"2023-10-10T10:00:00Z"`
}

type UpdateTaskRequestDTO struct {
	Id          uuid.UUID `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name        string    `json:"name" example:"Task name"`
	Description string    `json:"description" example:"Task description"`
	Deadline    time.Time `json:"deadline" example:"2023-10-10T10:00:00Z"`
}

func (u *UpdateTaskRequestDTO) ToServiceUpdateTask() *domain.UpdateTask {
	return &domain.UpdateTask{
		Id:          u.Id,
		Name:        u.Name,
		Description: u.Description,
		Deadline:    u.Deadline,
	}
}

func (t *TaskRequestDTO) ToServiceTask() *domain.Task {
	return &domain.Task{
		DeskId:      t.DeskId,
		Name:        t.Name,
		Description: t.Description,
		Done:        t.Done,
		Deadline:    t.Deadline,
	}
}
