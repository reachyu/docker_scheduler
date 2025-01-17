// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: docker_api.proto

package go_docker

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message containing the user's name.
type ReqContainerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerId string `protobuf:"bytes,1,opt,name=containerId,proto3" json:"containerId,omitempty"`
}

func (x *ReqContainerInfo) Reset() {
	*x = ReqContainerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqContainerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqContainerInfo) ProtoMessage() {}

func (x *ReqContainerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_docker_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqContainerInfo.ProtoReflect.Descriptor instead.
func (*ReqContainerInfo) Descriptor() ([]byte, []int) {
	return file_docker_api_proto_rawDescGZIP(), []int{0}
}

func (x *ReqContainerInfo) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

type ReqImageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageName string `protobuf:"bytes,1,opt,name=imageName,proto3" json:"imageName,omitempty"`
}

func (x *ReqImageInfo) Reset() {
	*x = ReqImageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqImageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqImageInfo) ProtoMessage() {}

func (x *ReqImageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_docker_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqImageInfo.ProtoReflect.Descriptor instead.
func (*ReqImageInfo) Descriptor() ([]byte, []int) {
	return file_docker_api_proto_rawDescGZIP(), []int{1}
}

func (x *ReqImageInfo) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

// The response message containing the greetings
type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_docker_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_docker_api_proto_rawDescGZIP(), []int{2}
}

func (x *RunResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StopResponse) Reset() {
	*x = StopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopResponse) ProtoMessage() {}

func (x *StopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_docker_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopResponse.ProtoReflect.Descriptor instead.
func (*StopResponse) Descriptor() ([]byte, []int) {
	return file_docker_api_proto_rawDescGZIP(), []int{3}
}

func (x *StopResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docker_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_docker_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_docker_api_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_docker_api_proto protoreflect.FileDescriptor

var file_docker_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x61, 0x70, 0x69, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0c, 0x52,
	0x65, 0x71, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x91, 0x02, 0x0a, 0x10, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a,
	0x0c, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x70, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x52,
	0x65, 0x71, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x12, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x52, 0x65, 0x71, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b,
	0x2e, 0x2f, 0x67, 0x6f, 0x5f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_docker_api_proto_rawDescOnce sync.Once
	file_docker_api_proto_rawDescData = file_docker_api_proto_rawDesc
)

func file_docker_api_proto_rawDescGZIP() []byte {
	file_docker_api_proto_rawDescOnce.Do(func() {
		file_docker_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_docker_api_proto_rawDescData)
	})
	return file_docker_api_proto_rawDescData
}

