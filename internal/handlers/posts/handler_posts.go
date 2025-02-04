package posts

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/novinbukannopin/fc-simple-forum/internal/middleware"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/posts"
)

type postService interface {
	CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postId, userId int64, request posts.CreateCommentRequest) error
	UpsertUserActivity(ctx context.Context, postId, userId int64, request posts.UserActivityRequest) error

	GetAllPost(ctx context.Context, pageSize, pageIncex int) (posts.GetAllPostsResponse, error)
	GetPostById(ctx context.Context, id int64) (*posts.GetPostResponse, error)
}

type Handler struct {
	*gin.Engine

	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create-post", h.CreatePost)
	route.POST("/create-comment/:postId", h.CreateComment)

	route.PUT("/user-activity/:postId", h.UpsertUserActivities)

	route.GET("/", h.GetAllPost)

	route.GET("/:postId", h.GetPostById)
}
