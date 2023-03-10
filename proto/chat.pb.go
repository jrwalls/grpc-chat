// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: api/chat.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserStatus int32

const (
	UserStatus_JOINED UserStatus = 0
	UserStatus_LEFT   UserStatus = 1
	UserStatus_AFK    UserStatus = 2
)

// Enum value maps for UserStatus.
var (
	UserStatus_name = map[int32]string{
		0: "JOINED",
		1: "LEFT",
		2: "AFK",
	}
	UserStatus_value = map[string]int32{
		"JOINED": 0,
		"LEFT":   1,
		"AFK":    2,
	}
)

func (x UserStatus) Enum() *UserStatus {
	p := new(UserStatus)
	*p = x
	return p
}

func (x UserStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_chat_proto_enumTypes[0].Descriptor()
}

func (UserStatus) Type() protoreflect.EnumType {
	return &file_api_chat_proto_enumTypes[0]
}

func (x UserStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatus.Descriptor instead.
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{0}
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body   string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Message) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

type ChangeUserStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string     `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Status UserStatus `protobuf:"varint,2,opt,name=status,proto3,enum=chat.UserStatus" json:"status,omitempty"`
}

func (x *ChangeUserStatus) Reset() {
	*x = ChangeUserStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeUserStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeUserStatus) ProtoMessage() {}

func (x *ChangeUserStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeUserStatus.ProtoReflect.Descriptor instead.
func (*ChangeUserStatus) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeUserStatus) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ChangeUserStatus) GetStatus() UserStatus {
	if x != nil {
		return x.Status
	}
	return UserStatus_JOINED
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Ask:
	//
	//	*Command_Message
	Ask isCommand_Ask `protobuf_oneof:"ask"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{4}
}

func (m *Command) GetAsk() isCommand_Ask {
	if m != nil {
		return m.Ask
	}
	return nil
}

func (x *Command) GetMessage() *Message {
	if x, ok := x.GetAsk().(*Command_Message); ok {
		return x.Message
	}
	return nil
}

type isCommand_Ask interface {
	isCommand_Ask()
}

type Command_Message struct {
	Message *Message `protobuf:"bytes,1,opt,name=message,proto3,oneof"`
}

func (*Command_Message) isCommand_Ask() {}

type Tell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Tell:
	//
	//	*Tell_Message
	//	*Tell_ChangeUserStatus
	Tell isTell_Tell `protobuf_oneof:"tell"`
}

func (x *Tell) Reset() {
	*x = Tell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tell) ProtoMessage() {}

func (x *Tell) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tell.ProtoReflect.Descriptor instead.
func (*Tell) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{5}
}

func (x *Tell) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *Tell) GetTell() isTell_Tell {
	if m != nil {
		return m.Tell
	}
	return nil
}

func (x *Tell) GetMessage() *Message {
	if x, ok := x.GetTell().(*Tell_Message); ok {
		return x.Message
	}
	return nil
}

func (x *Tell) GetChangeUserStatus() *ChangeUserStatus {
	if x, ok := x.GetTell().(*Tell_ChangeUserStatus); ok {
		return x.ChangeUserStatus
	}
	return nil
}

type isTell_Tell interface {
	isTell_Tell()
}

type Tell_Message struct {
	Message *Message `protobuf:"bytes,2,opt,name=message,proto3,oneof"`
}

type Tell_ChangeUserStatus struct {
	ChangeUserStatus *ChangeUserStatus `protobuf:"bytes,3,opt,name=changeUserStatus,proto3,oneof"`
}

func (*Tell_Message) isTell_Tell() {}

func (*Tell_ChangeUserStatus) isTell_Tell() {}

var File_api_chat_proto protoreflect.FileDescriptor

var file_api_chat_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x27, 0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x35, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22,
	0x54, 0x0a, 0x10, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3b, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x29, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x61,
	0x73, 0x6b, 0x22, 0x8f, 0x01, 0x0a, 0x04, 0x74, 0x65, 0x6c, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x10, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x06, 0x0a, 0x04,
	0x74, 0x65, 0x6c, 0x6c, 0x2a, 0x2b, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x4a, 0x4f, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x46, 0x4b, 0x10,
	0x02, 0x32, 0x78, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x32, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x54, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x0a, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x74, 0x65, 0x6c, 0x6c, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_chat_proto_rawDescOnce sync.Once
	file_api_chat_proto_rawDescData = file_api_chat_proto_rawDesc
)

func file_api_chat_proto_rawDescGZIP() []byte {
	file_api_chat_proto_rawDescOnce.Do(func() {
		file_api_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_chat_proto_rawDescData)
	})
	return file_api_chat_proto_rawDescData
}

var file_api_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_chat_proto_goTypes = []interface{}{
	(UserStatus)(0),          // 0: chat.userStatus
	(*LoginRequest)(nil),     // 1: chat.loginRequest
	(*LoginResponse)(nil),    // 2: chat.loginResponse
	(*Message)(nil),          // 3: chat.message
	(*ChangeUserStatus)(nil), // 4: chat.changeUserStatus
	(*Command)(nil),          // 5: chat.command
	(*Tell)(nil),             // 6: chat.tell
}
var file_api_chat_proto_depIdxs = []int32{
	0, // 0: chat.changeUserStatus.status:type_name -> chat.userStatus
	3, // 1: chat.command.message:type_name -> chat.message
	3, // 2: chat.tell.message:type_name -> chat.message
	4, // 3: chat.tell.changeUserStatus:type_name -> chat.changeUserStatus
	1, // 4: chat.ChatService.Login:input_type -> chat.loginRequest
	5, // 5: chat.ChatService.SubscribeToChannel:input_type -> chat.command
	2, // 6: chat.ChatService.Login:output_type -> chat.loginResponse
	6, // 7: chat.ChatService.SubscribeToChannel:output_type -> chat.tell
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_chat_proto_init() }
func file_api_chat_proto_init() {
	if File_api_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_api_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_api_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_api_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeUserStatus); i {
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
		file_api_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_api_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tell); i {
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
	file_api_chat_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Command_Message)(nil),
	}
	file_api_chat_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Tell_Message)(nil),
		(*Tell_ChangeUserStatus)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_chat_proto_goTypes,
		DependencyIndexes: file_api_chat_proto_depIdxs,
		EnumInfos:         file_api_chat_proto_enumTypes,
		MessageInfos:      file_api_chat_proto_msgTypes,
	}.Build()
	File_api_chat_proto = out.File
	file_api_chat_proto_rawDesc = nil
	file_api_chat_proto_goTypes = nil
	file_api_chat_proto_depIdxs = nil
}
