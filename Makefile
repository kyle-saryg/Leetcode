build:
	@go build -o bin/sandbox

run: build
	@./bin/sandbox