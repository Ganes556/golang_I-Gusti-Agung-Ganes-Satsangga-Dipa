package helpers

import "github.com/labstack/echo/v4"

func ResponseSuccess(ctx echo.Context, statusCode int, message string, data interface{}) error {
	return ctx.JSON(statusCode, map[string]interface{}{
		"message": message,
		"data":    data,
	})
}

func ResponseSuccessWithoutData(ctx echo.Context, statusCode int, message string) error {
	return ctx.JSON(statusCode, map[string]interface{}{
		"message": message,
	})
}

func ResponseError(ctx echo.Context, statusCode int, err error) error {
	return ctx.JSON(statusCode, map[string]interface{}{
		"message": err.Error(),
	})
}