proto:
	protoc --go_out=./proto ./proto/*.proto --go-grpc_out=./proto
.PHONY: proto

proxy:
	protoc -I . --grpc-gateway_out pb --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true ./proto/*.proto
.PHONY: proxy