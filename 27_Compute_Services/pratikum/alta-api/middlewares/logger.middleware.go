package middlewares

import (
	"io"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	// create log file if not exist and append to it if exist
	logFile, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
			log.Fatal(err)
	}
	// set log output to log file and console
	logOutput := io.MultiWriter(os.Stdout, logFile)
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${remote_ip} | ${method} ${status} ${uri} ${latency_human} | ${error}\n",
		Output: logOutput,
	})
}