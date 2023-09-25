BINARY_NAME=WatchGuardian

run:
	echo "Running ${BINARY_NAME}"
	go run .

compile:
	echo "Compiling for every OS and Platform"
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin .
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux .
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows .

build:
	echo "Building ${BINARY_NAME} for local OS and Platform"
	go build -o bin/${BINARY_NAME} .

clean:
	go clean
	rm -rf bin/

test:
	go test ./...