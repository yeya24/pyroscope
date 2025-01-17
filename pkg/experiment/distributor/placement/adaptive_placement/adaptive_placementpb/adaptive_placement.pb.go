// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: experiment/distributor/placement/adaptive_placement/adaptive_placementpb/adaptive_placement.proto

package adaptive_placementpb

import (
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

type LoadBalancing int32

const (
	LoadBalancing_LOAD_BALANCING_UNSPECIFIED LoadBalancing = 0
	LoadBalancing_LOAD_BALANCING_FINGERPRINT LoadBalancing = 1
	LoadBalancing_LOAD_BALANCING_ROUND_ROBIN LoadBalancing = 2
)

// Enum value maps for LoadBalancing.
var (
	LoadBalancing_name = map[int32]string{
		0: "LOAD_BALANCING_UNSPECIFIED",
		1: "LOAD_BALANCING_FINGERPRINT",
		2: "LOAD_BALANCING_ROUND_ROBIN",
	}
	LoadBalancing_value = map[string]int32{
		"LOAD_BALANCING_UNSPECIFIED": 0,
		"LOAD_BALANCING_FINGERPRINT": 1,
		"LOAD_BALANCING_ROUND_ROBIN": 2,
	}
)

func (x LoadBalancing) Enum() *LoadBalancing {
	p := new(LoadBalancing)
	*p = x
	return p
}

func (x LoadBalancing) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoadBalancing) Descriptor() protoreflect.EnumDescriptor {
	return file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_enumTypes[0].Descriptor()
}

func (LoadBalancing) Type() protoreflect.EnumType {
	return &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_enumTypes[0]
}

func (x LoadBalancing) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoadBalancing.Descriptor instead.
func (LoadBalancing) EnumDescriptor() ([]byte, []int) {
	return file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescGZIP(), []int{0}
}

// DistributionStats includes the data the Placement is built based on.
type DistributionStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenants   []*TenantStats  `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
	Datasets  []*DatasetStats `protobuf:"bytes,2,rep,name=datasets,proto3" json:"datasets,omitempty"`
	Shards    []*ShardStats   `protobuf:"bytes,3,rep,name=shards,proto3" json:"shards,omitempty"`
	CreatedAt int64           `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *DistributionStats) Reset() {
	*x = DistributionStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributionStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributionStats) ProtoMessage() {}

func (x *DistributionStats) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributionStats.ProtoReflect.Descriptor instead.
func (*DistributionStats) Descriptor() ([]byte, []int) {
	return file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescGZIP(), []int{0}
}

func (x *DistributionStats) GetTenants() []*TenantStats {
	if x != nil {
		return x.Tenants
	}
	return nil
}

func (x *DistributionStats) GetDatasets() []*DatasetStats {
	if x != nil {
		return x.Datasets
	}
	return nil
}

func (x *DistributionStats) GetShards() []*ShardStats {
	if x != nil {
		return x.Shards
	}
	return nil
}

func (x *DistributionStats) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type TenantStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *TenantStats) Reset() {
	*x = TenantStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantStats) ProtoMessage() {}

func (x *TenantStats) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantStats.ProtoReflect.Descriptor instead.
func (*TenantStats) Descriptor() ([]byte, []int) {
	return file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescGZIP(), []int{1}
}

func (x *TenantStats) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type DatasetStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant uint32 `protobuf:"varint,1,opt,name=tenant,proto3" json:"tenant,omitempty"` // Reference to TenantStats.
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Shard value is a reference to ShardStats.
	Shards []uint32 `protobuf:"varint,3,rep,packed,name=shards,proto3" json:"shards,omitempty"`
	// Data rate in bytes per second for each shard.
	// The dataset size is measured after being encoded
	// in the block wire format.
	Usage []uint64 `protobuf:"varint,4,rep,packed,name=usage,proto3" json:"usage,omitempty"`
	// Standard deviation of the data rate across shards
	// aggregated within a sliding time window.
	StdDev uint64 `protobuf:"varint,5,opt,name=std_dev,json=stdDev,proto3" json:"std_dev,omitempty"`
}

func (x *DatasetStats) Reset() {
	*x = DatasetStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetStats) ProtoMessage() {}

