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
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: api/proto/v1/service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetTrackingRequest set a tracking identifier
type GetTrackingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tracking_id is the identifier for tracking
	TrackingId string `protobuf:"bytes,1,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
}

func (x *GetTrackingRequest) Reset() {
	*x = GetTrackingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrackingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrackingRequest) ProtoMessage() {}

func (x *GetTrackingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrackingRequest.ProtoReflect.Descriptor instead.
func (*GetTrackingRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetTrackingRequest) GetTrackingId() string {
	if x != nil {
		return x.TrackingId
	}
	return ""
}

// Tracking contains information about a shipment
type Tracking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tracking_id is the identifier for tracking
	TrackingId string `protobuf:"bytes,1,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
	// status provides the current status of the shipment
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// created_time is when the tracking resource was created
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// signed is populated when the shipment is delivered
	Signed string `protobuf:"bytes,5,opt,name=signed,proto3" json:"signed,omitempty"`
	// updated_time is when the resource was last updated
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// weight of the shipment
	Weight string `protobuf:"bytes,6,opt,name=weight,proto3" json:"weight,omitempty"`
	// est_delivery_date estimated delivery time for shipment
	EstDeliveryTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=est_delivery_time,json=estDeliveryTime,proto3" json:"est_delivery_time,omitempty"`
	// carrier of the shipmemt
	Carrier string `protobuf:"bytes,8,opt,name=carrier,proto3" json:"carrier,omitempty"`
}

func (x *Tracking) Reset() {
	*x = Tracking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tracking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tracking) ProtoMessage() {}

func (x *Tracking) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tracking.ProtoReflect.Descriptor instead.
func (*Tracking) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *Tracking) GetTrackingId() string {
	if x != nil {
		return x.TrackingId
	}
	return ""
}

func (x *Tracking) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Tracking) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Tracking) GetSigned() string {
	if x != nil {
		return x.Signed
	}
	return ""
}

func (x *Tracking) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Tracking) GetWeight() string {
	if x != nil {
		return x.Weight
	}
	return ""
}

func (x *Tracking) GetEstDeliveryTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EstDeliveryTime
	}
	return nil
}

func (x *Tracking) GetCarrier() string {
	if x != nil {
		return x.Carrier
	}
	return ""
}

// ListTrackingResponse contains a list of tracking information
type ListTrackingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// trackings contains a repeated list of tracking information
	Trackings []*Tracking `protobuf:"bytes,1,rep,name=trackings,proto3" json:"trackings,omitempty"`
	// next_page_token is used for pagination
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListTrackingResponse) Reset() {
	*x = ListTrackingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTrackingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTrackingResponse) ProtoMessage() {}

func (x *ListTrackingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTrackingResponse.ProtoReflect.Descriptor instead.
func (*ListTrackingResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListTrackingResponse) GetTrackings() []*Tracking {
	if x != nil {
		return x.Trackings
	}
	return nil
}

func (x *ListTrackingResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_api_proto_v1_service_proto protoreflect.FileDescriptor

var file_api_proto_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22,
	0xcf, 0x02, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x46, 0x0a, 0x11, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x65, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0x22, 0x76, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xb5, 0x02, 0x0a, 0x08, 0x53, 0x68,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x6f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x61,
	0x70, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x62, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x24, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x54, 0x0a, 0x0e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x2e,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f,
	0x61, 0x70, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x06, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_v1_service_proto_rawDescOnce sync.Once
	file_api_proto_v1_service_proto_rawDescData = file_api_proto_v1_service_proto_rawDesc
)

func file_api_proto_v1_service_proto_rawDescGZIP() []byte {
	file_api_proto_v1_service_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1_service_proto_rawDescData)
	})
	return file_api_proto_v1_service_proto_rawDescData
}

var file_api_proto_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_proto_v1_service_proto_goTypes = []interface{}{
	(*GetTrackingRequest)(nil),    // 0: sample_apps.v1.GetTrackingRequest
	(*Tracking)(nil),              // 1: sample_apps.v1.Tracking
	(*ListTrackingResponse)(nil),  // 2: sample_apps.v1.ListTrackingResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_api_proto_v1_service_proto_depIdxs = []int32{
	3, // 0: sample_apps.v1.Tracking.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: sample_apps.v1.Tracking.update_time:type_name -> google.protobuf.Timestamp
	3, // 2: sample_apps.v1.Tracking.est_delivery_time:type_name -> google.protobuf.Timestamp
	1, // 3: sample_apps.v1.ListTrackingResponse.trackings:type_name -> sample_apps.v1.Tracking
	0, // 4: sample_apps.v1.Shipment.GetTracking:input_type -> sample_apps.v1.GetTrackingRequest
	4, // 5: sample_apps.v1.Shipment.ListTracking:input_type -> google.protobuf.Empty
	0, // 6: sample_apps.v1.Shipment.NotifyTracking:input_type -> sample_apps.v1.GetTrackingRequest
	1, // 7: sample_apps.v1.Shipment.GetTracking:output_type -> sample_apps.v1.Tracking
	2, // 8: sample_apps.v1.Shipment.ListTracking:output_type -> sample_apps.v1.ListTrackingResponse
	1, // 9: sample_apps.v1.Shipment.NotifyTracking:output_type -> sample_apps.v1.Tracking
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_v1_service_proto_init() }
func file_api_proto_v1_service_proto_init() {
	if File_api_proto_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrackingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tracking); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTrackingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_v1_service_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_service_proto_depIdxs,
		MessageInfos:      file_api_proto_v1_service_proto_msgTypes,
	}.Build()
	File_api_proto_v1_service_proto = out.File
	file_api_proto_v1_service_proto_rawDesc = nil
	file_api_proto_v1_service_proto_goTypes = nil
	file_api_proto_v1_service_proto_depIdxs = nil
}
