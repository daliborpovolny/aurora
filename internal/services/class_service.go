package services

import (
	db "aurora/database"
	gen "aurora/database/gen"
	"database/sql"

	"context"
)

type ClassServicer interface {
	ListClasses(ctx context.Context) ([]gen.ListClassesRow, error)
	GetClass(classId int64, ctx context.Context) (gen.GetClassRow, error)
}

var ClassService ClassServicer = ClassServiceStruct{}

type ClassServiceStruct struct {
}

func (t ClassServiceStruct) ListClasses(ctx context.Context) ([]gen.ListClassesRow, error) {
	return db.Queries.ListClasses(ctx)
}

func (t ClassServiceStruct) GetClass(classId int64, ctx context.Context) (gen.GetClassRow, error) {
	class, err := db.Queries.GetClass(ctx, classId)
	if err == sql.ErrNoRows {
		return gen.GetClassRow{}, &UnknownClassIdError{classId}
	}
	return class, nil
}
