package domain

import "github.com/labstack/echo/v4"

type IHandler interface {
	FetchByUsername(c echo.Context) error
}
