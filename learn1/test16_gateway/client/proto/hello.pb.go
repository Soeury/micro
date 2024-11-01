// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.2
// source: client/proto/hello.proto

package proto

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

type LoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *LoveRequest) Reset() {
	*x = LoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoveRequest) ProtoMessage() {}

func (x *LoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoveRequest.ProtoReflect.Descriptor instead.
func (*LoveRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_hello_proto_rawDescGZIP(), []int{0}
}

func (x *LoveRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=Reply,proto3" json:"Reply,omitempty"`
}

func (x *LoveResponse) Reset() {
	*x = LoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoveResponse) ProtoMessage() {}

func (x *LoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoveResponse.ProtoReflect.Descriptor instead.
func (*LoveResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_hello_proto_rawDescGZIP(), []int{1}
}

func (x *LoveResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_client_proto_hello_proto protoreflect.FileDescriptor

var file_client_proto_hello_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x4c, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x0c, 0x4c, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x3c, 0x0a, 0x04, 0x4c, 0x6f,
	0x76, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x53, 0x61, 0x79, 0x4c, 0x6f, 0x76, 0x65, 0x12, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x31, 0x36, 0x5f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_hello_proto_rawDescOnce sync.Once
	file_client_proto_hello_proto_rawDescData = file_client_proto_hello_proto_rawDesc
)

func file_client_proto_hello_proto_rawDescGZIP() []byte {
	file_client_proto_hello_proto_rawDescOnce.Do(func() {
		file_client_proto_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_hello_proto_rawDescData)
	})
	return file_client_proto_hello_proto_rawDescData
}

var file_client_proto_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_proto_hello_proto_goTypes = []interface{}{
	(*LoveRequest)(nil),  // 0: proto.LoveRequest
	(*LoveResponse)(nil), // 1: proto.LoveResponse
}
var file_client_proto_hello_proto_depIdxs = []int32{
	0, // 0: proto.Love.SayLove:input_type -> proto.LoveRequest
	1, // 1: proto.Love.SayLove:output_type -> proto.LoveResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_proto_hello_proto_init() }
func file_client_proto_hello_proto_init() {
	if File_client_proto_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_proto_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoveRequest); i {
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
		file_client_proto_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoveResponse); i {
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
			RawDescriptor: file_client_proto_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_proto_hello_proto_goTypes,
		DependencyIndexes: file_client_proto_hello_proto_depIdxs,
		MessageInfos:      file_client_proto_hello_proto_msgTypes,
	}.Build()
	File_client_proto_hello_proto = out.File
	file_client_proto_hello_proto_rawDesc = nil
	file_client_proto_hello_proto_goTypes = nil
	file_client_proto_hello_proto_depIdxs = nil
}
