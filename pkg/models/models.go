package models

import (
	"time"
)

// A struct that hold Blog
type Blog struct {
	ID           int
	FirstName    string
	LastName     string
	Email        string
	Subject      string
	Message      string
	Date_Created time.Time
}
