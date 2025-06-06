package commands

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var verbose bool

func NewRootCmd() (cmd *cobra.Command) {
	ctx := NewCmdCtx()

	cmd = &cobra.Command{
		Use:   "crafty [OPTIONS] COMMAND",
		Short: "Single executable Lua 5.1 interpreter",
		Long: `Single executable Lua 5.1 interpreter, baked with some useful modules.
		`,
		// Args: cobra.NoArgs,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if verbose {
				logLevel.Set(slog.LevelDebug)
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.RunE(newRunCmd(ctx), args)
		},
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Increase verbosity.")

	cmd.AddCommand(
		newVersionCmd(ctx),
		newRunCmd(ctx),
	)

	return cmd
}
