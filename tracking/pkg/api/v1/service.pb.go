// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type TrackingRequest struct {
	TrackingId           string   `protobuf:"bytes,1,opt,name=trackingId,proto3" json:"trackingId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackingRequest) Reset()         { *m = TrackingRequest{} }
func (m *TrackingRequest) String() string { return proto.CompactTextString(m) }
func (*TrackingRequest) ProtoMessage()    {}
func (*TrackingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *TrackingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackingRequest.Unmarshal(m, b)
}
func (m *TrackingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackingRequest.Marshal(b, m, deterministic)
}
func (m *TrackingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackingRequest.Merge(m, src)
}
func (m *TrackingRequest) XXX_Size() int {
	return xxx_messageInfo_TrackingRequest.Size(m)
}
func (m *TrackingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TrackingRequest proto.InternalMessageInfo

func (m *TrackingRequest) GetTrackingId() string {
	if m != nil {
		return m.TrackingId
	}
	return ""
}

type TrackingResponse struct {
	TrackingId           string   `protobuf:"bytes,1,opt,name=trackingId,proto3" json:"trackingId,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Created              string   `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Signed               string   `protobuf:"bytes,5,opt,name=signed,proto3" json:"signed,omitempty"`
	Updated              string   `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Weight               string   `protobuf:"bytes,6,opt,name=weight,proto3" json:"weight,omitempty"`
	EstDeliveryDate      string   `protobuf:"bytes,7,opt,name=est_delivery_date,json=estDeliveryDate,proto3" json:"est_delivery_date,omitempty"`
	Carrier              string   `protobuf:"bytes,8,opt,name=carrier,proto3" json:"carrier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackingResponse) Reset()         { *m = TrackingResponse{} }
func (m *TrackingResponse) String() string { return proto.CompactTextString(m) }
func (*TrackingResponse) ProtoMessage()    {}
func (*TrackingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *TrackingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackingResponse.Unmarshal(m, b)
}
func (m *TrackingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackingResponse.Marshal(b, m, deterministic)
}
func (m *TrackingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackingResponse.Merge(m, src)
}
func (m *TrackingResponse) XXX_Size() int {
	return xxx_messageInfo_TrackingResponse.Size(m)
}
func (m *TrackingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TrackingResponse proto.InternalMessageInfo

func (m *TrackingResponse) GetTrackingId() string {
	if m != nil {
		return m.TrackingId
	}
	return ""
}

func (m *TrackingResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TrackingResponse) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *TrackingResponse) GetSigned() string {
	if m != nil {
		return m.Signed
	}
	return ""
}

func (m *TrackingResponse) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

func (m *TrackingResponse) GetWeight() string {
	if m != nil {
		return m.Weight
	}
	return ""
}

func (m *TrackingResponse) GetEstDeliveryDate() string {
	if m != nil {
		return m.EstDeliveryDate
	}
	return ""
}

func (m *TrackingResponse) GetCarrier() string {
	if m != nil {
		return m.Carrier
	}
	return ""
}

