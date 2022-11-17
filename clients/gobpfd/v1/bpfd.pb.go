// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: bpfd.proto

package v1

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

type ProgramType int32

const (
	ProgramType_XDP        ProgramType = 0
	ProgramType_TC         ProgramType = 1
	ProgramType_TRACEPOINT ProgramType = 2
)

// Enum value maps for ProgramType.
var (
	ProgramType_name = map[int32]string{
		0: "XDP",
		1: "TC",
		2: "TRACEPOINT",
	}
	ProgramType_value = map[string]int32{
		"XDP":        0,
		"TC":         1,
		"TRACEPOINT": 2,
	}
)

func (x ProgramType) Enum() *ProgramType {
	p := new(ProgramType)
	*p = x
	return p
}

func (x ProgramType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProgramType) Descriptor() protoreflect.EnumDescriptor {
	return file_bpfd_proto_enumTypes[0].Descriptor()
}

func (ProgramType) Type() protoreflect.EnumType {
	return &file_bpfd_proto_enumTypes[0]
}

func (x ProgramType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProgramType.Descriptor instead.
func (ProgramType) EnumDescriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{0}
}

type Direction int32

const (
	Direction_NONE    Direction = 0
	Direction_INGRESS Direction = 1
	Direction_EGRESS  Direction = 2
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "NONE",
		1: "INGRESS",
		2: "EGRESS",
	}
	Direction_value = map[string]int32{
		"NONE":    0,
		"INGRESS": 1,
		"EGRESS":  2,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_bpfd_proto_enumTypes[1].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_bpfd_proto_enumTypes[1]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{1}
}

type ProceedOn int32

const (
	ProceedOn_ABORTED           ProceedOn = 0
	ProceedOn_DROP              ProceedOn = 1
	ProceedOn_PASS              ProceedOn = 2
	ProceedOn_TX                ProceedOn = 3
	ProceedOn_REDIRECT          ProceedOn = 4
	ProceedOn_DISPATCHER_RETURN ProceedOn = 31
)

// Enum value maps for ProceedOn.
var (
	ProceedOn_name = map[int32]string{
		0:  "ABORTED",
		1:  "DROP",
		2:  "PASS",
		3:  "TX",
		4:  "REDIRECT",
		31: "DISPATCHER_RETURN",
	}
	ProceedOn_value = map[string]int32{
		"ABORTED":           0,
		"DROP":              1,
		"PASS":              2,
		"TX":                3,
		"REDIRECT":          4,
		"DISPATCHER_RETURN": 31,
	}
)

func (x ProceedOn) Enum() *ProceedOn {
	p := new(ProceedOn)
	*p = x
	return p
}

func (x ProceedOn) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProceedOn) Descriptor() protoreflect.EnumDescriptor {
	return file_bpfd_proto_enumTypes[2].Descriptor()
}

func (ProceedOn) Type() protoreflect.EnumType {
	return &file_bpfd_proto_enumTypes[2]
}

func (x ProceedOn) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProceedOn.Descriptor instead.
func (ProceedOn) EnumDescriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{2}
}

type LoadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path        string      `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	FromImage   bool        `protobuf:"varint,2,opt,name=from_image,json=fromImage,proto3" json:"from_image,omitempty"`
	SectionName string      `protobuf:"bytes,3,opt,name=section_name,json=sectionName,proto3" json:"section_name,omitempty"`
	ProgramType ProgramType `protobuf:"varint,4,opt,name=program_type,json=programType,proto3,enum=bpfd.v1.ProgramType" json:"program_type,omitempty"`
	Direction   Direction   `protobuf:"varint,5,opt,name=direction,proto3,enum=bpfd.v1.Direction" json:"direction,omitempty"`
	// Types that are assignable to AttachType:
	//	*LoadRequest_NetworkMultiAttach
	//	*LoadRequest_SingleAttach
	AttachType isLoadRequest_AttachType `protobuf_oneof:"attach_type"`
}

func (x *LoadRequest) Reset() {
	*x = LoadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadRequest) ProtoMessage() {}

func (x *LoadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadRequest.ProtoReflect.Descriptor instead.
func (*LoadRequest) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{0}
}

func (x *LoadRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *LoadRequest) GetFromImage() bool {
	if x != nil {
		return x.FromImage
	}
	return false
}

func (x *LoadRequest) GetSectionName() string {
	if x != nil {
		return x.SectionName
	}
	return ""
}

func (x *LoadRequest) GetProgramType() ProgramType {
	if x != nil {
		return x.ProgramType
	}
	return ProgramType_XDP
}

func (x *LoadRequest) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_NONE
}

func (m *LoadRequest) GetAttachType() isLoadRequest_AttachType {
	if m != nil {
		return m.AttachType
	}
	return nil
}

func (x *LoadRequest) GetNetworkMultiAttach() *NetworkMultiAttach {
	if x, ok := x.GetAttachType().(*LoadRequest_NetworkMultiAttach); ok {
		return x.NetworkMultiAttach
	}
	return nil
}

func (x *LoadRequest) GetSingleAttach() *SingleAttach {
	if x, ok := x.GetAttachType().(*LoadRequest_SingleAttach); ok {
		return x.SingleAttach
	}
	return nil
}

type isLoadRequest_AttachType interface {
	isLoadRequest_AttachType()
}

type LoadRequest_NetworkMultiAttach struct {
	NetworkMultiAttach *NetworkMultiAttach `protobuf:"bytes,6,opt,name=network_multi_attach,json=networkMultiAttach,proto3,oneof"`
}

type LoadRequest_SingleAttach struct {
	SingleAttach *SingleAttach `protobuf:"bytes,7,opt,name=single_attach,json=singleAttach,proto3,oneof"`
}

func (*LoadRequest_NetworkMultiAttach) isLoadRequest_AttachType() {}

func (*LoadRequest_SingleAttach) isLoadRequest_AttachType() {}

type NetworkMultiAttach struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Priority  int32       `protobuf:"varint,1,opt,name=priority,proto3" json:"priority,omitempty"`
	Iface     string      `protobuf:"bytes,2,opt,name=iface,proto3" json:"iface,omitempty"`
	Position  int32       `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
	ProceedOn []ProceedOn `protobuf:"varint,4,rep,packed,name=proceed_on,json=proceedOn,proto3,enum=bpfd.v1.ProceedOn" json:"proceed_on,omitempty"`
}

func (x *NetworkMultiAttach) Reset() {
	*x = NetworkMultiAttach{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkMultiAttach) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkMultiAttach) ProtoMessage() {}

func (x *NetworkMultiAttach) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkMultiAttach.ProtoReflect.Descriptor instead.
func (*NetworkMultiAttach) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkMultiAttach) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *NetworkMultiAttach) GetIface() string {
	if x != nil {
		return x.Iface
	}
	return ""
}

func (x *NetworkMultiAttach) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *NetworkMultiAttach) GetProceedOn() []ProceedOn {
	if x != nil {
		return x.ProceedOn
	}
	return nil
}

type SingleAttach struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SingleAttach) Reset() {
	*x = SingleAttach{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleAttach) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleAttach) ProtoMessage() {}

func (x *SingleAttach) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleAttach.ProtoReflect.Descriptor instead.
func (*SingleAttach) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{2}
}

func (x *SingleAttach) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LoadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LoadResponse) Reset() {
	*x = LoadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadResponse) ProtoMessage() {}

func (x *LoadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadResponse.ProtoReflect.Descriptor instead.
func (*LoadResponse) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{3}
}

func (x *LoadResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UnloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UnloadRequest) Reset() {
	*x = UnloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnloadRequest) ProtoMessage() {}

func (x *UnloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnloadRequest.ProtoReflect.Descriptor instead.
func (*UnloadRequest) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{4}
}

func (x *UnloadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UnloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnloadResponse) Reset() {
	*x = UnloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnloadResponse) ProtoMessage() {}

func (x *UnloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnloadResponse.ProtoReflect.Descriptor instead.
func (*UnloadResponse) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{5}
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{6}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ListResponse_ListResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{7}
}

