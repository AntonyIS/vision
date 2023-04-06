build:
	go build -o bin/vision
serve: build
	./bin/vision
test:
	go test -v ./...