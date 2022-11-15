package report

// GenerateReports generates reports
func (rep *Report) GenerateReports(lastReports []string) (Reports, error) {
	var res Reports
	for _, report := range lastReports {
		reportAccount := GetReportAccount(report)
		reportUsers, err := rep.GetReportUsers(report)
		if err != nil {
			return nil, err
		}
		res[reportAccount] = reportUsers
	}
	return
}
