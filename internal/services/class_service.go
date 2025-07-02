package services

import (
	db "aurora/database"
	gen "aurora/database/gen"
	"database/sql"

	"context"
)

type ClassServiceStruct struct {
}

var ClassService ClassServiceStruct

func (t ClassServiceStruct) ListClasss(ctx context.Context) ([]gen.ListClassesRow, error) {
	return db.Queries.ListClasses(ctx)
}

func (t ClassServiceStruct) GetClass(ClassId int64, ctx context.Context) (gen.GetClassRow, error) {
	class, err := db.Queries.GetClass(ctx, ClassId)
	if err == sql.ErrNoRows {
		return gen.GetClassRow{}, &UnknownClassIdError{ClassId}
	}
	return class, nil
}
