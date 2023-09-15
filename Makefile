BINARY_NAME=vedirect-status

build:
	GOARCH=arm64 GOOS=linux go build -o ./build/${BINARY_NAME}-linux main.go

run: build
	./${BINARY_NAME}-linux

clean:
	go clean
	rm ${BINARY_NAME}-*

test:
	go test ./... ./vedirect-device