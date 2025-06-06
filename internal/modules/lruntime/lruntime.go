package lruntime

import (
	"runtime"

	lua "github.com/yuin/gopher-lua"
)

func Preload(state *lua.LState) {
	state.PreloadModule("runtime", Loader)
}

func Loader(state *lua.LState) int {
	t := state.NewTable()
	state.SetField(t, "OS", lua.LString(runtime.GOOS))
	state.SetField(t, "ARCH", lua.LString(runtime.GOARCH))
	state.SetField(t, "LUA_VERSION", lua.LString(lua.LuaVersion))

	state.Push(t)
	return 1
}
