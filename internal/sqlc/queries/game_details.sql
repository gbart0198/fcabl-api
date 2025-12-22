-- name: ListGameDetailsById :many
SELECT * from game_details
where id = $1;

-- name: ListGameDetailsByGameId :many
SELECT * from game_details
WHERE game_id = $1;

-- name: ListGameDetailsByTeamId :many
SELECT *
FROM game_details
WHERE PLAYER_ID IN 
  (SELECT id FROM players WHERE team_id = $1);

-- name: ListGameDetails :many
SELECT * FROM game_details;

-- name: ListGameDetailsVerbose :many
SELECT gd.player_id, gd.game_id, p.team_id, u.first_name, u.last_name, p.jersey_number, gd.score
FROM game_details as gd
INNER JOIN players as p ON gd.player_id = p.id
INNER JOIN users as u on u.id = p.user_id
order by game_id, team_id;

-- name: CreateGameDetails :one
INSERT INTO game_details 
(game_id, player_id, score)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateGameDetailsScore :exec
UPDATE game_details
SET score = $3
WHERE game_id = $1 AND player_id = $2;

-- name: DeleteGameDetailsByGameAndPlayer :exec
DELETE FROM game_details
WHERE game_id = $1 AND player_id = $2;

-- name: DeleteGameDetailsByGame :exec
DELETE FROM game_details
WHERE game_id = $1;

-- name: DeleteGameDetailsByPlayer :exec
DELETE FROM game_details
WHERE player_id = $1;

-- name: DeleteGameDetails :exec
DELETE FROM game_details
WHERE id = $1;

-- name: DeleteGameDetailsByTeam :exec
DELETE FROM game_details
WHERE player_id IN
  ( SELECT PLAYER_ID FROM players WHERE team_id = $1 );
