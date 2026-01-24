-- name: CreatePlayer :one
INSERT INTO players (user_id, team_id, fee_remainder, jersey_number, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW())
RETURNING *;

-- name: GetPlayerById :one
SELECT * FROM players WHERE id = $1;

-- name: GetPlayerByUserId :one
SELECT * FROM players WHERE user_id = $1;

-- name: ListPlayers :many
SELECT * FROM players
ORDER BY id;

-- name: ListPlayersByTeam :many
SELECT * FROM players
WHERE team_id = $1
ORDER BY jersey_number, id;

-- name: UpdatePlayer :exec
UPDATE players
SET team_id = $1, fee_remainder = $2, jersey_number = $3, updated_at = NOW()
WHERE id = $4;

-- name: UpdatePlayerTeam :exec
UPDATE players
SET team_id = $1, updated_at = NOW()
WHERE id = $2;

-- name: UpdatePlayerRegistrationFee :exec
UPDATE players
SET fee_remainder = $1, updated_at = NOW()
WHERE id = $2;

-- name: DeletePlayer :exec
DELETE FROM players
WHERE id = $1;

-- name: GetPlayerWithUserInfoDetailed :one
SELECT p.id, p.user_id, p.team_id, p.fee_remainder, p.jersey_number, p.created_at, p.updated_at,
u.email, u.phone_number, u.first_name, u.last_name, u.role
FROM players p
INNER JOIN users u ON p.user_id = u.id
WHERE p.id = $1;

-- name: GetPlayerWithUserInfo :one
SELECT p.id, p.user_id, p.team_id, p.jersey_number, u.first_name, u.last_name
FROM players p
INNER JOIN users u ON p.user_id = u.id
WHERE p.id = $1;

-- name: GetPlayerWithUserAndTeamInfo :one
SELECT p.id, p.user_id, p.team_id, t.name, p.jersey_number, u.first_name, u.last_name
FROM players p
INNER JOIN users u ON p.user_id = u.id
INNER JOIN teams t on t.id = p.team_id
WHERE p.id = $1;

-- name: ListPlayersWithUsersDetailed :many
SELECT p.id, p.user_id, p.team_id, p.fee_remainder, p.jersey_number, p.created_at, p.updated_at,
u.email, u.phone_number, u.first_name, u.last_name, u.role
FROM players p
INNER JOIN users u ON p.user_id = u.id
ORDER BY u.last_name, u.first_name;

-- name: ListPlayersWithUsers :many
SELECT p.id, p.user_id, p.team_id, p.jersey_number, u.first_name, u.last_name
FROM players p
INNER JOIN users u ON p.user_id = u.id
ORDER BY u.last_name, u.first_name;
