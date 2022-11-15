package report

import (
	"encoding/csv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetReport gets report
func (rep *Report) GetReportUsers(reportName string) ([]User, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(rep.awsRegion)},
	)
	if err != nil {
		return nil, err
	}
	svc := s3.New(sess)
	params := &s3.GetObjectInput{
		Bucket: aws.String(rep.bucketName),
		Key:    aws.String(reportName),
	}
	resp, err := svc.GetObject(params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	csvReader := csv.NewReader(resp.Body)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	var res []User
	skip := true
	for _, record := range records {
		if skip {
			skip = false
			continue
		}
		user := User{
			UserName:               record[0],
			Created:                record[2],
			PasswordEnabled:        record[3],
			passwordLastUsed:       record[4],
			passwordLastChanged:    record[5],
			accessKey1Active:       record[8],
			accessKey1LastRotated:  record[9],
			accessKey1LastUsedDate: record[10],
			accessKey2Active:       record[13],
			accessKey2LastRotated:  record[14],
			accessKey2LastUserDate: record[15],
		}
		res = append(res, user)
	}
	return res, nil
}
