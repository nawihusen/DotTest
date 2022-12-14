package helpers

import "github.com/labstack/echo/v4"

func ErrorBind(c echo.Context) error {
	return c.JSON(400, FailedResponseHelper("Failed To Bind Data"))
}

func ErrorInternal(c echo.Context, msg string) error {
	return c.JSON(500, FailedResponseHelper(msg))
}
