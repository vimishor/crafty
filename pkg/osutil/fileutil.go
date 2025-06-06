package osutil

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"os"
)

// CopyFlags represent options for CopyFile.
type CopyFlags struct {
	// Overwrite destination file if exists.
	FileOverwrite bool

	// Sync commits the current contents of the file to stable storage.
	Sync bool
}

type CopyFlag func(*CopyFlags) error

func defaultCopyFlags() *CopyFlags {
	return &CopyFlags{
		FileOverwrite: false,
	}
}

func WithOverwrite() CopyFlag {
	return func(cf *CopyFlags) error {
		cf.FileOverwrite = true
		return nil
	}
}

// CopyFile will copy a file while preserving permissions and modification time.
func CopyFile(src, dst string, flags ...CopyFlag) error {
	options := defaultCopyFlags()
	for _, opt := range flags {
		if err := opt(options); err != nil {
			return err
		}
	}

	srcF, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcF.Close()

	stat, err := srcF.Stat()
	if err != nil {
		return err
	}

	if !stat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	if IsDir(dst) {
		// If `dst` is a directory, append file name to `dst`
		dst = fmt.Sprintf("%s/%s", dst, stat.Name())
	}

	// Because we write to a temp file, it is a bit inconvenient to
	// use O_EXCL flag when opening the file, so we check here if
	// the destination file exists and if we can overwrite it.
	if !options.FileOverwrite && IsFile(dst) {
		return fmt.Errorf("file exists: %q", dst)
	}

	tmpFile := dst + ".tmp"

	dstF, err := os.OpenFile(tmpFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, stat.Mode()&os.ModePerm)
	if err != nil {
		return err
	}
	defer dstF.Close()

	_, err = io.Copy(dstF, srcF)
	if err != nil {
		return err
	}

	// check close error
	if err := dstF.Close(); err != nil {
		return err
	}

	if options.Sync {
		if err := dstF.Sync(); err != nil {
			return err
		}
	}

	if err := os.Chtimes(tmpFile, stat.ModTime(), stat.ModTime()); err != nil {
		return err
	}

	return os.Rename(tmpFile, dst)
}

// MoveFile moves a file.
//
// Moving will be done by copying the file to the new location
// and removing the old file.
func MoveFile(oldFile, newFile string) error {
	if err := CopyFile(oldFile, newFile); err != nil {
		return err
	}
	return os.Remove(oldFile)
}

// FileChecksum computes and returns file checksum of various lenghts,
// from SHA-2 family.
//
// If "hasher" is not specified, [sha256] will be used as default.
//
// [sha256]: https://pkg.go.dev/crypto/sha256
func FileChecksum(name string, hasher hash.Hash) (string, error) {
	if !IsFile(name) {
		return "", fmt.Errorf("no such file %s", name)
	}

	if hasher == nil {
		hasher = sha256.New()
	}

	f, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
