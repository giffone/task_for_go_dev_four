package domain

import "errors"

var (
	ErrUserName = errors.New("user name is empty")
	ErrTextBody = errors.New("text body is empty")
)
