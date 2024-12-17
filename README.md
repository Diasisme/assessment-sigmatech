## How to run

1. Clone repository `git clone https://github.com/Diasisme/assessment-sigmatech.git`
2. Open cmd/cli and enter to directory repository that already cloned
3. Install dependecy in cmd/cli with `go mod tidy`
4. Run `docker volume create assessment-sigmatech-minio-data` in cmd/cli for creating minIO volume in docker
5. For run the service, in cmd/cli type `docker compose up -d` then hit enter. Wait for few minutes to install all dependecy needed for this program for the first time (If install/download failed, please change your connection)
6. For stop the service, in cmd/cli type `docker compose down -v`
