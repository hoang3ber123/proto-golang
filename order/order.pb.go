// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0
// source: order/order.proto

package orderproto

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

// CreateOrderRequest message
type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TransactionId string                 `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	PaymentMethod string                 `protobuf:"bytes,2,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_order_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *CreateOrderRequest) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

// CreateOrderResponse message
type CreateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsCreated     bool                   `protobuf:"varint,1,opt,name=is_created,json=isCreated,proto3" json:"is_created,omitempty"`
	StatusCode    int32                  `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_order_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderResponse) GetIsCreated() bool {
	if x != nil {
		return x.IsCreated
	}
	return false
}

func (x *CreateOrderResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CreateOrderResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// GetProductIDsRequest message
type GetProductIDsRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	OrderId         string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`                            // Lọc theo order_id nếu có
	UserId          string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                               // Lọc theo user_id nếu có
	RelatedType     string                 `protobuf:"bytes,3,opt,name=related_type,json=relatedType,proto3" json:"related_type,omitempty"`                // Lọc theo related_type nếu có
	PaymentMethod   string                 `protobuf:"bytes,4,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`          // Chỉ cho phép "vnpay" hoặc "stripe"
	PaymentStatus   string                 `protobuf:"bytes,5,opt,name=payment_status,json=paymentStatus,proto3" json:"payment_status,omitempty"`          // Chỉ cho phép "success" hoặc "cancel"
	Page            int32                  `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`                                                // Trang hiện tại
	PageSize        int32                  `protobuf:"varint,7,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`                        // Số bản ghi mỗi trang
	MaxPrice        string                 `protobuf:"bytes,8,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`                         // Giá tối đa (chuỗi để dễ parse)
	MinPrice        string                 `protobuf:"bytes,9,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`                         // Giá tối thiểu (chuỗi để dễ parse)
	EndPaymentDay   string                 `protobuf:"bytes,10,opt,name=end_payment_day,json=endPaymentDay,proto3" json:"end_payment_day,omitempty"`       // Ngày kết thúc thanh toán (YYYY-MM-DD)
	StartPaymentDay string                 `protobuf:"bytes,11,opt,name=start_payment_day,json=startPaymentDay,proto3" json:"start_payment_day,omitempty"` // Ngày bắt đầu thanh toán (YYYY-MM-DD)
	PaymentDayOrder string                 `protobuf:"bytes,12,opt,name=payment_day_order,json=paymentDayOrder,proto3" json:"payment_day_order,omitempty"` // Sắp xếp theo payment_day: "asc" hoặc "desc"
	PriceOrder      string                 `protobuf:"bytes,13,opt,name=price_order,json=priceOrder,proto3" json:"price_order,omitempty"`                  // Sắp xếp theo giá: "asc" hoặc "desc"
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GetProductIDsRequest) Reset() {
	*x = GetProductIDsRequest{}
	mi := &file_order_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductIDsRequest) ProtoMessage() {}

func (x *GetProductIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductIDsRequest.ProtoReflect.Descriptor instead.
func (*GetProductIDsRequest) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductIDsRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *GetProductIDsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetProductIDsRequest) GetRelatedType() string {
	if x != nil {
		return x.RelatedType
	}
	return ""
}

func (x *GetProductIDsRequest) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *GetProductIDsRequest) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

func (x *GetProductIDsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetProductIDsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetProductIDsRequest) GetMaxPrice() string {
	if x != nil {
		return x.MaxPrice
	}
	return ""
}

func (x *GetProductIDsRequest) GetMinPrice() string {
	if x != nil {
		return x.MinPrice
	}
	return ""
}

func (x *GetProductIDsRequest) GetEndPaymentDay() string {
	if x != nil {
		return x.EndPaymentDay
	}
	return ""
}

func (x *GetProductIDsRequest) GetStartPaymentDay() string {
	if x != nil {
		return x.StartPaymentDay
	}
	return ""
}

func (x *GetProductIDsRequest) GetPaymentDayOrder() string {
	if x != nil {
		return x.PaymentDayOrder
	}
	return ""
}

func (x *GetProductIDsRequest) GetPriceOrder() string {
	if x != nil {
		return x.PriceOrder
	}
	return ""
}

// Pagination message
type Pagination struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`                            // Trang hiện tại
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`    // Số bản ghi mỗi trang
	Total         int32                  `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`                          // Tổng số bản ghi
	TotalPage     int32                  `protobuf:"varint,4,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"` // Tổng số trang
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	mi := &file_order_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{3}
}

func (x *Pagination) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Pagination) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Pagination) GetTotalPage() int32 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

