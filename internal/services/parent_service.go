package services

import (
	db "aurora/database"
	gen "aurora/database/gen"
	"database/sql"

	"context"
)

type ParentServiceStruct struct {
}

var ParentService ParentServiceStruct

func (t ParentServiceStruct) ListParents(ctx context.Context) ([]gen.ListParentsRow, error) {
	return db.Queries.ListParents(ctx)
}

func (t ParentServiceStruct) GetParent(ParentId int64, ctx context.Context) (gen.GetParentRow, error) {
	Parent, err := db.Queries.GetParent(ctx, ParentId)
	if err == sql.ErrNoRows {
		return gen.GetParentRow{}, &UnknownParentIdError{ParentId}
	}
	return Parent, nil
}
