package services

import "fmt"

type UnknownStudentIdError struct {
	id int64
}

func (e UnknownStudentIdError) Error() string {
	return fmt.Sprintf("unknown student id: %d", e.id)
}

var UnknownStudentIdErr error = UnknownStudentIdError{}
