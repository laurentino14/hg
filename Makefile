dev:
	docker-compose -f build/docker-compose.yml up air;

stop:
	docker-compose -f build/docker-compose.yml stop;


build: clean
	docker build -f build/Dockerfile -t sirius .;

preview: clean
	docker build -f build/Dockerfile -t sirius .;
	docker-compose -f build/docker-compose.yml up app;

migrate:
	docker-compose -f build/docker-compose.yml up migrate_db;

test/cover:
	go test ./... -coverprofile=coverage.out -covermode=count;

test:
	go test ./...;

clean:
	docker rmi sirius --force;

.PHONY: build clean