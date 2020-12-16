# multicluster-upgrade-operator-proto

The proto files and generated files of plugin server APIs
for [multicluster upgrade operator](https://github.com/taisho6339/multicluster-upgrade-operator)

Now, only supports Go

## How to use

### gRPC Server Sample

```go
package server

import (
	"github.com/taisho6339/multicluster-upgrade-operator-proto/go/plugin"
)

func Start() {
	...
	sv := grpc.NewServer()
	c := NewYourServer()
	plugin.RegisterClusterServer(sv, c)
	...
}
```

### gRPC Client Sample

```go
package main

import (
	"context"
	"github.com/taisho6339/multicluster-upgrade-operator-proto/go/plugin"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("your_endpoint")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := plugin.NewClusterClient(conn)
	ctx := context.Background()
	req := &plugin.GetClusterStatusRequest{
		ClusterID: "your cluster id",
	}
	res, err := c.GetClusterStatus(ctx, req)
	if err != nil {
		log.Printf("%#v", res)
	}
}
```

### Mock Test Sample

```go
package main_test

import (
	"github.com/taisho6339/multicluster-upgrade-operator-proto/go/plugin"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestMain(t *testing.T) {
	...
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	c := plugin.NewMockClusterClient(ctrl)
	c.EXPECT().
		GetClusterStatus(gomock.Any(), gomock.Any()).
		Return(plugin.ClusterStatus{}, nil)
	...
}
```

## How to build

If you modified proto files, you have to run generate command.

```
make generate-go
```

## Versioning

This library follows [Semantic Versioning](https://semver.org/).

## License

Apache Version 2.0