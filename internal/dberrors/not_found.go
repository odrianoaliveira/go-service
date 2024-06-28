package dberrors

import "fmt"

type NotFoundError struct {
	Entity string
	ID     string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("entity %s with ID %s not found", e.Entity, e.ID)
}
