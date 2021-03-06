// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/hotel_group_view.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A hotel group view.
type HotelGroupView struct {
	// The resource name of the hotel group view.
	// Hotel Group view resource names have the form:
	//
	// `customers/{customer_id}/hotelGroupViews/{ad_group_id}_{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotelGroupView) Reset()         { *m = HotelGroupView{} }
func (m *HotelGroupView) String() string { return proto.CompactTextString(m) }
func (*HotelGroupView) ProtoMessage()    {}
func (*HotelGroupView) Descriptor() ([]byte, []int) {
	return fileDescriptor_hotel_group_view_e87bc812956d3c2c, []int{0}
}
func (m *HotelGroupView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotelGroupView.Unmarshal(m, b)
}
func (m *HotelGroupView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotelGroupView.Marshal(b, m, deterministic)
}
func (dst *HotelGroupView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotelGroupView.Merge(dst, src)
}
func (m *HotelGroupView) XXX_Size() int {
	return xxx_messageInfo_HotelGroupView.Size(m)
}
func (m *HotelGroupView) XXX_DiscardUnknown() {
	xxx_messageInfo_HotelGroupView.DiscardUnknown(m)
}

var xxx_messageInfo_HotelGroupView proto.InternalMessageInfo

func (m *HotelGroupView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*HotelGroupView)(nil), "google.ads.googleads.v0.resources.HotelGroupView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/hotel_group_view.proto", fileDescriptor_hotel_group_view_e87bc812956d3c2c)
}

var fileDescriptor_hotel_group_view_e87bc812956d3c2c = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xa2, 0xd4,
	0xe2, 0xfc, 0xd2, 0xa2, 0xe4, 0xd4, 0x62, 0xfd, 0x8c, 0xfc, 0x92, 0xd4, 0x9c, 0xf8, 0xf4, 0xa2,
	0xfc, 0xd2, 0x82, 0xf8, 0xb2, 0xcc, 0xd4, 0x72, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x45,
	0x88, 0x72, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x4e, 0xbd, 0x32, 0x03, 0x3d, 0xb8, 0x4e, 0x25,
	0x53, 0x2e, 0x3e, 0x0f, 0x90, 0x66, 0x77, 0x90, 0xde, 0xb0, 0xcc, 0xd4, 0x72, 0x21, 0x65, 0x2e,
	0x5e, 0x98, 0x74, 0x7c, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x0f,
	0x4c, 0xd0, 0x2f, 0x31, 0x37, 0xd5, 0xe9, 0x06, 0x23, 0x97, 0x6a, 0x72, 0x7e, 0xae, 0x1e, 0x41,
	0x0b, 0x9c, 0x84, 0x51, 0x8d, 0x0f, 0x00, 0x39, 0x2c, 0x80, 0x31, 0xca, 0x0b, 0xaa, 0x33, 0x3d,
	0x3f, 0x27, 0x31, 0x2f, 0x5d, 0x2f, 0xbf, 0x28, 0x5d, 0x3f, 0x3d, 0x35, 0x0f, 0xec, 0x6c, 0x98,
	0x27, 0x0b, 0x32, 0x8b, 0xf1, 0xf8, 0xd9, 0x1a, 0xce, 0x5a, 0xc4, 0xc4, 0xec, 0xee, 0xe8, 0xb8,
	0x8a, 0x49, 0xd1, 0x1d, 0x62, 0xa4, 0x63, 0x4a, 0xb1, 0x1e, 0x84, 0x09, 0x62, 0x85, 0x19, 0xe8,
	0x05, 0xc1, 0x54, 0x9e, 0x82, 0xa9, 0x89, 0x71, 0x4c, 0x29, 0x8e, 0x81, 0xab, 0x89, 0x09, 0x33,
	0x88, 0x81, 0xab, 0x49, 0x62, 0x03, 0x3b, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x57,
	0x63, 0xbb, 0x77, 0x01, 0x00, 0x00,
}
