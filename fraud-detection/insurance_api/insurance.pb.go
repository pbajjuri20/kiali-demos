// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: insurance.proto

package insurance_api

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

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_insurance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_insurance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_insurance_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type InsuranceReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId  string  `protobuf:"bytes,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	Risk        float64 `protobuf:"fixed64,2,opt,name=risk,proto3" json:"risk,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *InsuranceReport) Reset() {
	*x = InsuranceReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_insurance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsuranceReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsuranceReport) ProtoMessage() {}

func (x *InsuranceReport) ProtoReflect() protoreflect.Message {
	mi := &file_insurance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsuranceReport.ProtoReflect.Descriptor instead.
func (*InsuranceReport) Descriptor() ([]byte, []int) {
	return file_insurance_proto_rawDescGZIP(), []int{1}
}

func (x *InsuranceReport) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *InsuranceReport) GetRisk() float64 {
	if x != nil {
		return x.Risk
	}
	return 0
}

func (x *InsuranceReport) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_insurance_proto protoreflect.FileDescriptor

var file_insurance_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x22, 0x2a, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x0f,
	0x49, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x69, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x72,
	0x69, 0x73, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x5c, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e,
	0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x46, 0x72, 0x61,
	0x75, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x69, 0x6e, 0x73, 0x75, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x49, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x69, 0x61, 0x6c, 0x69, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x66, 0x72,
	0x61, 0x75, 0x64, 0x2d, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e,
	0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_insurance_proto_rawDescOnce sync.Once
	file_insurance_proto_rawDescData = file_insurance_proto_rawDesc
)

func file_insurance_proto_rawDescGZIP() []byte {
	file_insurance_proto_rawDescOnce.Do(func() {
		file_insurance_proto_rawDescData = protoimpl.X.CompressGZIP(file_insurance_proto_rawDescData)
	})
	return file_insurance_proto_rawDescData
}

var file_insurance_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_insurance_proto_goTypes = []interface{}{
	(*Customer)(nil),        // 0: insurance_api.Customer
	(*InsuranceReport)(nil), // 1: insurance_api.InsuranceReport
}
var file_insurance_proto_depIdxs = []int32{
	0, // 0: insurance_api.InsuranceService.FraudReport:input_type -> insurance_api.Customer
	1, // 1: insurance_api.InsuranceService.FraudReport:output_type -> insurance_api.InsuranceReport
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_insurance_proto_init() }
func file_insurance_proto_init() {
	if File_insurance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_insurance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_insurance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsuranceReport); i {
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
			RawDescriptor: file_insurance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_insurance_proto_goTypes,
		DependencyIndexes: file_insurance_proto_depIdxs,
		MessageInfos:      file_insurance_proto_msgTypes,
	}.Build()
	File_insurance_proto = out.File
	file_insurance_proto_rawDesc = nil
	file_insurance_proto_goTypes = nil
	file_insurance_proto_depIdxs = nil
}
