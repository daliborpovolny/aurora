package services

import (
	db "aurora/database"
	gen "aurora/database/gen"
	"database/sql"

	"context"
)

type ParentServicer interface {
	ListParents(ctx context.Context) ([]gen.ListParentsRow, error)
	GetParent(parentId int64, ctx context.Context) (gen.GetParentRow, error)
}

var ParentService ParentServicer = ParentServiceStruct{}

type ParentServiceStruct struct {
}

func (t ParentServiceStruct) ListParents(ctx context.Context) ([]gen.ListParentsRow, error) {
	return db.Queries.ListParents(ctx)
}

func (t ParentServiceStruct) GetParent(parentId int64, ctx context.Context) (gen.GetParentRow, error) {
	parent, err := db.Queries.GetParent(ctx, parentId)
	if err == sql.ErrNoRows {
		return gen.GetParentRow{}, &UnknownParentIdError{parentId}
	}
	return parent, nil
}
