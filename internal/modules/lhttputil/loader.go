package lhttputil

import (
	lua "github.com/yuin/gopher-lua"
)

func Preload(l *lua.LState) {
	l.PreloadModule("http_util", Loader)
}

func Loader(l *lua.LState) int {
	t := l.NewTable()
	l.SetFuncs(t, map[string]lua.LGFunction{
		"file_download": fileDownloadFn,
	})
	l.Push(t)
	return 1
}
