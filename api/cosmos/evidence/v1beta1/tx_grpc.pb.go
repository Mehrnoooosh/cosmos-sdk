// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package evidencev1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// SubmitEvidence submits an arbitrary Evidence of misbehavior such as equivocation or
	// counterfactual signing.
	SubmitEvidence(ctx context.Context, in *MsgSubmitEvidence, opts ...grpc.CallOption) (*MsgSubmitEvidenceResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitEvidence(ctx context.Context, in *MsgSubmitEvidence, opts ...grpc.CallOption) (*MsgSubmitEvidenceResponse, error) {
	out := new(MsgSubmitEvidenceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.evidence.v1beta1.Msg/SubmitEvidence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// SubmitEvidence submits an arbitrary Evidence of misbehavior such as equivocation or
	// counterfactual signing.
	SubmitEvidence(context.Context, *MsgSubmitEvidence) (*MsgSubmitEvidenceResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitEvidence(context.Context, *MsgSubmitEvidence) (*MsgSubmitEvidenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitEvidence not implemented")
}
func (*UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

func RegisterMsgServer(s *grpc.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitEvidence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitEvidence)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitEvidence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.evidence.v1beta1.Msg/SubmitEvidence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitEvidence(ctx, req.(*MsgSubmitEvidence))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.evidence.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitEvidence",
			Handler:    _Msg_SubmitEvidence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/evidence/v1beta1/tx.proto",
}