func (x *DatasetStats) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetStats.ProtoReflect.Descriptor instead.
func (*DatasetStats) Descriptor() ([]byte, []int) {
	return file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescGZIP(), []int{2}
}

func (x *DatasetStats) GetTenant() uint32 {
	if x != nil {
		return x.Tenant
	}
	return 0
}

func (x *DatasetStats) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DatasetStats) GetShards() []uint32 {
	if x != nil {
		return x.Shards
	}
	return nil
}

func (x *DatasetStats) GetUsage() []uint64 {
	if x != nil {
		return x.Usage
	}
	return nil
}

func (x *DatasetStats) GetStdDev() uint64 {
	if x != nil {
		return x.StdDev
	}
	return 0
}

type ShardStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Shard ID.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Owner represents the node that hosted the shard.
	// There may be multiple entries for a single shard
	// if it was relocated across different nodes.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *ShardStats) Reset() {
	*x = ShardStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardStats) ProtoMessage() {}

func (x *ShardStats) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardStats.ProtoReflect.Descriptor instead.
func (*ShardStats) Descriptor() ([]byte, []int) {
	return file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescGZIP(), []int{3}
}

func (x *ShardStats) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ShardStats) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type PlacementRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenants   []*TenantPlacement  `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
	Datasets  []*DatasetPlacement `protobuf:"bytes,2,rep,name=datasets,proto3" json:"datasets,omitempty"`
	CreatedAt int64               `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PlacementRules) Reset() {
	*x = PlacementRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlacementRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlacementRules) ProtoMessage() {}

func (x *PlacementRules) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlacementRules.ProtoReflect.Descriptor instead.
func (*PlacementRules) Descriptor() ([]byte, []int) {
	return file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescGZIP(), []int{4}
}

func (x *PlacementRules) GetTenants() []*TenantPlacement {
	if x != nil {
		return x.Tenants
	}
	return nil
}

func (x *PlacementRules) GetDatasets() []*DatasetPlacement {
	if x != nil {
		return x.Datasets
	}
	return nil
}

func (x *PlacementRules) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type TenantPlacement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *TenantPlacement) Reset() {
	*x = TenantPlacement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantPlacement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantPlacement) ProtoMessage() {}

func (x *TenantPlacement) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantPlacement.ProtoReflect.Descriptor instead.
func (*TenantPlacement) Descriptor() ([]byte, []int) {
	return file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescGZIP(), []int{5}
}

