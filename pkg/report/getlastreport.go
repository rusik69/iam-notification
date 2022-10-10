package report

import "sort"

// GetLastReport gets the last report dir from the bucket.
func GetLastReport(files []string) string {
	if len(files) == 0 {
		return ""
	}
	sort.Strings(files)
	return files[len(files)-1]
}
