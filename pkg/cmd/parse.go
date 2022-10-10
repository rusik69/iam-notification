package cmd

import (
	"flag"

	"github.com/rusik69/iam-notification/pkg/types"
)

// ParseArgs parses the command line arguments.
func ParseArgs() types.Args {
	bucketName := flag.String("bucketName", "", "iam reports bucket name")
	awsRegion := flag.String("awsRegion", "us-east-1", "AWS region name")
	csvFileName := flag.String("csvFileName", "", "csv file name")
	flag.Parse()
	if *bucketName == "" || *csvFileName == "" {
		flag.PrintDefaults()
		panic("bucketName and csvFileName are required")
	}
	return types.Args{
		BucketName:  *bucketName,
		CsvFileName: *csvFileName,
		AwsRegion:   *awsRegion,
	}
}
