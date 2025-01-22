package model

import (
	"context"
	"database/sql"
	"errors"
)

const (
	CreateSongOK = "song added successfully"

	FilterSongsByName        = "filter_songs_by_name"
	FilterSongsByGroup       = "filter_songs_by_group"
	FilterSongsByLink        = "filter_songs_by_link"
	FilterSongsByText        = "filter_songs_by_text"
	FilterSongsByReleaseDate = "filter_songs_by_release_date"

	DefaultSortParams = "id desc"
)

var (
	ErrSongNotFound  = errors.New("song not found")
	ErrSongsNotFound = errors.New("songs not found")

	SongAlreadyExists = "song alredy exists"
	SongNotFound      = "song not found"
	SongWasDeleted    = "song was deleted successfully"
	SongWasUpdated    = "song was updated successfully"
)

type Song struct {
	ID int64 `db:"id"`

	Group       sql.NullString `db:"group"`
	Name        sql.NullString `db:"name"`
	ReleaseDate sql.NullString `db:"release_date"`
	Text        sql.NullString `db:"text"`
	Link        sql.NullString `db:"link"`
}

type ISongRepository interface {
	CreateSong(ctx context.Context, group string, song string) (int64, error)

	GetSongsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]Song, int64, error)
	GetCountByFilters(ctx context.Context, filters map[string]interface{}) (int64, error)
	GetSongByName(ctx context.Context, song string) (Song, error)
	GetSongByID(ctx context.Context, songID int64) (Song, error)

	UpdateSong(ctx context.Context, song Song) error

	DeleteSongByID(ctx context.Context, songID int64) error
}

type ISongUsecase interface {
	CreateSong(ctx context.Context, group string, song string) (int64, error)

	GetSongsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]Song, int64, error)
	GetCountByFilters(ctx context.Context, filters map[string]interface{}) (int64, error)
	GetSongByName(ctx context.Context, song string) (Song, error)
	GetSongByID(ctx context.Context, songID int64) (Song, error)

	UpdateSong(ctx context.Context, song Song) error

	DeleteSongByID(ctx context.Context, songID int64) error
}
