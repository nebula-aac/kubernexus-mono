package domain

import (
	"context"
	"encoding/gob"

	"github.com/labstack/echo/v4"
)

// Define the user domain entities and interfaces
type (
	// User represents the core data structure used in the business logic.
	User struct {
		ID       string `param:"id" query:"id" form:"id" json:"id"`
		Username string `param:"username" query:"username" form:"username" json:"username"`
		Password string `param:"password" query:"password" form:"password" json:"password"`
		Email    string `param:"email" query:"email" form:"email" json:"email"`
	}
	// UserEntity represents the structure specifically designed for persistence.
	UserEntity struct {
		ID       string
		Username string
		Password string
		Email    string
	}
	UserRepository interface {
		FetchByUsername(ctx context.Context, username string) (*UserEntity, error)
	}
	UserService interface {
		FetchByUsername(ctx context.Context, username string) (*User, error)
	}
	UserHandler interface {
		FetchByUsername(c echo.Context) error
	}
)

func init() {
	gob.Register(&User{})
}
