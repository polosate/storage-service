// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/storage/storage.proto

package storage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Storage struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87989e2875259d4, []int{0}
}

func (m *Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage.Unmarshal(m, b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return xxx_messageInfo_Storage.Size(m)
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

func (m *Storage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Storage) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Storage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Storage) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Storage) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87989e2875259d4, []int{1}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

type Response struct {
	Storage              *Storage   `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
	Storages             []*Storage `protobuf:"bytes,2,rep,name=storages,proto3" json:"storages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87989e2875259d4, []int{2}
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

func (m *Response) GetStorage() *Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *Response) GetStorages() []*Storage {
	if m != nil {
		return m.Storages
	}
	return nil
}

func init() {
	proto.RegisterType((*Storage)(nil), "storage.Storage")
	proto.RegisterType((*Specification)(nil), "storage.Specification")
	proto.RegisterType((*Response)(nil), "storage.Response")
}

func init() { proto.RegisterFile("proto/storage/storage.proto", fileDescriptor_c87989e2875259d4) }

var fileDescriptor_c87989e2875259d4 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xcd, 0xda, 0x3f, 0xe9, 0x48, 0x8b, 0xce, 0x41, 0x62, 0xf5, 0xb0, 0xec, 0x69, 0x51,
	0xa9, 0x50, 0xaf, 0x5e, 0xbc, 0x08, 0x5e, 0x3c, 0x64, 0x3f, 0x80, 0xa4, 0xc9, 0x28, 0x03, 0x35,
	0x59, 0x76, 0x97, 0x8a, 0x57, 0x3f, 0xb9, 0x10, 0x37, 0x5b, 0x0b, 0x9e, 0xf2, 0xde, 0x8f, 0x37,
	0xcc, 0xbc, 0xc0, 0x65, 0xdd, 0x84, 0x2e, 0xdc, 0xb5, 0x5d, 0x68, 0xcc, 0x3b, 0xa5, 0x77, 0x15,
	0x29, 0x4e, 0x7b, 0x5b, 0x7c, 0x0b, 0x98, 0x56, 0xbf, 0x1a, 0x17, 0x90, 0xb1, 0x53, 0x22, 0x17,
	0xe5, 0x58, 0x67, 0xec, 0x70, 0x09, 0xd2, 0x9a, 0xda, 0x58, 0xee, 0xbe, 0x54, 0x16, 0xe9, 0xe0,
	0x11, 0x61, 0xe4, 0xcd, 0x07, 0xa9, 0x51, 0x2e, 0xca, 0x99, 0x8e, 0x1a, 0xaf, 0x60, 0x66, 0x76,
	0x86, 0xb7, 0x66, 0xb3, 0x25, 0x35, 0xce, 0x45, 0x29, 0xf5, 0x1e, 0xe0, 0x05, 0xc8, 0xf0, 0xe9,
	0xa9, 0x79, 0x65, 0xa7, 0x26, 0x71, 0x6a, 0x1a, 0xfd, 0xb3, 0x2b, 0x6e, 0x60, 0x5e, 0xd5, 0x64,
	0xf9, 0x8d, 0xad, 0xe9, 0x38, 0xf8, 0x83, 0xcd, 0xe2, 0x70, 0x73, 0xe1, 0x40, 0x6a, 0x6a, 0xeb,
	0xe0, 0x5b, 0xc2, 0x6b, 0x48, 0x45, 0x62, 0xec, 0x64, 0x7d, 0xba, 0x4a, 0x3d, 0xfb, 0x52, 0x3a,
	0x05, 0xf0, 0x16, 0x64, 0x2f, 0x5b, 0x95, 0xe5, 0xc7, 0xff, 0x86, 0x87, 0xc4, 0xfa, 0x05, 0x16,
	0x3d, 0xac, 0xa8, 0xd9, 0xb1, 0x25, 0x7c, 0x80, 0xf9, 0x13, 0x7b, 0xf7, 0x38, 0x14, 0x3a, 0xdf,
	0x8f, 0xff, 0x3d, 0x7e, 0x79, 0x36, 0xf0, 0x74, 0x67, 0x71, 0xb4, 0x99, 0xc4, 0x7f, 0xbf, 0xff,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x24, 0xc3, 0x6a, 0xe8, 0x96, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for StorageService service

type StorageServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
}

type storageServiceClient struct {
	c           client.Client
	serviceName string
}

func NewStorageServiceClient(serviceName string, c client.Client) StorageServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "storage"
	}
	return &storageServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *storageServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "StorageService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StorageService service

type StorageServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
}

func RegisterStorageServiceHandler(s server.Server, hdlr StorageServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&StorageService{hdlr}, opts...))
}

type StorageService struct {
	StorageServiceHandler
}

func (h *StorageService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.StorageServiceHandler.FindAvailable(ctx, in, out)
}
