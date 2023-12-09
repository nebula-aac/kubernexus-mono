package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nebula-aac/kubernexus-mono/nexus-server/internal/domain"
)

type handler struct {
	svc domain.UserService
}

func NewHandler() domain.IHandler {
	return &handler{}
}

func (h *handler) FetchByUsername(c echo.Context) error {
	username := c.Param("username")

	// Fetch user from the service
	user, err := h.svc.FetchByUsername(c.Request().Context(), username)
	if err != nil {
		// Handle the error
		if user == nil {
			// If user is nil, it means user not found
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}

		// Handle other errors (log, etc.)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	// Return the user in the response
	return c.JSON(http.StatusOK, user)
}
