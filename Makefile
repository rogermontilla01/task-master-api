TARGET = task-master-api

build:
	go build -o $(TARGET) ./cmd/main.go

clean:
	rm -f $(TARGET)

dev:
	air

docker:
	docker-compose up -d