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
	// invoice status
	InvoiceStatus string `protobuf:"bytes,2,opt,name=invoice_status,json=invoiceStatus,proto3" json:"invoice_status,omitempty"`
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
	GrossAmount []byte `protobuf:"bytes,14,opt,name=gross_amount,json=grossAmount,proto3" json:"gross_amount,omitempty"`
	// invoice amount excluding tax
	NetAmount            []byte               `protobuf:"bytes,15,opt,name=net_amount,json=netAmount,proto3" json:"net_amount,omitempty"`
	TaxAmount            []byte               `protobuf:"bytes,16,opt,name=tax_amount,json=taxAmount,proto3" json:"tax_amount,omitempty"`
	TaxRate              []byte               `protobuf:"bytes,17,opt,name=tax_rate,json=taxRate,proto3" json:"tax_rate,omitempty"`
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

func (m *InvoiceData) GetInvoiceStatus() string {
	if m != nil {
		return m.InvoiceStatus
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

func (m *InvoiceData) GetGrossAmount() []byte {
	if m != nil {
		return m.GrossAmount
	}
	return nil
}

func (m *InvoiceData) GetNetAmount() []byte {
	if m != nil {
		return m.NetAmount
	}
	return nil
}

func (m *InvoiceData) GetTaxAmount() []byte {
	if m != nil {
		return m.TaxAmount
	}
	return nil
}

func (m *InvoiceData) GetTaxRate() []byte {
	if m != nil {
		return m.TaxRate
	}
	return nil
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
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcb, 0x6f, 0xd3, 0x40,
	0x10, 0x87, 0x15, 0xa0, 0x79, 0x8c, 0x9d, 0x3e, 0x96, 0xb6, 0x2c, 0x11, 0xa8, 0x01, 0x84, 0x14,
	0x84, 0x94, 0x48, 0x20, 0x8e, 0x1c, 0x68, 0xb8, 0x70, 0xa9, 0xaa, 0x94, 0x53, 0x2f, 0xd1, 0x66,
	0x3d, 0x44, 0x96, 0xb0, 0xd7, 0x5a, 0x8f, 0x51, 0xcc, 0x1f, 0xce, 0x19, 0xed, 0xec, 0xda, 0x71,
	0x7a, 0xe9, 0x29, 0xda, 0x6f, 0xbe, 0x9d, 0x9d, 0x5f, 0xac, 0x81, 0x8b, 0x34, 0xff, 0x63, 0x52,
	0x8d, 0x8b, 0xf0, 0x3b, 0x2f, 0xac, 0x21, 0x23, 0x06, 0xe1, 0x38, 0xb9, 0xda, 0x1a, 0xb3, 0xfd,
	0x8d, 0x0b, 0xc6, 0x9b, 0xea, 0xd7, 0x82, 0xd2, 0x0c, 0x4b, 0x52, 0x59, 0xe1, 0xcd, 0xb7, 0xff,
	0xfa, 0x10, 0xfd, 0xf0, 0xf2, 0x77, 0x45, 0x4a, 0xbc, 0x87, 0xe3, 0x70, 0x77, 0x9d, 0x57, 0xd9,
	0x06, 0xad, 0xec, 0x4d, 0x7b, 0xb3, 0xd1, 0x6a, 0x1c, 0xe8, 0x0d, 0xc3, 0xae, 0x56, 0x92, 0xa2,
	0xaa, 0x94, 0x4f, 0x0e, 0xb4, 0x3b, 0x86, 0xe2, 0x0a, 0xa2, 0x12, 0xf3, 0x04, 0xed, 0x3a, 0x57,
	0x19, 0xca, 0xa7, 0xec, 0x80, 0x47, 0x37, 0x2a, 0x43, 0xf1, 0x0e, 0xc6, 0x41, 0x28, 0xc9, 0x22,
	0x92, 0x7c, 0xc6, 0x4a, 0xec, 0xe1, 0x1d, 0xb3, 0x4e, 0x17, 0x9d, 0x52, 0x2d, 0x8f, 0xba, 0x5d,
	0x96, 0x29, 0xd5, 0x6e, 0x9a, 0x20, 0xfc, 0x4d, 0x0b, 0x6d, 0x12, 0x94, 0x7d, 0x3f, 0x8d, 0xa7,
	0xf7, 0x1e, 0x76, 0x34, 0x6d, 0xaa, 0x9c, 0x6c, 0x2d, 0x07, 0x5d, 0x6d, 0xe9, 0xa1, 0xd3, 0x2c,
	0xea, 0xb4, 0x48, 0x31, 0x27, 0x3f, 0xf7, 0xd0, 0x6b, 0x2d, 0xe5, 0xd1, 0x3f, 0xc0, 0xe9, 0x5e,
	0x0b, 0xd3, 0x8f, 0x58, 0x3c, 0x69, 0x79, 0x08, 0x70, 0xd0, 0x91, 0x33, 0xc0, 0x83, 0x8e, 0x1c,
	0xe3, 0x23, 0x9c, 0xed, 0xb5, 0x26, 0x49, 0xc4, 0xe6, 0xfe, 0xa9, 0x26, 0xcc, 0x81, 0xdc, 0xe4,
	0x89, 0x1f, 0xc8, 0x4d, 0xa4, 0x09, 0x0c, 0x75, 0x65, 0x2d, 0xe6, 0xba, 0x96, 0x63, 0x76, 0xda,
	0xb3, 0x78, 0x03, 0xf1, 0xd6, 0x9a, 0xb2, 0x5c, 0xab, 0xcc, 0xd9, 0xf2, 0x78, 0xda, 0x9b, 0xc5,
	0xab, 0x88, 0xd9, 0x37, 0x46, 0xe2, 0x35, 0x40, 0x8e, 0xd4, 0x08, 0x27, 0x2c, 0x8c, 0x72, 0xa4,
	0x7d, 0x99, 0xd4, 0xae, 0x29, 0x9f, 0xfa, 0x32, 0xa9, 0x5d, 0x28, 0xbf, 0x84, 0xa1, 0x2b, 0x5b,
	0x45, 0x28, 0xcf, 0xb8, 0x38, 0x20, 0xb5, 0x5b, 0x29, 0x42, 0xf1, 0x0a, 0x46, 0xed, 0xac, 0x52,
	0xf8, 0x8b, 0x2d, 0x10, 0x97, 0xd0, 0xf7, 0x5f, 0x46, 0x3e, 0xe7, 0x52, 0x38, 0x89, 0x73, 0x38,
	0x2a, 0x54, 0x8d, 0x28, 0xcf, 0x19, 0xfb, 0x83, 0x90, 0x30, 0xd0, 0x26, 0xcb, 0x5c, 0xa7, 0x0b,
	0x8e, 0xd8, 0x1c, 0xc5, 0x17, 0x18, 0x26, 0x15, 0xae, 0x13, 0x37, 0xc0, 0xe5, 0xb4, 0x37, 0x8b,
	0x3e, 0x4d, 0xe6, 0x7e, 0x2f, 0xe6, 0xcd, 0x5e, 0xcc, 0x7f, 0x36, 0x7b, 0xb1, 0x1a, 0x24, 0x95,
	0x5b, 0x05, 0x14, 0x5f, 0x21, 0x76, 0x57, 0xd6, 0xda, 0xa2, 0x22, 0x4c, 0xe4, 0x8b, 0x47, 0xaf,
	0x46, 0xce, 0x5f, 0x7a, 0xdd, 0xfd, 0x2b, 0xb8, 0x23, 0xab, 0xdc, 0xbb, 0x4a, 0x4a, 0x1f, 0x8e,
	0x89, 0x5b, 0xb4, 0xeb, 0x19, 0x44, 0xda, 0x64, 0xf3, 0xb0, 0x2f, 0xd7, 0x71, 0x58, 0xc2, 0x5b,
	0xd7, 0xf5, 0xb6, 0x77, 0x3f, 0x0a, 0x85, 0x62, 0xb3, 0xe9, 0xf3, 0x4b, 0x9f, 0xff, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xd6, 0xe0, 0xdc, 0x6d, 0xec, 0x03, 0x00, 0x00,
}
