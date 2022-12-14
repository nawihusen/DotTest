package main

import (
	"dot/config"
	"dot/factory"
	"dot/migrations"
	"dot/utils/database"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := database.InitDBmySql(cfg)

	migrations.InitMigrate(db)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.SERVER_PORT)))
}
