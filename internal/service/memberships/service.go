package memberships

import (
	"context"
	"github.com/novinbukannopin/fc-simple-forum/internal/model/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string, userID int64) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type Service struct {
	membershipRepo membershipRepository
}

func NewService(membershipRepo membershipRepository) *Service {
	return &Service{
		membershipRepo: membershipRepo,
	}
}
