package memberships

import (
	"context"
	"errors"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/memberships"
	"github.com/novinbukannopin/fc-simple-forum/pkg/jwt"
	tokenUtil "github.com/novinbukannopin/fc-simple-forum/pkg/token"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s *Service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("invalid credentials")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msg("failed to compare password")
		return "", "", errors.New("invalid credentials")
	}

	token, err := jwt.CreateToken(user.ID, user.Email, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("failed to create token")
		return "", "", err
	}

	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, user.ID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("failed to get refresh token")
		return "", "", err
	}

	if existingRefreshToken != nil {
		return token, existingRefreshToken.RefreshToken, nil

	}

	refreshToken := tokenUtil.GenerateRefreshToken()
	if err != nil {
		log.Error().Err(err).Msg("failed to generate refresh token")
		return token, "", errors.New("failed to generate refresh token")
	}

	err = s.membershipRepo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserId:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(time.Hour * 24 * 7),
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to insert refresh token")
		return token, refreshToken, err
	}
	return token, refreshToken, nil
}
