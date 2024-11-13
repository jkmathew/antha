// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/jkmathew/antha/driver/antha_driver_v1/driver.proto

package antha_driver_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TypeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeRequest) Reset()         { *m = TypeRequest{} }
func (m *TypeRequest) String() string { return proto.CompactTextString(m) }
func (*TypeRequest) ProtoMessage()    {}
func (*TypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_a025585186e2deaa, []int{0}
}
func (m *TypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeRequest.Unmarshal(m, b)
}
func (m *TypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeRequest.Marshal(b, m, deterministic)
}
func (dst *TypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeRequest.Merge(dst, src)
}
func (m *TypeRequest) XXX_Size() int {
	return xxx_messageInfo_TypeRequest.Size(m)
}
func (m *TypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TypeRequest proto.InternalMessageInfo

type TypeReply struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Subtypes             []string `protobuf:"bytes,2,rep,name=subtypes,proto3" json:"subtypes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeReply) Reset()         { *m = TypeReply{} }
func (m *TypeReply) String() string { return proto.CompactTextString(m) }
func (*TypeReply) ProtoMessage()    {}
func (*TypeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_a025585186e2deaa, []int{1}
}
func (m *TypeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeReply.Unmarshal(m, b)
}
func (m *TypeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeReply.Marshal(b, m, deterministic)
}
func (dst *TypeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeReply.Merge(dst, src)
}
func (m *TypeReply) XXX_Size() int {
	return xxx_messageInfo_TypeReply.Size(m)
}
func (m *TypeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeReply.DiscardUnknown(m)
}

var xxx_messageInfo_TypeReply proto.InternalMessageInfo

func (m *TypeReply) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TypeReply) GetSubtypes() []string {
	if m != nil {
		return m.Subtypes
	}
	return nil
}

type HttpHeader struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpHeader) Reset()         { *m = HttpHeader{} }
func (m *HttpHeader) String() string { return proto.CompactTextString(m) }
func (*HttpHeader) ProtoMessage()    {}
func (*HttpHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_a025585186e2deaa, []int{2}
}
func (m *HttpHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpHeader.Unmarshal(m, b)
}
func (m *HttpHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpHeader.Marshal(b, m, deterministic)
}
func (dst *HttpHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpHeader.Merge(dst, src)
}
func (m *HttpHeader) XXX_Size() int {
	return xxx_messageInfo_HttpHeader.Size(m)
}
func (m *HttpHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpHeader.DiscardUnknown(m)
}

var xxx_messageInfo_HttpHeader proto.InternalMessageInfo

func (m *HttpHeader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HttpHeader) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Remote Http call
type HttpCall struct {
	Url                  string        `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Method               string        `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Body                 []byte        `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Headers              []*HttpHeader `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HttpCall) Reset()         { *m = HttpCall{} }
func (m *HttpCall) String() string { return proto.CompactTextString(m) }
func (*HttpCall) ProtoMessage()    {}
func (*HttpCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_a025585186e2deaa, []int{3}
}
func (m *HttpCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpCall.Unmarshal(m, b)
}
func (m *HttpCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpCall.Marshal(b, m, deterministic)
}
func (dst *HttpCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpCall.Merge(dst, src)
}
func (m *HttpCall) XXX_Size() int {
	return xxx_messageInfo_HttpCall.Size(m)
}
func (m *HttpCall) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpCall.DiscardUnknown(m)
}

var xxx_messageInfo_HttpCall proto.InternalMessageInfo

func (m *HttpCall) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *HttpCall) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HttpCall) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *HttpCall) GetHeaders() []*HttpHeader {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*TypeRequest)(nil), "antha.driver.v1.TypeRequest")
	proto.RegisterType((*TypeReply)(nil), "antha.driver.v1.TypeReply")
	proto.RegisterType((*HttpHeader)(nil), "antha.driver.v1.HttpHeader")
	proto.RegisterType((*HttpCall)(nil), "antha.driver.v1.HttpCall")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DriverClient is the client API for Driver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DriverClient interface {
	DriverType(ctx context.Context, in *TypeRequest, opts ...grpc.CallOption) (*TypeReply, error)
}

type driverClient struct {
	cc *grpc.ClientConn
}

func NewDriverClient(cc *grpc.ClientConn) DriverClient {
	return &driverClient{cc}
}

func (c *driverClient) DriverType(ctx context.Context, in *TypeRequest, opts ...grpc.CallOption) (*TypeReply, error) {
	out := new(TypeReply)
	err := c.cc.Invoke(ctx, "/antha.driver.v1.Driver/DriverType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServer is the server API for Driver service.
type DriverServer interface {
	DriverType(context.Context, *TypeRequest) (*TypeReply, error)
}

func RegisterDriverServer(s *grpc.Server, srv DriverServer) {
	s.RegisterService(&_Driver_serviceDesc, srv)
}

func _Driver_DriverType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).DriverType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/antha.driver.v1.Driver/DriverType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).DriverType(ctx, req.(*TypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Driver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "antha.driver.v1.Driver",
	HandlerType: (*DriverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DriverType",
			Handler:    _Driver_DriverType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/jkmathew/antha/driver/antha_driver_v1/driver.proto",
}

func init() {
	proto.RegisterFile("github.com/jkmathew/antha/driver/antha_driver_v1/driver.proto", fileDescriptor_driver_a025585186e2deaa)
}

var fileDescriptor_driver_a025585186e2deaa = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcb, 0x4b, 0xc3, 0x40,
	0x10, 0x87, 0x49, 0x53, 0x63, 0x33, 0x55, 0x94, 0x41, 0x24, 0x44, 0x0f, 0x21, 0xa7, 0x5c, 0x4c,
	0x69, 0x45, 0x2f, 0x1e, 0x3c, 0x28, 0xd2, 0xa3, 0x04, 0xef, 0x65, 0x63, 0x96, 0xa6, 0xb0, 0x49,
	0xd6, 0x64, 0x37, 0xb0, 0x27, 0xff, 0x75, 0xd9, 0x87, 0x55, 0x94, 0xde, 0xbe, 0xdf, 0xcc, 0x7c,
	0xfb, 0x18, 0x78, 0xdc, 0xee, 0x44, 0x2d, 0xcb, 0xfc, 0xbd, 0x6b, 0x16, 0xa4, 0x15, 0x35, 0xb9,
	0x61, 0xa4, 0xdd, 0x5a, 0x5c, 0x54, 0xfd, 0x6e, 0xa4, 0xbd, 0x0d, 0x1b, 0x1b, 0x36, 0xe3, 0xd2,
	0x95, 0x73, 0xde, 0x77, 0xa2, 0xc3, 0x33, 0xd3, 0xcd, 0x5d, 0x6d, 0x5c, 0xa6, 0xa7, 0x30, 0x7f,
	0x53, 0x9c, 0x16, 0xf4, 0x43, 0xd2, 0x41, 0xa4, 0x0f, 0x10, 0xda, 0xc8, 0x99, 0x42, 0x84, 0xa9,
	0x50, 0x9c, 0x46, 0x5e, 0xe2, 0x65, 0x61, 0x61, 0x18, 0x63, 0x98, 0x0d, 0xb2, 0xd4, 0x38, 0x44,
	0x93, 0xc4, 0xcf, 0xc2, 0x62, 0x9f, 0xd3, 0x7b, 0x80, 0xb5, 0x10, 0x7c, 0x4d, 0x49, 0x45, 0x7b,
	0x6d, 0xb7, 0xa4, 0xd9, 0xdb, 0x9a, 0xf1, 0x02, 0x8e, 0x46, 0xc2, 0x24, 0x8d, 0x26, 0xa6, 0x68,
	0x43, 0xfa, 0x09, 0x33, 0xed, 0x3d, 0x11, 0xc6, 0xf0, 0x1c, 0x7c, 0xd9, 0x33, 0x27, 0x69, 0xc4,
	0x4b, 0x08, 0x1a, 0x2a, 0xea, 0xae, 0x72, 0x92, 0x4b, 0xfa, 0xfc, 0xb2, 0xab, 0x54, 0xe4, 0x27,
	0x5e, 0x76, 0x52, 0x18, 0xc6, 0x3b, 0x38, 0xae, 0xcd, 0xed, 0x43, 0x34, 0x4d, 0xfc, 0x6c, 0xbe,
	0xba, 0xca, 0xff, 0x7c, 0x38, 0xff, 0x79, 0x61, 0xf1, 0x3d, 0xbb, 0x7a, 0x85, 0xe0, 0xd9, 0x0c,
	0xe0, 0x0b, 0x80, 0x25, 0xbd, 0x05, 0xbc, 0xfe, 0x67, 0xff, 0xda, 0x55, 0x1c, 0x1f, 0xe8, 0x72,
	0xa6, 0xca, 0xc0, 0xac, 0xfb, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x88, 0x8d, 0x8d, 0x66, 0xb1,
	0x01, 0x00, 0x00,
}
