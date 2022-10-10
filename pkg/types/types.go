package types

// Args is a struct that holds the arguments for the command.
type Args struct {
	BucketName  string // Name of the bucket to get iam user reports from
	CsvFileName string // Name of the csv file to get account owners from
	AwsRegion   string // AWS region name
}
