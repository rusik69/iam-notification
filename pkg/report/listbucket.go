package report

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ListBucket lists the contents of an s3 bucket.
func (rep *Report) ListBucket() ([]string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(rep.awsRegion)},
	)
	if err != nil {
		return nil, err
	}
	svc := s3.New(sess)
	params := &s3.ListObjectsInput{
		Bucket: aws.String(rep.bucketName),
	}
	resp, err := svc.ListObjects(params)
	if err != nil {
		return []string{}, err
	}
	var keys []string
	for _, key := range resp.Contents {
		keys = append(keys, *key.Key)
	}
	return keys, nil
}
