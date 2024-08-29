// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/interchainquery/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QueryPendingQueriesRequest struct {
}

func (m *QueryPendingQueriesRequest) Reset()         { *m = QueryPendingQueriesRequest{} }
func (m *QueryPendingQueriesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryPendingQueriesRequest) ProtoMessage()    {}
func (*QueryPendingQueriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b720c147b9144d5b, []int{0}
}
func (m *QueryPendingQueriesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPendingQueriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPendingQueriesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPendingQueriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPendingQueriesRequest.Merge(m, src)
}
func (m *QueryPendingQueriesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryPendingQueriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPendingQueriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPendingQueriesRequest proto.InternalMessageInfo

type QueryPendingQueriesResponse struct {
	PendingQueries []Query `protobuf:"bytes,1,rep,name=pending_queries,json=pendingQueries,proto3" json:"pending_queries"`
}

func (m *QueryPendingQueriesResponse) Reset()         { *m = QueryPendingQueriesResponse{} }
func (m *QueryPendingQueriesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryPendingQueriesResponse) ProtoMessage()    {}
func (*QueryPendingQueriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b720c147b9144d5b, []int{1}
}
func (m *QueryPendingQueriesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPendingQueriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPendingQueriesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPendingQueriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPendingQueriesResponse.Merge(m, src)
}
func (m *QueryPendingQueriesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryPendingQueriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPendingQueriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPendingQueriesResponse proto.InternalMessageInfo

func (m *QueryPendingQueriesResponse) GetPendingQueries() []Query {
	if m != nil {
		return m.PendingQueries
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryPendingQueriesRequest)(nil), "stride.interchainquery.v1.QueryPendingQueriesRequest")
	proto.RegisterType((*QueryPendingQueriesResponse)(nil), "stride.interchainquery.v1.QueryPendingQueriesResponse")
}

func init() {
	proto.RegisterFile("stride/interchainquery/v1/query.proto", fileDescriptor_b720c147b9144d5b)
}

var fileDescriptor_b720c147b9144d5b = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x2e, 0x29, 0xca,
	0x4c, 0x49, 0xd5, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0x2b, 0x2c, 0x4d,
	0x2d, 0xaa, 0xd4, 0x2f, 0x33, 0xd4, 0x07, 0x33, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x24,
	0x21, 0xca, 0xf4, 0xd0, 0x94, 0xe9, 0x95, 0x19, 0x4a, 0xa9, 0xe3, 0x36, 0x21, 0x3d, 0x35, 0x2f,
	0xb5, 0x38, 0xb3, 0x18, 0x62, 0x86, 0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62,
	0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x1e, 0x4c, 0x56, 0x24,
	0x3d, 0x3f, 0x3d, 0x1f, 0xcc, 0xd4, 0x07, 0xb1, 0x20, 0xa2, 0x4a, 0x32, 0x5c, 0x52, 0x81, 0x20,
	0xd3, 0x02, 0x52, 0xf3, 0x52, 0x32, 0xf3, 0xd2, 0x41, 0xec, 0xcc, 0xd4, 0xe2, 0xa0, 0xd4, 0xc2,
	0xd2, 0xd4, 0xe2, 0x12, 0xa5, 0x3c, 0x2e, 0x69, 0xac, 0xb2, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9,
	0x42, 0xfe, 0x5c, 0xfc, 0x05, 0x10, 0x99, 0xf8, 0x42, 0x88, 0x94, 0x04, 0xa3, 0x02, 0xb3, 0x06,
	0xb7, 0x91, 0x82, 0x1e, 0x4e, 0xef, 0xe8, 0x81, 0x0d, 0x74, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21,
	0x88, 0xaf, 0x00, 0xc5, 0x60, 0xa3, 0xb3, 0x8c, 0x5c, 0x3c, 0x60, 0xf9, 0xe0, 0xd4, 0xa2, 0xb2,
	0xcc, 0xe4, 0x54, 0xa1, 0x3d, 0x8c, 0x5c, 0x7c, 0xa8, 0x96, 0x0b, 0x99, 0x12, 0x32, 0x1b, 0xab,
	0x57, 0xa4, 0xcc, 0x48, 0xd5, 0x06, 0xf1, 0xa3, 0x92, 0x75, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0x4c,
	0x85, 0x8c, 0xf5, 0x83, 0xc1, 0xfa, 0x75, 0x7d, 0x12, 0x93, 0x8a, 0xf5, 0x71, 0x44, 0x09, 0x5a,
	0x68, 0x38, 0x05, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x65, 0x7a,
	0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0x2e, 0x36, 0x83, 0xcb, 0x8c, 0x4c, 0xf4, 0x2b,
	0x30, 0x8c, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xc7, 0x9c, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x82, 0xbf, 0x49, 0xf2, 0x5a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	PendingQueries(ctx context.Context, in *QueryPendingQueriesRequest, opts ...grpc.CallOption) (*QueryPendingQueriesResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) PendingQueries(ctx context.Context, in *QueryPendingQueriesRequest, opts ...grpc.CallOption) (*QueryPendingQueriesResponse, error) {
	out := new(QueryPendingQueriesResponse)
	err := c.cc.Invoke(ctx, "/stride.interchainquery.v1.QueryService/PendingQueries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	PendingQueries(context.Context, *QueryPendingQueriesRequest) (*QueryPendingQueriesResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) PendingQueries(ctx context.Context, req *QueryPendingQueriesRequest) (*QueryPendingQueriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PendingQueries not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_PendingQueries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPendingQueriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).PendingQueries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stride.interchainquery.v1.QueryService/PendingQueries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).PendingQueries(ctx, req.(*QueryPendingQueriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stride.interchainquery.v1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PendingQueries",
			Handler:    _QueryService_PendingQueries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stride/interchainquery/v1/query.proto",
}

func (m *QueryPendingQueriesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPendingQueriesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPendingQueriesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryPendingQueriesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPendingQueriesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPendingQueriesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PendingQueries) > 0 {
		for iNdEx := len(m.PendingQueries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PendingQueries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryPendingQueriesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryPendingQueriesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PendingQueries) > 0 {
		for _, e := range m.PendingQueries {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryPendingQueriesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryPendingQueriesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPendingQueriesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryPendingQueriesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryPendingQueriesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPendingQueriesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingQueries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PendingQueries = append(m.PendingQueries, Query{})
			if err := m.PendingQueries[len(m.PendingQueries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
