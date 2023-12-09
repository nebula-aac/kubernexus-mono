package user

import (
	"context"

	"github.com/nebula-aac/kubernexus-mono/nexus-server/internal/domain"
)

type service struct {
	repo domain.UserRepository
}

// FetchByUsername fetches a user by username.
func (s *service) FetchByUsername(ctx context.Context, username string) (*domain.User, error) {
	userEntity, err := s.repo.FetchByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	// Convert UserEntity to User, assuming the structures are similar
	user := &domain.User{
		ID:       userEntity.ID,
		Username: userEntity.Username,
		Password: userEntity.Password,
		Email:    userEntity.Email,
	}

	return user, nil
}
