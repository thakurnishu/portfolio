build:
	@mkdir -p bin
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main
run: build
	@./bin/main
