package lfs

import lua "github.com/yuin/gopher-lua"

func Preload(l *lua.LState) {
	l.PreloadModule("fs", Loader)
}

func Loader(l *lua.LState) int {
	t := l.NewTable()

	l.SetFuncs(t, map[string]lua.LGFunction{
		"is_file":   isFileFn,
		"copy_file": copyFileFn,
		"move_file": moveFileFn,
		"is_dir":    isDirFn,
		"copy_dir":  copyDirFn,
		"list_dir":  listDirFn,
		"mkdir":     mkDirFn,
		"chmod":     chmodFn,
	})

	l.Push(t)
	return 1
}
