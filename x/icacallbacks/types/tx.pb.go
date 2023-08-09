// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/icacallbacks/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("stride/icacallbacks/tx.proto", fileDescriptor_c4981fec5f7fee51) }

var fileDescriptor_c4981fec5f7fee51 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x2e, 0x29, 0xca,
	0x4c, 0x49, 0xd5, 0xcf, 0x4c, 0x4e, 0x4c, 0x4e, 0xcc, 0xc9, 0x49, 0x4a, 0x4c, 0xce, 0x2e, 0xd6,
	0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86, 0xc8, 0xea, 0x21, 0xcb, 0x1a,
	0xb1, 0x72, 0x31, 0xfb, 0x16, 0xa7, 0x3b, 0x05, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1,
	0x1c, 0x43, 0x94, 0x59, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x30,
	0xd8, 0x00, 0x5d, 0x9f, 0xc4, 0xa4, 0x62, 0x7d, 0xa8, 0x55, 0x65, 0x86, 0x46, 0xfa, 0x15, 0x68,
	0x16, 0x56, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x2d, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xc7, 0xd9, 0xf7, 0x48, 0x94, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stride.icacallbacks.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "stride/icacallbacks/tx.proto",
}
