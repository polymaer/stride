// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/stakeibc/genesis.proto

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

// GenesisState defines the stakeibc module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId string `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	// list of zones that are registered by the protocol
	HostZoneList     []HostZone     `protobuf:"bytes,5,rep,name=host_zone_list,json=hostZoneList,proto3" json:"host_zone_list"`
	EpochTrackerList []EpochTracker `protobuf:"bytes,10,rep,name=epoch_tracker_list,json=epochTrackerList,proto3" json:"epoch_tracker_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea81129ed6fb77a, []int{0}
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

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetHostZoneList() []HostZone {
	if m != nil {
		return m.HostZoneList
	}
	return nil
}

func (m *GenesisState) GetEpochTrackerList() []EpochTracker {
	if m != nil {
		return m.EpochTrackerList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "stride.stakeibc.GenesisState")
}

func init() { proto.RegisterFile("stride/stakeibc/genesis.proto", fileDescriptor_dea81129ed6fb77a) }

var fileDescriptor_dea81129ed6fb77a = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0x5b, 0x18, 0x4a, 0x19, 0x88, 0x36, 0x8d, 0x09, 0x48, 0xa4, 0x10, 0xdd, 0xb0, 0xb1,
	0x35, 0x35, 0xbe, 0x00, 0x09, 0x51, 0x1b, 0x16, 0x0a, 0xae, 0xd8, 0x34, 0x6d, 0x99, 0xb4, 0x0d,
	0xc2, 0x34, 0x9d, 0xab, 0x51, 0x9f, 0xc2, 0x95, 0xcf, 0xc4, 0x92, 0xa5, 0x2b, 0x63, 0xe0, 0x45,
	0x4c, 0x3b, 0xe3, 0x5f, 0xd9, 0xcd, 0x9d, 0xf3, 0xe5, 0xbb, 0x27, 0x17, 0x77, 0x18, 0xa4, 0xf1,
	0x8c, 0x58, 0x0c, 0xbc, 0x39, 0x89, 0xfd, 0xc0, 0x0a, 0xc9, 0x92, 0xb0, 0x98, 0x99, 0x49, 0x4a,
	0x81, 0xea, 0xfb, 0x3c, 0x36, 0xbf, 0xe3, 0xf6, 0x41, 0x48, 0x43, 0x9a, 0x67, 0x56, 0xf6, 0xe2,
	0x58, 0xfb, 0xa8, 0x68, 0x49, 0xbc, 0xd4, 0x5b, 0x08, 0x49, 0xbb, 0x5b, 0x4c, 0x23, 0xca, 0xc0,
	0x7d, 0xa1, 0x4b, 0x22, 0x80, 0x93, 0x22, 0x40, 0x12, 0x1a, 0x44, 0x2e, 0xa4, 0x5e, 0x30, 0x27,
	0x29, 0x87, 0x8e, 0xdf, 0x4a, 0xb8, 0x71, 0xc9, 0xcb, 0x4d, 0xc0, 0x03, 0xa2, 0x5f, 0x60, 0x85,
	0xaf, 0x69, 0xc9, 0x3d, 0xb9, 0x5f, 0xb7, 0x9b, 0x66, 0xa1, 0xac, 0x79, 0x93, 0xc7, 0x03, 0xb4,
	0xfa, 0xe8, 0x4a, 0x63, 0x01, 0xeb, 0x4d, 0x5c, 0x4d, 0x68, 0x0a, 0x6e, 0x3c, 0x6b, 0x95, 0x7a,
	0x72, 0xbf, 0x36, 0x56, 0xb2, 0xf1, 0x7a, 0xa6, 0x0f, 0xf1, 0xde, 0x4f, 0x31, 0xf7, 0x3e, 0x66,
	0xd0, 0xaa, 0xf4, 0xca, 0xfd, 0xba, 0x7d, 0xb8, 0xe3, 0xbd, 0xa2, 0x0c, 0xa6, 0x74, 0x49, 0x84,
	0xb9, 0x11, 0x89, 0x79, 0x14, 0x33, 0xd0, 0x6f, 0xb1, 0xfe, 0xaf, 0x3e, 0x57, 0xe1, 0x5c, 0xd5,
	0xd9, 0x51, 0x0d, 0x33, 0xf4, 0x8e, 0x93, 0x42, 0xa7, 0x91, 0x3f, 0x7f, 0x99, 0xd2, 0x41, 0x6a,
	0x59, 0x43, 0x0e, 0x52, 0x91, 0x56, 0x71, 0x90, 0xaa, 0x68, 0x55, 0x07, 0xa9, 0x35, 0x0d, 0x3b,
	0x48, 0xad, 0x6b, 0x8d, 0xc1, 0x68, 0xb5, 0x31, 0xe4, 0xf5, 0xc6, 0x90, 0x3f, 0x37, 0x86, 0xfc,
	0xba, 0x35, 0xa4, 0xf5, 0xd6, 0x90, 0xde, 0xb7, 0x86, 0x34, 0xb5, 0xc3, 0x18, 0xa2, 0x07, 0xdf,
	0x0c, 0xe8, 0xc2, 0x9a, 0xe4, 0x8b, 0x4f, 0x47, 0x9e, 0xcf, 0x2c, 0x71, 0xee, 0x47, 0xfb, 0xcc,
	0x7a, 0xfa, 0x3d, 0x3a, 0x3c, 0x27, 0x84, 0xf9, 0x4a, 0x7e, 0xed, 0xf3, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa5, 0x67, 0x55, 0x8a, 0x19, 0x02, 0x00, 0x00,
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
	if len(m.EpochTrackerList) > 0 {
		for iNdEx := len(m.EpochTrackerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EpochTrackerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.HostZoneList) > 0 {
		for iNdEx := len(m.HostZoneList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HostZoneList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.HostZoneList) > 0 {
		for _, e := range m.HostZoneList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.EpochTrackerList) > 0 {
		for _, e := range m.EpochTrackerList {
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
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
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
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostZoneList", wireType)
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
			m.HostZoneList = append(m.HostZoneList, HostZone{})
			if err := m.HostZoneList[len(m.HostZoneList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochTrackerList", wireType)
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
			m.EpochTrackerList = append(m.EpochTrackerList, EpochTracker{})
			if err := m.EpochTrackerList[len(m.EpochTrackerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
