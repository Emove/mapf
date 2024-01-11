// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: common/data/v1/data.proto

package v1

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

type EnableStatus int32

const (
	EnableStatus_PLACEHOLDER_ENABLE_STATUS EnableStatus = 0
	EnableStatus_ENABLE                    EnableStatus = 1
	EnableStatus_DISABLE                   EnableStatus = 2
)

// Enum value maps for EnableStatus.
var (
	EnableStatus_name = map[int32]string{
		0: "PLACEHOLDER_ENABLE_STATUS",
		1: "ENABLE",
		2: "DISABLE",
	}
	EnableStatus_value = map[string]int32{
		"PLACEHOLDER_ENABLE_STATUS": 0,
		"ENABLE":                    1,
		"DISABLE":                   2,
	}
)

func (x EnableStatus) Enum() *EnableStatus {
	p := new(EnableStatus)
	*p = x
	return p
}

func (x EnableStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnableStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_data_v1_data_proto_enumTypes[0].Descriptor()
}

func (EnableStatus) Type() protoreflect.EnumType {
	return &file_common_data_v1_data_proto_enumTypes[0]
}

func (x EnableStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnableStatus.Descriptor instead.
func (EnableStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_data_v1_data_proto_rawDescGZIP(), []int{0}
}

type DefaultStatus int32

const (
	DefaultStatus_PLACEHOLDER_DEFAULT_STATUS DefaultStatus = 0
	DefaultStatus_IS_DEFAULT                 DefaultStatus = 1
	DefaultStatus_NOT_DEFAULT                DefaultStatus = 2
)

// Enum value maps for DefaultStatus.
var (
	DefaultStatus_name = map[int32]string{
		0: "PLACEHOLDER_DEFAULT_STATUS",
		1: "IS_DEFAULT",
		2: "NOT_DEFAULT",
	}
	DefaultStatus_value = map[string]int32{
		"PLACEHOLDER_DEFAULT_STATUS": 0,
		"IS_DEFAULT":                 1,
		"NOT_DEFAULT":                2,
	}
)

func (x DefaultStatus) Enum() *DefaultStatus {
	p := new(DefaultStatus)
	*p = x
	return p
}

func (x DefaultStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DefaultStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_data_v1_data_proto_enumTypes[1].Descriptor()
}

func (DefaultStatus) Type() protoreflect.EnumType {
	return &file_common_data_v1_data_proto_enumTypes[1]
}

func (x DefaultStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DefaultStatus.Descriptor instead.
func (DefaultStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_data_v1_data_proto_rawDescGZIP(), []int{1}
}

type ValidStatus int32

const (
	ValidStatus_PLACEHOLDER_VALID_STATUS ValidStatus = 0
	ValidStatus_VALID                    ValidStatus = 1
	ValidStatus_INVALID                  ValidStatus = 2
)

// Enum value maps for ValidStatus.
var (
	ValidStatus_name = map[int32]string{
		0: "PLACEHOLDER_VALID_STATUS",
		1: "VALID",
		2: "INVALID",
	}
	ValidStatus_value = map[string]int32{
		"PLACEHOLDER_VALID_STATUS": 0,
		"VALID":                    1,
		"INVALID":                  2,
	}
)

func (x ValidStatus) Enum() *ValidStatus {
	p := new(ValidStatus)
	*p = x
	return p
}

func (x ValidStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValidStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_data_v1_data_proto_enumTypes[2].Descriptor()
}

func (ValidStatus) Type() protoreflect.EnumType {
	return &file_common_data_v1_data_proto_enumTypes[2]
}

func (x ValidStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValidStatus.Descriptor instead.
func (ValidStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_data_v1_data_proto_rawDescGZIP(), []int{2}
}

var File_common_data_v1_data_proto protoreflect.FileDescriptor

var file_common_data_v1_data_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2a, 0x46, 0x0a, 0x0c, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x4e,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x02, 0x2a, 0x50, 0x0a, 0x0d, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c,
	0x44, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x53, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x41,
	0x55, 0x4c, 0x54, 0x10, 0x02, 0x2a, 0x43, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c,
	0x44, 0x45, 0x52, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x02, 0x42, 0x1c, 0x5a, 0x1a, 0x6d, 0x61,
	0x70, 0x66, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_data_v1_data_proto_rawDescOnce sync.Once
	file_common_data_v1_data_proto_rawDescData = file_common_data_v1_data_proto_rawDesc
)

func file_common_data_v1_data_proto_rawDescGZIP() []byte {
	file_common_data_v1_data_proto_rawDescOnce.Do(func() {
		file_common_data_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_data_v1_data_proto_rawDescData)
	})
	return file_common_data_v1_data_proto_rawDescData
}

var file_common_data_v1_data_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_common_data_v1_data_proto_goTypes = []interface{}{
	(EnableStatus)(0),  // 0: common.data.v1.EnableStatus
	(DefaultStatus)(0), // 1: common.data.v1.DefaultStatus
	(ValidStatus)(0),   // 2: common.data.v1.ValidStatus
}
var file_common_data_v1_data_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_data_v1_data_proto_init() }
func file_common_data_v1_data_proto_init() {
	if File_common_data_v1_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_data_v1_data_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_data_v1_data_proto_goTypes,
		DependencyIndexes: file_common_data_v1_data_proto_depIdxs,
		EnumInfos:         file_common_data_v1_data_proto_enumTypes,
	}.Build()
	File_common_data_v1_data_proto = out.File
	file_common_data_v1_data_proto_rawDesc = nil
	file_common_data_v1_data_proto_goTypes = nil
	file_common_data_v1_data_proto_depIdxs = nil
}
