build:
	@go build -o bin/filehive

run: build
	@./bin/fh

test:
	@go test -v ./...	


