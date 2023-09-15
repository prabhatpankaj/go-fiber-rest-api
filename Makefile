.PHONY: clean test security build run

APP_NAME = apiserver
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = postgres://postgres:password@localhost/postgres?sslmode=disable

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

docker.run: docker.network docker.postgres swag docker.go_fiber_rest_api migrate.up

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.go_fiber_rest_api.build:
	docker build -t go_fiber_rest_api .

docker.go_fiber_rest_api.remove:
	docker rm -f dev-go_fiber_rest_api && docker rmi go_fiber_rest_api


docker.go_fiber_rest_api: docker.go_fiber_rest_api.build
	docker run --rm -d \
		--name dev-go_fiber_rest_api \
		--network dev-network \
		-p 5000:5000 \
		go_fiber_rest_api
docker.go_fiber_rest_api.build-run: docker.go_fiber_rest_api.remove docker.go_fiber_rest_api.build
	docker run --rm -d \
		--name dev-go_fiber_rest_api \
		--network dev-network \
		-p 5000:5000 \
		go_fiber_rest_api
		
docker.postgres:
	docker run --rm -d \
		--name dev-postgres \
		--network dev-network \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=postgres \
		-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres

docker.stop: docker.stop.go_fiber_rest_api docker.stop.postgres

docker.stop.go_fiber_rest_api:
	docker stop dev-go_fiber_rest_api

docker.stop.postgres:
	docker stop dev-postgres

swag:
	swag init