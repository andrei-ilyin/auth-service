// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth_service

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

type Status_Code int32

const (
	Status_UNKNOWN         Status_Code = 0
	Status_OK              Status_Code = 1
	Status_ACCESS_DENIED   Status_Code = 2
	Status_INVALID_SESSION Status_Code = 3
)

var Status_Code_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "ACCESS_DENIED",
	3: "INVALID_SESSION",
}

var Status_Code_value = map[string]int32{
	"UNKNOWN":         0,
	"OK":              1,
	"ACCESS_DENIED":   2,
	"INVALID_SESSION": 3,
}

func (x Status_Code) String() string {
	return proto.EnumName(Status_Code_name, int32(x))
}

func (Status_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2, 0}
}

type Credentials struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credentials) Reset()         { *m = Credentials{} }
func (m *Credentials) String() string { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()    {}
func (*Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *Credentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Credentials.Unmarshal(m, b)
}
func (m *Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Credentials.Marshal(b, m, deterministic)
}
func (m *Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credentials.Merge(m, src)
}
func (m *Credentials) XXX_Size() int {
	return xxx_messageInfo_Credentials.Size(m)
}
func (m *Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_Credentials proto.InternalMessageInfo

func (m *Credentials) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Credentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Cookie struct {
	SessionId            int64    `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	HashKey              string   `protobuf:"bytes,2,opt,name=hash_key,json=hashKey,proto3" json:"hash_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cookie) Reset()         { *m = Cookie{} }
func (m *Cookie) String() string { return proto.CompactTextString(m) }
func (*Cookie) ProtoMessage()    {}
func (*Cookie) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *Cookie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cookie.Unmarshal(m, b)
}
func (m *Cookie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cookie.Marshal(b, m, deterministic)
}
func (m *Cookie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cookie.Merge(m, src)
}
func (m *Cookie) XXX_Size() int {
	return xxx_messageInfo_Cookie.Size(m)
}
func (m *Cookie) XXX_DiscardUnknown() {
	xxx_messageInfo_Cookie.DiscardUnknown(m)
}

var xxx_messageInfo_Cookie proto.InternalMessageInfo

func (m *Cookie) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *Cookie) GetHashKey() string {
	if m != nil {
		return m.HashKey
	}
	return ""
}

type Status struct {
	Code                 Status_Code `protobuf:"varint,1,opt,name=code,proto3,enum=auth.Status_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() Status_Code {
	if m != nil {
		return m.Code
	}
	return Status_UNKNOWN
}

type LoginRequest struct {
	Credentials          *Credentials `protobuf:"bytes,1,opt,name=credentials,proto3" json:"credentials,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetCredentials() *Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

type LoginResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Cookie               *Cookie  `protobuf:"bytes,2,opt,name=cookie,proto3" json:"cookie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *LoginResponse) GetCookie() *Cookie {
	if m != nil {
		return m.Cookie
	}
	return nil
}

type LogoutRequest struct {
	Cookie               *Cookie  `protobuf:"bytes,1,opt,name=cookie,proto3" json:"cookie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetCookie() *Cookie {
	if m != nil {
		return m.Cookie
	}
	return nil
}

type LogoutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{6}
}