// GetProductIDsResponse message
type GetProductIDsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductIds    []string               `protobuf:"bytes,1,rep,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`  // Danh sách ID sản phẩm
	Pagination    *Pagination            `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`                    // Thông tin phân trang
	StatusCode    int32                  `protobuf:"varint,3,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // Mã trạng thái HTTP
	Error         string                 `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`                              // Thông báo lỗi (nếu có)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductIDsResponse) Reset() {
	*x = GetProductIDsResponse{}
	mi := &file_order_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductIDsResponse) ProtoMessage() {}

func (x *GetProductIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductIDsResponse.ProtoReflect.Descriptor instead.
func (*GetProductIDsResponse) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductIDsResponse) GetProductIds() []string {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

func (x *GetProductIDsResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *GetProductIDsResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetProductIDsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// CheckBoughtProductRequest message
type CheckBoughtProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RelatedId     string                 `protobuf:"bytes,2,opt,name=related_id,json=relatedId,proto3" json:"related_id,omitempty"`
	RelatedType   string                 `protobuf:"bytes,3,opt,name=related_type,json=relatedType,proto3" json:"related_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckBoughtProductRequest) Reset() {
	*x = CheckBoughtProductRequest{}
	mi := &file_order_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckBoughtProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckBoughtProductRequest) ProtoMessage() {}

func (x *CheckBoughtProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckBoughtProductRequest.ProtoReflect.Descriptor instead.
func (*CheckBoughtProductRequest) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{5}
}

func (x *CheckBoughtProductRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CheckBoughtProductRequest) GetRelatedId() string {
	if x != nil {
		return x.RelatedId
	}
	return ""
}

func (x *CheckBoughtProductRequest) GetRelatedType() string {
	if x != nil {
		return x.RelatedType
	}
	return ""
}

// CheckBoughtProductResponse message
type CheckBoughtProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsBought      bool                   `protobuf:"varint,1,opt,name=is_bought,json=isBought,proto3" json:"is_bought,omitempty"`
	StatusCode    int32                  `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckBoughtProductResponse) Reset() {
	*x = CheckBoughtProductResponse{}
	mi := &file_order_order_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckBoughtProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckBoughtProductResponse) ProtoMessage() {}

func (x *CheckBoughtProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_order_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckBoughtProductResponse.ProtoReflect.Descriptor instead.
func (*CheckBoughtProductResponse) Descriptor() ([]byte, []int) {
	return file_order_order_proto_rawDescGZIP(), []int{6}
}

func (x *CheckBoughtProductResponse) GetIsBought() bool {
	if x != nil {
		return x.IsBought
	}
	return false
}

func (x *CheckBoughtProductResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CheckBoughtProductResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_order_order_proto protoreflect.FileDescriptor

var file_order_order_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x62, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x6b,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xc7, 0x03, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61,
	0x78, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65,
	0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x79, 0x12, 0x2a, 0x0a, 0x11,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61,
	0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x79, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x72, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x76,
	0x0a, 0x19, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x70, 0x0a, 0x1a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42,
	0x6f, 0x75, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x62, 0x6f, 0x75, 0x67, 0x68,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x42, 0x6f, 0x75, 0x67, 0x68,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x81, 0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x44, 0x73, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5b, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x42, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x61, 0x6e, 0x67,
	0x33, 0x62, 0x65, 0x72, 0x31, 0x32, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_order_order_proto_rawDescOnce sync.Once
	file_order_order_proto_rawDescData []byte
)

func file_order_order_proto_rawDescGZIP() []byte {
	file_order_order_proto_rawDescOnce.Do(func() {
		file_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_order_proto_rawDesc), len(file_order_order_proto_rawDesc)))
	})
	return file_order_order_proto_rawDescData
}

var file_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_order_order_proto_goTypes = []any{
	(*CreateOrderRequest)(nil),         // 0: order.CreateOrderRequest
	(*CreateOrderResponse)(nil),        // 1: order.CreateOrderResponse
	(*GetProductIDsRequest)(nil),       // 2: order.GetProductIDsRequest
	(*Pagination)(nil),                 // 3: order.Pagination
	(*GetProductIDsResponse)(nil),      // 4: order.GetProductIDsResponse
	(*CheckBoughtProductRequest)(nil),  // 5: order.CheckBoughtProductRequest
	(*CheckBoughtProductResponse)(nil), // 6: order.CheckBoughtProductResponse
}
var file_order_order_proto_depIdxs = []int32{
	3, // 0: order.GetProductIDsResponse.pagination:type_name -> order.Pagination
	0, // 1: order.OrderService.CreateOrder:input_type -> order.CreateOrderRequest
	2, // 2: order.OrderService.GetProductIDs:input_type -> order.GetProductIDsRequest
	5, // 3: order.OrderService.CheckBoughtProduct:input_type -> order.CheckBoughtProductRequest
	1, // 4: order.OrderService.CreateOrder:output_type -> order.CreateOrderResponse
	4, // 5: order.OrderService.GetProductIDs:output_type -> order.GetProductIDsResponse
	6, // 6: order.OrderService.CheckBoughtProduct:output_type -> order.CheckBoughtProductResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_order_order_proto_init() }
func file_order_order_proto_init() {
	if File_order_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_order_proto_rawDesc), len(file_order_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_order_proto_goTypes,
		DependencyIndexes: file_order_order_proto_depIdxs,
		MessageInfos:      file_order_order_proto_msgTypes,
	}.Build()
	File_order_order_proto = out.File
	file_order_order_proto_goTypes = nil
	file_order_order_proto_depIdxs = nil
}
