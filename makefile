test:
	go test ./... -v -bench . -failfast -cover -count=1
build:
	protoc --go-grpc_out=model/city  --go_out=model/city  model/city/*.proto
	protoc --go-grpc_out=model/country  --go_out=model/country  model/country/*.proto
	protoc --go-grpc_out=model/address  --go_out=model/address  model/address/*.proto
	protoc --go-grpc_out=model/customer  --go_out=model/customer  model/customer/*.proto
	protoc --go-grpc_out=model/city  --go_out=model/city  model/city/*.proto

