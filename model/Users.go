package model

import "time"

type Users struct {
	ID        int
	Username  string
	PasswordD string
	Pname     int
	Fullname  string
	Level     int
	CreateAt  time.Time
}
