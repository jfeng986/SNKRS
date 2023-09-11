// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: reserve.proto

package reserve

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

type ReserveOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId int64 `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *ReserveOrderRequest) Reset() {
	*x = ReserveOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reserve_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReserveOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveOrderRequest) ProtoMessage() {}

func (x *ReserveOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reserve_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveOrderRequest.ProtoReflect.Descriptor instead.
func (*ReserveOrderRequest) Descriptor() ([]byte, []int) {
	return file_reserve_proto_rawDescGZIP(), []int{0}
}

func (x *ReserveOrderRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReserveOrderRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type ReserveOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReserveOrderResponse) Reset() {
	*x = ReserveOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reserve_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReserveOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveOrderResponse) ProtoMessage() {}

func (x *ReserveOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reserve_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveOrderResponse.ProtoReflect.Descriptor instead.
func (*ReserveOrderResponse) Descriptor() ([]byte, []int) {
	return file_reserve_proto_rawDescGZIP(), []int{1}
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId  int64  `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc       string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Image      string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Stock      int64  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	CreateTime int64  `protobuf:"varint,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reserve_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_reserve_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_reserve_proto_rawDescGZIP(), []int{2}
}

func (x *Product) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Product) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Product) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *Product) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

var File_reserve_proto protoreflect.FileDescriptor

var file_reserve_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x22, 0x4d, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x9d, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32,
	0x5d, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_reserve_proto_rawDescOnce sync.Once
	file_reserve_proto_rawDescData = file_reserve_proto_rawDesc
)

func file_reserve_proto_rawDescGZIP() []byte {
	file_reserve_proto_rawDescOnce.Do(func() {
		file_reserve_proto_rawDescData = protoimpl.X.CompressGZIP(file_reserve_proto_rawDescData)
	})
	return file_reserve_proto_rawDescData
}

var file_reserve_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_reserve_proto_goTypes = []interface{}{
	(*ReserveOrderRequest)(nil),  // 0: reserve.ReserveOrderRequest
	(*ReserveOrderResponse)(nil), // 1: reserve.ReserveOrderResponse
	(*Product)(nil),              // 2: reserve.Product
}
var file_reserve_proto_depIdxs = []int32{
	0, // 0: reserve.ReserveService.ReserveOrder:input_type -> reserve.ReserveOrderRequest
	1, // 1: reserve.ReserveService.ReserveOrder:output_type -> reserve.ReserveOrderResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reserve_proto_init() }
func file_reserve_proto_init() {
	if File_reserve_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reserve_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReserveOrderRequest); i {
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
		file_reserve_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReserveOrderResponse); i {
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
		file_reserve_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_reserve_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reserve_proto_goTypes,
		DependencyIndexes: file_reserve_proto_depIdxs,
		MessageInfos:      file_reserve_proto_msgTypes,
	}.Build()
	File_reserve_proto = out.File
	file_reserve_proto_rawDesc = nil
	file_reserve_proto_goTypes = nil
	file_reserve_proto_depIdxs = nil
}
