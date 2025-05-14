
-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ?;

-- name: ListUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (
    first_name, last_name, email, hash
) VALUES (
    ?, ?, ?, ?
)
RETURNING *;

-- name: GetStudent :one
SELECT * FROM students
WHERE id = ?;

-- name: ListStudents :many
SELECT * FROM students;

-- name: CreateStudent :one
INSERT INTO students (
    user_id
) VALUES (
    ?
) RETURNING *;

-- name: GetTeachers :one
SELECT * FROM teachers
WHERE id = ?;

-- name: ListTeachers :many
SELECT * FROM teachers;

-- name: CreateTeacher :one
INSERT INTO teachers (
    user_id
) VALUES (
    ?
) RETURNING *;

-- name: GetParents :one
SELECT * FROM parents
WHERE id = ?;

-- name: ListParents :many
SELECT * FROM parents;

-- name: CreateParent :one
INSERT INTO parents (
    user_id
) VALUES (
    ?
) RETURNING *;

-- name: AssignStudentToParent :exec
INSERT INTO student_parent (
    parent_id, student_id
) VALUES (
    ?, ?
);

-- name: GetAdmins :one
SELECT * FROM admins
WHERE id = ?;

-- name: ListAdmins :many
SELECT * FROM admins;

-- name: CreateAdmins :one
INSERT INTO admins (
    user_id
) VALUES (
    ?
) RETURNING *;