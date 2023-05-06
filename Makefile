build:
	go build -o bin/server.exe
run: build
	./bin/server.exe
test:
	go test -v ./...