// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: badges.proto

package badges

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/lissteron/cloudsuny/pkg/types/v1"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateBadgesReq struct {
	UserId               string    `protobuf:"bytes,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	Type                 string    `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Point                *v1.Point `protobuf:"bytes,3,opt,name=point,proto3" json:"point,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateBadgesReq) Reset()         { *m = CreateBadgesReq{} }
func (m *CreateBadgesReq) String() string { return proto.CompactTextString(m) }
func (*CreateBadgesReq) ProtoMessage()    {}
func (*CreateBadgesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a06625c5fe5158, []int{0}
}
func (m *CreateBadgesReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateBadgesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateBadgesReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateBadgesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBadgesReq.Merge(m, src)
}
func (m *CreateBadgesReq) XXX_Size() int {
	return m.Size()
}
func (m *CreateBadgesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBadgesReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBadgesReq proto.InternalMessageInfo

func (m *CreateBadgesReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateBadgesReq) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateBadgesReq) GetPoint() *v1.Point {
	if m != nil {
		return m.Point
	}
	return nil
}

type CreateBadgesResp struct {
	Data                 *v1.Badge `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateBadgesResp) Reset()         { *m = CreateBadgesResp{} }
func (m *CreateBadgesResp) String() string { return proto.CompactTextString(m) }
func (*CreateBadgesResp) ProtoMessage()    {}
func (*CreateBadgesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a06625c5fe5158, []int{1}
}
func (m *CreateBadgesResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateBadgesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateBadgesResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateBadgesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBadgesResp.Merge(m, src)
}
func (m *CreateBadgesResp) XXX_Size() int {
	return m.Size()
}
func (m *CreateBadgesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBadgesResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBadgesResp proto.InternalMessageInfo

func (m *CreateBadgesResp) GetData() *v1.Badge {
	if m != nil {
		return m.Data
	}
	return nil
}

type UpdateBadgesReq struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Point                *v1.Point `protobuf:"bytes,2,opt,name=point,proto3" json:"point,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateBadgesReq) Reset()         { *m = UpdateBadgesReq{} }
func (m *UpdateBadgesReq) String() string { return proto.CompactTextString(m) }
func (*UpdateBadgesReq) ProtoMessage()    {}
func (*UpdateBadgesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a06625c5fe5158, []int{2}
}
func (m *UpdateBadgesReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateBadgesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateBadgesReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateBadgesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBadgesReq.Merge(m, src)
}
func (m *UpdateBadgesReq) XXX_Size() int {
	return m.Size()
}
func (m *UpdateBadgesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBadgesReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBadgesReq proto.InternalMessageInfo

func (m *UpdateBadgesReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateBadgesReq) GetPoint() *v1.Point {
	if m != nil {
		return m.Point
	}
	return nil
}

type UpdateBadgesResp struct {
	Data                 *v1.Badge `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateBadgesResp) Reset()         { *m = UpdateBadgesResp{} }
func (m *UpdateBadgesResp) String() string { return proto.CompactTextString(m) }
func (*UpdateBadgesResp) ProtoMessage()    {}
func (*UpdateBadgesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a06625c5fe5158, []int{3}
}
func (m *UpdateBadgesResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateBadgesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateBadgesResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateBadgesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBadgesResp.Merge(m, src)
}
func (m *UpdateBadgesResp) XXX_Size() int {
	return m.Size()
}
func (m *UpdateBadgesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBadgesResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBadgesResp proto.InternalMessageInfo

func (m *UpdateBadgesResp) GetData() *v1.Badge {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateBadgesReq)(nil), "badges.v1.CreateBadgesReq")
	proto.RegisterType((*CreateBadgesResp)(nil), "badges.v1.CreateBadgesResp")
	proto.RegisterType((*UpdateBadgesReq)(nil), "badges.v1.UpdateBadgesReq")
	proto.RegisterType((*UpdateBadgesResp)(nil), "badges.v1.UpdateBadgesResp")
}

func init() { proto.RegisterFile("badges.proto", fileDescriptor_94a06625c5fe5158) }

