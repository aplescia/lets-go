package secretsmanager

import (
	"github.com/aplescia-chwy/lets-go/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var (
	log, _ = util.InitLoggerWithLevel(nil)
)

//GetSecret retrieves a secret from Secrets Manager. Searches for a secret stored under the input
//	secretId
func GetSecret(secretId string) string {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(util.GetEnvOrDefault("AWS_REGION", "us-east-1")),
	})
	if err != nil {
		log.Fatalln(err)
	}
	svc := secretsmanager.New(sess)
	output, err := svc.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &secretId})
	if err != nil {
		log.Fatalln(err)
	}
	return *output.SecretString
}

//CreateSecret creates a Secrets Manager secret using a given CreateSecretInput. Return errors, if any.
func CreateSecret(opts secretsmanager.CreateSecretInput) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(util.GetEnvOrDefault("AWS_REGION", "us-east-1")),
	})
	if err != nil {
		log.Fatalln(err)
	}
	svc := secretsmanager.New(sess)
	_, err = svc.CreateSecret(&opts)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
