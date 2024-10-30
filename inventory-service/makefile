POSTGRES_URL= postgres://postgres:secret@localhost:5434/product-inventory?sslmode=disable

build:
	go build -o inventoryService ./cmd/*.go
run: build 
	./inventoryService

dbmigrateup:
	migrate -database ${POSTGRES_URL} -path db/migrations up
dbmigratedown:
	migrate -database ${POSTGRES_URL} -path db/migrations down