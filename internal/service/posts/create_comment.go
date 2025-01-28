package posts

import (
	"context"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

func (s *Service) CreateComment(ctx context.Context, postId, userId int64, request posts.CreateCommentRequest) error {
	now := time.Now()
	comment := posts.CommentModel{
		UserId:         userId,
		PostId:         postId,
		CommentContent: request.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(userId, 10),
		UpdatedBy:      strconv.FormatInt(userId, 10),
	}

	err := s.postRepository.CreateComment(ctx, comment)
	if err != nil {
		log.Error().Err(err).Msg("error create comment to repository")
		return err
	}
	return nil
}