func (x *ListResponse) GetResults() []*ListResponse_ListResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type ListResponse_ListResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Path        string      `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	ProgramType ProgramType `protobuf:"varint,4,opt,name=program_type,json=programType,proto3,enum=bpfd.v1.ProgramType" json:"program_type,omitempty"`
	Direction   Direction   `protobuf:"varint,5,opt,name=direction,proto3,enum=bpfd.v1.Direction" json:"direction,omitempty"`
	// Types that are assignable to AttachType:
	//	*ListResponse_ListResult_NetworkMultiAttach
	//	*ListResponse_ListResult_SingleAttach
	AttachType isListResponse_ListResult_AttachType `protobuf_oneof:"attach_type"`
}

func (x *ListResponse_ListResult) Reset() {
	*x = ListResponse_ListResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bpfd_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse_ListResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse_ListResult) ProtoMessage() {}

func (x *ListResponse_ListResult) ProtoReflect() protoreflect.Message {
	mi := &file_bpfd_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse_ListResult.ProtoReflect.Descriptor instead.
func (*ListResponse_ListResult) Descriptor() ([]byte, []int) {
	return file_bpfd_proto_rawDescGZIP(), []int{7, 0}
}

func (x *ListResponse_ListResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListResponse_ListResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListResponse_ListResult) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ListResponse_ListResult) GetProgramType() ProgramType {
	if x != nil {
		return x.ProgramType
	}
	return ProgramType_XDP
}

func (x *ListResponse_ListResult) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_NONE
}

func (m *ListResponse_ListResult) GetAttachType() isListResponse_ListResult_AttachType {
	if m != nil {
		return m.AttachType
	}
	return nil
}

func (x *ListResponse_ListResult) GetNetworkMultiAttach() *NetworkMultiAttach {
	if x, ok := x.GetAttachType().(*ListResponse_ListResult_NetworkMultiAttach); ok {
		return x.NetworkMultiAttach
	}
	return nil
}

func (x *ListResponse_ListResult) GetSingleAttach() *SingleAttach {
	if x, ok := x.GetAttachType().(*ListResponse_ListResult_SingleAttach); ok {
		return x.SingleAttach
	}
	return nil
}

type isListResponse_ListResult_AttachType interface {
	isListResponse_ListResult_AttachType()
}

type ListResponse_ListResult_NetworkMultiAttach struct {
	NetworkMultiAttach *NetworkMultiAttach `protobuf:"bytes,6,opt,name=network_multi_attach,json=networkMultiAttach,proto3,oneof"`
}

type ListResponse_ListResult_SingleAttach struct {
	SingleAttach *SingleAttach `protobuf:"bytes,7,opt,name=single_attach,json=singleAttach,proto3,oneof"`
}

func (*ListResponse_ListResult_NetworkMultiAttach) isListResponse_ListResult_AttachType() {}

func (*ListResponse_ListResult_SingleAttach) isListResponse_ListResult_AttachType() {}

var File_bpfd_proto protoreflect.FileDescriptor

var file_bpfd_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x70,
	0x66, 0x64, 0x2e, 0x76, 0x31, 0x22, 0xec, 0x02, 0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x66,
	0x72, 0x6f, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x14, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x48, 0x00, 0x52, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x12, 0x3c, 0x0a, 0x0d, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x42, 0x0d, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x12, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x66, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x66, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x65, 0x64, 0x4f,
	0x6e, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x65, 0x64, 0x4f, 0x6e, 0x22, 0x22, 0x0a, 0x0c,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x1e, 0x0a, 0x0c, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x1f, 0x0a, 0x0d, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x9a, 0x03, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a,
	0xcd, 0x02, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x37, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x62,
	0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x30, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x4f, 0x0a, 0x14, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x48, 0x00, 0x52, 0x12,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x12, 0x3c, 0x0a, 0x0d, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x70, 0x66, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x48, 0x00, 0x52, 0x0c, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x42, 0x0d, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2a,
	0x2e, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07,
	0x0a, 0x03, 0x58, 0x44, 0x50, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x54, 0x43, 0x10, 0x01, 0x12,
	0x0e, 0x0a, 0x0a, 0x54, 0x52, 0x41, 0x43, 0x45, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x02, 0x2a,
	0x2e, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x47, 0x52, 0x45, 0x53,
	0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x2a,
	0x59, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x0b, 0x0a, 0x07,
	0x41, 0x42, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x52, 0x4f,
	0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x53, 0x53, 0x10, 0x02, 0x12, 0x06, 0x0a,
	0x02, 0x54, 0x58, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43,
	0x54, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x49, 0x53, 0x50, 0x41, 0x54, 0x43, 0x48, 0x45,
	0x52, 0x5f, 0x52, 0x45, 0x54, 0x55, 0x52, 0x4e, 0x10, 0x1f, 0x32, 0xad, 0x01, 0x0a, 0x06, 0x4c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x2e,
	0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x55, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62,
	0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e,
	0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x70, 0x66, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x64, 0x68, 0x61, 0x74, 0x2d,
	0x65, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x62, 0x70, 0x66,
	0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bpfd_proto_rawDescOnce sync.Once
	file_bpfd_proto_rawDescData = file_bpfd_proto_rawDesc
)

func file_bpfd_proto_rawDescGZIP() []byte {
	file_bpfd_proto_rawDescOnce.Do(func() {
		file_bpfd_proto_rawDescData = protoimpl.X.CompressGZIP(file_bpfd_proto_rawDescData)
	})
	return file_bpfd_proto_rawDescData
}

var file_bpfd_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_bpfd_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_bpfd_proto_goTypes = []interface{}{
	(ProgramType)(0),                // 0: bpfd.v1.ProgramType
	(Direction)(0),                  // 1: bpfd.v1.Direction
	(ProceedOn)(0),                  // 2: bpfd.v1.ProceedOn
	(*LoadRequest)(nil),             // 3: bpfd.v1.LoadRequest
	(*NetworkMultiAttach)(nil),      // 4: bpfd.v1.NetworkMultiAttach
	(*SingleAttach)(nil),            // 5: bpfd.v1.SingleAttach
	(*LoadResponse)(nil),            // 6: bpfd.v1.LoadResponse
	(*UnloadRequest)(nil),           // 7: bpfd.v1.UnloadRequest
	(*UnloadResponse)(nil),          // 8: bpfd.v1.UnloadResponse
	(*ListRequest)(nil),             // 9: bpfd.v1.ListRequest
	(*ListResponse)(nil),            // 10: bpfd.v1.ListResponse
	(*ListResponse_ListResult)(nil), // 11: bpfd.v1.ListResponse.ListResult
}
var file_bpfd_proto_depIdxs = []int32{
	0,  // 0: bpfd.v1.LoadRequest.program_type:type_name -> bpfd.v1.ProgramType
	1,  // 1: bpfd.v1.LoadRequest.direction:type_name -> bpfd.v1.Direction
	4,  // 2: bpfd.v1.LoadRequest.network_multi_attach:type_name -> bpfd.v1.NetworkMultiAttach
	5,  // 3: bpfd.v1.LoadRequest.single_attach:type_name -> bpfd.v1.SingleAttach
	2,  // 4: bpfd.v1.NetworkMultiAttach.proceed_on:type_name -> bpfd.v1.ProceedOn
	11, // 5: bpfd.v1.ListResponse.results:type_name -> bpfd.v1.ListResponse.ListResult
	0,  // 6: bpfd.v1.ListResponse.ListResult.program_type:type_name -> bpfd.v1.ProgramType
	1,  // 7: bpfd.v1.ListResponse.ListResult.direction:type_name -> bpfd.v1.Direction
	4,  // 8: bpfd.v1.ListResponse.ListResult.network_multi_attach:type_name -> bpfd.v1.NetworkMultiAttach
	5,  // 9: bpfd.v1.ListResponse.ListResult.single_attach:type_name -> bpfd.v1.SingleAttach
	3,  // 10: bpfd.v1.Loader.Load:input_type -> bpfd.v1.LoadRequest
	7,  // 11: bpfd.v1.Loader.Unload:input_type -> bpfd.v1.UnloadRequest
	9,  // 12: bpfd.v1.Loader.List:input_type -> bpfd.v1.ListRequest
	6,  // 13: bpfd.v1.Loader.Load:output_type -> bpfd.v1.LoadResponse
	8,  // 14: bpfd.v1.Loader.Unload:output_type -> bpfd.v1.UnloadResponse
	10, // 15: bpfd.v1.Loader.List:output_type -> bpfd.v1.ListResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_bpfd_proto_init() }
func file_bpfd_proto_init() {
	if File_bpfd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bpfd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadRequest); i {
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
		file_bpfd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkMultiAttach); i {
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
		file_bpfd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleAttach); i {
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
		file_bpfd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadResponse); i {
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
		file_bpfd_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnloadRequest); i {
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
		file_bpfd_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnloadResponse); i {
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
		file_bpfd_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_bpfd_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_bpfd_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse_ListResult); i {
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
	file_bpfd_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LoadRequest_NetworkMultiAttach)(nil),
		(*LoadRequest_SingleAttach)(nil),
	}
	file_bpfd_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*ListResponse_ListResult_NetworkMultiAttach)(nil),
		(*ListResponse_ListResult_SingleAttach)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bpfd_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bpfd_proto_goTypes,
		DependencyIndexes: file_bpfd_proto_depIdxs,
		EnumInfos:         file_bpfd_proto_enumTypes,
		MessageInfos:      file_bpfd_proto_msgTypes,
	}.Build()
	File_bpfd_proto = out.File
	file_bpfd_proto_rawDesc = nil
	file_bpfd_proto_goTypes = nil
	file_bpfd_proto_depIdxs = nil
}
