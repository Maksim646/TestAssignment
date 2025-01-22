package handler

import (
	"context"

	"github.com/Maksim646/TestAssignment/internal/api/definition"
	"github.com/Maksim646/TestAssignment/internal/api/server/restapi/api"
	"github.com/Maksim646/TestAssignment/internal/model"
	"github.com/Maksim646/TestAssignment/pkg/useful"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) GetSongsHandler(req api.GetSongsParams) middleware.Responder {
	zap.L().Info("get songs request")
	ctx := req.HTTPRequest.Context()

	filters := make(map[string]interface{})

	if req.FilterSongByGroup != nil {
		filters[model.FilterSongsByGroup] = *req.FilterSongByGroup
	}

	if req.FilterSongByName != nil {
		filters[model.FilterSongsByName] = *req.FilterSongByName
	}

	if req.FilterSongByLink != nil {
		filters[model.FilterSongsByLink] = *req.FilterSongByLink
	}

	if req.FilterSongByText != nil {
		filters[model.FilterSongsByText] = *req.FilterSongByText
	}

	if req.FilterSongByReleaseDate != nil {
		filters[model.FilterSongsByReleaseDate] = *req.FilterSongByReleaseDate
	}

	songs, _, err := h.songUsecase.GetSongsByFilters(ctx, req.Offset, req.Limit, model.DefaultSortParams, filters)
	if err != nil {
		zap.L().Error("error fetch songs", zap.Error(err))
		return api.NewGetSongsInternalServerError()
	}

	resultSongs := h.SongsToDefinition(ctx, songs)

	return api.NewGetSongsOK().WithPayload(&definition.Songs{
		Count: useful.Int64Ptr(int64(len(resultSongs))),
		Songs: resultSongs,
	})

}

func (h *Handler) SongsToDefinition(ctx context.Context, songs []model.Song) []*definition.Song {
	songsData := make([]*definition.Song, len(songs))

	for i, _ := range songs {
		songsData[i] = &definition.Song{
			ID:          songs[i].ID,
			Name:        songs[i].Name.String,
			Group:       songs[i].Group.String,
			Link:        &songs[i].Link.String,
			Text:        &songs[i].Text.String,
			ReleaseDate: &songs[i].ReleaseDate.String,
		}
	}
	return songsData
}
