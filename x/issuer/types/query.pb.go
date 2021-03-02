// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: em/issuer/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type QueryIssuersRequest struct {
}

func (m *QueryIssuersRequest) Reset()         { *m = QueryIssuersRequest{} }
func (m *QueryIssuersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIssuersRequest) ProtoMessage()    {}
func (*QueryIssuersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef8f815dfaeaef83, []int{0}
}
func (m *QueryIssuersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIssuersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIssuersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIssuersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIssuersRequest.Merge(m, src)
}
func (m *QueryIssuersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIssuersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIssuersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIssuersRequest proto.InternalMessageInfo

type QueryIssuersResponse struct {
	Issuers []Issuer `protobuf:"bytes,1,rep,name=issuers,proto3" json:"issuers" yaml:"issuers"`
}

func (m *QueryIssuersResponse) Reset()         { *m = QueryIssuersResponse{} }
func (m *QueryIssuersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIssuersResponse) ProtoMessage()    {}
func (*QueryIssuersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef8f815dfaeaef83, []int{1}
}
func (m *QueryIssuersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIssuersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIssuersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIssuersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIssuersResponse.Merge(m, src)
}
func (m *QueryIssuersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIssuersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIssuersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIssuersResponse proto.InternalMessageInfo

func (m *QueryIssuersResponse) GetIssuers() []Issuer {
	if m != nil {
		return m.Issuers
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryIssuersRequest)(nil), "em.issuer.v1beta1.QueryIssuersRequest")
	proto.RegisterType((*QueryIssuersResponse)(nil), "em.issuer.v1beta1.QueryIssuersResponse")
}

func init() { proto.RegisterFile("em/issuer/v1beta1/query.proto", fileDescriptor_ef8f815dfaeaef83) }

var fileDescriptor_ef8f815dfaeaef83 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x93, 0xff, 0x1f, 0x2d, 0x44, 0x10, 0x8c, 0x55, 0xb4, 0xe8, 0x56, 0x73, 0xb0, 0x82,
	0x34, 0x4b, 0xeb, 0xcd, 0x63, 0xc1, 0x83, 0x78, 0xb2, 0x47, 0x6f, 0x9b, 0x38, 0xac, 0x81, 0xee,
	0x4e, 0x9a, 0xd9, 0x88, 0xb9, 0xea, 0x59, 0x10, 0xfc, 0x52, 0x3d, 0x16, 0xbc, 0x78, 0x2a, 0xd2,
	0xfa, 0x09, 0xfc, 0x04, 0xd2, 0x24, 0x2d, 0x48, 0x03, 0xde, 0x96, 0x79, 0xf3, 0x7e, 0xef, 0xed,
	0x38, 0x87, 0xa0, 0x78, 0x44, 0x94, 0x42, 0xc2, 0x1f, 0x3a, 0x01, 0x18, 0xd1, 0xe1, 0xc3, 0x14,
	0x92, 0xcc, 0x8f, 0x13, 0x34, 0xe8, 0x6e, 0x81, 0xf2, 0x0b, 0xd9, 0x2f, 0xe5, 0x46, 0x5d, 0xa2,
	0xc4, 0x5c, 0xe5, 0xf3, 0x57, 0xb1, 0xd8, 0x60, 0x21, 0x92, 0x42, 0xe2, 0x81, 0x20, 0x58, 0x92,
	0x42, 0x8c, 0x74, 0xa9, 0x1f, 0x48, 0x44, 0x39, 0x00, 0x2e, 0xe2, 0x88, 0x0b, 0xad, 0xd1, 0x08,
	0x13, 0xa1, 0xa6, 0x85, 0x7b, 0xb5, 0x45, 0x99, 0x9a, 0xeb, 0xde, 0x8e, 0xb3, 0x7d, 0x33, 0x6f,
	0x75, 0x95, 0x0f, 0xa9, 0x0f, 0xc3, 0x14, 0xc8, 0x78, 0xa1, 0x53, 0xff, 0x3d, 0xa6, 0x18, 0x35,
	0x81, 0x7b, 0xed, 0xd4, 0x0a, 0x3b, 0xed, 0xd9, 0x47, 0xff, 0x4f, 0x37, 0xba, 0xfb, 0xfe, 0xca,
	0x3f, 0xfc, 0xc2, 0xd4, 0xdb, 0x1d, 0x4d, 0x9a, 0xd6, 0xf7, 0xa4, 0xb9, 0x99, 0x09, 0x35, 0xb8,
	0xf0, 0x4a, 0x9f, 0xd7, 0x5f, 0x10, 0xba, 0x2f, 0xb6, 0xb3, 0x96, 0xa7, 0xb8, 0xcf, 0xb6, 0x53,
	0x2b, 0xa3, 0xdc, 0x93, 0x0a, 0x62, 0x45, 0xc5, 0x46, 0xeb, 0xcf, 0xbd, 0xa2, 0xb3, 0xd7, 0x7a,
	0x7a, 0xff, 0x7a, 0xfb, 0x77, 0xec, 0x36, 0x39, 0xb4, 0x15, 0x6a, 0xc8, 0xaa, 0x0f, 0x42, 0xbd,
	0xcb, 0xd1, 0x94, 0xd9, 0xe3, 0x29, 0xb3, 0x3f, 0xa7, 0xcc, 0x7e, 0x9d, 0x31, 0x6b, 0x3c, 0x63,
	0xd6, 0xc7, 0x8c, 0x59, 0xb7, 0x67, 0x32, 0x32, 0xf7, 0x69, 0xe0, 0x87, 0xa8, 0x96, 0x10, 0x50,
	0xed, 0x01, 0xdc, 0x49, 0x48, 0xf8, 0xe3, 0x02, 0x68, 0xb2, 0x18, 0x28, 0x58, 0xcf, 0x2f, 0x7b,
	0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xd0, 0x8e, 0xa9, 0x01, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Issuers(ctx context.Context, in *QueryIssuersRequest, opts ...grpc.CallOption) (*QueryIssuersResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Issuers(ctx context.Context, in *QueryIssuersRequest, opts ...grpc.CallOption) (*QueryIssuersResponse, error) {
	out := new(QueryIssuersResponse)
	err := c.cc.Invoke(ctx, "/em.issuer.v1beta1.Query/Issuers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Issuers(context.Context, *QueryIssuersRequest) (*QueryIssuersResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Issuers(ctx context.Context, req *QueryIssuersRequest) (*QueryIssuersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issuers not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Issuers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIssuersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Issuers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/em.issuer.v1beta1.Query/Issuers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Issuers(ctx, req.(*QueryIssuersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "em.issuer.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Issuers",
			Handler:    _Query_Issuers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "em/issuer/v1beta1/query.proto",
}

func (m *QueryIssuersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIssuersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIssuersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryIssuersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIssuersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIssuersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Issuers) > 0 {
		for iNdEx := len(m.Issuers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Issuers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryIssuersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryIssuersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Issuers) > 0 {
		for _, e := range m.Issuers {
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
func (m *QueryIssuersRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIssuersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIssuersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryIssuersResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIssuersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIssuersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuers", wireType)
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
			m.Issuers = append(m.Issuers, Issuer{})
			if err := m.Issuers[len(m.Issuers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
