package domains

import "time"

type User struct {
	Id         int
	Login      string
	Password   string
	Email      string
	Name       string
	Created_at time.Time
}
