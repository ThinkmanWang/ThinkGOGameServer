// Code generated by protoc-gen-go. DO NOT EDIT.
// source: GamePkg.proto

package serversdk

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type HeadType int32

const (
	HeadType_LOGIN_REQUEST      HeadType = 1
	HeadType_LOGIN_RESPONSE     HeadType = 2
	HeadType_LOGOUT_REQUEST     HeadType = 3
	HeadType_LOGOUT_RESPONSE    HeadType = 4
	HeadType_NOTIFICATION       HeadType = 5
	HeadType_HEARTBEAT_REQUEST  HeadType = 6
	HeadType_HEARTBEAT_RESPONSE HeadType = 7
	HeadType_SEND_TO_SERVER     HeadType = 8
	HeadType_SEND_TO_CLIENT     HeadType = 9
)

var HeadType_name = map[int32]string{
	1: "LOGIN_REQUEST",
	2: "LOGIN_RESPONSE",
	3: "LOGOUT_REQUEST",
	4: "LOGOUT_RESPONSE",
	5: "NOTIFICATION",
	6: "HEARTBEAT_REQUEST",
	7: "HEARTBEAT_RESPONSE",
	8: "SEND_TO_SERVER",
	9: "SEND_TO_CLIENT",
}

var HeadType_value = map[string]int32{
	"LOGIN_REQUEST":      1,
	"LOGIN_RESPONSE":     2,
	"LOGOUT_REQUEST":     3,
	"LOGOUT_RESPONSE":    4,
	"NOTIFICATION":       5,
	"HEARTBEAT_REQUEST":  6,
	"HEARTBEAT_RESPONSE": 7,
	"SEND_TO_SERVER":     8,
	"SEND_TO_CLIENT":     9,
}

func (x HeadType) Enum() *HeadType {
	p := new(HeadType)
	*p = x
	return p
}

func (x HeadType) String() string {
	return proto.EnumName(HeadType_name, int32(x))
}

func (x *HeadType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HeadType_value, data, "HeadType")
	if err != nil {
		return err
	}
	*x = HeadType(value)
	return nil
}

func (HeadType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6b3f2faac6163527, []int{0}
}

type LoginRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b3f2faac6163527, []int{0}
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

type LoginResponse struct {
	Code                 *int32   `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Msg                  *string  `protobuf:"bytes,2,req,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b3f2faac6163527, []int{1}
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

func (m *LoginResponse) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *LoginResponse) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

type LogoutRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b3f2faac6163527, []int{2}
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

type LogoutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b3f2faac6163527, []int{3}
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

type Notification struct {
	Msg                  *string  `protobuf:"bytes,1,req,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b3f2faac6163527, []int{4}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

type HeartbeatRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatRequest) Reset()         { *m = HeartbeatRequest{} }
func (m *HeartbeatRequest) String() string { return proto.CompactTextString(m) }
func (*HeartbeatRequest) ProtoMessage()    {}
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b3f2faac6163527, []int{5}
}

func (m *HeartbeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatRequest.Unmarshal(m, b)
}
func (m *HeartbeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatRequest.Marshal(b, m, deterministic)
}
func (m *HeartbeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatRequest.Merge(m, src)
}
func (m *HeartbeatRequest) XXX_Size() int {
	return xxx_messageInfo_HeartbeatRequest.Size(m)
}
func (m *HeartbeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatRequest proto.InternalMessageInfo

type HeartbeatResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatResponse) Reset()         { *m = HeartbeatResponse{} }
func (m *HeartbeatResponse) String() string { return proto.CompactTextString(m) }
func (*HeartbeatResponse) ProtoMessage()    {}
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b3f2faac6163527, []int{6}
}

func (m *HeartbeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatResponse.Unmarshal(m, b)
}
func (m *HeartbeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatResponse.Marshal(b, m, deterministic)
}
func (m *HeartbeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatResponse.Merge(m, src)
}
func (m *HeartbeatResponse) XXX_Size() int {
	return xxx_messageInfo_HeartbeatResponse.Size(m)
}
func (m *HeartbeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatResponse proto.InternalMessageInfo

type SendToServer struct {
	Data                 []byte   `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendToServer) Reset()         { *m = SendToServer{} }
func (m *SendToServer) String() string { return proto.CompactTextString(m) }
func (*SendToServer) ProtoMessage()    {}
func (*SendToServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b3f2faac6163527, []int{7}
}

