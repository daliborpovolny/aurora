package services

import "fmt"

type EmailInUseError struct {
	Email string
}

func (e *EmailInUseError) Error() string {
	return fmt.Sprintf("email in use: %s", e.Email)
}

var EmailInUseErr *EmailInUseError

type BadPasswordError struct {
	Reason string
}

func (e *BadPasswordError) Error() string {
	return e.Reason
}

var BadPasswordErr *BadPasswordError

type UnknownEmail struct {
	Email string
}

func (e *UnknownEmail) Error() string {
	return fmt.Sprintf("unknown email: %s", e.Email)
}
