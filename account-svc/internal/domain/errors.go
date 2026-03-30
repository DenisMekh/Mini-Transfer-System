package domain

import "errors"

var (
	ErrNotFound      = errors.New("не найден")
	ErrAlreadyExists = errors.New("уже существует")
)
