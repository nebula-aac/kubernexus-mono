package main_test

import (
	"database/sql"
	"log/slog"
	"os"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nebula-aac/kubernexus-mono/nexus-server/internal/user"
	slogecho "github.com/samber/slog-echo"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		logger.Error("failed opening connection to sqlite: %v", err)
	}
	defer db.Close()

	// Create the users table
	_, err = db.Exec(`
		CREATE TABLE users (
			id TEXT PRIMARY KEY,
			username TEXT NOT NULL,
			email TEXT NOT NULL
		)`)
	if err != nil {
		logger.Error("failed creating users table: %v", err)
	}

	userHandler := user.Wire(db)
	if err != nil {
		logger.Error(err.Error())
	}

	e := echo.New()

	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/error", func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusInternalServerError, "I'm angry")
	})

	e.GET("/user/:username", userHandler.FetchByUsername)

	// Test fetching a user
	t.Run("Fetch User", func(t *testing.T) {
		// Create a user for testing
		createUser := func() {
			_, err := db.Exec("INSERT INTO users (id, username, email) VALUES (?, ?, ?)", "1", "testuser", "test@example.com")
			if err != nil {
				t.Fatalf("failed to create user: %v", err)
			}
		}

		createUser()

		// Send a request to fetch the user
		req := httptest.NewRequest(http.MethodGet, "/user/testuser", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		// Validate the response
		assert.Equal(t, http.StatusOK, rec.Code)

		// Optionally, you can further validate the response body or headers
		// Example: assert.Contains(t, rec.Body.String(), "ExpectedContent")
		// Example: assert.Equal(t, "application/json; charset=UTF-8", rec.Header().Get("Content-Type"))
	})
}
