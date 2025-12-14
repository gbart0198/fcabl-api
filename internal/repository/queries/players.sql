-- name: CreatePlayer :one
INSERT INTO players (user_id, team_id, registration_fee_due, is_fully_registered, is_active, jersey_number, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
RETURNING *;

-- name: GetPlayerById :one
SELECT * FROM players WHERE id = $1;

-- name: GetPlayerByUserId :one
SELECT * FROM players WHERE user_id = $1;

-- name: ListPlayers :many
SELECT * FROM players
ORDER BY id;

-- name: ListActivePlayers :many
SELECT * FROM players
WHERE is_active = true
ORDER BY id;

-- name: ListPlayersByTeam :many
SELECT * FROM players
WHERE team_id = $1
ORDER BY jersey_number, id;

-- name: UpdatePlayer :exec
UPDATE players
SET team_id = $1, registration_fee_due = $2, is_fully_registered = $3, 
    is_active = $4, jersey_number = $5, updated_at = NOW()
WHERE id = $6;

-- name: UpdatePlayerTeam :exec
UPDATE players
SET team_id = $1, updated_at = NOW()
WHERE id = $2;

-- name: UpdatePlayerRegistrationStatus :exec
UPDATE players
SET registration_fee_due = $1, is_fully_registered = $2, updated_at = NOW()
WHERE id = $3;

-- name: DeletePlayer :exec
DELETE FROM players
WHERE id = $1;

-- name: GetPlayerWithUser :one
SELECT p.*, u.email, u.phone_number, u.first_name, u.last_name, u.role
FROM players p
INNER JOIN users u ON p.user_id = u.id
WHERE p.id = $1;

-- name: GetPlayerWithTeam :one
SELECT p.*, t.name as team_name
FROM players p
LEFT JOIN teams t ON p.team_id = t.id
WHERE p.id = $1;

-- name: ListPlayersWithUsers :many
SELECT p.*, u.email, u.first_name, u.last_name
FROM players p
INNER JOIN users u ON p.user_id = u.id
ORDER BY u.last_name, u.first_name;

-- name: ListFreeAgents :many
SELECT p.*, u.email, u.first_name, u.last_name
FROM players p
INNER JOIN users u ON p.user_id = u.id
WHERE p.team_id IS NULL AND p.is_active = true
ORDER BY u.last_name, u.first_name;
