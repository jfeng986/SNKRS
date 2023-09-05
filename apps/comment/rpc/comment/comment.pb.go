// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: comment.proto

package comment

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

type CommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId int64 `protobuf:"varint,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	Cursor   int64 `protobuf:"varint,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Ps       int32 `protobuf:"varint,3,opt,name=ps,proto3" json:"ps,omitempty"`
}

func (x *CommentsRequest) Reset() {
	*x = CommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentsRequest) ProtoMessage() {}

func (x *CommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentsRequest.ProtoReflect.Descriptor instead.
func (*CommentsRequest) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CommentsRequest) GetTargetId() int64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *CommentsRequest) GetCursor() int64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *CommentsRequest) GetPs() int32 {
	if x != nil {
		return x.Ps
	}
	return 0
}

type CommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments    []*CommentItem `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	IsEnd       bool           `protobuf:"varint,2,opt,name=is_end,json=isEnd,proto3" json:"is_end,omitempty"`
	CreatedTime int64          `protobuf:"varint,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
}

func (x *CommentsResponse) Reset() {
	*x = CommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentsResponse) ProtoMessage() {}

func (x *CommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentsResponse.ProtoReflect.Descriptor instead.
func (*CommentsResponse) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentsResponse) GetComments() []*CommentItem {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *CommentsResponse) GetIsEnd() bool {
	if x != nil {
		return x.IsEnd
	}
	return false
}

func (x *CommentsResponse) GetCreatedTime() int64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

type CommentItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Business      string `protobuf:"bytes,2,opt,name=business,proto3" json:"business,omitempty"`
	TargetId      int64  `protobuf:"varint,3,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	ReplyUserId   int64  `protobuf:"varint,4,opt,name=reply_user_id,json=replyUserId,proto3" json:"reply_user_id,omitempty"`
	BeReplyUserId int64  `protobuf:"varint,5,opt,name=be_reply_user_id,json=beReplyUserId,proto3" json:"be_reply_user_id,omitempty"`
	ParentId      int64  `protobuf:"varint,6,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Content       string `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	Image         string `protobuf:"bytes,8,opt,name=image,proto3" json:"image,omitempty"`
	CreateTime    int64  `protobuf:"varint,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime    int64  `protobuf:"varint,10,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *CommentItem) Reset() {
	*x = CommentItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentItem) ProtoMessage() {}

func (x *CommentItem) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentItem.ProtoReflect.Descriptor instead.
func (*CommentItem) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{2}
}

func (x *CommentItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentItem) GetBusiness() string {
	if x != nil {
		return x.Business
	}
	return ""
}

func (x *CommentItem) GetTargetId() int64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *CommentItem) GetReplyUserId() int64 {
	if x != nil {
		return x.ReplyUserId
	}
	return 0
}

func (x *CommentItem) GetBeReplyUserId() int64 {
	if x != nil {
		return x.BeReplyUserId
	}
	return 0
}

func (x *CommentItem) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *CommentItem) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommentItem) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CommentItem) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *CommentItem) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

var File_comment_proto protoreflect.FileDescriptor

var file_comment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x56, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x70, 0x73,
	0x22, 0x7e, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x45, 0x6e, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0xb2, 0x02, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a,
	0x10, 0x62, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x4a, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x3f, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_proto_rawDescOnce sync.Once
	file_comment_proto_rawDescData = file_comment_proto_rawDesc
)

func file_comment_proto_rawDescGZIP() []byte {
	file_comment_proto_rawDescOnce.Do(func() {
		file_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_proto_rawDescData)
	})
	return file_comment_proto_rawDescData
}

var file_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_comment_proto_goTypes = []interface{}{
	(*CommentsRequest)(nil),  // 0: comment.CommentsRequest
	(*CommentsResponse)(nil), // 1: comment.CommentsResponse
	(*CommentItem)(nil),      // 2: comment.CommentItem
}
var file_comment_proto_depIdxs = []int32{
	2, // 0: comment.CommentsResponse.comments:type_name -> comment.CommentItem
	0, // 1: comment.Comment.Comments:input_type -> comment.CommentsRequest
	1, // 2: comment.Comment.Comments:output_type -> comment.CommentsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_comment_proto_init() }
func file_comment_proto_init() {
	if File_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentsRequest); i {
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
		file_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentsResponse); i {
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
		file_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentItem); i {
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
			RawDescriptor: file_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comment_proto_goTypes,
		DependencyIndexes: file_comment_proto_depIdxs,
		MessageInfos:      file_comment_proto_msgTypes,
	}.Build()
	File_comment_proto = out.File
	file_comment_proto_rawDesc = nil
	file_comment_proto_goTypes = nil
	file_comment_proto_depIdxs = nil
}
