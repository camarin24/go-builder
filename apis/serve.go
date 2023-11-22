package apis

import (
	"log"

	"github.com/camarin24/go-studio/core"
	"github.com/fatih/color"
)

type ServerConfig struct {
	AllowedOrigins []string
	HttpAddr       string
}

func Server(app *core.App, config ServerConfig) error {
	if len(config.AllowedOrigins) == 0 {
		config.AllowedOrigins = []string{"*"}
	}

	// Migrations here

	server, err := InitApp(app)
	if err != nil {
		color.Red(err.Error())
		log.Fatal(err)
	}

	//TODO: Graceful shutdown
	//TODO: WTF is graceful shutdown

	return server.Start(config.HttpAddr)
}
