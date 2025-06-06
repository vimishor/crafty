package lhttputil

import (
	"crypto/sha256"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/adrg/xdg"
	lua "github.com/yuin/gopher-lua"

	"github.com/schollz/progressbar/v3"

	"github.com/vimishor/crafty/pkg/osutil"
)

var (
	checksum string = ""
	silent   bool   = false
)

func fileDownloadFn(l *lua.LState) int {
	src := l.CheckString(1)
	dst := l.OptString(2, xdg.CacheHome)
	var opts *lua.LTable
	if l.GetTop() > 2 {
		opts = l.CheckTable(3)
	}

	if opts != nil {
		if val := l.GetField(opts, `checksum`); val != lua.LNil {
			checksum = val.String()
		}

		if val := l.GetField(opts, `silent`); val != lua.LNil {
			silent = bool(val.(lua.LBool))
		}
	}

	dstFile, err := maybeDownload(src, dst)
	if err != nil {
		l.Push(lua.LString(""))
		l.Push(lua.LString(err.Error()))
		return 2
	}

	l.Push(lua.LString(dstFile))
	l.Push(lua.LNil)
	return 2
}

func maybeDownload(url, fpath string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	if fpath == "" {
		// FIXME: inherit from `--cache-dir`
		fpath = filepath.Join(xdg.CacheHome, "crafty", "download")
	}

	if osutil.IsDir(fpath) {
		fpath = filepath.Join(fpath, path.Base(req.URL.Path))
	}

	// If file is cached, verify integrity and skip download
	if checksum != "" && osutil.IsFile(fpath) {
		sum, err := osutil.FileChecksum(fpath, sha256.New())
		if err != nil {
			return "", fmt.Errorf("checksum failed for %q with %v", fpath, err)
		}
		if sum != checksum {
			return "", fmt.Errorf("invalid checksum for %q: expected %q and got %q", filepath.Base(fpath), checksum, sum)
		}

		return fpath, nil
	}

	return doDownload(req, fpath)
}

func doDownload(req *http.Request, fpath string) (string, error) {
	f, err := os.CreateTemp(os.TempDir(), "*.crafty.tmp")
	if err != nil {
		return "", err
	}
	defer f.Close()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad HTTP status: %q", resp.Status)
	}

	// fetch filename from response only if the user didn't specified a path with filename included
	if osutil.IsDir(fpath) {
		_, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
		if err == nil {
			if filename, ok := params["filename"]; ok {
				fpath = filepath.Join(filepath.Dir(fpath), filename)
			}
		}
	}

	writer := io.MultiWriter(f)
	if !silent {
		bar := progressbar.NewOptions64(resp.ContentLength,
			progressbar.OptionSetDescription(path.Base(fpath)),
			progressbar.OptionSetWriter(os.Stderr),
			progressbar.OptionShowBytes(true),
			progressbar.OptionShowTotalBytes(true),
			progressbar.OptionSetWidth(10),
			progressbar.OptionThrottle(65*time.Millisecond),
			progressbar.OptionShowCount(),
			progressbar.OptionOnCompletion(func() {
				fmt.Fprint(os.Stderr, "\n")
			}),
			progressbar.OptionSpinnerType(14),
			progressbar.OptionFullWidth(),
			progressbar.OptionSetRenderBlankState(true),
			// progressbar.OptionClearOnFinish(),
		)
		writer = io.MultiWriter(f, bar)
	}

	if _, err := io.Copy(writer, resp.Body); err != nil {
		return "", err
	}

	fmt.Print("\n")

	return fpath, os.Rename(f.Name(), fpath)
}
