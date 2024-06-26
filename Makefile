# commands for SERVER

run:
	cd server && docker-compose exec app go run server.go

test:
	cd server && docker-compose exec app go test ./test/... -v

test-srv:
	cd server && docker-compose exec app go test ./graph/service/... -v

test-coverage:
	cd server && docker-compose exec app go test ./test/... -v -cover -coverprofile=coverage.out -covermode=set

generate:
	cd server && docker-compose exec app go run github.com/99designs/gqlgen generate

start:
	cd server && docker-compose up -d

stop:
	cd server && docker-compose down

build:
	cd server && docker-compose build

