// Code generated by protoc-gen-go. DO NOT EDIT.
// source: audit/auditpb/audit.proto

package auditpb

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

type AuditMessage struct {
	Hashcode             string   `protobuf:"bytes,1,opt,name=hashcode,proto3" json:"hashcode,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditMessage) Reset()         { *m = AuditMessage{} }
func (m *AuditMessage) String() string { return proto.CompactTextString(m) }
func (*AuditMessage) ProtoMessage()    {}
func (*AuditMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d03d52e51e6ceb1, []int{0}
}

func (m *AuditMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditMessage.Unmarshal(m, b)
}
func (m *AuditMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditMessage.Marshal(b, m, deterministic)
}
func (m *AuditMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditMessage.Merge(m, src)
}
func (m *AuditMessage) XXX_Size() int {
	return xxx_messageInfo_AuditMessage.Size(m)
}
func (m *AuditMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AuditMessage proto.InternalMessageInfo

func (m *AuditMessage) GetHashcode() string {
	if m != nil {
		return m.Hashcode
	}
	return ""
}

func (m *AuditMessage) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type AuditRequest struct {
	Msg                  *AuditMessage `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AuditRequest) Reset()         { *m = AuditRequest{} }
func (m *AuditRequest) String() string { return proto.CompactTextString(m) }
func (*AuditRequest) ProtoMessage()    {}
func (*AuditRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d03d52e51e6ceb1, []int{1}
}

func (m *AuditRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditRequest.Unmarshal(m, b)
}
func (m *AuditRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditRequest.Marshal(b, m, deterministic)
}
func (m *AuditRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditRequest.Merge(m, src)
}
func (m *AuditRequest) XXX_Size() int {
	return xxx_messageInfo_AuditRequest.Size(m)
}
func (m *AuditRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuditRequest proto.InternalMessageInfo

func (m *AuditRequest) GetMsg() *AuditMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

type AuditResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditResponse) Reset()         { *m = AuditResponse{} }
func (m *AuditResponse) String() string { return proto.CompactTextString(m) }
func (*AuditResponse) ProtoMessage()    {}
func (*AuditResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d03d52e51e6ceb1, []int{2}
}

func (m *AuditResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditResponse.Unmarshal(m, b)
}
func (m *AuditResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditResponse.Marshal(b, m, deterministic)
}
func (m *AuditResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditResponse.Merge(m, src)
}
func (m *AuditResponse) XXX_Size() int {
	return xxx_messageInfo_AuditResponse.Size(m)
}
func (m *AuditResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuditResponse proto.InternalMessageInfo

func (m *AuditResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*AuditMessage)(nil), "example.AuditMessage")
	proto.RegisterType((*AuditRequest)(nil), "example.AuditRequest")
	proto.RegisterType((*AuditResponse)(nil), "example.AuditResponse")
}

func init() { proto.RegisterFile("audit/auditpb/audit.proto", fileDescriptor_1d03d52e51e6ceb1) }

var fileDescriptor_1d03d52e51e6ceb1 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x2c, 0x4d, 0xc9,
	0x2c, 0xd1, 0x07, 0x93, 0x05, 0x49, 0x10, 0x5a, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0x3d,
	0xb5, 0x22, 0x31, 0xb7, 0x20, 0x27, 0x55, 0xc9, 0x85, 0x8b, 0xc7, 0x11, 0x24, 0xee, 0x9b, 0x5a,
	0x5c, 0x9c, 0x98, 0x9e, 0x2a, 0x24, 0xc5, 0xc5, 0x91, 0x91, 0x58, 0x9c, 0x91, 0x9c, 0x9f, 0x92,
	0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7, 0x0b, 0x49, 0x70, 0xb1, 0x27, 0x26, 0x27,
	0xe7, 0x97, 0xe6, 0x95, 0x48, 0x30, 0x81, 0xa5, 0x60, 0x5c, 0x25, 0x73, 0xa8, 0x29, 0x41, 0xa9,
	0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xea, 0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x60, 0x03, 0xb8, 0x8d,
	0x44, 0xf5, 0xa0, 0x96, 0xe9, 0x21, 0xdb, 0x14, 0x04, 0x52, 0xa1, 0xa4, 0xce, 0xc5, 0x0b, 0xd5,
	0x58, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc6, 0xc5, 0x56, 0x94, 0x5a, 0x5c, 0x9a, 0x53,
	0x02, 0xb5, 0x1d, 0xca, 0x33, 0xf2, 0x80, 0xda, 0x10, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a,
	0x64, 0xc1, 0xc5, 0x0a, 0xe6, 0x0b, 0xa1, 0x99, 0x0e, 0x75, 0x81, 0x94, 0x18, 0xba, 0x30, 0xc4,
	0x7c, 0x25, 0x06, 0x27, 0xce, 0x28, 0x76, 0x68, 0x88, 0x24, 0xb1, 0x81, 0x03, 0xc3, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xb0, 0x81, 0xb8, 0x56, 0x29, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuditServiceClient is the client API for AuditService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuditServiceClient interface {
	Audit(ctx context.Context, in *AuditRequest, opts ...grpc.CallOption) (*AuditResponse, error)
}

type auditServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuditServiceClient(cc *grpc.ClientConn) AuditServiceClient {
	return &auditServiceClient{cc}
}

func (c *auditServiceClient) Audit(ctx context.Context, in *AuditRequest, opts ...grpc.CallOption) (*AuditResponse, error) {
	out := new(AuditResponse)
	err := c.cc.Invoke(ctx, "/example.AuditService/Audit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditServiceServer is the server API for AuditService service.
type AuditServiceServer interface {
	Audit(context.Context, *AuditRequest) (*AuditResponse, error)
}

// UnimplementedAuditServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuditServiceServer struct {
}

func (*UnimplementedAuditServiceServer) Audit(ctx context.Context, req *AuditRequest) (*AuditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Audit not implemented")
}

func RegisterAuditServiceServer(s *grpc.Server, srv AuditServiceServer) {
	s.RegisterService(&_AuditService_serviceDesc, srv)
}

func _AuditService_Audit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServiceServer).Audit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.AuditService/Audit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServiceServer).Audit(ctx, req.(*AuditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuditService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.AuditService",
	HandlerType: (*AuditServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Audit",
			Handler:    _AuditService_Audit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audit/auditpb/audit.proto",
}
