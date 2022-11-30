.PHONY: test coverage coverage-html build run compile

test:
	go test -v ./...

coverage:
	go test -cover -v ./...

coverage-html:
	go test -covermode=count -coverpkg=./... -coverprofile coverage/cover.out -v ./...
	go tool cover -html coverage/cover.out -o coverage/cover.html

build:
	go build -o bin/books-crud cmd/main.go

run:
	go run cmd/main.go

compile:
	echo "Compiling for every windows, linux and mac os x86_64 platform"
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 cmd/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64.exe cmd/main.go
