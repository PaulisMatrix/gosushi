// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: hello_world.proto

package hello_world

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

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_world_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_world_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_hello_world_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_world_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_hello_world_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_hello_world_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RepeatedHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *RepeatedHelloRequest) Reset() {
	*x = RepeatedHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_world_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatedHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedHelloRequest) ProtoMessage() {}

func (x *RepeatedHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_world_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedHelloRequest.ProtoReflect.Descriptor instead.
func (*RepeatedHelloRequest) Descriptor() ([]byte, []int) {
	return file_hello_world_proto_rawDescGZIP(), []int{2}
}

func (x *RepeatedHelloRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

var File_hello_world_proto protoreflect.FileDescriptor

var file_hello_world_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x14,
	0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x32, 0x9d, 0x01, 0x0a, 0x07, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x12, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x10, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x21, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x70, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_world_proto_rawDescOnce sync.Once
	file_hello_world_proto_rawDescData = file_hello_world_proto_rawDesc
)

func file_hello_world_proto_rawDescGZIP() []byte {
	file_hello_world_proto_rawDescOnce.Do(func() {
		file_hello_world_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_world_proto_rawDescData)
	})
	return file_hello_world_proto_rawDescData
}

var file_hello_world_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hello_world_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),         // 0: hello_world.HelloRequest
	(*HelloReply)(nil),           // 1: hello_world.HelloReply
	(*RepeatedHelloRequest)(nil), // 2: hello_world.RepeatedHelloRequest
}
var file_hello_world_proto_depIdxs = []int32{
	0, // 0: hello_world.Greeter.SayHello:input_type -> hello_world.HelloRequest
	2, // 1: hello_world.Greeter.SayHelloSpecific:input_type -> hello_world.RepeatedHelloRequest
	1, // 2: hello_world.Greeter.SayHello:output_type -> hello_world.HelloReply
	1, // 3: hello_world.Greeter.SayHelloSpecific:output_type -> hello_world.HelloReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_world_proto_init() }
func file_hello_world_proto_init() {
	if File_hello_world_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_world_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_hello_world_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_hello_world_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatedHelloRequest); i {
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
			RawDescriptor: file_hello_world_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_world_proto_goTypes,
		DependencyIndexes: file_hello_world_proto_depIdxs,
		MessageInfos:      file_hello_world_proto_msgTypes,
	}.Build()
	File_hello_world_proto = out.File
	file_hello_world_proto_rawDesc = nil
	file_hello_world_proto_goTypes = nil
	file_hello_world_proto_depIdxs = nil
}
