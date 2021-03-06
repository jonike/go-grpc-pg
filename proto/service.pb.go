// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	Row
	QueryRequest
	QueryResponse
	ExecRequest
	ExecResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Row struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *Row) Reset()                    { *m = Row{} }
func (m *Row) String() string            { return proto1.CompactTextString(m) }
func (*Row) ProtoMessage()               {}
func (*Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Row) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// The query statement
type QueryRequest struct {
	Stmt string `protobuf:"bytes,1,opt,name=stmt" json:"stmt,omitempty"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto1.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *QueryRequest) GetStmt() string {
	if m != nil {
		return m.Stmt
	}
	return ""
}

// The query response
type QueryResponse struct {
	Columns []string `protobuf:"bytes,1,rep,name=columns" json:"columns,omitempty"`
	Rows    []*Row   `protobuf:"bytes,3,rep,name=rows" json:"rows,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto1.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QueryResponse) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *QueryResponse) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

// The exec statement
type ExecRequest struct {
	Stmt string `protobuf:"bytes,1,opt,name=stmt" json:"stmt,omitempty"`
}

func (m *ExecRequest) Reset()                    { *m = ExecRequest{} }
func (m *ExecRequest) String() string            { return proto1.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()               {}
func (*ExecRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ExecRequest) GetStmt() string {
	if m != nil {
		return m.Stmt
	}
	return ""
}

// The exec response
type ExecResponse struct {
	LastInsertId int64 `protobuf:"varint,1,opt,name=lastInsertId" json:"lastInsertId,omitempty"`
	RowsAffected int64 `protobuf:"varint,2,opt,name=rowsAffected" json:"rowsAffected,omitempty"`
}

func (m *ExecResponse) Reset()                    { *m = ExecResponse{} }
func (m *ExecResponse) String() string            { return proto1.CompactTextString(m) }
func (*ExecResponse) ProtoMessage()               {}
func (*ExecResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ExecResponse) GetLastInsertId() int64 {
	if m != nil {
		return m.LastInsertId
	}
	return 0
}

func (m *ExecResponse) GetRowsAffected() int64 {
	if m != nil {
		return m.RowsAffected
	}
	return 0
}

func init() {
	proto1.RegisterType((*Row)(nil), "proto.Row")
	proto1.RegisterType((*QueryRequest)(nil), "proto.QueryRequest")
	proto1.RegisterType((*QueryResponse)(nil), "proto.QueryResponse")
	proto1.RegisterType((*ExecRequest)(nil), "proto.ExecRequest")
	proto1.RegisterType((*ExecResponse)(nil), "proto.ExecResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DBProvider service

type DBProviderClient interface {
	// Query executes a statement that reads data.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// Exec executes a statement that writes data.
	Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error)
}

type dBProviderClient struct {
	cc *grpc.ClientConn
}

func NewDBProviderClient(cc *grpc.ClientConn) DBProviderClient {
	return &dBProviderClient{cc}
}

func (c *dBProviderClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := grpc.Invoke(ctx, "/proto.DBProvider/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBProviderClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error) {
	out := new(ExecResponse)
	err := grpc.Invoke(ctx, "/proto.DBProvider/Exec", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DBProvider service

type DBProviderServer interface {
	// Query executes a statement that reads data.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	// Exec executes a statement that writes data.
	Exec(context.Context, *ExecRequest) (*ExecResponse, error)
}

func RegisterDBProviderServer(s *grpc.Server, srv DBProviderServer) {
	s.RegisterService(&_DBProvider_serviceDesc, srv)
}

func _DBProvider_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBProviderServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DBProvider/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBProviderServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBProvider_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBProviderServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DBProvider/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBProviderServer).Exec(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DBProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DBProvider",
	HandlerType: (*DBProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _DBProvider_Query_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _DBProvider_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto1.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x09, 0x4e, 0x8b, 0x7a, 0x4d, 0x97, 0x2b, 0x42, 0x56, 0x25, 0x50, 0xf0, 0x94, 0xa9,
	0x12, 0x85, 0x17, 0x00, 0xc1, 0x90, 0x0d, 0x3c, 0xb0, 0x97, 0xe4, 0x2a, 0x55, 0x4a, 0xe3, 0xe2,
	0x3f, 0x09, 0xbc, 0x3d, 0x8a, 0xe3, 0x48, 0xc9, 0xc2, 0x94, 0xdc, 0xe7, 0xcf, 0x77, 0xbf, 0x33,
	0xac, 0x0c, 0xe9, 0xe6, 0x58, 0xd0, 0xf6, 0xac, 0x95, 0x55, 0x38, 0xf3, 0x1f, 0x71, 0x0b, 0x4c,
	0xaa, 0x16, 0x6f, 0x60, 0xde, 0xec, 0x2b, 0x47, 0x86, 0x47, 0x29, 0xcb, 0x16, 0x32, 0x54, 0x42,
	0x40, 0xf2, 0xe1, 0x48, 0xff, 0x4a, 0xfa, 0x76, 0x64, 0x2c, 0x22, 0xc4, 0xc6, 0x9e, 0x2c, 0x8f,
	0xd2, 0x28, 0x5b, 0x48, 0xff, 0x2f, 0x72, 0x58, 0x05, 0xc7, 0x9c, 0x55, 0x6d, 0x08, 0x39, 0x5c,
	0x15, 0xaa, 0x72, 0xa7, 0x7a, 0xe8, 0x36, 0x94, 0x78, 0x07, 0xb1, 0x56, 0xad, 0xe1, 0x2c, 0x65,
	0xd9, 0x72, 0x07, 0x7d, 0x94, 0xad, 0x54, 0xad, 0xf4, 0x5c, 0xdc, 0xc3, 0xf2, 0xed, 0x87, 0x8a,
	0xff, 0xa6, 0x7d, 0x42, 0xd2, 0x2b, 0x61, 0x98, 0x80, 0xa4, 0xda, 0x1b, 0x9b, 0xd7, 0x86, 0xb4,
	0xcd, 0x4b, 0xef, 0x32, 0x39, 0x61, 0x9d, 0xd3, 0xb5, 0x7f, 0x3e, 0x1c, 0xa8, 0xb0, 0x54, 0xf2,
	0xcb, 0xde, 0x19, 0xb3, 0x9d, 0x03, 0x78, 0x7d, 0x79, 0xd7, 0xaa, 0x39, 0x96, 0xa4, 0xf1, 0x09,
	0x66, 0x7e, 0x27, 0x5c, 0x87, 0x8c, 0xe3, 0x57, 0xd8, 0x5c, 0x4f, 0x61, 0x9f, 0x44, 0x5c, 0xe0,
	0x03, 0xc4, 0x5d, 0x36, 0xc4, 0x70, 0x3e, 0xda, 0x65, 0xb3, 0x9e, 0xb0, 0xe1, 0xca, 0xd7, 0xdc,
	0xd3, 0xc7, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xb6, 0x66, 0x68, 0x9e, 0x01, 0x00, 0x00,
}
