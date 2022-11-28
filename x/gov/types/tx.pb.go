// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/gov/v1alpha1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/gov/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() { proto.RegisterFile("shentu/gov/v1alpha1/tx.proto", fileDescriptor_f5034648f58c6e59) }

var fileDescriptor_f5034648f58c6e59 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0x3b, 0xdf, 0x27, 0x0a, 0x41, 0x5c, 0x44, 0x37, 0xa6, 0x92, 0x8d, 0xe8, 0x46, 0x4c,
	0x18, 0x7d, 0x03, 0x71, 0xe1, 0xc2, 0x82, 0x7f, 0x40, 0xc1, 0x8d, 0x26, 0x6d, 0x9a, 0x09, 0xb4,
	0xb9, 0xa1, 0xc9, 0x0c, 0xed, 0x5b, 0xf8, 0x58, 0x2e, 0xbb, 0xec, 0x52, 0xa6, 0x5b, 0x1f, 0x42,
	0x3a, 0x99, 0x29, 0xb8, 0xe8, 0xd0, 0xdd, 0x9c, 0xfb, 0x3b, 0xe7, 0xcc, 0x85, 0x5c, 0x74, 0xe2,
	0x33, 0x65, 0x43, 0xce, 0x35, 0x14, 0xbc, 0x48, 0xc5, 0xc8, 0x65, 0x22, 0xe5, 0x61, 0xca, 0xdc,
	0x04, 0x02, 0xe0, 0xc3, 0x48, 0x99, 0x86, 0x82, 0x35, 0x94, 0xd0, 0x3e, 0xf8, 0x31, 0x78, 0x2e,
	0x85, 0x57, 0xbc, 0x48, 0xa5, 0x0a, 0x22, 0xe5, 0x7d, 0x30, 0x36, 0x86, 0x48, 0xb7, 0xe6, 0xb1,
	0x32, 0xe2, 0xa6, 0x91, 0x1c, 0x47, 0xf8, 0x5e, 0x29, 0x1e, 0x45, 0x8d, 0x8e, 0x34, 0x68, 0x88,
	0xf3, 0xd5, 0x57, 0x13, 0xd0, 0x00, 0x7a, 0xa4, 0x78, 0xa5, 0x64, 0x3e, 0xe4, 0xc2, 0xce, 0x22,
	0xba, 0xfa, 0xf9, 0x87, 0xfe, 0xf7, 0xbc, 0xc6, 0x43, 0x74, 0xf0, 0x9c, 0xcb, 0xb1, 0x09, 0x0f,
	0x13, 0x70, 0xe0, 0xc5, 0x08, 0x9f, 0xb1, 0xba, 0x39, 0x2e, 0x5e, 0xed, 0xc0, 0x7a, 0x5e, 0xff,
	0xb5, 0x91, 0xcb, 0xad, 0x6c, 0x4f, 0xca, 0x3b, 0xb0, 0x5e, 0xe1, 0x3b, 0xb4, 0xf3, 0x02, 0x41,
	0xe1, 0xee, 0x86, 0xd8, 0x0a, 0x92, 0xd3, 0x16, 0xb8, 0x6e, 0xfa, 0x40, 0xfb, 0x2b, 0xfd, 0xaa,
	0x8c, 0xce, 0x82, 0x1a, 0xe0, 0xb6, 0x50, 0x63, 0x22, 0x17, 0x5b, 0x98, 0xd6, 0x7f, 0x78, 0x44,
	0x7b, 0xb7, 0xca, 0x81, 0x37, 0x01, 0xd3, 0x0d, 0xb9, 0x9a, 0x93, 0xf3, 0x76, 0xde, 0x54, 0xde,
	0xdc, 0x7f, 0x95, 0x34, 0x99, 0x97, 0x34, 0x59, 0x94, 0x34, 0xf9, 0x2e, 0x69, 0xf2, 0xb9, 0xa4,
	0x9d, 0xf9, 0x92, 0x76, 0x16, 0x4b, 0xda, 0x79, 0x63, 0xda, 0x84, 0x2c, 0x97, 0xac, 0x0f, 0x63,
	0x1e, 0xaf, 0x66, 0x08, 0xb9, 0x1d, 0x88, 0x60, 0xc0, 0xd6, 0x03, 0x3e, 0xad, 0x6e, 0x22, 0xcc,
	0x9c, 0xf2, 0x72, 0xb7, 0x7a, 0xc3, 0xeb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0x0e, 0x95,
	0x95, 0x81, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// SubmitProposal defines a method to create new proposal given a content.
	SubmitProposal(ctx context.Context, in *types.MsgSubmitProposal, opts ...grpc.CallOption) (*types.MsgSubmitProposalResponse, error)
	// Vote defines a method to add a vote on a specific proposal.
	Vote(ctx context.Context, in *types.MsgVote, opts ...grpc.CallOption) (*types.MsgVoteResponse, error)
	// VoteWeighted defines a method to add a weighted vote on a specific proposal.
	//
	// Since: cosmos-sdk 0.43
	VoteWeighted(ctx context.Context, in *types.MsgVoteWeighted, opts ...grpc.CallOption) (*types.MsgVoteWeightedResponse, error)
	// Deposit defines a method to add deposit on a specific proposal.
	Deposit(ctx context.Context, in *types.MsgDeposit, opts ...grpc.CallOption) (*types.MsgDepositResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitProposal(ctx context.Context, in *types.MsgSubmitProposal, opts ...grpc.CallOption) (*types.MsgSubmitProposalResponse, error) {
	out := new(types.MsgSubmitProposalResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1alpha1.Msg/SubmitProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Vote(ctx context.Context, in *types.MsgVote, opts ...grpc.CallOption) (*types.MsgVoteResponse, error) {
	out := new(types.MsgVoteResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1alpha1.Msg/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) VoteWeighted(ctx context.Context, in *types.MsgVoteWeighted, opts ...grpc.CallOption) (*types.MsgVoteWeightedResponse, error) {
	out := new(types.MsgVoteWeightedResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1alpha1.Msg/VoteWeighted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Deposit(ctx context.Context, in *types.MsgDeposit, opts ...grpc.CallOption) (*types.MsgDepositResponse, error) {
	out := new(types.MsgDepositResponse)
	err := c.cc.Invoke(ctx, "/shentu.gov.v1alpha1.Msg/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SubmitProposal defines a method to create new proposal given a content.
	SubmitProposal(context.Context, *types.MsgSubmitProposal) (*types.MsgSubmitProposalResponse, error)
	// Vote defines a method to add a vote on a specific proposal.
	Vote(context.Context, *types.MsgVote) (*types.MsgVoteResponse, error)
	// VoteWeighted defines a method to add a weighted vote on a specific proposal.
	//
	// Since: cosmos-sdk 0.43
	VoteWeighted(context.Context, *types.MsgVoteWeighted) (*types.MsgVoteWeightedResponse, error)
	// Deposit defines a method to add deposit on a specific proposal.
	Deposit(context.Context, *types.MsgDeposit) (*types.MsgDepositResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitProposal(ctx context.Context, req *types.MsgSubmitProposal) (*types.MsgSubmitProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitProposal not implemented")
}
func (*UnimplementedMsgServer) Vote(ctx context.Context, req *types.MsgVote) (*types.MsgVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (*UnimplementedMsgServer) VoteWeighted(ctx context.Context, req *types.MsgVoteWeighted) (*types.MsgVoteWeightedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteWeighted not implemented")
}
func (*UnimplementedMsgServer) Deposit(ctx context.Context, req *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.MsgSubmitProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1alpha1.Msg/SubmitProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitProposal(ctx, req.(*types.MsgSubmitProposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.MsgVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1alpha1.Msg/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Vote(ctx, req.(*types.MsgVote))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_VoteWeighted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.MsgVoteWeighted)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VoteWeighted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1alpha1.Msg/VoteWeighted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VoteWeighted(ctx, req.(*types.MsgVoteWeighted))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.MsgDeposit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.gov.v1alpha1.Msg/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Deposit(ctx, req.(*types.MsgDeposit))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shentu.gov.v1alpha1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitProposal",
			Handler:    _Msg_SubmitProposal_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Msg_Vote_Handler,
		},
		{
			MethodName: "VoteWeighted",
			Handler:    _Msg_VoteWeighted_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Msg_Deposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shentu/gov/v1alpha1/tx.proto",
}
