package posts

import (
	"context"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
	"time"
)

func (s *Service) CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error {
	postsHashtags := strings.Join(req.PostHashtags, ",")
	model := posts.PostModel{
		UserId:       userId,
		PostTitle:    req.PostTitle,
		PostContent:  req.PostContent,
		PostHashtags: postsHashtags,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    strconv.FormatInt(userId, 10),
		UpdatedBy:    strconv.FormatInt(userId, 10),
	}
	err := s.postRepository.CreatePost(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error create post to repository")
		return err
	}
	return nil
}
