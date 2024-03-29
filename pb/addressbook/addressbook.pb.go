// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/addressbook.proto

package addressbook

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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e945a96804d5af6, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e945a96804d5af6, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=addressbook.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e945a96804d5af6, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e945a96804d5af6, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("addressbook.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "addressbook.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "addressbook.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "addressbook.AddressBook")
}

func init() { proto.RegisterFile("pb/addressbook.proto", fileDescriptor_0e945a96804d5af6) }

var fileDescriptor_0e945a96804d5af6 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0x9a, 0x06, 0x3b, 0x91, 0x52, 0xc6, 0x22, 0xa1, 0x20, 0x0d, 0x3d, 0x05, 0x0a,
	0x29, 0xd6, 0x83, 0x20, 0x78, 0xb0, 0x50, 0x50, 0xb4, 0xb6, 0x2c, 0x15, 0xbd, 0x49, 0x42, 0xc6,
	0x1a, 0x9a, 0x64, 0x97, 0xec, 0xf6, 0xd0, 0x9f, 0xe7, 0x3f, 0x93, 0xec, 0xa6, 0xd2, 0x83, 0xde,
	0xde, 0xcc, 0x7e, 0xbc, 0x7d, 0x6f, 0xa0, 0x2f, 0x92, 0x49, 0x9c, 0xa6, 0x15, 0x49, 0x99, 0x70,
	0xbe, 0x8d, 0x44, 0xc5, 0x15, 0x47, 0xef, 0x68, 0x35, 0x18, 0x6e, 0x38, 0xdf, 0xe4, 0x34, 0xd1,
	0x4f, 0xc9, 0xee, 0x73, 0xa2, 0xb2, 0x82, 0xa4, 0x8a, 0x0b, 0x61, 0xe8, 0xd1, 0xb7, 0x0d, 0xee,
	0x8a, 0x2a, 0xc9, 0x4b, 0x44, 0x70, 0xca, 0xb8, 0x20, 0xdf, 0x0a, 0xac, 0xb0, 0xc3, 0xb4, 0xc6,
	0x2e, 0xd8, 0x59, 0xea, 0xdb, 0x81, 0x15, 0xb6, 0x99, 0x9d, 0xa5, 0xd8, 0x87, 0x36, 0x15, 0x71,
	0x96, 0xfb, 0x2d, 0x0d, 0x99, 0x01, 0x6f, 0xc0, 0x15, 0x5f, 0xbc, 0x24, 0xe9, 0x3b, 0x41, 0x2b,
	0xf4, 0xa6, 0xc3, 0xe8, 0x38, 0x96, 0xb1, 0x8f, 0x56, 0x35, 0xf1, 0xb2, 0x2b, 0x12, 0xaa, 0x58,
	0x83, 0xe3, 0x1d, 0x9c, 0xe5, 0xb1, 0x54, 0x1f, 0x3b, 0x91, 0xc6, 0x8a, 0x52, 0xbf, 0x1d, 0x58,
	0xa1, 0x37, 0x1d, 0x44, 0x26, 0x75, 0x74, 0x48, 0x1d, 0xad, 0x0f, 0xa9, 0x99, 0x57, 0xf3, 0xaf,
	0x06, 0x1f, 0xbc, 0x83, 0x77, 0xe4, 0x8a, 0x17, 0xe0, 0x96, 0x5a, 0x35, 0x15, 0x9a, 0x09, 0xaf,
	0xc0, 0x51, 0x7b, 0x41, 0xba, 0x46, 0x77, 0x7a, 0xf9, 0x6f, 0xb8, 0xf5, 0x5e, 0x10, 0xd3, 0xe8,
	0x68, 0x0c, 0x9d, 0xdf, 0x15, 0x02, 0xb8, 0x8b, 0xe5, 0xec, 0xf1, 0x79, 0xde, 0x3b, 0xc1, 0x53,
	0x70, 0x1e, 0x96, 0x8b, 0x79, 0xcf, 0xaa, 0xd5, 0xdb, 0x92, 0x3d, 0xf5, 0xec, 0xd1, 0x2d, 0x78,
	0xf7, 0xc6, 0x72, 0xc6, 0xf9, 0x16, 0xc7, 0xe0, 0x0a, 0xe2, 0x22, 0xaf, 0x2f, 0x59, 0x5f, 0xe3,
	0xfc, 0x8f, 0x0f, 0x59, 0x83, 0x24, 0xae, 0xee, 0x78, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe8,
	0xe8, 0x85, 0x03, 0xcc, 0x01, 0x00, 0x00,
}
