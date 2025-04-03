// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: client/controller/v1/service.proto

package v1

import (
	endpoint "github.com/bayu-aditya/ideagate/backend/model/gen-go/core/endpoint"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetListEndpointRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetListEndpointRequest) Reset() {
	*x = GetListEndpointRequest{}
	mi := &file_client_controller_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListEndpointRequest) ProtoMessage() {}

func (x *GetListEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_controller_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListEndpointRequest.ProtoReflect.Descriptor instead.
func (*GetListEndpointRequest) Descriptor() ([]byte, []int) {
	return file_client_controller_v1_service_proto_rawDescGZIP(), []int{0}
}

type GetListEndpointResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Endpoints     []*endpoint.Endpoint   `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetListEndpointResponse) Reset() {
	*x = GetListEndpointResponse{}
	mi := &file_client_controller_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListEndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListEndpointResponse) ProtoMessage() {}

func (x *GetListEndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_controller_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListEndpointResponse.ProtoReflect.Descriptor instead.
func (*GetListEndpointResponse) Descriptor() ([]byte, []int) {
	return file_client_controller_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetListEndpointResponse) GetEndpoints() []*endpoint.Endpoint {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

var File_client_controller_v1_service_proto protoreflect.FileDescriptor

const file_client_controller_v1_service_proto_rawDesc = "" +
	"\n" +
	"\"client/controller/v1/service.proto\x12\x02v1\x1a\x1ccore/endpoint/endpoint.proto\"\x18\n" +
	"\x16GetListEndpointRequest\"K\n" +
	"\x17GetListEndpointResponse\x120\n" +
	"\tendpoints\x18\x01 \x03(\v2\x12.endpoint.EndpointR\tendpoints2_\n" +
	"\x11ControllerService\x12J\n" +
	"\x0fGetListEndpoint\x12\x1a.v1.GetListEndpointRequest\x1a\x1b.v1.GetListEndpointResponseB\x89\x01\n" +
	"\x06com.v1B\fServiceProtoP\x01ZIgithub.com/bayu-aditya/ideagate/backend/model/gen-go/client/controller/v1\xa2\x02\x03VXX\xaa\x02\x02V1\xca\x02\x02V1\xe2\x02\x0eV1\\GPBMetadata\xea\x02\x02V1b\x06proto3"

var (
	file_client_controller_v1_service_proto_rawDescOnce sync.Once
	file_client_controller_v1_service_proto_rawDescData []byte
)

func file_client_controller_v1_service_proto_rawDescGZIP() []byte {
	file_client_controller_v1_service_proto_rawDescOnce.Do(func() {
		file_client_controller_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_client_controller_v1_service_proto_rawDesc), len(file_client_controller_v1_service_proto_rawDesc)))
	})
	return file_client_controller_v1_service_proto_rawDescData
}

var file_client_controller_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_controller_v1_service_proto_goTypes = []any{
	(*GetListEndpointRequest)(nil),  // 0: v1.GetListEndpointRequest
	(*GetListEndpointResponse)(nil), // 1: v1.GetListEndpointResponse
	(*endpoint.Endpoint)(nil),       // 2: endpoint.Endpoint
}
var file_client_controller_v1_service_proto_depIdxs = []int32{
	2, // 0: v1.GetListEndpointResponse.endpoints:type_name -> endpoint.Endpoint
	0, // 1: v1.ControllerService.GetListEndpoint:input_type -> v1.GetListEndpointRequest
	1, // 2: v1.ControllerService.GetListEndpoint:output_type -> v1.GetListEndpointResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_client_controller_v1_service_proto_init() }
func file_client_controller_v1_service_proto_init() {
	if File_client_controller_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_client_controller_v1_service_proto_rawDesc), len(file_client_controller_v1_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_controller_v1_service_proto_goTypes,
		DependencyIndexes: file_client_controller_v1_service_proto_depIdxs,
		MessageInfos:      file_client_controller_v1_service_proto_msgTypes,
	}.Build()
	File_client_controller_v1_service_proto = out.File
	file_client_controller_v1_service_proto_goTypes = nil
	file_client_controller_v1_service_proto_depIdxs = nil
}
