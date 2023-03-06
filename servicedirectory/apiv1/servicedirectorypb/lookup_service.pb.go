// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: google/cloud/servicedirectory/v1/lookup_service.proto

package servicedirectorypb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message for
// [LookupService.ResolveService][google.cloud.servicedirectory.v1.LookupService.ResolveService].
// Looks up a service by its name, returns the service and its endpoints.
type ResolveServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the service to resolve.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The maximum number of endpoints to return. Defaults to 25.
	// Maximum is 100. If a value less than one is specified, the Default is used.
	// If a value greater than the Maximum is specified, the Maximum is used.
	MaxEndpoints int32 `protobuf:"varint,2,opt,name=max_endpoints,json=maxEndpoints,proto3" json:"max_endpoints,omitempty"`
	// Optional. The filter applied to the endpoints of the resolved service.
	//
	// General filter string syntax:
	// <field> <operator> <value> (<logical connector>)
	// <field> can be "name" or "metadata.<key>" for map field.
	// <operator> can be "<, >, <=, >=, !=, =, :". Of which ":" means HAS and is
	// roughly the same as "=".
	// <value> must be the same data type as the field.
	// <logical connector> can be "AND, OR, NOT".
	//
	// Examples of valid filters:
	// * "metadata.owner" returns Endpoints that have a label with the
	//   key "owner", this is the same as "metadata:owner"
	// * "metadata.protocol=gRPC" returns Endpoints that have key/value
	//   "protocol=gRPC"
	// * "metadata.owner!=sd AND metadata.foo=bar" returns
	//   Endpoints that have "owner" field in metadata with a value that is not
	//   "sd" AND have the key/value foo=bar.
	EndpointFilter string `protobuf:"bytes,3,opt,name=endpoint_filter,json=endpointFilter,proto3" json:"endpoint_filter,omitempty"`
}

func (x *ResolveServiceRequest) Reset() {
	*x = ResolveServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_servicedirectory_v1_lookup_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveServiceRequest) ProtoMessage() {}

func (x *ResolveServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_servicedirectory_v1_lookup_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveServiceRequest.ProtoReflect.Descriptor instead.
func (*ResolveServiceRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_servicedirectory_v1_lookup_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResolveServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResolveServiceRequest) GetMaxEndpoints() int32 {
	if x != nil {
		return x.MaxEndpoints
	}
	return 0
}

func (x *ResolveServiceRequest) GetEndpointFilter() string {
	if x != nil {
		return x.EndpointFilter
	}
	return ""
}

// The response message for
// [LookupService.ResolveService][google.cloud.servicedirectory.v1.LookupService.ResolveService].
type ResolveServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *ResolveServiceResponse) Reset() {
	*x = ResolveServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_servicedirectory_v1_lookup_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveServiceResponse) ProtoMessage() {}

func (x *ResolveServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_servicedirectory_v1_lookup_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveServiceResponse.ProtoReflect.Descriptor instead.
func (*ResolveServiceResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_servicedirectory_v1_lookup_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResolveServiceResponse) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

var File_google_cloud_servicedirectory_v1_lookup_service_proto protoreflect.FileDescriptor

var file_google_cloud_servicedirectory_v1_lookup_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a,
	0x15, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0d, 0x6d,
	0x61, 0x78, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x0f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x0e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x22, 0x5d, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x32, 0xb8, 0x02, 0x0a, 0x0d, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xd1, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x46, 0x22, 0x41, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f,
	0x2a, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x3a, 0x01, 0x2a, 0x1a, 0x53, 0xca, 0x41, 0x1f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xfd, 0x01,
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0xf8, 0x01,
	0x01, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_servicedirectory_v1_lookup_service_proto_rawDescOnce sync.Once
	file_google_cloud_servicedirectory_v1_lookup_service_proto_rawDescData = file_google_cloud_servicedirectory_v1_lookup_service_proto_rawDesc
)

func file_google_cloud_servicedirectory_v1_lookup_service_proto_rawDescGZIP() []byte {
	file_google_cloud_servicedirectory_v1_lookup_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_servicedirectory_v1_lookup_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_servicedirectory_v1_lookup_service_proto_rawDescData)
	})
	return file_google_cloud_servicedirectory_v1_lookup_service_proto_rawDescData
}

