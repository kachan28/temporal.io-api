// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/operatorservice/v1/service.proto

package operatorservice

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("temporal/api/operatorservice/v1/service.proto", fileDescriptor_a603f8bc80c66d47)
}

var fileDescriptor_a603f8bc80c66d47 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0xf7, 0xf5, 0xe0, 0x61, 0x2c, 0x14, 0x46, 0x51, 0xe9, 0x61, 0x0a, 0xb9, 0xbb, 0x21,
	0x6a, 0x55, 0x36, 0xd5, 0x76, 0xab, 0xde, 0xd4, 0x96, 0x56, 0x2b, 0x78, 0x91, 0xcd, 0xee, 0x53,
	0x07, 0x93, 0xce, 0x38, 0x33, 0x89, 0x5e, 0x04, 0x4f, 0x5e, 0x15, 0xc1, 0x93, 0x20, 0x78, 0x13,
	0x3f, 0x89, 0xc7, 0x1c, 0xeb, 0xcd, 0x6c, 0x40, 0xc4, 0x53, 0x3f, 0x82, 0xac, 0xbb, 0xb3, 0xb4,
	0x9b, 0x95, 0xdd, 0xac, 0xb7, 0x90, 0xf9, 0xff, 0xfe, 0xf3, 0x0b, 0x79, 0xc3, 0x23, 0xe7, 0x0d,
	0x0e, 0xa4, 0x50, 0x41, 0xbf, 0x1d, 0x48, 0xde, 0x16, 0x12, 0x55, 0x60, 0x84, 0xd2, 0xa8, 0x46,
	0x3c, 0xc4, 0xf6, 0xa8, 0xd3, 0xce, 0x3e, 0xba, 0x52, 0x09, 0x23, 0xe8, 0x8a, 0x8d, 0xbb, 0x81,
	0xe4, 0x6e, 0x21, 0xee, 0x8e, 0x3a, 0xcb, 0x97, 0xab, 0xfa, 0x14, 0x3e, 0x1f, 0xa2, 0x36, 0x8f,
	0x14, 0x6a, 0x29, 0xf6, 0x75, 0x56, 0x7c, 0xe1, 0xfb, 0x49, 0xb2, 0xb4, 0x95, 0xa5, 0x77, 0xd3,
	0x34, 0x7d, 0x0f, 0xe4, 0x94, 0x1f, 0x45, 0xbb, 0x18, 0xa8, 0xf0, 0xa9, 0x6f, 0x8c, 0xe2, 0xbd,
	0xa1, 0x41, 0x4d, 0xbb, 0x6e, 0x85, 0x85, 0x5b, 0x42, 0xed, 0xa4, 0xf7, 0x2e, 0xaf, 0x35, 0x83,
	0x53, 0xd9, 0x96, 0x43, 0x3f, 0x02, 0x39, 0xb3, 0x83, 0x03, 0x31, 0xc2, 0x19, 0xaf, 0xeb, 0x95,
	0xd5, 0xe5, 0xa0, 0x55, 0x5b, 0x6f, 0xcc, 0xe7, 0x76, 0x1f, 0x80, 0x9c, 0xbe, 0xcd, 0xb5, 0x99,
	0x71, 0xab, 0xfe, 0xd9, 0x65, 0x98, 0x35, 0xbb, 0xd6, 0x90, 0xce, 0xbd, 0xde, 0x00, 0x59, 0xba,
	0x89, 0x7d, 0x34, 0x78, 0x37, 0x18, 0xa0, 0x96, 0x41, 0x88, 0xf4, 0x4a, 0x65, 0x69, 0x81, 0xb0,
	0x36, 0x57, 0xe7, 0x07, 0x73, 0x91, 0x4f, 0x40, 0xce, 0xa6, 0xa7, 0x0f, 0x84, 0x7a, 0xf6, 0xb8,
	0x2f, 0x5e, 0xdc, 0x7a, 0x89, 0xe1, 0xd0, 0x70, 0xb1, 0x4f, 0xd7, 0x6b, 0xf6, 0xce, 0x90, 0x56,
	0x6c, 0xa3, 0x79, 0x41, 0x2e, 0xf8, 0x19, 0xc8, 0x39, 0x3f, 0x8a, 0xb6, 0xd4, 0x7d, 0x19, 0x05,
	0x06, 0x93, 0x7f, 0xdc, 0xe0, 0x8d, 0xfe, 0x50, 0x1b, 0x54, 0x74, 0xa3, 0xce, 0xf0, 0x96, 0xa2,
	0x56, 0xd1, 0xff, 0x8f, 0x86, 0xdc, 0x31, 0x79, 0x98, 0xe9, 0x28, 0x1e, 0xd7, 0xeb, 0xd6, 0x1c,
	0xe0, 0x52, 0xb3, 0xb5, 0x66, 0x70, 0x61, 0xc4, 0x74, 0xa8, 0x78, 0x2f, 0x17, 0xaa, 0x33, 0x62,
	0xc7, 0x88, 0x79, 0x46, 0xac, 0x00, 0xe6, 0x22, 0xaf, 0xc8, 0x62, 0xf2, 0x1a, 0xb2, 0x03, 0x4d,
	0x2f, 0xd5, 0x7a, 0x3c, 0x36, 0x6e, 0x0d, 0x56, 0xe7, 0xa4, 0xf2, 0xeb, 0xdf, 0x02, 0xa1, 0x47,
	0x8e, 0xee, 0xe0, 0xa0, 0x97, 0x58, 0x78, 0xf3, 0xf4, 0x65, 0x90, 0x75, 0xe9, 0x36, 0x62, 0xad,
	0xd1, 0xe6, 0x4f, 0x18, 0x4f, 0x98, 0x73, 0x30, 0x61, 0xce, 0xe1, 0x84, 0xc1, 0xeb, 0x98, 0xc1,
	0x97, 0x98, 0xc1, 0xb7, 0x98, 0xc1, 0x38, 0x66, 0xf0, 0x23, 0x66, 0xf0, 0x2b, 0x66, 0xce, 0x61,
	0xcc, 0xe0, 0xdd, 0x94, 0x39, 0xe3, 0x29, 0x73, 0x0e, 0xa6, 0xcc, 0x21, 0x2d, 0x2e, 0xaa, 0xee,
	0xdd, 0x5c, 0xcc, 0x76, 0xc5, 0x76, 0xb2, 0x44, 0xb6, 0xe1, 0xe1, 0xea, 0x93, 0x23, 0x0c, 0x17,
	0xff, 0x58, 0x42, 0xdd, 0xc2, 0x57, 0x5f, 0x17, 0x56, 0xee, 0x59, 0xc8, 0x97, 0xdc, 0x2d, 0x2c,
	0x22, 0x77, 0xaf, 0xf3, 0x7b, 0xa1, 0x65, 0x13, 0x9e, 0xe7, 0x4b, 0xee, 0x79, 0x85, 0x8c, 0xe7,
	0xed, 0x75, 0x7a, 0x27, 0xfe, 0xee, 0xb2, 0x8b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xac, 0xbd,
	0xdc, 0x06, 0x55, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OperatorServiceClient is the client API for OperatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperatorServiceClient interface {
	// AddSearchAttributes add custom search attributes.
	//
	// Returns ALREADY_EXISTS status code if a Search Attribute with any of the specified names already exists
	// Returns INTERNAL status code with temporal.api.errordetails.v1.SystemWorkflowFailure in Error Details if registration process fails,
	AddSearchAttributes(ctx context.Context, in *AddSearchAttributesRequest, opts ...grpc.CallOption) (*AddSearchAttributesResponse, error)
	// RemoveSearchAttributes removes custom search attributes.
	//
	// Returns NOT_FOUND status code if a Search Attribute with any of the specified names is not registered
	RemoveSearchAttributes(ctx context.Context, in *RemoveSearchAttributesRequest, opts ...grpc.CallOption) (*RemoveSearchAttributesResponse, error)
	// ListSearchAttributes returns comprehensive information about search attributes.
	ListSearchAttributes(ctx context.Context, in *ListSearchAttributesRequest, opts ...grpc.CallOption) (*ListSearchAttributesResponse, error)
	// DeleteNamespace synchronously deletes a namespace and asynchronously reclaims all namespace resources.
	// (-- api-linter: core::0135::method-signature=disabled
	//     aip.dev/not-precedent: DeleteNamespace RPC doesn't follow Google API format. --)
	// (-- api-linter: core::0135::response-message-name=disabled
	//     aip.dev/not-precedent: DeleteNamespace RPC doesn't follow Google API format. --)
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error)
	// DeleteWorkflowExecution asynchronously deletes a specific Workflow Execution (when
	// WorkflowExecution.run_id is provided) or the latest Workflow Execution (when
	// WorkflowExecution.run_id is not provided). If the Workflow Execution is Running, it will be
	// terminated before deletion.
	// (-- api-linter: core::0135::method-signature=disabled
	//     aip.dev/not-precedent: DeleteNamespace RPC doesn't follow Google API format. --)
	// (-- api-linter: core::0135::response-message-name=disabled
	//     aip.dev/not-precedent: DeleteNamespace RPC doesn't follow Google API format. --)
	DeleteWorkflowExecution(ctx context.Context, in *DeleteWorkflowExecutionRequest, opts ...grpc.CallOption) (*DeleteWorkflowExecutionResponse, error)
	// AddOrUpdateRemoteCluster adds or updates remote cluster.
	AddOrUpdateRemoteCluster(ctx context.Context, in *AddOrUpdateRemoteClusterRequest, opts ...grpc.CallOption) (*AddOrUpdateRemoteClusterResponse, error)
	// RemoveRemoteCluster removes remote cluster.
	RemoveRemoteCluster(ctx context.Context, in *RemoveRemoteClusterRequest, opts ...grpc.CallOption) (*RemoveRemoteClusterResponse, error)
	// DescribeCluster returns information about Temporal cluster.
	DescribeCluster(ctx context.Context, in *DescribeClusterRequest, opts ...grpc.CallOption) (*DescribeClusterResponse, error)
	// ListClusters returns information about Temporal clusters.
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	// ListClusterMembers returns information about Temporal cluster members.
	ListClusterMembers(ctx context.Context, in *ListClusterMembersRequest, opts ...grpc.CallOption) (*ListClusterMembersResponse, error)
}

type operatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewOperatorServiceClient(cc *grpc.ClientConn) OperatorServiceClient {
	return &operatorServiceClient{cc}
}

func (c *operatorServiceClient) AddSearchAttributes(ctx context.Context, in *AddSearchAttributesRequest, opts ...grpc.CallOption) (*AddSearchAttributesResponse, error) {
	out := new(AddSearchAttributesResponse)
	err := c.cc.Invoke(ctx, "/temporal.api.operatorservice.v1.OperatorService/AddSearchAttributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) RemoveSearchAttributes(ctx context.Context, in *RemoveSearchAttributesRequest, opts ...grpc.CallOption) (*RemoveSearchAttributesResponse, error) {
	out := new(RemoveSearchAttributesResponse)
	err := c.cc.Invoke(ctx, "/temporal.api.operatorservice.v1.OperatorService/RemoveSearchAttributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) ListSearchAttributes(ctx context.Context, in *ListSearchAttributesRequest, opts ...grpc.CallOption) (*ListSearchAttributesResponse, error) {
	out := new(ListSearchAttributesResponse)
	err := c.cc.Invoke(ctx, "/temporal.api.operatorservice.v1.OperatorService/ListSearchAttributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error) {
	out := new(DeleteNamespaceResponse)
	err := c.cc.Invoke(ctx, "/temporal.api.operatorservice.v1.OperatorService/DeleteNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) DeleteWorkflowExecution(ctx context.Context, in *DeleteWorkflowExecutionRequest, opts ...grpc.CallOption) (*DeleteWorkflowExecutionResponse, error) {
	out := new(DeleteWorkflowExecutionResponse)
	err := c.cc.Invoke(ctx, "/temporal.api.operatorservice.v1.OperatorService/DeleteWorkflowExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) AddOrUpdateRemoteCluster(ctx context.Context, in *AddOrUpdateRemoteClusterRequest, opts ...grpc.CallOption) (*AddOrUpdateRemoteClusterResponse, error) {
	out := new(AddOrUpdateRemoteClusterResponse)
	err := c.cc.Invoke(ctx, "/temporal.api.operatorservice.v1.OperatorService/AddOrUpdateRemoteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) RemoveRemoteCluster(ctx context.Context, in *RemoveRemoteClusterRequest, opts ...grpc.CallOption) (*RemoveRemoteClusterResponse, error) {
	out := new(RemoveRemoteClusterResponse)
	err := c.cc.Invoke(ctx, "/temporal.api.operatorservice.v1.OperatorService/RemoveRemoteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) DescribeCluster(ctx context.Context, in *DescribeClusterRequest, opts ...grpc.CallOption) (*DescribeClusterResponse, error) {
	out := new(DescribeClusterResponse)
	err := c.cc.Invoke(ctx, "/temporal.api.operatorservice.v1.OperatorService/DescribeCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := c.cc.Invoke(ctx, "/temporal.api.operatorservice.v1.OperatorService/ListClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) ListClusterMembers(ctx context.Context, in *ListClusterMembersRequest, opts ...grpc.CallOption) (*ListClusterMembersResponse, error) {
	out := new(ListClusterMembersResponse)
	err := c.cc.Invoke(ctx, "/temporal.api.operatorservice.v1.OperatorService/ListClusterMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatorServiceServer is the server API for OperatorService service.
type OperatorServiceServer interface {
	// AddSearchAttributes add custom search attributes.
	//
	// Returns ALREADY_EXISTS status code if a Search Attribute with any of the specified names already exists
	// Returns INTERNAL status code with temporal.api.errordetails.v1.SystemWorkflowFailure in Error Details if registration process fails,
	AddSearchAttributes(context.Context, *AddSearchAttributesRequest) (*AddSearchAttributesResponse, error)
	// RemoveSearchAttributes removes custom search attributes.
	//
	// Returns NOT_FOUND status code if a Search Attribute with any of the specified names is not registered
	RemoveSearchAttributes(context.Context, *RemoveSearchAttributesRequest) (*RemoveSearchAttributesResponse, error)
	// ListSearchAttributes returns comprehensive information about search attributes.
	ListSearchAttributes(context.Context, *ListSearchAttributesRequest) (*ListSearchAttributesResponse, error)
	// DeleteNamespace synchronously deletes a namespace and asynchronously reclaims all namespace resources.
	// (-- api-linter: core::0135::method-signature=disabled
	//     aip.dev/not-precedent: DeleteNamespace RPC doesn't follow Google API format. --)
	// (-- api-linter: core::0135::response-message-name=disabled
	//     aip.dev/not-precedent: DeleteNamespace RPC doesn't follow Google API format. --)
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error)
	// DeleteWorkflowExecution asynchronously deletes a specific Workflow Execution (when
	// WorkflowExecution.run_id is provided) or the latest Workflow Execution (when
	// WorkflowExecution.run_id is not provided). If the Workflow Execution is Running, it will be
	// terminated before deletion.
	// (-- api-linter: core::0135::method-signature=disabled
	//     aip.dev/not-precedent: DeleteNamespace RPC doesn't follow Google API format. --)
	// (-- api-linter: core::0135::response-message-name=disabled
	//     aip.dev/not-precedent: DeleteNamespace RPC doesn't follow Google API format. --)
	DeleteWorkflowExecution(context.Context, *DeleteWorkflowExecutionRequest) (*DeleteWorkflowExecutionResponse, error)
	// AddOrUpdateRemoteCluster adds or updates remote cluster.
	AddOrUpdateRemoteCluster(context.Context, *AddOrUpdateRemoteClusterRequest) (*AddOrUpdateRemoteClusterResponse, error)
	// RemoveRemoteCluster removes remote cluster.
	RemoveRemoteCluster(context.Context, *RemoveRemoteClusterRequest) (*RemoveRemoteClusterResponse, error)
	// DescribeCluster returns information about Temporal cluster.
	DescribeCluster(context.Context, *DescribeClusterRequest) (*DescribeClusterResponse, error)
	// ListClusters returns information about Temporal clusters.
	ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	// ListClusterMembers returns information about Temporal cluster members.
	ListClusterMembers(context.Context, *ListClusterMembersRequest) (*ListClusterMembersResponse, error)
}

// UnimplementedOperatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOperatorServiceServer struct {
}

func (*UnimplementedOperatorServiceServer) AddSearchAttributes(ctx context.Context, req *AddSearchAttributesRequest) (*AddSearchAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSearchAttributes not implemented")
}
func (*UnimplementedOperatorServiceServer) RemoveSearchAttributes(ctx context.Context, req *RemoveSearchAttributesRequest) (*RemoveSearchAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSearchAttributes not implemented")
}
func (*UnimplementedOperatorServiceServer) ListSearchAttributes(ctx context.Context, req *ListSearchAttributesRequest) (*ListSearchAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSearchAttributes not implemented")
}
func (*UnimplementedOperatorServiceServer) DeleteNamespace(ctx context.Context, req *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNamespace not implemented")
}
func (*UnimplementedOperatorServiceServer) DeleteWorkflowExecution(ctx context.Context, req *DeleteWorkflowExecutionRequest) (*DeleteWorkflowExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkflowExecution not implemented")
}
func (*UnimplementedOperatorServiceServer) AddOrUpdateRemoteCluster(ctx context.Context, req *AddOrUpdateRemoteClusterRequest) (*AddOrUpdateRemoteClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrUpdateRemoteCluster not implemented")
}
func (*UnimplementedOperatorServiceServer) RemoveRemoteCluster(ctx context.Context, req *RemoveRemoteClusterRequest) (*RemoveRemoteClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRemoteCluster not implemented")
}
func (*UnimplementedOperatorServiceServer) DescribeCluster(ctx context.Context, req *DescribeClusterRequest) (*DescribeClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCluster not implemented")
}
func (*UnimplementedOperatorServiceServer) ListClusters(ctx context.Context, req *ListClustersRequest) (*ListClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClusters not implemented")
}
func (*UnimplementedOperatorServiceServer) ListClusterMembers(ctx context.Context, req *ListClusterMembersRequest) (*ListClusterMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClusterMembers not implemented")
}

func RegisterOperatorServiceServer(s *grpc.Server, srv OperatorServiceServer) {
	s.RegisterService(&_OperatorService_serviceDesc, srv)
}

func _OperatorService_AddSearchAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSearchAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).AddSearchAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.api.operatorservice.v1.OperatorService/AddSearchAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).AddSearchAttributes(ctx, req.(*AddSearchAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_RemoveSearchAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSearchAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).RemoveSearchAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.api.operatorservice.v1.OperatorService/RemoveSearchAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).RemoveSearchAttributes(ctx, req.(*RemoveSearchAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_ListSearchAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSearchAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).ListSearchAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.api.operatorservice.v1.OperatorService/ListSearchAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).ListSearchAttributes(ctx, req.(*ListSearchAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_DeleteNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).DeleteNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.api.operatorservice.v1.OperatorService/DeleteNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).DeleteNamespace(ctx, req.(*DeleteNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_DeleteWorkflowExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkflowExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).DeleteWorkflowExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.api.operatorservice.v1.OperatorService/DeleteWorkflowExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).DeleteWorkflowExecution(ctx, req.(*DeleteWorkflowExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_AddOrUpdateRemoteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateRemoteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).AddOrUpdateRemoteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.api.operatorservice.v1.OperatorService/AddOrUpdateRemoteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).AddOrUpdateRemoteCluster(ctx, req.(*AddOrUpdateRemoteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_RemoveRemoteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRemoteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).RemoveRemoteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.api.operatorservice.v1.OperatorService/RemoveRemoteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).RemoveRemoteCluster(ctx, req.(*RemoveRemoteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_DescribeCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).DescribeCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.api.operatorservice.v1.OperatorService/DescribeCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).DescribeCluster(ctx, req.(*DescribeClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.api.operatorservice.v1.OperatorService/ListClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).ListClusters(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_ListClusterMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).ListClusterMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.api.operatorservice.v1.OperatorService/ListClusterMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).ListClusterMembers(ctx, req.(*ListClusterMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OperatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "temporal.api.operatorservice.v1.OperatorService",
	HandlerType: (*OperatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSearchAttributes",
			Handler:    _OperatorService_AddSearchAttributes_Handler,
		},
		{
			MethodName: "RemoveSearchAttributes",
			Handler:    _OperatorService_RemoveSearchAttributes_Handler,
		},
		{
			MethodName: "ListSearchAttributes",
			Handler:    _OperatorService_ListSearchAttributes_Handler,
		},
		{
			MethodName: "DeleteNamespace",
			Handler:    _OperatorService_DeleteNamespace_Handler,
		},
		{
			MethodName: "DeleteWorkflowExecution",
			Handler:    _OperatorService_DeleteWorkflowExecution_Handler,
		},
		{
			MethodName: "AddOrUpdateRemoteCluster",
			Handler:    _OperatorService_AddOrUpdateRemoteCluster_Handler,
		},
		{
			MethodName: "RemoveRemoteCluster",
			Handler:    _OperatorService_RemoveRemoteCluster_Handler,
		},
		{
			MethodName: "DescribeCluster",
			Handler:    _OperatorService_DescribeCluster_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _OperatorService_ListClusters_Handler,
		},
		{
			MethodName: "ListClusterMembers",
			Handler:    _OperatorService_ListClusterMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "temporal/api/operatorservice/v1/service.proto",
}
