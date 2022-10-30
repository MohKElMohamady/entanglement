// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: proto/physicist.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AllPhysicistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AllPhysicistsRequest) Reset() {
	*x = AllPhysicistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_physicist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllPhysicistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllPhysicistsRequest) ProtoMessage() {}

func (x *AllPhysicistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_physicist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllPhysicistsRequest.ProtoReflect.Descriptor instead.
func (*AllPhysicistsRequest) Descriptor() ([]byte, []int) {
	return file_proto_physicist_proto_rawDescGZIP(), []int{0}
}

type Physicist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhysicistId    *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=physicistId,proto3" json:"physicistId,omitempty"`
	FirstName      string                  `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName       string                  `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	CountryOfBirth string                  `protobuf:"bytes,4,opt,name=countryOfBirth,proto3" json:"countryOfBirth,omitempty"`
	DateOfBirth    int64                   `protobuf:"varint,5,opt,name=dateOfBirth,proto3" json:"dateOfBirth,omitempty"`
	DateOfDeath    int64                   `protobuf:"varint,6,opt,name=dateOfDeath,proto3" json:"dateOfDeath,omitempty"`
	Biography      string                  `protobuf:"bytes,7,opt,name=biography,proto3" json:"biography,omitempty"`
}

func (x *Physicist) Reset() {
	*x = Physicist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_physicist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Physicist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Physicist) ProtoMessage() {}

func (x *Physicist) ProtoReflect() protoreflect.Message {
	mi := &file_proto_physicist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Physicist.ProtoReflect.Descriptor instead.
func (*Physicist) Descriptor() ([]byte, []int) {
	return file_proto_physicist_proto_rawDescGZIP(), []int{1}
}

func (x *Physicist) GetPhysicistId() *wrapperspb.StringValue {
	if x != nil {
		return x.PhysicistId
	}
	return nil
}

func (x *Physicist) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Physicist) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Physicist) GetCountryOfBirth() string {
	if x != nil {
		return x.CountryOfBirth
	}
	return ""
}

func (x *Physicist) GetDateOfBirth() int64 {
	if x != nil {
		return x.DateOfBirth
	}
	return 0
}

func (x *Physicist) GetDateOfDeath() int64 {
	if x != nil {
		return x.DateOfDeath
	}
	return 0
}

func (x *Physicist) GetBiography() string {
	if x != nil {
		return x.Biography
	}
	return ""
}

type AllPhysicists struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Physicists []*Physicist `protobuf:"bytes,1,rep,name=physicists,proto3" json:"physicists,omitempty"`
}

func (x *AllPhysicists) Reset() {
	*x = AllPhysicists{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_physicist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllPhysicists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllPhysicists) ProtoMessage() {}

func (x *AllPhysicists) ProtoReflect() protoreflect.Message {
	mi := &file_proto_physicist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllPhysicists.ProtoReflect.Descriptor instead.
func (*AllPhysicists) Descriptor() ([]byte, []int) {
	return file_proto_physicist_proto_rawDescGZIP(), []int{2}
}

func (x *AllPhysicists) GetPhysicists() []*Physicist {
	if x != nil {
		return x.Physicists
	}
	return nil
}

var File_proto_physicist_proto protoreflect.FileDescriptor

var file_proto_physicist_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16,
	0x0a, 0x14, 0x41, 0x6c, 0x6c, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8f, 0x02, 0x0a, 0x09, 0x50, 0x68, 0x79, 0x73, 0x69,
	0x63, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0b, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f,
	0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66,
	0x42, 0x69, 0x72, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x66, 0x44, 0x65, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x66, 0x44, 0x65, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69,
	0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x22, 0x51, 0x0a, 0x0d, 0x41, 0x6c, 0x6c, 0x50,
	0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x68, 0x79,
	0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x52,
	0x0a, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x73, 0x32, 0x81, 0x03, 0x0a, 0x0e,
	0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x52,
	0x0a, 0x10, 0x67, 0x65, 0x74, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69,
	0x73, 0x74, 0x12, 0x52, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69,
	0x73, 0x74, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x68, 0x79, 0x73, 0x69,
	0x63, 0x69, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x6c, 0x6f, 0x67, 0x69,
	0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x68, 0x79,
	0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x12, 0x64, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x41,
	0x6c, 0x6c, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x73, 0x12, 0x61, 0x0a, 0x1d,
	0x67, 0x65, 0x74, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x20, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x2e, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x30, 0x01, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_physicist_proto_rawDescOnce sync.Once
	file_proto_physicist_proto_rawDescData = file_proto_physicist_proto_rawDesc
)

func file_proto_physicist_proto_rawDescGZIP() []byte {
	file_proto_physicist_proto_rawDescOnce.Do(func() {
		file_proto_physicist_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_physicist_proto_rawDescData)
	})
	return file_proto_physicist_proto_rawDescData
}

var file_proto_physicist_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_physicist_proto_goTypes = []interface{}{
	(*AllPhysicistsRequest)(nil),   // 0: cosmological_constant.AllPhysicistsRequest
	(*Physicist)(nil),              // 1: cosmological_constant.Physicist
	(*AllPhysicists)(nil),          // 2: cosmological_constant.AllPhysicists
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
}
var file_proto_physicist_proto_depIdxs = []int32{
	3, // 0: cosmological_constant.Physicist.physicistId:type_name -> google.protobuf.StringValue
	1, // 1: cosmological_constant.AllPhysicists.physicists:type_name -> cosmological_constant.Physicist
	3, // 2: cosmological_constant.PhysicistsInfo.getPhysicistById:input_type -> google.protobuf.StringValue
	1, // 3: cosmological_constant.PhysicistsInfo.addPhysicist:input_type -> cosmological_constant.Physicist
	0, // 4: cosmological_constant.PhysicistsInfo.getAllPhysicist:input_type -> cosmological_constant.AllPhysicistsRequest
	3, // 5: cosmological_constant.PhysicistsInfo.getPhysicistsByCountryOfBirth:input_type -> google.protobuf.StringValue
	1, // 6: cosmological_constant.PhysicistsInfo.getPhysicistById:output_type -> cosmological_constant.Physicist
	1, // 7: cosmological_constant.PhysicistsInfo.addPhysicist:output_type -> cosmological_constant.Physicist
	2, // 8: cosmological_constant.PhysicistsInfo.getAllPhysicist:output_type -> cosmological_constant.AllPhysicists
	1, // 9: cosmological_constant.PhysicistsInfo.getPhysicistsByCountryOfBirth:output_type -> cosmological_constant.Physicist
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_physicist_proto_init() }
func file_proto_physicist_proto_init() {
	if File_proto_physicist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_physicist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllPhysicistsRequest); i {
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
		file_proto_physicist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Physicist); i {
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
		file_proto_physicist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllPhysicists); i {
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
			RawDescriptor: file_proto_physicist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_physicist_proto_goTypes,
		DependencyIndexes: file_proto_physicist_proto_depIdxs,
		MessageInfos:      file_proto_physicist_proto_msgTypes,
	}.Build()
	File_proto_physicist_proto = out.File
	file_proto_physicist_proto_rawDesc = nil
	file_proto_physicist_proto_goTypes = nil
	file_proto_physicist_proto_depIdxs = nil
}
