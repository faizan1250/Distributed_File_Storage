run: build
	@./bin/main

build:
	@go build -o ./bin/main .

test:
	@go test ./... -v