var fileDescriptor_94a06625c5fe5158 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x4b, 0xf3, 0x30,
	0x18, 0x7f, 0xd3, 0x77, 0x56, 0x16, 0xc5, 0x49, 0xf0, 0x50, 0xab, 0xd4, 0x51, 0x11, 0xc6, 0x0e,
	0x2d, 0x9d, 0x07, 0x61, 0xc7, 0x79, 0xf1, 0x28, 0x05, 0x2f, 0x5e, 0x24, 0x5d, 0x62, 0x0d, 0xcc,
	0x24, 0x36, 0xe9, 0x60, 0x57, 0xbf, 0x82, 0x17, 0x3f, 0x92, 0xc7, 0x81, 0x5f, 0x40, 0xa6, 0x1f,
	0x44, 0x96, 0x6c, 0x93, 0x0c, 0x86, 0xde, 0x9a, 0xe7, 0xf9, 0xfd, 0x7b, 0x9e, 0xa7, 0x70, 0xb7,
	0xc0, 0xa4, 0xa4, 0x2a, 0x91, 0x95, 0xd0, 0x02, 0x35, 0x17, 0xaf, 0x71, 0x16, 0x1e, 0x97, 0x42,
	0x94, 0x23, 0x9a, 0x62, 0xc9, 0x52, 0xcc, 0xb9, 0xd0, 0x58, 0x33, 0xc1, 0x17, 0xc0, 0x30, 0x2b,
	0x99, 0x7e, 0xa8, 0x8b, 0x64, 0x28, 0x1e, 0xd3, 0x11, 0x53, 0x4a, 0xd3, 0x4a, 0xf0, 0x74, 0x38,
	0x12, 0x35, 0x51, 0x35, 0x9f, 0x18, 0xd6, 0x38, 0x4b, 0xf5, 0x44, 0x2e, 0xb5, 0xe3, 0x7b, 0xd8,
	0xba, 0xac, 0x28, 0xd6, 0x74, 0x60, 0x3c, 0x72, 0xfa, 0x84, 0x02, 0xb8, 0x5d, 0x2b, 0x5a, 0xdd,
	0x31, 0x12, 0x80, 0x36, 0xe8, 0x34, 0xf3, 0xe5, 0x13, 0x21, 0xd8, 0x98, 0x73, 0x03, 0xcf, 0x94,
	0xcd, 0x37, 0x3a, 0x83, 0x5b, 0x52, 0x30, 0xae, 0x83, 0xff, 0x6d, 0xd0, 0xd9, 0xe9, 0xb5, 0x12,
	0xab, 0x3e, 0xce, 0x92, 0xeb, 0x79, 0x39, 0xb7, 0xdd, 0xf8, 0x02, 0xee, 0xbb, 0x3e, 0x4a, 0xa2,
	0x53, 0xd8, 0x20, 0x58, 0x63, 0xe3, 0xe2, 0x30, 0x0d, 0x26, 0x37, 0xcd, 0xf8, 0x0a, 0xb6, 0x6e,
	0x24, 0x71, 0x02, 0xee, 0x41, 0x6f, 0x95, 0xcd, 0x63, 0xe4, 0x27, 0x82, 0xf7, 0x5b, 0x04, 0x57,
	0xe9, 0x8f, 0x11, 0x7a, 0x53, 0x00, 0x7d, 0xcb, 0x41, 0x05, 0xf4, 0xed, 0x18, 0x28, 0x4c, 0x56,
	0x57, 0x49, 0xd6, 0x36, 0x18, 0x1e, 0x6d, 0xec, 0x29, 0x19, 0x9f, 0x3c, 0xbf, 0x7f, 0xbd, 0x78,
	0x87, 0xf1, 0xc1, 0xf2, 0x1c, 0x06, 0x9b, 0x0e, 0x0d, 0xae, 0x0f, 0xba, 0x73, 0x0f, 0x9b, 0xd3,
	0xf1, 0x58, 0x5b, 0x82, 0xe3, 0xb1, 0x3e, 0xd6, 0x26, 0x8f, 0xda, 0xe0, 0xfa, 0xa0, 0x3b, 0x08,
	0xde, 0x66, 0x11, 0x98, 0xce, 0x22, 0xf0, 0x31, 0x8b, 0xc0, 0xeb, 0x67, 0xf4, 0xef, 0xd6, 0xb7,
	0x72, 0x85, 0x6f, 0xfe, 0x8b, 0xf3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x73, 0x9e, 0xb9,
	0x83, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BadgesClient is the client API for Badges service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BadgesClient interface {
	Create(ctx context.Context, in *CreateBadgesReq, opts ...grpc.CallOption) (*CreateBadgesResp, error)
	Update(ctx context.Context, in *UpdateBadgesReq, opts ...grpc.CallOption) (*UpdateBadgesResp, error)
}

type badgesClient struct {
	cc *grpc.ClientConn
}

func NewBadgesClient(cc *grpc.ClientConn) BadgesClient {
	return &badgesClient{cc}
}

