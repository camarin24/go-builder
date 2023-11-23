package cmd

import (
	"log"

	"github.com/camarin24/go-studio/apis"
	"github.com/camarin24/go-studio/core"
	"github.com/spf13/cobra"
)

func NewServeCommand(app *core.App) *cobra.Command {
	var httpAddr string
	command := cobra.Command{
		Use:   "serve [domain(s)]",
		Args:  cobra.ArbitraryArgs,
		Short: "Starts the web server (default to 127.0.0.1:8051 if no domain is specified)",
		Run: func(command *cobra.Command, args []string) {
			// set default listener addresses if at least one domain is specified
			if len(args) > 0 {
				if httpAddr == "" {
					httpAddr = "0.0.0.0:80"
				}
			} else {
				if httpAddr == "" {
					httpAddr = "127.0.0.1:8051"
				}
			}

			err := apis.Server(app, apis.ServerConfig{
				HttpAddr: httpAddr,
			})

			if err != nil {
				log.Fatalln(err)
			}
		},
	}

	command.LocalFlags().StringVar(
		&httpAddr,
		"http",
		"",
		"TCP address to listen for the HTTP server\n(if domain args are specified - default to 0.0.0.0:80, otherwise - default to 127.0.0.1:8090)",
	)

	return &command
}
