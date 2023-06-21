// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v4.23.2
// source: wokerd.proto

package entities

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

type Worker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID          string `protobuf:"bytes,1,opt,name=UID,proto3" json:"UID,omitempty"`                   // Unique id of the worker
	ExternalPath string `protobuf:"bytes,2,opt,name=ExternalPath,proto3" json:"ExternalPath,omitempty"` // External path of the worker, default is '/'
	HostName     string `protobuf:"bytes,3,opt,name=HostName,proto3" json:"HostName,omitempty"`         // the workerd runner host name, default is 'localhost'
	NodeName     string `protobuf:"bytes,4,opt,name=NodeName,proto3" json:"NodeName,omitempty"`         // for future HA feature, default is 'default'
	Port         int32  `protobuf:"varint,5,opt,name=Port,proto3" json:"Port,omitempty"`                // worker's port, platfrom will obtain free port while init worker
	Entry        string `protobuf:"bytes,6,opt,name=Entry,proto3" json:"Entry,omitempty"`               // worker's entry file, default is 'entry.js'
	Code         []byte `protobuf:"bytes,7,opt,name=Code,proto3" json:"Code,omitempty"`                 // worker's code
	Name         string `protobuf:"bytes,8,opt,name=Name,proto3" json:"Name,omitempty"`                 // worker's name, also use at worker routing, must be unique, default is UID
	TunnelID     string `protobuf:"bytes,9,opt,name=TunnelID,proto3" json:"TunnelID,omitempty"`         // worker's tunnel id
	UserID       uint64 `protobuf:"varint,10,opt,name=UserID,proto3" json:"UserID,omitempty"`           // worker's user id
}

func (x *Worker) Reset() {
	*x = Worker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wokerd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Worker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Worker) ProtoMessage() {}

func (x *Worker) ProtoReflect() protoreflect.Message {
	mi := &file_wokerd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Worker.ProtoReflect.Descriptor instead.
func (*Worker) Descriptor() ([]byte, []int) {
	return file_wokerd_proto_rawDescGZIP(), []int{0}
}

func (x *Worker) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

func (x *Worker) GetExternalPath() string {
	if x != nil {
		return x.ExternalPath
	}
	return ""
}

func (x *Worker) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *Worker) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *Worker) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Worker) GetEntry() string {
	if x != nil {
		return x.Entry
	}
	return ""
}

func (x *Worker) GetCode() []byte {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *Worker) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Worker) GetTunnelID() string {
	if x != nil {
		return x.TunnelID
	}
	return ""
}

func (x *Worker) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

// one WorkerList for one workerd instance
type WorkerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfName string    `protobuf:"bytes,1,opt,name=ConfName,proto3" json:"ConfName,omitempty"` // the name of the workerd instance
	Workers  []*Worker `protobuf:"bytes,2,rep,name=Workers,proto3" json:"Workers,omitempty"`
	NodeName string    `protobuf:"bytes,3,opt,name=NodeName,proto3" json:"NodeName,omitempty"` // workerd runner host name, for HA
}

func (x *WorkerList) Reset() {
	*x = WorkerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wokerd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerList) ProtoMessage() {}

func (x *WorkerList) ProtoReflect() protoreflect.Message {
	mi := &file_wokerd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerList.ProtoReflect.Descriptor instead.
func (*WorkerList) Descriptor() ([]byte, []int) {
	return file_wokerd_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerList) GetConfName() string {
	if x != nil {
		return x.ConfName
	}
	return ""
}

func (x *WorkerList) GetWorkers() []*Worker {
	if x != nil {
		return x.Workers
	}
	return nil
}

func (x *WorkerList) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

var File_wokerd_proto protoreflect.FileDescriptor

var file_wokerd_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x6f, 0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x64, 0x65, 0x66, 0x73, 0x22, 0xfc, 0x01, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x49,
	0x44, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x22, 0x6c, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x07, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x64, 0x65, 0x66, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x07, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wokerd_proto_rawDescOnce sync.Once
	file_wokerd_proto_rawDescData = file_wokerd_proto_rawDesc
)

func file_wokerd_proto_rawDescGZIP() []byte {
	file_wokerd_proto_rawDescOnce.Do(func() {
		file_wokerd_proto_rawDescData = protoimpl.X.CompressGZIP(file_wokerd_proto_rawDescData)
	})
	return file_wokerd_proto_rawDescData
}

var file_wokerd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wokerd_proto_goTypes = []interface{}{
	(*Worker)(nil),     // 0: defs.Worker
	(*WorkerList)(nil), // 1: defs.WorkerList
}
var file_wokerd_proto_depIdxs = []int32{
	0, // 0: defs.WorkerList.Workers:type_name -> defs.Worker
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wokerd_proto_init() }
func file_wokerd_proto_init() {
	if File_wokerd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wokerd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Worker); i {
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
		file_wokerd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerList); i {
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
			RawDescriptor: file_wokerd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wokerd_proto_goTypes,
		DependencyIndexes: file_wokerd_proto_depIdxs,
		MessageInfos:      file_wokerd_proto_msgTypes,
	}.Build()
	File_wokerd_proto = out.File
	file_wokerd_proto_rawDesc = nil
	file_wokerd_proto_goTypes = nil
	file_wokerd_proto_depIdxs = nil
}
