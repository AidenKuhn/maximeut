// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service-twirp-example.proto

package haberdasher

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

// A Hat is a piece of headwear made by a Haberdasher.
type Hat struct {
	// The size of a hat should always be in inches.
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// The color of a hat will never be 'invisible', but other than
	// that, anything is fair game.
	Color string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	// The name of a hat is it's type. Like, 'bowler', or something.
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hat) Reset()         { *m = Hat{} }
func (m *Hat) String() string { return proto.CompactTextString(m) }
func (*Hat) ProtoMessage()    {}
func (*Hat) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ff6cc22c260d56, []int{0}
}

func (m *Hat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hat.Unmarshal(m, b)
}
func (m *Hat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hat.Marshal(b, m, deterministic)
}
func (m *Hat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hat.Merge(m, src)
}
func (m *Hat) XXX_Size() int {
	return xxx_messageInfo_Hat.Size(m)
}
func (m *Hat) XXX_DiscardUnknown() {
	xxx_messageInfo_Hat.DiscardUnknown(m)
}

var xxx_messageInfo_Hat proto.InternalMessageInfo

func (m *Hat) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Hat) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Hat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Size is passed when requesting a new hat to be made. It's always measured in
// inches.
type Size struct {
	Inches               int32    `protobuf:"varint,1,opt,name=inches,proto3" json:"inches,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Size) Reset()         { *m = Size{} }
func (m *Size) String() string { return proto.CompactTextString(m) }
func (*Size) ProtoMessage()    {}
func (*Size) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ff6cc22c260d56, []int{1}
}

func (m *Size) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Size.Unmarshal(m, b)
}
func (m *Size) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Size.Marshal(b, m, deterministic)
}
func (m *Size) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Size.Merge(m, src)
}
func (m *Size) XXX_Size() int {
	return xxx_messageInfo_Size.Size(m)
}
func (m *Size) XXX_DiscardUnknown() {
	xxx_messageInfo_Size.DiscardUnknown(m)
}

var xxx_messageInfo_Size proto.InternalMessageInfo

func (m *Size) GetInches() int32 {
	if m != nil {
		return m.Inches
	}
	return 0
}

func init() {
	proto.RegisterType((*Hat)(nil), "spec.haberdasher.Hat")
	proto.RegisterType((*Size)(nil), "spec.haberdasher.Size")
}

func init() { proto.RegisterFile("service-twirp-example.proto", fileDescriptor_e3ff6cc22c260d56) }

var fileDescriptor_e3ff6cc22c260d56 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2d, 0x29, 0xcf, 0x2c, 0x2a, 0xd0, 0x4d, 0xad, 0x48, 0xcc, 0x2d, 0xc8,
	0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x28, 0x2e, 0x48, 0x4d, 0xd6, 0xcb, 0x48,
	0x4c, 0x4a, 0x2d, 0x4a, 0x49, 0x2c, 0xce, 0x48, 0x2d, 0x52, 0x72, 0xe6, 0x62, 0xf6, 0x48, 0x2c,
	0x11, 0x12, 0xe2, 0x62, 0x29, 0xce, 0xac, 0x4a, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02,
	0xb3, 0x85, 0x44, 0xb8, 0x58, 0x93, 0xf3, 0x73, 0xf2, 0x8b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38,
	0x83, 0x20, 0x1c, 0x90, 0xca, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x66, 0xb0, 0x20, 0x98, 0xad, 0x24,
	0xc7, 0xc5, 0x12, 0x0c, 0xd2, 0x21, 0xc6, 0xc5, 0x96, 0x99, 0x97, 0x9c, 0x91, 0x5a, 0x0c, 0x35,
	0x07, 0xca, 0x33, 0x72, 0xe7, 0xe2, 0xf6, 0x40, 0xd8, 0x29, 0x64, 0xc1, 0xc5, 0xee, 0x9b, 0x98,
	0x9d, 0x0a, 0xb2, 0x57, 0x4c, 0x0f, 0xdd, 0x45, 0x7a, 0x20, 0x93, 0xa4, 0x44, 0x31, 0xc5, 0x3d,
	0x12, 0x4b, 0x9c, 0x78, 0xa3, 0xb8, 0x91, 0x84, 0x92, 0xd8, 0xc0, 0xbe, 0x32, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xae, 0xdf, 0xdb, 0x33, 0xf4, 0x00, 0x00, 0x00,
}
