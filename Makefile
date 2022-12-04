build: 
	@go build -o bin/imbored 

run: 
	@./bin/imbored random

test: 
	go test -v ./...