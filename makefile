test:
	go test ./... -v -bench . -failfast -cover -count=1
grpc-build:
	protoc --go_out=. --go_opt=paths=source_relative \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
           model/**/grpc/**/*.proto