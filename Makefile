build-go:
	protoc --go_out=golang --go-grpc_out=golang --go_opt=paths=import --go-grpc_opt=paths=import proto/*.proto


build-typescript:
	./scripts/build-js.sh
