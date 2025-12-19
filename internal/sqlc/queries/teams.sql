-- name: GetTeamById :one
SELECT * FROM teams where id = $1;

-- name: ListTeams :many
SELECT * FROM teams
ORDER BY name;

-- name: CreateTeam :one
INSERT INTO teams (name, wins, losses, draws, points_for, points_against, created_at, updated_at)
values ($1, $2, $3, $4, $5, $6, NOW(), NOW())
RETURNING *;


-- name: UpdateTeam :exec
UPDATE teams
SET name = $1, wins = $2, losses = $3, draws = $4, points_for = $5, points_against = $6, updated_at = NOW()
WHERE id = $7;

-- name: DeleteTeam :exec
DELETE FROM teams
WHERE id = $1;

-- name: GetTeamWithPlayers :many
SELECT p.id, p.user_id, p.team_id, p.registration_fee_due, p.is_fully_registered,
       p.is_active, p.jersey_number, p.created_at, p.updated_at
FROM players p
WHERE p.team_id = $1
ORDER BY p.id;

-- name: ListTeamsWithPlayers :many
SELECT t.id, t.name, t.wins, t.losses, t.draws, t.points_for, t.points_against, t.created_at, t.updated_at, p.jersey_number,
u.first_name, u.last_name
FROM teams t
LEFT JOIN players p on t.id = p.team_id
LEFT JOIN users u on u.id = p.user_id
ORDER BY t.id;

-- name: GetTeamStats :one
SELECT t.*,
       COUNT(p.id) as player_count
FROM teams t
LEFT JOIN players p ON p.team_id = t.id AND p.is_active = true
WHERE t.id = $1
GROUP BY t.id;

-- name: GetTeamStandings :many
SELECT id, name, wins, losses, draws, points_for, points_against, created_at, updated_at,
       (wins * 3 + draws) as points,
       (points_for - points_against) as point_differential
FROM teams
ORDER BY points DESC, point_differential DESC, name ASC;
