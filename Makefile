server:
	go run cmd/server/main.go

####### Migrate #######

migrate-up:
	 migrate -path migration -database "postgres://postgres:1234@localhost:5432/notification?sslmode=disable" up

migrate-down:
	migrate -path migration -database "postgres://postgres:1234@localhost:5432/notification?sslmode=disable" down