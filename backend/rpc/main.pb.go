// Code generated by protoc-gen-go. DO NOT EDIT.
// source: main.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	main.proto

It has these top-level messages:
	TerminalStateRequest
	TerminalStateResponse
	TerminalBuyRequest
	TerminalBuyResponse
	TerminalScanRequest
	TerminalScanResponse
	AbortRequest
	AbortResponse
*/
package rpc

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

type TerminalStateRequest struct {
	TerminalID string `protobuf:"bytes,1,opt,name=TerminalID" json:"TerminalID,omitempty"`
}

func (m *TerminalStateRequest) Reset()                    { *m = TerminalStateRequest{} }
func (m *TerminalStateRequest) String() string            { return proto.CompactTextString(m) }
func (*TerminalStateRequest) ProtoMessage()               {}
func (*TerminalStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TerminalStateRequest) GetTerminalID() string {
	if m != nil {
		return m.TerminalID
	}
	return ""
}

type TerminalStateResponse struct {
	Accounts  []*TerminalStateResponse_Account `protobuf:"bytes,1,rep,name=Accounts" json:"Accounts,omitempty"`
	OrderList []*TerminalStateResponse_Order   `protobuf:"bytes,2,rep,name=OrderList" json:"OrderList,omitempty"`
}

func (m *TerminalStateResponse) Reset()                    { *m = TerminalStateResponse{} }
func (m *TerminalStateResponse) String() string            { return proto.CompactTextString(m) }
func (*TerminalStateResponse) ProtoMessage()               {}
func (*TerminalStateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TerminalStateResponse) GetAccounts() []*TerminalStateResponse_Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *TerminalStateResponse) GetOrderList() []*TerminalStateResponse_Order {
	if m != nil {
		return m.OrderList
	}
	return nil
}

type TerminalStateResponse_Account struct {
	ID          string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=DisplayName" json:"DisplayName,omitempty"`
	Balance     int32  `protobuf:"zigzag32,3,opt,name=Balance" json:"Balance,omitempty"`
}

func (m *TerminalStateResponse_Account) Reset()         { *m = TerminalStateResponse_Account{} }
func (m *TerminalStateResponse_Account) String() string { return proto.CompactTextString(m) }
func (*TerminalStateResponse_Account) ProtoMessage()    {}
func (*TerminalStateResponse_Account) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

func (m *TerminalStateResponse_Account) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *TerminalStateResponse_Account) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *TerminalStateResponse_Account) GetBalance() int32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type TerminalStateResponse_Order struct {
	DisplayName string `protobuf:"bytes,1,opt,name=DisplayName" json:"DisplayName,omitempty"`
	UnitPrice   int32  `protobuf:"zigzag32,2,opt,name=UnitPrice" json:"UnitPrice,omitempty"`
	Amount      uint32 `protobuf:"varint,3,opt,name=Amount" json:"Amount,omitempty"`
}

func (m *TerminalStateResponse_Order) Reset()                    { *m = TerminalStateResponse_Order{} }
func (m *TerminalStateResponse_Order) String() string            { return proto.CompactTextString(m) }
func (*TerminalStateResponse_Order) ProtoMessage()               {}
func (*TerminalStateResponse_Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

func (m *TerminalStateResponse_Order) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *TerminalStateResponse_Order) GetUnitPrice() int32 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *TerminalStateResponse_Order) GetAmount() uint32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type TerminalBuyRequest struct {
	TerminalID string `protobuf:"bytes,1,opt,name=TerminalID" json:"TerminalID,omitempty"`
	AccountID  string `protobuf:"bytes,2,opt,name=AccountID" json:"AccountID,omitempty"`
}

func (m *TerminalBuyRequest) Reset()                    { *m = TerminalBuyRequest{} }
func (m *TerminalBuyRequest) String() string            { return proto.CompactTextString(m) }
func (*TerminalBuyRequest) ProtoMessage()               {}
func (*TerminalBuyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TerminalBuyRequest) GetTerminalID() string {
	if m != nil {
		return m.TerminalID
	}
	return ""
}

func (m *TerminalBuyRequest) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

type TerminalBuyResponse struct {
	Error string `protobuf:"bytes,1,opt,name=Error" json:"Error,omitempty"`
}

