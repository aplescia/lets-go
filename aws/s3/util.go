package s3

import (
	"github.com/aplescia-chwy/lets-go/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"path/filepath"
)

var (
	log, _ = util.InitLoggerWithLevel(nil)
)

//DownloadFileFromS3EventRecord downloads a file from an S3 event's S3 Event Record. Takes in a pointer to a downloader object.
//		for _, eventObject := range event.Records {
//		file, err := s3.DownloadFileFromS3EventRecord(downloader, eventObject)
//		if err != nil {
//			log.Panicln(err)
//		}
//		ProcessFile(file.Name())
//	}
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
