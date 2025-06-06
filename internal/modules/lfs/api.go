package lfs

import (
	"os"
	"strconv"

	lua "github.com/yuin/gopher-lua"

	"github.com/vimishor/crafty/pkg/osutil"
)

func isFileFn(l *lua.LState) int {
	fpath := l.CheckString(1)

	l.Push(lua.LBool(osutil.IsFile(fpath)))
	return 1
}

func copyFileFn(l *lua.LState) int {
	src := l.CheckString(1)
	dst := l.CheckString(2)

	// FIXME: accept options for CopyFile from LUA
	if err := osutil.CopyFile(src, dst, osutil.WithOverwrite()); err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}

	l.Push(lua.LNil)
	return 1
}

func moveFileFn(l *lua.LState) int {
	src := l.CheckString(1)
	dst := l.CheckString(2)

	if err := osutil.MoveFile(src, dst); err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}

	l.Push(lua.LNil)
	return 1
}

func isDirFn(l *lua.LState) int {
	path := l.CheckString(1)

	l.Push(lua.LBool(osutil.IsDir(path)))
	return 1
}

func copyDirFn(l *lua.LState) int {
	src := l.CheckString(1)
	dst := l.CheckString(2)

	if err := osutil.CopyDirAll(src, dst); err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}

	l.Push(lua.LNil)
	return 1
}

func listDirFn(l *lua.LState) int {
	path := l.CheckString(1)

	files, err := osutil.ListDir(path)
	if err != nil {
		l.RaiseError("%q", err.Error())
		return 0
	}

	tb := l.NewTable()
	for index, file := range files {
		tb.RawSetString(strconv.Itoa(index), lua.LString(file))
	}

	l.Push(tb)
	return 1
}

func mkDirFn(l *lua.LState) int {
	path := l.CheckString(1)
	mode := l.OptInt(2, 0o700)

	if err := os.MkdirAll(path, os.FileMode(mode)); err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}

	l.Push(lua.LNil)
	return 1
}

func chmodFn(l *lua.LState) int {
	path := l.CheckString(1)
	mode := l.CheckInt(2)

	if err := os.Chmod(path, os.FileMode(mode)); err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}

	l.Push(lua.LNil)
	return 1
}
