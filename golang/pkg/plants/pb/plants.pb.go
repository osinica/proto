// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: proto/plants.proto

package pb

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_plants_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_plants_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_plants_proto_rawDescGZIP(), []int{0}
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PlantType   string  `protobuf:"bytes,2,opt,name=plant_type,json=plantType,proto3" json:"plant_type,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Lat         float32 `protobuf:"fixed32,4,opt,name=lat,proto3" json:"lat,omitempty"`
	Long        float32 `protobuf:"fixed32,5,opt,name=long,proto3" json:"long,omitempty"`
	LastChecked string  `protobuf:"bytes,6,opt,name=last_checked,json=lastChecked,proto3" json:"last_checked,omitempty"`
	Added       string  `protobuf:"bytes,7,opt,name=added,proto3" json:"added,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_plants_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_proto_plants_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_proto_plants_proto_rawDescGZIP(), []int{1}
}

func (x *Object) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Object) GetPlantType() string {
	if x != nil {
		return x.PlantType
	}
	return ""
}

func (x *Object) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Object) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Object) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *Object) GetLastChecked() string {
	if x != nil {
		return x.LastChecked
	}
	return ""
}

func (x *Object) GetAdded() string {
	if x != nil {
		return x.Added
	}
	return ""
}

type Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Type) Reset() {
	*x = Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_plants_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_proto_plants_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Type.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_proto_plants_proto_rawDescGZIP(), []int{2}
}

func (x *Type) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Type) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ObjectList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objects []*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *ObjectList) Reset() {
	*x = ObjectList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_plants_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectList) ProtoMessage() {}

func (x *ObjectList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_plants_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectList.ProtoReflect.Descriptor instead.
func (*ObjectList) Descriptor() ([]byte, []int) {
	return file_proto_plants_proto_rawDescGZIP(), []int{3}
}

func (x *ObjectList) GetObjects() []*Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

type TypeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Types []*Type `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
}

func (x *TypeList) Reset() {
	*x = TypeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_plants_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeList) ProtoMessage() {}

func (x *TypeList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_plants_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeList.ProtoReflect.Descriptor instead.
func (*TypeList) Descriptor() ([]byte, []int) {
	return file_proto_plants_proto_rawDescGZIP(), []int{4}
}

func (x *TypeList) GetTypes() []*Type {
	if x != nil {
		return x.Types
	}
	return nil
}

type RequestId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestId) Reset() {
	*x = RequestId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_plants_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestId) ProtoMessage() {}

func (x *RequestId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_plants_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestId.ProtoReflect.Descriptor instead.
func (*RequestId) Descriptor() ([]byte, []int) {
	return file_proto_plants_proto_rawDescGZIP(), []int{5}
}

func (x *RequestId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_plants_proto protoreflect.FileDescriptor

var file_proto_plants_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xb8, 0x01, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03,
	0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64,
	0x64, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64,
	0x22, 0x2a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x0a,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x32, 0x3c, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42,
	0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_plants_proto_rawDescOnce sync.Once
	file_proto_plants_proto_rawDescData = file_proto_plants_proto_rawDesc
)

func file_proto_plants_proto_rawDescGZIP() []byte {
	file_proto_plants_proto_rawDescOnce.Do(func() {
		file_proto_plants_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_plants_proto_rawDescData)
	})
	return file_proto_plants_proto_rawDescData
}

var file_proto_plants_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_plants_proto_goTypes = []interface{}{
	(*Empty)(nil),      // 0: plants.Empty
	(*Object)(nil),     // 1: plants.Object
	(*Type)(nil),       // 2: plants.Type
	(*ObjectList)(nil), // 3: plants.ObjectList
	(*TypeList)(nil),   // 4: plants.TypeList
	(*RequestId)(nil),  // 5: plants.RequestId
}
var file_proto_plants_proto_depIdxs = []int32{
	1, // 0: plants.ObjectList.objects:type_name -> plants.Object
	2, // 1: plants.TypeList.types:type_name -> plants.Type
	0, // 2: plants.PlantsService.List:input_type -> plants.Empty
	3, // 3: plants.PlantsService.List:output_type -> plants.ObjectList
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_plants_proto_init() }
func file_proto_plants_proto_init() {
	if File_proto_plants_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_plants_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_plants_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_proto_plants_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Type); i {
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
		file_proto_plants_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectList); i {
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
		file_proto_plants_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeList); i {
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
		file_proto_plants_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestId); i {
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
			RawDescriptor: file_proto_plants_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_plants_proto_goTypes,
		DependencyIndexes: file_proto_plants_proto_depIdxs,
		MessageInfos:      file_proto_plants_proto_msgTypes,
	}.Build()
	File_proto_plants_proto = out.File
	file_proto_plants_proto_rawDesc = nil
	file_proto_plants_proto_goTypes = nil
	file_proto_plants_proto_depIdxs = nil
}