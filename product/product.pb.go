// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0
// source: product/product.proto

package productproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ProductsInfoRequest Message
type ProductsInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RelatedId     string                 `protobuf:"bytes,1,opt,name=related_id,json=relatedId,proto3" json:"related_id,omitempty"`
	RelatedType   string                 `protobuf:"bytes,2,opt,name=related_type,json=relatedType,proto3" json:"related_type,omitempty"` // related_type (Example: "products", "services", ...)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductsInfoRequest) Reset() {
	*x = ProductsInfoRequest{}
	mi := &file_product_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductsInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsInfoRequest) ProtoMessage() {}

func (x *ProductsInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsInfoRequest.ProtoReflect.Descriptor instead.
func (*ProductsInfoRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductsInfoRequest) GetRelatedId() string {
	if x != nil {
		return x.RelatedId
	}
	return ""
}

func (x *ProductsInfoRequest) GetRelatedType() string {
	if x != nil {
		return x.RelatedType
	}
	return ""
}

// GetProductsInfoRequest message
type GetProductsInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*ProductsInfoRequest `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"` // Danh sách các loại sản phẩm/dịch vụ
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductsInfoRequest) Reset() {
	*x = GetProductsInfoRequest{}
	mi := &file_product_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductsInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsInfoRequest) ProtoMessage() {}

func (x *GetProductsInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsInfoRequest.ProtoReflect.Descriptor instead.
func (*GetProductsInfoRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductsInfoRequest) GetProducts() []*ProductsInfoRequest {
	if x != nil {
		return x.Products
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Slug          string                 `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Image         string                 `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	RelatedType   string                 `protobuf:"bytes,5,opt,name=related_type,json=relatedType,proto3" json:"related_type,omitempty"`
	Price         float64                `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_product_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[2]
	if x != nil {
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
	return file_product_product_proto_rawDescGZIP(), []int{2}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Product) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Product) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Product) GetRelatedType() string {
	if x != nil {
		return x.RelatedType
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type GetProductsInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*Product             `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	StatusCode    int32                  `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductsInfoResponse) Reset() {
	*x = GetProductsInfoResponse{}
	mi := &file_product_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductsInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsInfoResponse) ProtoMessage() {}

func (x *GetProductsInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsInfoResponse.ProtoReflect.Descriptor instead.
func (*GetProductsInfoResponse) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{3}
}

func (x *GetProductsInfoResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *GetProductsInfoResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetProductsInfoResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// ClearCartAfterCheckoutRequest message
type ClearCartAfterCheckoutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          string                 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Products      []*ProductsInfoRequest `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"` // Danh sách các loại sản phẩm/dịch vụ
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClearCartAfterCheckoutRequest) Reset() {
	*x = ClearCartAfterCheckoutRequest{}
	mi := &file_product_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClearCartAfterCheckoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearCartAfterCheckoutRequest) ProtoMessage() {}

func (x *ClearCartAfterCheckoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearCartAfterCheckoutRequest.ProtoReflect.Descriptor instead.
func (*ClearCartAfterCheckoutRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{4}
}

func (x *ClearCartAfterCheckoutRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ClearCartAfterCheckoutRequest) GetProducts() []*ProductsInfoRequest {
	if x != nil {
		return x.Products
	}
	return nil
}

type ClearCartAfterCheckoutResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusCode    int32                  `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClearCartAfterCheckoutResponse) Reset() {
	*x = ClearCartAfterCheckoutResponse{}
	mi := &file_product_product_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClearCartAfterCheckoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearCartAfterCheckoutResponse) ProtoMessage() {}

func (x *ClearCartAfterCheckoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearCartAfterCheckoutResponse.ProtoReflect.Descriptor instead.
func (*ClearCartAfterCheckoutResponse) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{5}
}

func (x *ClearCartAfterCheckoutResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ClearCartAfterCheckoutResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_product_product_proto protoreflect.FileDescriptor

var file_product_product_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x22, 0x57, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x52, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x92, 0x01,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x7e, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x6d, 0x0a, 0x1d, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x22, 0x57, 0x0a, 0x1e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x41, 0x66,
	0x74, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xd1, 0x01, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x16, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x26, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72,
	0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x61,
	0x6e, 0x67, 0x33, 0x62, 0x65, 0x72, 0x31, 0x32, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x3b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_product_product_proto_rawDescOnce sync.Once
	file_product_product_proto_rawDescData []byte
)

func file_product_product_proto_rawDescGZIP() []byte {
	file_product_product_proto_rawDescOnce.Do(func() {
		file_product_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_product_product_proto_rawDesc), len(file_product_product_proto_rawDesc)))
	})
	return file_product_product_proto_rawDescData
}

var file_product_product_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_product_product_proto_goTypes = []any{
	(*ProductsInfoRequest)(nil),            // 0: product.ProductsInfoRequest
	(*GetProductsInfoRequest)(nil),         // 1: product.GetProductsInfoRequest
	(*Product)(nil),                        // 2: product.Product
	(*GetProductsInfoResponse)(nil),        // 3: product.GetProductsInfoResponse
	(*ClearCartAfterCheckoutRequest)(nil),  // 4: product.ClearCartAfterCheckoutRequest
	(*ClearCartAfterCheckoutResponse)(nil), // 5: product.ClearCartAfterCheckoutResponse
}
var file_product_product_proto_depIdxs = []int32{
	0, // 0: product.GetProductsInfoRequest.products:type_name -> product.ProductsInfoRequest
	2, // 1: product.GetProductsInfoResponse.products:type_name -> product.Product
	0, // 2: product.ClearCartAfterCheckoutRequest.products:type_name -> product.ProductsInfoRequest
	1, // 3: product.ProductService.GetProductsInfo:input_type -> product.GetProductsInfoRequest
	4, // 4: product.ProductService.ClearCartAfterCheckout:input_type -> product.ClearCartAfterCheckoutRequest
	3, // 5: product.ProductService.GetProductsInfo:output_type -> product.GetProductsInfoResponse
	5, // 6: product.ProductService.ClearCartAfterCheckout:output_type -> product.ClearCartAfterCheckoutResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_product_product_proto_init() }
func file_product_product_proto_init() {
	if File_product_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_product_product_proto_rawDesc), len(file_product_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_product_proto_goTypes,
		DependencyIndexes: file_product_product_proto_depIdxs,
		MessageInfos:      file_product_product_proto_msgTypes,
	}.Build()
	File_product_product_proto = out.File
	file_product_product_proto_goTypes = nil
	file_product_product_proto_depIdxs = nil
}
