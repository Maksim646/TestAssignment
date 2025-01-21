package handler

import (
	"net/http"
	"strings"

	//"strings"

	"encoding/json"

	// "github.com/Maksim646/space_vpx_satellite/internal/api/definition"
	// "github.com/Maksim646/space_vpx_satellite/internal/model"
	// "github.com/Maksim646/space_vpx_satellite/pkg/jsonwebtoken"
	"go.uber.org/zap"

	// "github.com/Maksim646/TestAssignment/internal/api/server/restapi"
	"github.com/Maksim646/TestAssignment/internal/api/server/restapi/api"
	"github.com/Maksim646/TestAssignment/internal/model"
	"github.com/go-openapi/loads"
	"github.com/go-swagger/go-swagger/examples/oauth2/restapi"
)

type Handler struct {
	songUsecase model.ISongUsecase

	router       http.Handler
	HashSalt     string
	jwtSigninKey string
}

func New(
	songUsecase model.ISongUsecase,

	version string,
	HashSalt string,
	jwtSigninKey string,
) *Handler {

	withChangedVersion := strings.ReplaceAll(string(restapi.SwaggerJSON), "development", version)
	swagger, err := loads.Analyzed(json.RawMessage(withChangedVersion), "")
	if err != nil {
		panic(err)
	}

	h := &Handler{
		songUsecase: songUsecase,

		HashSalt:     HashSalt,
		jwtSigninKey: jwtSigninKey,
	}

	zap.L().Error("server http handler request")
	router := api.NewTestAssignmentBackendServiceAPI(swagger)
	router.UseSwaggerUI()
	router.Logger = zap.S().Infof
	// router.BearerAuth = h.ValidateHeader

	h.router = router.Serve(nil)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zap.L().Error("server http request")
	if h.router == nil {
		zap.L().Error("h.router is nil")
		return
	}
	zap.L().Error("h.router is not nil")
	h.router.ServeHTTP(w, r)
}

// func (h *Handler) ValidateHeader(bearerHeader string) (*definition.Principal, error) {
// 	ctx := context.Background()

// 	bearerToken := strings.TrimPrefix(bearerHeader, "Bearer ")
// 	userID, roleID, err := jsonwebtoken.ParseToken(bearerToken, h.jwtSigninKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if roleID == 0 {
// 		_, err = h.userUsecase.GetUserByID(ctx, userID)
// 		if err != nil {
// 			return nil, err
// 		}

// 	} else {
// 		_, err = h.adminUsecase.GetAdminByID(ctx, userID)
// 		if err != nil {
// 			return nil, err
// 		}

// 	}

// 	return &definition.Principal{ID: userID, Role: roleID}, nil
// }
