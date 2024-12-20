build:
	@go build -o bin/fh

run: build
	@./bin/fh

test:
	@go test -v ./...	


