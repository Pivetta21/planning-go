db_url = postgresql://postgres:1234@localhost:5432/planning_go?sslmode=disable

build:
	go mod download 
	go mod verify 
	go build -o ./build/bin -v ./cmd/api

start: build
	./build/bin

run:
	go mod tidy
	go run ./cmd/api

test:
	go test -v ./... -count=1

image:
	docker build -t planninggo .

postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1234 -d postgres:15.3-alpine

create-db:
	docker exec -it postgres15 createdb --username=postgres --owner=postgres planning_go

drop-db:
	docker exec -it postgres15 dropdb planning_go

.PHONY: build start run test image postgres create-db drop-db
