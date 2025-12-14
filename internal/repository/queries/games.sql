-- name: CreateGame :one
INSERT INTO games (home_team_id, away_team_id, game_time, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING *;

-- name: GetGameById :one
SELECT * FROM games WHERE id = $1;

-- name: ListGames :many
SELECT * FROM games
ORDER BY game_time;

-- name: ListUpcomingGames :many
SELECT * FROM games
WHERE game_time > NOW()
ORDER BY game_time;

-- name: ListPastGames :many
SELECT * FROM games
WHERE game_time <= NOW()
ORDER BY game_time DESC;

-- name: ListGamesByTeam :many
SELECT * FROM games
WHERE home_team_id = $1 OR away_team_id = $1
ORDER BY game_time;

-- name: UpdateGame :exec
UPDATE games
SET home_team_id = $1, away_team_id = $2, game_time = $3, updated_at = NOW()
WHERE id = $4;

-- name: UpdateGameTime :exec
UPDATE games
SET game_time = $1, updated_at = NOW()
WHERE id = $2;

-- name: DeleteGame :exec
DELETE FROM games
WHERE id = $1;

-- name: GetGameWithTeams :one
SELECT g.*, 
       ht.name as home_team_name, ht.wins as home_team_wins, ht.losses as home_team_losses,
       at.name as away_team_name, at.wins as away_team_wins, at.losses as away_team_losses
FROM games g
INNER JOIN teams ht ON g.home_team_id = ht.id
INNER JOIN teams at ON g.away_team_id = at.id
WHERE g.id = $1;

-- name: ListGamesWithTeams :many
SELECT g.*, 
       ht.name as home_team_name,
       at.name as away_team_name
FROM games g
INNER JOIN teams ht ON g.home_team_id = ht.id
INNER JOIN teams at ON g.away_team_id = at.id
ORDER BY g.game_time;

-- name: GetGameWithResult :one
SELECT g.*, gr.home_score, gr.away_score, gr.winning_team_id, gr.recorded_at
FROM games g
LEFT JOIN game_results gr ON g.id = gr.game_id
WHERE g.id = $1;

-- name: ListGamesWithResults :many
SELECT g.*, 
       ht.name as home_team_name,
       at.name as away_team_name,
       gr.home_score, gr.away_score, gr.winning_team_id
FROM games g
INNER JOIN teams ht ON g.home_team_id = ht.id
INNER JOIN teams at ON g.away_team_id = at.id
LEFT JOIN game_results gr ON g.id = gr.game_id
ORDER BY g.game_time DESC;

-- name: ListTeamSchedule :many
SELECT g.*, 
       ht.name as home_team_name,
       at.name as away_team_name,
       CASE 
           WHEN g.home_team_id = $1 THEN 'HOME'
           WHEN g.away_team_id = $1 THEN 'AWAY'
       END as team_location
FROM games g
INNER JOIN teams ht ON g.home_team_id = ht.id
INNER JOIN teams at ON g.away_team_id = at.id
WHERE g.home_team_id = $1 OR g.away_team_id = $1
ORDER BY g.game_time;
