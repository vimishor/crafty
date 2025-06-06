package lbase64

import (
	"encoding/base64"

	lua "github.com/yuin/gopher-lua"
)

func stdEncodeFn(l *lua.LState) int {
	s := l.CheckString(1)
	b := doEncodeString(s)
	l.Push(lua.LString(b))
	return 1
}

func doEncodeString(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func stdDecodeFn(l *lua.LState) int {
	s := l.CheckString(1)
	b, err := doDecodeString(s)
	if err != nil {
		l.Push(lua.LString(""))
		l.Push(lua.LString(err.Error()))
		return 2
	}

	l.Push(lua.LString(b))
	l.Push(lua.LNil)
	return 2
}

func doDecodeString(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func urlEncodeFn(l *lua.LState) int {
	s := l.CheckString(1)
	b := doEncodeUrl(s)
	l.Push(lua.LString(b))
	return 1
}

func doEncodeUrl(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

func urlDecodeFn(l *lua.LState) int {
	s := l.CheckString(1)
	b, err := doDecodeUrl(s)
	if err != nil {
		l.Push(lua.LString(""))
		l.Push(lua.LString(err.Error()))
		return 2
	}

	l.Push(lua.LString(b))
	l.Push(lua.LNil)
	return 2
}

func doDecodeUrl(data string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(data)
}
