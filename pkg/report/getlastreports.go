package report

import (
	"errors"
	"sort"
	"strings"
)

// GetLastReport gets the last report dir from the bucket.
func (rep *Report) GetLastReports(files []string) ([]string, error) {
	if len(files) == 0 {
		return nil, errors.New("no files found")
	}
	res := make([]string, 0)
	sort.Strings(files)
	lastFile := files[len(files)-1]
	lastPrefix := strings.Split(lastFile, "/")[0]
	for _, file := range files {
		if strings.HasPrefix(file, lastPrefix) {
			res = append(res, file)
		}
	}
	return res, nil
}
