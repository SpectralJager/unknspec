watch:
	./tailwindcss -i ./public/css/temp.css -o ./public/css/res.css --watch
build:
	go build -o cmd/server ./src/main.go 
run: build
	./cmd/server