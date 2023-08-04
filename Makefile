make:
	go build -o ./build/
PHONY: make

make-win:
	GOOS=windows GOARCH=amd64 go build -o ./build/
PHONY: make-win