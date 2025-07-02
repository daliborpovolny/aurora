package services

import "fmt"

type UnknownParentIdError struct {
	id int64
}

func (e UnknownParentIdError) Error() string {
	return fmt.Sprintf("unknown Parent id: %d", e.id)
}

var UnknownParentIdErr error = UnknownParentIdError{}
