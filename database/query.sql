
--* User

-- name: GetUser :one
SELECT * FROM user
WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM user
WHERE email = ?;

-- name: ListUsers :many
SELECT * FROM user;

-- name: CreateUser :one
INSERT INTO user (
    first_name, last_name, email, hash
) VALUES (
    ?, ?, ?, ?
)
RETURNING *;

-- name: GetUserType :one
SELECT
  CASE
    WHEN EXISTS (SELECT 1 FROM teacher WHERE user_id = user.id) THEN 'teacher'
    WHEN EXISTS (SELECT 1 FROM admin WHERE user_id = user.id) THEN 'admin'
    WHEN EXISTS (SELECT 1 FROM student WHERE user_id = user.id) THEN 'student'
    WHEN EXISTS (SELECT 1 FROM parent WHERE user_id = user.id) THEN 'parent'
    ELSE 'unknown'
  END AS role
FROM
  "user"
WHERE
  user.id = ?;

--* Session

-- name: GetUserBySessionCookie :one
SELECT * FROM user
JOIN session on user.id = session.user_id
WHERE session.cookie = ?;

-- name: CreateSession :one
INSERT INTO session (
    user_id, cookie, created_at, expires_at
) VALUES (
    ?, ?, ?, ?
) RETURNING *;


--* Student

-- name: GetStudent :one
SELECT 
    student.id AS student_id,
    student.class_id,
    user.id AS user_id,
    user.first_name, 
    user.last_name, 
    user.email
FROM student
JOIN user ON student.user_id = user.id
WHERE student.id = ?;

-- name: GetParentsOfStudent :many
SELECT
    parent.id AS parent_id,
    user.id AS user_id,
    user.first_name,
    user.last_name,
    user.email
FROM student_parent
JOIN parent ON parent.id = student_parent.parent_id
JOIN user ON user.id = parent.user_id
WHERE student_parent.student_id = ?;

-- name: ListStudents :many
SELECT 
    student.id AS student_id,
    user.id AS user_id,
    user.first_name,
    user.last_name,
    user.email
FROM student
JOIN user ON student.user_id = user.id
WHERE student.has_graduated = 0;

-- name: CreateStudent :one
INSERT INTO student (
    user_id, class_id, has_graduated
) VALUES (
    ?, ?, 0
) RETURNING *;

--* Graduation

-- name: GraduateClass :exec
UPDATE class
SET
    has_graduated = 1
WHERE
    class.id = ?;

-- name: GraduateStudentOfClass :exec
UPDATE student
SET
    has_graduated = 1
WHERE student.class_id = ?;


--* Teacher

-- name: GetTeacher :one
SELECT
    teacher.id AS teacher_id,
    user.id AS user_id,
    user.first_name,
    user.last_name,
    user.email
FROM teacher
JOIN user ON teacher.user_id = user.id
WHERE teacher.id = ?;

-- name: ListTeachers :many
SELECT
    teacher.id AS teacher_id,
    user.id AS user_id,
    user.first_name,
    user.last_name,
    user.email
FROM teacher
JOIN user ON teacher.user_id = user.id;

-- name: CreateTeacher :one
INSERT INTO teacher (
    user_id
) VALUES (
    ?
) RETURNING *;


--* Parent

-- name: GetParent :one
SELECT
    parent.id AS parent_id,
    user.id AS user_id,
    user.first_name,
    user.last_name,
    user.email
FROM parent
JOIN user on parent.user_id = user.id
WHERE parent.id = ?;

-- name: ListParents :many
SELECT
    parent.id AS parent_id,
    user.id AS user_id,
    user.first_name,
    user.last_name,
    user.email
FROM parent
JOIN user on parent.user_id = user.id;

-- name: CreateParent :one
INSERT INTO parent (
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

-- name: GetStudentOfParent :many
SELECT
    student.id,
    user.id AS user_id,
    user.first_name,
    user.last_name,
    user.email
FROM student
JOIN student_parent on student_parent.student_id = student.id
JOIN user on student.user_id = user.id
WHERE student_parent.parent_id = ?;


--* Admin

-- name: GetAdmin :one
SELECT
    admin.id AS admin_id,
    user.id AS user_id,
    user.first_name,
    user.last_name,
    user.email
FROM admin
JOIN user on admin.user_id = user.id
WHERE admin.id = ?;

-- name: ListAdmins :many
SELECT
    admin.id AS admin_id,
    user.id AS user_id,
    user.first_name,
    user.last_name,
    user.email
FROM admin
JOIN user on admin.user_id = user.id;

-- name: CreateAdmin :one
INSERT INTO admin (
    user_id
) VALUES (
    ?
) RETURNING *;


--* Class

-- name: GetClass :one
SELECT
    class.id AS class_id,
    class.name,
    class.teacher_id,
    class.room,
    class.start_year,
    class.graduation_year
FROM class
WHERE class.id = ?;

-- name: ListClasses :many
SELECT
    class.id AS class_id,
    class.name,
    class.teacher_id,
    class.room,
    class.start_year,
    class.graduation_year
FROM class
WHERE class.has_graduated = 0;

-- name: ListStudentsOfClass :many
SELECT 
    student.id AS student_id,
    user.id AS user_id,
    user.first_name,
    user.last_name,
    user.email
FROM student
JOIN user ON student.user_id = user.id
WHERE student.class_id = ?;

-- name: CreateClass :one
INSERT INTO class (
    teacher_id, name, room, start_year, graduation_year, has_graduated
) VALUES (
    ?, ?, ?, ?, ?, 0
) RETURNING *;