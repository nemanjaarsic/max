proto:
	cd ./db-svc && protoc --go_out=./proto ./proto/*.proto --go-grpc_out=./proto
	cd ./idempotency-svc && protoc --go_out=./proto ./proto/*.proto --go-grpc_out=./proto
	cd ./web-svc && protoc --go_out=./proto ./proto/*.proto --go-grpc_out=./proto
.PHONY: proto