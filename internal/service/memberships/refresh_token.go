package memberships

import (
	"context"
	"errors"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/memberships"
	"github.com/novinbukannopin/fc-simple-forum/pkg/jwt"
	"github.com/rs/zerolog/log"
	"time"
)

func (s *Service) ValidateRefreshToken(ctx context.Context, userId int64, request memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, userId, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get refresh token from repository")
		return "", err
	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token has been expired")
	}

	if existingRefreshToken.RefreshToken != request.Token {
		return "", errors.New("invalid refresh token")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userId)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	token, err := jwt.CreateToken(user.ID, user.Email, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("failed to create token")
		return "", err
	}
	
	return token, nil
}
