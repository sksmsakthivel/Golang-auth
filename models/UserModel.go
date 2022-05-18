package models

import (
	"time"
)

type User struct {
	Id       int
	Name     string
	UserName string
	Password string
	DOB      time.Time
}
