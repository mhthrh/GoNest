// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: model/customer/customer.proto

package customer

import (
	address "github.com/mhthrh/GoNest/model/address"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	gofeaturespb "google.golang.org/protobuf/types/gofeaturespb"
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

type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       *address.Address       `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Customer      *Customer              `protobuf:"bytes,2,opt,name=Customer,proto3" json:"Customer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_model_customer_customer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_model_customer_customer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_model_customer_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetAddress() *address.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Request) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Customer      *Customer              `protobuf:"bytes,1,opt,name=Customer,proto3" json:"Customer,omitempty"`
	Error         *gofeaturespb.Error    `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_model_customer_customer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_model_customer_customer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_model_customer_customer_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *Response) GetError() *gofeaturespb.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Customer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    string                 `protobuf:"bytes,1,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	IdType        int32                  `protobuf:"varint,2,opt,name=IdType,proto3" json:"IdType,omitempty"`
	UserName      string                 `protobuf:"bytes,3,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password      string                 `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	Email         string                 `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Mobile        string                 `protobuf:"bytes,6,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	Address       *address.Address       `protobuf:"bytes,7,opt,name=Address,proto3" json:"Address,omitempty"`
	FirstName     string                 `protobuf:"bytes,8,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	MiddleName    string                 `protobuf:"bytes,9,opt,name=MiddleName,proto3" json:"MiddleName,omitempty"`
	LastName      string                 `protobuf:"bytes,10,opt,name=LastName,proto3" json:"LastName,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Status        int32                  `protobuf:"varint,13,opt,name=Status,proto3" json:"Status,omitempty"`
	Picture       []byte                 `protobuf:"bytes,14,opt,name=Picture,proto3" json:"Picture,omitempty"`
	Document      []byte                 `protobuf:"bytes,15,opt,name=Document,proto3" json:"Document,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Customer) Reset() {
	*x = Customer{}
	mi := &file_model_customer_customer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_model_customer_customer_proto_msgTypes[2]
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
	return file_model_customer_customer_proto_rawDescGZIP(), []int{2}
}

func (x *Customer) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Customer) GetIdType() int32 {
	if x != nil {
		return x.IdType
	}
	return 0
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

func (x *Customer) GetAddress() *address.Address {
	if x != nil {
		return x.Address
	}
	return nil
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

func (x *Customer) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
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

var File_model_customer_customer_proto protoreflect.FileDescriptor

var file_model_customer_customer_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x65, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x5e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xf0, 0x03, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x46, 0x0a, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x68, 0x74, 0x68, 0x72, 0x68, 0x2f, 0x47, 0x6f, 0x4e, 0x65, 0x73, 0x74, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_model_customer_customer_proto_rawDescOnce sync.Once
	file_model_customer_customer_proto_rawDescData []byte
)

func file_model_customer_customer_proto_rawDescGZIP() []byte {
	file_model_customer_customer_proto_rawDescOnce.Do(func() {
		file_model_customer_customer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_model_customer_customer_proto_rawDesc), len(file_model_customer_customer_proto_rawDesc)))
	})
	return file_model_customer_customer_proto_rawDescData
}

var file_model_customer_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_model_customer_customer_proto_goTypes = []any{
	(*Request)(nil),               // 0: customer.Request
	(*Response)(nil),              // 1: customer.Response
	(*Customer)(nil),              // 2: customer.customer
	(*address.Address)(nil),       // 3: address.Address
	(*gofeaturespb.Error)(nil),    // 4: error.Error
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_model_customer_customer_proto_depIdxs = []int32{
	3, // 0: customer.Request.Address:type_name -> address.Address
	2, // 1: customer.Request.Customer:type_name -> customer.customer
	2, // 2: customer.Response.Customer:type_name -> customer.customer
	4, // 3: customer.Response.Error:type_name -> error.Error
	3, // 4: customer.customer.Address:type_name -> address.Address
	5, // 5: customer.customer.CreatedAt:type_name -> google.protobuf.Timestamp
	5, // 6: customer.customer.UpdatedAt:type_name -> google.protobuf.Timestamp
	0, // 7: customer.Service.RegisterCustomer:input_type -> customer.Request
	1, // 8: customer.Service.RegisterCustomer:output_type -> customer.Response
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_model_customer_customer_proto_init() }
func file_model_customer_customer_proto_init() {
	if File_model_customer_customer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_model_customer_customer_proto_rawDesc), len(file_model_customer_customer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_model_customer_customer_proto_goTypes,
		DependencyIndexes: file_model_customer_customer_proto_depIdxs,
		MessageInfos:      file_model_customer_customer_proto_msgTypes,
	}.Build()
	File_model_customer_customer_proto = out.File
	file_model_customer_customer_proto_goTypes = nil
	file_model_customer_customer_proto_depIdxs = nil
}
