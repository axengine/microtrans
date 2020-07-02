// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package order

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

type CreateOrderRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId"`
	Uid                  int64    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid"`
	Goods                int32    `protobuf:"varint,3,opt,name=goods,proto3" json:"goods"`
	Price                float64  `protobuf:"fixed64,4,opt,name=price,proto3" json:"price"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderRequest) Reset()         { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()    {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *CreateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRequest.Unmarshal(m, b)
}
func (m *CreateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRequest.Merge(m, src)
}
func (m *CreateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRequest.Size(m)
}
func (m *CreateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRequest proto.InternalMessageInfo

func (m *CreateOrderRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *CreateOrderRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *CreateOrderRequest) GetGoods() int32 {
	if m != nil {
		return m.Goods
	}
	return 0
}

func (m *CreateOrderRequest) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type SetOrderStatusRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetOrderStatusRequest) Reset()         { *m = SetOrderStatusRequest{} }
func (m *SetOrderStatusRequest) String() string { return proto.CompactTextString(m) }
func (*SetOrderStatusRequest) ProtoMessage()    {}
func (*SetOrderStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *SetOrderStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetOrderStatusRequest.Unmarshal(m, b)
}
func (m *SetOrderStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetOrderStatusRequest.Marshal(b, m, deterministic)
}
func (m *SetOrderStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetOrderStatusRequest.Merge(m, src)
}
func (m *SetOrderStatusRequest) XXX_Size() int {
	return xxx_messageInfo_SetOrderStatusRequest.Size(m)
}
func (m *SetOrderStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetOrderStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetOrderStatusRequest proto.InternalMessageInfo

func (m *SetOrderStatusRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *SetOrderStatusRequest) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type Response struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateOrderRequest)(nil), "order.CreateOrderRequest")
	proto.RegisterType((*SetOrderStatusRequest)(nil), "order.SetOrderStatusRequest")
	proto.RegisterType((*Response)(nil), "order.Response")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x8d, 0xdb, 0xd4, 0x76, 0x7a, 0x50, 0x06, 0x95, 0x28, 0x1e, 0x42, 0x4e, 0x39, 0x15,
	0xd1, 0x63, 0x8f, 0x9e, 0x7a, 0x12, 0xa6, 0xbf, 0xa0, 0x36, 0x43, 0x29, 0xa8, 0x59, 0x93, 0xec,
	0x0f, 0xf0, 0x9f, 0x4b, 0x66, 0x77, 0x41, 0x70, 0xc1, 0xdb, 0x7c, 0xfb, 0xf6, 0xf1, 0xf2, 0x1e,
	0xac, 0x62, 0x0a, 0x9c, 0xd6, 0x6d, 0x8a, 0x25, 0xa2, 0x16, 0x70, 0xef, 0x80, 0x2f, 0x89, 0xf7,
	0x85, 0x5f, 0x2b, 0x12, 0x7f, 0x75, 0x9c, 0x0b, 0x1a, 0xb8, 0x10, 0x79, 0x1b, 0x8c, 0xb2, 0xca,
	0x2f, 0x69, 0x44, 0xbc, 0x82, 0xa6, 0x3b, 0x05, 0x73, 0x6e, 0x95, 0x6f, 0xa8, 0x9e, 0x78, 0x0d,
	0xfa, 0x18, 0x63, 0xc8, 0xa6, 0xb1, 0xca, 0x6b, 0xea, 0xa1, 0x7e, 0x6d, 0xd3, 0xe9, 0xc0, 0x66,
	0x66, 0x95, 0x57, 0xd4, 0x83, 0xdb, 0xc2, 0xcd, 0x8e, 0x8b, 0x44, 0xed, 0xca, 0xbe, 0x74, 0xf9,
	0xff, 0xc0, 0x5b, 0x98, 0x67, 0xf9, 0x55, 0x32, 0x35, 0x0d, 0xe4, 0x1e, 0x61, 0x41, 0x9c, 0xdb,
	0xf8, 0x99, 0x19, 0x11, 0x66, 0x87, 0x18, 0x58, 0xac, 0x9a, 0xe4, 0xae, 0x0f, 0xfd, 0xc8, 0x47,
	0x31, 0x2d, 0xa9, 0x9e, 0x4f, 0xdf, 0x0a, 0xb4, 0x44, 0xe3, 0x06, 0x56, 0xbf, 0x4a, 0xe3, 0xdd,
	0xba, 0x1f, 0xe6, 0xef, 0x10, 0xf7, 0x97, 0x83, 0x34, 0x46, 0xb9, 0x33, 0xdc, 0xc0, 0x62, 0xec,
	0x80, 0x0f, 0x83, 0x3c, 0x59, 0x6a, 0xc2, 0xfc, 0x36, 0x97, 0xf1, 0x9f, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xe5, 0x8b, 0x58, 0x79, 0x8b, 0x01, 0x00, 0x00,
}
