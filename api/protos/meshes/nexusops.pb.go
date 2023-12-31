// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: nexusops.proto

package meshes

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type OperationCategory int32

const (
	OperationCategory_INSTALL   OperationCategory = 0
	OperationCategory_SAMPLE    OperationCategory = 1
	OperationCategory_CONFIGURE OperationCategory = 2
	OperationCategory_VALIDATE  OperationCategory = 3
	OperationCategory_CUSTOM    OperationCategory = 4
)

// Enum value maps for OperationCategory.
var (
	OperationCategory_name = map[int32]string{
		0: "INSTALL",
		1: "SAMPLE",
		2: "CONFIGURE",
		3: "VALIDATE",
		4: "CUSTOM",
	}
	OperationCategory_value = map[string]int32{
		"INSTALL":   0,
		"SAMPLE":    1,
		"CONFIGURE": 2,
		"VALIDATE":  3,
		"CUSTOM":    4,
	}
)

func (x OperationCategory) Enum() *OperationCategory {
	p := new(OperationCategory)
	*p = x
	return p
}

func (x OperationCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_nexusops_proto_enumTypes[0].Descriptor()
}

func (OperationCategory) Type() protoreflect.EnumType {
	return &file_nexusops_proto_enumTypes[0]
}

func (x OperationCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationCategory.Descriptor instead.
func (OperationCategory) EnumDescriptor() ([]byte, []int) {
	return file_nexusops_proto_rawDescGZIP(), []int{0}
}

type EventSeverity int32

const (
	EventSeverity_INFO  EventSeverity = 0
	EventSeverity_WARN  EventSeverity = 1
	EventSeverity_ERROR EventSeverity = 2
)

// Enum value maps for EventSeverity.
var (
	EventSeverity_name = map[int32]string{
		0: "INFO",
		1: "WARN",
		2: "ERROR",
	}
	EventSeverity_value = map[string]int32{
		"INFO":  0,
		"WARN":  1,
		"ERROR": 2,
	}
)

func (x EventSeverity) Enum() *EventSeverity {
	p := new(EventSeverity)
	*p = x
	return p
}

func (x EventSeverity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventSeverity) Descriptor() protoreflect.EnumDescriptor {
	return file_nexusops_proto_enumTypes[1].Descriptor()
}

func (EventSeverity) Type() protoreflect.EnumType {
	return &file_nexusops_proto_enumTypes[1]
}

func (x EventSeverity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventSeverity.Descriptor instead.
func (EventSeverity) EnumDescriptor() ([]byte, []int) {
	return file_nexusops_proto_rawDescGZIP(), []int{1}
}

type CloudEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloudEventsRequest) Reset() {
	*x = CloudEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nexusops_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudEventsRequest) ProtoMessage() {}

func (x *CloudEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nexusops_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudEventsRequest.ProtoReflect.Descriptor instead.
func (*CloudEventsRequest) Descriptor() ([]byte, []int) {
	return file_nexusops_proto_rawDescGZIP(), []int{0}
}

type CloudEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType EventSeverity `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=meshes.EventSeverity" json:"event_type,omitempty"`
	Summary   string        `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *CloudEventsResponse) Reset() {
	*x = CloudEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nexusops_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudEventsResponse) ProtoMessage() {}

func (x *CloudEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nexusops_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudEventsResponse.ProtoReflect.Descriptor instead.
func (*CloudEventsResponse) Descriptor() ([]byte, []int) {
	return file_nexusops_proto_rawDescGZIP(), []int{1}
}

func (x *CloudEventsResponse) GetEventType() EventSeverity {
	if x != nil {
		return x.EventType
	}
	return EventSeverity_INFO
}

func (x *CloudEventsResponse) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

type ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMessage string `protobuf:"bytes,1,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nexusops_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_nexusops_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_nexusops_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorMessage) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type NexusNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NexusNameRequest) Reset() {
	*x = NexusNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nexusops_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NexusNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NexusNameRequest) ProtoMessage() {}

func (x *NexusNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nexusops_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NexusNameRequest.ProtoReflect.Descriptor instead.
func (*NexusNameRequest) Descriptor() ([]byte, []int) {
	return file_nexusops_proto_rawDescGZIP(), []int{3}
}

type NexusNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NexusNameResponse) Reset() {
	*x = NexusNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nexusops_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NexusNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NexusNameResponse) ProtoMessage() {}

func (x *NexusNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nexusops_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NexusNameResponse.ProtoReflect.Descriptor instead.
func (*NexusNameResponse) Descriptor() ([]byte, []int) {
	return file_nexusops_proto_rawDescGZIP(), []int{4}
}

func (x *NexusNameResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nexusops_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nexusops_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_nexusops_proto_rawDescGZIP(), []int{5}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nexusops_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nexusops_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_nexusops_proto_rawDescGZIP(), []int{6}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_nexusops_proto protoreflect.FileDescriptor

var file_nexusops_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x65, 0x0a, 0x13,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x73,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x09,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x22, 0x33, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x4e, 0x65, 0x78, 0x75,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x11,
	0x4e, 0x65, 0x78, 0x75, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2a, 0x55, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x53,
	0x54, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x41, 0x4d, 0x50, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52, 0x45, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x04, 0x2a, 0x2e, 0x0a, 0x0d, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x08, 0x0a, 0x04,
	0x49, 0x4e, 0x46, 0x4f, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x32, 0xe6, 0x01, 0x0a, 0x0c,
	0x4e, 0x65, 0x78, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x05,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x73, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x12, 0x40, 0x0a, 0x09, 0x4e, 0x65, 0x78, 0x75, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x78, 0x75, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x65, 0x73, 0x2e, 0x4e, 0x65, 0x78, 0x75, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x65, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x73, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x75, 0x6c, 0x61, 0x2d, 0x61, 0x61, 0x63, 0x2f, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x73, 0x3b, 0x6d, 0x65, 0x73, 0x68, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nexusops_proto_rawDescOnce sync.Once
	file_nexusops_proto_rawDescData = file_nexusops_proto_rawDesc
)

func file_nexusops_proto_rawDescGZIP() []byte {
	file_nexusops_proto_rawDescOnce.Do(func() {
		file_nexusops_proto_rawDescData = protoimpl.X.CompressGZIP(file_nexusops_proto_rawDescData)
	})
	return file_nexusops_proto_rawDescData
}

var file_nexusops_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_nexusops_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_nexusops_proto_goTypes = []interface{}{
	(OperationCategory)(0),      // 0: meshes.OperationCategory
	(EventSeverity)(0),          // 1: meshes.EventSeverity
	(*CloudEventsRequest)(nil),  // 2: meshes.CloudEventsRequest
	(*CloudEventsResponse)(nil), // 3: meshes.CloudEventsResponse
	(*ErrorMessage)(nil),        // 4: meshes.ErrorMessage
	(*NexusNameRequest)(nil),    // 5: meshes.NexusNameRequest
	(*NexusNameResponse)(nil),   // 6: meshes.NexusNameResponse
	(*HelloRequest)(nil),        // 7: meshes.HelloRequest
	(*HelloResponse)(nil),       // 8: meshes.HelloResponse
}
var file_nexusops_proto_depIdxs = []int32{
	1, // 0: meshes.CloudEventsResponse.event_type:type_name -> meshes.EventSeverity
	7, // 1: meshes.NexusService.Hello:input_type -> meshes.HelloRequest
	5, // 2: meshes.NexusService.NexusName:input_type -> meshes.NexusNameRequest
	2, // 3: meshes.NexusService.StreamCloudEvents:input_type -> meshes.CloudEventsRequest
	8, // 4: meshes.NexusService.Hello:output_type -> meshes.HelloResponse
	6, // 5: meshes.NexusService.NexusName:output_type -> meshes.NexusNameResponse
	3, // 6: meshes.NexusService.StreamCloudEvents:output_type -> meshes.CloudEventsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nexusops_proto_init() }
func file_nexusops_proto_init() {
	if File_nexusops_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nexusops_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudEventsRequest); i {
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
		file_nexusops_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudEventsResponse); i {
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
		file_nexusops_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMessage); i {
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
		file_nexusops_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NexusNameRequest); i {
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
		file_nexusops_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NexusNameResponse); i {
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
		file_nexusops_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_nexusops_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
			RawDescriptor: file_nexusops_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nexusops_proto_goTypes,
		DependencyIndexes: file_nexusops_proto_depIdxs,
		EnumInfos:         file_nexusops_proto_enumTypes,
		MessageInfos:      file_nexusops_proto_msgTypes,
	}.Build()
	File_nexusops_proto = out.File
	file_nexusops_proto_rawDesc = nil
	file_nexusops_proto_goTypes = nil
	file_nexusops_proto_depIdxs = nil
}
