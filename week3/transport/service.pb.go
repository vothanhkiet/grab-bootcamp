// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transport/service.proto

package transport

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cf267de7b508487, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Paging struct {
	Total                int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Paging) Reset()         { *m = Paging{} }
func (m *Paging) String() string { return proto.CompactTextString(m) }
func (*Paging) ProtoMessage()    {}
func (*Paging) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cf267de7b508487, []int{1}
}

func (m *Paging) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Paging.Unmarshal(m, b)
}
func (m *Paging) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Paging.Marshal(b, m, deterministic)
}
func (m *Paging) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paging.Merge(m, src)
}
func (m *Paging) XXX_Size() int {
	return xxx_messageInfo_Paging.Size(m)
}
func (m *Paging) XXX_DiscardUnknown() {
	xxx_messageInfo_Paging.DiscardUnknown(m)
}

var xxx_messageInfo_Paging proto.InternalMessageInfo

func (m *Paging) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Paging) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Paging) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type PassengerFeedback struct {
	FeedbackID  string `protobuf:"bytes,1,opt,name=feedbackID,proto3" json:"feedbackID,omitempty"`
	BookingCode string `protobuf:"bytes,2,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	// Should be string?
	PassengerID          int32    `protobuf:"varint,3,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	Feedback             string   `protobuf:"bytes,4,opt,name=feedback,proto3" json:"feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassengerFeedback) Reset()         { *m = PassengerFeedback{} }
func (m *PassengerFeedback) String() string { return proto.CompactTextString(m) }
func (*PassengerFeedback) ProtoMessage()    {}
func (*PassengerFeedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cf267de7b508487, []int{2}
}

func (m *PassengerFeedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassengerFeedback.Unmarshal(m, b)
}
func (m *PassengerFeedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassengerFeedback.Marshal(b, m, deterministic)
}
func (m *PassengerFeedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassengerFeedback.Merge(m, src)
}
func (m *PassengerFeedback) XXX_Size() int {
	return xxx_messageInfo_PassengerFeedback.Size(m)
}
func (m *PassengerFeedback) XXX_DiscardUnknown() {
	xxx_messageInfo_PassengerFeedback.DiscardUnknown(m)
}

var xxx_messageInfo_PassengerFeedback proto.InternalMessageInfo

func (m *PassengerFeedback) GetFeedbackID() string {
	if m != nil {
		return m.FeedbackID
	}
	return ""
}

func (m *PassengerFeedback) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

func (m *PassengerFeedback) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

func (m *PassengerFeedback) GetFeedback() string {
	if m != nil {
		return m.Feedback
	}
	return ""
}

type AddPassengerFeedbackRequest struct {
	Feedback             *PassengerFeedback `protobuf:"bytes,1,opt,name=feedback,proto3" json:"feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AddPassengerFeedbackRequest) Reset()         { *m = AddPassengerFeedbackRequest{} }
func (m *AddPassengerFeedbackRequest) String() string { return proto.CompactTextString(m) }
func (*AddPassengerFeedbackRequest) ProtoMessage()    {}
func (*AddPassengerFeedbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cf267de7b508487, []int{3}
}

func (m *AddPassengerFeedbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPassengerFeedbackRequest.Unmarshal(m, b)
}
func (m *AddPassengerFeedbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPassengerFeedbackRequest.Marshal(b, m, deterministic)
}
func (m *AddPassengerFeedbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPassengerFeedbackRequest.Merge(m, src)
}
func (m *AddPassengerFeedbackRequest) XXX_Size() int {
	return xxx_messageInfo_AddPassengerFeedbackRequest.Size(m)
}
func (m *AddPassengerFeedbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPassengerFeedbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddPassengerFeedbackRequest proto.InternalMessageInfo

func (m *AddPassengerFeedbackRequest) GetFeedback() *PassengerFeedback {
	if m != nil {
		return m.Feedback
	}
	return nil
}

type AddPassengerFeedbackResponse struct {
	Errors               []*Error `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPassengerFeedbackResponse) Reset()         { *m = AddPassengerFeedbackResponse{} }
