package larchive

import (
	"errors"
	"testing"
)

func TestGoodArchives(t *testing.T) {
	for _, path := range []string{
		"testdata/tar/good1.tar",
		"testdata/zip/good1.zip",
	} {
		err := doExtract(path, "/some/dir", WithDryRun())
		if err != nil {
			t.Errorf("validatePath for %q: got err %v, want no error", path, err)
		}
	}
}

func TestAbsolutePath(t *testing.T) {
	for _, path := range []string{
		"testdata/tar/absolute1.tar",
		"testdata/tar/absolute2.tar",
		"testdata/zip/absolute1.zip",
		"testdata/zip/absolute2.zip",
	} {
		err := doExtract(path, "/some/dir", WithDryRun())
		if !errors.Is(err, ErrAbsolutePath) {
			t.Errorf("validatePath for %q: got err %v, want ErrAbsolutePath", path, err)
		}
	}
}

func TestRelativePath(t *testing.T) {
	for _, path := range []string{
		"testdata/tar/relative0.tar",
		"testdata/tar/relative2.tar",
		"testdata/zip/relative0.zip",
		"testdata/zip/relative2.zip",
	} {
		err := doExtract(path, "/some/dir", WithDryRun())
		if !errors.Is(err, ErrRelativePath) {
			t.Errorf("validatePath for %q: got err %v, want ErrRelativePath", path, err)
		}
	}
}

func TestSymlinkPath(t *testing.T) {
	for _, path := range []string{
		"testdata/tar/dirsymlink.tar",
		"testdata/tar/dirsymlink2a.tar",
		"testdata/tar/dirsymlink2b.tar",
		"testdata/tar/symlink.tar",
		"testdata/zip/dirsymlink.zip",
		"testdata/zip/dirsymlink2a.zip",
		"testdata/zip/dirsymlink2b.zip",
		"testdata/zip/symlink.zip",
	} {
		err := doExtract(path, "/some/dir", WithDryRun())
		if !errors.Is(err, ErrSymlinkInPath) {
			t.Errorf("validatePath for %q: got err %v, want ErrSymlinkInPath", path, err)
		}
	}
}
