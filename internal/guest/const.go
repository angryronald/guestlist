package guest

import "errors"

// Errors
var (
	ErrNotFound          = errors.New("guest not found")
	ErrAlreadyExist      = errors.New("guest already exists")
	ErrInsufficientSpace = errors.New("there is insufficient space")
	ErrAlreadyArrived    = errors.New("guest already arrived")
)
