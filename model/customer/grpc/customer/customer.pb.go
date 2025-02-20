// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: model/customer/grpc/customer/customer.proto

package customer

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

type Status int32

const (
	Status_Unknown  Status = 0
	Status_Active   Status = 1
	Status_Inactive Status = 2
	Status_Banned   Status = 3
	Status_Expired  Status = 4
	Status_Deceased Status = 5
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Unknown",
		1: "Active",
		2: "Inactive",
		3: "Banned",
		4: "Expired",
		5: "Deceased",
	}
	Status_value = map[string]int32{
		"Unknown":  0,
		"Active":   1,
		"Inactive": 2,
		"Banned":   3,
		"Expired":  4,
		"Deceased": 5,
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
	return file_model_customer_grpc_customer_customer_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_model_customer_grpc_customer_customer_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_model_customer_grpc_customer_customer_proto_rawDescGZIP(), []int{0}
}

type Type int32

const (
	Type_none            Type = 0
	Type_Card            Type = 1
	Type_Passport        Type = 2
	Type_GovernmentPaper Type = 3
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "none",
		1: "Card",
		2: "Passport",
		3: "GovernmentPaper",
	}
	Type_value = map[string]int32{
		"none":            0,
		"Card":            1,
		"Passport":        2,
		"GovernmentPaper": 3,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_model_customer_grpc_customer_customer_proto_enumTypes[1].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_model_customer_grpc_customer_customer_proto_enumTypes[1]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_model_customer_grpc_customer_customer_proto_rawDescGZIP(), []int{1}
}

type Customer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    string                 `protobuf:"bytes,1,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	IdType        Type                   `protobuf:"varint,2,opt,name=IdType,proto3,enum=customer.Type" json:"IdType,omitempty"`
	UserName      string                 `protobuf:"bytes,3,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password      string                 `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	Email         string                 `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Mobile        string                 `protobuf:"bytes,6,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	FirstName     string                 `protobuf:"bytes,7,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	MiddleName    string                 `protobuf:"bytes,8,opt,name=MiddleName,proto3" json:"MiddleName,omitempty"`
	LastName      string                 `protobuf:"bytes,9,opt,name=LastName,proto3" json:"LastName,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Status        Status                 `protobuf:"varint,12,opt,name=Status,proto3,enum=customer.Status" json:"Status,omitempty"`
	Picture       []byte                 `protobuf:"bytes,13,opt,name=Picture,proto3" json:"Picture,omitempty"`
	Document      []byte                 `protobuf:"bytes,14,opt,name=Document,proto3" json:"Document,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Customer) Reset() {
	*x = Customer{}
	mi := &file_model_customer_grpc_customer_customer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_model_customer_grpc_customer_customer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_model_customer_grpc_customer_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Customer) GetIdType() Type {
	if x != nil {
		return x.IdType
	}
	return Type_none
}

func (x *Customer) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Customer) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Customer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Customer) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *Customer) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Customer) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *Customer) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Customer) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Customer) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Customer) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Unknown
}

func (x *Customer) GetPicture() []byte {
	if x != nil {
		return x.Picture
	}
	return nil
}

func (x *Customer) GetDocument() []byte {
	if x != nil {
		return x.Document
	}
	return nil
}

var File_model_customer_grpc_customer_customer_proto protoreflect.FileDescriptor

var file_model_customer_grpc_customer_customer_proto_rawDesc = string([]byte{
	0x0a, 0x2b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x03, 0x0a, 0x08, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x28, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x50,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x50, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x2a, 0x56, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x10, 0x03, 0x12, 0x0b,
	0x0a, 0x07, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x65, 0x63, 0x65, 0x61, 0x73, 0x65, 0x64, 0x10, 0x05, 0x2a, 0x3d, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x43,
	0x61, 0x72, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x70, 0x65, 0x72, 0x10, 0x03, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x68, 0x74, 0x68, 0x72, 0x68, 0x2f, 0x47, 0x6f,
	0x4e, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_model_customer_grpc_customer_customer_proto_rawDescOnce sync.Once
	file_model_customer_grpc_customer_customer_proto_rawDescData []byte
)

func file_model_customer_grpc_customer_customer_proto_rawDescGZIP() []byte {
	file_model_customer_grpc_customer_customer_proto_rawDescOnce.Do(func() {
		file_model_customer_grpc_customer_customer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_model_customer_grpc_customer_customer_proto_rawDesc), len(file_model_customer_grpc_customer_customer_proto_rawDesc)))
	})
	return file_model_customer_grpc_customer_customer_proto_rawDescData
}

var file_model_customer_grpc_customer_customer_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_model_customer_grpc_customer_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_customer_grpc_customer_customer_proto_goTypes = []any{
	(Status)(0),                   // 0: customer.Status
	(Type)(0),                     // 1: customer.Type
	(*Customer)(nil),              // 2: customer.Customer
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_model_customer_grpc_customer_customer_proto_depIdxs = []int32{
	1, // 0: customer.Customer.IdType:type_name -> customer.Type
	3, // 1: customer.Customer.CreatedAt:type_name -> google.protobuf.Timestamp
	3, // 2: customer.Customer.UpdatedAt:type_name -> google.protobuf.Timestamp
	0, // 3: customer.Customer.Status:type_name -> customer.Status
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_model_customer_grpc_customer_customer_proto_init() }
func file_model_customer_grpc_customer_customer_proto_init() {
	if File_model_customer_grpc_customer_customer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_model_customer_grpc_customer_customer_proto_rawDesc), len(file_model_customer_grpc_customer_customer_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_customer_grpc_customer_customer_proto_goTypes,
		DependencyIndexes: file_model_customer_grpc_customer_customer_proto_depIdxs,
		EnumInfos:         file_model_customer_grpc_customer_customer_proto_enumTypes,
		MessageInfos:      file_model_customer_grpc_customer_customer_proto_msgTypes,
	}.Build()
	File_model_customer_grpc_customer_customer_proto = out.File
	file_model_customer_grpc_customer_customer_proto_goTypes = nil
	file_model_customer_grpc_customer_customer_proto_depIdxs = nil
}
