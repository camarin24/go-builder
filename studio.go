package gostudio

import (
	"os"
	"path/filepath"

	"github.com/camarin24/go-studio/cmd"
	"github.com/camarin24/go-studio/core"
	"github.com/spf13/cobra"
)

type GoStudio struct {
	core.App
	cmd *cobra.Command
}

func NewWithConfig() *GoStudio {
	return &GoStudio{
		App: *core.NewApp(),
		cmd: &cobra.Command{
			Use:     filepath.Base(os.Args[0]),
			Short:   "Go Studio CLI",
			Version: "v0.1",
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
		},
	}
}

func (gs *GoStudio) Start() error {
	gs.cmd.AddCommand(cmd.NewServeCommand(&gs.App))
	return gs.cmd.Execute()
}
