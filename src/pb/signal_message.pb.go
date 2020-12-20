// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: signal_message.proto

package pb

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

type Signal_SignalState int32

const (
	Signal_UNKNOWN       Signal_SignalState = 0
	Signal_BROKEN        Signal_SignalState = 1
	Signal_BLACK         Signal_SignalState = 2
	Signal_RED           Signal_SignalState = 3
	Signal_YELLOW        Signal_SignalState = 4
	Signal_DOUBLE_YELLOW Signal_SignalState = 5
	Signal_GREEN         Signal_SignalState = 6
	Signal_BLUE          Signal_SignalState = 7
	Signal_WHITE         Signal_SignalState = 8
)

// Enum value maps for Signal_SignalState.
var (
	Signal_SignalState_name = map[int32]string{
		0: "UNKNOWN",
		1: "BROKEN",
		2: "BLACK",
		3: "RED",
		4: "YELLOW",
		5: "DOUBLE_YELLOW",
		6: "GREEN",
		7: "BLUE",
		8: "WHITE",
	}
	Signal_SignalState_value = map[string]int32{
		"UNKNOWN":       0,
		"BROKEN":        1,
		"BLACK":         2,
		"RED":           3,
		"YELLOW":        4,
		"DOUBLE_YELLOW": 5,
		"GREEN":         6,
		"BLUE":          7,
		"WHITE":         8,
	}
)

func (x Signal_SignalState) Enum() *Signal_SignalState {
	p := new(Signal_SignalState)
	*p = x
	return p
}

func (x Signal_SignalState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Signal_SignalState) Descriptor() protoreflect.EnumDescriptor {
	return file_signal_message_proto_enumTypes[0].Descriptor()
}

func (Signal_SignalState) Type() protoreflect.EnumType {
	return &file_signal_message_proto_enumTypes[0]
}

func (x Signal_SignalState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Signal_SignalState.Descriptor instead.
func (Signal_SignalState) EnumDescriptor() ([]byte, []int) {
	return file_signal_message_proto_rawDescGZIP(), []int{0, 0}
}

type Signal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State Signal_SignalState `protobuf:"varint,2,opt,name=state,proto3,enum=prac.net.Signal_SignalState" json:"state,omitempty"`
}

func (x *Signal) Reset() {
	*x = Signal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signal_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signal) ProtoMessage() {}

func (x *Signal) ProtoReflect() protoreflect.Message {
	mi := &file_signal_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signal.ProtoReflect.Descriptor instead.
func (*Signal) Descriptor() ([]byte, []int) {
	return file_signal_message_proto_rawDescGZIP(), []int{0}
}

func (x *Signal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Signal) GetState() Signal_SignalState {
	if x != nil {
		return x.State
	}
	return Signal_UNKNOWN
}

var File_signal_message_proto protoreflect.FileDescriptor

var file_signal_message_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x61, 0x63, 0x2e, 0x6e, 0x65, 0x74,
	0x22, 0xc7, 0x01, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x61,
	0x63, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22,
	0x79, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x42,
	0x52, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x4c, 0x41, 0x43, 0x4b,
	0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x59,
	0x45, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x4f, 0x55, 0x42, 0x4c,
	0x45, 0x5f, 0x59, 0x45, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52,
	0x45, 0x45, 0x4e, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4c, 0x55, 0x45, 0x10, 0x07, 0x12,
	0x09, 0x0a, 0x05, 0x57, 0x48, 0x49, 0x54, 0x45, 0x10, 0x08, 0x42, 0x24, 0x0a, 0x1a, 0x64, 0x65,
	0x76, 0x2e, 0x67, 0x6c, 0x79, 0x63, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x62, 0x50, 0x01, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_signal_message_proto_rawDescOnce sync.Once
	file_signal_message_proto_rawDescData = file_signal_message_proto_rawDesc
)

func file_signal_message_proto_rawDescGZIP() []byte {
	file_signal_message_proto_rawDescOnce.Do(func() {
		file_signal_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_signal_message_proto_rawDescData)
	})
	return file_signal_message_proto_rawDescData
}

var file_signal_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_signal_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_signal_message_proto_goTypes = []interface{}{
	(Signal_SignalState)(0), // 0: prac.net.Signal.SignalState
	(*Signal)(nil),          // 1: prac.net.Signal
}
var file_signal_message_proto_depIdxs = []int32{
	0, // 0: prac.net.Signal.state:type_name -> prac.net.Signal.SignalState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_signal_message_proto_init() }
func file_signal_message_proto_init() {
	if File_signal_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_signal_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signal); i {
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
			RawDescriptor: file_signal_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_signal_message_proto_goTypes,
		DependencyIndexes: file_signal_message_proto_depIdxs,
		EnumInfos:         file_signal_message_proto_enumTypes,
		MessageInfos:      file_signal_message_proto_msgTypes,
	}.Build()
	File_signal_message_proto = out.File
	file_signal_message_proto_rawDesc = nil
	file_signal_message_proto_goTypes = nil
	file_signal_message_proto_depIdxs = nil
}