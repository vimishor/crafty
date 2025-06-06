package lbase64

import lua "github.com/yuin/gopher-lua"

func Preload(state *lua.LState) {
	state.PreloadModule("base64", Loader)
}

func Loader(state *lua.LState) int {
	t := state.NewTable()

	state.SetFuncs(t, map[string]lua.LGFunction{
		"encode":     stdEncodeFn,
		"encode_url": urlEncodeFn,
		"decode":     stdDecodeFn,
		"decode_url": urlDecodeFn,
	})

	state.Push(t)
	return 1
}
