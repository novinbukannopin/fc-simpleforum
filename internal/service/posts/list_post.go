package posts

import (
	"context"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *Service) GetAllPost(ctx context.Context, pageSize, pageIncex int) (posts.GetAllPostsResponse, error) {
	limit := pageSize
	offset := pageSize * (pageIncex - 1)

	response, err := s.postRepository.GetAllPost(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error get all post from repository")
		return response, err
	}
	return response, nil
}
