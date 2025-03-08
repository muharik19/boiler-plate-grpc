image:
	docker images

container:
	docker ps -a

volume:
	docker volume ls

dockerpostgres:
	docker run --name postgres-dashboard-ui -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret@123 -v postgres-data-dashboard-ui:/var/lib/postgresql/data -p 5432:5432 -d postgres

volumepg:
	docker volume inspect postgres-data-dashboard-ui

execpostgres:
	winpty docker exec -it postgres-dashboard-ui psql -U root

proto:
	protoc -I ./api/grpc/api/proto \
	--go_out ./api/grpc/api/pb --go_opt paths=source_relative \
	--go-grpc_out ./api/grpc/api/pb --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./api/grpc/api/pb --grpc-gateway_opt paths=source_relative \
	./api/grpc/api/proto/v1/role/role.proto \
	--go-grpc_opt=require_unimplemented_servers=false \
	--swagger_out ./api/grpc/api/proto/swagger \

server:
	go run main.go

.PHONY: image container volume dockerpostgres volumepg execpostgres proto server