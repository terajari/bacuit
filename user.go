package bacuit

import "time"

type UserRepo interface {
}

type User struct {
	ID        string
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
