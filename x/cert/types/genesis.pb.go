// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/cert/v1alpha1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type GenesisState struct {
	Certifiers        []Certifier   `protobuf:"bytes,1,rep,name=certifiers,proto3" json:"certifiers" yaml:"certifiers"`
	Validators        []Validator   `protobuf:"bytes,2,rep,name=validators,proto3" json:"validators" yaml:"validators"`
	Platforms         []Platform    `protobuf:"bytes,3,rep,name=platforms,proto3" json:"platforms" yaml:"platforms"`
	Certificates      []Certificate `protobuf:"bytes,4,rep,name=certificates,proto3" json:"certificates" yaml:"certificates"`
	Libraries         []Library     `protobuf:"bytes,5,rep,name=libraries,proto3" json:"libraries" yaml:"libraries"`
	NextCertificateId uint64        `protobuf:"varint,6,opt,name=next_certificate_id,json=nextCertificateId,proto3" json:"next_certificate_id,omitempty" yaml:"next_certificate_id"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_860284e2a718f650, []int{0}
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

func init() {
	proto.RegisterType((*GenesisState)(nil), "shentu.cert.v1alpha1.GenesisState")
}

func init() {
	proto.RegisterFile("shentu/cert/v1alpha1/genesis.proto", fileDescriptor_860284e2a718f650)
}

var fileDescriptor_860284e2a718f650 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xbf, 0x6e, 0x9c, 0x30,
	0x18, 0xc0, 0xa1, 0xb9, 0x46, 0x8d, 0x9b, 0xa1, 0x21, 0x19, 0xc8, 0x55, 0x35, 0xa9, 0xa7, 0x4c,
	0x58, 0xb4, 0x5b, 0x46, 0x3a, 0x54, 0x91, 0xaa, 0xaa, 0xa2, 0x6a, 0x86, 0x2c, 0x27, 0x03, 0x86,
	0xb3, 0x0a, 0x18, 0x61, 0x73, 0x0a, 0x6f, 0xd0, 0xb1, 0x8f, 0x90, 0x77, 0xe8, 0x4b, 0x64, 0xcc,
	0xd8, 0xe9, 0x54, 0xdd, 0x2d, 0x9d, 0xf3, 0x04, 0x11, 0x36, 0x04, 0x12, 0xa1, 0x6c, 0x98, 0xef,
	0xf7, 0xfd, 0xbe, 0x3f, 0xfa, 0x00, 0x12, 0x4b, 0x5a, 0xc8, 0x1a, 0x47, 0xb4, 0x92, 0x78, 0xe5,
	0x91, 0xac, 0x5c, 0x12, 0x0f, 0xa7, 0xb4, 0xa0, 0x82, 0x09, 0xb7, 0xac, 0xb8, 0xe4, 0xd6, 0x91,
	0x66, 0xdc, 0x96, 0x71, 0x7b, 0x66, 0x7e, 0x94, 0xf2, 0x94, 0x2b, 0x00, 0xb7, 0x5f, 0x9a, 0x9d,
	0x1f, 0xa7, 0x9c, 0xa7, 0x19, 0xc5, 0xea, 0x15, 0xd6, 0x09, 0x26, 0x45, 0xd3, 0x85, 0x60, 0xc4,
	0x45, 0xce, 0x05, 0x0e, 0x89, 0xa0, 0x78, 0xe5, 0x85, 0x54, 0x12, 0x0f, 0x47, 0x9c, 0x15, 0x7d,
	0xaa, 0x8e, 0x2f, 0xb4, 0x53, 0x3f, 0xba, 0x90, 0x33, 0xd9, 0xa5, 0xea, 0x47, 0x01, 0xe8, 0xcf,
	0x0c, 0xec, 0x7f, 0xd6, 0x4d, 0x7f, 0x97, 0x44, 0x52, 0xeb, 0x12, 0x80, 0x36, 0xcc, 0x12, 0x46,
	0x2b, 0x61, 0x9b, 0x27, 0x3b, 0xa7, 0xaf, 0x3f, 0x38, 0xee, 0xd4, 0x20, 0xee, 0xa7, 0x9e, 0xf3,
	0x8f, 0x6f, 0xd6, 0x8e, 0x71, 0xb7, 0x76, 0x0e, 0x1a, 0x92, 0x67, 0x67, 0x68, 0x10, 0xa0, 0x60,
	0x64, 0x6b, 0xdd, 0x2b, 0x92, 0xb1, 0x98, 0x48, 0x5e, 0x09, 0xfb, 0xc5, 0x73, 0xee, 0x8b, 0x9e,
	0x7b, 0xea, 0x1e, 0x04, 0x28, 0x18, 0xd9, 0xac, 0x0b, 0xb0, 0x57, 0x66, 0x44, 0x26, 0xbc, 0xca,
	0x85, 0xbd, 0xa3, 0xd4, 0x70, 0x5a, 0xfd, 0xad, 0xc3, 0x7c, 0xbb, 0x33, 0xbf, 0xd1, 0xe6, 0x87,
	0x74, 0x14, 0x0c, 0x2a, 0x2b, 0x04, 0xfb, 0xdd, 0x04, 0x11, 0x91, 0x54, 0xd8, 0x33, 0xa5, 0x7e,
	0xff, 0xec, 0x46, 0x5a, 0xd2, 0x7f, 0xdb, 0xd9, 0x0f, 0x1f, 0xed, 0x44, 0x49, 0x50, 0xf0, 0xc8,
	0x69, 0xfd, 0x00, 0x7b, 0x19, 0x0b, 0x2b, 0x52, 0x31, 0x2a, 0xec, 0x97, 0xaa, 0xc0, 0xbb, 0xe9,
	0x02, 0x5f, 0x14, 0xd6, 0x3c, 0x6d, 0xfd, 0x21, 0x1b, 0x05, 0x83, 0xc9, 0xfa, 0x0a, 0x0e, 0x0b,
	0x7a, 0x25, 0x17, 0xa3, 0x5a, 0x0b, 0x16, 0xdb, 0xbb, 0x27, 0xe6, 0xe9, 0xcc, 0x87, 0x77, 0x6b,
	0x67, 0xae, 0xb3, 0x27, 0x20, 0x14, 0x1c, 0xb4, 0x7f, 0x47, 0xf3, 0x9c, 0xc7, 0x67, 0xaf, 0x7e,
	0x5d, 0x3b, 0xc6, 0xff, 0x6b, 0xc7, 0xf0, 0xcf, 0x6f, 0x36, 0xd0, 0xbc, 0xdd, 0x40, 0xf3, 0xdf,
	0x06, 0x9a, 0xbf, 0xb7, 0xd0, 0xb8, 0xdd, 0x42, 0xe3, 0xef, 0x16, 0x1a, 0x97, 0x38, 0x65, 0x72,
	0x59, 0x87, 0x6e, 0xc4, 0x73, 0x75, 0x66, 0xec, 0x67, 0xc2, 0xeb, 0x22, 0x26, 0x92, 0xf1, 0x02,
	0x77, 0xc7, 0x78, 0xa5, 0xcf, 0x51, 0x36, 0x25, 0x15, 0xe1, 0xae, 0xba, 0xc3, 0x8f, 0xf7, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xac, 0x8e, 0x44, 0x7c, 0x50, 0x03, 0x00, 0x00,
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
	if m.NextCertificateId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextCertificateId))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Libraries) > 0 {
		for iNdEx := len(m.Libraries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Libraries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Certificates) > 0 {
		for iNdEx := len(m.Certificates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Certificates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Platforms) > 0 {
		for iNdEx := len(m.Platforms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Platforms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Certifiers) > 0 {
		for iNdEx := len(m.Certifiers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Certifiers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
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
	if len(m.Certifiers) > 0 {
		for _, e := range m.Certifiers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Platforms) > 0 {
		for _, e := range m.Platforms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Certificates) > 0 {
		for _, e := range m.Certificates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Libraries) > 0 {
		for _, e := range m.Libraries {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.NextCertificateId != 0 {
		n += 1 + sovGenesis(uint64(m.NextCertificateId))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Certifiers", wireType)
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
			m.Certifiers = append(m.Certifiers, Certifier{})
			if err := m.Certifiers[len(m.Certifiers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
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
			m.Validators = append(m.Validators, Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Platforms", wireType)
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
			m.Platforms = append(m.Platforms, Platform{})
			if err := m.Platforms[len(m.Platforms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificates", wireType)
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
			m.Certificates = append(m.Certificates, Certificate{})
			if err := m.Certificates[len(m.Certificates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Libraries", wireType)
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
			m.Libraries = append(m.Libraries, Library{})
			if err := m.Libraries[len(m.Libraries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextCertificateId", wireType)
			}
			m.NextCertificateId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextCertificateId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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