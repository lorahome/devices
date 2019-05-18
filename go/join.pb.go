// Code generated by protoc-gen-go. DO NOT EDIT.
// source: join.proto

package openiot

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

type JoinRequest struct {
	// Unique Device ID
	DeviceId uint64 `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Device details
	DeviceName         string `protobuf:"bytes,2,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	DeviceManufacturer string `protobuf:"bytes,3,opt,name=device_manufacturer,json=deviceManufacturer,proto3" json:"device_manufacturer,omitempty"`
	DeviceUrl          string `protobuf:"bytes,4,opt,name=device_url,json=deviceUrl,proto3" json:"device_url,omitempty"`
	// protobuf urls
	ProtoStatusMessageUrl string   `protobuf:"bytes,10,opt,name=proto_status_message_url,json=protoStatusMessageUrl,proto3" json:"proto_status_message_url,omitempty"`
	ProtoConfigMessageUrl string   `protobuf:"bytes,11,opt,name=proto_config_message_url,json=protoConfigMessageUrl,proto3" json:"proto_config_message_url,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0580e5afe87d2e40, []int{0}
}

func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetDeviceId() uint64 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *JoinRequest) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *JoinRequest) GetDeviceManufacturer() string {
	if m != nil {
		return m.DeviceManufacturer
	}
	return ""
}

func (m *JoinRequest) GetDeviceUrl() string {
	if m != nil {
		return m.DeviceUrl
	}
	return ""
}

func (m *JoinRequest) GetProtoStatusMessageUrl() string {
	if m != nil {
		return m.ProtoStatusMessageUrl
	}
	return ""
}

func (m *JoinRequest) GetProtoConfigMessageUrl() string {
	if m != nil {
		return m.ProtoConfigMessageUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*JoinRequest)(nil), "openiot.JoinRequest")
}

func init() { proto.RegisterFile("join.proto", fileDescriptor_0580e5afe87d2e40) }

var fileDescriptor_0580e5afe87d2e40 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xbf, 0x4a, 0xc7, 0x30,
	0x14, 0x46, 0x69, 0x2d, 0x6a, 0x6f, 0xb7, 0x88, 0x10, 0x10, 0xb1, 0x38, 0x75, 0xd2, 0xc1, 0xc1,
	0x07, 0x70, 0x52, 0xa8, 0x43, 0xc5, 0x39, 0xc4, 0xf6, 0xb6, 0x44, 0x9a, 0xa4, 0xe6, 0x8f, 0x6f,
	0xe1, 0x3b, 0x0b, 0x37, 0x81, 0x5f, 0xbb, 0x7e, 0xe7, 0x1c, 0xf8, 0x00, 0xbe, 0xad, 0x32, 0x0f,
	0x9b, 0xb3, 0xc1, 0xb2, 0x0b, 0xbb, 0xa1, 0x51, 0x36, 0xdc, 0xff, 0x95, 0xd0, 0xbc, 0x59, 0x65,
	0x06, 0xfc, 0x89, 0xe8, 0x03, 0xbb, 0x81, 0x7a, 0xc2, 0x5f, 0x35, 0xa2, 0x50, 0x13, 0x2f, 0xda,
	0xa2, 0xab, 0x86, 0xcb, 0x34, 0xbc, 0x4e, 0xec, 0x0e, 0x9a, 0x0c, 0x8d, 0xd4, 0xc8, 0xcb, 0xb6,
	0xe8, 0xea, 0x01, 0xd2, 0xf4, 0x2e, 0x35, 0xb2, 0x47, 0xb8, 0xca, 0x82, 0x96, 0x26, 0xce, 0x72,
	0x0c, 0xd1, 0xa1, 0xe3, 0x67, 0x24, 0xb2, 0x84, 0xfa, 0x1d, 0x61, 0xb7, 0x90, 0x73, 0x11, 0xdd,
	0xca, 0x2b, 0xf2, 0xf2, 0x81, 0x4f, 0xb7, 0xb2, 0x67, 0xe0, 0xf4, 0x57, 0xf8, 0x20, 0x43, 0xf4,
	0x42, 0xa3, 0xf7, 0x72, 0x49, 0x32, 0x90, 0x7c, 0x4d, 0xfc, 0x83, 0x70, 0x9f, 0xe8, 0x21, 0x1c,
	0xad, 0x99, 0xd5, 0x72, 0x08, 0x9b, 0x5d, 0xf8, 0x42, 0xf8, 0x14, 0x7e, 0x9d, 0xd3, 0xfc, 0xf4,
	0x1f, 0x00, 0x00, 0xff, 0xff, 0x37, 0xa6, 0x14, 0xd3, 0x2d, 0x01, 0x00, 0x00,
}
