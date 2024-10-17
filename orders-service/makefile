POSTGRES_URL=postgres://postgres:secret@localhost:5433/orders-adresses?sslmode=disable

build:
	go build -o ordersAdressesApp ./cmd/*.go

run: build
	./ordersAdressesApp
	
dbmigrateup:
	migrate -database ${POSTGRES_URL} -path db/migrations up

dbmigratedown:
	migrate -database ${POSTGRES_URL} -path db/migrations down