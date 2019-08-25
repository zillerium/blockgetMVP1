// Code generated by protoc-gen-go. DO NOT EDIT.
// source: store/storepb/store.proto

package storepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StoreMessage struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Id                   int32    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Parent               string   `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreMessage) Reset()         { *m = StoreMessage{} }
func (m *StoreMessage) String() string { return proto.CompactTextString(m) }
func (*StoreMessage) ProtoMessage()    {}
func (*StoreMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3885e847c506538, []int{0}
}

func (m *StoreMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreMessage.Unmarshal(m, b)
}
func (m *StoreMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreMessage.Marshal(b, m, deterministic)
}
func (m *StoreMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreMessage.Merge(m, src)
}
func (m *StoreMessage) XXX_Size() int {
	return xxx_messageInfo_StoreMessage.Size(m)
}
func (m *StoreMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StoreMessage proto.InternalMessageInfo

func (m *StoreMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *StoreMessage) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *StoreMessage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StoreMessage) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

type StoreRequest struct {
	Msg                  *StoreMessage `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3885e847c506538, []int{1}
}

func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (m *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(m, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

func (m *StoreRequest) GetMsg() *StoreMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

type StoreResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreResponse) Reset()         { *m = StoreResponse{} }
func (m *StoreResponse) String() string { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()    {}
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3885e847c506538, []int{2}
}

func (m *StoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreResponse.Unmarshal(m, b)
}
func (m *StoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreResponse.Marshal(b, m, deterministic)
}
func (m *StoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResponse.Merge(m, src)
}
func (m *StoreResponse) XXX_Size() int {
	return xxx_messageInfo_StoreResponse.Size(m)
}
func (m *StoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResponse proto.InternalMessageInfo

func (m *StoreResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type LongStoreRequest struct {
	Msg                  *StoreMessage `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LongStoreRequest) Reset()         { *m = LongStoreRequest{} }
func (m *LongStoreRequest) String() string { return proto.CompactTextString(m) }
func (*LongStoreRequest) ProtoMessage()    {}
func (*LongStoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3885e847c506538, []int{3}
}

func (m *LongStoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongStoreRequest.Unmarshal(m, b)
}
func (m *LongStoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongStoreRequest.Marshal(b, m, deterministic)
}
func (m *LongStoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongStoreRequest.Merge(m, src)
}
func (m *LongStoreRequest) XXX_Size() int {
	return xxx_messageInfo_LongStoreRequest.Size(m)
}
func (m *LongStoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LongStoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LongStoreRequest proto.InternalMessageInfo

func (m *LongStoreRequest) GetMsg() *StoreMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

type LongStoreResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongStoreResponse) Reset()         { *m = LongStoreResponse{} }
func (m *LongStoreResponse) String() string { return proto.CompactTextString(m) }
func (*LongStoreResponse) ProtoMessage()    {}
func (*LongStoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3885e847c506538, []int{4}
}

func (m *LongStoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongStoreResponse.Unmarshal(m, b)
}
func (m *LongStoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongStoreResponse.Marshal(b, m, deterministic)
}
func (m *LongStoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongStoreResponse.Merge(m, src)
}
func (m *LongStoreResponse) XXX_Size() int {
	return xxx_messageInfo_LongStoreResponse.Size(m)
}
func (m *LongStoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LongStoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LongStoreResponse proto.InternalMessageInfo

func (m *LongStoreResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*StoreMessage)(nil), "example.StoreMessage")
	proto.RegisterType((*StoreRequest)(nil), "example.StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "example.StoreResponse")
	proto.RegisterType((*LongStoreRequest)(nil), "example.LongStoreRequest")
	proto.RegisterType((*LongStoreResponse)(nil), "example.LongStoreResponse")
}

func init() { proto.RegisterFile("store/storepb/store.proto", fileDescriptor_f3885e847c506538) }

var fileDescriptor_f3885e847c506538 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x6a, 0x1b, 0x32, 0x7e, 0xa0, 0x0b, 0x96, 0x34, 0xa7, 0xb2, 0x97, 0x06, 0x84,
	0x08, 0xf5, 0xa0, 0xe0, 0xcd, 0x83, 0x27, 0xbd, 0xa4, 0x37, 0x6f, 0x69, 0x3a, 0x84, 0x48, 0xbb,
	0xbb, 0xee, 0x6e, 0xc4, 0x9f, 0xe1, 0x4f, 0x96, 0xdd, 0x4c, 0x4a, 0x09, 0xe2, 0xa1, 0x97, 0x5d,
	0xde, 0xf9, 0x7a, 0xde, 0x61, 0x60, 0x66, 0xac, 0xd4, 0x78, 0xe7, 0x5f, 0xb5, 0xee, 0xfe, 0x5c,
	0x69, 0x69, 0x25, 0x8b, 0xf0, 0xbb, 0xdc, 0xa9, 0x2d, 0xf2, 0x0f, 0x38, 0x5f, 0xb9, 0xf8, 0x1b,
	0x1a, 0x53, 0xd6, 0xc8, 0x12, 0x88, 0x2a, 0x29, 0x2c, 0x0a, 0x9b, 0x04, 0xf3, 0x20, 0x8b, 0x8b,
	0x5e, 0xba, 0x4c, 0x59, 0x55, 0xb2, 0x15, 0x36, 0x09, 0xbb, 0x0c, 0x49, 0x76, 0x09, 0x61, 0xb3,
	0x49, 0x46, 0xf3, 0x20, 0x1b, 0x17, 0x61, 0xb3, 0x61, 0x53, 0x98, 0xa8, 0x52, 0xbb, 0x11, 0xa7,
	0xbe, 0x90, 0x14, 0x7f, 0x20, 0x56, 0x81, 0x9f, 0x2d, 0x1a, 0xcb, 0x16, 0x30, 0xda, 0x99, 0xda,
	0x73, 0xce, 0x96, 0x37, 0x39, 0x59, 0xca, 0x0f, 0xfd, 0x14, 0xae, 0x82, 0x2f, 0xe0, 0x82, 0x1a,
	0x8d, 0x92, 0xc2, 0xa0, 0x23, 0x68, 0x34, 0xed, 0xb6, 0x37, 0x49, 0x8a, 0x3f, 0xc1, 0xd5, 0xab,
	0x14, 0xf5, 0x71, 0x94, 0x5b, 0xb8, 0x3e, 0x68, 0xfe, 0x9f, 0xb4, 0xfc, 0x09, 0x68, 0x99, 0x15,
	0xea, 0xaf, 0xa6, 0x42, 0xf6, 0x08, 0x63, 0xaf, 0xd9, 0x00, 0x41, 0x36, 0xd2, 0xe9, 0x30, 0xdc,
	0x01, 0xf8, 0x09, 0x7b, 0x81, 0x78, 0xcf, 0x65, 0xb3, 0x7d, 0xd9, 0x70, 0x91, 0x34, 0xfd, 0x2b,
	0xd5, 0x4f, 0xc9, 0x82, 0xe7, 0xf8, 0x3d, 0xa2, 0x53, 0xaf, 0x27, 0xfe, 0xca, 0xf7, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x9d, 0x2f, 0x19, 0x37, 0x02, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreServiceClient is the client API for StoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreServiceClient interface {
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	//client streaming
	LongStore(ctx context.Context, opts ...grpc.CallOption) (StoreService_LongStoreClient, error)
}

type storeServiceClient struct {
	cc *grpc.ClientConn
}

func NewStoreServiceClient(cc *grpc.ClientConn) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/example.StoreService/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) LongStore(ctx context.Context, opts ...grpc.CallOption) (StoreService_LongStoreClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StoreService_serviceDesc.Streams[0], "/example.StoreService/LongStore", opts...)
	if err != nil {
		return nil, err
	}
	x := &storeServiceLongStoreClient{stream}
	return x, nil
}

type StoreService_LongStoreClient interface {
	Send(*LongStoreRequest) error
	CloseAndRecv() (*LongStoreResponse, error)
	grpc.ClientStream
}

type storeServiceLongStoreClient struct {
	grpc.ClientStream
}

func (x *storeServiceLongStoreClient) Send(m *LongStoreRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storeServiceLongStoreClient) CloseAndRecv() (*LongStoreResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LongStoreResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StoreServiceServer is the server API for StoreService service.
type StoreServiceServer interface {
	Store(context.Context, *StoreRequest) (*StoreResponse, error)
	//client streaming
	LongStore(StoreService_LongStoreServer) error
}

// UnimplementedStoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStoreServiceServer struct {
}

func (*UnimplementedStoreServiceServer) Store(ctx context.Context, req *StoreRequest) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedStoreServiceServer) LongStore(srv StoreService_LongStoreServer) error {
	return status.Errorf(codes.Unimplemented, "method LongStore not implemented")
}

func RegisterStoreServiceServer(s *grpc.Server, srv StoreServiceServer) {
	s.RegisterService(&_StoreService_serviceDesc, srv)
}

func _StoreService_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.StoreService/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_LongStore_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StoreServiceServer).LongStore(&storeServiceLongStoreServer{stream})
}

type StoreService_LongStoreServer interface {
	SendAndClose(*LongStoreResponse) error
	Recv() (*LongStoreRequest, error)
	grpc.ServerStream
}

type storeServiceLongStoreServer struct {
	grpc.ServerStream
}

func (x *storeServiceLongStoreServer) SendAndClose(m *LongStoreResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storeServiceLongStoreServer) Recv() (*LongStoreRequest, error) {
	m := new(LongStoreRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _StoreService_Store_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LongStore",
			Handler:       _StoreService_LongStore_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "store/storepb/store.proto",
}
