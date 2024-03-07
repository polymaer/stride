// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/records/callbacks.proto

package types

import (
	fmt "fmt"
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

type TransferCallback struct {
	DepositRecordId uint64 `protobuf:"varint,1,opt,name=deposit_record_id,json=depositRecordId,proto3" json:"deposit_record_id,omitempty"`
}

func (m *TransferCallback) Reset()         { *m = TransferCallback{} }
func (m *TransferCallback) String() string { return proto.CompactTextString(m) }
func (*TransferCallback) ProtoMessage()    {}
func (*TransferCallback) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f7cdd5c1d8b3a46, []int{0}
}
func (m *TransferCallback) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransferCallback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransferCallback.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransferCallback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferCallback.Merge(m, src)
}
func (m *TransferCallback) XXX_Size() int {
	return m.Size()
}
func (m *TransferCallback) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferCallback.DiscardUnknown(m)
}

var xxx_messageInfo_TransferCallback proto.InternalMessageInfo

func (m *TransferCallback) GetDepositRecordId() uint64 {
	if m != nil {
		return m.DepositRecordId
	}
	return 0
}

type TransferLSMTokenCallback struct {
	Deposit *LSMTokenDeposit `protobuf:"bytes,1,opt,name=deposit,proto3" json:"deposit,omitempty"`
}

func (m *TransferLSMTokenCallback) Reset()         { *m = TransferLSMTokenCallback{} }
func (m *TransferLSMTokenCallback) String() string { return proto.CompactTextString(m) }
func (*TransferLSMTokenCallback) ProtoMessage()    {}
func (*TransferLSMTokenCallback) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f7cdd5c1d8b3a46, []int{1}
}
func (m *TransferLSMTokenCallback) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransferLSMTokenCallback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransferLSMTokenCallback.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransferLSMTokenCallback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferLSMTokenCallback.Merge(m, src)
}
func (m *TransferLSMTokenCallback) XXX_Size() int {
	return m.Size()
}
func (m *TransferLSMTokenCallback) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferLSMTokenCallback.DiscardUnknown(m)
}

var xxx_messageInfo_TransferLSMTokenCallback proto.InternalMessageInfo

func (m *TransferLSMTokenCallback) GetDeposit() *LSMTokenDeposit {
	if m != nil {
		return m.Deposit
	}
	return nil
}

func init() {
	proto.RegisterType((*TransferCallback)(nil), "stride.records.TransferCallback")
	proto.RegisterType((*TransferLSMTokenCallback)(nil), "stride.records.TransferLSMTokenCallback")
}

func init() { proto.RegisterFile("stride/records/callbacks.proto", fileDescriptor_6f7cdd5c1d8b3a46) }

var fileDescriptor_6f7cdd5c1d8b3a46 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x2e, 0x29, 0xca,
	0x4c, 0x49, 0xd5, 0x2f, 0x4a, 0x4d, 0xce, 0x2f, 0x4a, 0x29, 0xd6, 0x4f, 0x4e, 0xcc, 0xc9, 0x49,
	0x4a, 0x4c, 0xce, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0xc8, 0xeb, 0x41,
	0xe5, 0xa5, 0x64, 0xd0, 0xd4, 0x43, 0x69, 0x88, 0x6a, 0x25, 0x3b, 0x2e, 0x81, 0x90, 0xa2, 0xc4,
	0xbc, 0xe2, 0xb4, 0xd4, 0x22, 0x67, 0xa8, 0x41, 0x42, 0x5a, 0x5c, 0x82, 0x29, 0xa9, 0x05, 0xf9,
	0xc5, 0x99, 0x25, 0xf1, 0x10, 0xc5, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41,
	0xfc, 0x50, 0x89, 0x20, 0xb0, 0xb8, 0x67, 0x8a, 0x52, 0x28, 0x97, 0x04, 0x4c, 0xbf, 0x4f, 0xb0,
	0x6f, 0x48, 0x7e, 0x76, 0x6a, 0x1e, 0xdc, 0x1c, 0x4b, 0x2e, 0x76, 0xa8, 0x72, 0xb0, 0x6e, 0x6e,
	0x23, 0x79, 0x3d, 0x54, 0xb7, 0xe9, 0xc1, 0xb4, 0xb8, 0x40, 0x4d, 0x85, 0xa9, 0x77, 0xf2, 0x3e,
	0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63,
	0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xc3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2,
	0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0x60, 0xb0, 0x69, 0xba, 0x3e, 0x89, 0x49, 0xc5, 0xfa, 0x50,
	0x5f, 0x96, 0x19, 0x5a, 0xea, 0x57, 0xc0, 0xfd, 0x5a, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06,
	0xf6, 0xaa, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xef, 0xc7, 0x18, 0x3a, 0x01, 0x00, 0x00,
}

func (m *TransferCallback) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransferCallback) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransferCallback) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DepositRecordId != 0 {
		i = encodeVarintCallbacks(dAtA, i, uint64(m.DepositRecordId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TransferLSMTokenCallback) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransferLSMTokenCallback) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransferLSMTokenCallback) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Deposit != nil {
		{
			size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCallbacks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCallbacks(dAtA []byte, offset int, v uint64) int {
	offset -= sovCallbacks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransferCallback) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DepositRecordId != 0 {
		n += 1 + sovCallbacks(uint64(m.DepositRecordId))
	}
	return n
}

func (m *TransferLSMTokenCallback) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Deposit != nil {
		l = m.Deposit.Size()
		n += 1 + l + sovCallbacks(uint64(l))
	}
	return n
}

func sovCallbacks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCallbacks(x uint64) (n int) {
	return sovCallbacks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransferCallback) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCallbacks
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
			return fmt.Errorf("proto: TransferCallback: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransferCallback: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositRecordId", wireType)
			}
			m.DepositRecordId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbacks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DepositRecordId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCallbacks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCallbacks
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
func (m *TransferLSMTokenCallback) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCallbacks
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
			return fmt.Errorf("proto: TransferLSMTokenCallback: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransferLSMTokenCallback: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCallbacks
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
				return ErrInvalidLengthCallbacks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCallbacks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Deposit == nil {
				m.Deposit = &LSMTokenDeposit{}
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCallbacks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCallbacks
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
func skipCallbacks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCallbacks
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
					return 0, ErrIntOverflowCallbacks
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
					return 0, ErrIntOverflowCallbacks
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
				return 0, ErrInvalidLengthCallbacks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCallbacks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCallbacks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCallbacks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCallbacks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCallbacks = fmt.Errorf("proto: unexpected end of group")
)
