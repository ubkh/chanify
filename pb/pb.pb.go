// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.15.6
// source: pb.proto

package pb

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

type ChanType int32

const (
	ChanType_None ChanType = 0
	ChanType_Sys  ChanType = 1
	ChanType_User ChanType = 2
)

// Enum value maps for ChanType.
var (
	ChanType_name = map[int32]string{
		0: "None",
		1: "Sys",
		2: "User",
	}
	ChanType_value = map[string]int32{
		"None": 0,
		"Sys":  1,
		"User": 2,
	}
)

func (x ChanType) Enum() *ChanType {
	p := new(ChanType)
	*p = x
	return p
}

func (x ChanType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChanType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_proto_enumTypes[0].Descriptor()
}

func (ChanType) Type() protoreflect.EnumType {
	return &file_pb_proto_enumTypes[0]
}

func (x ChanType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChanType.Descriptor instead.
func (ChanType) EnumDescriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{0}
}

type ChanCode int32

const (
	ChanCode_Uncategorized ChanCode = 0
	ChanCode_Device        ChanCode = 1
)

// Enum value maps for ChanCode.
var (
	ChanCode_name = map[int32]string{
		0: "Uncategorized",
		1: "Device",
	}
	ChanCode_value = map[string]int32{
		"Uncategorized": 0,
		"Device":        1,
	}
)

func (x ChanCode) Enum() *ChanCode {
	p := new(ChanCode)
	*p = x
	return p
}

func (x ChanCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChanCode) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_proto_enumTypes[1].Descriptor()
}

func (ChanCode) Type() protoreflect.EnumType {
	return &file_pb_proto_enumTypes[1]
}

func (x ChanCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChanCode.Descriptor instead.
func (ChanCode) EnumDescriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{1}
}

type MsgType int32

const (
	MsgType_System MsgType = 0
	MsgType_Text   MsgType = 1
	MsgType_Image  MsgType = 2
	MsgType_Video  MsgType = 3
	MsgType_Audio  MsgType = 4
	MsgType_Link   MsgType = 5
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0: "System",
		1: "Text",
		2: "Image",
		3: "Video",
		4: "Audio",
		5: "Link",
	}
	MsgType_value = map[string]int32{
		"System": 0,
		"Text":   1,
		"Image":  2,
		"Video":  3,
		"Audio":  4,
		"Link":   5,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_proto_enumTypes[2].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_pb_proto_enumTypes[2]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{2}
}

type SoundType int32

const (
	SoundType_NormalSound   SoundType = 0
	SoundType_CriticalSound SoundType = 1
)

// Enum value maps for SoundType.
var (
	SoundType_name = map[int32]string{
		0: "NormalSound",
		1: "CriticalSound",
	}
	SoundType_value = map[string]int32{
		"NormalSound":   0,
		"CriticalSound": 1,
	}
)

func (x SoundType) Enum() *SoundType {
	p := new(SoundType)
	*p = x
	return p
}

func (x SoundType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SoundType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_proto_enumTypes[3].Descriptor()
}

func (SoundType) Type() protoreflect.EnumType {
	return &file_pb_proto_enumTypes[3]
}

func (x SoundType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SoundType.Descriptor instead.
func (SoundType) EnumDescriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{3}
}

type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ChanType `protobuf:"varint,1,opt,name=type,proto3,enum=net.chanify.model.ChanType" json:"type,omitempty"`
	Code ChanCode `protobuf:"varint,2,opt,name=code,proto3,enum=net.chanify.model.ChanCode" json:"code,omitempty"`
	Name string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Icon string   `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{0}
}

func (x *Channel) GetType() ChanType {
	if x != nil {
		return x.Type
	}
	return ChanType_None
}

func (x *Channel) GetCode() ChanCode {
	if x != nil {
		return x.Code
	}
	return ChanCode_Uncategorized
}

func (x *Channel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Channel) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expires  uint64 `protobuf:"varint,1,opt,name=expires,proto3" json:"expires,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeviceId []byte `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Channel  []byte `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"`
	NodeId   string `protobuf:"bytes,5,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	DataHash []byte `protobuf:"bytes,6,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetExpires() uint64 {
	if x != nil {
		return x.Expires
	}
	return 0
}

func (x *Token) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Token) GetDeviceId() []byte {
	if x != nil {
		return x.DeviceId
	}
	return nil
}

func (x *Token) GetChannel() []byte {
	if x != nil {
		return x.Channel
	}
	return nil
}

func (x *Token) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Token) GetDataHash() []byte {
	if x != nil {
		return x.DataHash
	}
	return nil
}

type Thumbnail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   uint32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Width  int32  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height int32  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Data   []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Thumbnail) Reset() {
	*x = Thumbnail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Thumbnail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thumbnail) ProtoMessage() {}

func (x *Thumbnail) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thumbnail.ProtoReflect.Descriptor instead.
func (*Thumbnail) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{2}
}

func (x *Thumbnail) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Thumbnail) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Thumbnail) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Thumbnail) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type MsgContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      MsgType    `protobuf:"varint,1,opt,name=type,proto3,enum=net.chanify.model.MsgType" json:"type,omitempty"`
	Text      string     `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	File      string     `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	Title     string     `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Thumbnail *Thumbnail `protobuf:"bytes,5,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	Duration  uint64     `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	Size      uint64     `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	Link      string     `protobuf:"bytes,8,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *MsgContent) Reset() {
	*x = MsgContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgContent) ProtoMessage() {}

func (x *MsgContent) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgContent.ProtoReflect.Descriptor instead.
func (*MsgContent) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{3}
}

