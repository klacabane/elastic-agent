// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: control_v2.proto

package cproto

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

// ElasticAgentControlClient is the client API for ElasticAgentControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElasticAgentControlClient interface {
	// Fetches the currently running version of the Elastic Agent.
	Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	// Fetches the currently states of the Elastic Agent.
	State(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StateResponse, error)
	// Restart restarts the current running Elastic Agent.
	Restart(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RestartResponse, error)
	// Upgrade starts the upgrade process of Elastic Agent.
	Upgrade(ctx context.Context, in *UpgradeRequest, opts ...grpc.CallOption) (*UpgradeResponse, error)
	// Gather diagnostic information for the running Elastic Agent.
	DiagnosticAgent(ctx context.Context, in *DiagnosticAgentRequest, opts ...grpc.CallOption) (*DiagnosticAgentResponse, error)
	// Gather diagnostic information for the running units.
	DiagnosticUnits(ctx context.Context, in *DiagnosticUnitsRequest, opts ...grpc.CallOption) (ElasticAgentControl_DiagnosticUnitsClient, error)
}

type elasticAgentControlClient struct {
	cc grpc.ClientConnInterface
}

func NewElasticAgentControlClient(cc grpc.ClientConnInterface) ElasticAgentControlClient {
	return &elasticAgentControlClient{cc}
}

