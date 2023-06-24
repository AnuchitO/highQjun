test-cover:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test:
	go test -v ./...

bench:
	go test -v -bench ./...

run:
	go run main.go