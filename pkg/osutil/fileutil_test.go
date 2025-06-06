package osutil

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCopyFile(t *testing.T) {
	type Test struct {
		name          string
		src           string
		skipCreateSrc bool
		dst           string
		dstIsDir      bool
		err           string
	}
	tests := []Test{
		{
			name: "Same name as source",
			src:  "src/foo",
			dst:  "dst/foo",
		},
		{
			name: "Different name than source",
			src:  "src/bar/baz",
			dst:  "dst/bar/foo",
		},
		{
			name: "Try to overwrite a file",
			src:  "src/foo",
			dst:  "dst/foo",
		},
		{
			name:     "To directory",
			src:      "src/baz/bar",
			dst:      "dst/foo",
			dstIsDir: true,
		},
		{
			name:          "Inexistent file",
			src:           "src/not_here",
			dst:           "dst/not_here",
			skipCreateSrc: true,
			err:           "no such file",
		},
		{
			name: "Dir instead of file",
			src:  "src/a_dir/",
			dst:  "dst/file",
			err:  "is not a regular file",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := t.TempDir()
			src := filepath.Join(root, test.src)
			dst := filepath.Join(root, test.dst)

			last_char := string(test.src[len(test.src)-1])
			if last_char == "/" {
				src += "/"
			}

			// +setup
			if !test.skipCreateSrc {
				if err := createFile(src); err != nil {
					t.Fatal(err)
				}
			}

			dstCreate := filepath.Dir(dst)
			if test.dstIsDir {
				dstCreate = dst
			}
			if err := os.MkdirAll(dstCreate, DefaultDirPerm); err != nil {
				t.Fatal(err)
			}
			// -setup

			// _ = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			// 	fmt.Printf("walk: %s\n", d)
			// 	return nil
			// })

			err := CopyFile(src, dst)
			if err != nil {
				if test.err != "" {
					if strings.Contains(err.Error(), test.err) {
						return
					}
					t.Fatalf("got err %q, want %q", err, test.err)
				}
				t.Fatal(err)
			}

			if test.err != "" {
				t.Fatalf("got no err, want %q", test.err)
			}

			if !IsFile(dst) {
				t.Fatalf("%s does not exists", dst)
			}
		})
	}
}

func createFile(fpath string) error {
	pathCreate := filepath.Dir(fpath)

	if fpath[len(fpath)-1:] == "/" {
		pathCreate = fpath
	}

	if err := os.MkdirAll(pathCreate, DefaultDirPerm); err != nil {
		return err
	}

	if fpath[len(fpath)-1:] == "/" {
		return nil
	}
	return os.WriteFile(fpath, []byte{'a'}, DefaultFilePerm)
}
