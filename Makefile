build:
	@go build -o bin/savings cmd/main.go

test:
	@go test -v ./..

run: build
	@./bin/savings