func (x *TenantPlacement) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type DatasetPlacement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant            uint32        `protobuf:"varint,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Name              string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TenantShardLimit  uint64        `protobuf:"varint,3,opt,name=tenant_shard_limit,json=tenantShardLimit,proto3" json:"tenant_shard_limit,omitempty"`
	DatasetShardLimit uint64        `protobuf:"varint,4,opt,name=dataset_shard_limit,json=datasetShardLimit,proto3" json:"dataset_shard_limit,omitempty"`
	LoadBalancing     LoadBalancing `protobuf:"varint,5,opt,name=load_balancing,json=loadBalancing,proto3,enum=adaptive_placement.LoadBalancing" json:"load_balancing,omitempty"`
}

func (x *DatasetPlacement) Reset() {
	*x = DatasetPlacement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetPlacement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetPlacement) ProtoMessage() {}

func (x *DatasetPlacement) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetPlacement.ProtoReflect.Descriptor instead.
func (*DatasetPlacement) Descriptor() ([]byte, []int) {
	return file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescGZIP(), []int{6}
}

func (x *DatasetPlacement) GetTenant() uint32 {
	if x != nil {
		return x.Tenant
	}
	return 0
}

func (x *DatasetPlacement) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DatasetPlacement) GetTenantShardLimit() uint64 {
	if x != nil {
		return x.TenantShardLimit
	}
	return 0
}

func (x *DatasetPlacement) GetDatasetShardLimit() uint64 {
	if x != nil {
		return x.DatasetShardLimit
	}
	return 0
}

func (x *DatasetPlacement) GetLoadBalancing() LoadBalancing {
	if x != nil {
		return x.LoadBalancing
	}
	return LoadBalancing_LOAD_BALANCING_UNSPECIFIED
}

var File_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto protoreflect.FileDescriptor

var file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDesc = []byte{
	0x0a, 0x61, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xe3, 0x01, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x39, 0x0a,
	0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x68, 0x61, 0x72,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2a, 0x0a,
	0x0b, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x0c, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x05, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x74, 0x64, 0x5f, 0x64, 0x65, 0x76, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x64, 0x44, 0x65, 0x76, 0x22, 0x32, 0x0a,
	0x0a, 0x53, 0x68, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x22, 0xb0, 0x01, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x2e, 0x0a, 0x0f, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0xe6, 0x01, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f,
	0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x10, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x68, 0x61, 0x72, 0x64, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x11, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x64, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x48, 0x0a, 0x0e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x61, 0x64,
	0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x0d,
	0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x2a, 0x6f, 0x0a,
	0x0d, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x1e,
	0x0a, 0x1a, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x42, 0x41, 0x4c, 0x41, 0x4e, 0x43, 0x49, 0x4e, 0x47,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e,
	0x0a, 0x1a, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x42, 0x41, 0x4c, 0x41, 0x4e, 0x43, 0x49, 0x4e, 0x47,
	0x5f, 0x46, 0x49, 0x4e, 0x47, 0x45, 0x52, 0x50, 0x52, 0x49, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x1e,
	0x0a, 0x1a, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x42, 0x41, 0x4c, 0x41, 0x4e, 0x43, 0x49, 0x4e, 0x47,
	0x5f, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x52, 0x4f, 0x42, 0x49, 0x4e, 0x10, 0x02, 0x42, 0xff,
	0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x16, 0x41, 0x64, 0x61, 0x70, 0x74,
	0x69, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x69, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0xa2, 0x02,
	0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x11, 0x41, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x11, 0x41, 0x64, 0x61, 0x70, 0x74,
	0x69, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0xe2, 0x02, 0x1d, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescOnce sync.Once
	file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescData = file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDesc
)

func file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescGZIP() []byte {
	file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescOnce.Do(func() {
		file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescData = protoimpl.X.CompressGZIP(file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescData)
	})
	return file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDescData
}

var file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_goTypes = []any{
	(LoadBalancing)(0),        // 0: adaptive_placement.LoadBalancing
	(*DistributionStats)(nil), // 1: adaptive_placement.DistributionStats
	(*TenantStats)(nil),       // 2: adaptive_placement.TenantStats
	(*DatasetStats)(nil),      // 3: adaptive_placement.DatasetStats
	(*ShardStats)(nil),        // 4: adaptive_placement.ShardStats
	(*PlacementRules)(nil),    // 5: adaptive_placement.PlacementRules
	(*TenantPlacement)(nil),   // 6: adaptive_placement.TenantPlacement
	(*DatasetPlacement)(nil),  // 7: adaptive_placement.DatasetPlacement
}
var file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_depIdxs = []int32{
	2, // 0: adaptive_placement.DistributionStats.tenants:type_name -> adaptive_placement.TenantStats
	3, // 1: adaptive_placement.DistributionStats.datasets:type_name -> adaptive_placement.DatasetStats
	4, // 2: adaptive_placement.DistributionStats.shards:type_name -> adaptive_placement.ShardStats
	6, // 3: adaptive_placement.PlacementRules.tenants:type_name -> adaptive_placement.TenantPlacement
	7, // 4: adaptive_placement.PlacementRules.datasets:type_name -> adaptive_placement.DatasetPlacement
	0, // 5: adaptive_placement.DatasetPlacement.load_balancing:type_name -> adaptive_placement.LoadBalancing
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_init()
}
func file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_init() {
	if File_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DistributionStats); i {
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
		file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TenantStats); i {
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
		file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DatasetStats); i {
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
		file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ShardStats); i {
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
		file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PlacementRules); i {
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
		file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TenantPlacement); i {
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
		file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DatasetPlacement); i {
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
			RawDescriptor: file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_goTypes,
		DependencyIndexes: file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_depIdxs,
		EnumInfos:         file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_enumTypes,
		MessageInfos:      file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_msgTypes,
	}.Build()
	File_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto = out.File
	file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_rawDesc = nil
	file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_goTypes = nil
	file_experiment_distributor_placement_adaptive_placement_adaptive_placementpb_adaptive_placement_proto_depIdxs = nil
}
