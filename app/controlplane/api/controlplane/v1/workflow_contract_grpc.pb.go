//
// Copyright 2023 The Chainloop Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: controlplane/v1/workflow_contract.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	WorkflowContractService_List_FullMethodName     = "/controlplane.v1.WorkflowContractService/List"
	WorkflowContractService_Create_FullMethodName   = "/controlplane.v1.WorkflowContractService/Create"
	WorkflowContractService_Update_FullMethodName   = "/controlplane.v1.WorkflowContractService/Update"
	WorkflowContractService_Describe_FullMethodName = "/controlplane.v1.WorkflowContractService/Describe"
	WorkflowContractService_Delete_FullMethodName   = "/controlplane.v1.WorkflowContractService/Delete"
)

// WorkflowContractServiceClient is the client API for WorkflowContractService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkflowContractServiceClient interface {
	List(ctx context.Context, in *WorkflowContractServiceListRequest, opts ...grpc.CallOption) (*WorkflowContractServiceListResponse, error)
	Create(ctx context.Context, in *WorkflowContractServiceCreateRequest, opts ...grpc.CallOption) (*WorkflowContractServiceCreateResponse, error)
	Update(ctx context.Context, in *WorkflowContractServiceUpdateRequest, opts ...grpc.CallOption) (*WorkflowContractServiceUpdateResponse, error)
	Describe(ctx context.Context, in *WorkflowContractServiceDescribeRequest, opts ...grpc.CallOption) (*WorkflowContractServiceDescribeResponse, error)
	Delete(ctx context.Context, in *WorkflowContractServiceDeleteRequest, opts ...grpc.CallOption) (*WorkflowContractServiceDeleteResponse, error)
}

type workflowContractServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkflowContractServiceClient(cc grpc.ClientConnInterface) WorkflowContractServiceClient {
	return &workflowContractServiceClient{cc}
}

func (c *workflowContractServiceClient) List(ctx context.Context, in *WorkflowContractServiceListRequest, opts ...grpc.CallOption) (*WorkflowContractServiceListResponse, error) {
	out := new(WorkflowContractServiceListResponse)
	err := c.cc.Invoke(ctx, WorkflowContractService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowContractServiceClient) Create(ctx context.Context, in *WorkflowContractServiceCreateRequest, opts ...grpc.CallOption) (*WorkflowContractServiceCreateResponse, error) {
	out := new(WorkflowContractServiceCreateResponse)
	err := c.cc.Invoke(ctx, WorkflowContractService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowContractServiceClient) Update(ctx context.Context, in *WorkflowContractServiceUpdateRequest, opts ...grpc.CallOption) (*WorkflowContractServiceUpdateResponse, error) {
	out := new(WorkflowContractServiceUpdateResponse)
	err := c.cc.Invoke(ctx, WorkflowContractService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowContractServiceClient) Describe(ctx context.Context, in *WorkflowContractServiceDescribeRequest, opts ...grpc.CallOption) (*WorkflowContractServiceDescribeResponse, error) {
	out := new(WorkflowContractServiceDescribeResponse)
	err := c.cc.Invoke(ctx, WorkflowContractService_Describe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowContractServiceClient) Delete(ctx context.Context, in *WorkflowContractServiceDeleteRequest, opts ...grpc.CallOption) (*WorkflowContractServiceDeleteResponse, error) {
	out := new(WorkflowContractServiceDeleteResponse)
	err := c.cc.Invoke(ctx, WorkflowContractService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowContractServiceServer is the server API for WorkflowContractService service.
// All implementations must embed UnimplementedWorkflowContractServiceServer
// for forward compatibility
type WorkflowContractServiceServer interface {
	List(context.Context, *WorkflowContractServiceListRequest) (*WorkflowContractServiceListResponse, error)
	Create(context.Context, *WorkflowContractServiceCreateRequest) (*WorkflowContractServiceCreateResponse, error)
	Update(context.Context, *WorkflowContractServiceUpdateRequest) (*WorkflowContractServiceUpdateResponse, error)
	Describe(context.Context, *WorkflowContractServiceDescribeRequest) (*WorkflowContractServiceDescribeResponse, error)
	Delete(context.Context, *WorkflowContractServiceDeleteRequest) (*WorkflowContractServiceDeleteResponse, error)
	mustEmbedUnimplementedWorkflowContractServiceServer()
}

// UnimplementedWorkflowContractServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkflowContractServiceServer struct {
}

func (UnimplementedWorkflowContractServiceServer) List(context.Context, *WorkflowContractServiceListRequest) (*WorkflowContractServiceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedWorkflowContractServiceServer) Create(context.Context, *WorkflowContractServiceCreateRequest) (*WorkflowContractServiceCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedWorkflowContractServiceServer) Update(context.Context, *WorkflowContractServiceUpdateRequest) (*WorkflowContractServiceUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedWorkflowContractServiceServer) Describe(context.Context, *WorkflowContractServiceDescribeRequest) (*WorkflowContractServiceDescribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Describe not implemented")
}
func (UnimplementedWorkflowContractServiceServer) Delete(context.Context, *WorkflowContractServiceDeleteRequest) (*WorkflowContractServiceDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedWorkflowContractServiceServer) mustEmbedUnimplementedWorkflowContractServiceServer() {
}

// UnsafeWorkflowContractServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkflowContractServiceServer will
// result in compilation errors.
type UnsafeWorkflowContractServiceServer interface {
	mustEmbedUnimplementedWorkflowContractServiceServer()
}

func RegisterWorkflowContractServiceServer(s grpc.ServiceRegistrar, srv WorkflowContractServiceServer) {
	s.RegisterService(&WorkflowContractService_ServiceDesc, srv)
}

func _WorkflowContractService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowContractServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowContractServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowContractService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowContractServiceServer).List(ctx, req.(*WorkflowContractServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowContractService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowContractServiceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowContractServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowContractService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowContractServiceServer).Create(ctx, req.(*WorkflowContractServiceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowContractService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowContractServiceUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowContractServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowContractService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowContractServiceServer).Update(ctx, req.(*WorkflowContractServiceUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowContractService_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowContractServiceDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowContractServiceServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowContractService_Describe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowContractServiceServer).Describe(ctx, req.(*WorkflowContractServiceDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowContractService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowContractServiceDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowContractServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowContractService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowContractServiceServer).Delete(ctx, req.(*WorkflowContractServiceDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkflowContractService_ServiceDesc is the grpc.ServiceDesc for WorkflowContractService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkflowContractService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controlplane.v1.WorkflowContractService",
	HandlerType: (*WorkflowContractServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _WorkflowContractService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _WorkflowContractService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _WorkflowContractService_Update_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _WorkflowContractService_Describe_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _WorkflowContractService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controlplane/v1/workflow_contract.proto",
}
