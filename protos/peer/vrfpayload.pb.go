// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vrfpayload.proto

package peer

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

type VrfPayload struct {
	VrfResult            []byte   `protobuf:"bytes,1,opt,name=vrf_result,json=vrfResult,proto3" json:"vrf_result,omitempty"`
	VrfProof             []byte   `protobuf:"bytes,2,opt,name=vrf_proof,json=vrfProof,proto3" json:"vrf_proof,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VrfPayload) Reset()         { *m = VrfPayload{} }
func (m *VrfPayload) String() string { return proto.CompactTextString(m) }
func (*VrfPayload) ProtoMessage()    {}
func (*VrfPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e9257332c5b4d0, []int{0}
}

func (m *VrfPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VrfPayload.Unmarshal(m, b)
}
func (m *VrfPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VrfPayload.Marshal(b, m, deterministic)
}
func (m *VrfPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VrfPayload.Merge(m, src)
}
func (m *VrfPayload) XXX_Size() int {
	return xxx_messageInfo_VrfPayload.Size(m)
}
func (m *VrfPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_VrfPayload.DiscardUnknown(m)
}

var xxx_messageInfo_VrfPayload proto.InternalMessageInfo

func (m *VrfPayload) GetVrfResult() []byte {
	if m != nil {
		return m.VrfResult
	}
	return nil
}

func (m *VrfPayload) GetVrfProof() []byte {
	if m != nil {
		return m.VrfProof
	}
	return nil
}

func (m *VrfPayload) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type ChaincodeResponsePayload struct {
	Payload              []byte            `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	VrfEndorsements      []*VrfEndorsement `protobuf:"bytes,2,rep,name=vrf_endorsements,json=vrfEndorsements,proto3" json:"vrf_endorsements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ChaincodeResponsePayload) Reset()         { *m = ChaincodeResponsePayload{} }
func (m *ChaincodeResponsePayload) String() string { return proto.CompactTextString(m) }
func (*ChaincodeResponsePayload) ProtoMessage()    {}
func (*ChaincodeResponsePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e9257332c5b4d0, []int{1}
}

func (m *ChaincodeResponsePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeResponsePayload.Unmarshal(m, b)
}
func (m *ChaincodeResponsePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeResponsePayload.Marshal(b, m, deterministic)
}
func (m *ChaincodeResponsePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeResponsePayload.Merge(m, src)
}
func (m *ChaincodeResponsePayload) XXX_Size() int {
	return xxx_messageInfo_ChaincodeResponsePayload.Size(m)
}
func (m *ChaincodeResponsePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeResponsePayload.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeResponsePayload proto.InternalMessageInfo

func (m *ChaincodeResponsePayload) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ChaincodeResponsePayload) GetVrfEndorsements() []*VrfEndorsement {
	if m != nil {
		return m.VrfEndorsements
	}
	return nil
}

type VrfEndorsement struct {
	Result               []byte   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Proof                []byte   `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VrfEndorsement) Reset()         { *m = VrfEndorsement{} }
func (m *VrfEndorsement) String() string { return proto.CompactTextString(m) }
func (*VrfEndorsement) ProtoMessage()    {}
func (*VrfEndorsement) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e9257332c5b4d0, []int{2}
}

func (m *VrfEndorsement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VrfEndorsement.Unmarshal(m, b)
}
func (m *VrfEndorsement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VrfEndorsement.Marshal(b, m, deterministic)
}
func (m *VrfEndorsement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VrfEndorsement.Merge(m, src)
}
func (m *VrfEndorsement) XXX_Size() int {
	return xxx_messageInfo_VrfEndorsement.Size(m)
}
func (m *VrfEndorsement) XXX_DiscardUnknown() {
	xxx_messageInfo_VrfEndorsement.DiscardUnknown(m)
}

var xxx_messageInfo_VrfEndorsement proto.InternalMessageInfo

func (m *VrfEndorsement) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VrfEndorsement) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterType((*VrfPayload)(nil), "protos.VrfPayload")
	proto.RegisterType((*ChaincodeResponsePayload)(nil), "protos.ChaincodeResponsePayload")
	proto.RegisterType((*VrfEndorsement)(nil), "protos.VrfEndorsement")
}

func init() { proto.RegisterFile("vrfpayload.proto", fileDescriptor_d1e9257332c5b4d0) }

var fileDescriptor_d1e9257332c5b4d0 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0x86, 0x53, 0x9f, 0xa2, 0x12, 0x64, 0x04, 0x44, 0x18, 0x3d, 0x4d, 0x84, 0x16,
	0xf4, 0x2e, 0xa8, 0x78, 0x1f, 0x3d, 0xec, 0xe0, 0xa5, 0xf4, 0xc7, 0xcb, 0x5a, 0xe8, 0xfa, 0xc2,
	0x4b, 0x17, 0xd9, 0x7f, 0x2f, 0x4b, 0x36, 0xd6, 0x9c, 0xc2, 0x27, 0x9f, 0x3c, 0xde, 0xf7, 0x1b,
	0x78, 0xb0, 0xac, 0x74, 0xb1, 0xef, 0xa8, 0xa8, 0x13, 0xcd, 0x34, 0x90, 0x98, 0xb9, 0xc3, 0xc4,
	0x25, 0xc0, 0x9a, 0xd5, 0xca, 0x3b, 0xf1, 0x0c, 0x60, 0x59, 0xe5, 0x8c, 0x66, 0xd7, 0x0d, 0x32,
	0x5a, 0x44, 0xcb, 0xdb, 0xec, 0xda, 0xb2, 0xca, 0xdc, 0x85, 0x78, 0x82, 0x03, 0xe4, 0x9a, 0x89,
	0x94, 0x9c, 0x38, 0x7b, 0x65, 0x59, 0xad, 0x0e, 0x2c, 0x24, 0x5c, 0x1e, 0x57, 0xc8, 0xa9, 0x53,
	0x27, 0x8c, 0xff, 0x40, 0x7e, 0x37, 0x45, 0xdb, 0x57, 0x54, 0x63, 0x86, 0x46, 0x53, 0x6f, 0xf0,
	0xb4, 0x71, 0x34, 0x15, 0x05, 0x53, 0xe2, 0xd3, 0xa5, 0xce, 0xb1, 0xaf, 0x89, 0x0d, 0x6e, 0xb1,
	0x1f, 0x8c, 0x9c, 0x2c, 0xa6, 0xcb, 0x9b, 0xb7, 0xb9, 0xef, 0x60, 0x92, 0x35, 0xab, 0x9f, 0xb3,
	0xce, 0xee, 0x6d, 0xc0, 0x26, 0xfe, 0x80, 0xbb, 0xf0, 0x89, 0x98, 0xc3, 0x2c, 0x28, 0x77, 0x24,
	0xf1, 0x08, 0x17, 0xe3, 0x56, 0x1e, 0xbe, 0x5e, 0x7f, 0x5f, 0x36, 0xed, 0xd0, 0xec, 0xca, 0xa4,
	0xa2, 0x6d, 0xda, 0xec, 0x35, 0x72, 0x87, 0xf5, 0x06, 0x39, 0x55, 0x45, 0xc9, 0x6d, 0x95, 0xfa,
	0x1c, 0xa9, 0x46, 0xe4, 0xd2, 0xff, 0xe8, 0xfb, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x1c,
	0xee, 0xde, 0x6c, 0x01, 0x00, 0x00,
}
