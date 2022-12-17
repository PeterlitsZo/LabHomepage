package databaseError

import "errors"

var (
	DbConnectionIsNil         = errors.New("database connection is nil")
	DbConnectionNotRegistered = errors.New("database connection is not registered")
)
