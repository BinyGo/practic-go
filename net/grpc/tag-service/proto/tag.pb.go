// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: proto/tag.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetTagListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State uint32 `protobuf:"varint,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *GetTagListRequest) Reset() {
	*x = GetTagListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagListRequest) ProtoMessage() {}

func (x *GetTagListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagListRequest.ProtoReflect.Descriptor instead.
func (*GetTagListRequest) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{0}
}

func (x *GetTagListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetTagListRequest) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State uint32 `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{1}
}

func (x *Tag) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tag) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

type Pager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	TotalRows int64 `protobuf:"varint,3,opt,name=total_rows,json=totalRows,proto3" json:"total_rows,omitempty"`
}

func (x *Pager) Reset() {
	*x = Pager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pager) ProtoMessage() {}

func (x *Pager) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pager.ProtoReflect.Descriptor instead.
func (*Pager) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{2}
}

func (x *Pager) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pager) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Pager) GetTotalRows() int64 {
	if x != nil {
		return x.TotalRows
	}
	return 0
}

type GetTagListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*Tag `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Pager *Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
}

func (x *GetTagListReply) Reset() {
	*x = GetTagListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagListReply) ProtoMessage() {}

func (x *GetTagListReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagListReply.ProtoReflect.Descriptor instead.
func (*GetTagListReply) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{3}
}

func (x *GetTagListReply) GetList() []*Tag {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *GetTagListReply) GetPager() *Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Detail  *anypb.Any `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_proto_tag_proto_rawDescGZIP(), []int{4}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetDetail() *anypb.Any {
	if x != nil {
		return x.Detail
	}
	return nil
}

var File_proto_tag_proto protoreflect.FileDescriptor

var file_proto_tag_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x74, 0x61, 0x67, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x22, 0x3f, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x57, 0x0a, 0x05, 0x50, 0x61, 0x67, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x6f, 0x77, 0x73, 0x22, 0x51, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x74, 0x61,
	0x67, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x61, 0x67,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x22, 0x63, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x32, 0x4a, 0x0a, 0x0a, 0x54, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16,
	0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_tag_proto_rawDescOnce sync.Once
	file_proto_tag_proto_rawDescData = file_proto_tag_proto_rawDesc
)

func file_proto_tag_proto_rawDescGZIP() []byte {
	file_proto_tag_proto_rawDescOnce.Do(func() {
		file_proto_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_tag_proto_rawDescData)
	})
	return file_proto_tag_proto_rawDescData
}

var file_proto_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_tag_proto_goTypes = []interface{}{
	(*GetTagListRequest)(nil), // 0: tag.GetTagListRequest
	(*Tag)(nil),               // 1: tag.Tag
	(*Pager)(nil),             // 2: tag.Pager
	(*GetTagListReply)(nil),   // 3: tag.GetTagListReply
	(*Error)(nil),             // 4: tag.Error
	(*anypb.Any)(nil),         // 5: google.protobuf.Any
}
var file_proto_tag_proto_depIdxs = []int32{
	1, // 0: tag.GetTagListReply.list:type_name -> tag.Tag
	2, // 1: tag.GetTagListReply.pager:type_name -> tag.Pager
	5, // 2: tag.Error.detail:type_name -> google.protobuf.Any
	0, // 3: tag.TagService.GetTagList:input_type -> tag.GetTagListRequest
	3, // 4: tag.TagService.GetTagList:output_type -> tag.GetTagListReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_tag_proto_init() }
func file_proto_tag_proto_init() {
	if File_proto_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagListRequest); i {
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
		file_proto_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_proto_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pager); i {
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
		file_proto_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagListReply); i {
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
		file_proto_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_proto_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_tag_proto_goTypes,
		DependencyIndexes: file_proto_tag_proto_depIdxs,
		MessageInfos:      file_proto_tag_proto_msgTypes,
	}.Build()
	File_proto_tag_proto = out.File
	file_proto_tag_proto_rawDesc = nil
	file_proto_tag_proto_goTypes = nil
	file_proto_tag_proto_depIdxs = nil
}
