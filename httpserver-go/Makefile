build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpserver-go ./src
	docker build . -t httpserver-go
	rm -f httpserver-go
build-macosx:
	go build -o httpserver-go-macosx ./src
	@echo "Run ./httpserver-go-macosx to test the Go webserver"
