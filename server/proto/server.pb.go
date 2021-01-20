// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package go_ms_server

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

type HandleRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Endpoint             string   `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Protocol             string   `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandleRequest) Reset()         { *m = HandleRequest{} }
func (m *HandleRequest) String() string { return proto.CompactTextString(m) }
func (*HandleRequest) ProtoMessage()    {}
func (*HandleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *HandleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleRequest.Unmarshal(m, b)
}
func (m *HandleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleRequest.Marshal(b, m, deterministic)
}
func (m *HandleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleRequest.Merge(m, src)
}
func (m *HandleRequest) XXX_Size() int {
	return xxx_messageInfo_HandleRequest.Size(m)
}
func (m *HandleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleRequest proto.InternalMessageInfo

func (m *HandleRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *HandleRequest) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *HandleRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type HandleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandleResponse) Reset()         { *m = HandleResponse{} }
func (m *HandleResponse) String() string { return proto.CompactTextString(m) }
func (*HandleResponse) ProtoMessage()    {}
func (*HandleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *HandleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleResponse.Unmarshal(m, b)
}
func (m *HandleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleResponse.Marshal(b, m, deterministic)
}
func (m *HandleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleResponse.Merge(m, src)
}
func (m *HandleResponse) XXX_Size() int {
	return xxx_messageInfo_HandleResponse.Size(m)
}
func (m *HandleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HandleResponse proto.InternalMessageInfo

type SubscribeRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type SubscribeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeResponse) Reset()         { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()    {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{3}
}

func (m *SubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeResponse.Unmarshal(m, b)
}
func (m *SubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeResponse.Merge(m, src)
}
func (m *SubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeResponse.Size(m)
}
func (m *SubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HandleRequest)(nil), "go.ms.server.HandleRequest")
	proto.RegisterType((*HandleResponse)(nil), "go.ms.server.HandleResponse")
	proto.RegisterType((*SubscribeRequest)(nil), "go.ms.server.SubscribeRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "go.ms.server.SubscribeResponse")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4f, 0xcf, 0xd7, 0xcb, 0xcd, 0x4c,
	0x2e, 0xca, 0xd7, 0x83, 0x08, 0x2b, 0x25, 0x72, 0xf1, 0x7a, 0x24, 0xe6, 0xa5, 0xe4, 0xa4, 0x06,
	0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x70, 0xb1, 0x83, 0xa4, 0x32, 0x93, 0x53, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x29, 0x2e, 0x8e, 0xd4, 0xbc, 0x94, 0x82,
	0xfc, 0xcc, 0xbc, 0x12, 0x09, 0x26, 0xb0, 0x14, 0x9c, 0x0f, 0x92, 0x03, 0x5b, 0x90, 0x9c, 0x9f,
	0x23, 0xc1, 0x0c, 0x91, 0x83, 0xf1, 0x95, 0x04, 0xb8, 0xf8, 0x60, 0x56, 0x14, 0x17, 0xe4, 0xe7,
	0x15, 0xa7, 0x2a, 0x69, 0x70, 0x09, 0x04, 0x97, 0x26, 0x15, 0x27, 0x17, 0x65, 0x26, 0xc1, 0xed,
	0x15, 0xe1, 0x62, 0x2d, 0xc9, 0x2f, 0xc8, 0x4c, 0x86, 0xda, 0x0a, 0xe1, 0x28, 0x09, 0x73, 0x09,
	0x22, 0xa9, 0x84, 0x68, 0x37, 0x5a, 0xcd, 0xc8, 0xc5, 0x16, 0x0c, 0x76, 0xbe, 0x90, 0x37, 0x17,
	0x1b, 0xc4, 0x6c, 0x21, 0x39, 0x3d, 0x34, 0xaf, 0xe9, 0xa1, 0xf8, 0x4b, 0x4a, 0x1e, 0xa7, 0x3c,
	0xd4, 0x51, 0x0c, 0x42, 0x21, 0x5c, 0x9c, 0x70, 0xcb, 0x84, 0x14, 0x31, 0xd4, 0xa3, 0x3b, 0x59,
	0x4a, 0x09, 0x9f, 0x12, 0x98, 0xa9, 0x49, 0x6c, 0xe0, 0x80, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xe0, 0x77, 0x9a, 0xe4, 0x89, 0x01, 0x00, 0x00,
}
