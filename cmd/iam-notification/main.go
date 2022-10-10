package main

import (
	"fmt"

	"github.com/rusik69/iam-notification/pkg/cmd"
	"github.com/rusik69/iam-notification/pkg/report"
)

func main() {
	args := cmd.ParseArgs()
	files, err := report.ListBucket(args.BucketName, args.AwsRegion)
	if err != nil {
		panic(err)
	}
	lastReport := report.GetLastReport(files)
	if lastReport == "" {
		panic("no reports found")
	}
	accountOwners, err := report.GetAccountOwners(args.CsvFileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(accountOwners)
	fmt.Println(lastReport)
}
