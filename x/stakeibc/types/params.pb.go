// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/stakeibc/params.proto

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

// Params defines the parameters for the module.
// next id: 18
type Params struct {
	// define epoch lengths, in stride_epochs
	RewardsInterval                   uint64 `protobuf:"varint,1,opt,name=rewards_interval,json=rewardsInterval,proto3" json:"rewards_interval,omitempty"`
	DelegateInterval                  uint64 `protobuf:"varint,6,opt,name=delegate_interval,json=delegateInterval,proto3" json:"delegate_interval,omitempty"`
	DepositInterval                   uint64 `protobuf:"varint,2,opt,name=deposit_interval,json=depositInterval,proto3" json:"deposit_interval,omitempty"`
	RedemptionRateInterval            uint64 `protobuf:"varint,3,opt,name=redemption_rate_interval,json=redemptionRateInterval,proto3" json:"redemption_rate_interval,omitempty"`
	StrideCommission                  uint64 `protobuf:"varint,4,opt,name=stride_commission,json=strideCommission,proto3" json:"stride_commission,omitempty"`
	ReinvestInterval                  uint64 `protobuf:"varint,7,opt,name=reinvest_interval,json=reinvestInterval,proto3" json:"reinvest_interval,omitempty"`
	IcaTimeoutNanos                   uint64 `protobuf:"varint,9,opt,name=ica_timeout_nanos,json=icaTimeoutNanos,proto3" json:"ica_timeout_nanos,omitempty"`
	BufferSize                        uint64 `protobuf:"varint,10,opt,name=buffer_size,json=bufferSize,proto3" json:"buffer_size,omitempty"`
	IbcTimeoutBlocks                  uint64 `protobuf:"varint,11,opt,name=ibc_timeout_blocks,json=ibcTimeoutBlocks,proto3" json:"ibc_timeout_blocks,omitempty"`
	FeeTransferTimeoutNanos           uint64 `protobuf:"varint,12,opt,name=fee_transfer_timeout_nanos,json=feeTransferTimeoutNanos,proto3" json:"fee_transfer_timeout_nanos,omitempty"`
	MaxStakeIcaCallsPerEpoch          uint64 `protobuf:"varint,13,opt,name=max_stake_ica_calls_per_epoch,json=maxStakeIcaCallsPerEpoch,proto3" json:"max_stake_ica_calls_per_epoch,omitempty"`
	DefaultMinRedemptionRateThreshold uint64 `protobuf:"varint,14,opt,name=default_min_redemption_rate_threshold,json=defaultMinRedemptionRateThreshold,proto3" json:"default_min_redemption_rate_threshold,omitempty"`
	DefaultMaxRedemptionRateThreshold uint64 `protobuf:"varint,15,opt,name=default_max_redemption_rate_threshold,json=defaultMaxRedemptionRateThreshold,proto3" json:"default_max_redemption_rate_threshold,omitempty"`
	IbcTransferTimeoutNanos           uint64 `protobuf:"varint,16,opt,name=ibc_transfer_timeout_nanos,json=ibcTransferTimeoutNanos,proto3" json:"ibc_transfer_timeout_nanos,omitempty"`
	SafetyNumValidators               uint64 `protobuf:"varint,17,opt,name=safety_num_validators,json=safetyNumValidators,proto3" json:"safety_num_validators,omitempty"`
	SafetyMaxSlashPercent             uint64 `protobuf:"varint,18,opt,name=safety_max_slash_percent,json=safetyMaxSlashPercent,proto3" json:"safety_max_slash_percent,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_5aeaab6a38c2b438, []int{0}
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

func (m *Params) GetRewardsInterval() uint64 {
	if m != nil {
		return m.RewardsInterval
	}
	return 0
}

func (m *Params) GetDelegateInterval() uint64 {
	if m != nil {
		return m.DelegateInterval
	}
	return 0
}

func (m *Params) GetDepositInterval() uint64 {
	if m != nil {
		return m.DepositInterval
	}
	return 0
}

func (m *Params) GetRedemptionRateInterval() uint64 {
	if m != nil {
		return m.RedemptionRateInterval
	}
	return 0
}

func (m *Params) GetStrideCommission() uint64 {
	if m != nil {
		return m.StrideCommission
	}
	return 0
}

func (m *Params) GetReinvestInterval() uint64 {
	if m != nil {
		return m.ReinvestInterval
	}
	return 0
}

func (m *Params) GetIcaTimeoutNanos() uint64 {
	if m != nil {
		return m.IcaTimeoutNanos
	}
	return 0
}

func (m *Params) GetBufferSize() uint64 {
	if m != nil {
		return m.BufferSize
	}
	return 0
}

func (m *Params) GetIbcTimeoutBlocks() uint64 {
	if m != nil {
		return m.IbcTimeoutBlocks
	}
	return 0
}

func (m *Params) GetFeeTransferTimeoutNanos() uint64 {
	if m != nil {
		return m.FeeTransferTimeoutNanos
	}
	return 0
}

func (m *Params) GetMaxStakeIcaCallsPerEpoch() uint64 {
	if m != nil {
		return m.MaxStakeIcaCallsPerEpoch
	}
	return 0
}

func (m *Params) GetDefaultMinRedemptionRateThreshold() uint64 {
	if m != nil {
		return m.DefaultMinRedemptionRateThreshold
	}
	return 0
}

func (m *Params) GetDefaultMaxRedemptionRateThreshold() uint64 {
	if m != nil {
		return m.DefaultMaxRedemptionRateThreshold
	}
	return 0
}

func (m *Params) GetIbcTransferTimeoutNanos() uint64 {
	if m != nil {
		return m.IbcTransferTimeoutNanos
	}
	return 0
}

func (m *Params) GetSafetyNumValidators() uint64 {
	if m != nil {
		return m.SafetyNumValidators
	}
	return 0
}

func (m *Params) GetSafetyMaxSlashPercent() uint64 {
	if m != nil {
		return m.SafetyMaxSlashPercent
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "stride.stakeibc.Params")
}

func init() { proto.RegisterFile("stride/stakeibc/params.proto", fileDescriptor_5aeaab6a38c2b438) }

var fileDescriptor_5aeaab6a38c2b438 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0x93, 0xef, 0x8b, 0x42, 0x3b, 0x05, 0x9a, 0x98, 0x7f, 0xa3, 0x0a, 0x5c, 0x40, 0x42,
	0xa2, 0x14, 0x62, 0x11, 0x16, 0x20, 0xba, 0x40, 0x6a, 0xc5, 0xa2, 0xa8, 0xad, 0xa2, 0x24, 0x62,
	0xc1, 0x66, 0x34, 0x1e, 0x5f, 0x27, 0xa3, 0xda, 0x1e, 0x6b, 0x66, 0x1c, 0xdc, 0x3e, 0x05, 0x4b,
	0x96, 0x3c, 0x0e, 0xcb, 0x2e, 0x59, 0xa2, 0xe4, 0x0d, 0x78, 0x02, 0xe4, 0x19, 0xc7, 0xa9, 0x51,
	0x61, 0x17, 0x9d, 0xf3, 0xbb, 0xc7, 0x77, 0xce, 0x64, 0xd0, 0x7d, 0xa5, 0x25, 0x0f, 0xc0, 0x53,
	0x9a, 0x9e, 0x02, 0xf7, 0x99, 0x97, 0x52, 0x49, 0x63, 0xd5, 0x4b, 0xa5, 0xd0, 0xc2, 0xd9, 0xb4,
	0x6e, 0x6f, 0xe9, 0x6e, 0xdd, 0x9e, 0x88, 0x89, 0x30, 0x9e, 0x57, 0xfc, 0xb2, 0xd8, 0xe3, 0x5f,
	0x6d, 0xd4, 0x1e, 0x98, 0x39, 0x67, 0x07, 0x75, 0x24, 0x7c, 0xa6, 0x32, 0x50, 0x84, 0x27, 0x1a,
	0xe4, 0x8c, 0x46, 0xb8, 0xf9, 0xb0, 0xf9, 0xb4, 0x35, 0xdc, 0x2c, 0xf5, 0xc3, 0x52, 0x76, 0x76,
	0x51, 0x37, 0x80, 0x08, 0x26, 0x54, 0xc3, 0x8a, 0x6d, 0x1b, 0xb6, 0xb3, 0x34, 0x2a, 0x78, 0x07,
	0x75, 0x02, 0x48, 0x85, 0xe2, 0x7a, 0xc5, 0xfe, 0x67, 0x73, 0x4b, 0xbd, 0x42, 0xdf, 0x20, 0x2c,
	0x21, 0x80, 0x38, 0xd5, 0x5c, 0x24, 0x44, 0xd6, 0xe2, 0xff, 0x37, 0x23, 0x77, 0x57, 0xfe, 0xf0,
	0xf2, 0x47, 0x76, 0x51, 0xd7, 0x1e, 0x98, 0x30, 0x11, 0xc7, 0x5c, 0x29, 0x2e, 0x12, 0xdc, 0xb2,
	0x1b, 0x59, 0xe3, 0xa0, 0xd2, 0x0b, 0x58, 0x02, 0x4f, 0x66, 0xa0, 0x2e, 0xad, 0x74, 0xcd, 0xc2,
	0x4b, 0xa3, 0x4a, 0x7e, 0x86, 0xba, 0x9c, 0x51, 0xa2, 0x79, 0x0c, 0x22, 0xd3, 0x24, 0xa1, 0x89,
	0x50, 0x78, 0xdd, 0xee, 0xcf, 0x19, 0x1d, 0x5b, 0xfd, 0xa4, 0x90, 0x9d, 0x6d, 0xb4, 0xe1, 0x67,
	0x61, 0x08, 0x92, 0x28, 0x7e, 0x0e, 0x18, 0x19, 0x0a, 0x59, 0x69, 0xc4, 0xcf, 0xc1, 0x79, 0x8e,
	0x1c, 0xee, 0xb3, 0x2a, 0xcc, 0x8f, 0x04, 0x3b, 0x55, 0x78, 0xc3, 0x7e, 0x9a, 0xfb, 0xac, 0x4c,
	0xdb, 0x37, 0xba, 0xb3, 0x87, 0xb6, 0x42, 0x00, 0xa2, 0x25, 0x4d, 0x54, 0x11, 0x5a, 0xdf, 0xe1,
	0xba, 0x99, 0xba, 0x17, 0x02, 0x8c, 0x4b, 0xa0, 0xb6, 0xcb, 0x3b, 0xf4, 0x20, 0xa6, 0x39, 0x31,
	0xf7, 0x4f, 0x8a, 0x13, 0x30, 0x1a, 0x45, 0x8a, 0xa4, 0x20, 0x09, 0xa4, 0x82, 0x4d, 0xf1, 0x0d,
	0x33, 0x8f, 0x63, 0x9a, 0x8f, 0x0a, 0xe6, 0x90, 0xd1, 0x83, 0x82, 0x18, 0x80, 0x7c, 0x5f, 0xf8,
	0xce, 0x00, 0x3d, 0x09, 0x20, 0xa4, 0x59, 0xa4, 0x49, 0xcc, 0x13, 0xf2, 0xe7, 0xc5, 0xe8, 0xa9,
	0x04, 0x35, 0x15, 0x51, 0x80, 0x6f, 0x9a, 0xa0, 0x47, 0x25, 0x7c, 0xcc, 0x93, 0x61, 0xed, 0x8e,
	0xc6, 0x4b, 0xb0, 0x96, 0x48, 0xf3, 0x7f, 0x24, 0x6e, 0xd6, 0x13, 0x69, 0xfe, 0xb7, 0xc4, 0x3d,
	0xb4, 0x65, 0xfa, 0xbc, 0xba, 0xa1, 0x8e, 0x6d, 0xa8, 0xe8, 0xf5, 0xaa, 0x86, 0xfa, 0xe8, 0x8e,
	0xa2, 0x21, 0xe8, 0x33, 0x92, 0x64, 0x31, 0x99, 0xd1, 0x88, 0x07, 0x54, 0x0b, 0xa9, 0x70, 0xd7,
	0xcc, 0xdd, 0xb2, 0xe6, 0x49, 0x16, 0x7f, 0xac, 0x2c, 0xe7, 0x35, 0xc2, 0xe5, 0x8c, 0x29, 0x37,
	0xa2, 0x6a, 0x5a, 0x54, 0xca, 0x20, 0xd1, 0xd8, 0x31, 0x63, 0x65, 0xe6, 0x31, 0xcd, 0x47, 0x85,
	0x3b, 0xb0, 0xe6, 0xdb, 0xd6, 0xd7, 0x6f, 0xdb, 0x8d, 0x0f, 0xad, 0xb5, 0xb5, 0xce, 0xfa, 0xfe,
	0xd1, 0xf7, 0xb9, 0xdb, 0xbc, 0x98, 0xbb, 0xcd, 0x9f, 0x73, 0xb7, 0xf9, 0x65, 0xe1, 0x36, 0x2e,
	0x16, 0x6e, 0xe3, 0xc7, 0xc2, 0x6d, 0x7c, 0xea, 0x4f, 0xb8, 0x9e, 0x66, 0x7e, 0x8f, 0x89, 0xd8,
	0x1b, 0x99, 0xbf, 0xed, 0x8b, 0x23, 0xea, 0x2b, 0xaf, 0x7c, 0xea, 0xb3, 0x97, 0x7d, 0x2f, 0x5f,
	0x3d, 0x78, 0x7d, 0x96, 0x82, 0xf2, 0xdb, 0xe6, 0x25, 0xbf, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0x38, 0x1e, 0x64, 0x27, 0x10, 0x04, 0x00, 0x00,
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
	if m.SafetyMaxSlashPercent != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SafetyMaxSlashPercent))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	if m.SafetyNumValidators != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SafetyNumValidators))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.IbcTransferTimeoutNanos != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.IbcTransferTimeoutNanos))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.DefaultMaxRedemptionRateThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DefaultMaxRedemptionRateThreshold))
		i--
		dAtA[i] = 0x78
	}
	if m.DefaultMinRedemptionRateThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DefaultMinRedemptionRateThreshold))
		i--
		dAtA[i] = 0x70
	}
	if m.MaxStakeIcaCallsPerEpoch != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxStakeIcaCallsPerEpoch))
		i--
		dAtA[i] = 0x68
	}
	if m.FeeTransferTimeoutNanos != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.FeeTransferTimeoutNanos))
		i--
		dAtA[i] = 0x60
	}
	if m.IbcTimeoutBlocks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.IbcTimeoutBlocks))
		i--
		dAtA[i] = 0x58
	}
	if m.BufferSize != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BufferSize))
		i--
		dAtA[i] = 0x50
	}
	if m.IcaTimeoutNanos != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.IcaTimeoutNanos))
		i--
		dAtA[i] = 0x48
	}
	if m.ReinvestInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ReinvestInterval))
		i--
		dAtA[i] = 0x38
	}
	if m.DelegateInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DelegateInterval))
		i--
		dAtA[i] = 0x30
	}
	if m.StrideCommission != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.StrideCommission))
		i--
		dAtA[i] = 0x20
	}
	if m.RedemptionRateInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RedemptionRateInterval))
		i--
		dAtA[i] = 0x18
	}
	if m.DepositInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DepositInterval))
		i--
		dAtA[i] = 0x10
	}
	if m.RewardsInterval != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RewardsInterval))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
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
	if m.RewardsInterval != 0 {
		n += 1 + sovParams(uint64(m.RewardsInterval))
	}
	if m.DepositInterval != 0 {
		n += 1 + sovParams(uint64(m.DepositInterval))
	}
	if m.RedemptionRateInterval != 0 {
		n += 1 + sovParams(uint64(m.RedemptionRateInterval))
	}
	if m.StrideCommission != 0 {
		n += 1 + sovParams(uint64(m.StrideCommission))
	}
	if m.DelegateInterval != 0 {
		n += 1 + sovParams(uint64(m.DelegateInterval))
	}
	if m.ReinvestInterval != 0 {
		n += 1 + sovParams(uint64(m.ReinvestInterval))
	}
	if m.IcaTimeoutNanos != 0 {
		n += 1 + sovParams(uint64(m.IcaTimeoutNanos))
	}
	if m.BufferSize != 0 {
		n += 1 + sovParams(uint64(m.BufferSize))
	}
	if m.IbcTimeoutBlocks != 0 {
		n += 1 + sovParams(uint64(m.IbcTimeoutBlocks))
	}
	if m.FeeTransferTimeoutNanos != 0 {
		n += 1 + sovParams(uint64(m.FeeTransferTimeoutNanos))
	}
	if m.MaxStakeIcaCallsPerEpoch != 0 {
		n += 1 + sovParams(uint64(m.MaxStakeIcaCallsPerEpoch))
	}
	if m.DefaultMinRedemptionRateThreshold != 0 {
		n += 1 + sovParams(uint64(m.DefaultMinRedemptionRateThreshold))
	}
	if m.DefaultMaxRedemptionRateThreshold != 0 {
		n += 1 + sovParams(uint64(m.DefaultMaxRedemptionRateThreshold))
	}
	if m.IbcTransferTimeoutNanos != 0 {
		n += 2 + sovParams(uint64(m.IbcTransferTimeoutNanos))
	}
	if m.SafetyNumValidators != 0 {
		n += 2 + sovParams(uint64(m.SafetyNumValidators))
	}
	if m.SafetyMaxSlashPercent != 0 {
		n += 2 + sovParams(uint64(m.SafetyMaxSlashPercent))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsInterval", wireType)
			}
			m.RewardsInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardsInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositInterval", wireType)
			}
			m.DepositInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DepositInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionRateInterval", wireType)
			}
			m.RedemptionRateInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RedemptionRateInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StrideCommission", wireType)
			}
			m.StrideCommission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StrideCommission |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateInterval", wireType)
			}
			m.DelegateInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelegateInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReinvestInterval", wireType)
			}
			m.ReinvestInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReinvestInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IcaTimeoutNanos", wireType)
			}
			m.IcaTimeoutNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IcaTimeoutNanos |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferSize", wireType)
			}
			m.BufferSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BufferSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcTimeoutBlocks", wireType)
			}
			m.IbcTimeoutBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IbcTimeoutBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeTransferTimeoutNanos", wireType)
			}
			m.FeeTransferTimeoutNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FeeTransferTimeoutNanos |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxStakeIcaCallsPerEpoch", wireType)
			}
			m.MaxStakeIcaCallsPerEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxStakeIcaCallsPerEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultMinRedemptionRateThreshold", wireType)
			}
			m.DefaultMinRedemptionRateThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DefaultMinRedemptionRateThreshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultMaxRedemptionRateThreshold", wireType)
			}
			m.DefaultMaxRedemptionRateThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DefaultMaxRedemptionRateThreshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcTransferTimeoutNanos", wireType)
			}
			m.IbcTransferTimeoutNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IbcTransferTimeoutNanos |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafetyNumValidators", wireType)
			}
			m.SafetyNumValidators = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SafetyNumValidators |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafetyMaxSlashPercent", wireType)
			}
			m.SafetyMaxSlashPercent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SafetyMaxSlashPercent |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
