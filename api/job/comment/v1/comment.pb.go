// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: job/comment/v1/comment.proto

package jobcommentv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
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

type CommentMessage_Type int32

const (
	CommentMessage_ADDREPLY CommentMessage_Type = 0
	CommentMessage_ADDLIKE  CommentMessage_Type = 1
)

// Enum value maps for CommentMessage_Type.
var (
	CommentMessage_Type_name = map[int32]string{
		0: "ADDREPLY",
		1: "ADDLIKE",
	}
	CommentMessage_Type_value = map[string]int32{
		"ADDREPLY": 0,
		"ADDLIKE":  1,
	}
)

func (x CommentMessage_Type) Enum() *CommentMessage_Type {
	p := new(CommentMessage_Type)
	*p = x
	return p
}

func (x CommentMessage_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommentMessage_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_job_comment_v1_comment_proto_enumTypes[0].Descriptor()
}

func (CommentMessage_Type) Type() protoreflect.EnumType {
	return &file_job_comment_v1_comment_proto_enumTypes[0]
}

func (x CommentMessage_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommentMessage_Type.Descriptor instead.
func (CommentMessage_Type) EnumDescriptor() ([]byte, []int) {
	return file_job_comment_v1_comment_proto_rawDescGZIP(), []int{0, 0}
}

type CommentMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      CommentMessage_Type `protobuf:"varint,1,opt,name=type,proto3,enum=job.comment.v1.CommentMessage_Type" json:"type,omitempty"`
	ReplyId   int64               `protobuf:"varint,2,opt,name=reply_id,json=replyId,proto3" json:"reply_id,omitempty"`
	Uid       int64               `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	OwnerUid  int64               `protobuf:"varint,4,opt,name=owner_uid,json=ownerUid,proto3" json:"owner_uid,omitempty"`
	SubjectId int64               `protobuf:"varint,5,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	BizType   string              `protobuf:"bytes,6,opt,name=biz_type,json=bizType,proto3" json:"biz_type,omitempty"`
	FloorId   int64               `protobuf:"varint,7,opt,name=floor_id,json=floorId,proto3" json:"floor_id,omitempty"`
	AtUid     int64               `protobuf:"varint,8,opt,name=at_uid,json=atUid,proto3" json:"at_uid,omitempty"`
	Content   string              `protobuf:"bytes,9,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CommentMessage) Reset() {
	*x = CommentMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_comment_v1_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentMessage) ProtoMessage() {}

func (x *CommentMessage) ProtoReflect() protoreflect.Message {
	mi := &file_job_comment_v1_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentMessage.ProtoReflect.Descriptor instead.
func (*CommentMessage) Descriptor() ([]byte, []int) {
	return file_job_comment_v1_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CommentMessage) GetType() CommentMessage_Type {
	if x != nil {
		return x.Type
	}
	return CommentMessage_ADDREPLY
}

func (x *CommentMessage) GetReplyId() int64 {
	if x != nil {
		return x.ReplyId
	}
	return 0
}

func (x *CommentMessage) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *CommentMessage) GetOwnerUid() int64 {
	if x != nil {
		return x.OwnerUid
	}
	return 0
}

func (x *CommentMessage) GetSubjectId() int64 {
	if x != nil {
		return x.SubjectId
	}
	return 0
}

func (x *CommentMessage) GetBizType() string {
	if x != nil {
		return x.BizType
	}
	return ""
}

func (x *CommentMessage) GetFloorId() int64 {
	if x != nil {
		return x.FloorId
	}
	return 0
}

func (x *CommentMessage) GetAtUid() int64 {
	if x != nil {
		return x.AtUid
	}
	return 0
}

func (x *CommentMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type AddLikeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplyId int64 `protobuf:"varint,1,opt,name=reply_id,json=replyId,proto3" json:"reply_id,omitempty"`
	Uid     int64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *AddLikeMessage) Reset() {
	*x = AddLikeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_comment_v1_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLikeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLikeMessage) ProtoMessage() {}

func (x *AddLikeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_job_comment_v1_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLikeMessage.ProtoReflect.Descriptor instead.
func (*AddLikeMessage) Descriptor() ([]byte, []int) {
	return file_job_comment_v1_comment_proto_rawDescGZIP(), []int{1}
}

func (x *AddLikeMessage) GetReplyId() int64 {
	if x != nil {
		return x.ReplyId
	}
	return 0
}

func (x *AddLikeMessage) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type AddReplyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectId int64  `protobuf:"varint,1,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	BizType   string `protobuf:"bytes,2,opt,name=biz_type,json=bizType,proto3" json:"biz_type,omitempty"`
	OwnerUid  int64  `protobuf:"varint,3,opt,name=owner_uid,json=ownerUid,proto3" json:"owner_uid,omitempty"`
	FloorId   int64  `protobuf:"varint,4,opt,name=floor_id,json=floorId,proto3" json:"floor_id,omitempty"`
	AtUid     int64  `protobuf:"varint,5,opt,name=at_uid,json=atUid,proto3" json:"at_uid,omitempty"`
	Content   string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	IsFloor   bool   `protobuf:"varint,7,opt,name=is_floor,json=isFloor,proto3" json:"is_floor,omitempty"`
}

func (x *AddReplyMessage) Reset() {
	*x = AddReplyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_comment_v1_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReplyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReplyMessage) ProtoMessage() {}

