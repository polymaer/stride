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
// next id: 20
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
	ValidatorSlashQueryThreshold      uint64 `protobuf:"varint,19,opt,name=validator_slash_query_threshold,json=validatorSlashQueryThreshold,proto3" json:"validator_slash_query_threshold,omitempty"`
	ValidatorWeightCap                uint64 `protobuf:"varint,20,opt,name=validator_weight_cap,json=validatorWeightCap,proto3" json:"validator_weight_cap,omitempty"`
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

func (m *Params) GetValidatorSlashQueryThreshold() uint64 {
	if m != nil {
		return m.ValidatorSlashQueryThreshold
	}
	return 0
}

func (m *Params) GetValidatorWeightCap() uint64 {
	if m != nil {
		return m.ValidatorWeightCap
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "stride.stakeibc.Params")
}

func init() { proto.RegisterFile("stride/stakeibc/params.proto", fileDescriptor_5aeaab6a38c2b438) }

var fileDescriptor_5aeaab6a38c2b438 = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x41, 0x6f, 0xd3, 0x3e,
	0x18, 0xc6, 0xdb, 0xff, 0x3f, 0xda, 0x3a, 0x0f, 0x58, 0x9a, 0x4d, 0x10, 0x4d, 0x23, 0x03, 0x24,
	0x24, 0xc6, 0x60, 0x81, 0x71, 0x41, 0xec, 0x80, 0xb4, 0x69, 0x87, 0x4d, 0x03, 0x95, 0xb6, 0x12,
	0x12, 0x17, 0xcb, 0x71, 0xde, 0x36, 0xd6, 0x92, 0x38, 0xd8, 0x6e, 0xd7, 0xed, 0x23, 0x70, 0xe2,
	0xc8, 0x91, 0x8f, 0xc3, 0x71, 0x47, 0x8e, 0xa8, 0xfd, 0x22, 0x28, 0x76, 0x9a, 0x34, 0x68, 0x70,
	0xab, 0x9e, 0xe7, 0xf7, 0x3e, 0x79, 0xf2, 0xba, 0x31, 0xda, 0x92, 0x4a, 0xb0, 0x10, 0x7c, 0xa9,
	0xc8, 0x39, 0xb0, 0x80, 0xfa, 0x19, 0x11, 0x24, 0x91, 0x7b, 0x99, 0xe0, 0x8a, 0x3b, 0x6b, 0xc6,
	0xdd, 0x9b, 0xbb, 0x9b, 0x1b, 0x43, 0x3e, 0xe4, 0xda, 0xf3, 0xf3, 0x5f, 0x06, 0x7b, 0xf4, 0x65,
	0x19, 0x2d, 0x75, 0xf4, 0x9c, 0xb3, 0x83, 0x6c, 0x01, 0x17, 0x44, 0x84, 0x12, 0xb3, 0x54, 0x81,
	0x18, 0x93, 0xd8, 0x6d, 0x3e, 0x68, 0x3e, 0xb1, 0xba, 0x6b, 0x85, 0x7e, 0x52, 0xc8, 0xce, 0x2e,
	0x6a, 0x87, 0x10, 0xc3, 0x90, 0x28, 0xa8, 0xd8, 0x25, 0xcd, 0xda, 0x73, 0xa3, 0x84, 0x77, 0x90,
	0x1d, 0x42, 0xc6, 0x25, 0x53, 0x15, 0xfb, 0x9f, 0xc9, 0x2d, 0xf4, 0x12, 0x7d, 0x8d, 0x5c, 0x01,
	0x21, 0x24, 0x99, 0x62, 0x3c, 0xc5, 0xa2, 0x16, 0xff, 0xbf, 0x1e, 0xb9, 0x5b, 0xf9, 0xdd, 0xc5,
	0x87, 0xec, 0xa2, 0xb6, 0x79, 0x61, 0x4c, 0x79, 0x92, 0x30, 0x29, 0x19, 0x4f, 0x5d, 0xcb, 0x34,
	0x32, 0xc6, 0x51, 0xa9, 0xe7, 0xb0, 0x00, 0x96, 0x8e, 0x41, 0x2e, 0x54, 0x5a, 0x36, 0xf0, 0xdc,
	0x28, 0x93, 0x9f, 0xa2, 0x36, 0xa3, 0x04, 0x2b, 0x96, 0x00, 0x1f, 0x29, 0x9c, 0x92, 0x94, 0x4b,
	0x77, 0xc5, 0xf4, 0x67, 0x94, 0xf4, 0x8d, 0xfe, 0x3e, 0x97, 0x9d, 0x6d, 0xb4, 0x1a, 0x8c, 0x06,
	0x03, 0x10, 0x58, 0xb2, 0x2b, 0x70, 0x91, 0xa6, 0x90, 0x91, 0x7a, 0xec, 0x0a, 0x9c, 0x67, 0xc8,
	0x61, 0x01, 0x2d, 0xc3, 0x82, 0x98, 0xd3, 0x73, 0xe9, 0xae, 0x9a, 0x47, 0xb3, 0x80, 0x16, 0x69,
	0x87, 0x5a, 0x77, 0x0e, 0xd0, 0xe6, 0x00, 0x00, 0x2b, 0x41, 0x52, 0x99, 0x87, 0xd6, 0x3b, 0xdc,
	0xd2, 0x53, 0xf7, 0x06, 0x00, 0xfd, 0x02, 0xa8, 0x75, 0x79, 0x8b, 0xee, 0x27, 0x64, 0x82, 0xf5,
	0xf9, 0xe3, 0xfc, 0x0d, 0x28, 0x89, 0x63, 0x89, 0x33, 0x10, 0x18, 0x32, 0x4e, 0x23, 0xf7, 0xb6,
	0x9e, 0x77, 0x13, 0x32, 0xe9, 0xe5, 0xcc, 0x09, 0x25, 0x47, 0x39, 0xd1, 0x01, 0x71, 0x9c, 0xfb,
	0x4e, 0x07, 0x3d, 0x0e, 0x61, 0x40, 0x46, 0xb1, 0xc2, 0x09, 0x4b, 0xf1, 0x9f, 0x07, 0xa3, 0x22,
	0x01, 0x32, 0xe2, 0x71, 0xe8, 0xde, 0xd1, 0x41, 0x0f, 0x0b, 0xf8, 0x1d, 0x4b, 0xbb, 0xb5, 0x33,
	0xea, 0xcf, 0xc1, 0x5a, 0x22, 0x99, 0xfc, 0x23, 0x71, 0xad, 0x9e, 0x48, 0x26, 0x7f, 0x4b, 0x3c,
	0x40, 0x9b, 0x7a, 0x9f, 0x37, 0x6f, 0xc8, 0x36, 0x1b, 0xca, 0xf7, 0x7a, 0xd3, 0x86, 0x8e, 0xd1,
	0xf6, 0x98, 0xc4, 0x2c, 0x24, 0x8a, 0x0b, 0x2c, 0x63, 0x22, 0x23, 0xfc, 0x79, 0x04, 0xe2, 0x72,
	0xa1, 0xc8, 0xba, 0x4e, 0xd8, 0x2a, 0xb1, 0x5e, 0x4e, 0x7d, 0xc8, 0xa1, 0xaa, 0xc3, 0x0b, 0xb4,
	0x51, 0xc5, 0x5c, 0x00, 0x1b, 0x46, 0x0a, 0x53, 0x92, 0xb9, 0x1b, 0x7a, 0xd6, 0x29, 0xbd, 0x8f,
	0xda, 0x3a, 0x22, 0xd9, 0x1b, 0xeb, 0xdb, 0xf7, 0xed, 0xc6, 0xa9, 0xd5, 0x6a, 0xd9, 0x2b, 0xa7,
	0x56, 0xab, 0x6d, 0x3b, 0xa7, 0x56, 0xcb, 0xb1, 0xd7, 0x0f, 0xcf, 0x7e, 0x4c, 0xbd, 0xe6, 0xf5,
	0xd4, 0x6b, 0xfe, 0x9a, 0x7a, 0xcd, 0xaf, 0x33, 0xaf, 0x71, 0x3d, 0xf3, 0x1a, 0x3f, 0x67, 0x5e,
	0xe3, 0xd3, 0xfe, 0x90, 0xa9, 0x68, 0x14, 0xec, 0x51, 0x9e, 0xf8, 0x3d, 0xfd, 0x77, 0x7e, 0x7e,
	0x46, 0x02, 0xe9, 0x17, 0x57, 0xc0, 0x78, 0xff, 0xa5, 0x3f, 0xa9, 0x2e, 0x02, 0x75, 0x99, 0x81,
	0x0c, 0x96, 0xf4, 0x17, 0xfe, 0xea, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0x1a, 0xb2, 0x1e,
	0x28, 0x04, 0x00, 0x00,
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
	if m.ValidatorWeightCap != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ValidatorWeightCap))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa0
	}
	if m.ValidatorSlashQueryThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ValidatorSlashQueryThreshold))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x98
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
	if m.ValidatorSlashQueryThreshold != 0 {
		n += 2 + sovParams(uint64(m.ValidatorSlashQueryThreshold))
	}
	if m.ValidatorWeightCap != 0 {
		n += 2 + sovParams(uint64(m.ValidatorWeightCap))
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
		case 19:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSlashQueryThreshold", wireType)
			}
			m.ValidatorSlashQueryThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidatorSlashQueryThreshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorWeightCap", wireType)
			}
			m.ValidatorWeightCap = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidatorWeightCap |= uint64(b&0x7F) << shift
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
