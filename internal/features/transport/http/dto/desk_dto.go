package dto

import (
	dn "Board_of_issuses/internal/core/domains"
	"time"
)

type Desk struct {
	Id         int
	Name       string
	Password   string
	OwnerId    int
	Created_at time.Time
}

func (d *Desk) ToServiceDeskr() *dn.Desk {
	return &dn.Desk{
		Id:         d.Id,
		Name:       d.Name,
		Password:   d.Password,
		OwnerId:    d.OwnerId,
		Created_at: d.Created_at,
	}

}

type UpdateDeskNameRequest struct {
	Name string `json:"name"`
}

type UpdateDeskPasswordRequest struct {
	Password string `json:"password"`
}

type UpdateDeskOwnerRequest struct {
	ID int `json:"new_owner_id"`
}
