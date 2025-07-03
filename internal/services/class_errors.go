package services

import "fmt"

type UnknownClassIdError struct {
	id int64
}

func (e *UnknownClassIdError) Error() string {
	return fmt.Sprintf("unknown class id: %d", e.id)
}

var UnknownClassIdErr error = &UnknownClassIdError{}
