// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package RegisterFound

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

type RegisterFoundREQMessage struct {
	//路由：取值 gateway_register 或 worker_register
	Cmd string `protobuf:"bytes,1,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
	//BusinessWorker注册
	BWRData *BusinessWorkerRegisterData `protobuf:"bytes,2,opt,name=BWRData,proto3" json:"BWRData,omitempty"`
	//Gateway注册
	GRData *GatewayRegisterData `protobuf:"bytes,3,opt,name=GRData,proto3" json:"GRData,omitempty"`
	//Gateway BusinessWorker 公共心跳
	HBDData *HeartBeat `protobuf:"bytes,4,opt,name=HBDData,proto3" json:"HBDData,omitempty"`
	//后台接口，展示
	BARData              *AdminGatewayList `protobuf:"bytes,5,opt,name=BARData,proto3" json:"BARData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RegisterFoundREQMessage) Reset()         { *m = RegisterFoundREQMessage{} }
func (m *RegisterFoundREQMessage) String() string { return proto.CompactTextString(m) }
func (*RegisterFoundREQMessage) ProtoMessage()    {}
func (*RegisterFoundREQMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *RegisterFoundREQMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterFoundREQMessage.Unmarshal(m, b)
}
func (m *RegisterFoundREQMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterFoundREQMessage.Marshal(b, m, deterministic)
}
func (m *RegisterFoundREQMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterFoundREQMessage.Merge(m, src)
}
func (m *RegisterFoundREQMessage) XXX_Size() int {
	return xxx_messageInfo_RegisterFoundREQMessage.Size(m)
}
func (m *RegisterFoundREQMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterFoundREQMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterFoundREQMessage proto.InternalMessageInfo

func (m *RegisterFoundREQMessage) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *RegisterFoundREQMessage) GetBWRData() *BusinessWorkerRegisterData {
	if m != nil {
		return m.BWRData
	}
	return nil
}

func (m *RegisterFoundREQMessage) GetGRData() *GatewayRegisterData {
	if m != nil {
		return m.GRData
	}
	return nil
}

func (m *RegisterFoundREQMessage) GetHBDData() *HeartBeat {
	if m != nil {
		return m.HBDData
	}
	return nil
}

func (m *RegisterFoundREQMessage) GetBARData() *AdminGatewayList {
	if m != nil {
		return m.BARData
	}
	return nil
}

type RegisterFoundRESMessage struct {
	//响应消息标识
	// connect  链接成功，建立连接时的响应
	// register 注册成功，work注册成功时会立即推送一次网关地址
	// broadcast_addresses 推送给work
	Cmd string `protobuf:"bytes,1,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
	// 1是成功   0是失败
	Code int32  `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,3,opt,name=Msg,proto3" json:"Msg,omitempty"`
	// 只有cmd=broadcast_addresses 时 BARData 才有数据
	BARData              *BroadcastAddressesResponseData `protobuf:"bytes,4,opt,name=BARData,proto3" json:"BARData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *RegisterFoundRESMessage) Reset()         { *m = RegisterFoundRESMessage{} }
func (m *RegisterFoundRESMessage) String() string { return proto.CompactTextString(m) }
func (*RegisterFoundRESMessage) ProtoMessage()    {}
func (*RegisterFoundRESMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *RegisterFoundRESMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterFoundRESMessage.Unmarshal(m, b)
}
func (m *RegisterFoundRESMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterFoundRESMessage.Marshal(b, m, deterministic)
}
func (m *RegisterFoundRESMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterFoundRESMessage.Merge(m, src)
}
func (m *RegisterFoundRESMessage) XXX_Size() int {
	return xxx_messageInfo_RegisterFoundRESMessage.Size(m)
}
func (m *RegisterFoundRESMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterFoundRESMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterFoundRESMessage proto.InternalMessageInfo

func (m *RegisterFoundRESMessage) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *RegisterFoundRESMessage) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegisterFoundRESMessage) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RegisterFoundRESMessage) GetBARData() *BroadcastAddressesResponseData {
	if m != nil {
		return m.BARData
	}
	return nil
}

//Gateway BusinessWorker 公共心跳
type HeartBeat struct {
	//和注册中心约定的key，当前全局都是空字符串
	SecretKey string `protobuf:"bytes,1,opt,name=SecretKey,proto3" json:"SecretKey,omitempty"`
	// 客户端发 ping ，服务端回复 pong
	Data                 string   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartBeat) Reset()         { *m = HeartBeat{} }
func (m *HeartBeat) String() string { return proto.CompactTextString(m) }
func (*HeartBeat) ProtoMessage()    {}
func (*HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *HeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeat.Unmarshal(m, b)
}
func (m *HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeat.Marshal(b, m, deterministic)
}
func (m *HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeat.Merge(m, src)
}
func (m *HeartBeat) XXX_Size() int {
	return xxx_messageInfo_HeartBeat.Size(m)
}
func (m *HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeat proto.InternalMessageInfo

func (m *HeartBeat) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *HeartBeat) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

//Gateway注册数据
type GatewayRegisterData struct {
	SecretKey string `protobuf:"bytes,1,opt,name=SecretKey,proto3" json:"SecretKey,omitempty"`
	// 业务类型，同组服务保持一致
	Business string `protobuf:"bytes,2,opt,name=Business,proto3" json:"Business,omitempty"`
	// 网关连接地址
	Address string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	// 节点名称
	Name                 string   `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayRegisterData) Reset()         { *m = GatewayRegisterData{} }
func (m *GatewayRegisterData) String() string { return proto.CompactTextString(m) }
func (*GatewayRegisterData) ProtoMessage()    {}
func (*GatewayRegisterData) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *GatewayRegisterData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayRegisterData.Unmarshal(m, b)
}
func (m *GatewayRegisterData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayRegisterData.Marshal(b, m, deterministic)
}
func (m *GatewayRegisterData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayRegisterData.Merge(m, src)
}
func (m *GatewayRegisterData) XXX_Size() int {
	return xxx_messageInfo_GatewayRegisterData.Size(m)
}
func (m *GatewayRegisterData) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayRegisterData.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayRegisterData proto.InternalMessageInfo

