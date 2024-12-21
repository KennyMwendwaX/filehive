build:
	@go build -o bin/filehive

run: build
	@./bin/filehive

test:
	@go test -v ./...	


