package osutil

import (
	"os"
	"path/filepath"
)

// CopyDirAll will copy a directory recursively.
func CopyDirAll(srcDir, dstDir string) error {
	if err := os.MkdirAll(dstDir, DefaultDirPerm); err != nil {
		return err
	}

	files, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		src := filepath.Join(srcDir, file.Name())
		dst := filepath.Join(dstDir, file.Name())

		if file.IsDir() {
			if err := CopyDirAll(src, dst); err != nil {
				return err
			}
			continue
		}
		if err := CopyFile(src, dst); err != nil {
			return err
		}
	}
	return nil
}

// ListDir will return all files in a directory.
func ListDir(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f.Readdirnames(-1)
}
