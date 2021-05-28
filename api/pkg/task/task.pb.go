//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: api/pkg/task/pb/task.proto

package task

import (
	_ "github.com/infraboard/mcube/pb/http"
	page "github.com/infraboard/mcube/pb/page"
	_ "github.com/infraboard/protoc-gen-go-ext/extension/tag"
	pipeline "github.com/infraboard/workflow/api/pkg/pipeline"
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

// PIPELINE_STATUS 流水线状态
type PIPELINE_STATUS int32

const (
	// 状态未知
	PIPELINE_STATUS_UNKNOW PIPELINE_STATUS = 0
	// 任务等待被调度
	PIPELINE_STATUS_PENDDING PIPELINE_STATUS = 1
	// 正在执行
	PIPELINE_STATUS_RUNNING PIPELINE_STATUS = 2
	// 执行成功
	PIPELINE_STATUS_SUCCEEDED PIPELINE_STATUS = 3
	// 执行失败
	PIPELINE_STATUS_FAILED PIPELINE_STATUS = 4
)

// Enum value maps for PIPELINE_STATUS.
var (
	PIPELINE_STATUS_name = map[int32]string{
		0: "UNKNOW",
		1: "PENDDING",
		2: "RUNNING",
		3: "SUCCEEDED",
		4: "FAILED",
	}
	PIPELINE_STATUS_value = map[string]int32{
		"UNKNOW":    0,
		"PENDDING":  1,
		"RUNNING":   2,
		"SUCCEEDED": 3,
		"FAILED":    4,
	}
)

func (x PIPELINE_STATUS) Enum() *PIPELINE_STATUS {
	p := new(PIPELINE_STATUS)
	*p = x
	return p
}

func (x PIPELINE_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PIPELINE_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_api_pkg_task_pb_task_proto_enumTypes[0].Descriptor()
}

func (PIPELINE_STATUS) Type() protoreflect.EnumType {
	return &file_api_pkg_task_pb_task_proto_enumTypes[0]
}

func (x PIPELINE_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PIPELINE_STATUS.Descriptor instead.
func (PIPELINE_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_api_pkg_task_pb_task_proto_rawDescGZIP(), []int{0}
}

// Stage todo
type PipelineTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 唯一ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 需要被运行的Pipeline对象
	Pipeline *pipeline.Pipeline `protobuf:"bytes,2,opt,name=pipeline,proto3" json:"pipeline" bson:"pipeline"`
	// 创建时间
	CreateAt int64 `protobuf:"varint,3,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 开始时间
	StartAt int64 `protobuf:"varint,4,opt,name=start_at,json=startAt,proto3" json:"start_at" bson:"start_at"`
	// 结束时间
	EndAt int64 `protobuf:"varint,5,opt,name=end_at,json=endAt,proto3" json:"end_at" bson:"end_at"`
	// 状态更新的时间
	UpdateAt int64 `protobuf:"varint,6,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 任务状态
	Status PIPELINE_STATUS `protobuf:"varint,7,opt,name=status,proto3,enum=workflow.task.PIPELINE_STATUS" json:"status" bson:"status"`
	// 执行结果
	Message string `protobuf:"bytes,8,opt,name=message,proto3" json:"message" bson:"message"`
}

func (x *PipelineTask) Reset() {
	*x = PipelineTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pkg_task_pb_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineTask) ProtoMessage() {}

func (x *PipelineTask) ProtoReflect() protoreflect.Message {
	mi := &file_api_pkg_task_pb_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineTask.ProtoReflect.Descriptor instead.
func (*PipelineTask) Descriptor() ([]byte, []int) {
	return file_api_pkg_task_pb_task_proto_rawDescGZIP(), []int{0}
}

func (x *PipelineTask) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PipelineTask) GetPipeline() *pipeline.Pipeline {
	if x != nil {
		return x.Pipeline
	}
	return nil
}

func (x *PipelineTask) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *PipelineTask) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *PipelineTask) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *PipelineTask) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *PipelineTask) GetStatus() PIPELINE_STATUS {
	if x != nil {
		return x.Status
	}
	return PIPELINE_STATUS_UNKNOW
}

func (x *PipelineTask) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// QueryPipelineTaskRequest 查询Book请求
type QueryPipelineTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Name string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *QueryPipelineTaskRequest) Reset() {
	*x = QueryPipelineTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pkg_task_pb_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPipelineTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPipelineTaskRequest) ProtoMessage() {}

