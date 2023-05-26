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
// source: controlplane/v1/integrations.proto

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
	IntegrationsService_AddDependencyTrack_FullMethodName = "/controlplane.v1.IntegrationsService/AddDependencyTrack"
	IntegrationsService_List_FullMethodName               = "/controlplane.v1.IntegrationsService/List"
	IntegrationsService_Delete_FullMethodName             = "/controlplane.v1.IntegrationsService/Delete"
	IntegrationsService_Attach_FullMethodName             = "/controlplane.v1.IntegrationsService/Attach"
	IntegrationsService_Detach_FullMethodName             = "/controlplane.v1.IntegrationsService/Detach"
	IntegrationsService_ListAttachments_FullMethodName    = "/controlplane.v1.IntegrationsService/ListAttachments"
)

// IntegrationsServiceClient is the client API for IntegrationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntegrationsServiceClient interface {
	// ORG related CRUD
	AddDependencyTrack(ctx context.Context, in *AddDependencyTrackRequest, opts ...grpc.CallOption) (*AddDependencyTrackResponse, error)
	List(ctx context.Context, in *IntegrationsServiceListRequest, opts ...grpc.CallOption) (*IntegrationsServiceListResponse, error)
	Delete(ctx context.Context, in *IntegrationsServiceDeleteRequest, opts ...grpc.CallOption) (*IntegrationsServiceDeleteResponse, error)
	// Workflow Related operations
	// Attach to a workflow
	Attach(ctx context.Context, in *IntegrationsServiceAttachRequest, opts ...grpc.CallOption) (*IntegrationsServiceAttachResponse, error)
	// Detach integration from a workflow
	Detach(ctx context.Context, in *IntegrationsServiceDetachRequest, opts ...grpc.CallOption) (*IntegrationsServiceDetachResponse, error)
	ListAttachments(ctx context.Context, in *ListAttachmentsRequest, opts ...grpc.CallOption) (*ListAttachmentsResponse, error)
}

type integrationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIntegrationsServiceClient(cc grpc.ClientConnInterface) IntegrationsServiceClient {
	return &integrationsServiceClient{cc}
}

