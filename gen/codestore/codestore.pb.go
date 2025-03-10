// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.3
// source: codestore/codestore.proto

package codestore

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

type PullRepoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PullRepoRequest) Reset() {
	*x = PullRepoRequest{}
	mi := &file_codestore_codestore_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PullRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullRepoRequest) ProtoMessage() {}

func (x *PullRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_codestore_codestore_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullRepoRequest.ProtoReflect.Descriptor instead.
func (*PullRepoRequest) Descriptor() ([]byte, []int) {
	return file_codestore_codestore_proto_rawDescGZIP(), []int{0}
}

func (x *PullRepoRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type PullRepoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reponame      string                 `protobuf:"bytes,1,opt,name=reponame,proto3" json:"reponame,omitempty"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Success       string                 `protobuf:"bytes,4,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PullRepoResponse) Reset() {
	*x = PullRepoResponse{}
	mi := &file_codestore_codestore_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PullRepoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullRepoResponse) ProtoMessage() {}

func (x *PullRepoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_codestore_codestore_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullRepoResponse.ProtoReflect.Descriptor instead.
func (*PullRepoResponse) Descriptor() ([]byte, []int) {
	return file_codestore_codestore_proto_rawDescGZIP(), []int{1}
}

func (x *PullRepoResponse) GetReponame() string {
	if x != nil {
		return x.Reponame
	}
	return ""
}

func (x *PullRepoResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PullRepoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PullRepoResponse) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

type FileChunk struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Filename      string                 `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	mi := &file_codestore_codestore_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_codestore_codestore_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_codestore_codestore_proto_rawDescGZIP(), []int{2}
}

func (x *FileChunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FileChunk) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type UploadCodeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Success       bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadCodeResponse) Reset() {
	*x = UploadCodeResponse{}
	mi := &file_codestore_codestore_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadCodeResponse) ProtoMessage() {}

func (x *UploadCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_codestore_codestore_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadCodeResponse.ProtoReflect.Descriptor instead.
func (*UploadCodeResponse) Descriptor() ([]byte, []int) {
	return file_codestore_codestore_proto_rawDescGZIP(), []int{3}
}

func (x *UploadCodeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UploadCodeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_codestore_codestore_proto protoreflect.FileDescriptor

var file_codestore_codestore_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x23, 0x0a, 0x0f, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65,
	0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x74, 0x0a, 0x10, 0x50,
	0x75, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x3b, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48,
	0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x95, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x64,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65,
	0x70, 0x6f, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50,
	0x75, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x52,
	0x65, 0x70, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a,
	0x1d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01,
	0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_codestore_codestore_proto_rawDescOnce sync.Once
	file_codestore_codestore_proto_rawDescData = file_codestore_codestore_proto_rawDesc
)

func file_codestore_codestore_proto_rawDescGZIP() []byte {
	file_codestore_codestore_proto_rawDescOnce.Do(func() {
		file_codestore_codestore_proto_rawDescData = protoimpl.X.CompressGZIP(file_codestore_codestore_proto_rawDescData)
	})
	return file_codestore_codestore_proto_rawDescData
}

var file_codestore_codestore_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_codestore_codestore_proto_goTypes = []any{
	(*PullRepoRequest)(nil),    // 0: codestore.PullRepoRequest
	(*PullRepoResponse)(nil),   // 1: codestore.PullRepoResponse
	(*FileChunk)(nil),          // 2: codestore.FileChunk
	(*UploadCodeResponse)(nil), // 3: codestore.UploadCodeResponse
}
var file_codestore_codestore_proto_depIdxs = []int32{
	0, // 0: codestore.CodeStore.PullRepo:input_type -> codestore.PullRepoRequest
	2, // 1: codestore.CodeStore.UploadCode:input_type -> codestore.FileChunk
	1, // 2: codestore.CodeStore.PullRepo:output_type -> codestore.PullRepoResponse
	3, // 3: codestore.CodeStore.UploadCode:output_type -> codestore.UploadCodeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_codestore_codestore_proto_init() }
func file_codestore_codestore_proto_init() {
	if File_codestore_codestore_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_codestore_codestore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_codestore_codestore_proto_goTypes,
		DependencyIndexes: file_codestore_codestore_proto_depIdxs,
		MessageInfos:      file_codestore_codestore_proto_msgTypes,
	}.Build()
	File_codestore_codestore_proto = out.File
	file_codestore_codestore_proto_rawDesc = nil
	file_codestore_codestore_proto_goTypes = nil
	file_codestore_codestore_proto_depIdxs = nil
}
