package models

import (
	"time"
)

type User struct {
	Email       string
	Password    string
	FirstName   string
	LastName    string
	AccessToken AccessToken
}

type AccessToken struct {
	Value     string
	TTL       time.Duration
	ExpiresAt time.Time
}
