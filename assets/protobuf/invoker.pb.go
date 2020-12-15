// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: invoker.proto

package invokerMessage

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type JobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FunctionName string   `protobuf:"bytes,2,opt,name=functionName,proto3" json:"functionName,omitempty"`
	Variables    []string `protobuf:"bytes,3,rep,name=variables,proto3" json:"variables,omitempty"`
	Files        []byte   `protobuf:"bytes,4,opt,name=files,proto3" json:"files,omitempty"`
}

func (x *JobRequest) Reset() {
	*x = JobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobRequest) ProtoMessage() {}

func (x *JobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invoker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobRequest.ProtoReflect.Descriptor instead.
func (*JobRequest) Descriptor() ([]byte, []int) {
	return file_invoker_proto_rawDescGZIP(), []int{0}
}

func (x *JobRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *JobRequest) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *JobRequest) GetVariables() []string {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *JobRequest) GetFiles() []byte {
	if x != nil {
		return x.Files
	}
	return nil
}

type JobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FunctionName string   `protobuf:"bytes,2,opt,name=functionName,proto3" json:"functionName,omitempty"`
	Results      []string `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty"`
	Files        []byte   `protobuf:"bytes,4,opt,name=files,proto3" json:"files,omitempty"`
	CacheHit     bool     `protobuf:"varint,5,opt,name=cacheHit,proto3" json:"cacheHit,omitempty"`
}

func (x *JobResponse) Reset() {
	*x = JobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobResponse) ProtoMessage() {}

func (x *JobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invoker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobResponse.ProtoReflect.Descriptor instead.
func (*JobResponse) Descriptor() ([]byte, []int) {
	return file_invoker_proto_rawDescGZIP(), []int{1}
}

func (x *JobResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *JobResponse) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *JobResponse) GetResults() []string {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *JobResponse) GetFiles() []byte {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *JobResponse) GetCacheHit() bool {
	if x != nil {
		return x.CacheHit
	}
	return false
}

var File_invoker_proto protoreflect.FileDescriptor

var file_invoker_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x74, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x48, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x48, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_invoker_proto_rawDescOnce sync.Once
	file_invoker_proto_rawDescData = file_invoker_proto_rawDesc
)

func file_invoker_proto_rawDescGZIP() []byte {
	file_invoker_proto_rawDescOnce.Do(func() {
		file_invoker_proto_rawDescData = protoimpl.X.CompressGZIP(file_invoker_proto_rawDescData)
	})
	return file_invoker_proto_rawDescData
}

var file_invoker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_invoker_proto_goTypes = []interface{}{
	(*JobRequest)(nil),  // 0: invokerMessage.JobRequest
	(*JobResponse)(nil), // 1: invokerMessage.JobResponse
}
var file_invoker_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_invoker_proto_init() }
func file_invoker_proto_init() {
	if File_invoker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_invoker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobRequest); i {
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
		file_invoker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobResponse); i {
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
			RawDescriptor: file_invoker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_invoker_proto_goTypes,
		DependencyIndexes: file_invoker_proto_depIdxs,
		MessageInfos:      file_invoker_proto_msgTypes,
	}.Build()
	File_invoker_proto = out.File
	file_invoker_proto_rawDesc = nil
	file_invoker_proto_goTypes = nil
	file_invoker_proto_depIdxs = nil
}