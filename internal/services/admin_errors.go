package services

import "fmt"

type UnknownAdminIdError struct {
	id int64
}

func (e *UnknownAdminIdError) Error() string {
	return fmt.Sprintf("unknown admin id: %d", e.id)
}

var UnknownAdminIdErr error = &UnknownAdminIdError{}
