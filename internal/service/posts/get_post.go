package posts

import (
	"context"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *Service) GetPostById(ctx context.Context, postId int64) (*posts.GetPostResponse, error) {
	detail, err := s.postRepository.GetPostByID(ctx, postId)
	if err != nil {
		log.Error().Err(err).Msg("error get post by id from repository")
		return nil, err
	}

	likeCount, err := s.postRepository.CountLikeByPostId(ctx, postId)
	if err != nil {
		log.Error().Err(err).Msg("error count like by post id from repository")
		return nil, err
	}

	comment, err := s.postRepository.GetCommentByPostId(ctx, postId)
	if err != nil {
		log.Error().Err(err).Msg("error get comment by post id from repository")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID:           detail.ID,
			UserId:       detail.UserId,
			Username:     detail.Username,
			PostTitle:    detail.PostTitle,
			PostContent:  detail.PostContent,
			PostHashtags: detail.PostHashtags,
			IsLiked:      detail.IsLiked,
		},
		LikeCount: likeCount,
		Comment:   comment,
	}, nil
}
