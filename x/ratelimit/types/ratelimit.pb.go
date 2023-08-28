// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/ratelimit/ratelimit.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type PacketDirection int32

const (
	PACKET_SEND PacketDirection = 0
	PACKET_RECV PacketDirection = 1
)

var PacketDirection_name = map[int32]string{
	0: "PACKET_SEND",
	1: "PACKET_RECV",
}

var PacketDirection_value = map[string]int32{
	"PACKET_SEND": 0,
	"PACKET_RECV": 1,
}

func (x PacketDirection) String() string {
	return proto.EnumName(PacketDirection_name, int32(x))
}

func (PacketDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a3e00ee2c967d747, []int{0}
}

type Path struct {
	Denom     string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (m *Path) Reset()         { *m = Path{} }
func (m *Path) String() string { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()    {}
func (*Path) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3e00ee2c967d747, []int{0}
}
func (m *Path) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Path) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Path.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Path) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Path.Merge(m, src)
}
func (m *Path) XXX_Size() int {
	return m.Size()
}
func (m *Path) XXX_DiscardUnknown() {
	xxx_messageInfo_Path.DiscardUnknown(m)
}

var xxx_messageInfo_Path proto.InternalMessageInfo

func (m *Path) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Path) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

type Quota struct {
	MaxPercentSend github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=max_percent_send,json=maxPercentSend,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_percent_send"`
	MaxPercentRecv github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=max_percent_recv,json=maxPercentRecv,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_percent_recv"`
	DurationHours  uint64                                 `protobuf:"varint,3,opt,name=duration_hours,json=durationHours,proto3" json:"duration_hours,omitempty"`
}

func (m *Quota) Reset()         { *m = Quota{} }
func (m *Quota) String() string { return proto.CompactTextString(m) }
func (*Quota) ProtoMessage()    {}
func (*Quota) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3e00ee2c967d747, []int{1}
}
func (m *Quota) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Quota) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Quota.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Quota) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quota.Merge(m, src)
}
func (m *Quota) XXX_Size() int {
	return m.Size()
}
func (m *Quota) XXX_DiscardUnknown() {
	xxx_messageInfo_Quota.DiscardUnknown(m)
}

var xxx_messageInfo_Quota proto.InternalMessageInfo

func (m *Quota) GetDurationHours() uint64 {
	if m != nil {
		return m.DurationHours
	}
	return 0
}

