// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package plugin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ClusterClient is the client API for Cluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterClient interface {
	// HealthCheck is the endpoint for readiness probe
	HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetVersion gets the given cluster's version
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*ClusterVersion, error)
	// GetClusterStatus gets the given cluster's status
	GetClusterStatus(ctx context.Context, in *GetClusterStatusRequest, opts ...grpc.CallOption) (*ClusterStatus, error)
	// GetOperationStatus gets the given operation's status
	GetOperationStatus(ctx context.Context, in *GetOperationStatusRequest, opts ...grpc.CallOption) (*OperationStatus, error)
	// ServiceIn perform the operation to route to the given cluster
	ServiceIn(ctx context.Context, in *ServiceInRequest, opts ...grpc.CallOption) (*Operation, error)
	// ServiceOut perform the operation to make the given cluster not be routed
	ServiceOut(ctx context.Context, in *ServiceOutRequest, opts ...grpc.CallOption) (*Operation, error)
	// UpgradeMaster perform the operation to upgrade master nodes in the given cluster.
	UpgradeMaster(ctx context.Context, in *MasterVersion, opts ...grpc.CallOption) (*Operation, error)
	// UpgradeNodePool perform the operation to upgrade worker nodes in the given cluster.
	UpgradeNodePool(ctx context.Context, in *NodePoolVersion, opts ...grpc.CallOption) (*Operation, error)
}

type clusterClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterClient(cc grpc.ClientConnInterface) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/plugin.Cluster/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*ClusterVersion, error) {
	out := new(ClusterVersion)
	err := c.cc.Invoke(ctx, "/plugin.Cluster/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) GetClusterStatus(ctx context.Context, in *GetClusterStatusRequest, opts ...grpc.CallOption) (*ClusterStatus, error) {
	out := new(ClusterStatus)
	err := c.cc.Invoke(ctx, "/plugin.Cluster/GetClusterStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) GetOperationStatus(ctx context.Context, in *GetOperationStatusRequest, opts ...grpc.CallOption) (*OperationStatus, error) {
	out := new(OperationStatus)
	err := c.cc.Invoke(ctx, "/plugin.Cluster/GetOperationStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) ServiceIn(ctx context.Context, in *ServiceInRequest, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := c.cc.Invoke(ctx, "/plugin.Cluster/ServiceIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) ServiceOut(ctx context.Context, in *ServiceOutRequest, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := c.cc.Invoke(ctx, "/plugin.Cluster/ServiceOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) UpgradeMaster(ctx context.Context, in *MasterVersion, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := c.cc.Invoke(ctx, "/plugin.Cluster/UpgradeMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) UpgradeNodePool(ctx context.Context, in *NodePoolVersion, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := c.cc.Invoke(ctx, "/plugin.Cluster/UpgradeNodePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServer is the server API for Cluster service.
// All implementations must embed UnimplementedClusterServer
// for forward compatibility
type ClusterServer interface {
	// HealthCheck is the endpoint for readiness probe
	HealthCheck(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// GetVersion gets the given cluster's version
	GetVersion(context.Context, *GetVersionRequest) (*ClusterVersion, error)
	// GetClusterStatus gets the given cluster's status
	GetClusterStatus(context.Context, *GetClusterStatusRequest) (*ClusterStatus, error)
	// GetOperationStatus gets the given operation's status
	GetOperationStatus(context.Context, *GetOperationStatusRequest) (*OperationStatus, error)
	// ServiceIn perform the operation to route to the given cluster
	ServiceIn(context.Context, *ServiceInRequest) (*Operation, error)
	// ServiceOut perform the operation to make the given cluster not be routed
	ServiceOut(context.Context, *ServiceOutRequest) (*Operation, error)
	// UpgradeMaster perform the operation to upgrade master nodes in the given cluster.
	UpgradeMaster(context.Context, *MasterVersion) (*Operation, error)
	// UpgradeNodePool perform the operation to upgrade worker nodes in the given cluster.
	UpgradeNodePool(context.Context, *NodePoolVersion) (*Operation, error)
	mustEmbedUnimplementedClusterServer()
}

// UnimplementedClusterServer must be embedded to have forward compatible implementations.
type UnimplementedClusterServer struct {
}

func (UnimplementedClusterServer) HealthCheck(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedClusterServer) GetVersion(context.Context, *GetVersionRequest) (*ClusterVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedClusterServer) GetClusterStatus(context.Context, *GetClusterStatusRequest) (*ClusterStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterStatus not implemented")
}
func (UnimplementedClusterServer) GetOperationStatus(context.Context, *GetOperationStatusRequest) (*OperationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperationStatus not implemented")
}
func (UnimplementedClusterServer) ServiceIn(context.Context, *ServiceInRequest) (*Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceIn not implemented")
}
func (UnimplementedClusterServer) ServiceOut(context.Context, *ServiceOutRequest) (*Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceOut not implemented")
}
func (UnimplementedClusterServer) UpgradeMaster(context.Context, *MasterVersion) (*Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeMaster not implemented")
}
func (UnimplementedClusterServer) UpgradeNodePool(context.Context, *NodePoolVersion) (*Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeNodePool not implemented")
}
func (UnimplementedClusterServer) mustEmbedUnimplementedClusterServer() {}

// UnsafeClusterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServer will
// result in compilation errors.
type UnsafeClusterServer interface {
	mustEmbedUnimplementedClusterServer()
}

func RegisterClusterServer(s grpc.ServiceRegistrar, srv ClusterServer) {
	s.RegisterService(&_Cluster_serviceDesc, srv)
}

func _Cluster_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.Cluster/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).HealthCheck(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.Cluster/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_GetClusterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).GetClusterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.Cluster/GetClusterStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).GetClusterStatus(ctx, req.(*GetClusterStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_GetOperationStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).GetOperationStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.Cluster/GetOperationStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).GetOperationStatus(ctx, req.(*GetOperationStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_ServiceIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).ServiceIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.Cluster/ServiceIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).ServiceIn(ctx, req.(*ServiceInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_ServiceOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).ServiceOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.Cluster/ServiceOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).ServiceOut(ctx, req.(*ServiceOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_UpgradeMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MasterVersion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).UpgradeMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.Cluster/UpgradeMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).UpgradeMaster(ctx, req.(*MasterVersion))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_UpgradeNodePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodePoolVersion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).UpgradeNodePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugin.Cluster/UpgradeNodePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).UpgradeNodePool(ctx, req.(*NodePoolVersion))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "plugin.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Cluster_HealthCheck_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Cluster_GetVersion_Handler,
		},
		{
			MethodName: "GetClusterStatus",
			Handler:    _Cluster_GetClusterStatus_Handler,
		},
		{
			MethodName: "GetOperationStatus",
			Handler:    _Cluster_GetOperationStatus_Handler,
		},
		{
			MethodName: "ServiceIn",
			Handler:    _Cluster_ServiceIn_Handler,
		},
		{
			MethodName: "ServiceOut",
			Handler:    _Cluster_ServiceOut_Handler,
		},
		{
			MethodName: "UpgradeMaster",
			Handler:    _Cluster_UpgradeMaster_Handler,
		},
		{
			MethodName: "UpgradeNodePool",
			Handler:    _Cluster_UpgradeNodePool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/plugin/plugin_api.proto",
}
