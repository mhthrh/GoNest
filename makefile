test:
	go test ./... -v -bench . -failfast -cover -count=1
build:
	protoc --go-grpc_out=api/proto/city/v1  --go_out=api/proto/city/v1  api/proto/city/v1/*.proto
	protoc --go-grpc_out=api/proto/country/v1  --go_out=api/proto/country/v1  api/proto/country/v1/*.proto
	protoc --go-grpc_out=api/proto/address/v1  --go_out=api/proto/address/v1  api/proto/address/v1/*.proto
	protoc --go-grpc_out=api/proto/customer/v1  --go_out=api/proto/customer/v1  api/proto/customer/v1/*.proto
bb:
	protoc --go_out=model/proto/gen/golang \
	  --go-grpc_out=model/proto/gen/golang \
	  model/**/grpc/v1/*.proto
kir:
	protoc --go_out=. --go_opt=paths=source_relative \
      model/**/grpc/**/*.proto