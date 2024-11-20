package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	postgresql "github.com/Maksim646/Bot/database/migrations/sql"
	"github.com/Maksim646/TestAssignment/api/server/restapi/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
	"go.uber.org/zap"
)

type SongRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.ISongRepository {
	return &SongRepository{sqalxConn: sqalxConn}
}

func (r *SongRepository) CreateSong(ctx context.Context, songName string) error {
	query, params, err := postgresql.Builder.Insert("songs").
		Columns(
			"song_name",
		).
		Values(
			songName,
		).
		ToSql()
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	fmt.Println(postgresql.BuildQuery(query, params))
	return err
}

func (r *SongRepository) GetSongByID(ctx context.Context, songID int64) (model.Song, error) {
	var song model.Song
	query, params, err := postgresql.Builder.Select(
		"songs.id",
		"songs.song_name",
		"songs.song_text",
	).
		From("users").
		Where(sq.Eq{"songs.id": songID}).
		ToSql()
	if err != nil {
		fmt.Println("error:", err)
		return song, err
	}
	//zap.L().Debug(postgresql.BuildQuery(query, params))

	if err = r.sqalxConn.GetContext(ctx, &song, query, params...); err != nil {
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return model.Song{}, fmt.Errorf("пользователь с ID %d не найден", songID)
			}
			return model.Song{}, fmt.Errorf("ошибка при получении пользователя: %w", err)
		}
	}
	return song, err
}

func (r *SongRepository) UpdateSong(ctx context.Context, songName string, songText string, songID int64) error {
	builder := postgresql.Builder.Update("songs")

	if songName != "" {
		builder = builder.Set("song_name", songName)
	}

	if songText != "" {
		builder = builder.Set("song_text", songText)
	}

	query, params, err := builder.
		Where(sq.Eq{"id": songID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}
