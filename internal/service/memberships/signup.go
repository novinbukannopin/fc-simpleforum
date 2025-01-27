package memberships

import (
	"context"
	"errors"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s *Service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username, 0)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("username or email already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	model := memberships.UserModel{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(pass),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	return s.membershipRepo.CreateUser(ctx, model)
}
