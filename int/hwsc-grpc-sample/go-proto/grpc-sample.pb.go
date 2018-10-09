// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc-sample.proto

package grpcSamplePb

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The SampleService request message containing the user's name.
type SampleServiceRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SampleServiceRequest) Reset()         { *m = SampleServiceRequest{} }
func (m *SampleServiceRequest) String() string { return proto.CompactTextString(m) }
func (*SampleServiceRequest) ProtoMessage()    {}
func (*SampleServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4c938bb013b3053, []int{0}
}

func (m *SampleServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleServiceRequest.Unmarshal(m, b)
}
func (m *SampleServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleServiceRequest.Marshal(b, m, deterministic)
}
func (m *SampleServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleServiceRequest.Merge(m, src)
}
func (m *SampleServiceRequest) XXX_Size() int {
	return xxx_messageInfo_SampleServiceRequest.Size(m)
}
func (m *SampleServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SampleServiceRequest proto.InternalMessageInfo

func (m *SampleServiceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The SampleService response message containing the greetings
type SampleServiceResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SampleServiceResponse) Reset()         { *m = SampleServiceResponse{} }
func (m *SampleServiceResponse) String() string { return proto.CompactTextString(m) }
func (*SampleServiceResponse) ProtoMessage()    {}
func (*SampleServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4c938bb013b3053, []int{1}
}

func (m *SampleServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleServiceResponse.Unmarshal(m, b)
}
func (m *SampleServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleServiceResponse.Marshal(b, m, deterministic)
}
func (m *SampleServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleServiceResponse.Merge(m, src)
}
func (m *SampleServiceResponse) XXX_Size() int {
	return xxx_messageInfo_SampleServiceResponse.Size(m)
}
func (m *SampleServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SampleServiceResponse proto.InternalMessageInfo

func (m *SampleServiceResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SampleServiceRequest)(nil), "sample.SampleServiceRequest")
	proto.RegisterType((*SampleServiceResponse)(nil), "sample.SampleServiceResponse")
}

func init() { proto.RegisterFile("grpc-sample.proto", fileDescriptor_c4c938bb013b3053) }

var fileDescriptor_c4c938bb013b3053 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2f, 0x2a, 0x48,
	0xd6, 0x2d, 0x4e, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83,
	0xf0, 0x94, 0xb4, 0xb8, 0x44, 0x82, 0xc1, 0xac, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0xd4, 0xa0,
	0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0xc9, 0x90, 0x4b, 0x14, 0x4d, 0x6d, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x3a, 0x4c, 0x3d, 0x8c,
	0x6b, 0x14, 0xc5, 0xc5, 0x8b, 0xa2, 0x45, 0xc8, 0x93, 0x8b, 0x23, 0x38, 0xb1, 0xd2, 0x23, 0x35,
	0x27, 0x27, 0x5f, 0x48, 0x46, 0x0f, 0xea, 0x24, 0x6c, 0x2e, 0x90, 0x92, 0xc5, 0x21, 0x0b, 0xb1,
	0x53, 0x89, 0xc1, 0x89, 0x2f, 0x8a, 0x07, 0xe4, 0x2f, 0x88, 0x74, 0x40, 0x52, 0x12, 0x1b, 0xd8,
	0x67, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xb6, 0x12, 0x8c, 0xee, 0x00, 0x00, 0x00,
}
