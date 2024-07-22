# commands for SERVER

run-s:
	cd server && docker-compose exec app go run server.go

test:
	cd server && docker-compose exec app go test ./test/... -v

test-srv:
	cd server && docker-compose exec app go test ./graph/service/... -v

test-coverage:
	cd server && docker-compose exec app go test ./test/... -v -cover -coverprofile=coverage.out -covermode=set

generate:
	cd server && docker-compose exec app go run github.com/99designs/gqlgen generate

start-ctn:
	cd server && docker-compose up -d

stop-ctn:
	cd server && docker-compose down

build-ctn:
	cd server && docker-compose build

# commands for CLIENT

run-c:
	cd client && pnpm run dev

codegen-c:
	cd client && pnpm run codegen

# GQL Schema

update-schema:
	cp -r -f graphql client && cd client && pnpm run codegen

