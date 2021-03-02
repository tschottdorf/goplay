// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: test.proto

package main

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

type A struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	B []byte `protobuf:"bytes,1,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *A) Reset() {
	*x = A{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *A) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*A) ProtoMessage() {}

func (x *A) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use A.ProtoReflect.Descriptor instead.
func (*A) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *A) GetB() []byte {
	if x != nil {
		return x.B
	}
	return nil
}

type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Experimental feature, not generally supported yet!
	I    int32  `protobuf:"varint,1,opt,name=i,proto3" json:"i,omitempty"`
	IOpt *int32 `protobuf:"varint,2,opt,name=i_opt,json=iOpt,proto3,oneof" json:"i_opt,omitempty"`
	A    *A     `protobuf:"bytes,3,opt,name=a,proto3" json:"a,omitempty"`
	AOpt *A     `protobuf:"bytes,4,opt,name=a_opt,json=aOpt,proto3,oneof" json:"a_opt,omitempty"`
}

func (x *Foo) Reset() {
	*x = Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *Foo) GetI() int32 {
	if x != nil {
		return x.I
	}
	return 0
}

func (x *Foo) GetIOpt() int32 {
	if x != nil && x.IOpt != nil {
		return *x.IOpt
	}
	return 0
}

func (x *Foo) GetA() *A {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *Foo) GetAOpt() *A {
	if x != nil {
		return x.AOpt
	}
	return nil
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x01,
	0x41, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x62, 0x22,
	0x71, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x0c, 0x0a, 0x01, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x69, 0x12, 0x18, 0x0a, 0x05, 0x69, 0x5f, 0x6f, 0x70, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x69, 0x4f, 0x70, 0x74, 0x88, 0x01, 0x01, 0x12, 0x10,
	0x0a, 0x01, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x02, 0x2e, 0x41, 0x52, 0x01, 0x61,
	0x12, 0x1c, 0x0a, 0x05, 0x61, 0x5f, 0x6f, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x02, 0x2e, 0x41, 0x48, 0x01, 0x52, 0x04, 0x61, 0x4f, 0x70, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x69, 0x5f, 0x6f, 0x70, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61, 0x5f, 0x6f,
	0x70, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_proto_goTypes = []interface{}{
	(*A)(nil),   // 0: A
	(*Foo)(nil), // 1: Foo
}
var file_test_proto_depIdxs = []int32{
	0, // 0: Foo.a:type_name -> A
	0, // 1: Foo.a_opt:type_name -> A
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*A); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo); i {
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
	file_test_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
