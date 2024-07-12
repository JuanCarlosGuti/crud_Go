package user

import (
	"errors"
	"fmt"
)

var ErrFirstNameRequired = errors.New("firstname is required")
var ErrLastNameRequired = errors.New("lastname is required")

type ErrorNotFound struct {
	ID uint64
}

func (e ErrorNotFound) Error() string {
	return fmt.Sprintf("user id '%d' not found", e.ID)
}
