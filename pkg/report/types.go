package report

// Report is a struct for report
type Report struct {
	awsRegion   string
	bucketName  string
	csvFileName string
}

// User is a struct for user
type User struct {
	UserName               string
	Created                string
	PasswordEnabled        string
	passwordLastUsed       string
	passwordLastChanged    string
	accessKey1Active       string
	accessKey1LastRotated  string
	accessKey1LastUsedDate string
	accessKey2Active       string
	accessKey2LastRotated  string
	accessKey2LastUserDate string
}

// Reports is a struct for reports
type Reports map[string]User
