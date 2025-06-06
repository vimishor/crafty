package commands

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"

	"github.com/vimishor/crafty/internal/modules"
)

var (
	configDir string
	cacheDir  string
	dataDir   string
)

func newRunCmd(ctx *CmdCtx) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "run <file>",
		Short: "Execute a lua file",
		Args:  cobra.MinimumNArgs(1),
		RunE:  ctx.doRunE,
	}

	cmd.LocalFlags().StringVar(&configDir, "config-dir", fmt.Sprintf("%s/crafty", xdg.ConfigHome), "Path to config directory")
	cmd.LocalFlags().StringVar(&cacheDir, "cache-dir", fmt.Sprintf("%s/crafty", xdg.CacheHome), "Path to cache directory")
	cmd.LocalFlags().StringVar(&dataDir, "data-dir", fmt.Sprintf("%s/crafty", xdg.DataHome), "Path to data directory")

	return cmd
}

func (ctx *CmdCtx) doRunE(cmd *cobra.Command, args []string) error {
	// create cacheDir
	if err := os.MkdirAll(cacheDir, os.ModePerm); err != nil {
		return fmt.Errorf("can't create cache directory %s: %q", cacheDir, err)
	}

	usr, err := user.Current()
	if err != nil {
		return err
	}

	// init LUA VM
	L := lua.NewState()
	defer L.Close()

	// register globals
	L.SetGlobal("USER", lua.LString(usr.Username))
	L.SetGlobal("HOME", lua.LString(usr.HomeDir))
	L.SetGlobal("XDG_CONFIG_HOME", lua.LString(configDir))
	L.SetGlobal("XDG_CACHE_HOME", lua.LString(cacheDir))
	L.SetGlobal("XDG_DATA_HOME", lua.LString(dataDir))

	// preload modules
	modules.PreloadAll(L)

	if len(args) > 0 && args[0] != "-" {
		// execute input file
		if err := L.DoFile(args[0]); err != nil {
			return fmt.Errorf("error executing %s: %v", args[0], err)
		}
	} else {
		script, err := scriptFromStdin()
		if err != nil {
			return err
		}

		if err := L.DoString(script); err != nil {
			return err
		}
	}

	return nil
}

func scriptFromStdin() (string, error) {
	var script []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Ignore shebang
		if strings.HasPrefix(scanner.Text(), "#") {
			continue
		}
		script = append(script, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return strings.Join(script, "\n"), nil
}
