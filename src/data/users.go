package users

import (
	"time"
)

type User struct {
	Name      string
	Password  string
	Email     string
	Admin     bool
	LastLogin time.Time
	LastNotif time.Time
}
