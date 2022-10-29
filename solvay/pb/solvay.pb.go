// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: proto/solvay.proto

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

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UuidInString string `protobuf:"bytes,1,opt,name=uuidInString,proto3" json:"uuidInString,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_solvay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solvay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_proto_solvay_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetUuidInString() string {
	if x != nil {
		return x.UuidInString
	}
	return ""
}

type Conference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConferenceId           *UUID                    `protobuf:"bytes,1,opt,name=conferenceId,proto3" json:"conferenceId,omitempty"`
	Name                   string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location               string                   `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Description            string                   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	StartTime              int64                    `protobuf:"varint,5,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime                int64                    `protobuf:"varint,6,opt,name=endTime,proto3" json:"endTime,omitempty"`
	PhysicistsInConference []*PhysicistInConference `protobuf:"bytes,7,rep,name=physicistsInConference,proto3" json:"physicistsInConference,omitempty"`
}

func (x *Conference) Reset() {
	*x = Conference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_solvay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conference) ProtoMessage() {}

func (x *Conference) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solvay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conference.ProtoReflect.Descriptor instead.
func (*Conference) Descriptor() ([]byte, []int) {
	return file_proto_solvay_proto_rawDescGZIP(), []int{1}
}

func (x *Conference) GetConferenceId() *UUID {
	if x != nil {
		return x.ConferenceId
	}
	return nil
}

func (x *Conference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Conference) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Conference) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Conference) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Conference) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *Conference) GetPhysicistsInConference() []*PhysicistInConference {
	if x != nil {
		return x.PhysicistsInConference
	}
	return nil
}

type DesignatedConferences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conferences []*Conference `protobuf:"bytes,1,rep,name=conferences,proto3" json:"conferences,omitempty"`
}

func (x *DesignatedConferences) Reset() {
	*x = DesignatedConferences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_solvay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DesignatedConferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DesignatedConferences) ProtoMessage() {}

func (x *DesignatedConferences) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solvay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DesignatedConferences.ProtoReflect.Descriptor instead.
func (*DesignatedConferences) Descriptor() ([]byte, []int) {
	return file_proto_solvay_proto_rawDescGZIP(), []int{2}
}

func (x *DesignatedConferences) GetConferences() []*Conference {
	if x != nil {
		return x.Conferences
	}
	return nil
}

//
// Temporary solution until I figure out how to import the physicist message from physicists-info
type PhysicistInConference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhysicistId    *UUID  `protobuf:"bytes,1,opt,name=physicistId,proto3" json:"physicistId,omitempty"`
	FirstName      string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName       string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	CountryOfBirth string `protobuf:"bytes,4,opt,name=countryOfBirth,proto3" json:"countryOfBirth,omitempty"`
}

func (x *PhysicistInConference) Reset() {
	*x = PhysicistInConference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_solvay_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhysicistInConference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhysicistInConference) ProtoMessage() {}

func (x *PhysicistInConference) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solvay_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhysicistInConference.ProtoReflect.Descriptor instead.
func (*PhysicistInConference) Descriptor() ([]byte, []int) {
	return file_proto_solvay_proto_rawDescGZIP(), []int{3}
}

func (x *PhysicistInConference) GetPhysicistId() *UUID {
	if x != nil {
		return x.PhysicistId
	}
	return nil
}

func (x *PhysicistInConference) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PhysicistInConference) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *PhysicistInConference) GetCountryOfBirth() string {
	if x != nil {
		return x.CountryOfBirth
	}
	return ""
}

var File_proto_solvay_proto protoreflect.FileDescriptor

var file_proto_solvay_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x76, 0x61, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x6f, 0x6c, 0x76, 0x61, 0x79, 0x22, 0x2a, 0x0a, 0x04,
	0x55, 0x55, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x75, 0x69, 0x64, 0x49, 0x6e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x75, 0x69, 0x64,
	0x49, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x9f, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6e,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x73, 0x6f, 0x6c, 0x76, 0x61, 0x79, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x55, 0x0a, 0x16, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74,
	0x73, 0x49, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x6f, 0x6c, 0x76, 0x61, 0x79, 0x2e, 0x50, 0x68, 0x79,
	0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x16, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x73, 0x49, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x4d, 0x0a, 0x15, 0x44, 0x65,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x6f, 0x6c, 0x76, 0x61,
	0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x15, 0x50, 0x68,
	0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x6f, 0x6c, 0x76, 0x61,
	0x79, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x0b, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x66,
	0x42, 0x69, 0x72, 0x74, 0x68, 0x32, 0x6e, 0x0a, 0x0d, 0x53, 0x6f, 0x6c, 0x76, 0x61, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x50, 0x68, 0x79, 0x73, 0x69,
	0x63, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x73, 0x6f, 0x6c, 0x76, 0x61, 0x79, 0x2e, 0x50,
	0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x1a, 0x1d, 0x2e, 0x73, 0x6f, 0x6c, 0x76, 0x61, 0x79, 0x2e, 0x44, 0x65,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x28, 0x01, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_solvay_proto_rawDescOnce sync.Once
	file_proto_solvay_proto_rawDescData = file_proto_solvay_proto_rawDesc
)

func file_proto_solvay_proto_rawDescGZIP() []byte {
	file_proto_solvay_proto_rawDescOnce.Do(func() {
		file_proto_solvay_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_solvay_proto_rawDescData)
	})
	return file_proto_solvay_proto_rawDescData
}

var file_proto_solvay_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_solvay_proto_goTypes = []interface{}{
	(*UUID)(nil),                  // 0: solvay.UUID
	(*Conference)(nil),            // 1: solvay.Conference
	(*DesignatedConferences)(nil), // 2: solvay.DesignatedConferences
	(*PhysicistInConference)(nil), // 3: solvay.PhysicistInConference
}
var file_proto_solvay_proto_depIdxs = []int32{
	0, // 0: solvay.Conference.conferenceId:type_name -> solvay.UUID
	3, // 1: solvay.Conference.physicistsInConference:type_name -> solvay.PhysicistInConference
	1, // 2: solvay.DesignatedConferences.conferences:type_name -> solvay.Conference
	0, // 3: solvay.PhysicistInConference.physicistId:type_name -> solvay.UUID
	3, // 4: solvay.SolvayService.GetConferencesForPhysicists:input_type -> solvay.PhysicistInConference
	2, // 5: solvay.SolvayService.GetConferencesForPhysicists:output_type -> solvay.DesignatedConferences
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_solvay_proto_init() }
func file_proto_solvay_proto_init() {
	if File_proto_solvay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_solvay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_proto_solvay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conference); i {
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
		file_proto_solvay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DesignatedConferences); i {
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
		file_proto_solvay_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhysicistInConference); i {
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
			RawDescriptor: file_proto_solvay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_solvay_proto_goTypes,
		DependencyIndexes: file_proto_solvay_proto_depIdxs,
		MessageInfos:      file_proto_solvay_proto_msgTypes,
	}.Build()
	File_proto_solvay_proto = out.File
	file_proto_solvay_proto_rawDesc = nil
	file_proto_solvay_proto_goTypes = nil
	file_proto_solvay_proto_depIdxs = nil
}