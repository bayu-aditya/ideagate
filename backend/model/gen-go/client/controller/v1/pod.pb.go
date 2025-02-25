// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: client/controller/v1/pod.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type PodStatus int32

const (
	PodStatus_POD_STATUS_UNSPECIFIED PodStatus = 0
	PodStatus_POD_STATUS_PENDING     PodStatus = 1
	PodStatus_POD_STATUS_RUNNING     PodStatus = 2
	PodStatus_POD_STATUS_SUCCEEDED   PodStatus = 3
	PodStatus_POD_STATUS_FAILED      PodStatus = 4
)

// Enum value maps for PodStatus.
var (
	PodStatus_name = map[int32]string{
		0: "POD_STATUS_UNSPECIFIED",
		1: "POD_STATUS_PENDING",
		2: "POD_STATUS_RUNNING",
		3: "POD_STATUS_SUCCEEDED",
		4: "POD_STATUS_FAILED",
	}
	PodStatus_value = map[string]int32{
		"POD_STATUS_UNSPECIFIED": 0,
		"POD_STATUS_PENDING":     1,
		"POD_STATUS_RUNNING":     2,
		"POD_STATUS_SUCCEEDED":   3,
		"POD_STATUS_FAILED":      4,
	}
)

func (x PodStatus) Enum() *PodStatus {
	p := new(PodStatus)
	*p = x
	return p
}

func (x PodStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PodStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_client_controller_v1_pod_proto_enumTypes[0].Descriptor()
}

func (PodStatus) Type() protoreflect.EnumType {
	return &file_client_controller_v1_pod_proto_enumTypes[0]
}

func (x PodStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PodStatus.Descriptor instead.
func (PodStatus) EnumDescriptor() ([]byte, []int) {
	return file_client_controller_v1_pod_proto_rawDescGZIP(), []int{0}
}

type Pod struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace     string                 `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Status        PodStatus              `protobuf:"varint,4,opt,name=status,proto3,enum=v1.PodStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pod) Reset() {
	*x = Pod{}
	mi := &file_client_controller_v1_pod_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pod) ProtoMessage() {}

func (x *Pod) ProtoReflect() protoreflect.Message {
	mi := &file_client_controller_v1_pod_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pod.ProtoReflect.Descriptor instead.
func (*Pod) Descriptor() ([]byte, []int) {
	return file_client_controller_v1_pod_proto_rawDescGZIP(), []int{0}
}

func (x *Pod) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pod) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Pod) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Pod) GetStatus() PodStatus {
	if x != nil {
		return x.Status
	}
	return PodStatus_POD_STATUS_UNSPECIFIED
}

var File_client_controller_v1_pod_proto protoreflect.FileDescriptor

var file_client_controller_v1_pod_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x03, 0x50, 0x6f, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2a, 0x88, 0x01, 0x0a, 0x09, 0x50, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1a, 0x0a, 0x16, 0x50, 0x4f, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x50,
	0x4f, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x4f, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x50,
	0x4f, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45,
	0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x4f, 0x44, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x42, 0x85, 0x01, 0x0a,
	0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x50, 0x6f, 0x64, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x61, 0x79, 0x75, 0x2d, 0x61, 0x64, 0x69, 0x74, 0x79, 0x61, 0x2f, 0x69, 0x64, 0x65, 0x61,
	0x67, 0x61, 0x74, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02,
	0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_client_controller_v1_pod_proto_rawDescOnce sync.Once
	file_client_controller_v1_pod_proto_rawDescData []byte
)

func file_client_controller_v1_pod_proto_rawDescGZIP() []byte {
	file_client_controller_v1_pod_proto_rawDescOnce.Do(func() {
		file_client_controller_v1_pod_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_client_controller_v1_pod_proto_rawDesc), len(file_client_controller_v1_pod_proto_rawDesc)))
	})
	return file_client_controller_v1_pod_proto_rawDescData
}

var file_client_controller_v1_pod_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_client_controller_v1_pod_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_client_controller_v1_pod_proto_goTypes = []any{
	(PodStatus)(0),                // 0: v1.PodStatus
	(*Pod)(nil),                   // 1: v1.Pod
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_client_controller_v1_pod_proto_depIdxs = []int32{
	2, // 0: v1.Pod.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: v1.Pod.status:type_name -> v1.PodStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_client_controller_v1_pod_proto_init() }
func file_client_controller_v1_pod_proto_init() {
	if File_client_controller_v1_pod_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_client_controller_v1_pod_proto_rawDesc), len(file_client_controller_v1_pod_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_controller_v1_pod_proto_goTypes,
		DependencyIndexes: file_client_controller_v1_pod_proto_depIdxs,
		EnumInfos:         file_client_controller_v1_pod_proto_enumTypes,
		MessageInfos:      file_client_controller_v1_pod_proto_msgTypes,
	}.Build()
	File_client_controller_v1_pod_proto = out.File
	file_client_controller_v1_pod_proto_goTypes = nil
	file_client_controller_v1_pod_proto_depIdxs = nil
}
