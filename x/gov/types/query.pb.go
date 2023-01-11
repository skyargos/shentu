// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/gov/v1alpha1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/gov/types"
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

// QueryProposalResponse is the response type for the Query/Proposal RPC method.
type QueryProposalResponse struct {
	Proposal types.Proposal `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal"`
	// cert_voted define a special proposal that has completed the vote of the cert.
	CertVoted bool `protobuf:"varint,2,opt,name=cert_voted,json=certVoted,proto3" json:"cert_voted,omitempty"`
}

func (m *QueryProposalResponse) Reset()         { *m = QueryProposalResponse{} }
func (m *QueryProposalResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProposalResponse) ProtoMessage()    {}
func (*QueryProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f945a4e1db5124e, []int{0}
}
func (m *QueryProposalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProposalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProposalResponse.Merge(m, src)
}
func (m *QueryProposalResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProposalResponse proto.InternalMessageInfo

func (m *QueryProposalResponse) GetProposal() types.Proposal {
	if m != nil {
		return m.Proposal
	}
	return types.Proposal{}
}

func (m *QueryProposalResponse) GetCertVoted() bool {
	if m != nil {
		return m.CertVoted
	}
	return false
}

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// voting_params defines the parameters related to voting.
	VotingParams types.VotingParams `protobuf:"bytes,1,opt,name=voting_params,json=votingParams,proto3" json:"voting_params"`
	// deposit_params defines the parameters related to deposit.
	DepositParams types.DepositParams `protobuf:"bytes,2,opt,name=deposit_params,json=depositParams,proto3" json:"deposit_params"`
	// tally_params defines the parameters related to tally.
	TallyParams types.TallyParams `protobuf:"bytes,3,opt,name=tally_params,json=tallyParams,proto3" json:"tally_params"`
	// custom_params defines the parameters related to custom.
	CustomParams CustomParams `protobuf:"bytes,4,opt,name=custom_params,json=customParams,proto3" json:"custom_params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f945a4e1db5124e, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetVotingParams() types.VotingParams {
	if m != nil {
		return m.VotingParams
	}
	return types.VotingParams{}
}

func (m *QueryParamsResponse) GetDepositParams() types.DepositParams {
	if m != nil {
		return m.DepositParams
	}
	return types.DepositParams{}
}

func (m *QueryParamsResponse) GetTallyParams() types.TallyParams {
	if m != nil {
		return m.TallyParams
	}
	return types.TallyParams{}
}

func (m *QueryParamsResponse) GetCustomParams() CustomParams {
	if m != nil {
		return m.CustomParams
	}
	return CustomParams{}
}

func init() {
	proto.RegisterType((*QueryProposalResponse)(nil), "shentu.gov.v1alpha1.QueryProposalResponse")
	proto.RegisterType((*QueryParamsResponse)(nil), "shentu.gov.v1alpha1.QueryParamsResponse")
}

func init() { proto.RegisterFile("shentu/gov/v1alpha1/query.proto", fileDescriptor_9f945a4e1db5124e) }

var fileDescriptor_9f945a4e1db5124e = []byte{
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xc1, 0x6b, 0x13, 0x41,
	0x14, 0xc6, 0x33, 0xb1, 0xad, 0xe9, 0xb4, 0xf5, 0xf0, 0xaa, 0x50, 0x42, 0xbb, 0xa9, 0xc1, 0x6a,
	0xda, 0xea, 0x8e, 0x69, 0x45, 0x28, 0x8a, 0x68, 0xf5, 0x50, 0x50, 0x44, 0x83, 0xf4, 0xe0, 0x25,
	0x6c, 0x93, 0x71, 0xbb, 0xb0, 0xd9, 0xd9, 0xee, 0xcc, 0x2e, 0x96, 0x90, 0x8b, 0x37, 0xc1, 0x83,
	0x22, 0x7a, 0xf3, 0x24, 0xf8, 0xb7, 0x14, 0x4f, 0x05, 0x41, 0x3c, 0x89, 0xb4, 0xfe, 0x21, 0x32,
	0xb3, 0x33, 0x69, 0x22, 0x9b, 0xa6, 0xd5, 0xd3, 0x0e, 0xef, 0x7d, 0xfb, 0x7d, 0xbf, 0x9d, 0x7d,
	0xb3, 0x8b, 0x4b, 0x7c, 0x9b, 0x06, 0x22, 0x26, 0x2e, 0x4b, 0x48, 0x52, 0x75, 0xfc, 0x70, 0xdb,
	0xa9, 0x92, 0x9d, 0x98, 0x46, 0xbb, 0x76, 0x18, 0x31, 0xc1, 0x60, 0x3a, 0x15, 0xd8, 0x2e, 0x4b,
	0x6c, 0x23, 0x28, 0x9e, 0x77, 0x99, 0xcb, 0x54, 0x9f, 0xc8, 0x55, 0x2a, 0x2d, 0xce, 0xba, 0x8c,
	0xb9, 0x3e, 0x25, 0x4e, 0xe8, 0x11, 0x27, 0x08, 0x98, 0x70, 0x84, 0xc7, 0x02, 0x6e, 0xba, 0x0d,
	0xc6, 0x5b, 0x8c, 0xeb, 0xa4, 0x2d, 0x2a, 0x9c, 0xaa, 0x5c, 0xeb, 0xae, 0x95, 0xd1, 0xed, 0xc1,
	0x28, 0xce, 0x65, 0x71, 0x76, 0x6f, 0x2f, 0x27, 0xf8, 0xc2, 0x53, 0xa9, 0x7e, 0x12, 0xb1, 0x90,
	0x71, 0xc7, 0xaf, 0x51, 0x1e, 0xb2, 0x80, 0x53, 0xb8, 0x83, 0x0b, 0xa1, 0xae, 0xcd, 0xa0, 0x79,
	0x54, 0x99, 0x58, 0x99, 0xb5, 0xd3, 0x28, 0xfd, 0x44, 0x2a, 0xca, 0x36, 0xf7, 0xad, 0x8f, 0xec,
	0xfd, 0x2c, 0xe5, 0x6a, 0xdd, 0x7b, 0x60, 0x0e, 0xe3, 0x06, 0x8d, 0x44, 0x3d, 0x61, 0x82, 0x36,
	0x67, 0xf2, 0xf3, 0xa8, 0x52, 0xa8, 0x8d, 0xcb, 0xca, 0xa6, 0x2c, 0x94, 0xbf, 0xe6, 0xf1, 0x74,
	0x1a, 0xec, 0x44, 0x4e, 0x8b, 0x77, 0x63, 0x1f, 0xe2, 0xa9, 0x84, 0x09, 0x2f, 0x70, 0xeb, 0xa1,
	0x6a, 0xe8, 0xec, 0xf9, 0xac, 0xec, 0x4d, 0x25, 0x4c, 0x0d, 0x74, 0xfe, 0x64, 0xd2, 0x53, 0x83,
	0xc7, 0xf8, 0x5c, 0x93, 0x86, 0x8c, 0x7b, 0xc2, 0xb8, 0xe5, 0x95, 0xdb, 0xc5, 0x2c, 0xb7, 0x07,
	0xa9, 0xb2, 0xcf, 0x6e, 0xaa, 0xd9, 0x5b, 0x84, 0x0d, 0x3c, 0x29, 0x1c, 0xdf, 0xdf, 0x35, 0x6e,
	0x67, 0x94, 0x5b, 0x29, 0xcb, 0xed, 0x99, 0xd4, 0xf5, 0x79, 0x4d, 0x88, 0xa3, 0x12, 0x3c, 0xc2,
	0x53, 0x8d, 0x98, 0x0b, 0xd6, 0x32, 0x56, 0x23, 0x1a, 0x2c, 0x63, 0x68, 0xec, 0xfb, 0x4a, 0xd9,
	0xff, 0x9c, 0x8d, 0x9e, 0xda, 0xca, 0xf7, 0x71, 0x3c, 0xaa, 0x36, 0x13, 0x3e, 0x20, 0x5c, 0x30,
	0xaf, 0x04, 0x2a, 0x59, 0x60, 0x7f, 0xbd, 0xed, 0x9d, 0x98, 0x72, 0x51, 0x5c, 0xca, 0xcc, 0xcd,
	0x1c, 0x8c, 0xf2, 0xea, 0xab, 0x6f, 0xbf, 0xdf, 0xe7, 0xaf, 0xc1, 0x32, 0xc9, 0x98, 0x3c, 0xf3,
	0xfa, 0x39, 0x69, 0x9b, 0x65, 0xdd, 0x6b, 0x76, 0xe0, 0x35, 0xc2, 0xe3, 0xc6, 0x89, 0xc3, 0xe2,
	0x50, 0x30, 0x7e, 0x44, 0x76, 0x02, 0xa9, 0x26, 0x5b, 0x50, 0x64, 0x25, 0x98, 0x3b, 0x96, 0x0c,
	0x3e, 0x22, 0x3c, 0x22, 0x87, 0x10, 0x2e, 0x0d, 0xf4, 0x96, 0x6d, 0x43, 0xb0, 0x30, 0x44, 0xa5,
	0xc3, 0xef, 0xa9, 0xf0, 0x5b, 0xb0, 0x76, 0x8a, 0x6d, 0x21, 0xf2, 0x70, 0x70, 0xd2, 0x96, 0x97,
	0xa8, 0x03, 0xef, 0x10, 0x1e, 0x95, 0x9e, 0x1c, 0x8e, 0xcf, 0xec, 0x6e, 0xce, 0xe5, 0x61, 0x32,
	0xcd, 0xb6, 0xa6, 0xd8, 0x56, 0xa1, 0x7a, 0x6a, 0x36, 0x78, 0x83, 0xf0, 0x98, 0x9e, 0xd9, 0xc1,
	0x69, 0xe6, 0x0c, 0xa7, 0x54, 0x95, 0x63, 0x86, 0xa9, 0xef, 0xb0, 0x97, 0xaf, 0x2b, 0xae, 0x25,
	0xa8, 0x64, 0x72, 0x29, 0x2d, 0x69, 0xa7, 0xd7, 0xba, 0xd8, 0x0d, 0x69, 0x07, 0xbe, 0x20, 0x7c,
	0x56, 0x1f, 0x54, 0xb8, 0x32, 0x90, 0x47, 0x2b, 0x8e, 0x80, 0x86, 0x0a, 0x35, 0xd0, 0x86, 0x02,
	0x5a, 0x87, 0xbb, 0xa7, 0xd9, 0x28, 0xfd, 0x8d, 0xe0, 0xa4, 0xad, 0x57, 0x2c, 0xea, 0xc0, 0x27,
	0x84, 0x0b, 0xda, 0x9d, 0xc3, 0x50, 0x80, 0xee, 0xde, 0x2d, 0x9e, 0x40, 0xa9, 0x59, 0x6f, 0x2b,
	0xd6, 0x9b, 0x70, 0xe3, 0x5f, 0x58, 0xe1, 0x33, 0xc2, 0x13, 0xea, 0x1b, 0x55, 0xa3, 0x3c, 0xf6,
	0x05, 0x2c, 0x0f, 0x0c, 0xee, 0x51, 0x19, 0xca, 0xab, 0x27, 0x13, 0xff, 0xcf, 0xf4, 0xa9, 0x8f,
	0xe5, 0xfa, 0xc6, 0xde, 0x81, 0x85, 0xf6, 0x0f, 0x2c, 0xf4, 0xeb, 0xc0, 0x42, 0x6f, 0x0f, 0xad,
	0xdc, 0xfe, 0xa1, 0x95, 0xfb, 0x71, 0x68, 0xe5, 0x9e, 0xdb, 0xae, 0x27, 0xb6, 0xe3, 0x2d, 0xbb,
	0xc1, 0x5a, 0x24, 0x1d, 0xb7, 0x17, 0x2c, 0x0e, 0x9a, 0xea, 0xc7, 0xa9, 0x0b, 0xe4, 0xa5, 0x4a,
	0x92, 0x83, 0xc3, 0xb7, 0xc6, 0xd4, 0xef, 0x6e, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2,
	0x9a, 0xc3, 0xad, 0xb7, 0x07, 0x00, 0x00,
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
	// Proposal queries proposal details based on ProposalID.
	Proposal(ctx context.Context, in *types.QueryProposalRequest, opts ...grpc.CallOption) (*QueryProposalResponse, error)
	// Proposals queries all proposals based on given status.
	Proposals(ctx context.Context, in *types.QueryProposalsRequest, opts ...grpc.CallOption) (*types.QueryProposalsResponse, error)
	// Vote queries voted information based on proposalID, voterAddr.
	Vote(ctx context.Context, in *types.QueryVoteRequest, opts ...grpc.CallOption) (*types.QueryVoteResponse, error)
	// Votes queries votes of a given proposal.
	Votes(ctx context.Context, in *types.QueryVotesRequest, opts ...grpc.CallOption) (*types.QueryVotesResponse, error)
	// Params queries all parameters of the gov module.
	Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Deposit queries single deposit information based proposalID, depositAddr.
	Deposit(ctx context.Context, in *types.QueryDepositRequest, opts ...grpc.CallOption) (*types.QueryDepositResponse, error)
	// Deposits queries all deposits of a single proposal.
	Deposits(ctx context.Context, in *types.QueryDepositsRequest, opts ...grpc.CallOption) (*types.QueryDepositsResponse, error)
	// TallyResult queries the tally of a proposal vote.
	TallyResult(ctx context.Context, in *types.QueryTallyResultRequest, opts ...grpc.CallOption) (*types.QueryTallyResultResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Proposal(ctx context.Context, in *types.QueryProposalRequest, opts ...grpc.CallOption) (*QueryProposalResponse, error) {
	out := new(QueryProposalResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1alpha1.Query/Proposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Proposals(ctx context.Context, in *types.QueryProposalsRequest, opts ...grpc.CallOption) (*types.QueryProposalsResponse, error) {
	out := new(types.QueryProposalsResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1alpha1.Query/Proposals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Vote(ctx context.Context, in *types.QueryVoteRequest, opts ...grpc.CallOption) (*types.QueryVoteResponse, error) {
	out := new(types.QueryVoteResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1alpha1.Query/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Votes(ctx context.Context, in *types.QueryVotesRequest, opts ...grpc.CallOption) (*types.QueryVotesResponse, error) {
	out := new(types.QueryVotesResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1alpha1.Query/Votes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1alpha1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Deposit(ctx context.Context, in *types.QueryDepositRequest, opts ...grpc.CallOption) (*types.QueryDepositResponse, error) {
	out := new(types.QueryDepositResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1alpha1.Query/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Deposits(ctx context.Context, in *types.QueryDepositsRequest, opts ...grpc.CallOption) (*types.QueryDepositsResponse, error) {
	out := new(types.QueryDepositsResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1alpha1.Query/Deposits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TallyResult(ctx context.Context, in *types.QueryTallyResultRequest, opts ...grpc.CallOption) (*types.QueryTallyResultResponse, error) {
	out := new(types.QueryTallyResultResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1alpha1.Query/TallyResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Proposal queries proposal details based on ProposalID.
	Proposal(context.Context, *types.QueryProposalRequest) (*QueryProposalResponse, error)
	// Proposals queries all proposals based on given status.
	Proposals(context.Context, *types.QueryProposalsRequest) (*types.QueryProposalsResponse, error)
	// Vote queries voted information based on proposalID, voterAddr.
	Vote(context.Context, *types.QueryVoteRequest) (*types.QueryVoteResponse, error)
	// Votes queries votes of a given proposal.
	Votes(context.Context, *types.QueryVotesRequest) (*types.QueryVotesResponse, error)
	// Params queries all parameters of the gov module.
	Params(context.Context, *types.QueryParamsRequest) (*QueryParamsResponse, error)
	// Deposit queries single deposit information based proposalID, depositAddr.
	Deposit(context.Context, *types.QueryDepositRequest) (*types.QueryDepositResponse, error)
	// Deposits queries all deposits of a single proposal.
	Deposits(context.Context, *types.QueryDepositsRequest) (*types.QueryDepositsResponse, error)
	// TallyResult queries the tally of a proposal vote.
	TallyResult(context.Context, *types.QueryTallyResultRequest) (*types.QueryTallyResultResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Proposal(ctx context.Context, req *types.QueryProposalRequest) (*QueryProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proposal not implemented")
}
func (*UnimplementedQueryServer) Proposals(ctx context.Context, req *types.QueryProposalsRequest) (*types.QueryProposalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proposals not implemented")
}
func (*UnimplementedQueryServer) Vote(ctx context.Context, req *types.QueryVoteRequest) (*types.QueryVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (*UnimplementedQueryServer) Votes(ctx context.Context, req *types.QueryVotesRequest) (*types.QueryVotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Votes not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *types.QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Deposit(ctx context.Context, req *types.QueryDepositRequest) (*types.QueryDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (*UnimplementedQueryServer) Deposits(ctx context.Context, req *types.QueryDepositsRequest) (*types.QueryDepositsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposits not implemented")
}
func (*UnimplementedQueryServer) TallyResult(ctx context.Context, req *types.QueryTallyResultRequest) (*types.QueryTallyResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TallyResult not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Proposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Proposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1alpha1.Query/Proposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Proposal(ctx, req.(*types.QueryProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Proposals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryProposalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Proposals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1alpha1.Query/Proposals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Proposals(ctx, req.(*types.QueryProposalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1alpha1.Query/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Vote(ctx, req.(*types.QueryVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Votes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryVotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Votes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1alpha1.Query/Votes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Votes(ctx, req.(*types.QueryVotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1alpha1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*types.QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1alpha1.Query/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Deposit(ctx, req.(*types.QueryDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Deposits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryDepositsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Deposits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1alpha1.Query/Deposits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Deposits(ctx, req.(*types.QueryDepositsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TallyResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryTallyResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TallyResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1alpha1.Query/TallyResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TallyResult(ctx, req.(*types.QueryTallyResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shentu.gov.v1alpha1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Proposal",
			Handler:    _Query_Proposal_Handler,
		},
		{
			MethodName: "Proposals",
			Handler:    _Query_Proposals_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Query_Vote_Handler,
		},
		{
			MethodName: "Votes",
			Handler:    _Query_Votes_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Query_Deposit_Handler,
		},
		{
			MethodName: "Deposits",
			Handler:    _Query_Deposits_Handler,
		},
		{
			MethodName: "TallyResult",
			Handler:    _Query_TallyResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shentu/gov/v1alpha1/query.proto",
}

func (m *QueryProposalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProposalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProposalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CertVoted {
		i--
		if m.CertVoted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Proposal.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CustomParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.TallyParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.DepositParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.VotingParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *QueryProposalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Proposal.Size()
	n += 1 + l + sovQuery(uint64(l))
	if m.CertVoted {
		n += 2
	}
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.VotingParams.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.DepositParams.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.TallyParams.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.CustomParams.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryProposalResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryProposalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProposalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposal", wireType)
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
			if err := m.Proposal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertVoted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
			m.CertVoted = bool(v != 0)
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingParams", wireType)
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
			if err := m.VotingParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositParams", wireType)
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
			if err := m.DepositParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TallyParams", wireType)
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
			if err := m.TallyParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomParams", wireType)
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
			if err := m.CustomParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
