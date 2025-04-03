// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0
// source: product/product.proto

package productproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Cart Message
type CartItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`       // related_type (Example: "product", "service", ...)
	Values        []string               `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"` // related_id
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartItemRequest) Reset() {
	*x = CartItemRequest{}
	mi := &file_product_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemRequest) ProtoMessage() {}

func (x *CartItemRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CartItemRequest.ProtoReflect.Descriptor instead.
func (*CartItemRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{0}
}

func (x *CartItemRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CartItemRequest) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// GetProductsInCartRequest message
type GetProductsInCartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          []*CartItemRequest     `protobuf:"bytes,1,rep,name=cart,proto3" json:"cart,omitempty"` // Danh sách các loại sản phẩm/dịch vụ
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductsInCartRequest) Reset() {
	*x = GetProductsInCartRequest{}
	mi := &file_product_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductsInCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsInCartRequest) ProtoMessage() {}

func (x *GetProductsInCartRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetProductsInCartRequest.ProtoReflect.Descriptor instead.
func (*GetProductsInCartRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductsInCartRequest) GetCart() []*CartItemRequest {
	if x != nil {
		return x.Cart
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Image         string                 `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	RelatedType   string                 `protobuf:"bytes,4,opt,name=related_type,json=relatedType,proto3" json:"related_type,omitempty"`
	Price         float64                `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
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

type GetProductsInCartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*Product             `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	StatusCode    int32                  `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductsInCartResponse) Reset() {
	*x = GetProductsInCartResponse{}
	mi := &file_product_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductsInCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsInCartResponse) ProtoMessage() {}

func (x *GetProductsInCartResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetProductsInCartResponse.ProtoReflect.Descriptor instead.
func (*GetProductsInCartResponse) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{3}
}

func (x *GetProductsInCartResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *GetProductsInCartResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetProductsInCartResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Category struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Category) Reset() {
	*x = Category{}
	mi := &file_product_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{4}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Category) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// GetCategoriesRequest mesage
type GetCategoriesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Categories    []*Category            `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCategoriesRequest) Reset() {
	*x = GetCategoriesRequest{}
	mi := &file_product_product_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoriesRequest) ProtoMessage() {}

func (x *GetCategoriesRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCategoriesRequest.ProtoReflect.Descriptor instead.
func (*GetCategoriesRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{5}
}

func (x *GetCategoriesRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *GetCategoriesRequest) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

type GetCategoriesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CategoriesId  []string               `protobuf:"bytes,1,rep,name=categories_id,json=categoriesId,proto3" json:"categories_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCategoriesResponse) Reset() {
	*x = GetCategoriesResponse{}
	mi := &file_product_product_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoriesResponse) ProtoMessage() {}

func (x *GetCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoriesResponse.ProtoReflect.Descriptor instead.
func (*GetCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{6}
}

func (x *GetCategoriesResponse) GetCategoriesId() []string {
	if x != nil {
		return x.CategoriesId
	}
	return nil
}

type ProductRecommend struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Categories    []string               `protobuf:"bytes,3,rep,name=categories,proto3" json:"categories,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Pricing       float64                `protobuf:"fixed64,5,opt,name=pricing,proto3" json:"pricing,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductRecommend) Reset() {
	*x = ProductRecommend{}
	mi := &file_product_product_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductRecommend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRecommend) ProtoMessage() {}

func (x *ProductRecommend) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRecommend.ProtoReflect.Descriptor instead.
func (*ProductRecommend) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{7}
}

func (x *ProductRecommend) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductRecommend) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProductRecommend) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *ProductRecommend) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ProductRecommend) GetPricing() float64 {
	if x != nil {
		return x.Pricing
	}
	return 0
}

type ClickDetail struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ClickTime     int32                  `protobuf:"varint,2,opt,name=click_time,json=clickTime,proto3" json:"click_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClickDetail) Reset() {
	*x = ClickDetail{}
	mi := &file_product_product_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClickDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickDetail) ProtoMessage() {}

func (x *ClickDetail) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickDetail.ProtoReflect.Descriptor instead.
func (*ClickDetail) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{8}
}

func (x *ClickDetail) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ClickDetail) GetClickTime() int32 {
	if x != nil {
		return x.ClickTime
	}
	return 0
}

type ViewProduct struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ViewTime      int32                  `protobuf:"varint,2,opt,name=view_time,json=viewTime,proto3" json:"view_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ViewProduct) Reset() {
	*x = ViewProduct{}
	mi := &file_product_product_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ViewProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewProduct) ProtoMessage() {}

func (x *ViewProduct) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewProduct.ProtoReflect.Descriptor instead.
func (*ViewProduct) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{9}
}

func (x *ViewProduct) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ViewProduct) GetViewTime() int32 {
	if x != nil {
		return x.ViewTime
	}
	return 0
}

type GetRecommendProductRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Categories     []*Category            `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	Products       []*ProductRecommend    `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	ClickDetails   []*ClickDetail         `protobuf:"bytes,3,rep,name=click_details,json=clickDetails,proto3" json:"click_details,omitempty"`
	ViewProducts   []*ViewProduct         `protobuf:"bytes,4,rep,name=view_products,json=viewProducts,proto3" json:"view_products,omitempty"`
	BoughtProducts []string               `protobuf:"bytes,5,rep,name=bought_products,json=boughtProducts,proto3" json:"bought_products,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetRecommendProductRequest) Reset() {
	*x = GetRecommendProductRequest{}
	mi := &file_product_product_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecommendProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecommendProductRequest) ProtoMessage() {}

func (x *GetRecommendProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecommendProductRequest.ProtoReflect.Descriptor instead.
func (*GetRecommendProductRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{10}
}

func (x *GetRecommendProductRequest) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *GetRecommendProductRequest) GetProducts() []*ProductRecommend {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *GetRecommendProductRequest) GetClickDetails() []*ClickDetail {
	if x != nil {
		return x.ClickDetails
	}
	return nil
}

func (x *GetRecommendProductRequest) GetViewProducts() []*ViewProduct {
	if x != nil {
		return x.ViewProducts
	}
	return nil
}

func (x *GetRecommendProductRequest) GetBoughtProducts() []string {
	if x != nil {
		return x.BoughtProducts
	}
	return nil
}

type GetRecommendProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductsId    []string               `protobuf:"bytes,1,rep,name=products_id,json=productsId,proto3" json:"products_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRecommendProductResponse) Reset() {
	*x = GetRecommendProductResponse{}
	mi := &file_product_product_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRecommendProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecommendProductResponse) ProtoMessage() {}

func (x *GetRecommendProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecommendProductResponse.ProtoReflect.Descriptor instead.
func (*GetRecommendProductResponse) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{11}
}

func (x *GetRecommendProductResponse) GetProductsId() []string {
	if x != nil {
		return x.ProductsId
	}
	return nil
}

var File_product_product_proto protoreflect.FileDescriptor

var file_product_product_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3b, 0x0a, 0x0f, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x48,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x49, 0x6e, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x63, 0x61,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x7e, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x49, 0x6e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x52, 0x0a, 0x08, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x5f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x22, 0x3c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x49, 0x64, 0x22, 0xad,
	0x01, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x22, 0x4b,
	0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6c, 0x69, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x0b, 0x56,
	0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x76, 0x69,
	0x65, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xa5, 0x02, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12,
	0x39, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0c, 0x63, 0x6c,
	0x69, 0x63, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x39, 0x0a, 0x0d, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x56, 0x69, 0x65, 0x77,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x0c, 0x76, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x62, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x3e,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x49, 0x64, 0x32, 0xa0,
	0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x49, 0x6e, 0x43, 0x61, 0x72, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x49, 0x6e, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x49,
	0x6e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x60, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x6f, 0x61, 0x6e, 0x67, 0x33, 0x62, 0x65, 0x72, 0x31, 0x32, 0x33, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x3b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_product_product_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_product_product_proto_goTypes = []any{
	(*CartItemRequest)(nil),             // 0: product.CartItemRequest
	(*GetProductsInCartRequest)(nil),    // 1: product.GetProductsInCartRequest
	(*Product)(nil),                     // 2: product.Product
	(*GetProductsInCartResponse)(nil),   // 3: product.GetProductsInCartResponse
	(*Category)(nil),                    // 4: product.Category
	(*GetCategoriesRequest)(nil),        // 5: product.GetCategoriesRequest
	(*GetCategoriesResponse)(nil),       // 6: product.GetCategoriesResponse
	(*ProductRecommend)(nil),            // 7: product.ProductRecommend
	(*ClickDetail)(nil),                 // 8: product.ClickDetail
	(*ViewProduct)(nil),                 // 9: product.ViewProduct
	(*GetRecommendProductRequest)(nil),  // 10: product.GetRecommendProductRequest
	(*GetRecommendProductResponse)(nil), // 11: product.GetRecommendProductResponse
	(*timestamppb.Timestamp)(nil),       // 12: google.protobuf.Timestamp
}
var file_product_product_proto_depIdxs = []int32{
	0,  // 0: product.GetProductsInCartRequest.cart:type_name -> product.CartItemRequest
	2,  // 1: product.GetProductsInCartResponse.products:type_name -> product.Product
	4,  // 2: product.GetCategoriesRequest.categories:type_name -> product.Category
	12, // 3: product.ProductRecommend.created_at:type_name -> google.protobuf.Timestamp
	4,  // 4: product.GetRecommendProductRequest.categories:type_name -> product.Category
	7,  // 5: product.GetRecommendProductRequest.products:type_name -> product.ProductRecommend
	8,  // 6: product.GetRecommendProductRequest.click_details:type_name -> product.ClickDetail
	9,  // 7: product.GetRecommendProductRequest.view_products:type_name -> product.ViewProduct
	1,  // 8: product.ProductService.GetProductsInCart:input_type -> product.GetProductsInCartRequest
	5,  // 9: product.ProductService.GetCategoryList:input_type -> product.GetCategoriesRequest
	10, // 10: product.ProductService.GetRecommendProduct:input_type -> product.GetRecommendProductRequest
	3,  // 11: product.ProductService.GetProductsInCart:output_type -> product.GetProductsInCartResponse
	6,  // 12: product.ProductService.GetCategoryList:output_type -> product.GetCategoriesResponse
	11, // 13: product.ProductService.GetRecommendProduct:output_type -> product.GetRecommendProductResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
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
			NumMessages:   12,
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
