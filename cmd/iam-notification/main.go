package main

import (
	"fmt"

	"github.com/rusik69/iam-notification/pkg/cmd"
	"github.com/rusik69/iam-notification/pkg/report"
)

func main() {
	args := cmd.ParseArgs()
	rep, err := report.NewReport(args.BucketName, args.AwsRegion, args.CsvFileName)
	if err != nil {
		panic(err)
	}
	files, err := rep.ListBucket()
	if err != nil {
		panic(err)
	}
	lastReports, err := rep.GetLastReports(files)
	if err != nil {
		panic(err)
	}
	accountOwners, err := rep.GetAccountOwners()
	if err != nil {
		panic(err)
	}
	reports, err := rep.GenerateReports(lastReports)
	if err != nil {
		panic(err)
	}
	fmt.Println(accountOwners)
	fmt.Println(lastReports)
}
