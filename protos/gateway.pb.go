// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: gateway.proto

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

type WindowSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval int32 `protobuf:"varint,1,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *WindowSize) Reset() {
	*x = WindowSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindowSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindowSize) ProtoMessage() {}

func (x *WindowSize) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindowSize.ProtoReflect.Descriptor instead.
func (*WindowSize) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *WindowSize) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

type WindowDataArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime string  `protobuf:"bytes,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	Time      uint64  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Open      float64 `protobuf:"fixed64,3,opt,name=open,proto3" json:"open,omitempty"`
	High      float64 `protobuf:"fixed64,4,opt,name=high,proto3" json:"high,omitempty"`
	Low       float64 `protobuf:"fixed64,5,opt,name=low,proto3" json:"low,omitempty"`
	Close     float64 `protobuf:"fixed64,6,opt,name=close,proto3" json:"close,omitempty"`
	Volume    float64 `protobuf:"fixed64,7,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *WindowDataArray) Reset() {
	*x = WindowDataArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindowDataArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindowDataArray) ProtoMessage() {}

func (x *WindowDataArray) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindowDataArray.ProtoReflect.Descriptor instead.
func (*WindowDataArray) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *WindowDataArray) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *WindowDataArray) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *WindowDataArray) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *WindowDataArray) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *WindowDataArray) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *WindowDataArray) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *WindowDataArray) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

var File_gateway_proto protoreflect.FileDescriptor

var file_gateway_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x28, 0x0a, 0x0a, 0x77, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x22, 0xab, 0x01, 0x0a, 0x0f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61,
	0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x69, 0x67, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f,
	0x77, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x32,
	0x4d, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x4f, 0x48, 0x4c, 0x43, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x53,
	0x69, 0x7a, 0x65, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x77, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x41, 0x72, 0x72, 0x61, 0x79, 0x22, 0x00, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x65,
	0x72, 0x69, 0x61, 0x2f, 0x67, 0x6f, 0x66, 0x74, 0x78, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_gateway_proto_rawDescOnce sync.Once
	file_gateway_proto_rawDescData = file_gateway_proto_rawDesc
)

func file_gateway_proto_rawDescGZIP() []byte {
	file_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_proto_rawDescData)
	})
	return file_gateway_proto_rawDescData
}

var file_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gateway_proto_goTypes = []interface{}{
	(*WindowSize)(nil),      // 0: protos.windowSize
	(*WindowDataArray)(nil), // 1: protos.windowDataArray
}
var file_gateway_proto_depIdxs = []int32{
	0, // 0: protos.chartDataExchange.getOHLC:input_type -> protos.windowSize
	1, // 1: protos.chartDataExchange.getOHLC:output_type -> protos.windowDataArray
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gateway_proto_init() }
func file_gateway_proto_init() {
	if File_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindowSize); i {
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
		file_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindowDataArray); i {
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
			RawDescriptor: file_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_proto_depIdxs,
		MessageInfos:      file_gateway_proto_msgTypes,
	}.Build()
	File_gateway_proto = out.File
	file_gateway_proto_rawDesc = nil
	file_gateway_proto_goTypes = nil
	file_gateway_proto_depIdxs = nil
}