func (m *AddPassengerFeedbackResponse) String() string { return proto.CompactTextString(m) }
func (*AddPassengerFeedbackResponse) ProtoMessage()    {}
func (*AddPassengerFeedbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cf267de7b508487, []int{4}
}

func (m *AddPassengerFeedbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPassengerFeedbackResponse.Unmarshal(m, b)
}
func (m *AddPassengerFeedbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPassengerFeedbackResponse.Marshal(b, m, deterministic)
}
func (m *AddPassengerFeedbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPassengerFeedbackResponse.Merge(m, src)
}
func (m *AddPassengerFeedbackResponse) XXX_Size() int {
	return xxx_messageInfo_AddPassengerFeedbackResponse.Size(m)
}
func (m *AddPassengerFeedbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPassengerFeedbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddPassengerFeedbackResponse proto.InternalMessageInfo

func (m *AddPassengerFeedbackResponse) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type ListPassengerFeedbackRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	BookingCode          string   `protobuf:"bytes,3,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	PassengerID          int32    `protobuf:"varint,4,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPassengerFeedbackRequest) Reset()         { *m = ListPassengerFeedbackRequest{} }
func (m *ListPassengerFeedbackRequest) String() string { return proto.CompactTextString(m) }
func (*ListPassengerFeedbackRequest) ProtoMessage()    {}
func (*ListPassengerFeedbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cf267de7b508487, []int{5}
}

func (m *ListPassengerFeedbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPassengerFeedbackRequest.Unmarshal(m, b)
}
func (m *ListPassengerFeedbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPassengerFeedbackRequest.Marshal(b, m, deterministic)
}
func (m *ListPassengerFeedbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPassengerFeedbackRequest.Merge(m, src)
}
func (m *ListPassengerFeedbackRequest) XXX_Size() int {
	return xxx_messageInfo_ListPassengerFeedbackRequest.Size(m)
}
func (m *ListPassengerFeedbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPassengerFeedbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPassengerFeedbackRequest proto.InternalMessageInfo

func (m *ListPassengerFeedbackRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListPassengerFeedbackRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListPassengerFeedbackRequest) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

func (m *ListPassengerFeedbackRequest) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

type ListPassengerFeedbackResponse struct {
	Errors               []*Error             `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	Paging               *Paging              `protobuf:"bytes,2,opt,name=paging,proto3" json:"paging,omitempty"`
	Data                 []*PassengerFeedback `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListPassengerFeedbackResponse) Reset()         { *m = ListPassengerFeedbackResponse{} }
func (m *ListPassengerFeedbackResponse) String() string { return proto.CompactTextString(m) }
func (*ListPassengerFeedbackResponse) ProtoMessage()    {}
func (*ListPassengerFeedbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cf267de7b508487, []int{6}
}

func (m *ListPassengerFeedbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPassengerFeedbackResponse.Unmarshal(m, b)
}
func (m *ListPassengerFeedbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPassengerFeedbackResponse.Marshal(b, m, deterministic)
}
func (m *ListPassengerFeedbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPassengerFeedbackResponse.Merge(m, src)
}
func (m *ListPassengerFeedbackResponse) XXX_Size() int {
	return xxx_messageInfo_ListPassengerFeedbackResponse.Size(m)
}
func (m *ListPassengerFeedbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPassengerFeedbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPassengerFeedbackResponse proto.InternalMessageInfo

func (m *ListPassengerFeedbackResponse) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *ListPassengerFeedbackResponse) GetPaging() *Paging {
	if m != nil {
		return m.Paging
	}
	return nil
}

func (m *ListPassengerFeedbackResponse) GetData() []*PassengerFeedback {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteFeedBackRequest struct {
	PassengerID          int32    `protobuf:"varint,4,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFeedBackRequest) Reset()         { *m = DeleteFeedBackRequest{} }
func (m *DeleteFeedBackRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteFeedBackRequest) ProtoMessage()    {}
func (*DeleteFeedBackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cf267de7b508487, []int{7}
}

func (m *DeleteFeedBackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFeedBackRequest.Unmarshal(m, b)
}
func (m *DeleteFeedBackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFeedBackRequest.Marshal(b, m, deterministic)
}
func (m *DeleteFeedBackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFeedBackRequest.Merge(m, src)
}
func (m *DeleteFeedBackRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteFeedBackRequest.Size(m)
}
func (m *DeleteFeedBackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFeedBackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFeedBackRequest proto.InternalMessageInfo

func (m *DeleteFeedBackRequest) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

type DeleteFeedBackResponse struct {
	Errors               []*Error `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	FeedbackIds          []string `protobuf:"bytes,2,rep,name=feedbackIds,proto3" json:"feedbackIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFeedBackResponse) Reset()         { *m = DeleteFeedBackResponse{} }
func (m *DeleteFeedBackResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteFeedBackResponse) ProtoMessage()    {}
func (*DeleteFeedBackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cf267de7b508487, []int{8}
}

func (m *DeleteFeedBackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFeedBackResponse.Unmarshal(m, b)
}
func (m *DeleteFeedBackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFeedBackResponse.Marshal(b, m, deterministic)
}
func (m *DeleteFeedBackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFeedBackResponse.Merge(m, src)
}
func (m *DeleteFeedBackResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteFeedBackResponse.Size(m)
}
func (m *DeleteFeedBackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFeedBackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFeedBackResponse proto.InternalMessageInfo

func (m *DeleteFeedBackResponse) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *DeleteFeedBackResponse) GetFeedbackIds() []string {
	if m != nil {
		return m.FeedbackIds
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "transport.Error")
	proto.RegisterType((*Paging)(nil), "transport.Paging")
	proto.RegisterType((*PassengerFeedback)(nil), "transport.PassengerFeedback")
	proto.RegisterType((*AddPassengerFeedbackRequest)(nil), "transport.AddPassengerFeedbackRequest")
	proto.RegisterType((*AddPassengerFeedbackResponse)(nil), "transport.AddPassengerFeedbackResponse")
	proto.RegisterType((*ListPassengerFeedbackRequest)(nil), "transport.ListPassengerFeedbackRequest")
	proto.RegisterType((*ListPassengerFeedbackResponse)(nil), "transport.ListPassengerFeedbackResponse")
	proto.RegisterType((*DeleteFeedBackRequest)(nil), "transport.DeleteFeedBackRequest")
	proto.RegisterType((*DeleteFeedBackResponse)(nil), "transport.DeleteFeedBackResponse")
}

func init() { proto.RegisterFile("transport/service.proto", fileDescriptor_1cf267de7b508487) }

var fileDescriptor_1cf267de7b508487 = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xd6, 0xc4, 0x49, 0xee, 0xe5, 0x84, 0x7b, 0x81, 0xd1, 0x05, 0x2c, 0x93, 0x5b, 0x19, 0x2f,
	0x68, 0x8a, 0x20, 0x06, 0x77, 0xd3, 0x66, 0xd5, 0x50, 0x8a, 0x8a, 0xca, 0x02, 0xb9, 0xa8, 0x5d,
	0x8f, 0xed, 0xc1, 0x71, 0x49, 0x3c, 0xae, 0x67, 0xa0, 0x52, 0x77, 0xed, 0x13, 0x54, 0xb0, 0xe8,
	0xa2, 0xea, 0xb6, 0xfb, 0x3e, 0x4b, 0x5f, 0xa1, 0x0f, 0x52, 0xcd, 0xf8, 0x07, 0x93, 0x40, 0x22,
	0x56, 0x30, 0xe7, 0x7c, 0xe7, 0x9b, 0xef, 0x7c, 0xe7, 0x78, 0x02, 0xab, 0x22, 0x25, 0x31, 0x4f,
	0x58, 0x2a, 0x6c, 0x4e, 0xd3, 0x8b, 0xc8, 0xa7, 0xdd, 0x24, 0x65, 0x82, 0xe1, 0xb9, 0x32, 0x61,
	0xb4, 0x43, 0xc6, 0xc2, 0x21, 0xb5, 0x49, 0x12, 0xd9, 0x24, 0x8e, 0x99, 0x20, 0x22, 0x62, 0x31,
	0xcf, 0x80, 0xc6, 0x96, 0xfa, 0xe3, 0x6f, 0x87, 0x34, 0xde, 0xe6, 0x1f, 0x48, 0x18, 0xd2, 0xd4,
	0x66, 0x89, 0x42, 0x4c, 0xa2, 0xad, 0x43, 0x68, 0xbc, 0x48, 0x53, 0x96, 0xe2, 0x0d, 0xa8, 0xfb,
	0x2c, 0xa0, 0x3a, 0x32, 0x51, 0xa7, 0xb1, 0x87, 0x2f, 0xfb, 0x0b, 0x0e, 0xa8, 0x8c, 0x29, 0xc3,
	0x3d, 0xb4, 0xe3, 0xaa, 0x3c, 0xd6, 0xe1, 0xaf, 0x11, 0xe5, 0x9c, 0x84, 0x54, 0xaf, 0x99, 0xa8,
	0x33, 0xe7, 0x16, 0x47, 0xeb, 0x08, 0x9a, 0xc7, 0x24, 0x8c, 0xe2, 0x10, 0xff, 0x07, 0x0d, 0xc1,
	0x04, 0x19, 0x66, 0x64, 0x6e, 0x76, 0xc0, 0x2b, 0xd0, 0x64, 0xa7, 0xa7, 0x9c, 0x0a, 0x55, 0xd8,
	0x70, 0xf3, 0x93, 0x44, 0x0f, 0xa3, 0x51, 0x24, 0x74, 0x2d, 0x43, 0xab, 0x83, 0x75, 0x85, 0x60,
	0xe9, 0x98, 0x70, 0x4e, 0xe3, 0x90, 0xa6, 0x07, 0x94, 0x06, 0x1e, 0xf1, 0xcf, 0xf0, 0x03, 0x80,
	0xd3, 0xfc, 0xff, 0xc3, 0x7d, 0x45, 0x3f, 0xe7, 0x56, 0x22, 0xd8, 0x84, 0x96, 0xc7, 0xd8, 0x59,
	0x14, 0x87, 0xcf, 0x65, 0x33, 0x99, 0xc2, 0x6a, 0x48, 0x22, 0x92, 0x82, 0xf6, 0x70, 0x3f, 0xbf,
	0xb3, 0x1a, 0xc2, 0x06, 0xfc, 0x5d, 0x30, 0xea, 0x75, 0x45, 0x50, 0x9e, 0xad, 0xb7, 0xb0, 0xd6,
	0x0f, 0x82, 0x09, 0x5d, 0x2e, 0x7d, 0x7f, 0x4e, 0xb9, 0xc0, 0x4f, 0x2a, 0xa5, 0x52, 0x5c, 0xcb,
	0x69, 0x77, 0xcb, 0xb9, 0x75, 0x27, 0xcb, 0xae, 0x89, 0x5f, 0x42, 0xfb, 0x76, 0x62, 0x9e, 0xb0,
	0x98, 0x53, 0xdc, 0x81, 0x26, 0x95, 0xd3, 0xe0, 0x3a, 0x32, 0xb5, 0x4e, 0xcb, 0x59, 0xac, 0xf0,
	0xaa, 0x31, 0xb9, 0x79, 0xde, 0xfa, 0x82, 0xa0, 0x7d, 0x14, 0x71, 0x71, 0xa7, 0xc8, 0xd2, 0x6f,
	0x54, 0xf1, 0xfb, 0xce, 0xe9, 0x8c, 0x39, 0xaa, 0xcd, 0x74, 0xb4, 0x3e, 0xe1, 0xa8, 0xf5, 0x03,
	0xc1, 0xff, 0x77, 0x48, 0xba, 0x6f, 0x7b, 0xf8, 0x11, 0x34, 0x13, 0xb5, 0x65, 0x4a, 0x67, 0xcb,
	0x59, 0xba, 0x61, 0xb0, 0x4c, 0xb8, 0x39, 0x00, 0xef, 0x40, 0x3d, 0x20, 0x82, 0xe8, 0x9a, 0xa2,
	0x9c, 0x3e, 0x09, 0x85, 0xb4, 0x9e, 0xc2, 0xf2, 0x3e, 0x1d, 0x52, 0x41, 0x65, 0x7c, 0xaf, 0xe2,
	0xd9, 0xec, 0x1e, 0x03, 0x58, 0x19, 0x2f, 0xbd, 0x77, 0x6f, 0x26, 0xb4, 0xca, 0x5d, 0x0e, 0xb8,
	0x5e, 0x33, 0x35, 0xe9, 0x75, 0x25, 0xe4, 0xfc, 0xd4, 0x60, 0xa1, 0xd0, 0xfc, 0x3a, 0x7b, 0x1f,
	0xf0, 0x27, 0x04, 0xad, 0x7e, 0x10, 0x94, 0xdf, 0xc8, 0x46, 0x85, 0x7f, 0xca, 0xb2, 0x1a, 0x0f,
	0x67, 0xe2, 0xb2, 0x06, 0xac, 0xf5, 0xcf, 0xbf, 0x7e, 0x5f, 0xd5, 0xd6, 0xac, 0x79, 0xfb, 0x62,
	0xd7, 0x2e, 0xd4, 0xf4, 0xca, 0xdd, 0xf5, 0xd0, 0x26, 0xfe, 0x08, 0xf3, 0x72, 0xc0, 0x45, 0xef,
	0xb8, 0xca, 0x3d, 0x6d, 0x19, 0x8d, 0xce, 0x6c, 0x60, 0xae, 0x62, 0x55, 0xa9, 0x58, 0xc2, 0x37,
	0x54, 0xc8, 0xbb, 0xbf, 0x23, 0xf8, 0xf7, 0xda, 0x7a, 0x65, 0x81, 0x59, 0x61, 0xbd, 0x75, 0xa0,
	0xc6, 0xfa, 0x14, 0x44, 0x7e, 0xe1, 0xc1, 0x65, 0x7f, 0xeb, 0xfa, 0x7b, 0xc6, 0x0b, 0x19, 0xcc,
	0x2c, 0x02, 0xc6, 0x78, 0x20, 0xd3, 0xb7, 0x39, 0xae, 0x6f, 0xef, 0x1b, 0xba, 0xec, 0x7f, 0x45,
	0x78, 0x04, 0x8b, 0x85, 0x3e, 0x33, 0x7f, 0xda, 0xad, 0x13, 0xf8, 0xe7, 0x55, 0x44, 0x85, 0x79,
	0x32, 0x20, 0xf1, 0xc0, 0x7c, 0xc3, 0xf0, 0xd6, 0x40, 0x88, 0x84, 0xf7, 0x6c, 0x3b, 0x8c, 0xc4,
	0xe0, 0xdc, 0xeb, 0xfa, 0x6c, 0x64, 0x5f, 0x30, 0x21, 0xd3, 0x67, 0x11, 0x15, 0x76, 0x98, 0x12,
	0x6f, 0xdb, 0x63, 0x4c, 0xf8, 0x64, 0x94, 0x18, 0xcb, 0x95, 0xd4, 0xb3, 0x70, 0x44, 0xa2, 0xa1,
	0x2c, 0x70, 0xb4, 0xdd, 0xee, 0xce, 0x66, 0x0d, 0xd5, 0x9c, 0x45, 0x92, 0x24, 0xc3, 0xc8, 0x57,
	0x0f, 0xbe, 0xfd, 0x8e, 0xb3, 0xb8, 0x37, 0x11, 0xf1, 0x9a, 0xea, 0x77, 0xe0, 0xf1, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x7e, 0x6e, 0x8e, 0x61, 0x79, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedbackServiceClient is the client API for FeedbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedbackServiceClient interface {
	AddFeedback(ctx context.Context, in *AddPassengerFeedbackRequest, opts ...grpc.CallOption) (*AddPassengerFeedbackResponse, error)
	ListFeedBack(ctx context.Context, in *ListPassengerFeedbackRequest, opts ...grpc.CallOption) (*ListPassengerFeedbackResponse, error)
	DeleteFeedback(ctx context.Context, in *DeleteFeedBackRequest, opts ...grpc.CallOption) (*DeleteFeedBackResponse, error)
}

type feedbackServiceClient struct {
	cc *grpc.ClientConn
}

func NewFeedbackServiceClient(cc *grpc.ClientConn) FeedbackServiceClient {
	return &feedbackServiceClient{cc}
}

func (c *feedbackServiceClient) AddFeedback(ctx context.Context, in *AddPassengerFeedbackRequest, opts ...grpc.CallOption) (*AddPassengerFeedbackResponse, error) {
	out := new(AddPassengerFeedbackResponse)
	err := c.cc.Invoke(ctx, "/transport.FeedbackService/AddFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackServiceClient) ListFeedBack(ctx context.Context, in *ListPassengerFeedbackRequest, opts ...grpc.CallOption) (*ListPassengerFeedbackResponse, error) {
	out := new(ListPassengerFeedbackResponse)
	err := c.cc.Invoke(ctx, "/transport.FeedbackService/ListFeedBack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackServiceClient) DeleteFeedback(ctx context.Context, in *DeleteFeedBackRequest, opts ...grpc.CallOption) (*DeleteFeedBackResponse, error) {
	out := new(DeleteFeedBackResponse)
	err := c.cc.Invoke(ctx, "/transport.FeedbackService/DeleteFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedbackServiceServer is the server API for FeedbackService service.
type FeedbackServiceServer interface {
	AddFeedback(context.Context, *AddPassengerFeedbackRequest) (*AddPassengerFeedbackResponse, error)
	ListFeedBack(context.Context, *ListPassengerFeedbackRequest) (*ListPassengerFeedbackResponse, error)
	DeleteFeedback(context.Context, *DeleteFeedBackRequest) (*DeleteFeedBackResponse, error)
}

// UnimplementedFeedbackServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFeedbackServiceServer struct {
}

func (*UnimplementedFeedbackServiceServer) AddFeedback(ctx context.Context, req *AddPassengerFeedbackRequest) (*AddPassengerFeedbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFeedback not implemented")
}
func (*UnimplementedFeedbackServiceServer) ListFeedBack(ctx context.Context, req *ListPassengerFeedbackRequest) (*ListPassengerFeedbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeedBack not implemented")
}
func (*UnimplementedFeedbackServiceServer) DeleteFeedback(ctx context.Context, req *DeleteFeedBackRequest) (*DeleteFeedBackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFeedback not implemented")
}

func RegisterFeedbackServiceServer(s *grpc.Server, srv FeedbackServiceServer) {
	s.RegisterService(&_FeedbackService_serviceDesc, srv)
}

func _FeedbackService_AddFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPassengerFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).AddFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.FeedbackService/AddFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).AddFeedback(ctx, req.(*AddPassengerFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedbackService_ListFeedBack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPassengerFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).ListFeedBack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.FeedbackService/ListFeedBack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).ListFeedBack(ctx, req.(*ListPassengerFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedbackService_DeleteFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeedBackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).DeleteFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.FeedbackService/DeleteFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).DeleteFeedback(ctx, req.(*DeleteFeedBackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedbackService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transport.FeedbackService",
	HandlerType: (*FeedbackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFeedback",
			Handler:    _FeedbackService_AddFeedback_Handler,
		},
		{
			MethodName: "ListFeedBack",
			Handler:    _FeedbackService_ListFeedBack_Handler,
		},
		{
			MethodName: "DeleteFeedback",
			Handler:    _FeedbackService_DeleteFeedback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport/service.proto",
}