func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (m *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(m, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

type ValidationRequest struct {
	Cookie               *Cookie  `protobuf:"bytes,1,opt,name=cookie,proto3" json:"cookie,omitempty"`
	Resource             string   `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidationRequest) Reset()         { *m = ValidationRequest{} }
func (m *ValidationRequest) String() string { return proto.CompactTextString(m) }
func (*ValidationRequest) ProtoMessage()    {}
func (*ValidationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{7}
}

func (m *ValidationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidationRequest.Unmarshal(m, b)
}
func (m *ValidationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidationRequest.Marshal(b, m, deterministic)
}
func (m *ValidationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationRequest.Merge(m, src)
}
func (m *ValidationRequest) XXX_Size() int {
	return xxx_messageInfo_ValidationRequest.Size(m)
}
func (m *ValidationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationRequest proto.InternalMessageInfo

func (m *ValidationRequest) GetCookie() *Cookie {
	if m != nil {
		return m.Cookie
	}
	return nil
}

func (m *ValidationRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

type ValidationResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidationResponse) Reset()         { *m = ValidationResponse{} }
func (m *ValidationResponse) String() string { return proto.CompactTextString(m) }
func (*ValidationResponse) ProtoMessage()    {}
func (*ValidationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{8}
}

func (m *ValidationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidationResponse.Unmarshal(m, b)
}
func (m *ValidationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidationResponse.Marshal(b, m, deterministic)
}
func (m *ValidationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationResponse.Merge(m, src)
}
func (m *ValidationResponse) XXX_Size() int {
	return xxx_messageInfo_ValidationResponse.Size(m)
}
func (m *ValidationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationResponse proto.InternalMessageInfo

func (m *ValidationResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterEnum("auth.Status_Code", Status_Code_name, Status_Code_value)
	proto.RegisterType((*Credentials)(nil), "auth.Credentials")
	proto.RegisterType((*Cookie)(nil), "auth.Cookie")
	proto.RegisterType((*Status)(nil), "auth.Status")
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "auth.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "auth.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "auth.LogoutResponse")
	proto.RegisterType((*ValidationRequest)(nil), "auth.ValidationRequest")
	proto.RegisterType((*ValidationResponse)(nil), "auth.ValidationResponse")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xdf, 0x6b, 0xd3, 0x50,
	0x14, 0x5e, 0xba, 0x9a, 0xb5, 0xa7, 0xeb, 0x4c, 0xcf, 0x04, 0x6b, 0x45, 0x90, 0x30, 0xc5, 0x17,
	0x3b, 0xe8, 0xd8, 0x8b, 0x2f, 0xd2, 0xa5, 0x15, 0x42, 0x47, 0x0a, 0x09, 0x9b, 0xa0, 0x0f, 0xe1,
	0x2e, 0x39, 0x2c, 0x97, 0xb5, 0xb9, 0x35, 0xf7, 0x46, 0xe9, 0xdf, 0xe6, 0x3f, 0x27, 0xb9, 0x49,
	0xd6, 0x8c, 0x81, 0xd0, 0xb7, 0x9c, 0xf3, 0xfd, 0xe0, 0xbb, 0xdf, 0x21, 0x00, 0x2c, 0x57, 0xc9,
	0x78, 0x93, 0x09, 0x25, 0xb0, 0x5d, 0x7c, 0xdb, 0xdf, 0xa0, 0xe7, 0x64, 0x14, 0x53, 0xaa, 0x38,
	0x5b, 0x49, 0x7c, 0x0b, 0xdd, 0x5c, 0x52, 0x16, 0xa6, 0x6c, 0x4d, 0x43, 0xe3, 0xbd, 0xf1, 0xa9,
	0xeb, 0x77, 0x8a, 0x85, 0xc7, 0xd6, 0x84, 0x23, 0xe8, 0x6c, 0x98, 0x94, 0x7f, 0x44, 0x16, 0x0f,
	0x5b, 0x25, 0x56, 0xcf, 0xf6, 0x15, 0x98, 0x8e, 0x10, 0x0f, 0x9c, 0xf0, 0x1d, 0x80, 0x24, 0x29,
	0xb9, 0x48, 0x43, 0x1e, 0x6b, 0x8f, 0x43, 0xbf, 0x5b, 0x6d, 0xdc, 0x18, 0xdf, 0x40, 0x27, 0x61,
	0x32, 0x09, 0x1f, 0x68, 0x5b, 0x99, 0x1c, 0x15, 0xf3, 0x82, 0xb6, 0xb6, 0x02, 0x33, 0x50, 0x4c,
	0xe5, 0x12, 0x3f, 0x40, 0x3b, 0x12, 0x71, 0x99, 0xe0, 0x64, 0x32, 0x18, 0xeb, 0xd8, 0x25, 0x36,
	0x76, 0x44, 0x4c, 0xbe, 0x86, 0x6d, 0x07, 0xda, 0xc5, 0x84, 0x3d, 0x38, 0xba, 0xf1, 0x16, 0xde,
	0xf2, 0xbb, 0x67, 0x1d, 0xa0, 0x09, 0xad, 0xe5, 0xc2, 0x32, 0x70, 0x00, 0xfd, 0xa9, 0xe3, 0xcc,
	0x83, 0x20, 0x9c, 0xcd, 0x3d, 0x77, 0x3e, 0xb3, 0x5a, 0x78, 0x0a, 0x2f, 0x5d, 0xef, 0x76, 0x7a,
	0xed, 0xce, 0xc2, 0x60, 0x1e, 0x04, 0xee, 0xd2, 0xb3, 0x0e, 0x6d, 0x07, 0x8e, 0xaf, 0xc5, 0x3d,
	0x4f, 0x7d, 0xfa, 0x95, 0x93, 0x54, 0x78, 0x01, 0xbd, 0x68, 0xd7, 0x88, 0x8e, 0xd0, 0xab, 0x23,
	0x34, 0xaa, 0xf2, 0x9b, 0x2c, 0xfb, 0x27, 0xf4, 0x2b, 0x13, 0xb9, 0x11, 0xa9, 0x24, 0x3c, 0x03,
	0x53, 0xea, 0xbc, 0x95, 0xc1, 0x71, 0xf3, 0x0d, 0x7e, 0x85, 0x15, 0xac, 0x48, 0xb7, 0xa6, 0xab,
	0x78, 0x64, 0x95, 0x4d, 0xfa, 0x15, 0x66, 0x5f, 0x6a, 0x73, 0x91, 0xab, 0x3a, 0xe2, 0x4e, 0x66,
	0xfc, 0x47, 0x66, 0xc1, 0x49, 0x2d, 0x2b, 0x43, 0xd9, 0x37, 0x30, 0xb8, 0x65, 0x2b, 0x1e, 0x33,
	0xc5, 0x45, 0xba, 0x97, 0x59, 0x71, 0xfb, 0x8c, 0xa4, 0xc8, 0xb3, 0x88, 0xea, 0xdb, 0xd7, 0xb3,
	0xfd, 0x05, 0xb0, 0x69, 0xbb, 0x4f, 0x03, 0x93, 0xbf, 0x06, 0xf4, 0xa7, 0xb9, 0x4a, 0x8a, 0x22,
	0x23, 0xa6, 0x44, 0x86, 0x13, 0x78, 0xa1, 0xab, 0x44, 0x2c, 0x05, 0xcd, 0xe3, 0x8c, 0x4e, 0x9f,
	0xec, 0xaa, 0x67, 0x1d, 0xe0, 0x25, 0x98, 0xe5, 0x53, 0x71, 0x47, 0xd8, 0xf5, 0x35, 0x7a, 0xf5,
	0x74, 0xf9, 0x28, 0xfb, 0x0a, 0x9d, 0x2a, 0x38, 0xe1, 0xeb, 0x92, 0xf3, 0xac, 0x9f, 0xd1, 0xf0,
	0x39, 0x50, 0x1b, 0x5c, 0x7d, 0xfc, 0x71, 0x76, 0xcf, 0x55, 0x92, 0xdf, 0x8d, 0x23, 0xb1, 0x3e,
	0x67, 0x69, 0x9c, 0x11, 0x0f, 0xf9, 0x6a, 0xcb, 0xd3, 0xf3, 0x42, 0xf4, 0x59, 0x52, 0xf6, 0x9b,
	0x47, 0x74, 0x67, 0xea, 0x5f, 0xee, 0xe2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x32, 0x85,
	0xf9, 0x80, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticatorClient is the client API for Authenticator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticatorClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	Validate(ctx context.Context, in *ValidationRequest, opts ...grpc.CallOption) (*ValidationResponse, error)
}

type authenticatorClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticatorClient(cc *grpc.ClientConn) AuthenticatorClient {
	return &authenticatorClient{cc}
}

func (c *authenticatorClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth.Authenticator/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/auth.Authenticator/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorClient) Validate(ctx context.Context, in *ValidationRequest, opts ...grpc.CallOption) (*ValidationResponse, error) {
	out := new(ValidationResponse)
	err := c.cc.Invoke(ctx, "/auth.Authenticator/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticatorServer is the server API for Authenticator service.
type AuthenticatorServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Validate(context.Context, *ValidationRequest) (*ValidationResponse, error)
}

// UnimplementedAuthenticatorServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticatorServer struct {
}

func (*UnimplementedAuthenticatorServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthenticatorServer) Logout(ctx context.Context, req *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (*UnimplementedAuthenticatorServer) Validate(ctx context.Context, req *ValidationRequest) (*ValidationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}

func RegisterAuthenticatorServer(s *grpc.Server, srv AuthenticatorServer) {
	s.RegisterService(&_Authenticator_serviceDesc, srv)
}

func _Authenticator_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Authenticator/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticator_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Authenticator/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticator_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Authenticator/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).Validate(ctx, req.(*ValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authenticator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Authenticator",
	HandlerType: (*AuthenticatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Authenticator_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Authenticator_Logout_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _Authenticator_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}