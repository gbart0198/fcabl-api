-- name: GetTeamById :one
SELECT * FROM teams where id = $1;

-- name: ListTeams :many
SELECT * FROM teams
ORDER BY name;

-- name: CreateTeam :one
INSERT INTO teams (name, created_at, updated_at)
values ($1, NOW(), NOW())
RETURNING *;

-- name: UpdateTeamName :exec
UPDATE teams
SET name = $1, updated_at = NOW()
WHERE id = $2;

-- name: DeleteTeam :exec
DELETE FROM teams
WHERE id = $1;

-- name: GetTeamWithPlayersDetailed :many
SELECT t.id, t.name, t.created_at, t.updated_at, p.user_id, u.first_name, u.last_name, u.email, u.phone_number, 
        p.id as player_id, p.fee_remainder, p.jersey_number
FROM players p
INNER JOIN teams t on p.team_id = t.id
INNER JOIN users u on u.id = p.user_id
WHERE p.team_id = $1
ORDER BY p.id;

-- name: GetTeamWithPlayers :one
SELECT t.id, t.name, p.user_id, u.first_name, u.last_name, p.id as player_id, p.jersey_number
FROM players p
INNER JOIN teams t on p.team_id = t.id
INNER JOIN users u on u.id = p.user_id
WHERE p.team_id = $1
ORDER BY p.id;

-- name: ListTeamsWithPlayersDetailed :many
SELECT t.id, t.name, t.created_at, t.updated_at, p.user_id, u.first_name, u.last_name, u.email, u.phone_number, 
        p.id as player_id, p.fee_remainder, p.jersey_number
FROM teams t
LEFT JOIN players p on t.id = p.team_id
LEFT JOIN users u on u.id = p.user_id
ORDER BY t.id;

-- name: ListTeamsWithPlayers :many
SELECT t.id, t.name, p.user_id, u.first_name, u.last_name, p.id as player_id, p.jersey_number
FROM teams t
LEFT JOIN players p on t.id = p.team_id
LEFT JOIN users u on u.id = p.user_id
ORDER BY t.id;

-- name: ListTeamSchedule :many
SELECT g.id, g.home_team_id, g.away_team_id, g.home_score, g.away_score,
       g.game_time, g.created_at, g.updated_at, g.status,
       ht.name as home_team_name,
       at.name as away_team_name
FROM games g
INNER JOIN teams ht ON g.home_team_id = ht.id
INNER JOIN teams at ON g.away_team_id = at.id
WHERE g.home_team_id = $1 OR g.away_team_id = $1
ORDER BY g.game_time;
