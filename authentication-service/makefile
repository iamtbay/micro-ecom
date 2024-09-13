POSTGRES_URL= postgres://postgres:secret@localhost:5432/users?sslmode=disable
run: build
	./authService
build:
	go build -o authService ./cmd/*.go

dbmigrateup:
	migrate -database ${POSTGRES_URL} -path db/migrations up
dbmigratedown:
	migrate -database ${POSTGRES_URL} -path db/migrations down
