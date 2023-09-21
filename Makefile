BINARY_NAME=jam

clean:
	@rm -rf /home/miles/go/jam/cleaned.csv /home/miles/go/jam/dump

build: clean
	@/usr/local/go/bin/go build -o bin/${BINARY_NAME} ./src/main.go

run: build
	@./bin/${BINARY_NAME}

test:
	@go test -v ./...

