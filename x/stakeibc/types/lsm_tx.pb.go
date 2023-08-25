// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/staking/v1beta1/lsm_tx.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// MsgRedeemTokensForShares redeems a tokenized share back into a native
// delegation
type MsgRedeemTokensForShares struct {
	DelegatorAddress string     `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty" yaml:"delegator_address"`
	Amount           types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgRedeemTokensForShares) Reset()         { *m = MsgRedeemTokensForShares{} }
func (m *MsgRedeemTokensForShares) String() string { return proto.CompactTextString(m) }
func (*MsgRedeemTokensForShares) ProtoMessage()    {}
func (*MsgRedeemTokensForShares) Descriptor() ([]byte, []int) {
	return fileDescriptor_34c3b474a863e424, []int{0}
}
func (m *MsgRedeemTokensForShares) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRedeemTokensForShares) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRedeemTokensForShares.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRedeemTokensForShares) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRedeemTokensForShares.Merge(m, src)
}
func (m *MsgRedeemTokensForShares) XXX_Size() int {
	return m.Size()
}
func (m *MsgRedeemTokensForShares) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRedeemTokensForShares.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRedeemTokensForShares proto.InternalMessageInfo

// MsgRedeemTokensForSharesResponse defines the Msg/MsgRedeemTokensForShares
// response type.
type MsgRedeemTokensForSharesResponse struct {
	Amount types.Coin `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgRedeemTokensForSharesResponse) Reset()         { *m = MsgRedeemTokensForSharesResponse{} }
func (m *MsgRedeemTokensForSharesResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRedeemTokensForSharesResponse) ProtoMessage()    {}
func (*MsgRedeemTokensForSharesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_34c3b474a863e424, []int{1}
}
func (m *MsgRedeemTokensForSharesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRedeemTokensForSharesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRedeemTokensForSharesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRedeemTokensForSharesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRedeemTokensForSharesResponse.Merge(m, src)
}
func (m *MsgRedeemTokensForSharesResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRedeemTokensForSharesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRedeemTokensForSharesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRedeemTokensForSharesResponse proto.InternalMessageInfo

func (m *MsgRedeemTokensForSharesResponse) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*MsgRedeemTokensForShares)(nil), "cosmos.staking.v1beta1.MsgRedeemTokensForShares")
	proto.RegisterType((*MsgRedeemTokensForSharesResponse)(nil), "cosmos.staking.v1beta1.MsgRedeemTokensForSharesResponse")
}

func init() {
	proto.RegisterFile("cosmos/staking/v1beta1/lsm_tx.proto", fileDescriptor_34c3b474a863e424)
}

var fileDescriptor_34c3b474a863e424 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x2f, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0xd4, 0xcf, 0x29, 0xce, 0x8d, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x83, 0x28, 0xd2, 0x83, 0x2a, 0xd2, 0x83, 0x2a, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0x2b, 0xd1, 0x07, 0xb1, 0x20, 0xaa, 0xa5, 0xe4, 0xa0, 0x46, 0x26, 0x25, 0x16, 0xa7, 0xc2, 0xcd,
	0x4b, 0xce, 0xcf, 0xcc, 0x83, 0xc8, 0x2b, 0xad, 0x60, 0xe4, 0x92, 0xf0, 0x2d, 0x4e, 0x0f, 0x4a,
	0x4d, 0x49, 0x4d, 0xcd, 0x0d, 0xc9, 0xcf, 0x4e, 0xcd, 0x2b, 0x76, 0xcb, 0x2f, 0x0a, 0xce, 0x48,
	0x2c, 0x4a, 0x2d, 0x16, 0xf2, 0xe4, 0x12, 0x4c, 0x49, 0xcd, 0x49, 0x4d, 0x4f, 0x2c, 0xc9, 0x2f,
	0x8a, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0x92,
	0xf9, 0x74, 0x4f, 0x5e, 0xa2, 0x32, 0x31, 0x37, 0xc7, 0x4a, 0x09, 0x43, 0x89, 0x52, 0x90, 0x00,
	0x5c, 0xcc, 0x11, 0x22, 0x24, 0x64, 0xce, 0xc5, 0x96, 0x98, 0x9b, 0x5f, 0x9a, 0x57, 0x22, 0xc1,
	0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa9, 0x07, 0xf5, 0x06, 0xc8, 0x61, 0x30, 0x3f, 0xe8, 0x39,
	0xe7, 0x67, 0xe6, 0x39, 0xb1, 0x9c, 0xb8, 0x27, 0xcf, 0x10, 0x04, 0x55, 0x6e, 0xc5, 0xd1, 0xb1,
	0x40, 0x9e, 0xe1, 0xc5, 0x02, 0x79, 0x06, 0xa5, 0x68, 0x2e, 0x05, 0x5c, 0x2e, 0x0d, 0x4a, 0x2d,
	0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x45, 0xb2, 0x86, 0x91, 0x24, 0x6b, 0x9c, 0x7c, 0x4e, 0x3c, 0x92,
	0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c,
	0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x28, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f,
	0x39, 0x3f, 0x57, 0x3f, 0xb8, 0xa4, 0x28, 0x33, 0x25, 0x55, 0xd7, 0x27, 0x31, 0x09, 0x14, 0x49,
	0x20, 0xb6, 0x7e, 0x99, 0xa1, 0xb1, 0x7e, 0x05, 0x38, 0xc6, 0x52, 0x33, 0x93, 0x92, 0xf5, 0x4b,
	0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x81, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x00,
	0x6c, 0xa7, 0xb3, 0xd1, 0x01, 0x00, 0x00,
}

func (m *MsgRedeemTokensForShares) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRedeemTokensForShares) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRedeemTokensForShares) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLsmTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintLsmTx(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRedeemTokensForSharesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRedeemTokensForSharesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRedeemTokensForSharesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLsmTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintLsmTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovLsmTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRedeemTokensForShares) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovLsmTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovLsmTx(uint64(l))
	return n
}

func (m *MsgRedeemTokensForSharesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovLsmTx(uint64(l))
	return n
}

func sovLsmTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLsmTx(x uint64) (n int) {
	return sovLsmTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRedeemTokensForShares) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLsmTx
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
			return fmt.Errorf("proto: MsgRedeemTokensForShares: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRedeemTokensForShares: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLsmTx
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
				return ErrInvalidLengthLsmTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLsmTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLsmTx
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
				return ErrInvalidLengthLsmTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLsmTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLsmTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLsmTx
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
func (m *MsgRedeemTokensForSharesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLsmTx
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
			return fmt.Errorf("proto: MsgRedeemTokensForSharesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRedeemTokensForSharesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLsmTx
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
				return ErrInvalidLengthLsmTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLsmTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLsmTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLsmTx
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
func skipLsmTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLsmTx
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
					return 0, ErrIntOverflowLsmTx
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
					return 0, ErrIntOverflowLsmTx
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
				return 0, ErrInvalidLengthLsmTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLsmTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLsmTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLsmTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLsmTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLsmTx = fmt.Errorf("proto: unexpected end of group")
)
