package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/vimishor/crafty/internal/version"
)

func newVersionCmd(_ *CmdCtx) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Long: `Print application version.

For a more detailed version information, --verbose flag can be used.
		`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.String())

			if verbose {
				fmt.Println("Build Date:", version.BuildDate)
				fmt.Println("Git Commit:", version.GitCommit)
				fmt.Println("Go Version:", version.GoVersion)
				fmt.Println("OS / Arch:", version.OsArch)
			}
		},
	}

	return cmd
}
