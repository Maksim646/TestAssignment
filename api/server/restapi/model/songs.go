package model

import (
	"context"
	"database/sql"
)

type Song struct {
	ID       sql.NullInt32  `db:"id"`
	SongName sql.NullString `db:"song_name"`
	SongText sql.NullString `db:"song_text"`
}

type ISongRepository interface {
	CreateSong(ctx context.Context, songName string) error
	GetSongByID(ctx context.Context, songID int64) (Song, error)
	UpdateSong(ctx context.Context, songName string, songText string, songID int64) error
}

type ISongUsecase interface {
	CreateSong(ctx context.Context, songName string) error
	GetSongByID(ctx context.Context, songID int64) (Song, error)
	UpdateSong(ctx context.Context, songName string, songText string, songID int64) error
}
