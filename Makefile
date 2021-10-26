build:
	go build -o ./bin/go-search cmd/go-search/main.go
run: build
	./bin/go-search