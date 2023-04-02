package utils

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Id2Int(idStr string, idPtr *int) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}
	*idPtr = id
	return nil
}