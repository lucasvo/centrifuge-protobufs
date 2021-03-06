// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification/notification.proto

package notificationpb

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

// NotificationMessage wraps a single CoreDocument to be notified to upstream services
type NotificationMessage struct {
	EventType    uint32               `protobuf:"varint,1,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	Recorded     *timestamp.Timestamp `protobuf:"bytes,3,opt,name=recorded" json:"recorded,omitempty"`
	DocumentType string               `protobuf:"bytes,4,opt,name=document_type,json=documentType" json:"document_type,omitempty"`
	DocumentId   string               `protobuf:"bytes,7,opt,name=document_id,json=documentId" json:"document_id,omitempty"`
	// account_id is the account associated to webhook
	AccountId string `protobuf:"bytes,8,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	// from_id if provided, original trigger of the event
	FromId string `protobuf:"bytes,9,opt,name=from_id,json=fromId" json:"from_id,omitempty"`
	// to_id if provided, final destination of the event
	ToId                 string   `protobuf:"bytes,10,opt,name=to_id,json=toId" json:"to_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationMessage) Reset()         { *m = NotificationMessage{} }
func (m *NotificationMessage) String() string { return proto.CompactTextString(m) }
func (*NotificationMessage) ProtoMessage()    {}
func (*NotificationMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_notification_f0e64ddbcd9b013c, []int{0}
}
func (m *NotificationMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationMessage.Unmarshal(m, b)
}
func (m *NotificationMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationMessage.Marshal(b, m, deterministic)
}
func (dst *NotificationMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationMessage.Merge(dst, src)
}
func (m *NotificationMessage) XXX_Size() int {
	return xxx_messageInfo_NotificationMessage.Size(m)
}
func (m *NotificationMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationMessage.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationMessage proto.InternalMessageInfo

func (m *NotificationMessage) GetEventType() uint32 {
	if m != nil {
		return m.EventType
	}
	return 0
}

func (m *NotificationMessage) GetRecorded() *timestamp.Timestamp {
	if m != nil {
		return m.Recorded
	}
	return nil
}

func (m *NotificationMessage) GetDocumentType() string {
	if m != nil {
		return m.DocumentType
	}
	return ""
}

func (m *NotificationMessage) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
	}
	return ""
}

func (m *NotificationMessage) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *NotificationMessage) GetFromId() string {
	if m != nil {
		return m.FromId
	}
	return ""
}

func (m *NotificationMessage) GetToId() string {
	if m != nil {
		return m.ToId
	}
	return ""
}

func init() {
	proto.RegisterType((*NotificationMessage)(nil), "notification.NotificationMessage")
}

func init() {
	proto.RegisterFile("notification/notification.proto", fileDescriptor_notification_f0e64ddbcd9b013c)
}

var fileDescriptor_notification_f0e64ddbcd9b013c = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x46, 0x59, 0xac, 0x6d, 0x33, 0x6d, 0x45, 0xb7, 0x07, 0x43, 0x41, 0x12, 0xf4, 0x92, 0x53,
	0x02, 0x0a, 0x7a, 0xef, 0x2d, 0x07, 0xa5, 0x84, 0x9e, 0xbc, 0x48, 0xb2, 0x3b, 0x09, 0x01, 0x93,
	0x59, 0x92, 0x8d, 0xd0, 0x5f, 0xaf, 0xec, 0xc6, 0x84, 0x3d, 0xce, 0xfb, 0xde, 0xce, 0xec, 0x0c,
	0x04, 0x2d, 0xe9, 0xba, 0xac, 0x45, 0xae, 0x6b, 0x6a, 0x13, 0xb7, 0x88, 0x55, 0x47, 0x9a, 0xf8,
	0xd6, 0x65, 0x87, 0xa0, 0x22, 0xaa, 0xbe, 0x31, 0xb1, 0x59, 0x31, 0x94, 0x89, 0xae, 0x1b, 0xec,
	0x75, 0xde, 0xa8, 0x51, 0x7f, 0xfc, 0x65, 0xb0, 0xff, 0x70, 0x5e, 0xbc, 0x63, 0xdf, 0xe7, 0x15,
	0xf2, 0x07, 0x00, 0xfc, 0xc1, 0x56, 0x7f, 0xe9, 0x8b, 0x42, 0x9f, 0x85, 0x2c, 0xda, 0x65, 0x9e,
	0x25, 0xe7, 0x8b, 0x42, 0xfe, 0x0a, 0xeb, 0x0e, 0x05, 0x75, 0x12, 0xa5, 0x7f, 0x15, 0xb2, 0x68,
	0xf3, 0x7c, 0x88, 0xc7, 0x51, 0xf1, 0x34, 0x2a, 0x3e, 0x4f, 0xa3, 0xb2, 0xd9, 0xe5, 0x4f, 0xb0,
	0x93, 0x24, 0x86, 0x66, 0xee, 0xbc, 0x08, 0x59, 0xe4, 0x65, 0xdb, 0x09, 0xda, 0xe6, 0x01, 0x6c,
	0x66, 0xa9, 0x96, 0xfe, 0xca, 0x2a, 0x30, 0xa1, 0x54, 0x9a, 0xcf, 0xe5, 0x42, 0xd0, 0x30, 0xe6,
	0x6b, 0x9b, 0x7b, 0xff, 0x24, 0x95, 0xfc, 0x1e, 0x56, 0x65, 0x47, 0x8d, 0xc9, 0x3c, 0x9b, 0x2d,
	0x4d, 0x99, 0x4a, 0xbe, 0x87, 0x6b, 0x4d, 0x06, 0x83, 0xc5, 0x0b, 0x4d, 0xa9, 0x3c, 0xbe, 0xc1,
	0xad, 0xa0, 0x26, 0x76, 0xcf, 0x76, 0xbc, 0x73, 0x4f, 0x72, 0x32, 0x0b, 0x9d, 0xd8, 0xe7, 0x8d,
	0xab, 0xa8, 0xa2, 0x58, 0xda, 0x4d, 0x5f, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x93, 0x54,
	0x87, 0x93, 0x01, 0x00, 0x00,
}
