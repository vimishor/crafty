package modules

import (
	lua "github.com/yuin/gopher-lua"

	lhttp "github.com/vadv/gopher-lua-libs/http"
	ljson "github.com/vadv/gopher-lua-libs/json"
	ltime "github.com/vadv/gopher-lua-libs/time"

	"github.com/vimishor/crafty/internal/modules/larchive"
	"github.com/vimishor/crafty/internal/modules/lbase64"
	"github.com/vimishor/crafty/internal/modules/lfs"
	"github.com/vimishor/crafty/internal/modules/lhttputil"
	"github.com/vimishor/crafty/internal/modules/lruntime"
	"github.com/vimishor/crafty/internal/modules/luuid"
)

// PreloadAll will preload all crafty lua packages
func PreloadAll(state *lua.LState) {
	lruntime.Preload(state)
	lbase64.Preload(state)
	ltime.Preload(state)
	luuid.Preload(state)
	lfs.Preload(state)
	larchive.Preload(state)
	lhttp.Preload(state)
	lhttputil.Preload(state)
	ljson.Preload(state)
}