func (c *integrationsServiceClient) AddDependencyTrack(ctx context.Context, in *AddDependencyTrackRequest, opts ...grpc.CallOption) (*AddDependencyTrackResponse, error) {
	out := new(AddDependencyTrackResponse)
	err := c.cc.Invoke(ctx, IntegrationsService_AddDependencyTrack_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationsServiceClient) List(ctx context.Context, in *IntegrationsServiceListRequest, opts ...grpc.CallOption) (*IntegrationsServiceListResponse, error) {
	out := new(IntegrationsServiceListResponse)
	err := c.cc.Invoke(ctx, IntegrationsService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationsServiceClient) Delete(ctx context.Context, in *IntegrationsServiceDeleteRequest, opts ...grpc.CallOption) (*IntegrationsServiceDeleteResponse, error) {
	out := new(IntegrationsServiceDeleteResponse)
	err := c.cc.Invoke(ctx, IntegrationsService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationsServiceClient) Attach(ctx context.Context, in *IntegrationsServiceAttachRequest, opts ...grpc.CallOption) (*IntegrationsServiceAttachResponse, error) {
	out := new(IntegrationsServiceAttachResponse)
	err := c.cc.Invoke(ctx, IntegrationsService_Attach_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationsServiceClient) Detach(ctx context.Context, in *IntegrationsServiceDetachRequest, opts ...grpc.CallOption) (*IntegrationsServiceDetachResponse, error) {
	out := new(IntegrationsServiceDetachResponse)
	err := c.cc.Invoke(ctx, IntegrationsService_Detach_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationsServiceClient) ListAttachments(ctx context.Context, in *ListAttachmentsRequest, opts ...grpc.CallOption) (*ListAttachmentsResponse, error) {
	out := new(ListAttachmentsResponse)
	err := c.cc.Invoke(ctx, IntegrationsService_ListAttachments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntegrationsServiceServer is the server API for IntegrationsService service.
// All implementations must embed UnimplementedIntegrationsServiceServer
// for forward compatibility
type IntegrationsServiceServer interface {
	// ORG related CRUD
	AddDependencyTrack(context.Context, *AddDependencyTrackRequest) (*AddDependencyTrackResponse, error)
	List(context.Context, *IntegrationsServiceListRequest) (*IntegrationsServiceListResponse, error)
	Delete(context.Context, *IntegrationsServiceDeleteRequest) (*IntegrationsServiceDeleteResponse, error)
	// Workflow Related operations
	// Attach to a workflow
	Attach(context.Context, *IntegrationsServiceAttachRequest) (*IntegrationsServiceAttachResponse, error)
	// Detach integration from a workflow
	Detach(context.Context, *IntegrationsServiceDetachRequest) (*IntegrationsServiceDetachResponse, error)
	ListAttachments(context.Context, *ListAttachmentsRequest) (*ListAttachmentsResponse, error)
	mustEmbedUnimplementedIntegrationsServiceServer()
}

// UnimplementedIntegrationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIntegrationsServiceServer struct {
}

func (UnimplementedIntegrationsServiceServer) AddDependencyTrack(context.Context, *AddDependencyTrackRequest) (*AddDependencyTrackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDependencyTrack not implemented")
}
func (UnimplementedIntegrationsServiceServer) List(context.Context, *IntegrationsServiceListRequest) (*IntegrationsServiceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedIntegrationsServiceServer) Delete(context.Context, *IntegrationsServiceDeleteRequest) (*IntegrationsServiceDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedIntegrationsServiceServer) Attach(context.Context, *IntegrationsServiceAttachRequest) (*IntegrationsServiceAttachResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Attach not implemented")
}
func (UnimplementedIntegrationsServiceServer) Detach(context.Context, *IntegrationsServiceDetachRequest) (*IntegrationsServiceDetachResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detach not implemented")
}
func (UnimplementedIntegrationsServiceServer) ListAttachments(context.Context, *ListAttachmentsRequest) (*ListAttachmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAttachments not implemented")
}
func (UnimplementedIntegrationsServiceServer) mustEmbedUnimplementedIntegrationsServiceServer() {}

// UnsafeIntegrationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntegrationsServiceServer will
// result in compilation errors.
type UnsafeIntegrationsServiceServer interface {
	mustEmbedUnimplementedIntegrationsServiceServer()
}

func RegisterIntegrationsServiceServer(s grpc.ServiceRegistrar, srv IntegrationsServiceServer) {
	s.RegisterService(&IntegrationsService_ServiceDesc, srv)
}

func _IntegrationsService_AddDependencyTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDependencyTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationsServiceServer).AddDependencyTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationsService_AddDependencyTrack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationsServiceServer).AddDependencyTrack(ctx, req.(*AddDependencyTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationsService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntegrationsServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationsServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationsService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationsServiceServer).List(ctx, req.(*IntegrationsServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntegrationsServiceDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationsService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationsServiceServer).Delete(ctx, req.(*IntegrationsServiceDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationsService_Attach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntegrationsServiceAttachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationsServiceServer).Attach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationsService_Attach_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationsServiceServer).Attach(ctx, req.(*IntegrationsServiceAttachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationsService_Detach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntegrationsServiceDetachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationsServiceServer).Detach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationsService_Detach_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationsServiceServer).Detach(ctx, req.(*IntegrationsServiceDetachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationsService_ListAttachments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAttachmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationsServiceServer).ListAttachments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationsService_ListAttachments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationsServiceServer).ListAttachments(ctx, req.(*ListAttachmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IntegrationsService_ServiceDesc is the grpc.ServiceDesc for IntegrationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IntegrationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controlplane.v1.IntegrationsService",
	HandlerType: (*IntegrationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDependencyTrack",
			Handler:    _IntegrationsService_AddDependencyTrack_Handler,
		},
		{
			MethodName: "List",
			Handler:    _IntegrationsService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _IntegrationsService_Delete_Handler,
		},
		{
			MethodName: "Attach",
			Handler:    _IntegrationsService_Attach_Handler,
		},
		{
			MethodName: "Detach",
			Handler:    _IntegrationsService_Detach_Handler,
		},
		{
			MethodName: "ListAttachments",
			Handler:    _IntegrationsService_ListAttachments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controlplane/v1/integrations.proto",
}
