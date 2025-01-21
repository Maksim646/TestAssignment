package model

import (
	"context"
	"database/sql"
)

const (
	CreateSongOK = "song added successfully"
)

type Song struct {
	ID    sql.NullInt32  `db:"id"`
	Group sql.NullString `db:"group"`
	Song  sql.NullString `db:"song"`
}

type ISongRepository interface {
	CreateSong(ctx context.Context, group string, song string) error
	GetSongByID(ctx context.Context, songID int64) (Song, error)
	UpdateSong(ctx context.Context, songName string, songText string, songID int64) error
}

type ISongUsecase interface {
	CreateSong(ctx context.Context, group string, song string) error
	GetSongByID(ctx context.Context, songID int64) (Song, error)
	UpdateSong(ctx context.Context, songName string, songText string, songID int64) error
}
