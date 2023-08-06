build: build-go build-typescript build-swift

build-go:
	protoc --go_out=golang --go-grpc_out=golang --go_opt=paths=import --go-grpc_opt=paths=import proto/*.proto


build-typescript:
	./scripts/build-js.sh

build-swift:
	protoc --grpc-swift_opt=Client=true,Server=true --grpc-swift_out=swift --proto_path=. --swift_opt=Visibility=Public --swift_out=swift **/*.proto
	mv swift/proto/* swift
	rm -rf swift/proto
