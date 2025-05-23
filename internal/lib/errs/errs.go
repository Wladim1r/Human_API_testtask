package errs

import "errors"

var (
	ErrInvalidParam = errors.New("invalid parameter value")
	ErrInvalidSize  = errors.New("invalid size value")
	ErrInvalidID    = errors.New("invalid ID format")
	ErrNotFound     = errors.New("record not found")
	ErrDBOperation  = errors.New("database operation failed")
)