func (m *GatewayRegisterData) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *GatewayRegisterData) GetBusiness() string {
	if m != nil {
		return m.Business
	}
	return ""
}

func (m *GatewayRegisterData) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GatewayRegisterData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

//BusinessWorker注册数据
type BusinessWorkerRegisterData struct {
	SecretKey string `protobuf:"bytes,1,opt,name=SecretKey,proto3" json:"SecretKey,omitempty"`
	// 通过业务类型获得相应的所有网关地址，
	Business             string   `protobuf:"bytes,2,opt,name=Business,proto3" json:"Business,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BusinessWorkerRegisterData) Reset()         { *m = BusinessWorkerRegisterData{} }
func (m *BusinessWorkerRegisterData) String() string { return proto.CompactTextString(m) }
func (*BusinessWorkerRegisterData) ProtoMessage()    {}
func (*BusinessWorkerRegisterData) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *BusinessWorkerRegisterData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BusinessWorkerRegisterData.Unmarshal(m, b)
}
func (m *BusinessWorkerRegisterData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BusinessWorkerRegisterData.Marshal(b, m, deterministic)
}
func (m *BusinessWorkerRegisterData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BusinessWorkerRegisterData.Merge(m, src)
}
func (m *BusinessWorkerRegisterData) XXX_Size() int {
	return xxx_messageInfo_BusinessWorkerRegisterData.Size(m)
}
func (m *BusinessWorkerRegisterData) XXX_DiscardUnknown() {
	xxx_messageInfo_BusinessWorkerRegisterData.DiscardUnknown(m)
}

var xxx_messageInfo_BusinessWorkerRegisterData proto.InternalMessageInfo

func (m *BusinessWorkerRegisterData) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *BusinessWorkerRegisterData) GetBusiness() string {
	if m != nil {
		return m.Business
	}
	return ""
}

func (m *BusinessWorkerRegisterData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type BroadcastAddressesResponseData struct {
	// 网关注册的地址原文
	AddressList          []string `protobuf:"bytes,1,rep,name=AddressList,proto3" json:"AddressList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadcastAddressesResponseData) Reset()         { *m = BroadcastAddressesResponseData{} }
func (m *BroadcastAddressesResponseData) String() string { return proto.CompactTextString(m) }
func (*BroadcastAddressesResponseData) ProtoMessage()    {}
func (*BroadcastAddressesResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *BroadcastAddressesResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadcastAddressesResponseData.Unmarshal(m, b)
}
func (m *BroadcastAddressesResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadcastAddressesResponseData.Marshal(b, m, deterministic)
}
func (m *BroadcastAddressesResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastAddressesResponseData.Merge(m, src)
}
func (m *BroadcastAddressesResponseData) XXX_Size() int {
	return xxx_messageInfo_BroadcastAddressesResponseData.Size(m)
}
func (m *BroadcastAddressesResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastAddressesResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastAddressesResponseData proto.InternalMessageInfo

func (m *BroadcastAddressesResponseData) GetAddressList() []string {
	if m != nil {
		return m.AddressList
	}
	return nil
}

//后台获取gateway列表
type AdminGatewayList struct {
	SecretKey            string   `protobuf:"bytes,1,opt,name=SecretKey,proto3" json:"SecretKey,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdminGatewayList) Reset()         { *m = AdminGatewayList{} }
func (m *AdminGatewayList) String() string { return proto.CompactTextString(m) }
func (*AdminGatewayList) ProtoMessage()    {}
func (*AdminGatewayList) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{6}
}

func (m *AdminGatewayList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminGatewayList.Unmarshal(m, b)
}
func (m *AdminGatewayList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminGatewayList.Marshal(b, m, deterministic)
}
func (m *AdminGatewayList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminGatewayList.Merge(m, src)
}
func (m *AdminGatewayList) XXX_Size() int {
	return xxx_messageInfo_AdminGatewayList.Size(m)
}
func (m *AdminGatewayList) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminGatewayList.DiscardUnknown(m)
}

var xxx_messageInfo_AdminGatewayList proto.InternalMessageInfo

func (m *AdminGatewayList) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *AdminGatewayList) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterFoundREQMessage)(nil), "RegisterFound.RegisterFoundREQMessage")
	proto.RegisterType((*RegisterFoundRESMessage)(nil), "RegisterFound.RegisterFoundRESMessage")
	proto.RegisterType((*HeartBeat)(nil), "RegisterFound.HeartBeat")
	proto.RegisterType((*GatewayRegisterData)(nil), "RegisterFound.GatewayRegisterData")
	proto.RegisterType((*BusinessWorkerRegisterData)(nil), "RegisterFound.BusinessWorkerRegisterData")
	proto.RegisterType((*BroadcastAddressesResponseData)(nil), "RegisterFound.BroadcastAddressesResponseData")
	proto.RegisterType((*AdminGatewayList)(nil), "RegisterFound.AdminGatewayList")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x4f, 0xea, 0x40,
	0x10, 0xc7, 0x53, 0x0a, 0xf4, 0x75, 0x08, 0x09, 0xd9, 0x77, 0x78, 0x0d, 0x79, 0xd1, 0xa6, 0x27,
	0x3c, 0xc8, 0x01, 0x4f, 0x9a, 0x78, 0xa0, 0xa0, 0x90, 0x28, 0x26, 0x2e, 0x07, 0xce, 0x2b, 0x9d,
	0x90, 0x6a, 0xda, 0x92, 0xdd, 0x25, 0x86, 0x83, 0xff, 0x88, 0xfe, 0xb3, 0xa6, 0xdb, 0x5d, 0x7e,
	0x54, 0xd1, 0x18, 0x6f, 0xd3, 0xdd, 0xf9, 0xcc, 0x7c, 0xe7, 0xbb, 0x53, 0x68, 0x26, 0x28, 0x04,
	0x5b, 0x60, 0x77, 0xc9, 0x33, 0x99, 0x91, 0x26, 0xc5, 0x45, 0x2c, 0x24, 0xf2, 0xeb, 0x6c, 0x95,
	0x46, 0xc1, 0x5b, 0x05, 0xfe, 0xed, 0x9d, 0xd0, 0xab, 0xfb, 0x49, 0x01, 0x90, 0x16, 0xd8, 0x83,
	0x24, 0xf2, 0x2c, 0xdf, 0xea, 0xb8, 0x34, 0x0f, 0xc9, 0x00, 0x9c, 0x70, 0x46, 0x87, 0x4c, 0x32,
	0xaf, 0xe2, 0x5b, 0x9d, 0x46, 0xef, 0xa4, 0xbb, 0x07, 0x77, 0xc3, 0x95, 0x88, 0x53, 0x14, 0x62,
	0x96, 0xf1, 0x27, 0xe4, 0xe6, 0x2e, 0x07, 0xa8, 0x21, 0xc9, 0x05, 0xd4, 0x47, 0x45, 0x0d, 0x5b,
	0xd5, 0x08, 0x4a, 0x35, 0x46, 0x4c, 0xe2, 0x33, 0x5b, 0xef, 0xc1, 0x9a, 0x20, 0x3d, 0x70, 0xc6,
	0xe1, 0x50, 0xc1, 0x55, 0x05, 0x7b, 0x25, 0x78, 0x8c, 0x8c, 0xcb, 0x10, 0x99, 0xa4, 0x26, 0x91,
	0x9c, 0x83, 0x13, 0xf6, 0x8b, 0x86, 0x35, 0xc5, 0x1c, 0x97, 0x98, 0x7e, 0x94, 0xc4, 0xa9, 0xee,
	0x7a, 0x1b, 0x0b, 0x49, 0x4d, 0x7e, 0xf0, 0x6a, 0x7d, 0x70, 0x67, 0x7a, 0xd8, 0x1d, 0x02, 0xd5,
	0x41, 0x16, 0xa1, 0xb2, 0xa6, 0x46, 0x55, 0x9c, 0x67, 0x4d, 0xc4, 0x42, 0x4d, 0xea, 0xd2, 0x3c,
	0x24, 0xa3, 0xad, 0x9c, 0x62, 0x84, 0xd3, 0xb2, 0x87, 0x3c, 0x63, 0xd1, 0x9c, 0x09, 0xd9, 0x8f,
	0x22, 0x8e, 0x42, 0xa0, 0xa0, 0x28, 0x96, 0x59, 0x2a, 0x50, 0xfb, 0xa8, 0xc5, 0x5d, 0x82, 0xbb,
	0x99, 0x96, 0xfc, 0x07, 0x77, 0x8a, 0x73, 0x8e, 0xf2, 0x06, 0xd7, 0x5a, 0xd3, 0xf6, 0x20, 0x57,
	0xb6, 0x79, 0x34, 0x97, 0xaa, 0x38, 0x78, 0x81, 0xbf, 0x9f, 0x38, 0xfd, 0x4d, 0xa1, 0x36, 0xfc,
	0x31, 0x4f, 0xac, 0x8b, 0x6d, 0xbe, 0x89, 0x07, 0x8e, 0x56, 0xac, 0xc7, 0x35, 0x9f, 0x79, 0xfb,
	0x3b, 0x96, 0xa0, 0x9a, 0xd7, 0xa5, 0x2a, 0x0e, 0x1e, 0xa1, 0x7d, 0x78, 0x59, 0x7e, 0xa1, 0xc2,
	0xf4, 0xb2, 0x77, 0x7a, 0x85, 0x70, 0xf4, 0xb5, 0xa9, 0xc4, 0x87, 0x86, 0xbe, 0xc8, 0x17, 0xc0,
	0xb3, 0x7c, 0xbb, 0xe3, 0xd2, 0xdd, 0xa3, 0x60, 0x08, 0xad, 0xf2, 0x9e, 0xfc, 0xdc, 0xf4, 0x87,
	0xba, 0xfa, 0x09, 0xcf, 0xde, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x51, 0x5e, 0x38, 0x95, 0x03,
	0x00, 0x00,
}