func (m *SendToServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendToServer.Unmarshal(m, b)
}
func (m *SendToServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendToServer.Marshal(b, m, deterministic)
}
func (m *SendToServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendToServer.Merge(m, src)
}
func (m *SendToServer) XXX_Size() int {
	return xxx_messageInfo_SendToServer.Size(m)
}
func (m *SendToServer) XXX_DiscardUnknown() {
	xxx_messageInfo_SendToServer.DiscardUnknown(m)
}

var xxx_messageInfo_SendToServer proto.InternalMessageInfo

func (m *SendToServer) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SendToClient struct {
	Data                 []byte   `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendToClient) Reset()         { *m = SendToClient{} }
func (m *SendToClient) String() string { return proto.CompactTextString(m) }
func (*SendToClient) ProtoMessage()    {}
func (*SendToClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b3f2faac6163527, []int{8}
}

func (m *SendToClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendToClient.Unmarshal(m, b)
}
func (m *SendToClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendToClient.Marshal(b, m, deterministic)
}
func (m *SendToClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendToClient.Merge(m, src)
}
func (m *SendToClient) XXX_Size() int {
	return xxx_messageInfo_SendToClient.Size(m)
}
func (m *SendToClient) XXX_DiscardUnknown() {
	xxx_messageInfo_SendToClient.DiscardUnknown(m)
}

var xxx_messageInfo_SendToClient proto.InternalMessageInfo

func (m *SendToClient) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GamePkg struct {
	Type                 *HeadType          `protobuf:"varint,1,req,name=type,enum=serversdk.HeadType" json:"type,omitempty"`
	AppId                *int32             `protobuf:"varint,2,req,name=appId" json:"appId,omitempty"`
	Uid                  *string            `protobuf:"bytes,3,req,name=uid" json:"uid,omitempty"`
	Timestamp            *int64             `protobuf:"varint,4,req,name=timestamp" json:"timestamp,omitempty"`
	LoginRequest         *LoginRequest      `protobuf:"bytes,5,opt,name=loginRequest" json:"loginRequest,omitempty"`
	LoginResponse        *LoginResponse     `protobuf:"bytes,6,opt,name=loginResponse" json:"loginResponse,omitempty"`
	LogoutRequest        *LogoutRequest     `protobuf:"bytes,7,opt,name=logoutRequest" json:"logoutRequest,omitempty"`
	LogoutResponse       *LogoutResponse    `protobuf:"bytes,8,opt,name=logoutResponse" json:"logoutResponse,omitempty"`
	Notification         *Notification      `protobuf:"bytes,9,opt,name=notification" json:"notification,omitempty"`
	HeartbeatRequest     *HeartbeatRequest  `protobuf:"bytes,10,opt,name=heartbeatRequest" json:"heartbeatRequest,omitempty"`
	HeartbeatResponse    *HeartbeatResponse `protobuf:"bytes,11,opt,name=heartbeatResponse" json:"heartbeatResponse,omitempty"`
	SendToServer         *SendToServer      `protobuf:"bytes,12,opt,name=sendToServer" json:"sendToServer,omitempty"`
	SendToClient         *SendToClient      `protobuf:"bytes,13,opt,name=sendToClient" json:"sendToClient,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GamePkg) Reset()         { *m = GamePkg{} }
func (m *GamePkg) String() string { return proto.CompactTextString(m) }
func (*GamePkg) ProtoMessage()    {}
func (*GamePkg) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b3f2faac6163527, []int{9}
}

func (m *GamePkg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GamePkg.Unmarshal(m, b)
}
func (m *GamePkg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GamePkg.Marshal(b, m, deterministic)
}
func (m *GamePkg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GamePkg.Merge(m, src)
}
func (m *GamePkg) XXX_Size() int {
	return xxx_messageInfo_GamePkg.Size(m)
}
func (m *GamePkg) XXX_DiscardUnknown() {
	xxx_messageInfo_GamePkg.DiscardUnknown(m)
}

var xxx_messageInfo_GamePkg proto.InternalMessageInfo

func (m *GamePkg) GetType() HeadType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HeadType_LOGIN_REQUEST
}

func (m *GamePkg) GetAppId() int32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func (m *GamePkg) GetUid() string {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return ""
}

func (m *GamePkg) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *GamePkg) GetLoginRequest() *LoginRequest {
	if m != nil {
		return m.LoginRequest
	}
	return nil
}

func (m *GamePkg) GetLoginResponse() *LoginResponse {
	if m != nil {
		return m.LoginResponse
	}
	return nil
}

