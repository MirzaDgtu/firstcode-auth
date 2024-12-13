package storage

import (
	"errors"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound  = errors.New("app not found")
)