func (c *badgesClient) Create(ctx context.Context, in *CreateBadgesReq, opts ...grpc.CallOption) (*CreateBadgesResp, error) {
	out := new(CreateBadgesResp)
	err := c.cc.Invoke(ctx, "/badges.v1.Badges/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgesClient) Update(ctx context.Context, in *UpdateBadgesReq, opts ...grpc.CallOption) (*UpdateBadgesResp, error) {
	out := new(UpdateBadgesResp)
	err := c.cc.Invoke(ctx, "/badges.v1.Badges/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BadgesServer is the server API for Badges service.
type BadgesServer interface {
	Create(context.Context, *CreateBadgesReq) (*CreateBadgesResp, error)
	Update(context.Context, *UpdateBadgesReq) (*UpdateBadgesResp, error)
}

// UnimplementedBadgesServer can be embedded to have forward compatible implementations.
type UnimplementedBadgesServer struct {
}

func (*UnimplementedBadgesServer) Create(ctx context.Context, req *CreateBadgesReq) (*CreateBadgesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedBadgesServer) Update(ctx context.Context, req *UpdateBadgesReq) (*UpdateBadgesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterBadgesServer(s *grpc.Server, srv BadgesServer) {
	s.RegisterService(&_Badges_serviceDesc, srv)
}

func _Badges_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBadgesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/badges.v1.Badges/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgesServer).Create(ctx, req.(*CreateBadgesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Badges_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBadgesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/badges.v1.Badges/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgesServer).Update(ctx, req.(*UpdateBadgesReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Badges_serviceDesc = grpc.ServiceDesc{
	ServiceName: "badges.v1.Badges",
	HandlerType: (*BadgesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Badges_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Badges_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "badges.proto",
}

func (m *CreateBadgesReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateBadgesReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateBadgesReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Point != nil {
		{
			size, err := m.Point.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBadges(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintBadges(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UserId) > 0 {
		i -= len(m.UserId)
		copy(dAtA[i:], m.UserId)
		i = encodeVarintBadges(dAtA, i, uint64(len(m.UserId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreateBadgesResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateBadgesResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateBadgesResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBadges(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateBadgesReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateBadgesReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateBadgesReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Point != nil {
		{
			size, err := m.Point.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBadges(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintBadges(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateBadgesResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateBadgesResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateBadgesResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBadges(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBadges(dAtA []byte, offset int, v uint64) int {
	offset -= sovBadges(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreateBadgesReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovBadges(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovBadges(uint64(l))
	}
	if m.Point != nil {
		l = m.Point.Size()
		n += 1 + l + sovBadges(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CreateBadgesResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovBadges(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UpdateBadgesReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovBadges(uint64(l))
	}
	if m.Point != nil {
		l = m.Point.Size()
		n += 1 + l + sovBadges(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UpdateBadgesResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovBadges(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBadges(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBadges(x uint64) (n int) {
	return sovBadges(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateBadgesReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBadges
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
			return fmt.Errorf("proto: CreateBadgesReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateBadgesReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadges
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
				return ErrInvalidLengthBadges
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBadges
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadges
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
				return ErrInvalidLengthBadges
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBadges
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Point", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadges
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
				return ErrInvalidLengthBadges
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBadges
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Point == nil {
				m.Point = &v1.Point{}
			}
			if err := m.Point.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBadges(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBadges
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBadges
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateBadgesResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBadges
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
			return fmt.Errorf("proto: CreateBadgesResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateBadgesResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadges
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
				return ErrInvalidLengthBadges
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBadges
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &v1.Badge{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBadges(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBadges
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBadges
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpdateBadgesReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBadges
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
			return fmt.Errorf("proto: UpdateBadgesReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateBadgesReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadges
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
				return ErrInvalidLengthBadges
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBadges
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Point", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadges
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
				return ErrInvalidLengthBadges
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBadges
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Point == nil {
				m.Point = &v1.Point{}
			}
			if err := m.Point.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBadges(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBadges
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBadges
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpdateBadgesResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBadges
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
			return fmt.Errorf("proto: UpdateBadgesResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateBadgesResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadges
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
				return ErrInvalidLengthBadges
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBadges
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &v1.Badge{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBadges(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBadges
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBadges
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBadges(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBadges
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
					return 0, ErrIntOverflowBadges
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
					return 0, ErrIntOverflowBadges
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
				return 0, ErrInvalidLengthBadges
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBadges
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBadges
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBadges        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBadges          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBadges = fmt.Errorf("proto: unexpected end of group")
)
