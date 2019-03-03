// Code generated by protoc-gen-go. DO NOT EDIT.
// source: invoice/invoice.proto

package invoicepb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// InvoiceData is the default invoice schema
type InvoiceData struct {
	// invoice number or reference number
	InvoiceNumber string `protobuf:"bytes,1,opt,name=invoice_number,json=invoiceNumber,proto3" json:"invoice_number,omitempty"`
	// name of the sender company
	SenderName string `protobuf:"bytes,3,opt,name=sender_name,json=senderName,proto3" json:"sender_name,omitempty"`
	// street and address details of the sender company
	SenderStreet  string `protobuf:"bytes,4,opt,name=sender_street,json=senderStreet,proto3" json:"sender_street,omitempty"`
	SenderCity    string `protobuf:"bytes,5,opt,name=sender_city,json=senderCity,proto3" json:"sender_city,omitempty"`
	SenderZipcode string `protobuf:"bytes,6,opt,name=sender_zipcode,json=senderZipcode,proto3" json:"sender_zipcode,omitempty"`
	// country ISO code of the sender of this invoice
	SenderCountry string `protobuf:"bytes,7,opt,name=sender_country,json=senderCountry,proto3" json:"sender_country,omitempty"`
	// name of the recipient company
	RecipientName    string `protobuf:"bytes,8,opt,name=recipient_name,json=recipientName,proto3" json:"recipient_name,omitempty"`
	RecipientStreet  string `protobuf:"bytes,9,opt,name=recipient_street,json=recipientStreet,proto3" json:"recipient_street,omitempty"`
	RecipientCity    string `protobuf:"bytes,10,opt,name=recipient_city,json=recipientCity,proto3" json:"recipient_city,omitempty"`
	RecipientZipcode string `protobuf:"bytes,11,opt,name=recipient_zipcode,json=recipientZipcode,proto3" json:"recipient_zipcode,omitempty"`
	// country ISO code of the receipient of this invoice
	RecipientCountry string `protobuf:"bytes,12,opt,name=recipient_country,json=recipientCountry,proto3" json:"recipient_country,omitempty"`
	// ISO currency code
	Currency string `protobuf:"bytes,13,opt,name=currency,proto3" json:"currency,omitempty"`
	// invoice amount including tax
	GrossAmount int64 `protobuf:"varint,14,opt,name=gross_amount,json=grossAmount,proto3" json:"gross_amount,omitempty"`
	// invoice amount excluding tax
	NetAmount            int64                `protobuf:"varint,15,opt,name=net_amount,json=netAmount,proto3" json:"net_amount,omitempty"`
	TaxAmount            int64                `protobuf:"varint,16,opt,name=tax_amount,json=taxAmount,proto3" json:"tax_amount,omitempty"`
	TaxRate              int64                `protobuf:"varint,17,opt,name=tax_rate,json=taxRate,proto3" json:"tax_rate,omitempty"`
	Recipient            []byte               `protobuf:"bytes,18,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Sender               []byte               `protobuf:"bytes,19,opt,name=sender,proto3" json:"sender,omitempty"`
	Payee                []byte               `protobuf:"bytes,20,opt,name=payee,proto3" json:"payee,omitempty"`
	Comment              string               `protobuf:"bytes,21,opt,name=comment,proto3" json:"comment,omitempty"`
	DueDate              *timestamp.Timestamp `protobuf:"bytes,22,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	DateCreated          *timestamp.Timestamp `protobuf:"bytes,23,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	ExtraData            []byte               `protobuf:"bytes,24,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *InvoiceData) Reset()         { *m = InvoiceData{} }
func (m *InvoiceData) String() string { return proto.CompactTextString(m) }
func (*InvoiceData) ProtoMessage()    {}
func (*InvoiceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3e2b5ce0fcadd51, []int{0}
}

func (m *InvoiceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceData.Unmarshal(m, b)
}
func (m *InvoiceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceData.Marshal(b, m, deterministic)
}
func (m *InvoiceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceData.Merge(m, src)
}
func (m *InvoiceData) XXX_Size() int {
	return xxx_messageInfo_InvoiceData.Size(m)
}
func (m *InvoiceData) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceData.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceData proto.InternalMessageInfo

func (m *InvoiceData) GetInvoiceNumber() string {
	if m != nil {
		return m.InvoiceNumber
	}
	return ""
}

func (m *InvoiceData) GetSenderName() string {
	if m != nil {
		return m.SenderName
	}
	return ""
}

func (m *InvoiceData) GetSenderStreet() string {
	if m != nil {
		return m.SenderStreet
	}
	return ""
}

func (m *InvoiceData) GetSenderCity() string {
	if m != nil {
		return m.SenderCity
	}
	return ""
}

func (m *InvoiceData) GetSenderZipcode() string {
	if m != nil {
		return m.SenderZipcode
	}
	return ""
}

func (m *InvoiceData) GetSenderCountry() string {
	if m != nil {
		return m.SenderCountry
	}
	return ""
}

func (m *InvoiceData) GetRecipientName() string {
	if m != nil {
		return m.RecipientName
	}
	return ""
}

func (m *InvoiceData) GetRecipientStreet() string {
	if m != nil {
		return m.RecipientStreet
	}
	return ""
}

func (m *InvoiceData) GetRecipientCity() string {
	if m != nil {
		return m.RecipientCity
	}
	return ""
}

func (m *InvoiceData) GetRecipientZipcode() string {
	if m != nil {
		return m.RecipientZipcode
	}
	return ""
}

func (m *InvoiceData) GetRecipientCountry() string {
	if m != nil {
		return m.RecipientCountry
	}
	return ""
}

func (m *InvoiceData) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *InvoiceData) GetGrossAmount() int64 {
	if m != nil {
		return m.GrossAmount
	}
	return 0
}

func (m *InvoiceData) GetNetAmount() int64 {
	if m != nil {
		return m.NetAmount
	}
	return 0
}

func (m *InvoiceData) GetTaxAmount() int64 {
	if m != nil {
		return m.TaxAmount
	}
	return 0
}

func (m *InvoiceData) GetTaxRate() int64 {
	if m != nil {
		return m.TaxRate
	}
	return 0
}

func (m *InvoiceData) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *InvoiceData) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *InvoiceData) GetPayee() []byte {
	if m != nil {
		return m.Payee
	}
	return nil
}

func (m *InvoiceData) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *InvoiceData) GetDueDate() *timestamp.Timestamp {
	if m != nil {
		return m.DueDate
	}
	return nil
}

func (m *InvoiceData) GetDateCreated() *timestamp.Timestamp {
	if m != nil {
		return m.DateCreated
	}
	return nil
}

func (m *InvoiceData) GetExtraData() []byte {
	if m != nil {
		return m.ExtraData
	}
	return nil
}

func init() {
	proto.RegisterType((*InvoiceData)(nil), "invoice.InvoiceData")
}

func init() { proto.RegisterFile("invoice/invoice.proto", fileDescriptor_b3e2b5ce0fcadd51) }

var fileDescriptor_b3e2b5ce0fcadd51 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x65, 0x95, 0x26, 0xf1, 0xd8, 0xe9, 0x9f, 0xa5, 0x2d, 0x4b, 0x04, 0x6a, 0x00, 0x21,
	0x19, 0x21, 0x39, 0x12, 0x88, 0x23, 0x07, 0x1a, 0x2e, 0x5c, 0xaa, 0xca, 0x70, 0xea, 0xc5, 0xda,
	0xac, 0x87, 0xc8, 0x12, 0xf6, 0x5a, 0xeb, 0x35, 0x8a, 0xf9, 0x92, 0x7c, 0x25, 0xb4, 0xb3, 0xeb,
	0x38, 0xed, 0x85, 0x53, 0x34, 0xef, 0xfd, 0x76, 0x76, 0x5e, 0x3c, 0x0b, 0x97, 0x65, 0xfd, 0x5b,
	0x95, 0x12, 0x57, 0xfe, 0x37, 0x6d, 0xb4, 0x32, 0x8a, 0x4d, 0x7d, 0xb9, 0xb8, 0xde, 0x2a, 0xb5,
	0xfd, 0x85, 0x2b, 0x92, 0x37, 0xdd, 0xcf, 0x95, 0x29, 0x2b, 0x6c, 0x8d, 0xa8, 0x1a, 0x47, 0xbe,
	0xfe, 0x3b, 0x81, 0xe8, 0x9b, 0x83, 0xbf, 0x0a, 0x23, 0xd8, 0x5b, 0x38, 0xf1, 0x67, 0xf3, 0xba,
	0xab, 0x36, 0xa8, 0x79, 0xb0, 0x0c, 0x92, 0x30, 0x9b, 0x7b, 0xf5, 0x96, 0x44, 0x76, 0x0d, 0x51,
	0x8b, 0x75, 0x81, 0x3a, 0xaf, 0x45, 0x85, 0xfc, 0x88, 0x18, 0x70, 0xd2, 0xad, 0xa8, 0x90, 0xbd,
	0x81, 0xb9, 0x07, 0x5a, 0xa3, 0x11, 0x0d, 0x7f, 0x42, 0x48, 0xec, 0xc4, 0xef, 0xa4, 0x1d, 0x74,
	0x91, 0xa5, 0xe9, 0xf9, 0xf1, 0x61, 0x97, 0x75, 0x69, 0x7a, 0x3b, 0x8d, 0x07, 0xfe, 0x94, 0x8d,
	0x54, 0x05, 0xf2, 0x89, 0x9b, 0xc6, 0xa9, 0xf7, 0x4e, 0x3c, 0xc0, 0xa4, 0xea, 0x6a, 0xa3, 0x7b,
	0x3e, 0x3d, 0xc4, 0xd6, 0x4e, 0xb4, 0x98, 0x46, 0x59, 0x36, 0x25, 0xd6, 0xc6, 0xcd, 0x3d, 0x73,
	0xd8, 0x5e, 0xa5, 0xd1, 0xdf, 0xc1, 0xd9, 0x88, 0xf9, 0xe9, 0x43, 0x02, 0x4f, 0xf7, 0xba, 0x0f,
	0xf0, 0xa0, 0x23, 0x65, 0x80, 0x47, 0x1d, 0x29, 0xc6, 0x7b, 0x38, 0x1f, 0xb1, 0x21, 0x49, 0x44,
	0xe4, 0x78, 0xd5, 0x10, 0xe6, 0x01, 0x3c, 0xe4, 0x89, 0x1f, 0xc1, 0x43, 0xa4, 0x05, 0xcc, 0x64,
	0xa7, 0x35, 0xd6, 0xb2, 0xe7, 0x73, 0x62, 0xf6, 0x35, 0x7b, 0x05, 0xf1, 0x56, 0xab, 0xb6, 0xcd,
	0x45, 0x65, 0x69, 0x7e, 0xb2, 0x0c, 0x92, 0xa3, 0x2c, 0x22, 0xed, 0x0b, 0x49, 0xec, 0x25, 0x40,
	0x8d, 0x66, 0x00, 0x4e, 0x09, 0x08, 0x6b, 0x34, 0xa3, 0x6d, 0xc4, 0x6e, 0xb0, 0xcf, 0x9c, 0x6d,
	0xc4, 0xce, 0xdb, 0xcf, 0x61, 0x66, 0x6d, 0x2d, 0x0c, 0xf2, 0x73, 0x32, 0xa7, 0x46, 0xec, 0x32,
	0x61, 0x90, 0xbd, 0x80, 0x70, 0x3f, 0x2b, 0x67, 0xcb, 0x20, 0x89, 0xb3, 0x51, 0x60, 0x57, 0x30,
	0x71, 0x5f, 0x86, 0x3f, 0x25, 0xcb, 0x57, 0xec, 0x02, 0x8e, 0x1b, 0xd1, 0x23, 0xf2, 0x0b, 0x92,
	0x5d, 0xc1, 0x38, 0x4c, 0xa5, 0xaa, 0x2a, 0xdb, 0xe9, 0x92, 0x22, 0x0e, 0x25, 0xfb, 0x04, 0xb3,
	0xa2, 0xc3, 0xbc, 0xb0, 0x03, 0x5c, 0x2d, 0x83, 0x24, 0xfa, 0xb0, 0x48, 0xdd, 0xc2, 0xa7, 0xc3,
	0xc2, 0xa7, 0x3f, 0x86, 0x85, 0xcf, 0xa6, 0x45, 0x67, 0x77, 0x1c, 0xd9, 0x67, 0x88, 0xed, 0x91,
	0x5c, 0x6a, 0x14, 0x06, 0x0b, 0xfe, 0xec, 0xbf, 0x47, 0x23, 0xcb, 0xaf, 0x1d, 0x6e, 0xff, 0x15,
	0xdc, 0x19, 0x2d, 0xec, 0xbd, 0x82, 0x73, 0x17, 0x8e, 0x14, 0xfb, 0x82, 0x6e, 0x12, 0x88, 0xa4,
	0xaa, 0x52, 0xff, 0x5e, 0x6e, 0x62, 0xff, 0xba, 0xee, 0x6c, 0xd7, 0xbb, 0xe0, 0x3e, 0xf4, 0x46,
	0xb3, 0xd9, 0x4c, 0xe8, 0xa6, 0x8f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x88, 0x14, 0x37, 0x55,
	0xc5, 0x03, 0x00, 0x00,
}
