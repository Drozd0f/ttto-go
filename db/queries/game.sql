-- name: CreateGame :one
WITH g AS ( INSERT INTO games (owner_id) VALUES ($1) RETURNING * )
SELECT
    g.*,
    ow.username AS owner_name,
    op.username AS opponent_name,
    win.username AS winner_name,
    cp.username AS current_player_name
FROM g
LEFT JOIN users ow ON ow.id = g.owner_id
LEFT JOIN users op ON op.id = g.opponent_id
LEFT JOIN users win ON win.id = g.winner_id
LEFT JOIN users cp ON cp.id = g.current_player_id;

-- name: GetGameById :one
SELECT * FROM games_with_usernames WHERE id=$1;

-- name: GetGames :many
SELECT *
FROM games_with_usernames
ORDER BY id LIMIT $2 OFFSET $1;

-- name: GetTotalGames :one
SELECT count(id) FROM games;

-- name: UpdateGameById :exec
UPDATE games
SET
    owner_id=$2,
    opponent_id=$3,
    current_player_id=$4,
    step_count=$5,
    winner_id=$6,
    field=$7,
    current_state=$8
WHERE id=$1;