func (c *elasticAgentControlClient) Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/cproto.ElasticAgentControl/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticAgentControlClient) State(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StateResponse, error) {
	out := new(StateResponse)
	err := c.cc.Invoke(ctx, "/cproto.ElasticAgentControl/State", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticAgentControlClient) Restart(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RestartResponse, error) {
	out := new(RestartResponse)
	err := c.cc.Invoke(ctx, "/cproto.ElasticAgentControl/Restart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticAgentControlClient) Upgrade(ctx context.Context, in *UpgradeRequest, opts ...grpc.CallOption) (*UpgradeResponse, error) {
	out := new(UpgradeResponse)
	err := c.cc.Invoke(ctx, "/cproto.ElasticAgentControl/Upgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticAgentControlClient) DiagnosticAgent(ctx context.Context, in *DiagnosticAgentRequest, opts ...grpc.CallOption) (*DiagnosticAgentResponse, error) {
	out := new(DiagnosticAgentResponse)
	err := c.cc.Invoke(ctx, "/cproto.ElasticAgentControl/DiagnosticAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticAgentControlClient) DiagnosticUnits(ctx context.Context, in *DiagnosticUnitsRequest, opts ...grpc.CallOption) (ElasticAgentControl_DiagnosticUnitsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ElasticAgentControl_ServiceDesc.Streams[0], "/cproto.ElasticAgentControl/DiagnosticUnits", opts...)
	if err != nil {
		return nil, err
	}
	x := &elasticAgentControlDiagnosticUnitsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ElasticAgentControl_DiagnosticUnitsClient interface {
	Recv() (*DiagnosticUnitResponse, error)
	grpc.ClientStream
}

type elasticAgentControlDiagnosticUnitsClient struct {
	grpc.ClientStream
}

func (x *elasticAgentControlDiagnosticUnitsClient) Recv() (*DiagnosticUnitResponse, error) {
	m := new(DiagnosticUnitResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ElasticAgentControlServer is the server API for ElasticAgentControl service.
// All implementations must embed UnimplementedElasticAgentControlServer
// for forward compatibility
type ElasticAgentControlServer interface {
	// Fetches the currently running version of the Elastic Agent.
	Version(context.Context, *Empty) (*VersionResponse, error)
	// Fetches the currently states of the Elastic Agent.
	State(context.Context, *Empty) (*StateResponse, error)
	// Restart restarts the current running Elastic Agent.
	Restart(context.Context, *Empty) (*RestartResponse, error)
	// Upgrade starts the upgrade process of Elastic Agent.
	Upgrade(context.Context, *UpgradeRequest) (*UpgradeResponse, error)
	// Gather diagnostic information for the running Elastic Agent.
	DiagnosticAgent(context.Context, *DiagnosticAgentRequest) (*DiagnosticAgentResponse, error)
	// Gather diagnostic information for the running units.
	DiagnosticUnits(*DiagnosticUnitsRequest, ElasticAgentControl_DiagnosticUnitsServer) error
	mustEmbedUnimplementedElasticAgentControlServer()
}

// UnimplementedElasticAgentControlServer must be embedded to have forward compatible implementations.
type UnimplementedElasticAgentControlServer struct {
}

func (UnimplementedElasticAgentControlServer) Version(context.Context, *Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedElasticAgentControlServer) State(context.Context, *Empty) (*StateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method State not implemented")
}
func (UnimplementedElasticAgentControlServer) Restart(context.Context, *Empty) (*RestartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedElasticAgentControlServer) Upgrade(context.Context, *UpgradeRequest) (*UpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upgrade not implemented")
}
func (UnimplementedElasticAgentControlServer) DiagnosticAgent(context.Context, *DiagnosticAgentRequest) (*DiagnosticAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiagnosticAgent not implemented")
}
func (UnimplementedElasticAgentControlServer) DiagnosticUnits(*DiagnosticUnitsRequest, ElasticAgentControl_DiagnosticUnitsServer) error {
	return status.Errorf(codes.Unimplemented, "method DiagnosticUnits not implemented")
}
func (UnimplementedElasticAgentControlServer) mustEmbedUnimplementedElasticAgentControlServer() {}

// UnsafeElasticAgentControlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElasticAgentControlServer will
// result in compilation errors.
type UnsafeElasticAgentControlServer interface {
	mustEmbedUnimplementedElasticAgentControlServer()
}

func RegisterElasticAgentControlServer(s grpc.ServiceRegistrar, srv ElasticAgentControlServer) {
	s.RegisterService(&ElasticAgentControl_ServiceDesc, srv)
}

func _ElasticAgentControl_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticAgentControlServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cproto.ElasticAgentControl/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticAgentControlServer).Version(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticAgentControl_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticAgentControlServer).State(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cproto.ElasticAgentControl/State",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticAgentControlServer).State(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticAgentControl_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticAgentControlServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cproto.ElasticAgentControl/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticAgentControlServer).Restart(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticAgentControl_Upgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticAgentControlServer).Upgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cproto.ElasticAgentControl/Upgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticAgentControlServer).Upgrade(ctx, req.(*UpgradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticAgentControl_DiagnosticAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiagnosticAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticAgentControlServer).DiagnosticAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cproto.ElasticAgentControl/DiagnosticAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticAgentControlServer).DiagnosticAgent(ctx, req.(*DiagnosticAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticAgentControl_DiagnosticUnits_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DiagnosticUnitsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ElasticAgentControlServer).DiagnosticUnits(m, &elasticAgentControlDiagnosticUnitsServer{stream})
}

type ElasticAgentControl_DiagnosticUnitsServer interface {
	Send(*DiagnosticUnitResponse) error
	grpc.ServerStream
}

type elasticAgentControlDiagnosticUnitsServer struct {
	grpc.ServerStream
}

func (x *elasticAgentControlDiagnosticUnitsServer) Send(m *DiagnosticUnitResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ElasticAgentControl_ServiceDesc is the grpc.ServiceDesc for ElasticAgentControl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ElasticAgentControl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cproto.ElasticAgentControl",
	HandlerType: (*ElasticAgentControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _ElasticAgentControl_Version_Handler,
		},
		{
			MethodName: "State",
			Handler:    _ElasticAgentControl_State_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _ElasticAgentControl_Restart_Handler,
		},
		{
			MethodName: "Upgrade",
			Handler:    _ElasticAgentControl_Upgrade_Handler,
		},
		{
			MethodName: "DiagnosticAgent",
			Handler:    _ElasticAgentControl_DiagnosticAgent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DiagnosticUnits",
			Handler:       _ElasticAgentControl_DiagnosticUnits_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "control_v2.proto",
}
