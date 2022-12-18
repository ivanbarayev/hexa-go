#!make

echos:
	echo "$$service"

run:
	swag init
	go run cmd/api/main.go
	clear

swag:
	 swag init -g .\cmd\http\main.go

genproto:
	@echo Generating es microservice order gRPC proto
	protoc --proto_path=proto --grpc-gateway_out -I proto\google --go_out=proto\auth --go-grpc_out=proto\auth --go-grpc_opt=require_unimplemented_servers=false --grpc-gateway_opt paths=source_relative .\proto\auth\auth.proto

rebuild:
	go build -ldflags "-s -w" -o service
	systemctl restart ${APP_SERVICE_NAME}.service
	clear
	journalctl -xe -u ${APP_SERVICE_NAME} -f

build:
	swag init -g .\cmd\http\main.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/app-amd64-linux cmd/http/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/app-amd64-win.exe cmd/http/main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/app-amd64-mac cmd/http/main.go

buildl:
	swag init -g .\cmd\http\main.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/app-amd64-linux cmd/api/main.go

buildw:
	swag init -g .\cmd\http\main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/app-amd64-win.exe cmd/api/main.go

buildm:
	swag init -g .\cmd\http\main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/app-amd64-mac cmd/api/main.go

ports:
	sudo netstat -tulpn | grep LISTEN