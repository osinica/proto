// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: proto/meal.proto

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

type Time int32

const (
	Time_SHORT  Time = 0
	Time_MEDIUM Time = 1
	Time_LONG   Time = 2
)

// Enum value maps for Time.
var (
	Time_name = map[int32]string{
		0: "SHORT",
		1: "MEDIUM",
		2: "LONG",
	}
	Time_value = map[string]int32{
		"SHORT":  0,
		"MEDIUM": 1,
		"LONG":   2,
	}
)

func (x Time) Enum() *Time {
	p := new(Time)
	*p = x
	return p
}

func (x Time) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Time) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_meal_proto_enumTypes[0].Descriptor()
}

func (Time) Type() protoreflect.EnumType {
	return &file_proto_meal_proto_enumTypes[0]
}

func (x Time) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Time.Descriptor instead.
func (Time) EnumDescriptor() ([]byte, []int) {
	return file_proto_meal_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_SUCCESS Status = 0
	Status_ERROR   Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "SUCCESS",
		1: "ERROR",
	}
	Status_value = map[string]int32{
		"SUCCESS": 0,
		"ERROR":   1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_meal_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_proto_meal_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_meal_proto_rawDescGZIP(), []int{1}
}

type Meal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Duration Time   `protobuf:"varint,3,opt,name=duration,proto3,enum=meal.Time" json:"duration,omitempty"`
}

func (x *Meal) Reset() {
	*x = Meal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meal) ProtoMessage() {}

func (x *Meal) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meal.ProtoReflect.Descriptor instead.
func (*Meal) Descriptor() ([]byte, []int) {
	return file_proto_meal_proto_rawDescGZIP(), []int{0}
}

func (x *Meal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Meal) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Meal) GetDuration() Time {
	if x != nil {
		return x.Duration
	}
	return Time_SHORT
}

// create
type CreateMealRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Duration Time   `protobuf:"varint,2,opt,name=duration,proto3,enum=meal.Time" json:"duration,omitempty"`
}

func (x *CreateMealRequest) Reset() {
	*x = CreateMealRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMealRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMealRequest) ProtoMessage() {}

func (x *CreateMealRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMealRequest.ProtoReflect.Descriptor instead.
func (*CreateMealRequest) Descriptor() ([]byte, []int) {
	return file_proto_meal_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMealRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMealRequest) GetDuration() Time {
	if x != nil {
		return x.Duration
	}
	return Time_SHORT
}

type CreateMealResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=meal.Status" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   *Meal  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateMealResponse) Reset() {
	*x = CreateMealResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMealResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMealResponse) ProtoMessage() {}

func (x *CreateMealResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMealResponse.ProtoReflect.Descriptor instead.
func (*CreateMealResponse) Descriptor() ([]byte, []int) {
	return file_proto_meal_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMealResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_SUCCESS
}

func (x *CreateMealResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateMealResponse) GetData() *Meal {
	if x != nil {
		return x.Data
	}
	return nil
}

// find one
type FindOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindOneRequest) Reset() {
	*x = FindOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneRequest) ProtoMessage() {}

func (x *FindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneRequest.ProtoReflect.Descriptor instead.
func (*FindOneRequest) Descriptor() ([]byte, []int) {
	return file_proto_meal_proto_rawDescGZIP(), []int{3}
}

func (x *FindOneRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=meal.Status" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   *Meal  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FindOneResponse) Reset() {
	*x = FindOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneResponse) ProtoMessage() {}

func (x *FindOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneResponse.ProtoReflect.Descriptor instead.
func (*FindOneResponse) Descriptor() ([]byte, []int) {
	return file_proto_meal_proto_rawDescGZIP(), []int{4}
}

func (x *FindOneResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_SUCCESS
}

func (x *FindOneResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindOneResponse) GetData() *Meal {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_meal_proto protoreflect.FileDescriptor

var file_proto_meal_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6d, 0x65, 0x61, 0x6c, 0x22, 0x52, 0x0a, 0x04, 0x4d, 0x65, 0x61, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x70, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x6d, 0x65, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x61, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x20, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x6d, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x61, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x2a, 0x27, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x48, 0x4f, 0x52,
	0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x02, 0x2a, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x32, 0x8a, 0x01, 0x0a, 0x0b,
	0x4d, 0x65, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x61, 0x6c, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x61, 0x6c,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x65, 0x61, 0x6c,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_meal_proto_rawDescOnce sync.Once
	file_proto_meal_proto_rawDescData = file_proto_meal_proto_rawDesc
)

func file_proto_meal_proto_rawDescGZIP() []byte {
	file_proto_meal_proto_rawDescOnce.Do(func() {
		file_proto_meal_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_meal_proto_rawDescData)
	})
	return file_proto_meal_proto_rawDescData
}

var file_proto_meal_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_meal_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_meal_proto_goTypes = []interface{}{
	(Time)(0),                  // 0: meal.Time
	(Status)(0),                // 1: meal.Status
	(*Meal)(nil),               // 2: meal.Meal
	(*CreateMealRequest)(nil),  // 3: meal.CreateMealRequest
	(*CreateMealResponse)(nil), // 4: meal.CreateMealResponse
	(*FindOneRequest)(nil),     // 5: meal.FindOneRequest
	(*FindOneResponse)(nil),    // 6: meal.FindOneResponse
}
var file_proto_meal_proto_depIdxs = []int32{
	0, // 0: meal.Meal.duration:type_name -> meal.Time
	0, // 1: meal.CreateMealRequest.duration:type_name -> meal.Time
	1, // 2: meal.CreateMealResponse.status:type_name -> meal.Status
	2, // 3: meal.CreateMealResponse.data:type_name -> meal.Meal
	1, // 4: meal.FindOneResponse.status:type_name -> meal.Status
	2, // 5: meal.FindOneResponse.data:type_name -> meal.Meal
	3, // 6: meal.MealService.CreateMeal:input_type -> meal.CreateMealRequest
	5, // 7: meal.MealService.FindOne:input_type -> meal.FindOneRequest
	4, // 8: meal.MealService.CreateMeal:output_type -> meal.CreateMealResponse
	6, // 9: meal.MealService.FindOne:output_type -> meal.FindOneResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_meal_proto_init() }
func file_proto_meal_proto_init() {
	if File_proto_meal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_meal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meal); i {
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
		file_proto_meal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMealRequest); i {
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
		file_proto_meal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMealResponse); i {
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
		file_proto_meal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneRequest); i {
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
		file_proto_meal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneResponse); i {
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
			RawDescriptor: file_proto_meal_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_meal_proto_goTypes,
		DependencyIndexes: file_proto_meal_proto_depIdxs,
		EnumInfos:         file_proto_meal_proto_enumTypes,
		MessageInfos:      file_proto_meal_proto_msgTypes,
	}.Build()
	File_proto_meal_proto = out.File
	file_proto_meal_proto_rawDesc = nil
	file_proto_meal_proto_goTypes = nil
	file_proto_meal_proto_depIdxs = nil
}