migration:
	migrate -path ./storage/schema -database 'postgres://postgres:posgres@0.0.0.0:5432/postgres?sslmode=disable' up

build:
	docker-compose build

run:
	docker-compose up