package models

import (
	"errors"
)

var (
	ErrNoRecord           = errors.New("models: No Matching Record Found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)
