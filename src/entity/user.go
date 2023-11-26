package entity

import "time"

type User struct {
	Id        string
	Name      string
	Password  string
	Token     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
