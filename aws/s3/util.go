package s3

import (
	"github.com/Chewy-Inc/lets-go/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"path/filepath"
)

var (
	log, _ = util.InitLoggerWithLevel(nil)
)

func DownloadFileFromS3EventRecord(downloader *s3manager.Downloader, record events.S3EventRecord) (*os.File, error) {
	file, err := os.Create(filepath.Join("tmp", record.S3.Object.Key))
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	_, err = downloader.Download(
		file, &s3.GetObjectInput{Bucket: &record.S3.Bucket.Name,
			Key: &record.S3.Object.Key})
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	return file, err
}
