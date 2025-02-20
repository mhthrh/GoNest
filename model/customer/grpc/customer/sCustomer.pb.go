// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: model/customer/grpc/customer/sCustomer.proto

package customer

import (
	address "github.com/mhthrh/GoNest/model/address/grpc/address"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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
	mi := &file_model_customer_grpc_customer_sCustomer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_model_customer_grpc_customer_sCustomer_proto_msgTypes[0]
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
	return file_model_customer_grpc_customer_sCustomer_proto_rawDescGZIP(), []int{0}
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
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_model_customer_grpc_customer_sCustomer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_model_customer_grpc_customer_sCustomer_proto_msgTypes[1]
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
	return file_model_customer_grpc_customer_sCustomer_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

var File_model_customer_grpc_customer_sCustomer_proto protoreflect.FileDescriptor

var file_model_customer_grpc_customer_sCustomer_proto_rawDesc = string([]byte{
	0x0a, 0x2c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x73,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x28, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x65, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x3a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x32, 0x46, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a,
	0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x11, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x68, 0x74, 0x68, 0x72, 0x68, 0x2f,
	0x47, 0x6f, 0x4e, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_model_customer_grpc_customer_sCustomer_proto_rawDescOnce sync.Once
	file_model_customer_grpc_customer_sCustomer_proto_rawDescData []byte
)

func file_model_customer_grpc_customer_sCustomer_proto_rawDescGZIP() []byte {
	file_model_customer_grpc_customer_sCustomer_proto_rawDescOnce.Do(func() {
		file_model_customer_grpc_customer_sCustomer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_model_customer_grpc_customer_sCustomer_proto_rawDesc), len(file_model_customer_grpc_customer_sCustomer_proto_rawDesc)))
	})
	return file_model_customer_grpc_customer_sCustomer_proto_rawDescData
}

var file_model_customer_grpc_customer_sCustomer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_customer_grpc_customer_sCustomer_proto_goTypes = []any{
	(*Request)(nil),         // 0: customer.Request
	(*Response)(nil),        // 1: customer.Response
	(*address.Address)(nil), // 2: address.Address
	(*Customer)(nil),        // 3: customer.Customer
}
var file_model_customer_grpc_customer_sCustomer_proto_depIdxs = []int32{
	2, // 0: customer.Request.Address:type_name -> address.Address
	3, // 1: customer.Request.Customer:type_name -> customer.Customer
	3, // 2: customer.Response.Customer:type_name -> customer.Customer
	0, // 3: customer.Service.RegisterCustomer:input_type -> customer.Request
	1, // 4: customer.Service.RegisterCustomer:output_type -> customer.Response
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_model_customer_grpc_customer_sCustomer_proto_init() }
func file_model_customer_grpc_customer_sCustomer_proto_init() {
	if File_model_customer_grpc_customer_sCustomer_proto != nil {
		return
	}
	file_model_customer_grpc_customer_customer_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_model_customer_grpc_customer_sCustomer_proto_rawDesc), len(file_model_customer_grpc_customer_sCustomer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_model_customer_grpc_customer_sCustomer_proto_goTypes,
		DependencyIndexes: file_model_customer_grpc_customer_sCustomer_proto_depIdxs,
		MessageInfos:      file_model_customer_grpc_customer_sCustomer_proto_msgTypes,
	}.Build()
	File_model_customer_grpc_customer_sCustomer_proto = out.File
	file_model_customer_grpc_customer_sCustomer_proto_goTypes = nil
	file_model_customer_grpc_customer_sCustomer_proto_depIdxs = nil
}
