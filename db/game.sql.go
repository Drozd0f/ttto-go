// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: game.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/tabbed/pqtype"
)

const createGame = `-- name: CreateGame :one
WITH g AS ( INSERT INTO games (id, owner_id) VALUES ($1, $2) RETURNING id, owner_id, opponent_id, current_player_id, step_count, winner_id, field, current_state )
SELECT
    g.id, g.owner_id, g.opponent_id, g.current_player_id, g.step_count, g.winner_id, g.field, g.current_state,
    ow.username AS owner_name,
    op.username AS opponent_name,
    win.username AS winner_name,
    cp.username AS current_player_name
FROM g
LEFT JOIN users ow ON ow.id = g.owner_id
LEFT JOIN users op ON op.id = g.opponent_id
LEFT JOIN users win ON win.id = g.winner_id
LEFT JOIN users cp ON cp.id = g.current_player_id
`

type CreateGameParams struct {
	ID      uuid.UUID `json:"id"`
	OwnerID uuid.UUID `json:"owner_id"`
}

type CreateGameRow struct {
	ID                uuid.UUID             `json:"id"`
	OwnerID           uuid.UUID             `json:"owner_id"`
	OpponentID        uuid.NullUUID         `json:"opponent_id"`
	CurrentPlayerID   uuid.NullUUID         `json:"current_player_id"`
	StepCount         int32                 `json:"step_count"`
	WinnerID          uuid.NullUUID         `json:"winner_id"`
	Field             pqtype.NullRawMessage `json:"field"`
	CurrentState      int16                 `json:"current_state"`
	OwnerName         sql.NullString        `json:"owner_name"`
	OpponentName      sql.NullString        `json:"opponent_name"`
	WinnerName        sql.NullString        `json:"winner_name"`
	CurrentPlayerName sql.NullString        `json:"current_player_name"`
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) (CreateGameRow, error) {
	row := q.queryRow(ctx, q.createGameStmt, createGame, arg.ID, arg.OwnerID)
	var i CreateGameRow
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.OpponentID,
		&i.CurrentPlayerID,
		&i.StepCount,
		&i.WinnerID,
		&i.Field,
		&i.CurrentState,
		&i.OwnerName,
		&i.OpponentName,
		&i.WinnerName,
		&i.CurrentPlayerName,
	)
	return i, err
}

const getGameById = `-- name: GetGameById :one
SELECT id, owner_id, opponent_id, current_player_id, step_count, winner_id, field, current_state, owner_name, opponent_name, current_player_name, winner_name FROM games_with_usernames WHERE id=$1
`

func (q *Queries) GetGameById(ctx context.Context, id uuid.UUID) (GamesWithUsername, error) {
	row := q.queryRow(ctx, q.getGameByIdStmt, getGameById, id)
	var i GamesWithUsername
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.OpponentID,
		&i.CurrentPlayerID,
		&i.StepCount,
		&i.WinnerID,
		&i.Field,
		&i.CurrentState,
		&i.OwnerName,
		&i.OpponentName,
		&i.CurrentPlayerName,
		&i.WinnerName,
	)
	return i, err
}

const getGames = `-- name: GetGames :many
SELECT id, owner_id, opponent_id, current_player_id, step_count, winner_id, field, current_state, owner_name, opponent_name, current_player_name, winner_name
FROM games_with_usernames
ORDER BY id LIMIT $2 OFFSET $1
`

type GetGamesParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) GetGames(ctx context.Context, arg GetGamesParams) ([]GamesWithUsername, error) {
	rows, err := q.query(ctx, q.getGamesStmt, getGames, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GamesWithUsername
	for rows.Next() {
		var i GamesWithUsername
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.OpponentID,
			&i.CurrentPlayerID,
			&i.StepCount,
			&i.WinnerID,
			&i.Field,
			&i.CurrentState,
			&i.OwnerName,
			&i.OpponentName,
			&i.CurrentPlayerName,
			&i.WinnerName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTotalGames = `-- name: GetTotalGames :one
SELECT count(id) FROM games
`

func (q *Queries) GetTotalGames(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.getTotalGamesStmt, getTotalGames)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateGameById = `-- name: UpdateGameById :exec
UPDATE games
SET
    owner_id=$2,
    opponent_id=$3,
    current_player_id=$4,
    step_count=$5,
    winner_id=$6,
    field=$7,
    current_state=$8
WHERE id=$1
`

type UpdateGameByIdParams struct {
	ID              uuid.UUID             `json:"id"`
	OwnerID         uuid.UUID             `json:"owner_id"`
	OpponentID      uuid.NullUUID         `json:"opponent_id"`
	CurrentPlayerID uuid.NullUUID         `json:"current_player_id"`
	StepCount       int32                 `json:"step_count"`
	WinnerID        uuid.NullUUID         `json:"winner_id"`
	Field           pqtype.NullRawMessage `json:"field"`
	CurrentState    int16                 `json:"current_state"`
}

func (q *Queries) UpdateGameById(ctx context.Context, arg UpdateGameByIdParams) error {
	_, err := q.exec(ctx, q.updateGameByIdStmt, updateGameById,
		arg.ID,
		arg.OwnerID,
		arg.OpponentID,
		arg.CurrentPlayerID,
		arg.StepCount,
		arg.WinnerID,
		arg.Field,
		arg.CurrentState,
	)
	return err
}