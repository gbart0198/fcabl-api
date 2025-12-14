-- name: CreateUser :one
INSERT INTO users (email, phone_number, password_hash, first_name, last_name, role, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT id, email, phone_number, first_name, last_name, role, created_at
FROM users
order by last_name;

-- name: UpdateUser :exec
UPDATE users
SET email = $1, phone_number = $2, first_name = $3, last_name = $4, role = $5, updated_at = $6
WHERE id = $7;


