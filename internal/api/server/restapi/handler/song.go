package handler

import (
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

	err := h.songUsecase.CreateSong(ctx, *req.CreateSong.Group, *req.CreateSong.Song)
	if err != nil {
		zap.L().Error("error create song", zap.Error(err))
		return api.NewCreateSongInternalServerError()
	}

	return api.NewCreateSongOK().WithPayload(&definition.Error{
		Message: useful.StrPtr(model.CreateSongOK),
	})
}
