// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/multiplier/service.proto

package multiplier

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

type Arguments struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Arguments) Reset()         { *m = Arguments{} }
func (m *Arguments) String() string { return proto.CompactTextString(m) }
func (*Arguments) ProtoMessage()    {}
func (*Arguments) Descriptor() ([]byte, []int) {
	return fileDescriptor_53868c93e730d590, []int{0}
}

func (m *Arguments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Arguments.Unmarshal(m, b)
}
func (m *Arguments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Arguments.Marshal(b, m, deterministic)
}
func (m *Arguments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Arguments.Merge(m, src)
}
func (m *Arguments) XXX_Size() int {
	return xxx_messageInfo_Arguments.Size(m)
}
func (m *Arguments) XXX_DiscardUnknown() {
	xxx_messageInfo_Arguments.DiscardUnknown(m)
}

var xxx_messageInfo_Arguments proto.InternalMessageInfo

func (m *Arguments) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Arguments) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type Product struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_53868c93e730d590, []int{1}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Arguments)(nil), "twirp.example.multiplier.Arguments")
	proto.RegisterType((*Product)(nil), "twirp.example.multiplier.Product")
}

func init() { proto.RegisterFile("rpc/multiplier/service.proto", fileDescriptor_53868c93e730d590) }

var fileDescriptor_53868c93e730d590 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x2a, 0x48, 0xd6,
	0xcf, 0x2d, 0xcd, 0x29, 0xc9, 0x2c, 0xc8, 0xc9, 0x4c, 0x2d, 0xd2, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x28, 0x29, 0xcf, 0x2c, 0x2a, 0xd0,
	0x4b, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x43, 0xa8, 0x53, 0x52, 0xe7, 0xe2, 0x74, 0x2c,
	0x4a, 0x2f, 0xcd, 0x4d, 0xcd, 0x2b, 0x29, 0x16, 0xe2, 0xe1, 0x62, 0x4c, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0d, 0x62, 0x4c, 0x04, 0xf1, 0x92, 0x24, 0x98, 0x20, 0xbc, 0x24, 0x25, 0x79, 0x2e,
	0xf6, 0x80, 0xa2, 0xfc, 0x94, 0xd2, 0xe4, 0x12, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2,
	0x54, 0xa8, 0x52, 0x08, 0xc7, 0x28, 0x81, 0x8b, 0xcb, 0x17, 0x6e, 0xae, 0x50, 0x10, 0x17, 0x07,
	0x94, 0x57, 0x29, 0xa4, 0xac, 0x87, 0xcb, 0x7a, 0x3d, 0xb8, 0xdd, 0x52, 0x8a, 0xb8, 0x15, 0x41,
	0xed, 0x75, 0xe2, 0x89, 0xe2, 0x42, 0x88, 0x26, 0xb1, 0x81, 0xbd, 0x66, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x68, 0xfb, 0x66, 0x5b, 0xfa, 0x00, 0x00, 0x00,
}
