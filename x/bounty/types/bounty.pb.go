// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/bounty/v1/bounty.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type Program struct {
	ProgramId         uint64                                 `protobuf:"varint,1,opt,name=program_id,json=programId,proto3" json:"id" yaml:"id"`
	CreatorAddress    string                                 `protobuf:"bytes,2,opt,name=creator_address,json=creatorAddress,proto3" json:"creator_address,omitempty" yaml:"creator_address"`
	SubmissionEndTime time.Time                              `protobuf:"bytes,3,opt,name=submission_end_time,json=submissionEndTime,proto3,stdtime" json:"submission_end_time" yaml:"submission_end_time"`
	JudgingEndTime    time.Time                              `protobuf:"bytes,4,opt,name=judging_end_time,json=judgingEndTime,proto3,stdtime" json:"judging_end_time" yaml:"judging_end_time"`
	ClaimEndTime      time.Time                              `protobuf:"bytes,5,opt,name=claim_end_time,json=claimEndTime,proto3,stdtime" json:"claim_end_time" yaml:"claim_end_time"`
	Description       string                                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	EncryptionKey     *types.Any                             `protobuf:"bytes,7,opt,name=encryption_key,json=encryptionKey,proto3" json:"encryption_key,omitempty" yaml:"encryption_key"`
	Deposit           []types1.Coin                          `protobuf:"bytes,8,rep,name=deposit,proto3" json:"deposit" yaml:"deposit"`
	CommissionRate    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=commission_rate,json=commissionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"commission_rate"`
}

func (m *Program) Reset()         { *m = Program{} }
func (m *Program) String() string { return proto.CompactTextString(m) }
func (*Program) ProtoMessage()    {}
func (*Program) Descriptor() ([]byte, []int) {
	return fileDescriptor_36e6d679af1b94c6, []int{0}
}
func (m *Program) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Program) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Program.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Program) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Program.Merge(m, src)
}
func (m *Program) XXX_Size() int {
	return m.Size()
}
func (m *Program) XXX_DiscardUnknown() {
	xxx_messageInfo_Program.DiscardUnknown(m)
}

var xxx_messageInfo_Program proto.InternalMessageInfo

func (m *Program) GetProgramId() uint64 {
	if m != nil {
		return m.ProgramId
	}
	return 0
}

func (m *Program) GetCreatorAddress() string {
	if m != nil {
		return m.CreatorAddress
	}
	return ""
}

func (m *Program) GetSubmissionEndTime() time.Time {
	if m != nil {
		return m.SubmissionEndTime
	}
	return time.Time{}
}

func (m *Program) GetJudgingEndTime() time.Time {
	if m != nil {
		return m.JudgingEndTime
	}
	return time.Time{}
}

func (m *Program) GetClaimEndTime() time.Time {
	if m != nil {
		return m.ClaimEndTime
	}
	return time.Time{}
}

func (m *Program) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Program) GetEncryptionKey() *types.Any {
	if m != nil {
		return m.EncryptionKey
	}
	return nil
}

func (m *Program) GetDeposit() []types1.Coin {
	if m != nil {
		return m.Deposit
	}
	return nil
}

func init() {
	proto.RegisterType((*Program)(nil), "shentu.bounty.v1.Program")
}

func init() { proto.RegisterFile("shentu/bounty/v1/bounty.proto", fileDescriptor_36e6d679af1b94c6) }

var fileDescriptor_36e6d679af1b94c6 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x4f, 0x4f, 0xd4, 0x4c,
	0x1c, 0xc7, 0xb7, 0xc0, 0x03, 0xcf, 0x0e, 0xba, 0x60, 0x41, 0x2c, 0x9b, 0xd8, 0xae, 0x35, 0x21,
	0x7b, 0x61, 0x9a, 0xc5, 0x8b, 0xe1, 0x46, 0xd1, 0x83, 0xd9, 0x0b, 0x69, 0x4c, 0x4c, 0xbc, 0x34,
	0xd3, 0xce, 0x50, 0x46, 0xe8, 0x4c, 0xd3, 0x99, 0x12, 0xfb, 0x2e, 0x78, 0x01, 0x1e, 0x7c, 0x11,
	0xbe, 0x08, 0xe2, 0x89, 0xa3, 0xf1, 0x50, 0x0d, 0x5c, 0x8c, 0xc7, 0x7d, 0x05, 0xa6, 0x9d, 0x69,
	0x0a, 0xab, 0x09, 0xa7, 0x9d, 0xf9, 0xfd, 0xf9, 0x7e, 0xe6, 0x3b, 0xf3, 0xdb, 0x82, 0xa7, 0xe2,
	0x84, 0x30, 0x59, 0x78, 0x11, 0x2f, 0x98, 0x2c, 0xbd, 0xf3, 0x89, 0x5e, 0xc1, 0x2c, 0xe7, 0x92,
	0x9b, 0xeb, 0x2a, 0x0d, 0x75, 0xf0, 0x7c, 0x32, 0xdc, 0x4c, 0x78, 0xc2, 0x9b, 0xa4, 0x57, 0xaf,
	0x54, 0xdd, 0x70, 0x3b, 0xe1, 0x3c, 0x39, 0x23, 0x5e, 0xb3, 0x8b, 0x8a, 0x63, 0x0f, 0x31, 0x2d,
	0x31, 0x74, 0xe6, 0x53, 0x92, 0xa6, 0x44, 0x48, 0x94, 0x66, 0xba, 0xc0, 0x8e, 0xb9, 0x48, 0xb9,
	0xf0, 0x22, 0x24, 0x88, 0x77, 0x3e, 0x89, 0x88, 0x44, 0x13, 0x2f, 0xe6, 0x94, 0xb5, 0xda, 0x2a,
	0x1f, 0x2a, 0xa8, 0xda, 0xa8, 0x94, 0xfb, 0x69, 0x19, 0xac, 0x1c, 0xe5, 0x3c, 0xc9, 0x51, 0x6a,
	0xee, 0x01, 0x90, 0xa9, 0x65, 0x48, 0xb1, 0x65, 0x8c, 0x8c, 0xf1, 0x92, 0xbf, 0xf1, 0xbb, 0x72,
	0x16, 0x28, 0x9e, 0x55, 0x4e, 0xbf, 0x44, 0xe9, 0xd9, 0xbe, 0x4b, 0xb1, 0x1b, 0xf4, 0x75, 0xd9,
	0x1b, 0x6c, 0x1e, 0x82, 0xb5, 0x38, 0x27, 0x48, 0xf2, 0x3c, 0x44, 0x18, 0xe7, 0x44, 0x08, 0x6b,
	0x61, 0x64, 0x8c, 0xfb, 0xfe, 0x70, 0x56, 0x39, 0x5b, 0xaa, 0x65, 0xae, 0xc0, 0x0d, 0x06, 0x3a,
	0x72, 0xa0, 0x02, 0x66, 0x0e, 0x36, 0x44, 0x11, 0xa5, 0x54, 0x08, 0xca, 0x59, 0x48, 0x18, 0x0e,
	0x6b, 0x87, 0xd6, 0xe2, 0xc8, 0x18, 0xaf, 0xee, 0x0d, 0xa1, 0xb2, 0x0f, 0x5b, 0xfb, 0xf0, 0x6d,
	0x6b, 0xdf, 0xdf, 0xb9, 0xac, 0x9c, 0xde, 0xac, 0x72, 0x86, 0x0a, 0xf4, 0x0f, 0x11, 0xf7, 0xe2,
	0x87, 0x63, 0x04, 0x8f, 0xba, 0xcc, 0x6b, 0x86, 0xeb, 0x7e, 0x93, 0x82, 0xf5, 0x0f, 0x05, 0x4e,
	0x28, 0x4b, 0x3a, 0xe0, 0xd2, 0xbd, 0xc0, 0xe7, 0x1a, 0xf8, 0x44, 0x01, 0xe7, 0x15, 0x14, 0x6d,
	0xa0, 0xc3, 0x2d, 0x2a, 0x06, 0x83, 0xf8, 0x0c, 0xd1, 0xb4, 0x03, 0xfd, 0x77, 0x2f, 0xe8, 0x99,
	0x06, 0x3d, 0xd6, 0x57, 0x78, 0xa7, 0x5f, 0x61, 0x1e, 0x34, 0xc1, 0x16, 0xf2, 0x12, 0xac, 0x62,
	0x22, 0xe2, 0x9c, 0x66, 0x92, 0x72, 0x66, 0x2d, 0x37, 0x8f, 0xb0, 0x35, 0xab, 0x1c, 0x53, 0x29,
	0xdc, 0x4a, 0xba, 0xc1, 0xed, 0x52, 0x93, 0x81, 0x01, 0x61, 0x71, 0x5e, 0x36, 0xbb, 0xf0, 0x94,
	0x94, 0xd6, 0x4a, 0x73, 0xbc, 0xcd, 0xbf, 0x8e, 0x77, 0xc0, 0x4a, 0x7f, 0xd2, 0x1d, 0xea, 0x6e,
	0x97, 0xfb, 0xf5, 0xcb, 0xee, 0xa6, 0x9e, 0xad, 0x26, 0xce, 0xe1, 0x51, 0x11, 0x4d, 0x49, 0x19,
	0x3c, 0xec, 0x0a, 0xa7, 0xa4, 0x34, 0xa7, 0x60, 0x05, 0x93, 0x8c, 0x0b, 0x2a, 0xad, 0xff, 0x47,
	0x8b, 0xe3, 0xd5, 0xbd, 0x6d, 0xa8, 0xdb, 0xea, 0xf9, 0x85, 0x7a, 0x7e, 0xe1, 0x21, 0xa7, 0xcc,
	0xdf, 0xd2, 0xd7, 0x30, 0x68, 0x4d, 0x34, 0x7d, 0x6e, 0xd0, 0x2a, 0x98, 0xef, 0xc0, 0x5a, 0xcc,
	0xd3, 0xf6, 0xd5, 0x73, 0x24, 0x89, 0xd5, 0x6f, 0xac, 0xc3, 0xba, 0xf3, 0x7b, 0xe5, 0xec, 0x24,
	0x54, 0x9e, 0x14, 0x11, 0x8c, 0x79, 0xaa, 0x27, 0x5f, 0xff, 0xec, 0x0a, 0x7c, 0xea, 0xc9, 0x32,
	0x23, 0x02, 0xbe, 0x22, 0x71, 0x30, 0xe8, 0x64, 0x02, 0x24, 0xc9, 0xfe, 0xd2, 0xaf, 0xcf, 0x4e,
	0xcf, 0x9f, 0x5e, 0x5e, 0xdb, 0xc6, 0xd5, 0xb5, 0x6d, 0xfc, 0xbc, 0xb6, 0x8d, 0x8b, 0x1b, 0xbb,
	0x77, 0x75, 0x63, 0xf7, 0xbe, 0xdd, 0xd8, 0xbd, 0xf7, 0x93, 0x5b, 0xba, 0xea, 0x2f, 0x7e, 0xcc,
	0x0b, 0x86, 0x51, 0xed, 0x52, 0x07, 0xbc, 0x8f, 0xed, 0x47, 0xa1, 0xc1, 0x44, 0xcb, 0xcd, 0x45,
	0xbe, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x5f, 0x3e, 0xb6, 0x32, 0x04, 0x00, 0x00,
}

func (m *Program) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Program) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Program) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CommissionRate.Size()
		i -= size
		if _, err := m.CommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBounty(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.Deposit) > 0 {
		for iNdEx := len(m.Deposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBounty(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.EncryptionKey != nil {
		{
			size, err := m.EncryptionKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBounty(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintBounty(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x32
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ClaimEndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ClaimEndTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintBounty(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.JudgingEndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.JudgingEndTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintBounty(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.SubmissionEndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.SubmissionEndTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintBounty(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x1a
	if len(m.CreatorAddress) > 0 {
		i -= len(m.CreatorAddress)
		copy(dAtA[i:], m.CreatorAddress)
		i = encodeVarintBounty(dAtA, i, uint64(len(m.CreatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProgramId != 0 {
		i = encodeVarintBounty(dAtA, i, uint64(m.ProgramId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBounty(dAtA []byte, offset int, v uint64) int {
	offset -= sovBounty(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Program) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProgramId != 0 {
		n += 1 + sovBounty(uint64(m.ProgramId))
	}
	l = len(m.CreatorAddress)
	if l > 0 {
		n += 1 + l + sovBounty(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.SubmissionEndTime)
	n += 1 + l + sovBounty(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.JudgingEndTime)
	n += 1 + l + sovBounty(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ClaimEndTime)
	n += 1 + l + sovBounty(uint64(l))
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovBounty(uint64(l))
	}
	if m.EncryptionKey != nil {
		l = m.EncryptionKey.Size()
		n += 1 + l + sovBounty(uint64(l))
	}
	if len(m.Deposit) > 0 {
		for _, e := range m.Deposit {
			l = e.Size()
			n += 1 + l + sovBounty(uint64(l))
		}
	}
	l = m.CommissionRate.Size()
	n += 1 + l + sovBounty(uint64(l))
	return n
}

func sovBounty(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBounty(x uint64) (n int) {
	return sovBounty(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Program) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBounty
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
			return fmt.Errorf("proto: Program: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Program: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProgramId", wireType)
			}
			m.ProgramId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProgramId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
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
				return ErrInvalidLengthBounty
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBounty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmissionEndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
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
				return ErrInvalidLengthBounty
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBounty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.SubmissionEndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JudgingEndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
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
				return ErrInvalidLengthBounty
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBounty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.JudgingEndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimEndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
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
				return ErrInvalidLengthBounty
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBounty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ClaimEndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
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
				return ErrInvalidLengthBounty
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBounty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptionKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
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
				return ErrInvalidLengthBounty
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBounty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EncryptionKey == nil {
				m.EncryptionKey = &types.Any{}
			}
			if err := m.EncryptionKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
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
				return ErrInvalidLengthBounty
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBounty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposit = append(m.Deposit, types1.Coin{})
			if err := m.Deposit[len(m.Deposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBounty
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
				return ErrInvalidLengthBounty
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBounty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBounty(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBounty
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBounty
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
func skipBounty(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBounty
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
					return 0, ErrIntOverflowBounty
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
					return 0, ErrIntOverflowBounty
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
				return 0, ErrInvalidLengthBounty
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBounty
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBounty
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBounty        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBounty          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBounty = fmt.Errorf("proto: unexpected end of group")
)