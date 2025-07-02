package services

import (
	db "aurora/database"
	gen "aurora/database/gen"
	"database/sql"

	"context"
)

type StudentServiceStruct struct {
}

var StudentService StudentServiceStruct

func (t StudentServiceStruct) ListStudents(ctx context.Context) ([]gen.ListStudentsRow, error) {
	return db.Queries.ListStudents(ctx)
}

func (t StudentServiceStruct) GetStudent(StudentId int64, ctx context.Context) (gen.GetStudentRow, error) {
	Student, err := db.Queries.GetStudent(ctx, StudentId)
	if err != nil {
		if err == sql.ErrNoRows {
			return gen.GetStudentRow{}, &UnknownStudentIdError{StudentId}
		}
		return gen.GetStudentRow{}, err
	}
	return Student, nil
}
