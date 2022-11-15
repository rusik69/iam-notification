package report

import "strings"

func GetReportAccount(report string) string {
	reportSplit := strings.Split(report, "_")
	reportFile := reportSplit[1]
	reportSplit = strings.Split(reportFile, ".")
	reportFile = reportSplit[0]
	return reportFile
}
