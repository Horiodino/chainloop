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
// source: controlplane/v1/robot_accounts.proto

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
	RobotAccountService_Create_FullMethodName = "/controlplane.v1.RobotAccountService/Create"
	RobotAccountService_List_FullMethodName   = "/controlplane.v1.RobotAccountService/List"
	RobotAccountService_Revoke_FullMethodName = "/controlplane.v1.RobotAccountService/Revoke"
)

// RobotAccountServiceClient is the client API for RobotAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RobotAccountServiceClient interface {
	Create(ctx context.Context, in *RobotAccountServiceCreateRequest, opts ...grpc.CallOption) (*RobotAccountServiceCreateResponse, error)
	List(ctx context.Context, in *RobotAccountServiceListRequest, opts ...grpc.CallOption) (*RobotAccountServiceListResponse, error)
	Revoke(ctx context.Context, in *RobotAccountServiceRevokeRequest, opts ...grpc.CallOption) (*RobotAccountServiceRevokeResponse, error)
}

type robotAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRobotAccountServiceClient(cc grpc.ClientConnInterface) RobotAccountServiceClient {
	return &robotAccountServiceClient{cc}
}

func (c *robotAccountServiceClient) Create(ctx context.Context, in *RobotAccountServiceCreateRequest, opts ...grpc.CallOption) (*RobotAccountServiceCreateResponse, error) {
	out := new(RobotAccountServiceCreateResponse)
	err := c.cc.Invoke(ctx, RobotAccountService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotAccountServiceClient) List(ctx context.Context, in *RobotAccountServiceListRequest, opts ...grpc.CallOption) (*RobotAccountServiceListResponse, error) {
	out := new(RobotAccountServiceListResponse)
	err := c.cc.Invoke(ctx, RobotAccountService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotAccountServiceClient) Revoke(ctx context.Context, in *RobotAccountServiceRevokeRequest, opts ...grpc.CallOption) (*RobotAccountServiceRevokeResponse, error) {
	out := new(RobotAccountServiceRevokeResponse)
	err := c.cc.Invoke(ctx, RobotAccountService_Revoke_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotAccountServiceServer is the server API for RobotAccountService service.
// All implementations must embed UnimplementedRobotAccountServiceServer
// for forward compatibility
type RobotAccountServiceServer interface {
	Create(context.Context, *RobotAccountServiceCreateRequest) (*RobotAccountServiceCreateResponse, error)
	List(context.Context, *RobotAccountServiceListRequest) (*RobotAccountServiceListResponse, error)
	Revoke(context.Context, *RobotAccountServiceRevokeRequest) (*RobotAccountServiceRevokeResponse, error)
	mustEmbedUnimplementedRobotAccountServiceServer()
}

// UnimplementedRobotAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRobotAccountServiceServer struct {
}

func (UnimplementedRobotAccountServiceServer) Create(context.Context, *RobotAccountServiceCreateRequest) (*RobotAccountServiceCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRobotAccountServiceServer) List(context.Context, *RobotAccountServiceListRequest) (*RobotAccountServiceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRobotAccountServiceServer) Revoke(context.Context, *RobotAccountServiceRevokeRequest) (*RobotAccountServiceRevokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Revoke not implemented")
}
func (UnimplementedRobotAccountServiceServer) mustEmbedUnimplementedRobotAccountServiceServer() {}

// UnsafeRobotAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RobotAccountServiceServer will
// result in compilation errors.
type UnsafeRobotAccountServiceServer interface {
	mustEmbedUnimplementedRobotAccountServiceServer()
}

func RegisterRobotAccountServiceServer(s grpc.ServiceRegistrar, srv RobotAccountServiceServer) {
	s.RegisterService(&RobotAccountService_ServiceDesc, srv)
}

func _RobotAccountService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RobotAccountServiceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotAccountServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RobotAccountService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotAccountServiceServer).Create(ctx, req.(*RobotAccountServiceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotAccountService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RobotAccountServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotAccountServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RobotAccountService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotAccountServiceServer).List(ctx, req.(*RobotAccountServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotAccountService_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RobotAccountServiceRevokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotAccountServiceServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RobotAccountService_Revoke_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotAccountServiceServer).Revoke(ctx, req.(*RobotAccountServiceRevokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RobotAccountService_ServiceDesc is the grpc.ServiceDesc for RobotAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RobotAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controlplane.v1.RobotAccountService",
	HandlerType: (*RobotAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RobotAccountService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RobotAccountService_List_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _RobotAccountService_Revoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controlplane/v1/robot_accounts.proto",
}
