package database

import (
	database "aurora/database/gen"
	"context"
	_ "embed"
	"os"

	"database/sql"

	_ "modernc.org/sqlite"
)

var Queries *database.Queries

//go:embed schema.sql
var ddl string

func Initialize() *sql.DB {

	resetDB := os.Getenv("RESET_DB")

	if resetDB == "true" {
		os.Remove("./database/db_file.db")
		Seed()
	}

	ctx := context.Background()

	db, err := sql.Open("sqlite", "./database/db_file.db")
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		panic(err)
	}

	Queries = database.New(db)

	return db
}

func Seed() {
	ctx := context.Background()
	if users, _ := Queries.ListUsers(ctx); len(users) == 0 {
		Queries.CreateUser(ctx, database.CreateUserParams{
			FirstName: "par",
			LastName:  "ent",
			Email:     "par.ent@gmail.com",
			Hash:      "parentHash",
		})

		Queries.CreateUser(ctx, database.CreateUserParams{
			FirstName: "tea",
			LastName:  "cher",
			Email:     "tea.cher@isic.com",
			Hash:      "teachingHash",
		})

		Queries.CreateUser(ctx, database.CreateUserParams{
			FirstName: "stu",
			LastName:  "end",
			Email:     "stu.dent@school.com",
			Hash:      "studentsHash",
		})

		Queries.CreateUser(ctx, database.CreateUserParams{
			FirstName: "ad",
			LastName:  "min",
			Email:     "ad.min@manage.com",
			Hash:      "adminsHash",
		})

		_, err := Queries.CreateParent(ctx, 1)
		if err != nil {
			panic(err)
		}

		teacher, err := Queries.CreateTeacher(ctx, 2)
		if err != nil {
			panic(err)
		}

		class, err := Queries.CreateClass(ctx, database.CreateClassParams{
			TeacherID:      teacher.ID,
			Room:           "A24",
			StartYear:      2024,
			GraduationYear: 2032,
		})
		if err != nil {
			panic(err)
		}

		Queries.CreateStudent(ctx, database.CreateStudentParams{
			UserID:  3,
			ClassID: class.ID,
		})

		Queries.CreateAdmin(ctx, 4)

		Queries.AssignStudentToParent(ctx, database.AssignStudentToParentParams{
			ParentID:  1,
			StudentID: 1,
		})
	}
}
