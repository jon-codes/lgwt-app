build:
	go build -o ./bin/lgwt-app
start:
	go build -o ./bin/lgwt-app && ./bin/lgwt-app
test:
	go test ./...
cover:
	go test -cover ./...