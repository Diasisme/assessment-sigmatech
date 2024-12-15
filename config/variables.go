package config

import (
	"assesment-sigmatech/service/models"

	"github.com/spf13/viper"
)

func NewEnvVar(viper *viper.Viper) (envVar models.VarEnviroment) {

	envVar = models.VarEnviroment{
		Host:        viper.GetString("POSTGRES_HOST"),
		Port:        viper.GetInt32("POSTGRES_DB_PORT"),
		User:        viper.GetString("POSTGRES_USER"),
		Pass:        viper.GetString("POSTGRES_PASSWORD"),
		DB:          viper.GetString("POSTGRES_DB"),
		ServicePort: viper.GetString("SVC_PORT"),
		Service:     viper.GetString("CONTAINER_ID_NAME"),
		MinioPort:   viper.GetString("MINIO_API_PORT"),
		MinioUser:   viper.GetString("MINIO_ROOT_USER"),
		MinioPass:   viper.GetString("MINIO_ROOT_PASSWORD"),
		MinioBucket: viper.GetString("MINIO_BUCKET"),
	}

	return
}
