// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.2
// source: learn2/test6_summary/proto/sum/sum.proto

package sum

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

type AddPRCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int64 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int64 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *AddPRCRequest) Reset() {
	*x = AddPRCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_learn2_test6_summary_proto_sum_sum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPRCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPRCRequest) ProtoMessage() {}

func (x *AddPRCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_learn2_test6_summary_proto_sum_sum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPRCRequest.ProtoReflect.Descriptor instead.
func (*AddPRCRequest) Descriptor() ([]byte, []int) {
	return file_learn2_test6_summary_proto_sum_sum_proto_rawDescGZIP(), []int{0}
}

func (x *AddPRCRequest) GetNum1() int64 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *AddPRCRequest) GetNum2() int64 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type AddRPCResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret int64  `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *AddRPCResponse) Reset() {
	*x = AddRPCResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_learn2_test6_summary_proto_sum_sum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRPCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRPCResponse) ProtoMessage() {}

func (x *AddRPCResponse) ProtoReflect() protoreflect.Message {
	mi := &file_learn2_test6_summary_proto_sum_sum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRPCResponse.ProtoReflect.Descriptor instead.
func (*AddRPCResponse) Descriptor() ([]byte, []int) {
	return file_learn2_test6_summary_proto_sum_sum_proto_rawDescGZIP(), []int{1}
}

func (x *AddRPCResponse) GetRet() int64 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *AddRPCResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type AppendRPCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str1 string `protobuf:"bytes,1,opt,name=str1,proto3" json:"str1,omitempty"`
	Str2 string `protobuf:"bytes,2,opt,name=str2,proto3" json:"str2,omitempty"`
}

func (x *AppendRPCRequest) Reset() {
	*x = AppendRPCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_learn2_test6_summary_proto_sum_sum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendRPCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendRPCRequest) ProtoMessage() {}

func (x *AppendRPCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_learn2_test6_summary_proto_sum_sum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendRPCRequest.ProtoReflect.Descriptor instead.
func (*AppendRPCRequest) Descriptor() ([]byte, []int) {
	return file_learn2_test6_summary_proto_sum_sum_proto_rawDescGZIP(), []int{2}
}

func (x *AppendRPCRequest) GetStr1() string {
	if x != nil {
		return x.Str1
	}
	return ""
}

func (x *AppendRPCRequest) GetStr2() string {
	if x != nil {
		return x.Str2
	}
	return ""
}

type AppendRPCResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret string `protobuf:"bytes,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *AppendRPCResponse) Reset() {
	*x = AppendRPCResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_learn2_test6_summary_proto_sum_sum_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendRPCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendRPCResponse) ProtoMessage() {}

func (x *AppendRPCResponse) ProtoReflect() protoreflect.Message {
	mi := &file_learn2_test6_summary_proto_sum_sum_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendRPCResponse.ProtoReflect.Descriptor instead.
func (*AppendRPCResponse) Descriptor() ([]byte, []int) {
	return file_learn2_test6_summary_proto_sum_sum_proto_rawDescGZIP(), []int{3}
}

func (x *AppendRPCResponse) GetRet() string {
	if x != nil {
		return x.Ret
	}
	return ""
}

func (x *AppendRPCResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_learn2_test6_summary_proto_sum_sum_proto protoreflect.FileDescriptor

var file_learn2_test6_summary_proto_sum_sum_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x32, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x36, 0x5f, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x6d,
	0x2f, 0x73, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x75, 0x6d, 0x22,
	0x37, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x50, 0x52, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x6e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x22, 0x34, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x52,
	0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x3a,
	0x0a, 0x10, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x72, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x74, 0x72, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x72, 0x32, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x74, 0x72, 0x32, 0x22, 0x37, 0x0a, 0x11, 0x41, 0x70,
	0x70, 0x65, 0x6e, 0x64, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x65, 0x72, 0x72, 0x32, 0x77, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x12,
	0x30, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x12, 0x2e, 0x73, 0x75, 0x6d, 0x2e, 0x41, 0x64, 0x64,
	0x50, 0x52, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x75, 0x6d,
	0x2e, 0x41, 0x64, 0x64, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x39, 0x0a, 0x06, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x15, 0x2e, 0x73, 0x75,
	0x6d, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x75, 0x6d, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x52,
	0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x32, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x36, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_learn2_test6_summary_proto_sum_sum_proto_rawDescOnce sync.Once
	file_learn2_test6_summary_proto_sum_sum_proto_rawDescData = file_learn2_test6_summary_proto_sum_sum_proto_rawDesc
)

func file_learn2_test6_summary_proto_sum_sum_proto_rawDescGZIP() []byte {
	file_learn2_test6_summary_proto_sum_sum_proto_rawDescOnce.Do(func() {
		file_learn2_test6_summary_proto_sum_sum_proto_rawDescData = protoimpl.X.CompressGZIP(file_learn2_test6_summary_proto_sum_sum_proto_rawDescData)
	})
	return file_learn2_test6_summary_proto_sum_sum_proto_rawDescData
}

var file_learn2_test6_summary_proto_sum_sum_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_learn2_test6_summary_proto_sum_sum_proto_goTypes = []interface{}{
	(*AddPRCRequest)(nil),     // 0: sum.AddPRCRequest
	(*AddRPCResponse)(nil),    // 1: sum.AddRPCResponse
	(*AppendRPCRequest)(nil),  // 2: sum.AppendRPCRequest
	(*AppendRPCResponse)(nil), // 3: sum.AppendRPCResponse
}
var file_learn2_test6_summary_proto_sum_sum_proto_depIdxs = []int32{
	0, // 0: sum.Computer.Add:input_type -> sum.AddPRCRequest
	2, // 1: sum.Computer.Append:input_type -> sum.AppendRPCRequest
	1, // 2: sum.Computer.Add:output_type -> sum.AddRPCResponse
	3, // 3: sum.Computer.Append:output_type -> sum.AppendRPCResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_learn2_test6_summary_proto_sum_sum_proto_init() }
func file_learn2_test6_summary_proto_sum_sum_proto_init() {
	if File_learn2_test6_summary_proto_sum_sum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_learn2_test6_summary_proto_sum_sum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPRCRequest); i {
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
		file_learn2_test6_summary_proto_sum_sum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRPCResponse); i {
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
		file_learn2_test6_summary_proto_sum_sum_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendRPCRequest); i {
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
		file_learn2_test6_summary_proto_sum_sum_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendRPCResponse); i {
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
			RawDescriptor: file_learn2_test6_summary_proto_sum_sum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_learn2_test6_summary_proto_sum_sum_proto_goTypes,
		DependencyIndexes: file_learn2_test6_summary_proto_sum_sum_proto_depIdxs,
		MessageInfos:      file_learn2_test6_summary_proto_sum_sum_proto_msgTypes,
	}.Build()
	File_learn2_test6_summary_proto_sum_sum_proto = out.File
	file_learn2_test6_summary_proto_sum_sum_proto_rawDesc = nil
	file_learn2_test6_summary_proto_sum_sum_proto_goTypes = nil
	file_learn2_test6_summary_proto_sum_sum_proto_depIdxs = nil
}