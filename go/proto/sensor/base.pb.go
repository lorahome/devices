// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/sensor/base.proto

package sensors

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

type Temperature struct {
	ValueC               float32  `protobuf:"fixed32,1,opt,name=value_c,json=valueC,proto3" json:"value_c,omitempty"`
	ValueF               float32  `protobuf:"fixed32,2,opt,name=value_f,json=valueF,proto3" json:"value_f,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Temperature) Reset()         { *m = Temperature{} }
func (m *Temperature) String() string { return proto.CompactTextString(m) }
func (*Temperature) ProtoMessage()    {}
func (*Temperature) Descriptor() ([]byte, []int) {
	return fileDescriptor_a708dc921c96b02f, []int{0}
}

func (m *Temperature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Temperature.Unmarshal(m, b)
}
func (m *Temperature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Temperature.Marshal(b, m, deterministic)
}
func (m *Temperature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Temperature.Merge(m, src)
}
func (m *Temperature) XXX_Size() int {
	return xxx_messageInfo_Temperature.Size(m)
}
func (m *Temperature) XXX_DiscardUnknown() {
	xxx_messageInfo_Temperature.DiscardUnknown(m)
}

var xxx_messageInfo_Temperature proto.InternalMessageInfo

func (m *Temperature) GetValueC() float32 {
	if m != nil {
		return m.ValueC
	}
	return 0
}

func (m *Temperature) GetValueF() float32 {
	if m != nil {
		return m.ValueF
	}
	return 0
}

type Humidity struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Humidity) Reset()         { *m = Humidity{} }
func (m *Humidity) String() string { return proto.CompactTextString(m) }
func (*Humidity) ProtoMessage()    {}
func (*Humidity) Descriptor() ([]byte, []int) {
	return fileDescriptor_a708dc921c96b02f, []int{1}
}

func (m *Humidity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Humidity.Unmarshal(m, b)
}
func (m *Humidity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Humidity.Marshal(b, m, deterministic)
}
func (m *Humidity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Humidity.Merge(m, src)
}
func (m *Humidity) XXX_Size() int {
	return xxx_messageInfo_Humidity.Size(m)
}
func (m *Humidity) XXX_DiscardUnknown() {
	xxx_messageInfo_Humidity.DiscardUnknown(m)
}

var xxx_messageInfo_Humidity proto.InternalMessageInfo

func (m *Humidity) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type AmbientLight struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AmbientLight) Reset()         { *m = AmbientLight{} }
func (m *AmbientLight) String() string { return proto.CompactTextString(m) }
func (*AmbientLight) ProtoMessage()    {}
func (*AmbientLight) Descriptor() ([]byte, []int) {
	return fileDescriptor_a708dc921c96b02f, []int{2}
}

func (m *AmbientLight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AmbientLight.Unmarshal(m, b)
}
func (m *AmbientLight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AmbientLight.Marshal(b, m, deterministic)
}
func (m *AmbientLight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AmbientLight.Merge(m, src)
}
func (m *AmbientLight) XXX_Size() int {
	return xxx_messageInfo_AmbientLight.Size(m)
}
func (m *AmbientLight) XXX_DiscardUnknown() {
	xxx_messageInfo_AmbientLight.DiscardUnknown(m)
}

var xxx_messageInfo_AmbientLight proto.InternalMessageInfo

func (m *AmbientLight) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Battery struct {
	Percent              uint32   `protobuf:"varint,1,opt,name=percent,proto3" json:"percent,omitempty"`
	VoltageMv            uint32   `protobuf:"varint,2,opt,name=voltage_mv,json=voltageMv,proto3" json:"voltage_mv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Battery) Reset()         { *m = Battery{} }
func (m *Battery) String() string { return proto.CompactTextString(m) }
func (*Battery) ProtoMessage()    {}
func (*Battery) Descriptor() ([]byte, []int) {
	return fileDescriptor_a708dc921c96b02f, []int{3}
}

func (m *Battery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Battery.Unmarshal(m, b)
}
func (m *Battery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Battery.Marshal(b, m, deterministic)
}
func (m *Battery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Battery.Merge(m, src)
}
func (m *Battery) XXX_Size() int {
	return xxx_messageInfo_Battery.Size(m)
}
func (m *Battery) XXX_DiscardUnknown() {
	xxx_messageInfo_Battery.DiscardUnknown(m)
}

var xxx_messageInfo_Battery proto.InternalMessageInfo

func (m *Battery) GetPercent() uint32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *Battery) GetVoltageMv() uint32 {
	if m != nil {
		return m.VoltageMv
	}
	return 0
}

func init() {
	proto.RegisterType((*Temperature)(nil), "sensors.Temperature")
	proto.RegisterType((*Humidity)(nil), "sensors.Humidity")
	proto.RegisterType((*AmbientLight)(nil), "sensors.AmbientLight")
	proto.RegisterType((*Battery)(nil), "sensors.Battery")
}

func init() { proto.RegisterFile("proto/sensor/base.proto", fileDescriptor_a708dc921c96b02f) }

var fileDescriptor_a708dc921c96b02f = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0xcd, 0x2b, 0xce, 0x2f, 0xd2, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x03, 0x8b,
	0x08, 0xb1, 0x43, 0x84, 0x8a, 0x95, 0xec, 0xb9, 0xb8, 0x43, 0x52, 0x73, 0x0b, 0x52, 0x8b, 0x12,
	0x4b, 0x4a, 0x8b, 0x52, 0x85, 0xc4, 0xb9, 0xd8, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0xe3, 0x93, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x98, 0x82, 0xd8, 0xc0, 0x5c, 0x67, 0x84, 0x44, 0x9a, 0x04, 0x13, 0x92,
	0x84, 0x9b, 0x92, 0x02, 0x17, 0x87, 0x47, 0x69, 0x6e, 0x66, 0x4a, 0x66, 0x49, 0xa5, 0x90, 0x08,
	0x17, 0x2b, 0x58, 0x14, 0xaa, 0x17, 0xc2, 0x51, 0x52, 0xe1, 0xe2, 0x71, 0xcc, 0x4d, 0xca, 0x4c,
	0xcd, 0x2b, 0xf1, 0xc9, 0x4c, 0xcf, 0x28, 0x41, 0x55, 0xc5, 0x0b, 0x53, 0xe5, 0xc4, 0xc5, 0xee,
	0x94, 0x58, 0x52, 0x92, 0x5a, 0x54, 0x29, 0x24, 0xc1, 0xc5, 0x5e, 0x90, 0x5a, 0x94, 0x9c, 0x9a,
	0x57, 0x02, 0x55, 0x02, 0xe3, 0x0a, 0xc9, 0x72, 0x71, 0x95, 0xe5, 0xe7, 0x94, 0x24, 0xa6, 0xa7,
	0xc6, 0xe7, 0x96, 0x81, 0x1d, 0xc2, 0x1b, 0xc4, 0x09, 0x15, 0xf1, 0x2d, 0x4b, 0x62, 0x03, 0x7b,
	0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x86, 0x7f, 0x63, 0x30, 0xf7, 0x00, 0x00, 0x00,
}