func (x *AddReplyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_job_comment_v1_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReplyMessage.ProtoReflect.Descriptor instead.
func (*AddReplyMessage) Descriptor() ([]byte, []int) {
	return file_job_comment_v1_comment_proto_rawDescGZIP(), []int{2}
}

func (x *AddReplyMessage) GetSubjectId() int64 {
	if x != nil {
		return x.SubjectId
	}
	return 0
}

func (x *AddReplyMessage) GetBizType() string {
	if x != nil {
		return x.BizType
	}
	return ""
}

func (x *AddReplyMessage) GetOwnerUid() int64 {
	if x != nil {
		return x.OwnerUid
	}
	return 0
}

func (x *AddReplyMessage) GetFloorId() int64 {
	if x != nil {
		return x.FloorId
	}
	return 0
}

func (x *AddReplyMessage) GetAtUid() int64 {
	if x != nil {
		return x.AtUid
	}
	return 0
}

func (x *AddReplyMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AddReplyMessage) GetIsFloor() bool {
	if x != nil {
		return x.IsFloor
	}
	return false
}

var File_job_comment_v1_comment_proto protoreflect.FileDescriptor

var file_job_comment_v1_comment_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6a, 0x6f, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x6a, 0x6f, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x26,
	0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x88, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x15, 0xea, 0xde, 0x1f,
	0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0xfa, 0x42, 0x04, 0x22, 0x02,
	0x20, 0x00, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a,
	0x08, 0x62, 0x69, 0x7a, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0xea, 0xde, 0x1f, 0x08, 0x62, 0x69, 0x7a, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x07, 0x62,
	0x69, 0x7a, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0c, 0xea, 0xde, 0x1f, 0x08, 0x66, 0x6c,
	0x6f, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x52, 0x07, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x06, 0x61, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x0a, 0xea, 0xde, 0x1f, 0x06, 0x61, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x52, 0x05, 0x61, 0x74, 0x55,
	0x69, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0b, 0xea, 0xde, 0x1f, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x21, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x44, 0x44, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x41, 0x44, 0x44, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x01, 0x22, 0x3d, 0x0a, 0x0e,
	0x41, 0x64, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0xa9, 0x02, 0x0a, 0x0f,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x34, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x15, 0xea, 0xde, 0x1f, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x08, 0x62, 0x69, 0x7a, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xea, 0xde, 0x1f, 0x08, 0x62, 0x69, 0x7a,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x07, 0x62, 0x69, 0x7a, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x08, 0x66,
	0x6c, 0x6f, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0c, 0xea,
	0xde, 0x1f, 0x08, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x52, 0x07, 0x66, 0x6c, 0x6f,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x06, 0x61, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xea, 0xde, 0x1f, 0x06, 0x61, 0x74, 0x5f, 0x75, 0x69, 0x64,
	0x52, 0x05, 0x61, 0x74, 0x55, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xea, 0xde, 0x1f, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x27,
	0x0a, 0x08, 0x69, 0x73, 0x5f, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x0c, 0xea, 0xde, 0x1f, 0x08, 0x69, 0x73, 0x5f, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x52, 0x07,
	0x69, 0x73, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x42, 0x29, 0x5a, 0x27, 0x72, 0x65, 0x64, 0x62, 0x6f,
	0x6f, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x6a, 0x6f, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_comment_v1_comment_proto_rawDescOnce sync.Once
	file_job_comment_v1_comment_proto_rawDescData = file_job_comment_v1_comment_proto_rawDesc
)

func file_job_comment_v1_comment_proto_rawDescGZIP() []byte {
	file_job_comment_v1_comment_proto_rawDescOnce.Do(func() {
		file_job_comment_v1_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_comment_v1_comment_proto_rawDescData)
	})
	return file_job_comment_v1_comment_proto_rawDescData
}

var file_job_comment_v1_comment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_job_comment_v1_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_job_comment_v1_comment_proto_goTypes = []interface{}{
	(CommentMessage_Type)(0), // 0: job.comment.v1.CommentMessage.Type
	(*CommentMessage)(nil),   // 1: job.comment.v1.CommentMessage
	(*AddLikeMessage)(nil),   // 2: job.comment.v1.AddLikeMessage
	(*AddReplyMessage)(nil),  // 3: job.comment.v1.AddReplyMessage
}
var file_job_comment_v1_comment_proto_depIdxs = []int32{
	0, // 0: job.comment.v1.CommentMessage.type:type_name -> job.comment.v1.CommentMessage.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_job_comment_v1_comment_proto_init() }
func file_job_comment_v1_comment_proto_init() {
	if File_job_comment_v1_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_comment_v1_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentMessage); i {
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
		file_job_comment_v1_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddLikeMessage); i {
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
		file_job_comment_v1_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReplyMessage); i {
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
			RawDescriptor: file_job_comment_v1_comment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_job_comment_v1_comment_proto_goTypes,
		DependencyIndexes: file_job_comment_v1_comment_proto_depIdxs,
		EnumInfos:         file_job_comment_v1_comment_proto_enumTypes,
		MessageInfos:      file_job_comment_v1_comment_proto_msgTypes,
	}.Build()
	File_job_comment_v1_comment_proto = out.File
	file_job_comment_v1_comment_proto_rawDesc = nil
	file_job_comment_v1_comment_proto_goTypes = nil
	file_job_comment_v1_comment_proto_depIdxs = nil
}
