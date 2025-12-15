-- name: CreateGameResult :one
INSERT INTO game_results (game_id, home_score, away_score, winning_team_id, recorded_at)
VALUES ($1, $2, $3, $4, NOW())
RETURNING *;

-- name: GetGameResultById :one
SELECT * FROM game_results WHERE id = $1;

-- name: GetGameResultByGameId :one
SELECT * FROM game_results WHERE game_id = $1;

-- name: ListGameResults :many
SELECT * FROM game_results
ORDER BY recorded_at DESC;

-- name: UpdateGameResult :exec
UPDATE game_results
SET home_score = $1, away_score = $2, winning_team_id = $3, recorded_at = NOW()
WHERE id = $4;

-- name: DeleteGameResult :exec
DELETE FROM game_results
WHERE id = $1;

-- name: GetGameResultWithTeams :one
SELECT gr.*, g.game_time,
       ht.name as home_team_name,
       at.name as away_team_name,
       wt.name as winning_team_name
FROM game_results gr
INNER JOIN games g ON gr.game_id = g.id
INNER JOIN teams ht ON g.home_team_id = ht.id
INNER JOIN teams at ON g.away_team_id = at.id
LEFT JOIN teams wt ON gr.winning_team_id = wt.id
WHERE gr.id = $1;

-- name: ListGameResultsWithTeams :many
SELECT gr.*, g.game_time,
       ht.name as home_team_name,
       at.name as away_team_name,
       wt.name as winning_team_name
FROM game_results gr
INNER JOIN games g ON gr.game_id = g.id
INNER JOIN teams ht ON g.home_team_id = ht.id
INNER JOIN teams at ON g.away_team_id = at.id
LEFT JOIN teams wt ON gr.winning_team_id = wt.id
ORDER BY gr.recorded_at DESC;

-- name: ListTeamGameResults :many
SELECT gr.*, g.game_time, g.home_team_id, g.away_team_id,
       ht.name as home_team_name,
       at.name as away_team_name,
       CASE 
           WHEN g.home_team_id = $1 THEN gr.home_score
           WHEN g.away_team_id = $1 THEN gr.away_score
       END as team_score,
       CASE 
           WHEN g.home_team_id = $1 THEN gr.away_score
           WHEN g.away_team_id = $1 THEN gr.home_score
       END as opponent_score,
       CASE 
           WHEN gr.winning_team_id = $1 THEN 'WIN'
           WHEN gr.winning_team_id IS NULL THEN 'DRAW'
           ELSE 'LOSS'
       END as result
FROM game_results gr
INNER JOIN games g ON gr.game_id = g.id
INNER JOIN teams ht ON g.home_team_id = ht.id
INNER JOIN teams at ON g.away_team_id = at.id
WHERE g.home_team_id = $1 OR g.away_team_id = $1
ORDER BY g.game_time DESC;

-- name: GetTeamRecord :one
SELECT 
    COUNT(*) as games_played,
    SUM(CASE WHEN gr.winning_team_id = $1 THEN 1 ELSE 0 END) as wins,
    SUM(CASE WHEN gr.winning_team_id IS NULL THEN 1 ELSE 0 END) as draws,
    SUM(CASE WHEN gr.winning_team_id IS NOT NULL AND gr.winning_team_id != $1 THEN 1 ELSE 0 END) as losses,
    SUM(CASE 
        WHEN g.home_team_id = $1 THEN gr.home_score
        WHEN g.away_team_id = $1 THEN gr.away_score
        ELSE 0
    END) as points_for,
    SUM(CASE 
        WHEN g.home_team_id = $1 THEN gr.away_score
        WHEN g.away_team_id = $1 THEN gr.home_score
        ELSE 0
    END) as points_against
FROM game_results gr
INNER JOIN games g ON gr.game_id = g.id
WHERE g.home_team_id = $1 OR g.away_team_id = $1;
