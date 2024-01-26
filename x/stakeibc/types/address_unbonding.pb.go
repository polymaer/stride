// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/stakeibc/address_unbonding.proto

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

type AddressUnbonding struct {
	Address                string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Receiver               string                                 `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	UnbondingEstimatedTime string                                 `protobuf:"bytes,3,opt,name=unbonding_estimated_time,json=unbondingEstimatedTime,proto3" json:"unbonding_estimated_time,omitempty"`
	Amount                 github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	Denom                  string                                 `protobuf:"bytes,5,opt,name=denom,proto3" json:"denom,omitempty"`
	ClaimIsPending         bool                                   `protobuf:"varint,8,opt,name=claim_is_pending,json=claimIsPending,proto3" json:"claim_is_pending,omitempty"`
	EpochNumber            uint64                                 `protobuf:"varint,9,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
}

func (m *AddressUnbonding) Reset()         { *m = AddressUnbonding{} }
func (m *AddressUnbonding) String() string { return proto.CompactTextString(m) }
func (*AddressUnbonding) ProtoMessage()    {}
func (*AddressUnbonding) Descriptor() ([]byte, []int) {
	return fileDescriptor_6491082275bcb473, []int{0}
}
func (m *AddressUnbonding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddressUnbonding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddressUnbonding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddressUnbonding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressUnbonding.Merge(m, src)
}
func (m *AddressUnbonding) XXX_Size() int {
	return m.Size()
}
func (m *AddressUnbonding) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressUnbonding.DiscardUnknown(m)
}

var xxx_messageInfo_AddressUnbonding proto.InternalMessageInfo

func (m *AddressUnbonding) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddressUnbonding) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *AddressUnbonding) GetUnbondingEstimatedTime() string {
	if m != nil {
		return m.UnbondingEstimatedTime
	}
	return ""
}

func (m *AddressUnbonding) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *AddressUnbonding) GetClaimIsPending() bool {
	if m != nil {
		return m.ClaimIsPending
	}
	return false
}

func (m *AddressUnbonding) GetEpochNumber() uint64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*AddressUnbonding)(nil), "stride.stakeibc.AddressUnbonding")
}

func init() {
	proto.RegisterFile("stride/stakeibc/address_unbonding.proto", fileDescriptor_6491082275bcb473)
}

var fileDescriptor_6491082275bcb473 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0xcd, 0x4e, 0xea, 0x40,
	0x18, 0x6d, 0xb9, 0xc0, 0x85, 0xb9, 0x37, 0x4a, 0x26, 0xc4, 0x4c, 0x58, 0x14, 0x74, 0xa1, 0xdd,
	0xd0, 0x46, 0xdd, 0xb0, 0x95, 0x44, 0x13, 0x12, 0x62, 0x4c, 0xd5, 0x8d, 0x9b, 0xa6, 0x3f, 0x5f,
	0xca, 0x04, 0x67, 0xa6, 0xe9, 0x4c, 0x89, 0xbe, 0x85, 0x2f, 0xe2, 0x7b, 0xb0, 0x64, 0x69, 0x5c,
	0x10, 0x03, 0x2f, 0x62, 0x98, 0xb6, 0xe8, 0x6a, 0xe6, 0x3b, 0xe7, 0xcc, 0xf9, 0x4e, 0xe6, 0xa0,
	0x33, 0xa9, 0x32, 0x1a, 0x83, 0x2b, 0x55, 0x30, 0x07, 0x1a, 0x46, 0x6e, 0x10, 0xc7, 0x19, 0x48,
	0xe9, 0xe7, 0x3c, 0x14, 0x3c, 0xa6, 0x3c, 0x71, 0xd2, 0x4c, 0x28, 0x81, 0x0f, 0x0b, 0xa1, 0x53,
	0x09, 0x7b, 0xdd, 0x44, 0x24, 0x42, 0x73, 0xee, 0xee, 0x56, 0xc8, 0x4e, 0xde, 0x6b, 0xa8, 0x73,
	0x55, 0x58, 0x3c, 0x56, 0x0e, 0x98, 0xa0, 0xbf, 0xa5, 0x2d, 0x31, 0x07, 0xa6, 0xdd, 0xf6, 0xaa,
	0x11, 0xf7, 0x50, 0x2b, 0x83, 0x08, 0xe8, 0x02, 0x32, 0x52, 0xd3, 0xd4, 0x7e, 0xc6, 0x23, 0x44,
	0xf6, 0x21, 0x7c, 0x90, 0x8a, 0xb2, 0x40, 0x41, 0xec, 0x2b, 0xca, 0x80, 0xfc, 0xd1, 0xda, 0xa3,
	0x3d, 0x7f, 0x5d, 0xd1, 0x0f, 0x94, 0x01, 0xbe, 0x41, 0xcd, 0x80, 0x89, 0x9c, 0x2b, 0x52, 0xdf,
	0xe9, 0xc6, 0xce, 0x72, 0xdd, 0x37, 0x3e, 0xd7, 0xfd, 0xd3, 0x84, 0xaa, 0x59, 0x1e, 0x3a, 0x91,
	0x60, 0x6e, 0x24, 0x24, 0x13, 0xb2, 0x3c, 0x86, 0x32, 0x9e, 0xbb, 0xea, 0x35, 0x05, 0xe9, 0x4c,
	0xb8, 0xf2, 0xca, 0xd7, 0xb8, 0x8b, 0x1a, 0x31, 0x70, 0xc1, 0x48, 0x43, 0xaf, 0x2b, 0x06, 0x6c,
	0xa3, 0x4e, 0xf4, 0x1c, 0x50, 0xe6, 0x53, 0xe9, 0xa7, 0xa0, 0xd7, 0x93, 0xd6, 0xc0, 0xb4, 0x5b,
	0xde, 0x81, 0xc6, 0x27, 0xf2, 0xae, 0x40, 0xf1, 0x31, 0xfa, 0x0f, 0xa9, 0x88, 0x66, 0x3e, 0xcf,
	0x59, 0x08, 0x19, 0x69, 0x0f, 0x4c, 0xbb, 0xee, 0xfd, 0xd3, 0xd8, 0xad, 0x86, 0xc6, 0xd3, 0xe5,
	0xc6, 0x32, 0x57, 0x1b, 0xcb, 0xfc, 0xda, 0x58, 0xe6, 0xdb, 0xd6, 0x32, 0x56, 0x5b, 0xcb, 0xf8,
	0xd8, 0x5a, 0xc6, 0xd3, 0xc5, 0xaf, 0xb0, 0xf7, 0xfa, 0xef, 0x87, 0xd3, 0x20, 0x94, 0x6e, 0x59,
	0xd8, 0xe2, 0x7c, 0xe4, 0xbe, 0xfc, 0xd4, 0xa6, 0xc3, 0x87, 0x4d, 0x5d, 0xc2, 0xe5, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x62, 0xf7, 0xee, 0x59, 0xd6, 0x01, 0x00, 0x00,
}

