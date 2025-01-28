package posts

import (
	"context"
	"github.com/novinbukannopin/fc-simple-forum/internal/configs"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error

	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error

	GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostsResponse, error)
	GetPostByID(ctx context.Context, id int64) (*posts.Post, error)

	CountLikeByPostId(ctx context.Context, postId int64) (int, error)

	GetCommentByPostId(ctx context.Context, postId int64) ([]posts.Comment, error)
}

type Service struct {
	postRepository postRepository
	cfg            *configs.Config
}

func NewService(cfg *configs.Config, postRepo postRepository) *Service {
	return &Service{
		cfg:            cfg,
		postRepository: postRepo,
	}
}
