-- name: CreateGameWithoutScore :one
INSERT INTO games (home_team_id, away_team_id, game_time, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING *;

-- name: CreateGameWithScore :one
INSERT INTO games (home_team_id, away_team_id, game_time, home_score, away_score, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
RETURNING *;

-- name: GetGameById :one
SELECT * FROM games WHERE id = $1;

-- name: ListGames :many
SELECT * FROM games
ORDER BY game_time;

-- name: ListGamesByTeam :many
SELECT g.id, g.home_team_id, g.away_team_id, g.home_score, g.away_score, g.game_time, g.created_at, g.updated_at, g.status, 
t_home.name home_name, t_away.name away_name
FROM games g
INNER JOIN teams t_home on t_home.id = g.home_team_id
INNER JOIN teams t_away on t_away.id = g.away_team_id
WHERE home_team_id = $1 OR away_team_id = $1
ORDER BY game_time;

-- name: UpdateGame :exec
UPDATE games
SET home_team_id = $1, away_team_id = $2, game_time = $3, home_score = $4, away_score = $5, status = $6, updated_at = NOW()
WHERE id = $7;

-- name: UpdateGameTime :exec
UPDATE games
SET game_time = $1, updated_at = NOW()
WHERE id = $2;

-- name: UpdateGameScoreAndStatus :exec
UPDATE games
SET home_score = $1, away_score = $2, status = $3, updated_at = NOW()
WHERE id = $4;

-- name: DeleteGame :exec
DELETE FROM games
WHERE id = $1;

-- name: GetGameWithTeams :one
SELECT g.*, ht.name as home_team_name, at.name as away_team_name
FROM games g
INNER JOIN teams ht ON g.home_team_id = ht.id
INNER JOIN teams at ON g.away_team_id = at.id
WHERE g.id = $1;

-- name: ListGamesWithTeams :many
SELECT g.id, g.home_team_id, g.away_team_id, g.home_score, g.away_score, 
       g.game_time, g.created_at, g.updated_at, g.status,
       ht.name as home_team_name,
       at.name as away_team_name
FROM games g
INNER JOIN teams ht ON g.home_team_id = ht.id
INNER JOIN teams at ON g.away_team_id = at.id
ORDER BY g.game_time;