func (m *AddressUnbonding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddressUnbonding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddressUnbonding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EpochNumber != 0 {
		i = encodeVarintAddressUnbonding(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x48
	}
	if m.ClaimIsPending {
		i--
		if m.ClaimIsPending {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintAddressUnbonding(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAddressUnbonding(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.UnbondingEstimatedTime) > 0 {
		i -= len(m.UnbondingEstimatedTime)
		copy(dAtA[i:], m.UnbondingEstimatedTime)
		i = encodeVarintAddressUnbonding(dAtA, i, uint64(len(m.UnbondingEstimatedTime)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintAddressUnbonding(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAddressUnbonding(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAddressUnbonding(dAtA []byte, offset int, v uint64) int {
	offset -= sovAddressUnbonding(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AddressUnbonding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAddressUnbonding(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovAddressUnbonding(uint64(l))
	}
	l = len(m.UnbondingEstimatedTime)
	if l > 0 {
		n += 1 + l + sovAddressUnbonding(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovAddressUnbonding(uint64(l))
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovAddressUnbonding(uint64(l))
	}
	if m.ClaimIsPending {
		n += 2
	}
	if m.EpochNumber != 0 {
		n += 1 + sovAddressUnbonding(uint64(m.EpochNumber))
	}
	return n
}

func sovAddressUnbonding(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAddressUnbonding(x uint64) (n int) {
	return sovAddressUnbonding(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddressUnbonding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddressUnbonding
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
			return fmt.Errorf("proto: AddressUnbonding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddressUnbonding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddressUnbonding
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
				return ErrInvalidLengthAddressUnbonding
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddressUnbonding
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddressUnbonding
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
				return ErrInvalidLengthAddressUnbonding
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddressUnbonding
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingEstimatedTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddressUnbonding
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
				return ErrInvalidLengthAddressUnbonding
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddressUnbonding
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnbondingEstimatedTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddressUnbonding
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
				return ErrInvalidLengthAddressUnbonding
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddressUnbonding
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddressUnbonding
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
				return ErrInvalidLengthAddressUnbonding
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddressUnbonding
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimIsPending", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddressUnbonding
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ClaimIsPending = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddressUnbonding
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
			skippy, err := skipAddressUnbonding(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAddressUnbonding
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
func skipAddressUnbonding(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAddressUnbonding
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
					return 0, ErrIntOverflowAddressUnbonding
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
					return 0, ErrIntOverflowAddressUnbonding
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
				return 0, ErrInvalidLengthAddressUnbonding
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAddressUnbonding
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAddressUnbonding
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAddressUnbonding        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAddressUnbonding          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAddressUnbonding = fmt.Errorf("proto: unexpected end of group")
)
