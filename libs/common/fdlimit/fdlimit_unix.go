// +build linux darwin netbsd openbsd solaris

package fdlimit

import "syscall"

// Raise tries to maximize the file descriptor allowance of this process
// to the maximum hard-limit allowed by the OS.
// Returns the size it was set to (may differ from the desired 'max')
func Raise(max uint64) (uint64, error) {
	// Get the current limit
	var limit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		return 0, err
	}
	// Try to update the limit to the max allowance
	limit.Cur = limit.Max
	if limit.Cur > max {
		limit.Cur = max
	}
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		return 0, err
	}
	// MacOS can silently apply further caps, so retrieve the actually set limit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		return 0, err
	}
	return limit.Cur, nil
}

// Current retrieves the number of file descriptors allowed to be opened by this
// process.
func Current() (int, error) {
	var limit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		return 0, err
	}
	return int(limit.Cur), nil
}

// Maximum retrieves the maximum number of file descriptors this process is
// allowed to request for itself.
func Maximum() (int, error) {
	var limit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		return 0, err
	}
	return int(limit.Max), nil
}
