package larchive

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mholt/archives"
	lua "github.com/yuin/gopher-lua"
)

var (
	ErrAbsolutePath  = errors.New("absolute file path")
	ErrRelativePath  = errors.New("relative file path")
	ErrSymlinkInPath = errors.New("symlink in file path")
)

func extractFn(l *lua.LState) int {
	src := l.CheckString(1)
	dst := l.CheckString(2)
	defOptions := l.NewTable()
	opts := l.OptTable(3, defOptions)

	var options []ExtractOption
	if opts.RawGetString("dry-run").Type() == lua.LTBool {
		options = append(options, WithDryRun())
	}

	if opts.RawGetString("allow-symlink").Type() == lua.LTBool {
		options = append(options, WithAllowSymlink())
	}

	if err := doExtract(src, dst, options...); err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}

	l.Push(lua.LNil)
	return 0
}

func doExtract(src, dst string, opts ...ExtractOption) error {
	e := &extractor{
		DryRun:       false,
		AllowSymlink: false,
	}
	for _, opt := range opts {
		opt(e)
	}

	if len(dst) == 0 {
		return fmt.Errorf("no destination was specified")
	}

	srcRoot := filepath.Dir(src)
	srcFile := filepath.Base(src)

	f, err := os.OpenInRoot(srcRoot, srcFile)
	if err != nil {
		return err
	}
	defer f.Close()

	ctx := context.Background()
	format, reader, err := archives.Identify(ctx, src, f)
	if err != nil {
		return err
	}

	root := &os.Root{}

	if !e.DryRun {
		// Ensure the parent directory exists
		if err := os.MkdirAll(dst, 0o700); err != nil {
			return err
		}

		root, err = os.OpenRoot(dst)
		if err != nil {
			return err
		}
		defer root.Close()
	}

	return format.(archives.Extractor).Extract(ctx, reader, func(ctx context.Context, info archives.FileInfo) error {
		return handleFile(info, root, dst, e)
	})
}

func handleFile(file archives.FileInfo, root *os.Root, dst string, e *extractor) error {
	fileInArchive := file.NameInArchive
	if fileInArchive == "" {
		fileInArchive = file.Name()
	}

	if err := validatePath(file, fileInArchive, dst, e); err != nil {
		return err
	}

	// Handle directories
	if !e.DryRun && file.IsDir() {
		return root.Mkdir(fileInArchive, 0o700)
	}

	if !e.DryRun {
		// Handle regular files
		reader, err := file.Open()
		if err != nil {
			return err
		}
		defer reader.Close()

		writer, err := root.OpenFile(fileInArchive, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, file.Mode()&os.ModePerm)
		if err != nil {
			return err
		}
		defer writer.Close()

		if _, err := io.Copy(writer, reader); err != nil {
			return err
		}
	}

	return nil
}

// 1. path should not be absolute
// 2. path should be local (i.e: no "../")
// 3. path should not be a symlink
func validatePath(f archives.FileInfo, path, dst string, e *extractor) error {
	if filepath.IsAbs(path) {
		return ErrAbsolutePath
	}

	rel, err := filepath.Rel(dst, filepath.Join(dst, path))
	if err != nil {
		return err
	}
	if !filepath.IsLocal(rel) {
		return ErrRelativePath
	}

	if !e.AllowSymlink && (f.LinkTarget != "" || f.Mode()&os.ModeType == os.ModeSymlink) {
		return ErrSymlinkInPath
	}

	return nil
}
