package report

// NewReport creates new report object
func NewReport(bucketName string, awsRegion string, csvFileName string) (Report, error) {
	return Report{
		awsRegion:   awsRegion,
		bucketName:  bucketName,
		csvFileName: csvFileName,
	}, nil
}
