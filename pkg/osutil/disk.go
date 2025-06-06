package osutil

import "syscall"

// UsageInfo holds data about disk usage, like total size,
// amount of space used and remaining free space.
type UsageInfo struct {
	Total uint64
	Used  uint64
	Free  uint64
}

// DiskUsage returns disk usage for the given path.
func DiskUsage(path string) (UsageInfo, error) {
	var fs_t syscall.Statfs_t
	usage := UsageInfo{}
	if err := syscall.Statfs(path, &fs_t); err != nil {
		return usage, err
	}

	usage.Total = fs_t.Blocks * uint64(fs_t.Bsize)
	usage.Free = fs_t.Bfree * uint64(fs_t.Bsize)
	usage.Used = usage.Total - usage.Free

	return usage, nil
}
