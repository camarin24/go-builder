package utils

import (
	"github.com/camarin24/go-studio/types"
	"github.com/labstack/echo/v4"
)

func StandardError(c echo.Context, error  string) error {
	return c.JSON(500, types.DefaultErrorResponse{
		Message: error,
		Status:  500,
	})
}
