generate-go:
	go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc \
         google.golang.org/grpc \
         github.com/golang/mock/mockgen@v1.4.4
	protoc  \
		--go_out=./go/plugin \
		--go-grpc_out=./go/plugin \
    ./src/plugin/plugin_api.proto
	mv ./go/plugin/github.com/taisho6339/multicluster-upgrade-operator-proto/go/plugin/*.go ./go/plugin
	rm -r ./go/plugin/github.com
	mockgen -source ./go/plugin/plugin_api_grpc.pb.go -destination ./go/plugin/mock_grpc_client.go -package=plugin