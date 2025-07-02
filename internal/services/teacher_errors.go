package services

import "fmt"

type UnknownTeacherIdError struct {
	id int64
}

func (e *UnknownTeacherIdError) Error() string {
	return fmt.Sprintf("unknown teacher id: %d", e.id)
}

var UnknownTeacherIdErr *UnknownTeacherIdError