type Flow struct {
	Inflow       github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=inflow,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"inflow"`
	Outflow      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=outflow,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"outflow"`
	ChannelValue github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=channel_value,json=channelValue,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"channel_value"`
}

func (m *Flow) Reset()         { *m = Flow{} }
func (m *Flow) String() string { return proto.CompactTextString(m) }
func (*Flow) ProtoMessage()    {}
func (*Flow) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3e00ee2c967d747, []int{2}
}
func (m *Flow) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Flow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Flow.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Flow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Flow.Merge(m, src)
}
func (m *Flow) XXX_Size() int {
	return m.Size()
}
func (m *Flow) XXX_DiscardUnknown() {
	xxx_messageInfo_Flow.DiscardUnknown(m)
}

var xxx_messageInfo_Flow proto.InternalMessageInfo

type RateLimit struct {
	Path  *Path  `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Quota *Quota `protobuf:"bytes,2,opt,name=quota,proto3" json:"quota,omitempty"`
	Flow  *Flow  `protobuf:"bytes,3,opt,name=flow,proto3" json:"flow,omitempty"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3e00ee2c967d747, []int{3}
}
func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return m.Size()
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *RateLimit) GetQuota() *Quota {
	if m != nil {
		return m.Quota
	}
	return nil
}

func (m *RateLimit) GetFlow() *Flow {
	if m != nil {
		return m.Flow
	}
	return nil
}

type WhitelistedAddressPair struct {
	Sender   string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
}

func (m *WhitelistedAddressPair) Reset()         { *m = WhitelistedAddressPair{} }
func (m *WhitelistedAddressPair) String() string { return proto.CompactTextString(m) }
func (*WhitelistedAddressPair) ProtoMessage()    {}
func (*WhitelistedAddressPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3e00ee2c967d747, []int{4}
}
func (m *WhitelistedAddressPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhitelistedAddressPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhitelistedAddressPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhitelistedAddressPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhitelistedAddressPair.Merge(m, src)
}
func (m *WhitelistedAddressPair) XXX_Size() int {
	return m.Size()
}
func (m *WhitelistedAddressPair) XXX_DiscardUnknown() {
	xxx_messageInfo_WhitelistedAddressPair.DiscardUnknown(m)
}

var xxx_messageInfo_WhitelistedAddressPair proto.InternalMessageInfo

func (m *WhitelistedAddressPair) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *WhitelistedAddressPair) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func init() {
	proto.RegisterEnum("stride.ratelimit.PacketDirection", PacketDirection_name, PacketDirection_value)
	proto.RegisterType((*Path)(nil), "stride.ratelimit.Path")
	proto.RegisterType((*Quota)(nil), "stride.ratelimit.Quota")
	proto.RegisterType((*Flow)(nil), "stride.ratelimit.Flow")
	proto.RegisterType((*RateLimit)(nil), "stride.ratelimit.RateLimit")
	proto.RegisterType((*WhitelistedAddressPair)(nil), "stride.ratelimit.WhitelistedAddressPair")
}

func init() { proto.RegisterFile("stride/ratelimit/ratelimit.proto", fileDescriptor_a3e00ee2c967d747) }

var fileDescriptor_a3e00ee2c967d747 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xb5, 0xbf, 0x26, 0xf9, 0xc8, 0x84, 0xb6, 0xd1, 0xa8, 0x0a, 0x51, 0x24, 0xdc, 0x28, 0x12,
	0xa8, 0xaa, 0x14, 0x5b, 0xb4, 0x6c, 0x10, 0xab, 0xfe, 0xa4, 0x6a, 0x45, 0x40, 0xc1, 0x41, 0x05,
	0xb1, 0x89, 0x26, 0x9e, 0x4b, 0x3c, 0x6a, 0xec, 0x09, 0x33, 0x63, 0x37, 0xbc, 0x01, 0x4b, 0xc4,
	0x2b, 0xf0, 0x32, 0x5d, 0x76, 0x89, 0x58, 0x54, 0x28, 0x59, 0xf3, 0x0e, 0x68, 0xc6, 0x4e, 0x89,
	0x8a, 0x58, 0x50, 0x56, 0x9e, 0x7b, 0xee, 0x99, 0x23, 0x9f, 0x3b, 0xe7, 0xa2, 0xa6, 0x54, 0x82,
	0x51, 0xf0, 0x04, 0x51, 0x30, 0x66, 0x11, 0x53, 0xbf, 0x4e, 0xee, 0x44, 0x70, 0xc5, 0x71, 0x35,
	0x63, 0xb8, 0xd7, 0x78, 0x63, 0x63, 0xc4, 0x47, 0xdc, 0x34, 0x3d, 0x7d, 0xca, 0x78, 0xad, 0xa7,
	0xa8, 0xd0, 0x23, 0x2a, 0xc4, 0x1b, 0xa8, 0x48, 0x21, 0xe6, 0x51, 0xdd, 0x6e, 0xda, 0x5b, 0x65,
	0x3f, 0x2b, 0xf0, 0x7d, 0x84, 0x82, 0x90, 0xc4, 0x31, 0x8c, 0x07, 0x8c, 0xd6, 0xff, 0x33, 0xad,
	0x72, 0x8e, 0x9c, 0xd0, 0xd6, 0xcc, 0x46, 0xc5, 0x97, 0x09, 0x57, 0x04, 0xbf, 0x41, 0xd5, 0x88,
	0x4c, 0x07, 0x13, 0x10, 0x01, 0xc4, 0x6a, 0x20, 0x21, 0xa6, 0x99, 0xd2, 0xbe, 0x7b, 0x71, 0xb5,
	0x69, 0x7d, 0xbb, 0xda, 0x7c, 0x38, 0x62, 0x2a, 0x4c, 0x86, 0x6e, 0xc0, 0x23, 0x2f, 0xe0, 0x32,
	0xe2, 0x32, 0xff, 0xb4, 0x25, 0x3d, 0xf3, 0xd4, 0x87, 0x09, 0x48, 0xf7, 0x24, 0x56, 0xfe, 0x5a,
	0x44, 0xa6, 0xbd, 0x4c, 0xa6, 0x0f, 0x31, 0xbd, 0xa9, 0x2c, 0x20, 0x48, 0xb3, 0x1f, 0xf9, 0x17,
	0x65, 0x1f, 0x82, 0x14, 0x3f, 0x40, 0x6b, 0x34, 0x11, 0x44, 0x31, 0x1e, 0x0f, 0x42, 0x9e, 0x08,
	0x59, 0x5f, 0x69, 0xda, 0x5b, 0x05, 0x7f, 0x75, 0x81, 0x1e, 0x6b, 0xb0, 0xf5, 0xc3, 0x46, 0x85,
	0xa3, 0x31, 0x3f, 0xc7, 0x47, 0xa8, 0xc4, 0xe2, 0x77, 0x63, 0x7e, 0x7e, 0x4b, 0x67, 0xf9, 0x6d,
	0x7c, 0x8c, 0xfe, 0xe7, 0x89, 0x32, 0x42, 0xb7, 0x33, 0xb2, 0xb8, 0x8e, 0xfb, 0x68, 0x75, 0xf1,
	0x3c, 0x29, 0x19, 0x27, 0x60, 0x0c, 0xfc, 0xbd, 0xde, 0xdd, 0x5c, 0xe4, 0x54, 0x6b, 0xb4, 0x3e,
	0xdb, 0xa8, 0xec, 0x13, 0x05, 0x5d, 0x9d, 0x1a, 0xbc, 0x8d, 0x0a, 0x13, 0xa2, 0x42, 0x63, 0xb9,
	0xb2, 0x53, 0x73, 0x6f, 0xc6, 0xca, 0xd5, 0xe9, 0xf1, 0x0d, 0x07, 0xb7, 0x51, 0xf1, 0xbd, 0x4e,
	0x83, 0xb1, 0x55, 0xd9, 0xb9, 0xf7, 0x3b, 0xd9, 0x84, 0xc5, 0xcf, 0x58, 0x5a, 0xda, 0x0c, 0x61,
	0xe5, 0x4f, 0xd2, 0x7a, 0xea, 0xbe, 0xe1, 0xb4, 0xba, 0xa8, 0xf6, 0x3a, 0x64, 0xba, 0x21, 0x15,
	0xd0, 0x3d, 0x4a, 0x05, 0x48, 0xd9, 0x23, 0x4c, 0xe0, 0x1a, 0x2a, 0xe9, 0xb4, 0x81, 0xc8, 0x93,
	0x9b, 0x57, 0xb8, 0x81, 0xee, 0x08, 0x08, 0x80, 0xa5, 0x20, 0xf2, 0xe0, 0x5e, 0xd7, 0xdb, 0x4f,
	0xd0, 0x7a, 0x8f, 0x04, 0x67, 0xa0, 0x0e, 0x99, 0x80, 0x40, 0x3f, 0x35, 0x5e, 0x47, 0x95, 0xde,
	0xde, 0xc1, 0xb3, 0xce, 0xab, 0x41, 0xbf, 0xf3, 0xe2, 0xb0, 0x6a, 0x2d, 0x01, 0x7e, 0xe7, 0xe0,
	0xb4, 0x6a, 0x37, 0x0a, 0x1f, 0xbf, 0x38, 0xd6, 0xfe, 0xf3, 0x8b, 0x99, 0x63, 0x5f, 0xce, 0x1c,
	0xfb, 0xfb, 0xcc, 0xb1, 0x3f, 0xcd, 0x1d, 0xeb, 0x72, 0xee, 0x58, 0x5f, 0xe7, 0x8e, 0xf5, 0x76,
	0x77, 0x69, 0xda, 0x7d, 0x63, 0xa5, 0xdd, 0x25, 0x43, 0xe9, 0xe5, 0xab, 0x9a, 0x3e, 0x7a, 0xec,
	0x4d, 0x97, 0x16, 0xd6, 0x8c, 0x7f, 0x58, 0x32, 0x5b, 0xb8, 0xfb, 0x33, 0x00, 0x00, 0xff, 0xff,
	0x2b, 0x0e, 0xae, 0xb0, 0xd1, 0x03, 0x00, 0x00,
}

func (m *Path) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Path) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Path) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintRatelimit(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintRatelimit(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Quota) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Quota) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Quota) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DurationHours != 0 {
		i = encodeVarintRatelimit(dAtA, i, uint64(m.DurationHours))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.MaxPercentRecv.Size()
		i -= size
		if _, err := m.MaxPercentRecv.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRatelimit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MaxPercentSend.Size()
		i -= size
		if _, err := m.MaxPercentSend.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRatelimit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Flow) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Flow) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Flow) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ChannelValue.Size()
		i -= size
		if _, err := m.ChannelValue.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRatelimit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Outflow.Size()
		i -= size
		if _, err := m.Outflow.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRatelimit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Inflow.Size()
		i -= size
		if _, err := m.Inflow.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRatelimit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RateLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RateLimit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Flow != nil {
		{
			size, err := m.Flow.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRatelimit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Quota != nil {
		{
			size, err := m.Quota.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRatelimit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Path != nil {
		{
			size, err := m.Path.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRatelimit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WhitelistedAddressPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhitelistedAddressPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhitelistedAddressPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintRatelimit(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintRatelimit(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRatelimit(dAtA []byte, offset int, v uint64) int {
	offset -= sovRatelimit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Path) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovRatelimit(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovRatelimit(uint64(l))
	}
	return n
}

func (m *Quota) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MaxPercentSend.Size()
	n += 1 + l + sovRatelimit(uint64(l))
	l = m.MaxPercentRecv.Size()
	n += 1 + l + sovRatelimit(uint64(l))
	if m.DurationHours != 0 {
		n += 1 + sovRatelimit(uint64(m.DurationHours))
	}
	return n
}

func (m *Flow) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflow.Size()
	n += 1 + l + sovRatelimit(uint64(l))
	l = m.Outflow.Size()
	n += 1 + l + sovRatelimit(uint64(l))
	l = m.ChannelValue.Size()
	n += 1 + l + sovRatelimit(uint64(l))
	return n
}

func (m *RateLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Path != nil {
		l = m.Path.Size()
		n += 1 + l + sovRatelimit(uint64(l))
	}
	if m.Quota != nil {
		l = m.Quota.Size()
		n += 1 + l + sovRatelimit(uint64(l))
	}
	if m.Flow != nil {
		l = m.Flow.Size()
		n += 1 + l + sovRatelimit(uint64(l))
	}
	return n
}

func (m *WhitelistedAddressPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovRatelimit(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovRatelimit(uint64(l))
	}
	return n
}

func sovRatelimit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRatelimit(x uint64) (n int) {
	return sovRatelimit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Path) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRatelimit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Path: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Path: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRatelimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRatelimit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Quota) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRatelimit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Quota: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Quota: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPercentSend", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxPercentSend.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPercentRecv", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxPercentRecv.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationHours", wireType)
			}
			m.DurationHours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DurationHours |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRatelimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRatelimit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Flow) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRatelimit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Flow: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Flow: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflow", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Outflow", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Outflow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChannelValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRatelimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRatelimit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RateLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRatelimit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RateLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Path == nil {
				m.Path = &Path{}
			}
			if err := m.Path.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quota", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Quota == nil {
				m.Quota = &Quota{}
			}
			if err := m.Quota.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flow", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Flow == nil {
				m.Flow = &Flow{}
			}
			if err := m.Flow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRatelimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRatelimit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WhitelistedAddressPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRatelimit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WhitelistedAddressPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhitelistedAddressPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRatelimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRatelimit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRatelimit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRatelimit
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRatelimit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRatelimit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRatelimit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRatelimit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRatelimit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRatelimit = fmt.Errorf("proto: unexpected end of group")
)
