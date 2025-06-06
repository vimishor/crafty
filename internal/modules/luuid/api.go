package luuid

import (
	"github.com/google/uuid"
	lua "github.com/yuin/gopher-lua"
)

func createV1Fn(l *lua.LState) int {
	t, err := uuid.NewUUID()
	if err != nil {
		l.Push(lua.LString(""))
		l.Push(lua.LString(err.Error()))
		return 2
	}

	l.Push(lua.LString(t.String()))
	l.Push(lua.LNil)
	return 2
}

func createV4Fn(l *lua.LState) int {
	t, err := uuid.NewRandom()
	if err != nil {
		l.Push(lua.LString(""))
		l.Push(lua.LString(err.Error()))
		return 2
	}

	l.Push(lua.LString(t.String()))
	l.Push(lua.LNil)
	return 2
}

func createV6Fn(l *lua.LState) int {
	t, err := uuid.NewV6()
	if err != nil {
		l.Push(lua.LString(""))
		l.Push(lua.LString(err.Error()))
		return 2
	}

	l.Push(lua.LString(t.String()))
	l.Push(lua.LNil)
	return 2
}

func createV7Fn(l *lua.LState) int {
	t, err := uuid.NewV7()
	if err != nil {
		l.Push(lua.LString(""))
		l.Push(lua.LString(err.Error()))
		return 2
	}

	l.Push(lua.LString(t.String()))
	l.Push(lua.LNil)
	return 2
}

func isValidFn(l *lua.LState) int {
	b := l.CheckString(1)
	err := uuid.Validate(b)
	if err != nil {
		l.Push(lua.LFalse)
		l.Push(lua.LString(err.Error()))
		return 2
	}

	l.Push(lua.LTrue)
	return 1
}
