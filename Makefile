run:
	GIN_MODE=release go run cmd/hrbrain-go-assignment/main.go

test:
	go test ./internal/services/...
