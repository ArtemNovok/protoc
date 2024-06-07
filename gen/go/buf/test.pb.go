// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: buf/test.proto

package absv1

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

type AbsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *AbsRequest) Reset() {
	*x = AbsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbsRequest) ProtoMessage() {}

func (x *AbsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbsRequest.ProtoReflect.Descriptor instead.
func (*AbsRequest) Descriptor() ([]byte, []int) {
	return file_buf_test_proto_rawDescGZIP(), []int{0}
}

func (x *AbsRequest) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type AbsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num     int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	IsError bool  `protobuf:"varint,2,opt,name=is_error,json=isError,proto3" json:"is_error,omitempty"`
}

func (x *AbsResponse) Reset() {
	*x = AbsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbsResponse) ProtoMessage() {}

func (x *AbsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbsResponse.ProtoReflect.Descriptor instead.
func (*AbsResponse) Descriptor() ([]byte, []int) {
	return file_buf_test_proto_rawDescGZIP(), []int{1}
}

func (x *AbsResponse) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *AbsResponse) GetIsError() bool {
	if x != nil {
		return x.IsError
	}
	return false
}

var File_buf_test_proto protoreflect.FileDescriptor

var file_buf_test_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x62, 0x6c, 0x75, 0x70, 0x22, 0x1e, 0x0a, 0x0a, 0x41, 0x62, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x3a, 0x0a, 0x0b, 0x41, 0x62, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0x36, 0x0a, 0x03, 0x41, 0x62, 0x73, 0x12, 0x2f, 0x0a, 0x08, 0x41, 0x62, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x62, 0x6c, 0x75, 0x70, 0x2e, 0x41, 0x62, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x62, 0x6c, 0x75, 0x70, 0x2e, 0x41,
	0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x61, 0x72,
	0x74, 0x65, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x3b, 0x61, 0x62, 0x73, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_test_proto_rawDescOnce sync.Once
	file_buf_test_proto_rawDescData = file_buf_test_proto_rawDesc
)

func file_buf_test_proto_rawDescGZIP() []byte {
	file_buf_test_proto_rawDescOnce.Do(func() {
		file_buf_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_test_proto_rawDescData)
	})
	return file_buf_test_proto_rawDescData
}

var file_buf_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_buf_test_proto_goTypes = []interface{}{
	(*AbsRequest)(nil),  // 0: blup.AbsRequest
	(*AbsResponse)(nil), // 1: blup.AbsResponse
}
var file_buf_test_proto_depIdxs = []int32{
	0, // 0: blup.Abs.Absolute:input_type -> blup.AbsRequest
	1, // 1: blup.Abs.Absolute:output_type -> blup.AbsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_buf_test_proto_init() }
func file_buf_test_proto_init() {
	if File_buf_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbsRequest); i {
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
		file_buf_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbsResponse); i {
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
			RawDescriptor: file_buf_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_test_proto_goTypes,
		DependencyIndexes: file_buf_test_proto_depIdxs,
		MessageInfos:      file_buf_test_proto_msgTypes,
	}.Build()
	File_buf_test_proto = out.File
	file_buf_test_proto_rawDesc = nil
	file_buf_test_proto_goTypes = nil
	file_buf_test_proto_depIdxs = nil
}
