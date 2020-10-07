binary:
	GOOS=linux GOARCH=amd64 go build -o ./docker/artifacts/tatanka ./cmd/tatanka

## Generates mock golang interfaces for testing
mock:
	go install github.com/golang/mock/mockgen
	mockgen -destination internal/cassandra/mock_cassandra/mock.go github.com/cpurta/tatanka/internal/cassandra Client

.PHONY: test
test:
	go test ./...

all: mock test
