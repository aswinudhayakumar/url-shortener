dev:
	go run main.go

dep:
	go mod tidy
	go mod vendor

build:
	go build .

run-build: 
	./url-shortener.exe

run: build
	./url-shortener.exe