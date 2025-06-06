package larchive

import (
	"io"
	"io/fs"
	"os"
	"regexp"

	"github.com/mholt/archives"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

func pickFn(l *lua.LState) int {
	src := l.CheckString(1)
	inputFilters := l.CheckTable(2)
	var filters map[string]string

	mapper := gluamapper.NewMapper(gluamapper.Option{
		NameFunc: gluamapper.Id,
	})

	if err := mapper.Map(inputFilters, &filters); err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}

	if err := mountArchive(src, filters); err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}

	l.Push(lua.LNil)
	return 1
}

func mountArchive(src string, filters map[string]string) error {
	fsys := &archives.DeepFS{Root: src}

	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if path == "." {
			return nil
		}

		for filter, dst := range filters {
			// 1. check for exact match
			if path == filter {
				err := copyFile(path, dst, fsys)
				if err != nil {
					return err
				}
				continue
			}

			// 2. regex match
			match, _ := regexp.MatchString(filter, path)
			if match {
				err := copyFile(path, dst, fsys)
				if err != nil {
					return err
				}
				continue
			}
		}

		return nil
	})
}

func copyFile(src, dst string, fsys fs.FS) error {
	reader, err := fsys.Open(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	stat, err := reader.Stat()
	if err != nil {
		return err
	}

	writer, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, stat.Mode()&os.ModePerm)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, reader)
	return err
}
