// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.0
// source: test.proto

package test

import (
	_ "github.com/wangzhe1991/protoc-gen-go-errorx/gerr"
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

type TestErrorReason int32

const (
	// 自定义错误
	TestErrorReason_OK           TestErrorReason = 0
	TestErrorReason_Err          TestErrorReason = 1
	TestErrorReason_Unknown      TestErrorReason = 2
	TestErrorReason_DB           TestErrorReason = 3
	TestErrorReason_Redis        TestErrorReason = 4
	TestErrorReason_GRPC         TestErrorReason = 5
	TestErrorReason_NotFound     TestErrorReason = 6
	TestErrorReason_NoPermission TestErrorReason = 7
	TestErrorReason_Existed      TestErrorReason = 8
	TestErrorReason_ParamInvalid TestErrorReason = 9
)

// Enum value maps for TestErrorReason.
var (
	TestErrorReason_name = map[int32]string{
		0: "OK",
		1: "Err",
		2: "Unknown",
		3: "DB",
		4: "Redis",
		5: "GRPC",
		6: "NotFound",
		7: "NoPermission",
		8: "Existed",
		9: "ParamInvalid",
	}
	TestErrorReason_value = map[string]int32{
		"OK":           0,
		"Err":          1,
		"Unknown":      2,
		"DB":           3,
		"Redis":        4,
		"GRPC":         5,
		"NotFound":     6,
		"NoPermission": 7,
		"Existed":      8,
		"ParamInvalid": 9,
	}
)

func (x TestErrorReason) Enum() *TestErrorReason {
	p := new(TestErrorReason)
	*p = x
	return p
}

func (x TestErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[0].Descriptor()
}

func (TestErrorReason) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[0]
}

func (x TestErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestErrorReason.Descriptor instead.
func (TestErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xbf, 0x03, 0x0a, 0x0f, 0x54, 0x65, 0x73, 0x74,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x00, 0x1a, 0x0e, 0xaa, 0x45, 0x0b, 0x08, 0xc8, 0x01, 0x10, 0xa2, 0x8d, 0x06, 0x1a,
	0x02, 0x6f, 0x6b, 0x12, 0x2b, 0x0a, 0x03, 0x45, 0x72, 0x72, 0x10, 0x01, 0x1a, 0x22, 0xaa, 0x45,
	0x1f, 0x08, 0xf4, 0x03, 0x10, 0xa3, 0x8d, 0x06, 0x1a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x12, 0x27, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x02, 0x1a, 0x1a, 0xaa,
	0x45, 0x17, 0x08, 0xf4, 0x03, 0x10, 0xa4, 0x8d, 0x06, 0x1a, 0x0e, 0x75, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x20, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x02, 0x44, 0x42, 0x10,
	0x03, 0x1a, 0x1e, 0xaa, 0x45, 0x1b, 0x08, 0xf4, 0x03, 0x10, 0xa5, 0x8d, 0x06, 0x1a, 0x12, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x64, 0x62, 0x20, 0x66, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x12, 0x2c, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x10, 0x04, 0x1a, 0x21, 0xaa, 0x45,
	0x1e, 0x08, 0xf4, 0x03, 0x10, 0xa6, 0x8d, 0x06, 0x1a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x20, 0x72, 0x65, 0x64, 0x69, 0x73, 0x20, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12,
	0x2a, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43, 0x10, 0x05, 0x1a, 0x20, 0xaa, 0x45, 0x1d, 0x08, 0xf4,
	0x03, 0x10, 0xa7, 0x8d, 0x06, 0x1a, 0x14, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20,
	0x67, 0x72, 0x70, 0x63, 0x20, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x4e,
	0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x06, 0x1a, 0x15, 0xaa, 0x45, 0x12, 0x08, 0x90,
	0x03, 0x10, 0xa8, 0x8d, 0x06, 0x1a, 0x09, 0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x12, 0x2b, 0x0a, 0x0c, 0x4e, 0x6f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x10, 0x07, 0x1a, 0x19, 0xaa, 0x45, 0x16, 0x08, 0x90, 0x03, 0x10, 0xa9, 0x8d, 0x06, 0x1a, 0x0d,
	0x6e, 0x6f, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a,
	0x07, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x64, 0x10, 0x08, 0x1a, 0x1b, 0xaa, 0x45, 0x18, 0x08,
	0x90, 0x03, 0x10, 0xaa, 0x8d, 0x06, 0x1a, 0x0f, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x20,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x09, 0x1a, 0x19, 0xaa, 0x45, 0x16, 0x08, 0x90,
	0x03, 0x10, 0xab, 0x8d, 0x06, 0x1a, 0x0d, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x1a, 0x13, 0xa2, 0x45, 0x10, 0x08, 0xf4, 0x03, 0x10, 0xa1, 0x8d, 0x06,
	0x1a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x42, 0x12, 0x5a, 0x10, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_proto_goTypes = []interface{}{
	(TestErrorReason)(0), // 0: errors.test.TestErrorReason
}
var file_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		EnumInfos:         file_test_proto_enumTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
