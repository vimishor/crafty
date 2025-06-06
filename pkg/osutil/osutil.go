// Package osutil provides filesystem utility functions
package osutil

import "os"

const (
	DefaultDirPerm  = 0o755
	DefaultFilePerm = 0o644
)

// IsFile returns true if the file exists.
func IsFile(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

// IsDir returns true if the directory exists.
func IsDir(name string) bool {
	info, err := os.Stat(name)
	return err == nil && info.IsDir()
}
