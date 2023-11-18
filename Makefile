.DEFAULT_GOAL := run-dev

run-dev:
	go run cmd/main.go

build:
	go build -o ./bin/main cmd/main.go

compile:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/main-darwin cmd/main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/main-linux cmd/main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/main-windows cmd/main.go

run-prod: compile
	./bin/main-linux


clean:
	go clean
	rm -rf ./bin