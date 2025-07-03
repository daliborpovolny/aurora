package services

import (
	db "aurora/database"
	gen "aurora/database/gen"
	"database/sql"

	"context"
)

type StudentServicer interface {
	ListStudents(ctx context.Context) ([]gen.ListStudentsRow, error)
	GetStudent(studentId int64, ctx context.Context) (gen.GetStudentRow, error)
}

var StudentService StudentServicer = StudentServiceStruct{}

type StudentServiceStruct struct {
}

func (t StudentServiceStruct) ListStudents(ctx context.Context) ([]gen.ListStudentsRow, error) {
	return db.Queries.ListStudents(ctx)
}

func (t StudentServiceStruct) GetStudent(studentId int64, ctx context.Context) (gen.GetStudentRow, error) {
	student, err := db.Queries.GetStudent(ctx, studentId)
	if err != nil {
		if err == sql.ErrNoRows {
			return gen.GetStudentRow{}, &UnknownStudentIdError{studentId}
		}
		return gen.GetStudentRow{}, err
	}
	return student, nil
}