func (m *GamePkg) GetLogoutRequest() *LogoutRequest {
	if m != nil {
		return m.LogoutRequest
	}
	return nil
}

func (m *GamePkg) GetLogoutResponse() *LogoutResponse {
	if m != nil {
		return m.LogoutResponse
	}
	return nil
}

func (m *GamePkg) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

func (m *GamePkg) GetHeartbeatRequest() *HeartbeatRequest {
	if m != nil {
		return m.HeartbeatRequest
	}
	return nil
}

func (m *GamePkg) GetHeartbeatResponse() *HeartbeatResponse {
	if m != nil {
		return m.HeartbeatResponse
	}
	return nil
}

func (m *GamePkg) GetSendToServer() *SendToServer {
	if m != nil {
		return m.SendToServer
	}
	return nil
}

func (m *GamePkg) GetSendToClient() *SendToClient {
	if m != nil {
		return m.SendToClient
	}
	return nil
}

func init() {
	proto.RegisterEnum("serversdk.HeadType", HeadType_name, HeadType_value)
	proto.RegisterType((*LoginRequest)(nil), "serversdk.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "serversdk.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "serversdk.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "serversdk.LogoutResponse")
	proto.RegisterType((*Notification)(nil), "serversdk.Notification")
	proto.RegisterType((*HeartbeatRequest)(nil), "serversdk.HeartbeatRequest")
	proto.RegisterType((*HeartbeatResponse)(nil), "serversdk.HeartbeatResponse")
	proto.RegisterType((*SendToServer)(nil), "serversdk.SendToServer")
	proto.RegisterType((*SendToClient)(nil), "serversdk.SendToClient")
	proto.RegisterType((*GamePkg)(nil), "serversdk.GamePkg")
}

func init() {
	proto.RegisterFile("GamePkg.proto", fileDescriptor_6b3f2faac6163527)
}

var fileDescriptor_6b3f2faac6163527 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x8f, 0xda, 0x30,
	0x14, 0x14, 0x1f, 0x59, 0xe0, 0x6d, 0x60, 0xcd, 0xa3, 0x1f, 0xae, 0xba, 0x07, 0xc4, 0xa5, 0xa8,
	0x07, 0x0e, 0x2b, 0xf5, 0x54, 0xa9, 0x12, 0xa5, 0x29, 0xa4, 0x42, 0xc9, 0xd6, 0xc9, 0xf6, 0x8a,
	0xd2, 0x8d, 0xcb, 0x46, 0x0b, 0x49, 0x4a, 0x4c, 0xa5, 0xfd, 0x7d, 0xed, 0x0f, 0xab, 0xe2, 0x04,
	0xe2, 0x10, 0xf6, 0x66, 0x8f, 0x67, 0xc6, 0x1e, 0xbf, 0x81, 0xee, 0xdc, 0xdb, 0xf2, 0xdb, 0xc7,
	0xf5, 0x24, 0xde, 0x45, 0x22, 0xc2, 0x4e, 0xc2, 0x77, 0x7f, 0xf8, 0x2e, 0xf1, 0x1f, 0x47, 0x3d,
	0xd0, 0x97, 0xd1, 0x3a, 0x08, 0x19, 0xff, 0xbd, 0xe7, 0x89, 0x18, 0x7d, 0x80, 0x6e, 0xbe, 0x4f,
	0xe2, 0x28, 0x4c, 0x38, 0x22, 0x34, 0xef, 0x23, 0x9f, 0xd3, 0xda, 0xb0, 0x3e, 0xd6, 0x98, 0x5c,
	0x23, 0x81, 0xc6, 0x36, 0x59, 0xd3, 0xfa, 0xb0, 0x3e, 0xee, 0xb0, 0x74, 0x39, 0xba, 0x92, 0xb2,
	0x68, 0x2f, 0x0e, 0x3e, 0x04, 0x7a, 0x07, 0x20, 0x33, 0x1a, 0x0d, 0x41, 0xb7, 0x22, 0x11, 0xfc,
	0x0a, 0xee, 0x3d, 0x11, 0x44, 0xe1, 0xc1, 0xa4, 0x56, 0x98, 0x20, 0x90, 0x05, 0xf7, 0x76, 0xe2,
	0x27, 0xf7, 0x8e, 0x3e, 0x03, 0xe8, 0x2b, 0x58, 0x6e, 0x35, 0x02, 0xdd, 0xe1, 0xa1, 0xef, 0x46,
	0x8e, 0xcc, 0x91, 0xbe, 0xd1, 0xf7, 0x84, 0x27, 0xbd, 0x74, 0x26, 0xd7, 0x05, 0x67, 0xb6, 0x09,
	0x78, 0x28, 0xce, 0x72, 0xfe, 0x69, 0xd0, 0xca, 0x7f, 0x06, 0xdf, 0x41, 0x53, 0x3c, 0xc5, 0x59,
	0xce, 0xde, 0xcd, 0x60, 0x72, 0xfc, 0xa2, 0xc9, 0x82, 0x7b, 0xbe, 0xfb, 0x14, 0x73, 0x26, 0x09,
	0xf8, 0x02, 0x34, 0x2f, 0x8e, 0x4d, 0x5f, 0xc6, 0xd7, 0x58, 0xb6, 0x49, 0xd3, 0xec, 0x03, 0x9f,
	0x36, 0xb2, 0x34, 0xfb, 0xc0, 0xc7, 0x6b, 0xe8, 0x88, 0x60, 0xcb, 0x13, 0xe1, 0x6d, 0x63, 0xda,
	0x1c, 0xd6, 0xc7, 0x0d, 0x56, 0x00, 0xf8, 0x11, 0xf4, 0x8d, 0xf2, 0xef, 0x54, 0x1b, 0xd6, 0xc6,
	0x97, 0x37, 0xaf, 0x95, 0x6b, 0xd5, 0xb1, 0xb0, 0x12, 0x19, 0x3f, 0x41, 0x77, 0xa3, 0x0e, 0x89,
	0x5e, 0x48, 0x35, 0xad, 0xaa, 0xb3, 0x73, 0x56, 0xa6, 0xe7, 0xfa, 0x62, 0x5a, 0xb4, 0x75, 0x4e,
	0x5f, 0x9c, 0xb3, 0x32, 0x1d, 0xa7, 0xd0, 0xdb, 0x94, 0x86, 0x4b, 0xdb, 0xd2, 0xe0, 0xcd, 0x19,
	0x83, 0xfc, 0x05, 0x27, 0x82, 0x34, 0x7f, 0xa8, 0xb4, 0x81, 0x76, 0x2a, 0xf9, 0xd5, 0xb2, 0xb0,
	0x12, 0x19, 0xe7, 0x40, 0x1e, 0x4e, 0x8a, 0x42, 0x41, 0x1a, 0xbc, 0x2d, 0xcf, 0xad, 0x44, 0x61,
	0x15, 0x11, 0x7e, 0x83, 0xfe, 0xc3, 0x69, 0xbb, 0xe8, 0xa5, 0x74, 0xba, 0x3e, 0xef, 0x94, 0xc7,
	0xa9, 0xca, 0xd2, 0x44, 0x89, 0x52, 0x4a, 0xaa, 0x57, 0x12, 0xa9, 0x9d, 0x65, 0x25, 0x72, 0x21,
	0xce, 0xda, 0x4a, 0xbb, 0xcf, 0x88, 0xb3, 0x63, 0x56, 0x22, 0xbf, 0xff, 0x5b, 0x83, 0xf6, 0xa1,
	0xa4, 0xd8, 0x87, 0xee, 0xd2, 0x9e, 0x9b, 0xd6, 0x8a, 0x19, 0xdf, 0xef, 0x0c, 0xc7, 0x25, 0x35,
	0x44, 0xe8, 0x1d, 0x20, 0xe7, 0xd6, 0xb6, 0x1c, 0x83, 0xd4, 0x73, 0xcc, 0xbe, 0x73, 0x8f, 0xbc,
	0x06, 0x0e, 0xe0, 0xea, 0x88, 0xe5, 0xc4, 0x26, 0x12, 0xd0, 0x2d, 0xdb, 0x35, 0xbf, 0x9a, 0xb3,
	0xa9, 0x6b, 0xda, 0x16, 0xd1, 0xf0, 0x25, 0xf4, 0x17, 0xc6, 0x94, 0xb9, 0x9f, 0x8d, 0x69, 0xa1,
	0xbe, 0xc0, 0x57, 0x80, 0x2a, 0x9c, 0x1b, 0xb4, 0xd2, 0x9b, 0x1c, 0xc3, 0xfa, 0xb2, 0x72, 0xed,
	0x95, 0x63, 0xb0, 0x1f, 0x06, 0x23, 0x6d, 0x15, 0x9b, 0x2d, 0x4d, 0xc3, 0x72, 0x49, 0xe7, 0x7f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0xbf, 0x68, 0x1c, 0xa4, 0x04, 0x00, 0x00,
}
