package gogetsimples3

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var message_error_format string = "%s \n %s"

//
type GoObjectS3Bucket struct {
	clientAwsS3 *s3.Client
	bucketName  string
	objectName  string
}

//
func NewGoObjectS3Bucket(strBucketName, strObjectName string) (GoObjectS3Bucket, error) {
	goObjectS3Bucket := GoObjectS3Bucket{}
	config, errConfig := config.LoadDefaultConfig(context.TODO())

	if errConfig != nil {
		strMessageI := "Não foi possível inicializar o s3 Client."
		return goObjectS3Bucket, errors.New(fmt.Sprintf(message_error_format, strMessageI, errConfig.Error()))
	}

	goObjectS3Bucket.bucketName = strBucketName
	goObjectS3Bucket.objectName = strObjectName
	goObjectS3Bucket.clientAwsS3 = s3.NewFromConfig(config)

	return goObjectS3Bucket, nil
}

//
func (ref GoObjectS3Bucket) GetObjectStrInBucket() (dataStr string, iError error) {

	obInput := &s3.GetObjectInput{
		Bucket: aws.String(ref.bucketName),
		Key:    aws.String(ref.objectName),
	}

	result, iErr := ref.clientAwsS3.GetObject(context.TODO(), obInput)

	if iErr != nil {
		strMessageI := "Erro ao Carregar recurso do Bucket S3"
		return "", errors.New(fmt.Sprintf(message_error_format, strMessageI, iErr.Error()))
	}

	defer result.Body.Close()

	body, errReadBody := ioutil.ReadAll(result.Body)

	if errReadBody != nil {
		strMessageI := "Erro ao ler estrutura do objeto S3"
		return "", errors.New(fmt.Sprintf(message_error_format, strMessageI, errReadBody.Error()))
	}

	strReturn := string(body)

	return strReturn, nil
}
