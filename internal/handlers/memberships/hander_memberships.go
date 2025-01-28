package memberships

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/novinbukannopin/fc-simple-forum/internal/middleware"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/memberships"
)

type memberShipService interface {
	SignUp(c context.Context, req memberships.SignUpRequest) error
	Login(c context.Context, req memberships.LoginRequest) (string, string, error)

	ValidateRefreshToken(ctx context.Context, userId int64, request memberships.RefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine

	membershipSvc memberShipService
}

func NewHandler(api *gin.Engine, membershipSvc memberShipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/memberships")
	route.GET("/ping", h.ping)
	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)

	routeRefresh := route.Group("/memberships")
	routeRefresh.Use(middleware.AuthRefreshMiddleware())
	routeRefresh.POST("/refresh", h.Refresh)
}
