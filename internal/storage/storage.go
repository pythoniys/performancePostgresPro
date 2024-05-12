package postgres

import "errors"

var (
	ErrScriptNotFound      = errors.New("Script ot found")
	ErrScriptAlreadyExists = errors.New("Script already exists")
)
