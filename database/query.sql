
-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ?;

-- name: GetUserBySessionCookie :one
SELECT * FROM users
JOIN sessions on users.id = sessions.user_id
WHERE sessions.cookie = ?;

-- name: CreateSession :one
INSERT INTO sessions (
    user_id, cookie, created_at, expires_at
) VALUES (
    ?, ?, ?, ?
) RETURNING *;

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
SELECT 
    students.id,
    users.id AS user_id,
    users.first_name, 
    users.last_name, 
    users.email
FROM students
JOIN users ON students.user_id = users.id
WHERE students.id = ?;

-- name: ListStudents :many
SELECT 
    students.id,
    users.id AS user_id,
    users.first_name, 
    users.last_name, 
    users.email
FROM students
JOIN users ON students.user_id = users.id;

-- name: CreateStudent :one
INSERT INTO students (
    user_id
) VALUES (
    ?
) RETURNING *;

-- name: GetTeacher :one
SELECT
    teachers.id,
    users.id AS user_id,
    users.first_name,
    users.last_name,
    users.email
FROM teachers
JOIN users ON teacher.user_id = users.id
WHERE teachers.id = ?;

-- name: ListTeachers :many
SELECT
    teachers.id,
    users.id AS user_id,
    users.first_name,
    users.last_name,
    users.email
FROM teachers
JOIN users ON teachers.user_id = users.id;

-- name: CreateTeacher :one
INSERT INTO teachers (
    user_id
) VALUES (
    ?
) RETURNING *;

-- name: GetParent :one
SELECT
    parents.id,
    users.id AS users_id,
    users.first_name,
    users.last_name,
    users.email
FROM parents
JOIN users on parents.user_id = users.id
WHERE parents.id = ?;

-- name: ListParents :many
SELECT
    parents.id,
    users.id AS users_id,
    users.first_name,
    users.last_name,
    users.email
FROM parents
JOIN users on parents.user_id = users.id;

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

-- name: GetStudentsOfParent :many
SELECT
    students.id,
    users.id AS user_id,
    users.first_name,
    users.last_name,
    users.email
FROM student_parent
JOIN students on student_parent.student_id = students.id
JOIN users on students.user_id = users.id
WHERE student_parent.parent_id = ?;

-- name: GetAdmin :one
SELECT
    admins.id,
    users.id AS user_id,
    users.first_name,
    users.last_name,
    users.email
FROM admins
JOIN users on admins.user_id = users.id
WHERE admins.id = ?;

-- name: ListAdmins :many
SELECT
    admins.id,
    users.id AS user_id,
    users.first_name,
    users.last_name,
    users.email
FROM admins
JOIN users on admins.user_id = users.id;

-- name: CreateAdmin :one
INSERT INTO admins (
    user_id
) VALUES (
    ?
) RETURNING *;