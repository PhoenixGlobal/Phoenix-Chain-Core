package metrics

// getProcessCPUTime returns 0 on Windows as there is no system call to resolve
// the actual process' CPU time.
func getProcessCPUTime() int64 {
	return 0
}
