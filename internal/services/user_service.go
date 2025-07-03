package services

import (
	db "aurora/database"
	gen "aurora/database/gen"

	"context"
)

type UserServicer interface {
	ListUsers(ctx context.Context) ([]gen.User, error)
}

var UserService UserServicer = UserServiceStruct{}

type UserServiceStruct struct {
}

func (u UserServiceStruct) ListUsers(ctx context.Context) ([]gen.User, error) {
	return db.Queries.ListUsers(ctx)
}
