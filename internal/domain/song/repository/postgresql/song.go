package postgresql

import (
	"context"
	"database/sql"
	"errors"

	postgresql "github.com/Maksim646/Bot/database/migrations/sql"
	"github.com/Maksim646/TestAssignment/internal/model"
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

func (r *SongRepository) GetSongBuilder() sq.SelectBuilder {
	builder := postgresql.Builder.Select(
		"songs.id",
		"songs.group",
		"songs.name",
		"songs.release_date",
		"songs.text",
		"songs.link",
	).
		From("songs")

	return builder
}

func (r *SongRepository) CreateSong(ctx context.Context, group string, name string) (int64, error) {
	query, params, err := postgresql.Builder.Insert("songs").
		Columns(
			"\"group\"",
			"name",
		).
		Values(
			group,
			name,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}
	zap.L().Debug(postgresql.BuildQuery(query, params))

	var newSongID int64
	err = r.sqalxConn.GetContext(ctx, &newSongID, query, params...)
	return newSongID, err
}

func (r *SongRepository) GetSongByName(ctx context.Context, name string) (model.Song, error) {
	var song model.Song
	query, params, err := postgresql.Builder.Select(
		"songs.id",
		"songs.group",
		"songs.release_date",
		"songs.text",
		"songs.link",
	).
		From("songs").
		Where(sq.Eq{"songs.name": name}).ToSql()
	if err != nil {
		return song, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &song, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return song, model.ErrSongNotFound
		}
	}
	return song, err
}

func (r *SongRepository) GetSongByID(ctx context.Context, songID int64) (model.Song, error) {
	var song model.Song
	query, params, err := postgresql.Builder.Select(
		"songs.id",
		"songs.name",
		"songs.group",
		"songs.release_date",
		"songs.text",
		"songs.link",
	).
		From("songs").
		Where(sq.Eq{"songs.id": songID}).ToSql()
	if err != nil {
		return song, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &song, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return song, model.ErrSongNotFound
		}
	}
	return song, err
}

func (r *SongRepository) GetSongsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.Song, int64, error) {
	songsCount, err := r.GetCountByFilters(ctx, filters)
	if err != nil {
		return []model.Song{}, songsCount, err
	}

	builder := r.GetSongBuilder()

	builder = r.ApplyFilters(builder, filters)

	query, params, err := builder.
		OrderBy(sortParams).
		Offset(uint64(offset)).
		Limit(uint64(limit)).ToSql()

	if err != nil {
		return []model.Song{}, songsCount, err
	}

	var songs []model.Song
	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.SelectContext(ctx, &songs, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return songs, songsCount, model.ErrSongsNotFound
		}
	}

	return songs, songsCount, nil
}

func (r *SongRepository) GetCountByFilters(ctx context.Context, filters map[string]interface{}) (int64, error) {
	var count int64
	builder := postgresql.Builder.Select("COUNT(id)").From("songs")
	builder = r.ApplyFilters(builder, filters)
	query, params, err := builder.ToSql()
	if err != nil {
		return count, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &count, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return count, model.ErrSongsNotFound
		}
	}
	return count, err
}

func (r *SongRepository) ApplyFilters(builder sq.SelectBuilder, filters map[string]interface{}) sq.SelectBuilder {
	for filterKey, filterValue := range filters {
		switch filterKey {
		case model.FilterSongsByName:
			if name, ok := filterValue.(string); ok {
				builder = builder.Where(sq.Eq{"songs.name": name})
			}

		case model.FilterSongsByGroup:
			if group, ok := filterValue.(string); ok {
				builder = builder.Where(sq.Eq{"songs.group": group})
			}

		case model.FilterSongsByLink:
			if link, ok := filterValue.(string); ok {
				builder = builder.Where(sq.Eq{"songs.link": link})
			}

		case model.FilterSongsByReleaseDate:
			if releaseDate, ok := filterValue.(string); ok {
				builder = builder.Where(sq.Eq{"songs.release_date": releaseDate})
			}

		case model.FilterSongsByText:
			if text, ok := filterValue.(string); ok {
				builder = builder.Where(sq.Eq{"songs.text": text})
			}

		default:
			zap.L().Info("Unknown filter", zap.String("filterKey", filterKey))
			return builder
		}
	}

	return builder
}

func (r *SongRepository) DeleteSongByID(ctx context.Context, songID int64) error {
	query, params, err := postgresql.Builder.Delete("songs").
		Where(sq.Eq{"songs.id": songID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

func (r *SongRepository) UpdateSong(ctx context.Context, song model.Song) error {
	query, params, err := postgresql.Builder.Update("songs").
		Set("name", song.Name.String).
		Set("\"group\"", song.Group.String).
		Set("link", song.Link.String).
		Set("text", song.Text.String).
		Set("release_date", song.ReleaseDate).
		Where(sq.Eq{"id": song.ID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}
