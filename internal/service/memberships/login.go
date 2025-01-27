package memberships

import (
	"context"
	"errors"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/memberships"
	"github.com/novinbukannopin/fc-simple-forum/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("invalid credentials")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msg("failed to compare password")
		return "", errors.New("invalid credentials")
	}

	token, err := jwt.CreateToken(user.ID, user.Email, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("failed to create token")
		return "", err
	}
	return token, nil
}
