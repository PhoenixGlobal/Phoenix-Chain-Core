// +build iospackage metrics

// ReadCPUStats retrieves the current CPU stats. Internally this uses `gosigar`,
// which is not supported on the platforms in this file.
func ReadCPUStats(stats *CPUStats) {}
