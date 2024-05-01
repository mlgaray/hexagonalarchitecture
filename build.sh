set -xe

go get -u  ./...

GOOS=linux GOARCH=amd64 go build -o bin/application -ldflags="-s -w"
