package s3adapter

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

type S3Adapter struct {
	S3         *s3.S3
	Domain     string
	BucketName string
}

func (s *S3Adapter) New() error {
	fmt.Println(config.C.Filestore.Region)
	ssn, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.C.Filestore.Region),
		Endpoint:    aws.String(config.C.Filestore.Domain),
		Credentials: credentials.NewStaticCredentials(config.C.Filestore.KeyId, config.C.Filestore.Secret, ""),
	})

	if err != nil {
		return err
	}

	s.S3 = s3.New(ssn)
	s.Domain = config.C.Filestore.Domain
	s.BucketName = config.C.Filestore.BucketName
	return nil
}

func (s *S3Adapter) GetDomain() string {
	return s.Domain
}

func (s *S3Adapter) Delete(key string, recordID uuid.UUID) error {
	_, err := s.S3.DeleteObject(&s3.DeleteObjectInput{Bucket: &s.BucketName, Key: &key})

	return err
}

func (s *S3Adapter) Save(file multipart.File, filename string, path string, recordID uuid.UUID) (string, error) {
	key := filepath.Join(path, filename)

	_, err := s.S3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	return key, nil
}

func (s *S3Adapter) GetTemporaryURL(key string, recordID uuid.UUID) (string, error) {
	if key == "" {
		return "", nil
	}
	req, _ := s.S3.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(key),
	})

	urlStr, err := req.Presign(24 * time.Hour) // Tempo de expiração da URL (24 horas neste exemplo)
	if err != nil {
		return "", err
	}

	return urlStr, nil
}

func (s *S3Adapter) SaveAndGetTemporaryURL(file multipart.File, filename string, path string, recordID uuid.UUID) (string, string, error) {
	key, err := s.Save(file, filename, path, recordID)
	if err != nil {
		return "", "", err
	}

	url, err := s.GetTemporaryURL(key, recordID)
	if err != nil {
		return "", "", err
	}

	return key, url, nil
}
