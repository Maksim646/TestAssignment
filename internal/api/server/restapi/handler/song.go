package handler

import (
	"fmt"
	_ "fmt"
	"strings"

	"github.com/Maksim646/TestAssignment/internal/api/definition"
	"github.com/Maksim646/TestAssignment/internal/api/server/restapi/api"
	"github.com/Maksim646/TestAssignment/internal/model"
	"github.com/Maksim646/TestAssignment/pkg/useful"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) CreateSongHandler(req api.CreateSongParams) middleware.Responder {
	zap.L().Info("create new song request")
	ctx := req.HTTPRequest.Context()

	filters := make(map[string]interface{})
	filters[model.FilterSongsByName] = *req.CreateSong.Song
	filters[model.FilterSongsByGroup] = *req.CreateSong.Group

	limit, err := h.songUsecase.GetCountByFilters(ctx, filters)
	if err != nil {
		zap.L().Error("error count songs", zap.Error(err))
		return api.NewCreateSongInternalServerError()
	}

	if limit == 0 {
		_, err = h.songUsecase.CreateSong(ctx, *req.CreateSong.Group, *req.CreateSong.Song)
		if err != nil {
			zap.L().Error("error create song", zap.Error(err))
			return api.NewCreateSongInternalServerError()
		}

		return api.NewCreateSongOK().WithPayload(&definition.Error{
			Message: useful.StrPtr(model.CreateSongOK),
		})
	} else {
		return api.NewCreateSongBadRequest().WithPayload(&definition.Error{
			Message: &model.SongAlreadyExists,
		})
	}

}

func (h *Handler) DeleteSongHandler(req api.DeleteSongParams) middleware.Responder {
	zap.L().Info("delete song request")
	ctx := req.HTTPRequest.Context()

	_, err := h.songUsecase.GetSongByID(ctx, req.ID)
	if err != nil {
		return api.NewDeleteSongBadRequest().WithPayload(&definition.Error{
			Message: &model.SongNotFound,
		})
	}

	err = h.songUsecase.DeleteSongByID(ctx, req.ID)
	if err != nil {
		zap.L().Error("error delete song", zap.Error(err))
		return api.NewDeleteSongInternalServerError()
	}

	return api.NewDeleteSongOK().WithPayload(&definition.Error{
		Message: &model.SongWasDeleted,
	})
}

func (h *Handler) UpdateSongHAndler(req api.UpdateSongParams) middleware.Responder {
	zap.L().Info("update song request")
	ctx := req.HTTPRequest.Context()

	fmt.Println(req.ID)

	song, err := h.songUsecase.GetSongByID(ctx, req.ID)
	if err != nil {
		return api.NewUpdateSongBadRequest().WithPayload(&definition.Error{
			Message: &model.SongNotFound,
		})
	}

	fmt.Println(song.ID, song.Name)

	if req.UpdateSong.Name != "" {
		song.Name.String = req.UpdateSong.Name
	}

	if req.UpdateSong.Group != "" {
		song.Group.String = req.UpdateSong.Group
	}

	if req.UpdateSong.Link != "" {
		song.Link.String = req.UpdateSong.Link
	}

	if req.UpdateSong.ReleaseDate != "" {
		song.ReleaseDate.String = req.UpdateSong.ReleaseDate
	}

	if req.UpdateSong.Text != "" {
		song.Text.String = req.UpdateSong.Text
	}

	fmt.Println(song.ID)

	err = h.songUsecase.UpdateSong(ctx, song)
	if err != nil {
		zap.L().Error("error update song", zap.Error(err))
		return api.NewUpdateSongInternalServerError()
	}

	return api.NewUpdateSongOK().WithPayload(&definition.Error{
		Message: &model.SongWasUpdated,
	})
}

func (h *Handler) GetSongVerseHandler(req api.GetSongVerseParams) middleware.Responder {
	zap.L().Info("get song verse request")
	ctx := req.HTTPRequest.Context()

	song, err := h.songUsecase.GetSongByID(ctx, req.ID)
	if err != nil {
		return api.NewCreateSongBadRequest().WithPayload(&definition.Error{
			Message: &model.SongNotFound,
		})
	}

	if !song.Text.Valid {
		return api.NewCreateSongBadRequest().WithPayload(&definition.Error{
			Message: useful.StrPtr("Song text is empty"),
		})
	}

	verses := strings.Split(song.Text.String, "\n\n")

	for i := range verses {
		verses[i] = strings.TrimSpace(verses[i])
	}

	start := req.Offset
	end := start + req.Limit

	if start >= int64(len(verses)) {
		return api.NewCreateSongBadRequest().WithPayload(&definition.Error{
			Message: useful.StrPtr("Offset exceeds the number of verses"),
		})
	}

	if end > int64(len(verses)) {
		end = int64(len(verses))
	}

	var songVerses []*definition.SongVerse
	for i := start; i < end; i++ {
		songVerses = append(songVerses, &definition.SongVerse{
			Verse: &verses[i],
		})
	}

	return api.NewGetSongVerseOK().WithPayload(&definition.SongVersesResponse{
		Verses: songVerses,
	})
}
