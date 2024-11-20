package handler

import (
	"time"

	"github.com/Maksim646/TestAssignment/api/server/restapi/model"
	"github.com/patrickmn/go-cache"
)

type Handler struct {
	cache       *cache.Cache
	songUsecase model.ISongUsecase
}

func New(
	version string,
	songUsecase model.ISongUsecase,
) *Handler {

	h := &Handler{
		cache:       cache.New(cache.NoExpiration, time.Minute),
		songUsecase: songUsecase,
	}
	return h
}
