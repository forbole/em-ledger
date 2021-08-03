// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: em/authority/v1/authority.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Authority struct {
	Address       string    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	FormerAddress string    `protobuf:"bytes,2,opt,name=former_address,json=formerAddress,proto3" json:"former_address,omitempty" yaml:"former_address"`
	LastModified  time.Time `protobuf:"bytes,3,opt,name=last_modified,json=lastModified,proto3,stdtime" json:"last_modified" yaml:"last_modified"`
}

func (m *Authority) Reset()         { *m = Authority{} }
func (m *Authority) String() string { return proto.CompactTextString(m) }
func (*Authority) ProtoMessage()    {}
func (*Authority) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f91f8bbecb83881, []int{0}
}
func (m *Authority) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Authority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Authority.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Authority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authority.Merge(m, src)
}
func (m *Authority) XXX_Size() int {
	return m.Size()
}
func (m *Authority) XXX_DiscardUnknown() {
	xxx_messageInfo_Authority.DiscardUnknown(m)
}

var xxx_messageInfo_Authority proto.InternalMessageInfo

func (m *Authority) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Authority) GetFormerAddress() string {
	if m != nil {
		return m.FormerAddress
	}
	return ""
}

func (m *Authority) GetLastModified() time.Time {
	if m != nil {
		return m.LastModified
	}
	return time.Time{}
}

