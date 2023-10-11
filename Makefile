server:
	go run cmd/server/main.go

mail:
	go run cmd/mail/main.go

####### Migrate #######

migrate-up:
	 migrate -path migration -database "postgres://postgres:postgres@localhost:5435/notification?sslmode=disable" up

migrate-down:
	 migrate -path migration -database "postgres://postgres:postgres@localhost:5435/notification?sslmode=disable" down

migrate-force:
	 migrate -path migration -database "postgres://postgres:postgres@localhost:5435/notification?sslmode=disable" force 1

#######  Docker #######

docker-up:
	docker compose -f docker-compose.yaml up

docker-down:
	docker compose -f docker-compose.yaml down

docker-build:
	docker compose up --build

docker-clear:
	docker system prune