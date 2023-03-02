// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/stakeibc/validator.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Validator_ValidatorStatus int32

const (
	Validator_ACTIVE   Validator_ValidatorStatus = 0
	Validator_INACTIVE Validator_ValidatorStatus = 1
)

var Validator_ValidatorStatus_name = map[int32]string{
	0: "ACTIVE",
	1: "INACTIVE",
}

var Validator_ValidatorStatus_value = map[string]int32{
	"ACTIVE":   0,
	"INACTIVE": 1,
}

func (x Validator_ValidatorStatus) String() string {
	return proto.EnumName(Validator_ValidatorStatus_name, int32(x))
}

func (Validator_ValidatorStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5d2f32e16bd6ab8f, []int{1, 0}
}

type ValidatorExchangeRate struct {
	InternalTokensToSharesRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=internal_tokens_to_shares_rate,json=internalTokensToSharesRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"internal_tokens_to_shares_rate"`
	EpochNumber                uint64                                 `protobuf:"varint,2,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
}

func (m *ValidatorExchangeRate) Reset()         { *m = ValidatorExchangeRate{} }
func (m *ValidatorExchangeRate) String() string { return proto.CompactTextString(m) }
func (*ValidatorExchangeRate) ProtoMessage()    {}
func (*ValidatorExchangeRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d2f32e16bd6ab8f, []int{0}
}
func (m *ValidatorExchangeRate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorExchangeRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorExchangeRate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorExchangeRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorExchangeRate.Merge(m, src)
}
func (m *ValidatorExchangeRate) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorExchangeRate) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorExchangeRate.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorExchangeRate proto.InternalMessageInfo

func (m *ValidatorExchangeRate) GetEpochNumber() uint64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

