// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: rpcService.proto

package protos

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

type OHLCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryString string `protobuf:"bytes,1,opt,name=queryString,proto3" json:"queryString,omitempty"`
}

func (x *OHLCRequest) Reset() {
	*x = OHLCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OHLCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OHLCRequest) ProtoMessage() {}

func (x *OHLCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OHLCRequest.ProtoReflect.Descriptor instead.
func (*OHLCRequest) Descriptor() ([]byte, []int) {
	return file_rpcService_proto_rawDescGZIP(), []int{0}
}

func (x *OHLCRequest) GetQueryString() string {
	if x != nil {
		return x.QueryString
	}
	return ""
}

type OHLCResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryResponse string `protobuf:"bytes,1,opt,name=queryResponse,proto3" json:"queryResponse,omitempty"`
}

func (x *OHLCResponse) Reset() {
	*x = OHLCResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OHLCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OHLCResponse) ProtoMessage() {}

func (x *OHLCResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OHLCResponse.ProtoReflect.Descriptor instead.
func (*OHLCResponse) Descriptor() ([]byte, []int) {
	return file_rpcService_proto_rawDescGZIP(), []int{1}
}

func (x *OHLCResponse) GetQueryResponse() string {
	if x != nil {
		return x.QueryResponse
	}
	return ""
}

var File_rpcService_proto protoreflect.FileDescriptor

var file_rpcService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x2f, 0x0a, 0x0b, 0x4f, 0x48,
	0x4c, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x34, 0x0a, 0x0c, 0x4f,
	0x48, 0x4c, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x40, 0x0a, 0x0b, 0x4f, 0x48, 0x4c, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x31, 0x0a, 0x04, 0x4f, 0x48, 0x4c, 0x43, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x4f, 0x48, 0x4c, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4f, 0x48, 0x4c, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x6f, 0x65, 0x72, 0x69, 0x61, 0x2f, 0x67, 0x6f, 0x66, 0x74, 0x78, 0x6c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpcService_proto_rawDescOnce sync.Once
	file_rpcService_proto_rawDescData = file_rpcService_proto_rawDesc
)

func file_rpcService_proto_rawDescGZIP() []byte {
	file_rpcService_proto_rawDescOnce.Do(func() {
		file_rpcService_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpcService_proto_rawDescData)
	})
	return file_rpcService_proto_rawDescData
}

var file_rpcService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpcService_proto_goTypes = []interface{}{
	(*OHLCRequest)(nil),  // 0: protos.OHLCRequest
	(*OHLCResponse)(nil), // 1: protos.OHLCResponse
}
var file_rpcService_proto_depIdxs = []int32{
	0, // 0: protos.OHLCService.OHLC:input_type -> protos.OHLCRequest
	1, // 1: protos.OHLCService.OHLC:output_type -> protos.OHLCResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpcService_proto_init() }
func file_rpcService_proto_init() {
	if File_rpcService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpcService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OHLCRequest); i {
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
		file_rpcService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OHLCResponse); i {
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
			RawDescriptor: file_rpcService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpcService_proto_goTypes,
		DependencyIndexes: file_rpcService_proto_depIdxs,
		MessageInfos:      file_rpcService_proto_msgTypes,
	}.Build()
	File_rpcService_proto = out.File
	file_rpcService_proto_rawDesc = nil
	file_rpcService_proto_goTypes = nil
	file_rpcService_proto_depIdxs = nil
}