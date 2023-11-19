build:
	@go build -o bin/gosearchengine
run: build
	@./bin/gosearchengine
test:
	@go test -v ./...