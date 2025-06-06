package larchive

import lua "github.com/yuin/gopher-lua"

func Preload(l *lua.LState) {
	l.PreloadModule("archive", Loader)
}

func Loader(l *lua.LState) int {
	t := l.NewTable()

	l.SetFuncs(t, map[string]lua.LGFunction{
		"extract":    extractFn,
		"decompress": decompressFn,
		"pick":       pickFn,
	})

	l.Push(t)
	return 1
}
