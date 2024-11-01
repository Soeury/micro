// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.2
// source: learn1/test7/book/wrap.proto

package book

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=Author,proto3" json:"Author,omitempty"`
	// int64 price = 3; // 这里无法判断  -1.输入的值是零值  -2.还是未输入值
	Price     *wrapperspb.Int64Value  `protobuf:"bytes,3,opt,name=Price,proto3" json:"Price,omitempty"`
	SalePrice *wrapperspb.DoubleValue `protobuf:"bytes,4,opt,name=SalePrice,proto3" json:"SalePrice,omitempty"`
	Content   *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	Another   *Book_Info              `protobuf:"bytes,6,opt,name=another,proto3" json:"another,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_learn1_test7_book_wrap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_learn1_test7_book_wrap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_learn1_test7_book_wrap_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPrice() *wrapperspb.Int64Value {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *Book) GetSalePrice() *wrapperspb.DoubleValue {
	if x != nil {
		return x.SalePrice
	}
	return nil
}

func (x *Book) GetContent() *wrapperspb.StringValue {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Book) GetAnother() *Book_Info {
	if x != nil {
		return x.Another
	}
	return nil
}

type Book2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string                 `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Author string                 `protobuf:"bytes,2,opt,name=Author,proto3" json:"Author,omitempty"`
	Price  *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *Book2) Reset() {
	*x = Book2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_learn1_test7_book_wrap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book2) ProtoMessage() {}

func (x *Book2) ProtoReflect() protoreflect.Message {
	mi := &file_learn1_test7_book_wrap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book2.ProtoReflect.Descriptor instead.
func (*Book2) Descriptor() ([]byte, []int) {
	return file_learn1_test7_book_wrap_proto_rawDescGZIP(), []int{1}
}

func (x *Book2) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book2) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book2) GetPrice() *wrapperspb.Int64Value {
	if x != nil {
		return x.Price
	}
	return nil
}

type UpdataBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 操作人
	Op string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	// 要更新的书籍
	Book *Book `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
	// 添加这一条来记录需要更新的字段
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdataBook) Reset() {
	*x = UpdataBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_learn1_test7_book_wrap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdataBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdataBook) ProtoMessage() {}

func (x *UpdataBook) ProtoReflect() protoreflect.Message {
	mi := &file_learn1_test7_book_wrap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdataBook.ProtoReflect.Descriptor instead.
func (*UpdataBook) Descriptor() ([]byte, []int) {
	return file_learn1_test7_book_wrap_proto_rawDescGZIP(), []int{2}
}

func (x *UpdataBook) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *UpdataBook) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

func (x *UpdataBook) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// 这里给 book 嵌套一个消息
type Book_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teacher string `protobuf:"bytes,1,opt,name=teacher,proto3" json:"teacher,omitempty"`
	Art     string `protobuf:"bytes,2,opt,name=art,proto3" json:"art,omitempty"`
}

func (x *Book_Info) Reset() {
	*x = Book_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_learn1_test7_book_wrap_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book_Info) ProtoMessage() {}

func (x *Book_Info) ProtoReflect() protoreflect.Message {
	mi := &file_learn1_test7_book_wrap_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book_Info.ProtoReflect.Descriptor instead.
func (*Book_Info) Descriptor() ([]byte, []int) {
	return file_learn1_test7_book_wrap_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Book_Info) GetTeacher() string {
	if x != nil {
		return x.Teacher
	}
	return ""
}

func (x *Book_Info) GetArt() string {
	if x != nil {
		return x.Art
	}
	return ""
}

var File_learn1_test7_book_wrap_proto protoreflect.FileDescriptor

var file_learn1_test7_book_wrap_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x37, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x02, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x31, 0x0a,
	0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x3a, 0x0a, 0x09, 0x53, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x09, 0x53, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x1a,
	0x32, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x72, 0x74, 0x22, 0x68, 0x0a, 0x05, 0x42, 0x6f, 0x6f, 0x6b, 0x32, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x31, 0x0a, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x79, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x61, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x1e, 0x0a, 0x04, 0x62,
	0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x13, 0x5a, 0x11, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x37, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_learn1_test7_book_wrap_proto_rawDescOnce sync.Once
	file_learn1_test7_book_wrap_proto_rawDescData = file_learn1_test7_book_wrap_proto_rawDesc
)

func file_learn1_test7_book_wrap_proto_rawDescGZIP() []byte {
	file_learn1_test7_book_wrap_proto_rawDescOnce.Do(func() {
		file_learn1_test7_book_wrap_proto_rawDescData = protoimpl.X.CompressGZIP(file_learn1_test7_book_wrap_proto_rawDescData)
	})
	return file_learn1_test7_book_wrap_proto_rawDescData
}

var file_learn1_test7_book_wrap_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_learn1_test7_book_wrap_proto_goTypes = []interface{}{
	(*Book)(nil),                   // 0: book.Book
	(*Book2)(nil),                  // 1: book.Book2
	(*UpdataBook)(nil),             // 2: book.UpdataBook
	(*Book_Info)(nil),              // 3: book.Book.Info
	(*wrapperspb.Int64Value)(nil),  // 4: google.protobuf.Int64Value
	(*wrapperspb.DoubleValue)(nil), // 5: google.protobuf.DoubleValue
	(*wrapperspb.StringValue)(nil), // 6: google.protobuf.StringValue
	(*fieldmaskpb.FieldMask)(nil),  // 7: google.protobuf.FieldMask
}
var file_learn1_test7_book_wrap_proto_depIdxs = []int32{
	4, // 0: book.Book.Price:type_name -> google.protobuf.Int64Value
	5, // 1: book.Book.SalePrice:type_name -> google.protobuf.DoubleValue
	6, // 2: book.Book.Content:type_name -> google.protobuf.StringValue
	3, // 3: book.Book.another:type_name -> book.Book.Info
	4, // 4: book.Book2.Price:type_name -> google.protobuf.Int64Value
	0, // 5: book.UpdataBook.book:type_name -> book.Book
	7, // 6: book.UpdataBook.update_mask:type_name -> google.protobuf.FieldMask
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_learn1_test7_book_wrap_proto_init() }
func file_learn1_test7_book_wrap_proto_init() {
	if File_learn1_test7_book_wrap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_learn1_test7_book_wrap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_learn1_test7_book_wrap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book2); i {
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
		file_learn1_test7_book_wrap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdataBook); i {
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
		file_learn1_test7_book_wrap_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book_Info); i {
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
			RawDescriptor: file_learn1_test7_book_wrap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_learn1_test7_book_wrap_proto_goTypes,
		DependencyIndexes: file_learn1_test7_book_wrap_proto_depIdxs,
		MessageInfos:      file_learn1_test7_book_wrap_proto_msgTypes,
	}.Build()
	File_learn1_test7_book_wrap_proto = out.File
	file_learn1_test7_book_wrap_proto_rawDesc = nil
	file_learn1_test7_book_wrap_proto_goTypes = nil
	file_learn1_test7_book_wrap_proto_depIdxs = nil
}
