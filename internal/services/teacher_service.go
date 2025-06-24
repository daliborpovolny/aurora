package services

import (
	db "aurora/database"
	gen "aurora/database/gen"
	"database/sql"

	"context"
)

type TeacherServiceStruct struct {
}

var TeacherService TeacherServiceStruct

func (t TeacherServiceStruct) ListTeachers(ctx context.Context) ([]gen.ListTeachersRow, error) {
	return db.Queries.ListTeachers(ctx)
}

func (t TeacherServiceStruct) GetTeacher(teacherId int64, ctx context.Context) (gen.GetTeacherRow, error) {
	teacher, err := db.Queries.GetTeacher(ctx, teacherId)
	if err == sql.ErrNoRows {
		return gen.GetTeacherRow{}, &UnknownTeacherIdError{teacherId}
	}
	return teacher, nil
}
