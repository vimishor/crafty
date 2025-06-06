package luuid

import lua "github.com/yuin/gopher-lua"

const luaUuidType = "uuid"

func Preload(state *lua.LState) {
	state.PreloadModule("uuid", Loader)
}

func Loader(state *lua.LState) int {
	mt := state.NewTypeMetatable(luaUuidType)
	state.SetFuncs(mt, map[string]lua.LGFunction{
		"__call": createV4Fn,
	})

	t := state.NewTable()
	t.Metatable = mt

	state.SetFuncs(t, map[string]lua.LGFunction{
		"V1":       createV1Fn,
		"V4":       createV4Fn,
		"V6":       createV6Fn,
		"V7":       createV7Fn,
		"is_valid": isValidFn,
	})

	state.Push(t)
	return 1
}
