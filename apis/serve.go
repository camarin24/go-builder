package apis

import (
	"log"

	"github.com/camarin24/go-studio/core"
	"github.com/camarin24/go-studio/migrations"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4/middleware"
)

type ServerConfig struct {
	AllowedOrigins []string
	HttpAddr       string
}

func Server(app *core.App, config ServerConfig) error {
	if len(config.AllowedOrigins) == 0 {
		config.AllowedOrigins = []string{"*"}
	}

	migrations.Migrate(app)

	server, err := InitApp(app)
	if err != nil {
		color.Red(err.Error())
		log.Fatal(err)
	}

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.AllowedOrigins,
		AllowHeaders: []string{"*"},
	}))

	//TODO: Graceful shutdown
	//TODO: WTF is graceful shutdown

	return server.Start(config.HttpAddr)
}
