// Code generated by protoc-gen-go. DO NOT EDIT.
// source: purchaseorder/purchaseorder.proto

package purchaseorderpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// PurchaseOrderData is the default schema for a purchase order
type PurchaseOrderData struct {
	// purchase order number or reference number
	PoNumber string `protobuf:"bytes,1,opt,name=po_number,json=poNumber,proto3" json:"po_number,omitempty"`
	// name of the ordering company
	OrderName string `protobuf:"bytes,2,opt,name=order_name,json=orderName,proto3" json:"order_name,omitempty"`
	// street and address details of the ordering company
	OrderStreet  string `protobuf:"bytes,3,opt,name=order_street,json=orderStreet,proto3" json:"order_street,omitempty"`
	OrderCity    string `protobuf:"bytes,4,opt,name=order_city,json=orderCity,proto3" json:"order_city,omitempty"`
	OrderZipcode string `protobuf:"bytes,5,opt,name=order_zipcode,json=orderZipcode,proto3" json:"order_zipcode,omitempty"`
	// country ISO code of the ordering company of this purchase order
	OrderCountry string `protobuf:"bytes,6,opt,name=order_country,json=orderCountry,proto3" json:"order_country,omitempty"`
	// name of the recipient company
	RecipientName    string `protobuf:"bytes,7,opt,name=recipient_name,json=recipientName,proto3" json:"recipient_name,omitempty"`
	RecipientStreet  string `protobuf:"bytes,8,opt,name=recipient_street,json=recipientStreet,proto3" json:"recipient_street,omitempty"`
	RecipientCity    string `protobuf:"bytes,9,opt,name=recipient_city,json=recipientCity,proto3" json:"recipient_city,omitempty"`
	RecipientZipcode string `protobuf:"bytes,10,opt,name=recipient_zipcode,json=recipientZipcode,proto3" json:"recipient_zipcode,omitempty"`
	// country ISO code of the receipient of this purchase order
	RecipientCountry string `protobuf:"bytes,11,opt,name=recipient_country,json=recipientCountry,proto3" json:"recipient_country,omitempty"`
	// ISO currency code
	Currency string `protobuf:"bytes,12,opt,name=currency,proto3" json:"currency,omitempty"`
	// ordering gross amount including tax
	OrderAmount []byte `protobuf:"bytes,13,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount,omitempty"`
	// invoice amount excluding tax
	NetAmount []byte `protobuf:"bytes,14,opt,name=net_amount,json=netAmount,proto3" json:"net_amount,omitempty"`
	TaxAmount []byte `protobuf:"bytes,15,opt,name=tax_amount,json=taxAmount,proto3" json:"tax_amount,omitempty"`
	TaxRate   []byte `protobuf:"bytes,16,opt,name=tax_rate,json=taxRate,proto3" json:"tax_rate,omitempty"`
	Recipient []byte `protobuf:"bytes,17,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Order     []byte `protobuf:"bytes,18,opt,name=order,proto3" json:"order,omitempty"`
	// contact or requester or purchaser at the ordering company
	OrderContact string `protobuf:"bytes,19,opt,name=order_contact,json=orderContact,proto3" json:"order_contact,omitempty"`
	Comment      string `protobuf:"bytes,20,opt,name=comment,proto3" json:"comment,omitempty"`
	// requested delivery date
	DeliveryDate *timestamp.Timestamp `protobuf:"bytes,21,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	// purchase order date
	DateCreated *timestamp.Timestamp `protobuf:"bytes,22,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	ExtraData   []byte               `protobuf:"bytes,23,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
	// purchase order status
	PoStatus             string   `protobuf:"bytes,24,opt,name=po_status,json=poStatus,proto3" json:"po_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PurchaseOrderData) Reset()         { *m = PurchaseOrderData{} }
func (m *PurchaseOrderData) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderData) ProtoMessage()    {}
func (*PurchaseOrderData) Descriptor() ([]byte, []int) {
	return fileDescriptor_purchaseorder_0a1b86b207679c65, []int{0}
}
func (m *PurchaseOrderData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderData.Unmarshal(m, b)
}
func (m *PurchaseOrderData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderData.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderData.Merge(dst, src)
}
func (m *PurchaseOrderData) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderData.Size(m)
}
func (m *PurchaseOrderData) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderData.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderData proto.InternalMessageInfo

func (m *PurchaseOrderData) GetPoNumber() string {
	if m != nil {
		return m.PoNumber
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderName() string {
	if m != nil {
		return m.OrderName
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderStreet() string {
	if m != nil {
		return m.OrderStreet
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderCity() string {
	if m != nil {
		return m.OrderCity
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderZipcode() string {
	if m != nil {
		return m.OrderZipcode
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderCountry() string {
	if m != nil {
		return m.OrderCountry
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientName() string {
	if m != nil {
		return m.RecipientName
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientStreet() string {
	if m != nil {
		return m.RecipientStreet
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientCity() string {
	if m != nil {
		return m.RecipientCity
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientZipcode() string {
	if m != nil {
		return m.RecipientZipcode
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientCountry() string {
	if m != nil {
		return m.RecipientCountry
	}
	return ""
}

func (m *PurchaseOrderData) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderAmount() []byte {
	if m != nil {
		return m.OrderAmount
	}
	return nil
}

func (m *PurchaseOrderData) GetNetAmount() []byte {
	if m != nil {
		return m.NetAmount
	}
	return nil
}

func (m *PurchaseOrderData) GetTaxAmount() []byte {
	if m != nil {
		return m.TaxAmount
	}
	return nil
}

func (m *PurchaseOrderData) GetTaxRate() []byte {
	if m != nil {
		return m.TaxRate
	}
	return nil
}

func (m *PurchaseOrderData) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *PurchaseOrderData) GetOrder() []byte {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *PurchaseOrderData) GetOrderContact() string {
	if m != nil {
		return m.OrderContact
	}
	return ""
}

func (m *PurchaseOrderData) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *PurchaseOrderData) GetDeliveryDate() *timestamp.Timestamp {
	if m != nil {
		return m.DeliveryDate
	}
	return nil
}

func (m *PurchaseOrderData) GetDateCreated() *timestamp.Timestamp {
	if m != nil {
		return m.DateCreated
	}
	return nil
}

func (m *PurchaseOrderData) GetExtraData() []byte {
	if m != nil {
		return m.ExtraData
	}
	return nil
}

func (m *PurchaseOrderData) GetPoStatus() string {
	if m != nil {
		return m.PoStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*PurchaseOrderData)(nil), "purchaseorder.PurchaseOrderData")
}

func init() {
	proto.RegisterFile("purchaseorder/purchaseorder.proto", fileDescriptor_purchaseorder_0a1b86b207679c65)
}

var fileDescriptor_purchaseorder_0a1b86b207679c65 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x55, 0x60, 0x6b, 0xe3, 0x26, 0xeb, 0x6a, 0x06, 0x98, 0x02, 0xa2, 0x03, 0x21, 0x15,
	0x21, 0xa5, 0x12, 0xdc, 0x90, 0x10, 0x62, 0xdd, 0x79, 0x54, 0x1d, 0xa7, 0x5d, 0x2a, 0xd7, 0x79,
	0x8c, 0x48, 0x4b, 0x1c, 0x39, 0x2f, 0xa8, 0xe1, 0xff, 0xe6, 0x8e, 0xfc, 0x1c, 0x37, 0xc9, 0x2e,
	0x1c, 0xdf, 0xf7, 0xfd, 0xde, 0xcb, 0xf7, 0x1c, 0x9b, 0x9d, 0x17, 0x95, 0x51, 0xbf, 0x64, 0x09,
	0xda, 0x24, 0x60, 0x96, 0xbd, 0x2a, 0x2e, 0x8c, 0x46, 0xcd, 0xa3, 0x9e, 0x38, 0x7b, 0x7d, 0xab,
	0xf5, 0xed, 0x1d, 0x2c, 0xc9, 0xdc, 0x55, 0x3f, 0x97, 0x98, 0x66, 0x50, 0xa2, 0xcc, 0x0a, 0xc7,
	0xbf, 0xf9, 0x7b, 0xcc, 0xa6, 0xeb, 0xa6, 0xe5, 0xbb, 0x6d, 0xb9, 0x94, 0x28, 0xf9, 0x0b, 0x16,
	0x14, 0x7a, 0x9b, 0x57, 0xd9, 0x0e, 0x8c, 0x18, 0xcc, 0x07, 0x8b, 0x60, 0x33, 0x2a, 0xf4, 0x15,
	0xd5, 0xfc, 0x15, 0x63, 0x34, 0x7c, 0x9b, 0xcb, 0x0c, 0xc4, 0x03, 0x72, 0x03, 0x52, 0xae, 0x64,
	0x06, 0xfc, 0x9c, 0x85, 0xce, 0x2e, 0xd1, 0x00, 0xa0, 0x78, 0x48, 0xc0, 0x98, 0xb4, 0x6b, 0x92,
	0xda, 0x09, 0x2a, 0xc5, 0x5a, 0x3c, 0xea, 0x4c, 0x58, 0xa5, 0x58, 0xf3, 0xb7, 0x2c, 0x72, 0xf6,
	0x9f, 0xb4, 0x50, 0x3a, 0x01, 0x71, 0x44, 0x84, 0x1b, 0x7b, 0xe3, 0xb4, 0x16, 0x52, 0xba, 0xca,
	0xd1, 0xd4, 0xe2, 0xb8, 0x03, 0xad, 0x9c, 0xc6, 0xdf, 0xb1, 0x13, 0x03, 0x2a, 0x2d, 0x52, 0xc8,
	0xd1, 0xc5, 0x1d, 0x12, 0x15, 0x1d, 0x54, 0x8a, 0xfc, 0x9e, 0x9d, 0xb6, 0x58, 0x13, 0x7b, 0x44,
	0xe0, 0xe4, 0xa0, 0x37, 0xd1, 0x7b, 0x13, 0x29, 0x7e, 0x70, 0x6f, 0x22, 0xad, 0xf0, 0x81, 0x4d,
	0x5b, 0xcc, 0xaf, 0xc1, 0x88, 0x6c, 0x3f, 0xe5, 0x57, 0xe9, 0xc1, 0x7e, 0x9d, 0xf1, 0x3d, 0xd8,
	0xaf, 0x34, 0x63, 0x23, 0x55, 0x19, 0x03, 0xb9, 0xaa, 0x45, 0xe8, 0xfe, 0x8c, 0xaf, 0xdb, 0xa3,
	0x97, 0x99, 0xa5, 0x45, 0x34, 0x1f, 0x2c, 0xc2, 0xe6, 0xe8, 0xbf, 0x91, 0x64, 0x8f, 0x3e, 0x07,
	0xf4, 0xc0, 0x09, 0x01, 0x41, 0x0e, 0xd8, 0xda, 0x28, 0xf7, 0xde, 0x9e, 0x38, 0x1b, 0xe5, 0xbe,
	0xb1, 0x9f, 0xb3, 0x91, 0xb5, 0x8d, 0x44, 0x10, 0xa7, 0x64, 0x0e, 0x51, 0xee, 0x37, 0x12, 0x81,
	0xbf, 0x64, 0xc1, 0x21, 0xab, 0x98, 0xba, 0xc6, 0x83, 0xc0, 0xcf, 0xd8, 0x11, 0xa5, 0x10, 0x9c,
	0x1c, 0x57, 0x74, 0xff, 0x61, 0x8e, 0x52, 0xa1, 0x78, 0xdc, 0xfb, 0x87, 0xa4, 0x71, 0xc1, 0x86,
	0x4a, 0x67, 0x99, 0x1d, 0x7b, 0x46, 0xb6, 0x2f, 0xf9, 0x57, 0x16, 0x25, 0x70, 0x97, 0xfe, 0x06,
	0x53, 0x6f, 0x13, 0x1b, 0xe9, 0xc9, 0x7c, 0xb0, 0x18, 0x7f, 0x9c, 0xc5, 0xee, 0xd2, 0xc7, 0xfe,
	0xd2, 0xc7, 0x3f, 0xfc, 0xa5, 0xdf, 0x84, 0xbe, 0xe1, 0xd2, 0x66, 0xfe, 0xc2, 0x42, 0xdb, 0xb7,
	0x55, 0x06, 0x24, 0x42, 0x22, 0x9e, 0xfe, 0xb7, 0x7f, 0x6c, 0xf9, 0x95, 0xc3, 0xed, 0x61, 0xc1,
	0x1e, 0x8d, 0xb4, 0x1f, 0x97, 0xe2, 0x99, 0xdb, 0x99, 0x94, 0xce, 0x23, 0x2a, 0x51, 0x62, 0x55,
	0x0a, 0xe1, 0x1f, 0xd1, 0x35, 0xd5, 0x17, 0x9f, 0xd9, 0x54, 0xe9, 0x2c, 0xee, 0xbd, 0xd6, 0x0b,
	0xbe, 0xee, 0x96, 0x6b, 0xfb, 0xf9, 0xf5, 0xe0, 0x66, 0xd2, 0x83, 0x8a, 0xdd, 0xee, 0x98, 0x82,
	0x7d, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0x02, 0x37, 0x3f, 0x4f, 0x0f, 0x04, 0x00, 0x00,
}
