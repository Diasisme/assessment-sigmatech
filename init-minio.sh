#!/bin/sh
mc alias set minio-sigmatech http://$MINIO_PATH_URL:$MINIO_API_PORT $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD
mc mb -p minio-sigmatec/$MINIO_BUCKET
mc policy set public minio-sigmatec/$MINIO_BUCKET