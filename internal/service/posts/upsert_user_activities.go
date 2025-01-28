package posts

import (
	"context"
	"errors"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

func (s *Service) UpsertUserActivity(ctx context.Context, postId, userId int64, request posts.UserActivityRequest) error {
	now := time.Now()
	model := posts.UserActivityModel{
		PostId:    postId,
		UserId:    userId,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userId, 10),
		UpdatedBy: strconv.FormatInt(userId, 10),
	}
	userActivity, err := s.postRepository.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error get user activity from repository")
		return err
	}

	if userActivity == nil {
		if !request.IsLiked {
			return errors.New("anda belum menyukai post ini")
		}
		err = s.postRepository.CreateUserActivity(ctx, model)
	} else {
		err = s.postRepository.UpdateUserActivity(ctx, model)
	}

	if err != nil {
		log.Error().Err(err).Msg("error upsert user activity to repository")
		return err
	}
	return nil
}