type GasPrices struct {
	Minimum github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=minimum,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"minimum" yaml:"minimum"`
}

func (m *GasPrices) Reset()         { *m = GasPrices{} }
func (m *GasPrices) String() string { return proto.CompactTextString(m) }
func (*GasPrices) ProtoMessage()    {}
func (*GasPrices) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f91f8bbecb83881, []int{1}
}
func (m *GasPrices) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GasPrices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasPrices.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GasPrices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasPrices.Merge(m, src)
}
func (m *GasPrices) XXX_Size() int {
	return m.Size()
}
func (m *GasPrices) XXX_DiscardUnknown() {
	xxx_messageInfo_GasPrices.DiscardUnknown(m)
}

var xxx_messageInfo_GasPrices proto.InternalMessageInfo

func (m *GasPrices) GetMinimum() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Minimum
	}
	return nil
}

func init() {
	proto.RegisterType((*Authority)(nil), "em.authority.v1.Authority")
	proto.RegisterType((*GasPrices)(nil), "em.authority.v1.GasPrices")
}

func init() { proto.RegisterFile("em/authority/v1/authority.proto", fileDescriptor_3f91f8bbecb83881) }

var fileDescriptor_3f91f8bbecb83881 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0xc1, 0x6e, 0xd4, 0x30,
	0x14, 0x8c, 0xa9, 0x44, 0xb5, 0x29, 0x2d, 0x52, 0x54, 0xa4, 0x65, 0x85, 0xe2, 0x28, 0xa7, 0x95,
	0x60, 0x6d, 0xa5, 0xdc, 0x38, 0xd1, 0x00, 0x82, 0x0b, 0x12, 0x8a, 0x38, 0x71, 0xa9, 0x9c, 0xe4,
	0x6d, 0x6a, 0x11, 0xc7, 0xab, 0xd8, 0x59, 0x91, 0x03, 0x3f, 0xc0, 0xa9, 0xdf, 0xc1, 0x97, 0xf4,
	0xd8, 0x23, 0xa7, 0x14, 0xed, 0xfe, 0xc1, 0x7e, 0x41, 0xb5, 0xb1, 0xad, 0xb6, 0xa7, 0xe4, 0xcd,
	0xbc, 0x99, 0x79, 0x9a, 0xc4, 0xc7, 0x20, 0x28, 0xeb, 0xf4, 0xa5, 0x6c, 0xb9, 0xee, 0xe9, 0x3a,
	0xb9, 0x1f, 0xc8, 0xaa, 0x95, 0x5a, 0x06, 0xcf, 0x41, 0x90, 0x7b, 0x6c, 0x9d, 0xcc, 0x4e, 0x2b,
	0x59, 0xc9, 0x91, 0xa3, 0xfb, 0x37, 0xb3, 0x36, 0x0b, 0x0b, 0xa9, 0x84, 0x54, 0x34, 0x67, 0x0a,
	0xe8, 0x3a, 0xc9, 0x41, 0xb3, 0x84, 0x16, 0x92, 0x37, 0x96, 0xc7, 0x95, 0x94, 0x55, 0x0d, 0x74,
	0x9c, 0xf2, 0x6e, 0x49, 0x35, 0x17, 0xa0, 0x34, 0x13, 0x2b, 0xb3, 0x10, 0x0f, 0xc8, 0x9f, 0x9c,
	0xbb, 0x9c, 0xe0, 0x8d, 0x7f, 0xc8, 0xca, 0xb2, 0x05, 0xa5, 0xa6, 0x28, 0x42, 0xf3, 0x49, 0x1a,
	0xec, 0x06, 0x7c, 0xd2, 0x33, 0x51, 0xbf, 0x8b, 0x2d, 0x11, 0x67, 0x6e, 0x25, 0x78, 0xef, 0x9f,
	0x2c, 0x65, 0x2b, 0xa0, 0xbd, 0x70, 0xa2, 0x27, 0xa3, 0xe8, 0xe5, 0x6e, 0xc0, 0x2f, 0x8c, 0xe8,
	0x31, 0x1f, 0x67, 0xc7, 0x06, 0x38, 0xb7, 0x0e, 0xcc, 0x3f, 0xae, 0x99, 0xd2, 0x17, 0x42, 0x96,
	0x7c, 0xc9, 0xa1, 0x9c, 0x1e, 0x44, 0x68, 0x7e, 0x74, 0x36, 0x23, 0xe6, 0x6c, 0xe2, 0xce, 0x26,
	0xdf, 0xdd, 0xd9, 0x69, 0x74, 0x3d, 0x60, 0x6f, 0x37, 0xe0, 0x53, 0x13, 0xf0, 0x48, 0x1e, 0x5f,
	0xdd, 0x62, 0x94, 0x3d, 0xdb, 0x63, 0x5f, 0x1d, 0xf4, 0x07, 0xf9, 0x93, 0xcf, 0x4c, 0x7d, 0x6b,
	0x79, 0x01, 0x2a, 0xf8, 0xed, 0x1f, 0x0a, 0xde, 0x70, 0xd1, 0x89, 0x29, 0x8a, 0x0e, 0xe6, 0x47,
	0x67, 0xaf, 0x88, 0x69, 0x90, 0xec, 0x1b, 0x24, 0xb6, 0x41, 0xf2, 0x11, 0x8a, 0x0f, 0x92, 0x37,
	0xe9, 0x27, 0x1b, 0x66, 0x2b, 0xb0, 0xd2, 0xf8, 0xef, 0x2d, 0x7e, 0x5d, 0x71, 0x7d, 0xd9, 0xe5,
	0xa4, 0x90, 0x82, 0xda, 0x6f, 0x60, 0x1e, 0x0b, 0x55, 0xfe, 0xa4, 0xba, 0x5f, 0x81, 0x72, 0x2e,
	0x2a, 0x73, 0x99, 0xe9, 0x97, 0xeb, 0x4d, 0x88, 0x6e, 0x36, 0x21, 0xfa, 0xbf, 0x09, 0xd1, 0xd5,
	0x36, 0xf4, 0x6e, 0xb6, 0xa1, 0xf7, 0x6f, 0x1b, 0x7a, 0x3f, 0xc8, 0x03, 0x3f, 0x58, 0x08, 0xd9,
	0x40, 0x4f, 0x41, 0x2c, 0x6a, 0x28, 0x2b, 0x68, 0xe9, 0xaf, 0x07, 0x3f, 0xcb, 0xe8, 0x9d, 0x3f,
	0x1d, 0xab, 0x79, 0x7b, 0x17, 0x00, 0x00, 0xff, 0xff, 0x05, 0xd5, 0xd3, 0x36, 0x49, 0x02, 0x00,
	0x00,
}

func (m *Authority) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Authority) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Authority) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastModified, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastModified):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintAuthority(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if len(m.FormerAddress) > 0 {
		i -= len(m.FormerAddress)
		copy(dAtA[i:], m.FormerAddress)
		i = encodeVarintAuthority(dAtA, i, uint64(len(m.FormerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAuthority(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GasPrices) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasPrices) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasPrices) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Minimum) > 0 {
		for iNdEx := len(m.Minimum) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Minimum[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuthority(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthority(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthority(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Authority) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAuthority(uint64(l))
	}
	l = len(m.FormerAddress)
	if l > 0 {
		n += 1 + l + sovAuthority(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastModified)
	n += 1 + l + sovAuthority(uint64(l))
	return n
}

func (m *GasPrices) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Minimum) > 0 {
		for _, e := range m.Minimum {
			l = e.Size()
			n += 1 + l + sovAuthority(uint64(l))
		}
	}
	return n
}

func sovAuthority(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthority(x uint64) (n int) {
	return sovAuthority(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Authority) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthority
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
			return fmt.Errorf("proto: Authority: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Authority: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthority
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
				return ErrInvalidLengthAuthority
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthority
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FormerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthority
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
				return ErrInvalidLengthAuthority
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthority
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FormerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastModified", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthority
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
				return ErrInvalidLengthAuthority
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthority
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastModified, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthority(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthority
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
func (m *GasPrices) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthority
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
			return fmt.Errorf("proto: GasPrices: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasPrices: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minimum", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthority
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
				return ErrInvalidLengthAuthority
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthority
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Minimum = append(m.Minimum, types.DecCoin{})
			if err := m.Minimum[len(m.Minimum)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthority(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthority
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
func skipAuthority(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthority
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
					return 0, ErrIntOverflowAuthority
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
					return 0, ErrIntOverflowAuthority
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
				return 0, ErrInvalidLengthAuthority
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthority
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthority
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthority        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthority          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthority = fmt.Errorf("proto: unexpected end of group")
)
