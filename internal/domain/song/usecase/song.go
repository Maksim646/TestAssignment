package usecase

import (
	"context"

	"github.com/Maksim646/TestAssignment/internal/model"
)

type Usecase struct {
	songRepository model.ISongRepository
}

func New(songRepository model.ISongRepository) model.ISongUsecase {
	return &Usecase{
		songRepository: songRepository,
	}
}

func (u *Usecase) CreateSong(ctx context.Context, group string, song string) (int64, error) {
	return u.songRepository.CreateSong(ctx, group, song)
}

func (u *Usecase) GetSongByName(ctx context.Context, song string) (model.Song, error) {
	return u.songRepository.GetSongByName(ctx, song)
}

func (u *Usecase) GetSongByID(ctx context.Context, songID int64) (model.Song, error) {
	return u.songRepository.GetSongByID(ctx, songID)
}

func (u *Usecase) GetSongsByFilters(ctx context.Context, offset int64, limit int64, sortParams string, filters map[string]interface{}) ([]model.Song, int64, error) {
	return u.songRepository.GetSongsByFilters(ctx, offset, limit, sortParams, filters)
}

func (u *Usecase) GetCountByFilters(ctx context.Context, filters map[string]interface{}) (int64, error) {
	return u.songRepository.GetCountByFilters(ctx, filters)
}

func (u *Usecase) UpdateSong(ctx context.Context, song model.Song) error {
	return u.songRepository.UpdateSong(ctx, song)
}

func (u *Usecase) DeleteSongByID(ctx context.Context, songID int64) error {
	return u.songRepository.DeleteSongByID(ctx, songID)
}
