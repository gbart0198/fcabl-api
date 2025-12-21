-- name: CreateUser :one
INSERT INTO users (email, phone_number, password_hash, first_name, last_name, role)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: GetUserByEmail :one
SELECT id, email, phone_number, first_name, last_name, role, created_at
FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT id, email, phone_number, first_name, last_name, role, created_at
FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT id, email, phone_number, first_name, last_name, role, created_at
FROM users
order by id;

-- name: UpdateUser :exec
UPDATE users
SET email = $1, phone_number = $2, first_name = $3, last_name = $4, role = $5, updated_at = $6
WHERE id = $7;

-- name: DeleteUser :exec
DELETE FROM users
where id = $1;

-- name: GetUserByEmailWithPassword :one
SELECT id, email, phone_number, password_hash, first_name, last_name, role, created_at, updated_at
FROM users WHERE email = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET password_hash = $1, updated_at = $2
WHERE id = $3;