func (x *QueryPipelineTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pkg_task_pb_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPipelineTaskRequest.ProtoReflect.Descriptor instead.
func (*QueryPipelineTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_pkg_task_pb_task_proto_rawDescGZIP(), []int{1}
}

func (x *QueryPipelineTaskRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryPipelineTaskRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// PipelineTaskSet todo
type PipelineTaskSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64           `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	Items []*PipelineTask `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *PipelineTaskSet) Reset() {
	*x = PipelineTaskSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pkg_task_pb_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineTaskSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineTaskSet) ProtoMessage() {}

func (x *PipelineTaskSet) ProtoReflect() protoreflect.Message {
	mi := &file_api_pkg_task_pb_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineTaskSet.ProtoReflect.Descriptor instead.
func (*PipelineTaskSet) Descriptor() ([]byte, []int) {
	return file_api_pkg_task_pb_task_proto_rawDescGZIP(), []int{2}
}

func (x *PipelineTaskSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PipelineTaskSet) GetItems() []*PipelineTask {
	if x != nil {
		return x.Items
	}
	return nil
}

// RunPipelineRequest todo
type RunPipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 唯一ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
}

func (x *RunPipelineRequest) Reset() {
	*x = RunPipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pkg_task_pb_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunPipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunPipelineRequest) ProtoMessage() {}

func (x *RunPipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pkg_task_pb_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunPipelineRequest.ProtoReflect.Descriptor instead.
func (*RunPipelineRequest) Descriptor() ([]byte, []int) {
	return file_api_pkg_task_pb_task_proto_rawDescGZIP(), []int{3}
}

func (x *RunPipelineRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_api_pkg_task_pb_task_proto protoreflect.FileDescriptor

var file_api_pkg_task_pb_task_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70,
	0x62, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x3f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2d, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74,
	0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61,
	0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbc, 0x04, 0x0a, 0x0c, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a,
	0xc2, 0xde, 0x1f, 0x16, 0x0a, 0x14, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x5e,
	0x0a, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x25, 0xc2,
	0xde, 0x1f, 0x21, 0x0a, 0x1f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x22, 0x52, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x44,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x25, 0xc2, 0xde, 0x1f, 0x21, 0x0a, 0x1f, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x07, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x21, 0xc2, 0xde, 0x1f, 0x1d, 0x0a, 0x1b, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74,
	0x12, 0x44, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x08, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x59, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x42, 0x21, 0xc2, 0xde, 0x1f, 0x1d, 0x0a, 0x1b, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x3d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x23, 0xc2, 0xde, 0x1f, 0x1f, 0x0a, 0x1d, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x55, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x0f, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b,
	0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x52, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x1f,
	0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x40, 0x0a, 0x12, 0x52, 0x75, 0x6e, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xc2, 0xde, 0x1f, 0x16, 0x0a, 0x14,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x53, 0x0a, 0x0f, 0x50, 0x49, 0x50, 0x45,
	0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x0a, 0x0a, 0x06, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x45, 0x4e, 0x44, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x32, 0xbb, 0x02,
	0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x0b, 0x52, 0x75,
	0x6e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x21, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x75, 0x6e, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x42, 0xa2, 0xa3, 0x8c, 0x4d, 0x3d,
	0x1a, 0x10, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x2f, 0x22, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x2a, 0x0d, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x30, 0x01, 0x38, 0x01, 0x42, 0x10, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x9b, 0x01,
	0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x27, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x74, 0x22, 0x3d, 0xa2, 0xa3,
	0x8c, 0x4d, 0x38, 0x1a, 0x10, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2f, 0x22, 0x03, 0x47, 0x45, 0x54, 0x2a, 0x0d, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x30, 0x01, 0x42, 0x0e, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_pkg_task_pb_task_proto_rawDescOnce sync.Once
	file_api_pkg_task_pb_task_proto_rawDescData = file_api_pkg_task_pb_task_proto_rawDesc
)

func file_api_pkg_task_pb_task_proto_rawDescGZIP() []byte {
	file_api_pkg_task_pb_task_proto_rawDescOnce.Do(func() {
		file_api_pkg_task_pb_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_pkg_task_pb_task_proto_rawDescData)
	})
	return file_api_pkg_task_pb_task_proto_rawDescData
}

var file_api_pkg_task_pb_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_pkg_task_pb_task_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_pkg_task_pb_task_proto_goTypes = []interface{}{
	(PIPELINE_STATUS)(0),             // 0: workflow.task.PIPELINE_STATUS
	(*PipelineTask)(nil),             // 1: workflow.task.PipelineTask
	(*QueryPipelineTaskRequest)(nil), // 2: workflow.task.QueryPipelineTaskRequest
	(*PipelineTaskSet)(nil),          // 3: workflow.task.PipelineTaskSet
	(*RunPipelineRequest)(nil),       // 4: workflow.task.RunPipelineRequest
	(*pipeline.Pipeline)(nil),        // 5: workflow.pipeline.Pipeline
	(*page.PageRequest)(nil),         // 6: page.PageRequest
}
var file_api_pkg_task_pb_task_proto_depIdxs = []int32{
	5, // 0: workflow.task.PipelineTask.pipeline:type_name -> workflow.pipeline.Pipeline
	0, // 1: workflow.task.PipelineTask.status:type_name -> workflow.task.PIPELINE_STATUS
	6, // 2: workflow.task.QueryPipelineTaskRequest.page:type_name -> page.PageRequest
	1, // 3: workflow.task.PipelineTaskSet.items:type_name -> workflow.task.PipelineTask
	4, // 4: workflow.task.Service.RunPipeline:input_type -> workflow.task.RunPipelineRequest
	2, // 5: workflow.task.Service.QueryPipelineTask:input_type -> workflow.task.QueryPipelineTaskRequest
	1, // 6: workflow.task.Service.RunPipeline:output_type -> workflow.task.PipelineTask
	3, // 7: workflow.task.Service.QueryPipelineTask:output_type -> workflow.task.PipelineTaskSet
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_pkg_task_pb_task_proto_init() }
func file_api_pkg_task_pb_task_proto_init() {
	if File_api_pkg_task_pb_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_pkg_task_pb_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineTask); i {
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
		file_api_pkg_task_pb_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPipelineTaskRequest); i {
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
		file_api_pkg_task_pb_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineTaskSet); i {
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
		file_api_pkg_task_pb_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunPipelineRequest); i {
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
			RawDescriptor: file_api_pkg_task_pb_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_pkg_task_pb_task_proto_goTypes,
		DependencyIndexes: file_api_pkg_task_pb_task_proto_depIdxs,
		EnumInfos:         file_api_pkg_task_pb_task_proto_enumTypes,
		MessageInfos:      file_api_pkg_task_pb_task_proto_msgTypes,
	}.Build()
	File_api_pkg_task_pb_task_proto = out.File
	file_api_pkg_task_pb_task_proto_rawDesc = nil
	file_api_pkg_task_pb_task_proto_goTypes = nil
	file_api_pkg_task_pb_task_proto_depIdxs = nil
}
