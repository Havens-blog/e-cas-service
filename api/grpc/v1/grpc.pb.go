// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: grpc/v1/grpc.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/google/gnostic/openapiv3"
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

// 获取用户信息请求结构体
type UserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名
	UserName string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *UserInfoRequest) Reset() {
	*x = UserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_v1_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRequest) ProtoMessage() {}

func (x *UserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_v1_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRequest.ProtoReflect.Descriptor instead.
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return file_grpc_v1_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfoRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

// 获取用户信息返回结构体
type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_v1_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_v1_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_grpc_v1_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfoResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_grpc_v1_grpc_proto protoreflect.FileDescriptor

var file_grpc_v1_grpc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18, 0x0b, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x52, 0x0a, 0x04, 0x47, 0x72,
	0x70, 0x63, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa2,
	0x03, 0xba, 0x47, 0xdb, 0x02, 0x12, 0xcd, 0x01, 0x0a, 0x0e, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x1b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x20, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x61, 0x6f, 0x68, 0x75,
	0x62, 0x61, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x6c, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x1a, 0x15, 0x78, 0x69, 0x61, 0x6f, 0x68, 0x75, 0x62, 0x61, 0x69, 0x40, 0x6f, 0x75,
	0x74, 0x6c, 0x6f, 0x6f, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x42, 0x0a, 0x0b, 0x4d, 0x49, 0x54,
	0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x61,
	0x6f, 0x68, 0x75, 0x62, 0x61, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x6c,
	0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x32, 0x06, 0x76,
	0x31, 0x2e, 0x30, 0x2e, 0x30, 0x1a, 0x26, 0x0a, 0x16, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f,
	0x31, 0x37, 0x32, 0x2e, 0x32, 0x31, 0x2e, 0x30, 0x2e, 0x32, 0x3a, 0x38, 0x30, 0x30, 0x30, 0x12,
	0x0c, 0xe6, 0xb5, 0x8b, 0xe8, 0xaf, 0x95, 0xe7, 0x8e, 0xaf, 0xe5, 0xa2, 0x83, 0x1a, 0x26, 0x0a,
	0x16, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x31, 0x37, 0x32, 0x2e, 0x32, 0x31, 0x2e, 0x30,
	0x2e, 0x32, 0x3a, 0x38, 0x30, 0x30, 0x30, 0x12, 0x0c, 0xe7, 0xba, 0xbf, 0xe4, 0xb8, 0x8a, 0xe7,
	0x8e, 0xaf, 0xe5, 0xa2, 0x83, 0x2a, 0x27, 0x3a, 0x25, 0x0a, 0x23, 0x0a, 0x0a, 0x62, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x15, 0x0a, 0x13, 0x0a, 0x04, 0x68, 0x74, 0x74,
	0x70, 0x2a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x32, 0x03, 0x4a, 0x57, 0x54, 0x32, 0x10,
	0x0a, 0x0e, 0x0a, 0x0a, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00,
	0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x61, 0x6f,
	0x68, 0x75, 0x62, 0x61, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x6c, 0x61,
	0x79, 0x6f, 0x75, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_v1_grpc_proto_rawDescOnce sync.Once
	file_grpc_v1_grpc_proto_rawDescData = file_grpc_v1_grpc_proto_rawDesc
)

func file_grpc_v1_grpc_proto_rawDescGZIP() []byte {
	file_grpc_v1_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_v1_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_v1_grpc_proto_rawDescData)
	})
	return file_grpc_v1_grpc_proto_rawDescData
}

var file_grpc_v1_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_v1_grpc_proto_goTypes = []interface{}{
	(*UserInfoRequest)(nil),  // 0: api.grpc.v1.UserInfoRequest
	(*UserInfoResponse)(nil), // 1: api.grpc.v1.UserInfoResponse
}
var file_grpc_v1_grpc_proto_depIdxs = []int32{
	0, // 0: api.grpc.v1.Grpc.GetUserInfo:input_type -> api.grpc.v1.UserInfoRequest
	1, // 1: api.grpc.v1.Grpc.GetUserInfo:output_type -> api.grpc.v1.UserInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_v1_grpc_proto_init() }
func file_grpc_v1_grpc_proto_init() {
	if File_grpc_v1_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_v1_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoRequest); i {
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
		file_grpc_v1_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResponse); i {
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
			RawDescriptor: file_grpc_v1_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_v1_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_v1_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_v1_grpc_proto_msgTypes,
	}.Build()
	File_grpc_v1_grpc_proto = out.File
	file_grpc_v1_grpc_proto_rawDesc = nil
	file_grpc_v1_grpc_proto_goTypes = nil
	file_grpc_v1_grpc_proto_depIdxs = nil
}