type Validator struct {
	Name                 string                                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              string                                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Status               Validator_ValidatorStatus              `protobuf:"varint,3,opt,name=status,proto3,enum=stride.stakeibc.Validator_ValidatorStatus" json:"status,omitempty"`
	CommissionRate       uint64                                 `protobuf:"varint,4,opt,name=commission_rate,json=commissionRate,proto3" json:"commission_rate,omitempty"`
	DelegationAmt        github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=delegation_amt,json=delegationAmt,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"delegation_amt"`
	Weight               uint64                                 `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	InternalExchangeRate *ValidatorExchangeRate                 `protobuf:"bytes,7,opt,name=internal_exchange_rate,json=internalExchangeRate,proto3" json:"internal_exchange_rate,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d2f32e16bd6ab8f, []int{1}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Validator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Validator) GetStatus() Validator_ValidatorStatus {
	if m != nil {
		return m.Status
	}
	return Validator_ACTIVE
}

func (m *Validator) GetCommissionRate() uint64 {
	if m != nil {
		return m.CommissionRate
	}
	return 0
}

func (m *Validator) GetWeight() uint64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Validator) GetInternalExchangeRate() *ValidatorExchangeRate {
	if m != nil {
		return m.InternalExchangeRate
	}
	return nil
}

func init() {
	proto.RegisterEnum("stride.stakeibc.Validator_ValidatorStatus", Validator_ValidatorStatus_name, Validator_ValidatorStatus_value)
	proto.RegisterType((*ValidatorExchangeRate)(nil), "stride.stakeibc.ValidatorExchangeRate")
	proto.RegisterType((*Validator)(nil), "stride.stakeibc.Validator")
}

func init() { proto.RegisterFile("stride/stakeibc/validator.proto", fileDescriptor_5d2f32e16bd6ab8f) }

var fileDescriptor_5d2f32e16bd6ab8f = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xf6, 0xd1, 0xe0, 0x92, 0x6b, 0x49, 0xd0, 0x29, 0x54, 0x26, 0x83, 0x13, 0x32, 0x94, 0x08,
	0x14, 0x5b, 0x04, 0x89, 0x89, 0x25, 0xa1, 0x1d, 0x22, 0x50, 0x07, 0x27, 0x74, 0x40, 0x48, 0xd6,
	0xd9, 0x3e, 0xd9, 0x56, 0x62, 0x5f, 0x74, 0xf7, 0x52, 0xca, 0xc6, 0x4f, 0xe0, 0xc7, 0x74, 0x42,
	0x62, 0xef, 0x58, 0x75, 0x42, 0x0c, 0x15, 0x4a, 0xfe, 0x08, 0xea, 0xf9, 0x9c, 0x44, 0x95, 0x18,
	0x3a, 0xf9, 0xbd, 0xcf, 0xef, 0xbd, 0xef, 0x7b, 0xdf, 0xdd, 0xe1, 0x96, 0x04, 0x91, 0x46, 0xcc,
	0x95, 0x40, 0xa7, 0x2c, 0x0d, 0x42, 0xf7, 0x8c, 0xce, 0xd2, 0x88, 0x02, 0x17, 0xce, 0x5c, 0x70,
	0xe0, 0xa4, 0x5e, 0x14, 0x38, 0x65, 0x41, 0xf3, 0x59, 0xc8, 0x65, 0xc6, 0xa5, 0xaf, 0x7e, 0xbb,
	0x45, 0x52, 0xd4, 0x36, 0x1b, 0x31, 0x8f, 0x79, 0x81, 0xdf, 0x46, 0x05, 0xda, 0xf9, 0x85, 0xf0,
	0xd3, 0xd3, 0x72, 0xea, 0xf1, 0x79, 0x98, 0xd0, 0x3c, 0x66, 0x1e, 0x05, 0x46, 0xbe, 0x23, 0x6c,
	0xa7, 0x39, 0x30, 0x91, 0xd3, 0x99, 0x0f, 0x7c, 0xca, 0x72, 0xe9, 0x03, 0xf7, 0x65, 0x42, 0x05,
	0x93, 0xbe, 0xa0, 0xc0, 0x2c, 0xd4, 0x46, 0xdd, 0xea, 0xf0, 0xdd, 0xe5, 0x4d, 0xcb, 0xf8, 0x73,
	0xd3, 0x3a, 0x8c, 0x53, 0x48, 0x16, 0x81, 0x13, 0xf2, 0x4c, 0x33, 0xeb, 0x4f, 0x4f, 0x46, 0x53,
	0x17, 0xbe, 0xcd, 0x99, 0x74, 0x8e, 0x58, 0x78, 0x7d, 0xd1, 0xc3, 0x5a, 0xd8, 0x11, 0x0b, 0xbd,
	0x66, 0xc9, 0x31, 0x51, 0x14, 0x13, 0x3e, 0x56, 0x04, 0x4a, 0xc2, 0x73, 0xbc, 0xcf, 0xe6, 0x3c,
	0x4c, 0xfc, 0x7c, 0x91, 0x05, 0x4c, 0x58, 0x0f, 0xda, 0xa8, 0x5b, 0xf1, 0xf6, 0x14, 0x76, 0xa2,
	0xa0, 0xce, 0xcf, 0x1d, 0x5c, 0x5d, 0xeb, 0x27, 0x04, 0x57, 0x72, 0x9a, 0x69, 0x61, 0x9e, 0x8a,
	0x49, 0x1f, 0xef, 0xd2, 0x28, 0x12, 0x4c, 0x4a, 0xd5, 0x5f, 0x1d, 0x5a, 0xd7, 0x17, 0xbd, 0x86,
	0x56, 0x30, 0x28, 0xfe, 0x8c, 0x41, 0xa4, 0x79, 0xec, 0x95, 0x85, 0x64, 0x88, 0x4d, 0x09, 0x14,
	0x16, 0xd2, 0xda, 0x69, 0xa3, 0x6e, 0xad, 0xff, 0xd2, 0xb9, 0x63, 0xb4, 0xb3, 0xe6, 0xdc, 0x44,
	0x63, 0xd5, 0xe1, 0xe9, 0x4e, 0xf2, 0x02, 0xd7, 0x43, 0x9e, 0x65, 0xa9, 0x94, 0x29, 0xcf, 0x0b,
	0xbf, 0x2a, 0x4a, 0x7f, 0x6d, 0x03, 0xab, 0x2d, 0x3f, 0xe1, 0x5a, 0xc4, 0x66, 0x2c, 0xa6, 0x70,
	0x5b, 0x48, 0x33, 0xb0, 0x1e, 0x2a, 0x9d, 0xce, 0x3d, 0x7c, 0x1d, 0xe5, 0xe0, 0x3d, 0xde, 0x4c,
	0x19, 0x64, 0x40, 0x0e, 0xb0, 0xf9, 0x95, 0xa5, 0x71, 0x02, 0x96, 0xa9, 0x68, 0x75, 0x46, 0xbe,
	0xe0, 0x83, 0xf5, 0xb1, 0x32, 0x7d, 0xe0, 0x85, 0xbc, 0xdd, 0x36, 0xea, 0xee, 0xf5, 0x0f, 0xff,
	0xbf, 0xeb, 0xf6, 0xfd, 0xf0, 0x1a, 0xe5, 0x94, 0x6d, 0xb4, 0xf3, 0x0a, 0xd7, 0xef, 0x18, 0x42,
	0x30, 0x36, 0x07, 0xef, 0x27, 0xa3, 0xd3, 0xe3, 0x27, 0x06, 0xd9, 0xc7, 0x8f, 0x46, 0x27, 0x3a,
	0x43, 0xc3, 0x0f, 0x97, 0x4b, 0x1b, 0x5d, 0x2d, 0x6d, 0xf4, 0x77, 0x69, 0xa3, 0x1f, 0x2b, 0xdb,
	0xb8, 0x5a, 0xd9, 0xc6, 0xef, 0x95, 0x6d, 0x7c, 0x7e, 0xbd, 0xb5, 0xf3, 0x58, 0xc9, 0xe9, 0x7d,
	0xa4, 0x81, 0x74, 0xf5, 0x83, 0x38, 0x7b, 0xeb, 0x9e, 0x6f, 0x5e, 0x85, 0xb2, 0x20, 0x30, 0xd5,
	0x85, 0x7e, 0xf3, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x2f, 0xfa, 0xeb, 0x35, 0x03, 0x00, 0x00,
}

func (m *ValidatorExchangeRate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorExchangeRate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorExchangeRate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EpochNumber != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.InternalTokensToSharesRate.Size()
		i -= size
		if _, err := m.InternalTokensToSharesRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InternalExchangeRate != nil {
		{
			size, err := m.InternalExchangeRate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Weight != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.Weight))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.DelegationAmt.Size()
		i -= size
		if _, err := m.DelegationAmt.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.CommissionRate != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.CommissionRate))
		i--
		dAtA[i] = 0x20
	}
	if m.Status != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorExchangeRate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.InternalTokensToSharesRate.Size()
	n += 1 + l + sovValidator(uint64(l))
	if m.EpochNumber != 0 {
		n += 1 + sovValidator(uint64(m.EpochNumber))
	}
	return n
}

func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovValidator(uint64(m.Status))
	}
	if m.CommissionRate != 0 {
		n += 1 + sovValidator(uint64(m.CommissionRate))
	}
	l = m.DelegationAmt.Size()
	n += 1 + l + sovValidator(uint64(l))
	if m.Weight != 0 {
		n += 1 + sovValidator(uint64(m.Weight))
	}
	if m.InternalExchangeRate != nil {
		l = m.InternalExchangeRate.Size()
		n += 1 + l + sovValidator(uint64(l))
	}
	return n
}

func sovValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidator(x uint64) (n int) {
	return sovValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorExchangeRate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: ValidatorExchangeRate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorExchangeRate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalTokensToSharesRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InternalTokensToSharesRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Validator_ValidatorStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionRate", wireType)
			}
			m.CommissionRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommissionRate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationAmt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DelegationAmt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Weight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalExchangeRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InternalExchangeRate == nil {
				m.InternalExchangeRate = &ValidatorExchangeRate{}
			}
			if err := m.InternalExchangeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func skipValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
				return 0, ErrInvalidLengthValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidator = fmt.Errorf("proto: unexpected end of group")
)