type TrackingListResponse struct {
	TrackingResponse     []*TrackingResponse `protobuf:"bytes,1,rep,name=trackingResponse,proto3" json:"trackingResponse,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TrackingListResponse) Reset()         { *m = TrackingListResponse{} }
func (m *TrackingListResponse) String() string { return proto.CompactTextString(m) }
func (*TrackingListResponse) ProtoMessage()    {}
func (*TrackingListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *TrackingListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackingListResponse.Unmarshal(m, b)
}
func (m *TrackingListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackingListResponse.Marshal(b, m, deterministic)
}
func (m *TrackingListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackingListResponse.Merge(m, src)
}
func (m *TrackingListResponse) XXX_Size() int {
	return xxx_messageInfo_TrackingListResponse.Size(m)
}
func (m *TrackingListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackingListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TrackingListResponse proto.InternalMessageInfo

func (m *TrackingListResponse) GetTrackingResponse() []*TrackingResponse {
	if m != nil {
		return m.TrackingResponse
	}
	return nil
}

func init() {
	proto.RegisterType((*TrackingRequest)(nil), "v1.TrackingRequest")
	proto.RegisterType((*TrackingResponse)(nil), "v1.TrackingResponse")
	proto.RegisterType((*TrackingListResponse)(nil), "v1.TrackingListResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4e, 0x02, 0x31,
	0x10, 0x86, 0x59, 0x50, 0xc0, 0x31, 0x06, 0x2c, 0x84, 0x34, 0x98, 0x18, 0xb2, 0x27, 0xe2, 0xa1,
	0x04, 0x7c, 0x01, 0x4d, 0x30, 0x86, 0xc4, 0x13, 0xf1, 0xe0, 0x8d, 0x14, 0x76, 0x5c, 0x1b, 0x91,
	0x5d, 0xdb, 0xd9, 0x35, 0xbc, 0x89, 0x0f, 0xea, 0x03, 0x98, 0x6d, 0xbb, 0x8a, 0x68, 0xe2, 0xf1,
	0x9f, 0xfe, 0xdf, 0x9f, 0x99, 0xe9, 0xc0, 0x89, 0x41, 0x9d, 0xab, 0x15, 0x8a, 0x54, 0x27, 0x94,
	0xb0, 0x6a, 0x3e, 0xee, 0x9f, 0xc5, 0x49, 0x12, 0xaf, 0x71, 0x64, 0x2b, 0xcb, 0xec, 0x71, 0x84,
	0x2f, 0x29, 0x6d, 0x9d, 0x21, 0x1c, 0x43, 0xeb, 0x5e, 0xcb, 0xd5, 0xb3, 0xda, 0xc4, 0x73, 0x7c,
	0xcd, 0xd0, 0x10, 0x3b, 0x07, 0x20, 0x5f, 0x9a, 0x45, 0x3c, 0x18, 0x04, 0xc3, 0xa3, 0xf9, 0x4e,
	0x25, 0xfc, 0x08, 0xa0, 0xfd, 0xcd, 0x98, 0x34, 0xd9, 0x18, 0xfc, 0x0f, 0x62, 0x3d, 0xa8, 0x1b,
	0x92, 0x94, 0x19, 0x5e, 0xb5, 0x6f, 0x5e, 0x31, 0x0e, 0x8d, 0x95, 0x46, 0x49, 0x18, 0xf1, 0x9a,
	0x7d, 0x28, 0xa5, 0x25, 0x54, 0xbc, 0xc1, 0x88, 0x1f, 0x7a, 0xc2, 0xaa, 0x82, 0xc8, 0xd2, 0xc8,
	0x12, 0x07, 0x8e, 0xf0, 0xb2, 0x20, 0xde, 0x50, 0xc5, 0x4f, 0xc4, 0xeb, 0x8e, 0x70, 0x8a, 0x5d,
	0xc0, 0x29, 0x1a, 0x5a, 0x44, 0xb8, 0x56, 0x39, 0xea, 0xed, 0xa2, 0x70, 0xf3, 0x86, 0xb5, 0xb4,
	0xd0, 0xd0, 0xd4, 0xd7, 0xa7, 0x92, 0xd0, 0xf6, 0x23, 0xb5, 0x56, 0xa8, 0x79, 0xd3, 0xf7, 0xe3,
	0x64, 0xf8, 0x00, 0xdd, 0x72, 0xea, 0x3b, 0x65, 0xe8, 0x6b, 0xf2, 0x2b, 0x68, 0xd3, 0xde, 0x36,
	0x78, 0x30, 0xa8, 0x0d, 0x8f, 0x27, 0x5d, 0x91, 0x8f, 0xc5, 0xfe, 0xa6, 0xe6, 0xbf, 0xdc, 0x93,
	0xf7, 0x00, 0x9a, 0xa5, 0x8d, 0x5d, 0x03, 0xbb, 0x45, 0x2a, 0xe5, 0x14, 0x49, 0xaa, 0xb5, 0x61,
	0x9d, 0x9f, 0x51, 0xf6, 0xa3, 0xfa, 0x7f, 0xe6, 0x87, 0x15, 0x36, 0x83, 0x4e, 0xd1, 0xe1, 0x7e,
	0x46, 0x4f, 0xb8, 0x43, 0x10, 0xe5, 0x21, 0x88, 0x9b, 0xe2, 0x10, 0xfa, 0x7c, 0x37, 0x66, 0x77,
	0xb4, 0xb0, 0xb2, 0xac, 0x5b, 0xef, 0xe5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x40, 0x32,
	0x45, 0x57, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TrackingClient is the client API for Tracking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrackingClient interface {
	GetTrackingDetails(ctx context.Context, in *TrackingRequest, opts ...grpc.CallOption) (*TrackingResponse, error)
	ListTrackingDetails(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TrackingListResponse, error)
}

type trackingClient struct {
	cc *grpc.ClientConn
}

func NewTrackingClient(cc *grpc.ClientConn) TrackingClient {
	return &trackingClient{cc}
}

func (c *trackingClient) GetTrackingDetails(ctx context.Context, in *TrackingRequest, opts ...grpc.CallOption) (*TrackingResponse, error) {
	out := new(TrackingResponse)
	err := c.cc.Invoke(ctx, "/v1.Tracking/GetTrackingDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackingClient) ListTrackingDetails(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TrackingListResponse, error) {
	out := new(TrackingListResponse)
	err := c.cc.Invoke(ctx, "/v1.Tracking/ListTrackingDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackingServer is the server API for Tracking service.
type TrackingServer interface {
	GetTrackingDetails(context.Context, *TrackingRequest) (*TrackingResponse, error)
	ListTrackingDetails(context.Context, *empty.Empty) (*TrackingListResponse, error)
}

// UnimplementedTrackingServer can be embedded to have forward compatible implementations.
type UnimplementedTrackingServer struct {
}

func (*UnimplementedTrackingServer) GetTrackingDetails(ctx context.Context, req *TrackingRequest) (*TrackingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrackingDetails not implemented")
}
func (*UnimplementedTrackingServer) ListTrackingDetails(ctx context.Context, req *empty.Empty) (*TrackingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTrackingDetails not implemented")
}

func RegisterTrackingServer(s *grpc.Server, srv TrackingServer) {
	s.RegisterService(&_Tracking_serviceDesc, srv)
}

func _Tracking_GetTrackingDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingServer).GetTrackingDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Tracking/GetTrackingDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingServer).GetTrackingDetails(ctx, req.(*TrackingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracking_ListTrackingDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingServer).ListTrackingDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Tracking/ListTrackingDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingServer).ListTrackingDetails(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tracking_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Tracking",
	HandlerType: (*TrackingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTrackingDetails",
			Handler:    _Tracking_GetTrackingDetails_Handler,
		},
		{
			MethodName: "ListTrackingDetails",
			Handler:    _Tracking_ListTrackingDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
