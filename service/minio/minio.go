package minio

import (
	"assesment-sigmatech/config/logging"
	"assesment-sigmatech/service/models"
	"bytes"
	"context"
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

type MinioData struct {
	minio             *minio.Client
	log               *logging.Logger
	endpoint          string
	access_key_id     string
	secret_access_key string
	minio_bucket      string
}

func InitMinio(varenv models.VarEnviroment, log *logging.Logger) *MinioData {
	var err error

	endpoint := "172.26.0.3:9000"
	accessKeyID := varenv.MinioUser
	secretAccessKey := varenv.MinioPass
	useSSL := false // Ubah jika menggunakan SSL

	log.Info(logrus.Fields{
		"endpoint":        endpoint,
		"accessKeyID":     accessKeyID,
		"secretAccessKey": secretAccessKey,
	}, nil, "info log")

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Success connected to MinIO")

	return &MinioData{
		minio:             client,
		log:               log,
		access_key_id:     accessKeyID,
		secret_access_key: secretAccessKey,
		minio_bucket:      varenv.MinioBucket,
		endpoint:          endpoint,
	}
}

func (f *MinioData) UploadToCloud(c echo.Context, folder, filename, extension string) (result models.UploadToCloudResult, err error) {

	contentTypeMap := map[string]string{
		"JPG": "image/jpeg",
		"PNG": "image/png",
	}

	var (
		objectName  = fmt.Sprintf("%s/%s", folder, filename)
		contentType = contentTypeMap[extension]
	)

	fmt.Printf("bucket: %s", f.minio_bucket)

	file, err := c.FormFile("file")
	if err != nil {
		return
	}

	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err = buf.ReadFrom(src); err != nil {
		return
	}

	fileSize := file.Size

	uploadInfo, err := f.minio.PutObject(
		context.Background(),
		f.minio_bucket,
		objectName,
		bytes.NewReader(buf.Bytes()),
		fileSize,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return
	}

	fileURL := fmt.Sprintf("http://%s/%s/%s", f.endpoint, f.minio_bucket, objectName)

	f.log.Info(logrus.Fields{"uploadInfo": uploadInfo}, nil, "Successfully uploaded report to cloud")

	result = models.UploadToCloudResult{
		Path:   fmt.Sprintf("%s/%s", folder, filename),
		Bucket: f.minio_bucket,
		Url:    fileURL,
	}

	return
}