func (m *TerminalBuyResponse) Reset()                    { *m = TerminalBuyResponse{} }
func (m *TerminalBuyResponse) String() string            { return proto.CompactTextString(m) }
func (*TerminalBuyResponse) ProtoMessage()               {}
func (*TerminalBuyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TerminalBuyResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type TerminalScanRequest struct {
	TerminalID string `protobuf:"bytes,1,opt,name=TerminalID" json:"TerminalID,omitempty"`
	ProductID  string `protobuf:"bytes,2,opt,name=ProductID" json:"ProductID,omitempty"`
}

func (m *TerminalScanRequest) Reset()                    { *m = TerminalScanRequest{} }
func (m *TerminalScanRequest) String() string            { return proto.CompactTextString(m) }
func (*TerminalScanRequest) ProtoMessage()               {}
func (*TerminalScanRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TerminalScanRequest) GetTerminalID() string {
	if m != nil {
		return m.TerminalID
	}
	return ""
}

func (m *TerminalScanRequest) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

type TerminalScanResponse struct {
	Error string `protobuf:"bytes,1,opt,name=Error" json:"Error,omitempty"`
}

func (m *TerminalScanResponse) Reset()                    { *m = TerminalScanResponse{} }
func (m *TerminalScanResponse) String() string            { return proto.CompactTextString(m) }
func (*TerminalScanResponse) ProtoMessage()               {}
func (*TerminalScanResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TerminalScanResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type AbortRequest struct {
	TerminalID string `protobuf:"bytes,1,opt,name=TerminalID" json:"TerminalID,omitempty"`
}

func (m *AbortRequest) Reset()                    { *m = AbortRequest{} }
func (m *AbortRequest) String() string            { return proto.CompactTextString(m) }
func (*AbortRequest) ProtoMessage()               {}
func (*AbortRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AbortRequest) GetTerminalID() string {
	if m != nil {
		return m.TerminalID
	}
	return ""
}

type AbortResponse struct {
	Error string `protobuf:"bytes,1,opt,name=Error" json:"Error,omitempty"`
}

func (m *AbortResponse) Reset()                    { *m = AbortResponse{} }
func (m *AbortResponse) String() string            { return proto.CompactTextString(m) }
func (*AbortResponse) ProtoMessage()               {}
func (*AbortResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AbortResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*TerminalStateRequest)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalStateRequest")
	proto.RegisterType((*TerminalStateResponse)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalStateResponse")
	proto.RegisterType((*TerminalStateResponse_Account)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalStateResponse.Account")
	proto.RegisterType((*TerminalStateResponse_Order)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalStateResponse.Order")
	proto.RegisterType((*TerminalBuyRequest)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalBuyRequest")
	proto.RegisterType((*TerminalBuyResponse)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalBuyResponse")
	proto.RegisterType((*TerminalScanRequest)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalScanRequest")
	proto.RegisterType((*TerminalScanResponse)(nil), "i6getraenkeabrechnungssystem3000.rpc.TerminalScanResponse")
	proto.RegisterType((*AbortRequest)(nil), "i6getraenkeabrechnungssystem3000.rpc.AbortRequest")
	proto.RegisterType((*AbortResponse)(nil), "i6getraenkeabrechnungssystem3000.rpc.AbortResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TerminalBackend service

type TerminalBackendClient interface {
	GetState(ctx context.Context, in *TerminalStateRequest, opts ...grpc.CallOption) (*TerminalStateResponse, error)
	Buy(ctx context.Context, in *TerminalBuyRequest, opts ...grpc.CallOption) (*TerminalBuyResponse, error)
	Scan(ctx context.Context, in *TerminalScanRequest, opts ...grpc.CallOption) (*TerminalScanResponse, error)
	Abort(ctx context.Context, in *AbortRequest, opts ...grpc.CallOption) (*AbortResponse, error)
}

type terminalBackendClient struct {
	cc *grpc.ClientConn
}

func NewTerminalBackendClient(cc *grpc.ClientConn) TerminalBackendClient {
	return &terminalBackendClient{cc}
}

func (c *terminalBackendClient) GetState(ctx context.Context, in *TerminalStateRequest, opts ...grpc.CallOption) (*TerminalStateResponse, error) {
	out := new(TerminalStateResponse)
	err := grpc.Invoke(ctx, "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/GetState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terminalBackendClient) Buy(ctx context.Context, in *TerminalBuyRequest, opts ...grpc.CallOption) (*TerminalBuyResponse, error) {
	out := new(TerminalBuyResponse)
	err := grpc.Invoke(ctx, "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Buy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terminalBackendClient) Scan(ctx context.Context, in *TerminalScanRequest, opts ...grpc.CallOption) (*TerminalScanResponse, error) {
	out := new(TerminalScanResponse)
	err := grpc.Invoke(ctx, "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Scan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terminalBackendClient) Abort(ctx context.Context, in *AbortRequest, opts ...grpc.CallOption) (*AbortResponse, error) {
	out := new(AbortResponse)
	err := grpc.Invoke(ctx, "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Abort", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TerminalBackend service

type TerminalBackendServer interface {
	GetState(context.Context, *TerminalStateRequest) (*TerminalStateResponse, error)
	Buy(context.Context, *TerminalBuyRequest) (*TerminalBuyResponse, error)
	Scan(context.Context, *TerminalScanRequest) (*TerminalScanResponse, error)
	Abort(context.Context, *AbortRequest) (*AbortResponse, error)
}

func RegisterTerminalBackendServer(s *grpc.Server, srv TerminalBackendServer) {
	s.RegisterService(&_TerminalBackend_serviceDesc, srv)
}

func _TerminalBackend_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminalStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalBackendServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalBackendServer).GetState(ctx, req.(*TerminalStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerminalBackend_Buy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminalBuyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalBackendServer).Buy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Buy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalBackendServer).Buy(ctx, req.(*TerminalBuyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerminalBackend_Scan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminalScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalBackendServer).Scan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Scan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalBackendServer).Scan(ctx, req.(*TerminalScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerminalBackend_Abort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AbortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalBackendServer).Abort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/i6getraenkeabrechnungssystem3000.rpc.TerminalBackend/Abort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalBackendServer).Abort(ctx, req.(*AbortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TerminalBackend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "i6getraenkeabrechnungssystem3000.rpc.TerminalBackend",
	HandlerType: (*TerminalBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetState",
			Handler:    _TerminalBackend_GetState_Handler,
		},
		{
			MethodName: "Buy",
			Handler:    _TerminalBackend_Buy_Handler,
		},
		{
			MethodName: "Scan",
			Handler:    _TerminalBackend_Scan_Handler,
		},
		{
			MethodName: "Abort",
			Handler:    _TerminalBackend_Abort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}

func init() { proto.RegisterFile("main.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0xcd, 0xa6, 0xdb, 0x76, 0x5f, 0xad, 0xd2, 0xb1, 0x4a, 0x08, 0x22, 0x4b, 0x50, 0x58,
	0x50, 0xc2, 0xb2, 0x0b, 0x45, 0xeb, 0x29, 0x31, 0x22, 0x0b, 0xa2, 0x65, 0xd6, 0x5e, 0xbc, 0x2c,
	0xb3, 0xb3, 0x43, 0x0d, 0xdd, 0x4c, 0xe2, 0xcc, 0xe4, 0x10, 0x10, 0x3c, 0x79, 0xf3, 0xe2, 0x7f,
	0x2c, 0x3b, 0x9d, 0xfc, 0xb0, 0x96, 0x92, 0x6c, 0x8f, 0xef, 0xed, 0xbc, 0xcf, 0xf7, 0xfb, 0xf6,
	0x7d, 0x09, 0x40, 0x42, 0x62, 0xee, 0x67, 0x22, 0x55, 0x29, 0x7a, 0x1e, 0x9f, 0x5c, 0x30, 0x25,
	0x08, 0xe3, 0x97, 0x8c, 0x2c, 0x05, 0xa3, 0xdf, 0x78, 0xce, 0x2f, 0xa4, 0x2c, 0xa4, 0x62, 0xc9,
	0x74, 0x3c, 0x1e, 0xfb, 0x22, 0xa3, 0xde, 0x09, 0x1c, 0x7f, 0x61, 0x22, 0x89, 0x39, 0x59, 0xcf,
	0x15, 0x51, 0x0c, 0xb3, 0xef, 0x39, 0x93, 0x0a, 0x3d, 0x03, 0x28, 0xfb, 0xb3, 0xc8, 0xb1, 0x86,
	0xd6, 0x68, 0x80, 0x1b, 0x1d, 0xef, 0x8f, 0x0d, 0x8f, 0xaf, 0x0d, 0xca, 0x2c, 0xe5, 0x92, 0xa1,
	0x05, 0xec, 0x07, 0x94, 0xa6, 0x39, 0x57, 0xd2, 0xb1, 0x86, 0xf6, 0xe8, 0x60, 0xf2, 0xce, 0x6f,
	0x63, 0xc5, 0xbf, 0x11, 0xe7, 0x1b, 0x16, 0xae, 0xa0, 0x68, 0x01, 0x83, 0xcf, 0x62, 0xc5, 0xc4,
	0xc7, 0x58, 0x2a, 0xa7, 0xa7, 0x15, 0x82, 0xbb, 0x28, 0x68, 0x18, 0xae, 0x99, 0xee, 0x39, 0xec,
	0x19, 0x31, 0xf4, 0x00, 0x7a, 0xd5, 0xfa, 0xbd, 0x59, 0x84, 0x86, 0x70, 0x10, 0xc5, 0x32, 0x5b,
	0x93, 0xe2, 0x13, 0x49, 0x98, 0xd3, 0xd3, 0x3f, 0x34, 0x5b, 0xc8, 0x81, 0xbd, 0x90, 0xac, 0x09,
	0xa7, 0xcc, 0xb1, 0x87, 0xd6, 0xe8, 0x08, 0x97, 0xa5, 0xbb, 0x80, 0xbe, 0xd6, 0xb8, 0x0e, 0xb1,
	0xfe, 0x87, 0x3c, 0x85, 0xc1, 0x39, 0x8f, 0xd5, 0x99, 0x88, 0xe9, 0x95, 0xc8, 0x11, 0xae, 0x1b,
	0xe8, 0x09, 0xec, 0x06, 0xc9, 0xc6, 0x9e, 0x56, 0x38, 0xc4, 0xa6, 0xf2, 0x30, 0xa0, 0x72, 0xc3,
	0x30, 0x2f, 0x5a, 0x5e, 0x72, 0xa3, 0x65, 0xb6, 0x9d, 0x45, 0x66, 0xa1, 0xba, 0xe1, 0xbd, 0x84,
	0x47, 0xff, 0x30, 0xcd, 0x91, 0x8f, 0xa1, 0xff, 0x5e, 0x88, 0x54, 0x18, 0xde, 0x55, 0xe1, 0xcd,
	0xeb, 0xc7, 0x73, 0x4a, 0x78, 0x07, 0x07, 0x67, 0x22, 0x5d, 0xe5, 0xb4, 0xe1, 0xa0, 0x6a, 0x78,
	0xaf, 0x1a, 0x09, 0xd5, 0xd0, 0x5b, 0x2d, 0xf8, 0x70, 0x3f, 0x58, 0xa6, 0x42, 0xb5, 0xcd, 0xf1,
	0x0b, 0x38, 0x34, 0xef, 0x6f, 0xc3, 0x4e, 0x7e, 0xef, 0xc0, 0xc3, 0xea, 0x7f, 0x20, 0xf4, 0x92,
	0xf1, 0x15, 0xfa, 0x65, 0xc1, 0xfe, 0x07, 0xa6, 0x74, 0x98, 0xd0, 0xe9, 0x56, 0x09, 0xd4, 0x1e,
	0xdd, 0xb7, 0x77, 0x48, 0xaf, 0x77, 0x0f, 0xfd, 0x00, 0x3b, 0xcc, 0x0b, 0xf4, 0xba, 0x1b, 0xa5,
	0x4e, 0x88, 0xfb, 0x66, 0x8b, 0xc9, 0x4a, 0xfd, 0x27, 0xec, 0x6c, 0xce, 0x82, 0x3a, 0x42, 0x1a,
	0xf9, 0x70, 0x4f, 0xb7, 0x19, 0xad, 0x0c, 0x08, 0xe8, 0xeb, 0x0b, 0xa2, 0x49, 0x3b, 0x4c, 0x33,
	0x1e, 0xee, 0xb4, 0xd3, 0x4c, 0xa9, 0x19, 0xf6, 0xbf, 0xda, 0x22, 0xa3, 0xcb, 0x5d, 0xfd, 0xa5,
	0x9d, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xde, 0x09, 0x71, 0x77, 0x05, 0x00, 0x00,
}
