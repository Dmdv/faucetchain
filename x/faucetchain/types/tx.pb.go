// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: faucetchain/faucetchain/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ec35cfe8dcc823, []int{0}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ec35cfe8dcc823, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

type MsgRequestTokens struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Amount  uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *MsgRequestTokens) Reset()         { *m = MsgRequestTokens{} }
func (m *MsgRequestTokens) String() string { return proto.CompactTextString(m) }
func (*MsgRequestTokens) ProtoMessage()    {}
func (*MsgRequestTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ec35cfe8dcc823, []int{2}
}
func (m *MsgRequestTokens) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestTokens.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestTokens.Merge(m, src)
}
func (m *MsgRequestTokens) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestTokens.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestTokens proto.InternalMessageInfo

func (m *MsgRequestTokens) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgRequestTokens) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type MsgRequestTokensResponse struct {
}

func (m *MsgRequestTokensResponse) Reset()         { *m = MsgRequestTokensResponse{} }
func (m *MsgRequestTokensResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRequestTokensResponse) ProtoMessage()    {}
func (*MsgRequestTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ec35cfe8dcc823, []int{3}
}
func (m *MsgRequestTokensResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestTokensResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestTokensResponse.Merge(m, src)
}
func (m *MsgRequestTokensResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestTokensResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateParams)(nil), "faucetchain.faucetchain.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "faucetchain.faucetchain.MsgUpdateParamsResponse")
	proto.RegisterType((*MsgRequestTokens)(nil), "faucetchain.faucetchain.MsgRequestTokens")
	proto.RegisterType((*MsgRequestTokensResponse)(nil), "faucetchain.faucetchain.MsgRequestTokensResponse")
}

func init() { proto.RegisterFile("faucetchain/faucetchain/tx.proto", fileDescriptor_52ec35cfe8dcc823) }

var fileDescriptor_52ec35cfe8dcc823 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0x4b, 0x2c, 0x4d,
	0x4e, 0x2d, 0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x47, 0x66, 0x97, 0x54, 0xe8, 0x15, 0x14, 0xe5,
	0x97, 0xe4, 0x0b, 0x89, 0x23, 0x89, 0xea, 0x21, 0xb1, 0xa5, 0x04, 0x13, 0x73, 0x33, 0xf3, 0xf2,
	0xf5, 0xc1, 0x24, 0x44, 0xad, 0x94, 0x78, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0xb1, 0x7e, 0x6e, 0x71,
	0xba, 0x7e, 0x99, 0x21, 0x88, 0x82, 0x4a, 0x48, 0x42, 0x24, 0xe2, 0xc1, 0x3c, 0x7d, 0x08, 0x07,
	0x2a, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x11, 0x07, 0xb1, 0xa0, 0xa2, 0x2a, 0xb8, 0xdc, 0x55,
	0x90, 0x58, 0x94, 0x98, 0x0b, 0xd5, 0xab, 0x74, 0x9a, 0x91, 0x8b, 0xdf, 0xb7, 0x38, 0x3d, 0xb4,
	0x20, 0x25, 0xb1, 0x24, 0x35, 0x00, 0x2c, 0x23, 0x64, 0xc6, 0xc5, 0x99, 0x58, 0x5a, 0x92, 0x91,
	0x5f, 0x94, 0x59, 0x52, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0x24, 0x71, 0x69, 0x8b, 0xae,
	0x08, 0xd4, 0x52, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0xe2, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0xf4,
	0x20, 0x84, 0x52, 0x21, 0x27, 0x2e, 0x36, 0x88, 0xd9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46,
	0xf2, 0x7a, 0x38, 0x3c, 0xae, 0x07, 0xb1, 0xc8, 0x89, 0xf3, 0xc4, 0x3d, 0x79, 0x86, 0x15, 0xcf,
	0x37, 0x68, 0x31, 0x06, 0x41, 0x75, 0x5a, 0xd9, 0x34, 0x3d, 0xdf, 0xa0, 0x85, 0x30, 0xb3, 0xeb,
	0xf9, 0x06, 0x2d, 0x4d, 0x64, 0xc7, 0x57, 0xa0, 0x78, 0x05, 0xcd, 0xe5, 0x4a, 0x92, 0x5c, 0xe2,
	0x68, 0x42, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x4a, 0x41, 0x5c, 0x02, 0xbe, 0xc5,
	0xe9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x21, 0xf9, 0xd9, 0xa9, 0x79, 0xc5, 0x42, 0x12,
	0x5c, 0xec, 0xc9, 0x45, 0xa9, 0x89, 0x25, 0xf9, 0x45, 0x10, 0x6f, 0x06, 0xc1, 0xb8, 0x42, 0x62,
	0x5c, 0x6c, 0x89, 0xb9, 0xf9, 0xa5, 0x79, 0x25, 0x60, 0xaf, 0xb0, 0x04, 0x41, 0x79, 0x56, 0x3c,
	0x20, 0xe7, 0xc1, 0x54, 0x29, 0x49, 0x71, 0x49, 0xa0, 0x9b, 0x09, 0xb3, 0xcf, 0xe8, 0x39, 0x23,
	0x17, 0xb3, 0x6f, 0x71, 0xba, 0x50, 0x16, 0x17, 0x0f, 0x4a, 0xe0, 0x6a, 0xe0, 0x0c, 0x14, 0x34,
	0x97, 0x4b, 0x19, 0x10, 0xab, 0x12, 0x66, 0xa7, 0x50, 0x2e, 0x17, 0x2f, 0xaa, 0x07, 0x35, 0xf1,
	0x19, 0x81, 0xa2, 0x54, 0xca, 0x90, 0x68, 0xa5, 0x30, 0xeb, 0xa4, 0x58, 0x1b, 0x40, 0x51, 0xe7,
	0x64, 0x79, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xf2, 0xb8, 0x63, 0xae,
	0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x9c, 0x08, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf1, 0x76, 0xf4, 0x6a, 0x44, 0x03, 0x00, 0x00,
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
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	RequestTokens(ctx context.Context, in *MsgRequestTokens, opts ...grpc.CallOption) (*MsgRequestTokensResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/faucetchain.faucetchain.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RequestTokens(ctx context.Context, in *MsgRequestTokens, opts ...grpc.CallOption) (*MsgRequestTokensResponse, error) {
	out := new(MsgRequestTokensResponse)
	err := c.cc.Invoke(ctx, "/faucetchain.faucetchain.Msg/RequestTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	RequestTokens(context.Context, *MsgRequestTokens) (*MsgRequestTokensResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (*UnimplementedMsgServer) RequestTokens(ctx context.Context, req *MsgRequestTokens) (*MsgRequestTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestTokens not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faucetchain.faucetchain.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RequestTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestTokens)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faucetchain.faucetchain.Msg/RequestTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestTokens(ctx, req.(*MsgRequestTokens))
	}
	return interceptor(ctx, in, info, handler)
}

var Msg_serviceDesc = _Msg_serviceDesc
var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "faucetchain.faucetchain.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "RequestTokens",
			Handler:    _Msg_RequestTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "faucetchain/faucetchain/tx.proto",
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRequestTokens) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestTokens) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestTokens) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRequestTokensResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestTokensResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestTokensResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRequestTokens) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovTx(uint64(m.Amount))
	}
	return n
}

func (m *MsgRequestTokensResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRequestTokens) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRequestTokens: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestTokens: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRequestTokensResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRequestTokensResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestTokensResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
