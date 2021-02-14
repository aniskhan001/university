package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const EchoLogFormat = "" +
	`{"time":"${time_rfc3339_nano}", "method":"${method}", "uri":"${uri}", ` +
	`"status":${status},"latency":"${latency_human}", ` +
	`"host":"${host}", "remote_ip":"${remote_ip}", ` +
	`"user_agent":"${user_agent}", "error":"${error}"}` + "\n"

// Attach middlewares required for the application
func Attach(e *echo.Echo) error {

	// echo middlewares, todo: add color to the log
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: EchoLogFormat}))
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	return nil
}
