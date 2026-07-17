run:
	go run cmd/server/main.go

test:
	go test ./...

test-race:
	go test -race ./...