package larchive

import (
	"context"
	"io"
	"os"

	"github.com/mholt/archives"
	lua "github.com/yuin/gopher-lua"
)

func decompressFn(l *lua.LState) int {
	src := l.CheckString(1)
	dst := l.CheckString(2)

	if err := doDecompress(src, dst); err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}

	l.Push(lua.LNil)
	return 0
}

func doDecompress(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	ctx := context.Background()
	format, r, err := archives.Identify(ctx, src, f)
	if err != nil {
		return err
	}

	reader, err := format.(archives.Decompressor).OpenReader(r)
	if err != nil {
		return err
	}
	defer reader.Close()

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	writer, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, stat.Mode()&os.ModePerm)
	if err != nil {
		return err
	}
	defer writer.Close()

	if _, err := io.Copy(writer, reader); err != nil {
		return err
	}

	return nil
}
