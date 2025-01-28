package memberships

import (
	"context"
	"github.com/novinbukannopin/fc-simple-forum/internal/configs"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/memberships"
	"time"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string, userID int64) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error

	GetRefreshToken(ctx context.Context, userId int64, now time.Time) (*memberships.RefreshTokenModel, error)
	InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error
}

type Service struct {
	membershipRepo membershipRepository
	cfg            *configs.Config
}

func NewService(cfg *configs.Config, membershipRepo membershipRepository) *Service {
	return &Service{
		cfg:            cfg,
		membershipRepo: membershipRepo,
	}
}
