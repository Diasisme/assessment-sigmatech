package app

import (
	"assesment-sigmatech/config/logging"
	"assesment-sigmatech/service/minio"
)

type accountApp struct {
	accRepo     AccountDatastore
	log         *logging.Logger
	minioClient *minio.MinioData
}

func InitApp(db AccountDatastore, log *logging.Logger, minio *minio.MinioData) AccountApp {
	return &accountApp{
		accRepo:     db,
		log:         log,
		minioClient: minio,
	}
}