var file_google_cloud_servicedirectory_v1_lookup_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_servicedirectory_v1_lookup_service_proto_goTypes = []interface{}{
	(*ResolveServiceRequest)(nil),  // 0: google.cloud.servicedirectory.v1.ResolveServiceRequest
	(*ResolveServiceResponse)(nil), // 1: google.cloud.servicedirectory.v1.ResolveServiceResponse
	(*Service)(nil),                // 2: google.cloud.servicedirectory.v1.Service
}
var file_google_cloud_servicedirectory_v1_lookup_service_proto_depIdxs = []int32{
	2, // 0: google.cloud.servicedirectory.v1.ResolveServiceResponse.service:type_name -> google.cloud.servicedirectory.v1.Service
	0, // 1: google.cloud.servicedirectory.v1.LookupService.ResolveService:input_type -> google.cloud.servicedirectory.v1.ResolveServiceRequest
	1, // 2: google.cloud.servicedirectory.v1.LookupService.ResolveService:output_type -> google.cloud.servicedirectory.v1.ResolveServiceResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_servicedirectory_v1_lookup_service_proto_init() }
func file_google_cloud_servicedirectory_v1_lookup_service_proto_init() {
	if File_google_cloud_servicedirectory_v1_lookup_service_proto != nil {
		return
	}
	file_google_cloud_servicedirectory_v1_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_servicedirectory_v1_lookup_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveServiceRequest); i {
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
		file_google_cloud_servicedirectory_v1_lookup_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveServiceResponse); i {
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
			RawDescriptor: file_google_cloud_servicedirectory_v1_lookup_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_servicedirectory_v1_lookup_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_servicedirectory_v1_lookup_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_servicedirectory_v1_lookup_service_proto_msgTypes,
	}.Build()
	File_google_cloud_servicedirectory_v1_lookup_service_proto = out.File
	file_google_cloud_servicedirectory_v1_lookup_service_proto_rawDesc = nil
	file_google_cloud_servicedirectory_v1_lookup_service_proto_goTypes = nil
	file_google_cloud_servicedirectory_v1_lookup_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LookupServiceClient is the client API for LookupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LookupServiceClient interface {
	// Returns a [service][google.cloud.servicedirectory.v1.Service] and its
	// associated endpoints.
	// Resolving a service is not considered an active developer method.
	ResolveService(ctx context.Context, in *ResolveServiceRequest, opts ...grpc.CallOption) (*ResolveServiceResponse, error)
}

type lookupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLookupServiceClient(cc grpc.ClientConnInterface) LookupServiceClient {
	return &lookupServiceClient{cc}
}

func (c *lookupServiceClient) ResolveService(ctx context.Context, in *ResolveServiceRequest, opts ...grpc.CallOption) (*ResolveServiceResponse, error) {
	out := new(ResolveServiceResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.servicedirectory.v1.LookupService/ResolveService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LookupServiceServer is the server API for LookupService service.
type LookupServiceServer interface {
	// Returns a [service][google.cloud.servicedirectory.v1.Service] and its
	// associated endpoints.
	// Resolving a service is not considered an active developer method.
	ResolveService(context.Context, *ResolveServiceRequest) (*ResolveServiceResponse, error)
}

// UnimplementedLookupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLookupServiceServer struct {
}

func (*UnimplementedLookupServiceServer) ResolveService(context.Context, *ResolveServiceRequest) (*ResolveServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveService not implemented")
}

func RegisterLookupServiceServer(s *grpc.Server, srv LookupServiceServer) {
	s.RegisterService(&_LookupService_serviceDesc, srv)
}

func _LookupService_ResolveService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServiceServer).ResolveService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.servicedirectory.v1.LookupService/ResolveService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServiceServer).ResolveService(ctx, req.(*ResolveServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LookupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.servicedirectory.v1.LookupService",
	HandlerType: (*LookupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResolveService",
			Handler:    _LookupService_ResolveService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/servicedirectory/v1/lookup_service.proto",
}
