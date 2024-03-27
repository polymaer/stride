// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/staketia/genesis.proto

package types

import (
	fmt "fmt"
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

// Params defines the staketia module parameters.
type Params struct {
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2938dee39d417bb, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// TransferInProgressRecordIds stores record IDs for delegation records
// that have a transfer in progress
type TransferInProgressRecordIds struct {
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Sequence  uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	RecordId  uint64 `protobuf:"varint,3,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
}

func (m *TransferInProgressRecordIds) Reset()         { *m = TransferInProgressRecordIds{} }
func (m *TransferInProgressRecordIds) String() string { return proto.CompactTextString(m) }
func (*TransferInProgressRecordIds) ProtoMessage()    {}
func (*TransferInProgressRecordIds) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2938dee39d417bb, []int{1}
}
func (m *TransferInProgressRecordIds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransferInProgressRecordIds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransferInProgressRecordIds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransferInProgressRecordIds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferInProgressRecordIds.Merge(m, src)
}
func (m *TransferInProgressRecordIds) XXX_Size() int {
	return m.Size()
}
func (m *TransferInProgressRecordIds) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferInProgressRecordIds.DiscardUnknown(m)
}

var xxx_messageInfo_TransferInProgressRecordIds proto.InternalMessageInfo

func (m *TransferInProgressRecordIds) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *TransferInProgressRecordIds) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *TransferInProgressRecordIds) GetRecordId() uint64 {
	if m != nil {
		return m.RecordId
	}
	return 0
}

// GenesisState defines the staketia module's genesis state.
type GenesisState struct {
	Params                      Params                        `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params"`
	HostZone                    HostZone                      `protobuf:"bytes,2,opt,name=host_zone,json=hostZone,proto3" json:"host_zone"`
	DelegationRecords           []DelegationRecord            `protobuf:"bytes,3,rep,name=delegation_records,json=delegationRecords,proto3" json:"delegation_records"`
	UnbondingRecords            []UnbondingRecord             `protobuf:"bytes,4,rep,name=unbonding_records,json=unbondingRecords,proto3" json:"unbonding_records"`
	RedemptionRecords           []RedemptionRecord            `protobuf:"bytes,5,rep,name=redemption_records,json=redemptionRecords,proto3" json:"redemption_records"`
	SlashRecords                []SlashRecord                 `protobuf:"bytes,6,rep,name=slash_records,json=slashRecords,proto3" json:"slash_records"`
	TransferInProgressRecordIds []TransferInProgressRecordIds `protobuf:"bytes,7,rep,name=transfer_in_progress_record_ids,json=transferInProgressRecordIds,proto3" json:"transfer_in_progress_record_ids"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2938dee39d417bb, []int{2}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetHostZone() HostZone {
	if m != nil {
		return m.HostZone
	}
	return HostZone{}
}

func (m *GenesisState) GetDelegationRecords() []DelegationRecord {
	if m != nil {
		return m.DelegationRecords
	}
	return nil
}

func (m *GenesisState) GetUnbondingRecords() []UnbondingRecord {
	if m != nil {
		return m.UnbondingRecords
	}
	return nil
}

func (m *GenesisState) GetRedemptionRecords() []RedemptionRecord {
	if m != nil {
		return m.RedemptionRecords
	}
	return nil
}

func (m *GenesisState) GetSlashRecords() []SlashRecord {
	if m != nil {
		return m.SlashRecords
	}
	return nil
}

func (m *GenesisState) GetTransferInProgressRecordIds() []TransferInProgressRecordIds {
	if m != nil {
		return m.TransferInProgressRecordIds
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "stride.staketia.Params")
	proto.RegisterType((*TransferInProgressRecordIds)(nil), "stride.staketia.TransferInProgressRecordIds")
	proto.RegisterType((*GenesisState)(nil), "stride.staketia.GenesisState")
}

func init() { proto.RegisterFile("stride/staketia/genesis.proto", fileDescriptor_e2938dee39d417bb) }

var fileDescriptor_e2938dee39d417bb = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x63, 0x12, 0x42, 0xb2, 0x69, 0x05, 0x5d, 0x81, 0x08, 0x09, 0x75, 0x82, 0x4f, 0x39,
	0x80, 0x2d, 0xcc, 0x0d, 0x71, 0xaa, 0x10, 0x25, 0x52, 0x0f, 0x95, 0x03, 0x1c, 0x7a, 0xb1, 0x36,
	0xf1, 0x60, 0x5b, 0x24, 0xbb, 0x66, 0x67, 0x8d, 0x5a, 0x9e, 0x82, 0xc7, 0xea, 0xb1, 0x47, 0x0e,
	0xa8, 0x42, 0xc9, 0x1b, 0xf0, 0x04, 0xc8, 0x5e, 0xc7, 0x95, 0x63, 0xb5, 0xb7, 0xdd, 0x99, 0x7f,
	0xbe, 0xf9, 0x77, 0x35, 0x43, 0x0e, 0x51, 0xc9, 0x38, 0x00, 0x07, 0x15, 0xfb, 0x06, 0x2a, 0x66,
	0x4e, 0x08, 0x1c, 0x30, 0x46, 0x3b, 0x91, 0x42, 0x09, 0xfa, 0x50, 0xa7, 0xed, 0x6d, 0x7a, 0xf0,
	0x38, 0x14, 0xa1, 0xc8, 0x73, 0x4e, 0x76, 0xd2, 0xb2, 0x81, 0xb9, 0x4b, 0xd9, 0x1e, 0x74, 0xde,
	0xea, 0x90, 0xf6, 0x29, 0x93, 0x6c, 0x85, 0x56, 0x4a, 0x86, 0x9f, 0x24, 0xe3, 0xf8, 0x15, 0xe4,
	0x94, 0x9f, 0x4a, 0x11, 0x4a, 0x40, 0xf4, 0x60, 0x21, 0x64, 0x30, 0x0d, 0x90, 0x1e, 0x12, 0xb2,
	0x88, 0x18, 0xe7, 0xb0, 0xf4, 0xe3, 0xa0, 0x6f, 0x8c, 0x8d, 0x49, 0xd7, 0xeb, 0x16, 0x91, 0x69,
	0x40, 0x07, 0xa4, 0x83, 0xf0, 0x3d, 0x05, 0xbe, 0x80, 0xfe, 0xbd, 0xb1, 0x31, 0x69, 0x79, 0xe5,
	0x9d, 0x0e, 0x49, 0x57, 0xe6, 0x9c, 0xac, 0xb2, 0xa9, 0x93, 0xb2, 0x00, 0x5b, 0x7f, 0x5a, 0x64,
	0xef, 0x58, 0xbf, 0x6c, 0xa6, 0x98, 0x02, 0xfa, 0x81, 0xb4, 0x93, 0xdc, 0x51, 0xde, 0xa4, 0xe7,
	0x3e, 0xb5, 0x77, 0x5e, 0x6a, 0x6b, 0xc3, 0x47, 0x4f, 0x2e, 0xaf, 0x47, 0x8d, 0x7f, 0xd7, 0xa3,
	0xfd, 0x0b, 0xb6, 0x5a, 0xbe, 0xb5, 0x74, 0x91, 0xe5, 0x15, 0xd5, 0xf4, 0x1d, 0xe9, 0x46, 0x02,
	0x95, 0xff, 0x53, 0x70, 0x6d, 0xa9, 0xe7, 0x3e, 0xab, 0xa1, 0x3e, 0x0a, 0x54, 0x67, 0x82, 0xc3,
	0x51, 0x2b, 0x83, 0x79, 0x9d, 0xa8, 0xb8, 0xd3, 0x2f, 0x84, 0x06, 0xb0, 0x84, 0x90, 0xa9, 0x58,
	0x70, 0x5f, 0xbb, 0xc5, 0x7e, 0x73, 0xdc, 0x9c, 0xf4, 0xdc, 0x17, 0x35, 0xcc, 0xfb, 0x52, 0xaa,
	0x3f, 0xac, 0xc0, 0x1d, 0x04, 0x3b, 0x71, 0xa4, 0x33, 0x72, 0x90, 0xf2, 0xb9, 0xe0, 0x41, 0xcc,
	0xc3, 0x12, 0xdb, 0xca, 0xb1, 0xe3, 0x1a, 0xf6, 0xf3, 0x56, 0x59, 0xa1, 0x3e, 0x4a, 0xab, 0x61,
	0xcc, 0xcc, 0x4a, 0x08, 0x60, 0x95, 0x54, 0xcc, 0xde, 0xbf, 0xc5, 0xac, 0x57, 0x4a, 0xab, 0x66,
	0xe5, 0x4e, 0x1c, 0xe9, 0x31, 0xd9, 0xc7, 0x25, 0xc3, 0xa8, 0x44, 0xb6, 0x73, 0xe4, 0xf3, 0x1a,
	0x72, 0x96, 0xa9, 0x2a, 0xb4, 0x3d, 0xbc, 0x09, 0x21, 0x3d, 0x27, 0x23, 0x55, 0xcc, 0x96, 0x1f,
	0x73, 0x3f, 0x29, 0xa6, 0xcb, 0x2f, 0xc7, 0x02, 0xfb, 0x0f, 0x72, 0xf4, 0xcb, 0x1a, 0xfa, 0x8e,
	0x99, 0x2c, 0x5a, 0x0d, 0xd5, 0x1d, 0x92, 0x93, 0xcb, 0xb5, 0x69, 0x5c, 0xad, 0x4d, 0xe3, 0xef,
	0xda, 0x34, 0x7e, 0x6d, 0xcc, 0xc6, 0xd5, 0xc6, 0x6c, 0xfc, 0xde, 0x98, 0x8d, 0x33, 0x37, 0x8c,
	0x55, 0x94, 0xce, 0xed, 0x85, 0x58, 0x39, 0xb3, 0xbc, 0xe9, 0xab, 0x13, 0x36, 0x47, 0xa7, 0x58,
	0x98, 0x1f, 0xee, 0x6b, 0xe7, 0xfc, 0x66, 0x6d, 0xd4, 0x45, 0x02, 0x38, 0x6f, 0xe7, 0x4b, 0xf3,
	0xe6, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x3d, 0xbd, 0x4a, 0x9c, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *TransferInProgressRecordIds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransferInProgressRecordIds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransferInProgressRecordIds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RecordId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RecordId))
		i--
		dAtA[i] = 0x18
	}
	if m.Sequence != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TransferInProgressRecordIds) > 0 {
		for iNdEx := len(m.TransferInProgressRecordIds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TransferInProgressRecordIds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.SlashRecords) > 0 {
		for iNdEx := len(m.SlashRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SlashRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.RedemptionRecords) > 0 {
		for iNdEx := len(m.RedemptionRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RedemptionRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.UnbondingRecords) > 0 {
		for iNdEx := len(m.UnbondingRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UnbondingRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.DelegationRecords) > 0 {
		for iNdEx := len(m.DelegationRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegationRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.HostZone.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *TransferInProgressRecordIds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovGenesis(uint64(m.Sequence))
	}
	if m.RecordId != 0 {
		n += 1 + sovGenesis(uint64(m.RecordId))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.HostZone.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.DelegationRecords) > 0 {
		for _, e := range m.DelegationRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UnbondingRecords) > 0 {
		for _, e := range m.UnbondingRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RedemptionRecords) > 0 {
		for _, e := range m.RedemptionRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SlashRecords) > 0 {
		for _, e := range m.SlashRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TransferInProgressRecordIds) > 0 {
		for _, e := range m.TransferInProgressRecordIds {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *TransferInProgressRecordIds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: TransferInProgressRecordIds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransferInProgressRecordIds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordId", wireType)
			}
			m.RecordId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RecordId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostZone", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HostZone.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegationRecords = append(m.DelegationRecords, DelegationRecord{})
			if err := m.DelegationRecords[len(m.DelegationRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnbondingRecords = append(m.UnbondingRecords, UnbondingRecord{})
			if err := m.UnbondingRecords[len(m.UnbondingRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RedemptionRecords = append(m.RedemptionRecords, RedemptionRecord{})
			if err := m.RedemptionRecords[len(m.RedemptionRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SlashRecords = append(m.SlashRecords, SlashRecord{})
			if err := m.SlashRecords[len(m.SlashRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferInProgressRecordIds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransferInProgressRecordIds = append(m.TransferInProgressRecordIds, TransferInProgressRecordIds{})
			if err := m.TransferInProgressRecordIds[len(m.TransferInProgressRecordIds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
