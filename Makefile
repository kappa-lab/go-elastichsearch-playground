run:
	go run cmd/main.go

test:
	go test -v -count=1 ./...

start-docker:
	docker-compose up -d	