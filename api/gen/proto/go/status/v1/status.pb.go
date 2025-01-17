// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: status/v1/status.proto

package statusv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetBuildInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBuildInfoRequest) Reset() {
	*x = GetBuildInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_v1_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBuildInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBuildInfoRequest) ProtoMessage() {}

func (x *GetBuildInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_status_v1_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBuildInfoRequest.ProtoReflect.Descriptor instead.
func (*GetBuildInfoRequest) Descriptor() ([]byte, []int) {
	return file_status_v1_status_proto_rawDescGZIP(), []int{0}
}

type GetBuildInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string            `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data   *GetBuildInfoData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetBuildInfoResponse) Reset() {
	*x = GetBuildInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_v1_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBuildInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBuildInfoResponse) ProtoMessage() {}

func (x *GetBuildInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_status_v1_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBuildInfoResponse.ProtoReflect.Descriptor instead.
func (*GetBuildInfoResponse) Descriptor() ([]byte, []int) {
	return file_status_v1_status_proto_rawDescGZIP(), []int{1}
}

func (x *GetBuildInfoResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetBuildInfoResponse) GetData() *GetBuildInfoData {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetBuildInfoData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Revision  string `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
	Branch    string `protobuf:"bytes,3,opt,name=branch,proto3" json:"branch,omitempty"`
	BuildUser string `protobuf:"bytes,4,opt,name=build_user,json=buildUser,proto3" json:"build_user,omitempty"`
	BuildDate string `protobuf:"bytes,5,opt,name=build_date,json=buildDate,proto3" json:"build_date,omitempty"`
	GoVersion string `protobuf:"bytes,6,opt,name=go_version,json=goVersion,proto3" json:"go_version,omitempty"`
}

func (x *GetBuildInfoData) Reset() {
	*x = GetBuildInfoData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_v1_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBuildInfoData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBuildInfoData) ProtoMessage() {}

func (x *GetBuildInfoData) ProtoReflect() protoreflect.Message {
	mi := &file_status_v1_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBuildInfoData.ProtoReflect.Descriptor instead.
func (*GetBuildInfoData) Descriptor() ([]byte, []int) {
	return file_status_v1_status_proto_rawDescGZIP(), []int{2}
}

func (x *GetBuildInfoData) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GetBuildInfoData) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *GetBuildInfoData) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *GetBuildInfoData) GetBuildUser() string {
	if x != nil {
		return x.BuildUser
	}
	return ""
}

func (x *GetBuildInfoData) GetBuildDate() string {
	if x != nil {
		return x.BuildDate
	}
	return ""
}

func (x *GetBuildInfoData) GetGoVersion() string {
	if x != nil {
		return x.GoVersion
	}
	return ""
}

type GetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_v1_status_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_status_v1_status_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigRequest.ProtoReflect.Descriptor instead.
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return file_status_v1_status_proto_rawDescGZIP(), []int{3}
}

type GetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConfigResponse) Reset() {
	*x = GetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_v1_status_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigResponse) ProtoMessage() {}

func (x *GetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_status_v1_status_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigResponse.ProtoReflect.Descriptor instead.
func (*GetConfigResponse) Descriptor() ([]byte, []int) {
	return file_status_v1_status_proto_rawDescGZIP(), []int{4}
}

var File_status_v1_status_proto protoreflect.FileDescriptor

var file_status_v1_status_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x5f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xbd, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x6f, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb7, 0x03,
	0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x71, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x5d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f,
	0x64, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x66, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x69, 0x66, 0x66, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74,
	0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x69, 0x66, 0x66, 0x12, 0x6c, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79,
	0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0xa3, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x79, 0x72,
	0x6f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58,
	0xaa, 0x02, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_status_v1_status_proto_rawDescOnce sync.Once
	file_status_v1_status_proto_rawDescData = file_status_v1_status_proto_rawDesc
)

func file_status_v1_status_proto_rawDescGZIP() []byte {
	file_status_v1_status_proto_rawDescOnce.Do(func() {
		file_status_v1_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_status_v1_status_proto_rawDescData)
	})
	return file_status_v1_status_proto_rawDescData
}

var file_status_v1_status_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_status_v1_status_proto_goTypes = []any{
	(*GetBuildInfoRequest)(nil),  // 0: status.v1.GetBuildInfoRequest
	(*GetBuildInfoResponse)(nil), // 1: status.v1.GetBuildInfoResponse
	(*GetBuildInfoData)(nil),     // 2: status.v1.GetBuildInfoData
	(*GetConfigRequest)(nil),     // 3: status.v1.GetConfigRequest
	(*GetConfigResponse)(nil),    // 4: status.v1.GetConfigResponse
	(*httpbody.HttpBody)(nil),    // 5: google.api.HttpBody
}
var file_status_v1_status_proto_depIdxs = []int32{
	2, // 0: status.v1.GetBuildInfoResponse.data:type_name -> status.v1.GetBuildInfoData
	0, // 1: status.v1.StatusService.GetBuildInfo:input_type -> status.v1.GetBuildInfoRequest
	3, // 2: status.v1.StatusService.GetConfig:input_type -> status.v1.GetConfigRequest
	3, // 3: status.v1.StatusService.GetDiffConfig:input_type -> status.v1.GetConfigRequest
	3, // 4: status.v1.StatusService.GetDefaultConfig:input_type -> status.v1.GetConfigRequest
	1, // 5: status.v1.StatusService.GetBuildInfo:output_type -> status.v1.GetBuildInfoResponse
	5, // 6: status.v1.StatusService.GetConfig:output_type -> google.api.HttpBody
	5, // 7: status.v1.StatusService.GetDiffConfig:output_type -> google.api.HttpBody
	5, // 8: status.v1.StatusService.GetDefaultConfig:output_type -> google.api.HttpBody
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_status_v1_status_proto_init() }
func file_status_v1_status_proto_init() {
	if File_status_v1_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_status_v1_status_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetBuildInfoRequest); i {
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
		file_status_v1_status_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetBuildInfoResponse); i {
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
		file_status_v1_status_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetBuildInfoData); i {
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
		file_status_v1_status_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetConfigRequest); i {
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
		file_status_v1_status_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetConfigResponse); i {
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
			RawDescriptor: file_status_v1_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_status_v1_status_proto_goTypes,
		DependencyIndexes: file_status_v1_status_proto_depIdxs,
		MessageInfos:      file_status_v1_status_proto_msgTypes,
	}.Build()
	File_status_v1_status_proto = out.File
	file_status_v1_status_proto_rawDesc = nil
	file_status_v1_status_proto_goTypes = nil
	file_status_v1_status_proto_depIdxs = nil
}
