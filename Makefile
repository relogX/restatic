lint:
	gofmt -w .

build:
	go build -o ./bin/restatic ./cmd/restatic
