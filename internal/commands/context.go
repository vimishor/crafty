package commands

import (
	"context"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var logLevel *slog.LevelVar = new(slog.LevelVar)

type CmdCtx struct {
	context.Context
	log *slog.Logger
}

func NewCmdCtx() *CmdCtx {
	ctx := context.Background()

	return &CmdCtx{
		Context: ctx,
		log: slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: logLevel,
		})),
	}
}

// CobraRunECmd describes a function that can be used as a *cobra.Command RunE, PreRunE, or PostRunE.
type CobraRunECmd func(cmd *cobra.Command, args []string) (err error)

// ChainRunE runs multiple CobraRunECmd funcs one after the other returning errors.
func (ctx *CmdCtx) ChainRunE(cmdRunEs ...CobraRunECmd) CobraRunECmd {
	return func(cmd *cobra.Command, args []string) (err error) {
		for _, cmdRunE := range cmdRunEs {
			if err = cmdRunE(cmd, args); err != nil {
				return err
			}
		}

		return nil
	}
}
