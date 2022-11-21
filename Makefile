build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go

dev: 
	nodemon --exec go run cmd/main.go --signal SIGTERM



