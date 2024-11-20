package usecase

import (
	"context"

	"github.com/Maksim646/TestAssignment/api/server/restapi/model"
)

type Usecase struct {
	songRepository model.ISongRepository
}

func New(songRepository model.ISongRepository) model.ISongUsecase {
	return &Usecase{
		songRepository: songRepository,
	}
}

func (u *Usecase) CreateSong(ctx context.Context, songName string) error {
	return u.songRepository.CreateSong(ctx, songName)
}

func (u *Usecase) GetSongByID(ctx context.Context, songID int64) (model.Song, error) {
	return u.songRepository.GetSongByID(ctx, songID)
}

func (u *Usecase) UpdateSong(ctx context.Context, songName string, songText string, songID int64) error {
	return u.songRepository.UpdateSong(ctx, songName, songText, songID)
}
