package user

import (
	"context"
	"database/sql"

	"github.com/nebula-aac/kubernexus-mono/nexus-server/internal/domain"
)

type repository struct {
	db *sql.DB
}

// Retrieve username from database
func (r *repository) FetchByUsername(ctx context.Context, username string) (*domain.UserEntity, error) {
	// Implement the logic to fetch a user by username from the database.
	// For example:
	row := r.db.QueryRowContext(ctx, "SELECT id, username, email FROM users WHERE username = ?", username)

	userEntity := &domain.UserEntity{}
	err := row.Scan(&userEntity.ID, &userEntity.Username, &userEntity.Email)
	if err != nil {
		// Handle the error, e.g., log it and return an error response.
		return nil, err
	}

	return userEntity, nil
}
