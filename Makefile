server:
	go run cmd/server/main.go

mail:
	go run cmd/mail/main.go

####### Migrate #######

migrate-up:
	 migrate -path migration -database "postgres://postgres:postgres@localhost:5435/notification?sslmode=disable" up

migrate-down:
	 migrate -path migration -database "postgres://postgres:postgres@localhost:5435/notification?sslmode=disable" down


#######  Docker #######

docker-up:
	docker compose -f docker-compose.yaml up

docker-down:
	docker compose -f docker-compose.yaml down