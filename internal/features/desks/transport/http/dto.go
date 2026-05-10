package http

import (
	"N1ktarchik/Board_of_issues/internal/core/domain"

	"github.com/google/uuid"
)

type DeskRequestDTO struct {
	Name     string `json:"name" example:"My Desk"`
	Password string `json:"password" example:"mysecretpassword"`
}

type DeskUpdateRequestDTO struct {
	Id       uuid.UUID `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name     string    `json:"name" example:"My Desk"`
	Password string    `json:"password" example:"mysecretpassword"`
}

type DeskConnectRequestDTO struct {
	Password string    `json:"password" example:"mysecretpassword"`
	Id       uuid.UUID `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *DeskRequestDTO) ToServiceDesk() *domain.Desk {
	return &domain.Desk{
		Name:     d.Name,
		Password: d.Password,
	}
}

func (d *DeskConnectRequestDTO) ToServiceDesk() *domain.Desk {
	return &domain.Desk{
		Id:       d.Id,
		Password: d.Password,
	}
}

func (d *DeskUpdateRequestDTO) ToServiceUpdateDesk() *domain.Desk {
	return &domain.Desk{
		Id:       d.Id,
		Name:     d.Name,
		Password: d.Password,
	}
}