var file_docker_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_docker_api_proto_goTypes = []interface{}{
	(*ReqContainerInfo)(nil), // 0: apis.ReqContainerInfo
	(*ReqImageInfo)(nil),     // 1: apis.ReqImageInfo
	(*RunResponse)(nil),      // 2: apis.RunResponse
	(*StopResponse)(nil),     // 3: apis.StopResponse
	(*DeleteResponse)(nil),   // 4: apis.DeleteResponse
	(*empty.Empty)(nil),      // 5: google.protobuf.Empty
}
var file_docker_api_proto_depIdxs = []int32{
	1, // 0: apis.ContainerService.RunContainer:input_type -> apis.ReqImageInfo
	0, // 1: apis.ContainerService.StopContainer:input_type -> apis.ReqContainerInfo
	0, // 2: apis.ContainerService.DeleteContainer:input_type -> apis.ReqContainerInfo
	5, // 3: apis.ContainerService.ListContainers:input_type -> google.protobuf.Empty
	2, // 4: apis.ContainerService.RunContainer:output_type -> apis.RunResponse
	3, // 5: apis.ContainerService.StopContainer:output_type -> apis.StopResponse
	4, // 6: apis.ContainerService.DeleteContainer:output_type -> apis.DeleteResponse
	5, // 7: apis.ContainerService.ListContainers:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_docker_api_proto_init() }
func file_docker_api_proto_init() {
	if File_docker_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_docker_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqContainerInfo); i {
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
		file_docker_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqImageInfo); i {
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
		file_docker_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunResponse); i {
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
		file_docker_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopResponse); i {
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
		file_docker_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_docker_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_docker_api_proto_goTypes,
		DependencyIndexes: file_docker_api_proto_depIdxs,
		MessageInfos:      file_docker_api_proto_msgTypes,
	}.Build()
	File_docker_api_proto = out.File
	file_docker_api_proto_rawDesc = nil
	file_docker_api_proto_goTypes = nil
	file_docker_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ContainerServiceClient is the client API for ContainerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContainerServiceClient interface {
	// Sends a greeting
	RunContainer(ctx context.Context, in *ReqImageInfo, opts ...grpc.CallOption) (*RunResponse, error)
	StopContainer(ctx context.Context, in *ReqContainerInfo, opts ...grpc.CallOption) (*StopResponse, error)
	DeleteContainer(ctx context.Context, in *ReqContainerInfo, opts ...grpc.CallOption) (*DeleteResponse, error)
	ListContainers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type containerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContainerServiceClient(cc grpc.ClientConnInterface) ContainerServiceClient {
	return &containerServiceClient{cc}
}

func (c *containerServiceClient) RunContainer(ctx context.Context, in *ReqImageInfo, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := c.cc.Invoke(ctx, "/apis.ContainerService/RunContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) StopContainer(ctx context.Context, in *ReqContainerInfo, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/apis.ContainerService/StopContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) DeleteContainer(ctx context.Context, in *ReqContainerInfo, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/apis.ContainerService/DeleteContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) ListContainers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/apis.ContainerService/ListContainers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContainerServiceServer is the server API for ContainerService service.
type ContainerServiceServer interface {
	// Sends a greeting
	RunContainer(context.Context, *ReqImageInfo) (*RunResponse, error)
	StopContainer(context.Context, *ReqContainerInfo) (*StopResponse, error)
	DeleteContainer(context.Context, *ReqContainerInfo) (*DeleteResponse, error)
	ListContainers(context.Context, *empty.Empty) (*empty.Empty, error)
}

// UnimplementedContainerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedContainerServiceServer struct {
}

func (*UnimplementedContainerServiceServer) RunContainer(context.Context, *ReqImageInfo) (*RunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunContainer not implemented")
}
func (*UnimplementedContainerServiceServer) StopContainer(context.Context, *ReqContainerInfo) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopContainer not implemented")
}
func (*UnimplementedContainerServiceServer) DeleteContainer(context.Context, *ReqContainerInfo) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContainer not implemented")
}
func (*UnimplementedContainerServiceServer) ListContainers(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContainers not implemented")
}

func RegisterContainerServiceServer(s *grpc.Server, srv ContainerServiceServer) {
	s.RegisterService(&_ContainerService_serviceDesc, srv)
}

func _ContainerService_RunContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqImageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).RunContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.ContainerService/RunContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).RunContainer(ctx, req.(*ReqImageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_StopContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqContainerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).StopContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.ContainerService/StopContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).StopContainer(ctx, req.(*ReqContainerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_DeleteContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqContainerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).DeleteContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.ContainerService/DeleteContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).DeleteContainer(ctx, req.(*ReqContainerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_ListContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).ListContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.ContainerService/ListContainers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).ListContainers(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContainerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apis.ContainerService",
	HandlerType: (*ContainerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunContainer",
			Handler:    _ContainerService_RunContainer_Handler,
		},
		{
			MethodName: "StopContainer",
			Handler:    _ContainerService_StopContainer_Handler,
		},
		{
			MethodName: "DeleteContainer",
			Handler:    _ContainerService_DeleteContainer_Handler,
		},
		{
			MethodName: "ListContainers",
			Handler:    _ContainerService_ListContainers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "docker_api.proto",
}