func (x *MsgContent) GetType() MsgType {
	if x != nil {
		return x.Type
	}
	return MsgType_System
}

func (x *MsgContent) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *MsgContent) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *MsgContent) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MsgContent) GetThumbnail() *Thumbnail {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

func (x *MsgContent) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *MsgContent) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *MsgContent) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type Sound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   SoundType `protobuf:"varint,1,opt,name=type,proto3,enum=net.chanify.model.SoundType" json:"type,omitempty"`
	Name   string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Volume float32   `protobuf:"fixed32,3,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *Sound) Reset() {
	*x = Sound{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sound) ProtoMessage() {}

func (x *Sound) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sound.ProtoReflect.Descriptor instead.
func (*Sound) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{4}
}

func (x *Sound) GetType() SoundType {
	if x != nil {
		return x.Type
	}
	return SoundType_NormalSound
}

func (x *Sound) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sound) GetVolume() float32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From       []byte `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Channel    []byte `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`       // Channel
	Content    []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`       // MsgContent
	Ciphertext []byte `protobuf:"bytes,4,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"` // Encrypt MsgContent
	Priority   int32  `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
	Sound      *Sound `protobuf:"bytes,6,opt,name=sound,proto3" json:"sound,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[5]
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
	return file_pb_proto_rawDescGZIP(), []int{5}
}

func (x *Message) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Message) GetChannel() []byte {
	if x != nil {
		return x.Channel
	}
	return nil
}

func (x *Message) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Message) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

func (x *Message) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Message) GetSound() *Sound {
	if x != nil {
		return x.Sound
	}
	return nil
}

var File_pb_proto protoreflect.FileDescriptor

var file_pb_proto_rawDesc = []byte{
	0x0a, 0x08, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x65, 0x74, 0x2e,
	0x63, 0x68, 0x61, 0x6e, 0x69, 0x66, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x93, 0x01,
	0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x69, 0x66, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63,
	0x68, 0x61, 0x6e, 0x69, 0x66, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x22, 0xa7, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x48, 0x61, 0x73, 0x68, 0x22, 0x61, 0x0a,
	0x09, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xfa, 0x01, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x6e, 0x65, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x66, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3a, 0x0a,
	0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x66, 0x79, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x09,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x65, 0x0a,
	0x05, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x69,
	0x66, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x69,
	0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x66,
	0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x05, 0x73,
	0x6f, 0x75, 0x6e, 0x64, 0x2a, 0x27, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x79,
	0x73, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x10, 0x02, 0x2a, 0x29, 0x0a,
	0x08, 0x43, 0x68, 0x61, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x6e, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x10, 0x01, 0x2a, 0x4a, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x03, 0x12,
	0x09, 0x0a, 0x05, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x69,
	0x6e, 0x6b, 0x10, 0x05, 0x2a, 0x2f, 0x0a, 0x09, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x53, 0x6f, 0x75, 0x6e, 0x64,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x6f,
	0x75, 0x6e, 0x64, 0x10, 0x01, 0x42, 0x0f, 0x48, 0x03, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0xa2,
	0x02, 0x04, 0x43, 0x48, 0x54, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_proto_rawDescOnce sync.Once
	file_pb_proto_rawDescData = file_pb_proto_rawDesc
)

func file_pb_proto_rawDescGZIP() []byte {
	file_pb_proto_rawDescOnce.Do(func() {
		file_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_proto_rawDescData)
	})
	return file_pb_proto_rawDescData
}

var file_pb_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_proto_goTypes = []interface{}{
	(ChanType)(0),      // 0: net.chanify.model.ChanType
	(ChanCode)(0),      // 1: net.chanify.model.ChanCode
	(MsgType)(0),       // 2: net.chanify.model.MsgType
	(SoundType)(0),     // 3: net.chanify.model.SoundType
	(*Channel)(nil),    // 4: net.chanify.model.Channel
	(*Token)(nil),      // 5: net.chanify.model.Token
	(*Thumbnail)(nil),  // 6: net.chanify.model.Thumbnail
	(*MsgContent)(nil), // 7: net.chanify.model.MsgContent
	(*Sound)(nil),      // 8: net.chanify.model.Sound
	(*Message)(nil),    // 9: net.chanify.model.Message
}
var file_pb_proto_depIdxs = []int32{
	0, // 0: net.chanify.model.Channel.type:type_name -> net.chanify.model.ChanType
	1, // 1: net.chanify.model.Channel.code:type_name -> net.chanify.model.ChanCode
	2, // 2: net.chanify.model.MsgContent.type:type_name -> net.chanify.model.MsgType
	6, // 3: net.chanify.model.MsgContent.thumbnail:type_name -> net.chanify.model.Thumbnail
	3, // 4: net.chanify.model.Sound.type:type_name -> net.chanify.model.SoundType
	8, // 5: net.chanify.model.Message.sound:type_name -> net.chanify.model.Sound
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pb_proto_init() }
func file_pb_proto_init() {
	if File_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channel); i {
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
		file_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Thumbnail); i {
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
		file_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgContent); i {
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
		file_pb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sound); i {
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
		file_pb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_proto_goTypes,
		DependencyIndexes: file_pb_proto_depIdxs,
		EnumInfos:         file_pb_proto_enumTypes,
		MessageInfos:      file_pb_proto_msgTypes,
	}.Build()
	File_pb_proto = out.File
	file_pb_proto_rawDesc = nil
	file_pb_proto_goTypes = nil
	file_pb_proto_depIdxs = nil
}
