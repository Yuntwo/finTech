// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: merchant.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type JoinMerchantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopName string `protobuf:"bytes,1,opt,name=shop_name,json=shopName,proto3" json:"shop_name,omitempty"`
	ShopLogo string `protobuf:"bytes,2,opt,name=shop_logo,json=shopLogo,proto3" json:"shop_logo,omitempty"`
	Mobile   string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Address  string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Remark   string `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	UserId   int64  `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *JoinMerchantRequest) Reset() {
	*x = JoinMerchantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchants_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinMerchantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinMerchantRequest) ProtoMessage() {}

func (x *JoinMerchantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_merchants_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinMerchantRequest.ProtoReflect.Descriptor instead.
func (*JoinMerchantRequest) Descriptor() ([]byte, []int) {
	return file_merchants_proto_rawDescGZIP(), []int{0}
}

func (x *JoinMerchantRequest) GetShopName() string {
	if x != nil {
		return x.ShopName
	}
	return ""
}

func (x *JoinMerchantRequest) GetShopLogo() string {
	if x != nil {
		return x.ShopLogo
	}
	return ""
}

func (x *JoinMerchantRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *JoinMerchantRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *JoinMerchantRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *JoinMerchantRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type JoinMerchantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *JoinMerchantResponse) Reset() {
	*x = JoinMerchantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchants_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinMerchantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinMerchantResponse) ProtoMessage() {}

func (x *JoinMerchantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_merchants_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinMerchantResponse.ProtoReflect.Descriptor instead.
func (*JoinMerchantResponse) Descriptor() ([]byte, []int) {
	return file_merchants_proto_rawDescGZIP(), []int{1}
}

func (x *JoinMerchantResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type UpdateMerchantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopName string `protobuf:"bytes,1,opt,name=shop_name,json=shopName,proto3" json:"shop_name,omitempty"`
	ShopLogo string `protobuf:"bytes,2,opt,name=shop_logo,json=shopLogo,proto3" json:"shop_logo,omitempty"`
	Mobile   string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Address  string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Remark   string `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	UserId   int64  `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ShopId   int64  `protobuf:"varint,7,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *UpdateMerchantRequest) Reset() {
	*x = UpdateMerchantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchants_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMerchantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMerchantRequest) ProtoMessage() {}

func (x *UpdateMerchantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_merchants_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMerchantRequest.ProtoReflect.Descriptor instead.
func (*UpdateMerchantRequest) Descriptor() ([]byte, []int) {
	return file_merchants_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateMerchantRequest) GetShopName() string {
	if x != nil {
		return x.ShopName
	}
	return ""
}

func (x *UpdateMerchantRequest) GetShopLogo() string {
	if x != nil {
		return x.ShopLogo
	}
	return ""
}

func (x *UpdateMerchantRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UpdateMerchantRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateMerchantRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *UpdateMerchantRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateMerchantRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

type UpdateMerchantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateMerchantResponse) Reset() {
	*x = UpdateMerchantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchants_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMerchantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMerchantResponse) ProtoMessage() {}

func (x *UpdateMerchantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_merchants_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMerchantResponse.ProtoReflect.Descriptor instead.
func (*UpdateMerchantResponse) Descriptor() ([]byte, []int) {
	return file_merchants_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateMerchantResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type CloseMerchantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId int64 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *CloseMerchantRequest) Reset() {
	*x = CloseMerchantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchants_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseMerchantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseMerchantRequest) ProtoMessage() {}

func (x *CloseMerchantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_merchants_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseMerchantRequest.ProtoReflect.Descriptor instead.
func (*CloseMerchantRequest) Descriptor() ([]byte, []int) {
	return file_merchants_proto_rawDescGZIP(), []int{4}
}

func (x *CloseMerchantRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

type CloseMerchantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CloseMerchantResponse) Reset() {
	*x = CloseMerchantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchants_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseMerchantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseMerchantResponse) ProtoMessage() {}

func (x *CloseMerchantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_merchants_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseMerchantResponse.ProtoReflect.Descriptor instead.
func (*CloseMerchantResponse) Descriptor() ([]byte, []int) {
	return file_merchants_proto_rawDescGZIP(), []int{5}
}

func (x *CloseMerchantResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type GetMerchantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetMerchantRequest) Reset() {
	*x = GetMerchantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchants_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMerchantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMerchantRequest) ProtoMessage() {}

func (x *GetMerchantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_merchants_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMerchantRequest.ProtoReflect.Descriptor instead.
func (*GetMerchantRequest) Descriptor() ([]byte, []int) {
	return file_merchants_proto_rawDescGZIP(), []int{6}
}

func (x *GetMerchantRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetMerchantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId   int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ShopName string `protobuf:"bytes,3,opt,name=shop_name,json=shopName,proto3" json:"shop_name,omitempty"`
	ShopLogo string `protobuf:"bytes,4,opt,name=shop_logo,json=shopLogo,proto3" json:"shop_logo,omitempty"`
	Mobile   string `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Address  string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Remark   string `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
	Sort     int32  `protobuf:"varint,8,opt,name=sort,proto3" json:"sort,omitempty"`
	IsHide   int32  `protobuf:"varint,9,opt,name=is_hide,json=isHide,proto3" json:"is_hide,omitempty"`
	Status   int64  `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetMerchantResponse) Reset() {
	*x = GetMerchantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchants_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMerchantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMerchantResponse) ProtoMessage() {}

func (x *GetMerchantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_merchants_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMerchantResponse.ProtoReflect.Descriptor instead.
func (*GetMerchantResponse) Descriptor() ([]byte, []int) {
	return file_merchants_proto_rawDescGZIP(), []int{7}
}

func (x *GetMerchantResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetMerchantResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetMerchantResponse) GetShopName() string {
	if x != nil {
		return x.ShopName
	}
	return ""
}

func (x *GetMerchantResponse) GetShopLogo() string {
	if x != nil {
		return x.ShopLogo
	}
	return ""
}

func (x *GetMerchantResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *GetMerchantResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetMerchantResponse) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *GetMerchantResponse) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *GetMerchantResponse) GetIsHide() int32 {
	if x != nil {
		return x.IsHide
	}
	return 0
}

func (x *GetMerchantResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_merchants_proto protoreflect.FileDescriptor

var file_merchants_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x6d, 0x61, 0x6c, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x13, 0x4a, 0x6f, 0x69, 0x6e, 0x4d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68,
	0x6f, 0x70, 0x5f, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x68, 0x6f, 0x70, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x14, 0x4a, 0x6f,
	0x69, 0x6e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xcd, 0x01, 0x0a, 0x15, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x16, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2f, 0x0a, 0x14,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0x2f, 0x0a,
	0x15, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x87, 0x02,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x68, 0x6f, 0x70, 0x5f, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x68, 0x6f, 0x70, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x68, 0x69, 0x64,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x73, 0x48, 0x69, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xa5, 0x03, 0x0a, 0x09, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x63, 0x0a, 0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x73, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x29, 0x2e, 0x6d,
	0x61, 0x6c, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x6d,
	0x61, 0x6c, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x61, 0x6c, 0x6c, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x06, 0x5a, 0x04, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_merchants_proto_rawDescOnce sync.Once
	file_merchants_proto_rawDescData = file_merchants_proto_rawDesc
)

func file_merchants_proto_rawDescGZIP() []byte {
	file_merchants_proto_rawDescOnce.Do(func() {
		file_merchants_proto_rawDescData = protoimpl.X.CompressGZIP(file_merchants_proto_rawDescData)
	})
	return file_merchants_proto_rawDescData
}

var file_merchants_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_merchants_proto_goTypes = []interface{}{
	(*JoinMerchantRequest)(nil),    // 0: mall.rpc.merchant.JoinMerchantRequest
	(*JoinMerchantResponse)(nil),   // 1: mall.rpc.merchant.JoinMerchantResponse
	(*UpdateMerchantRequest)(nil),  // 2: mall.rpc.merchant.UpdateMerchantRequest
	(*UpdateMerchantResponse)(nil), // 3: mall.rpc.merchant.UpdateMerchantResponse
	(*CloseMerchantRequest)(nil),   // 4: mall.rpc.merchant.CloseMerchantRequest
	(*CloseMerchantResponse)(nil),  // 5: mall.rpc.merchant.CloseMerchantResponse
	(*GetMerchantRequest)(nil),     // 6: mall.rpc.merchant.GetMerchantRequest
	(*GetMerchantResponse)(nil),    // 7: mall.rpc.merchant.GetMerchantResponse
}
var file_merchants_proto_depIdxs = []int32{
	0, // 0: mall.rpc.merchant.Merchants.JoinMerchant:input_type -> mall.rpc.merchant.JoinMerchantRequest
	2, // 1: mall.rpc.merchant.Merchants.UpdateMerchant:input_type -> mall.rpc.merchant.UpdateMerchantRequest
	4, // 2: mall.rpc.merchant.Merchants.CloseMerchant:input_type -> mall.rpc.merchant.CloseMerchantRequest
	6, // 3: mall.rpc.merchant.Merchants.GetMerchant:input_type -> mall.rpc.merchant.GetMerchantRequest
	1, // 4: mall.rpc.merchant.Merchants.JoinMerchant:output_type -> mall.rpc.merchant.JoinMerchantResponse
	3, // 5: mall.rpc.merchant.Merchants.UpdateMerchant:output_type -> mall.rpc.merchant.UpdateMerchantResponse
	5, // 6: mall.rpc.merchant.Merchants.CloseMerchant:output_type -> mall.rpc.merchant.CloseMerchantResponse
	7, // 7: mall.rpc.merchant.Merchants.GetMerchant:output_type -> mall.rpc.merchant.GetMerchantResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_merchants_proto_init() }
func file_merchants_proto_init() {
	if File_merchants_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_merchants_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinMerchantRequest); i {
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
		file_merchants_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinMerchantResponse); i {
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
		file_merchants_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMerchantRequest); i {
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
		file_merchants_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMerchantResponse); i {
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
		file_merchants_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseMerchantRequest); i {
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
		file_merchants_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseMerchantResponse); i {
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
		file_merchants_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMerchantRequest); i {
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
		file_merchants_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMerchantResponse); i {
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
			RawDescriptor: file_merchants_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_merchants_proto_goTypes,
		DependencyIndexes: file_merchants_proto_depIdxs,
		MessageInfos:      file_merchants_proto_msgTypes,
	}.Build()
	File_merchants_proto = out.File
	file_merchants_proto_rawDesc = nil
	file_merchants_proto_goTypes = nil
	file_merchants_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MerchantsClient is the client API for Merchants service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MerchantsClient interface {
	// 加入商户
	JoinMerchant(ctx context.Context, in *JoinMerchantRequest, opts ...grpc.CallOption) (*JoinMerchantResponse, error)
	// 更新商户信息
	UpdateMerchant(ctx context.Context, in *UpdateMerchantRequest, opts ...grpc.CallOption) (*UpdateMerchantResponse, error)
	// 关闭商户
	CloseMerchant(ctx context.Context, in *CloseMerchantRequest, opts ...grpc.CallOption) (*CloseMerchantResponse, error)
	// 查询商户
	GetMerchant(ctx context.Context, in *GetMerchantRequest, opts ...grpc.CallOption) (*GetMerchantResponse, error)
}

type merchantsClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantsClient(cc grpc.ClientConnInterface) MerchantsClient {
	return &merchantsClient{cc}
}

func (c *merchantsClient) JoinMerchant(ctx context.Context, in *JoinMerchantRequest, opts ...grpc.CallOption) (*JoinMerchantResponse, error) {
	out := new(JoinMerchantResponse)
	err := c.cc.Invoke(ctx, "/mall.rpc.merchant.Merchants/JoinMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantsClient) UpdateMerchant(ctx context.Context, in *UpdateMerchantRequest, opts ...grpc.CallOption) (*UpdateMerchantResponse, error) {
	out := new(UpdateMerchantResponse)
	err := c.cc.Invoke(ctx, "/mall.rpc.merchant.Merchants/UpdateMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantsClient) CloseMerchant(ctx context.Context, in *CloseMerchantRequest, opts ...grpc.CallOption) (*CloseMerchantResponse, error) {
	out := new(CloseMerchantResponse)
	err := c.cc.Invoke(ctx, "/mall.rpc.merchant.Merchants/CloseMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantsClient) GetMerchant(ctx context.Context, in *GetMerchantRequest, opts ...grpc.CallOption) (*GetMerchantResponse, error) {
	out := new(GetMerchantResponse)
	err := c.cc.Invoke(ctx, "/mall.rpc.merchant.Merchants/GetMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantsServer is the server API for Merchants service.
type MerchantsServer interface {
	// 加入商户
	JoinMerchant(context.Context, *JoinMerchantRequest) (*JoinMerchantResponse, error)
	// 更新商户信息
	UpdateMerchant(context.Context, *UpdateMerchantRequest) (*UpdateMerchantResponse, error)
	// 关闭商户
	CloseMerchant(context.Context, *CloseMerchantRequest) (*CloseMerchantResponse, error)
	// 查询商户
	GetMerchant(context.Context, *GetMerchantRequest) (*GetMerchantResponse, error)
}

// UnimplementedMerchantsServer can be embedded to have forward compatible implementations.
type UnimplementedMerchantsServer struct {
}

func (*UnimplementedMerchantsServer) JoinMerchant(context.Context, *JoinMerchantRequest) (*JoinMerchantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinMerchant not implemented")
}
func (*UnimplementedMerchantsServer) UpdateMerchant(context.Context, *UpdateMerchantRequest) (*UpdateMerchantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMerchant not implemented")
}
func (*UnimplementedMerchantsServer) CloseMerchant(context.Context, *CloseMerchantRequest) (*CloseMerchantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseMerchant not implemented")
}
func (*UnimplementedMerchantsServer) GetMerchant(context.Context, *GetMerchantRequest) (*GetMerchantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMerchant not implemented")
}

func RegisterMerchantsServer(s *grpc.Server, srv MerchantsServer) {
	s.RegisterService(&_Merchants_serviceDesc, srv)
}

func _Merchants_JoinMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinMerchantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantsServer).JoinMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.rpc.merchant.Merchants/JoinMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantsServer).JoinMerchant(ctx, req.(*JoinMerchantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merchants_UpdateMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMerchantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantsServer).UpdateMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.rpc.merchant.Merchants/UpdateMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantsServer).UpdateMerchant(ctx, req.(*UpdateMerchantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merchants_CloseMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseMerchantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantsServer).CloseMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.rpc.merchant.Merchants/CloseMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantsServer).CloseMerchant(ctx, req.(*CloseMerchantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merchants_GetMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMerchantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantsServer).GetMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.rpc.merchant.Merchants/GetMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantsServer).GetMerchant(ctx, req.(*GetMerchantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Merchants_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mall.rpc.merchant.Merchants",
	HandlerType: (*MerchantsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinMerchant",
			Handler:    _Merchants_JoinMerchant_Handler,
		},
		{
			MethodName: "UpdateMerchant",
			Handler:    _Merchants_UpdateMerchant_Handler,
		},
		{
			MethodName: "CloseMerchant",
			Handler:    _Merchants_CloseMerchant_Handler,
		},
		{
			MethodName: "GetMerchant",
			Handler:    _Merchants_GetMerchant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "merchant.proto",
}
