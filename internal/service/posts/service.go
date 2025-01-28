package posts

import (
	"context"
	"github.com/novinbukannopin/fc-simple-forum/internal/configs"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
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
