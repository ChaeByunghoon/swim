// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Seq string `protobuf:"bytes,1,opt,name=seq,proto3" json:"seq,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*Message_Ping
	//	*Message_Ack
	//	*Message_IndirectPing
	Payload              isMessage_Payload `protobuf_oneof:"payload"`
	PiggyBack            *PiggyBack        `protobuf:"bytes,5,opt,name=piggyBack,proto3" json:"piggyBack,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSeq() string {
	if m != nil {
		return m.Seq
	}
	return ""
}

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_Ping struct {
	Ping *Ping `protobuf:"bytes,2,opt,name=ping,proto3,oneof"`
}

type Message_Ack struct {
	Ack *Ack `protobuf:"bytes,3,opt,name=ack,proto3,oneof"`
}

type Message_IndirectPing struct {
	IndirectPing *IndirectPing `protobuf:"bytes,4,opt,name=indirect_ping,json=indirectPing,proto3,oneof"`
}

func (*Message_Ping) isMessage_Payload() {}

func (*Message_Ack) isMessage_Payload() {}

func (*Message_IndirectPing) isMessage_Payload() {}

func (m *Message) GetPayload() isMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetPing() *Ping {
	if x, ok := m.GetPayload().(*Message_Ping); ok {
		return x.Ping
	}
	return nil
}

func (m *Message) GetAck() *Ack {
	if x, ok := m.GetPayload().(*Message_Ack); ok {
		return x.Ack
	}
	return nil
}

func (m *Message) GetIndirectPing() *IndirectPing {
	if x, ok := m.GetPayload().(*Message_IndirectPing); ok {
		return x.IndirectPing
	}
	return nil
}

func (m *Message) GetPiggyBack() *PiggyBack {
	if m != nil {
		return m.PiggyBack
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_Ping)(nil),
		(*Message_Ack)(nil),
		(*Message_IndirectPing)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// payload
	switch x := m.Payload.(type) {
	case *Message_Ping:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ping); err != nil {
			return err
		}
	case *Message_Ack:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ack); err != nil {
			return err
		}
	case *Message_IndirectPing:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.IndirectPing); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Payload has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 2: // payload.ping
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Ping)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Ping{msg}
		return true, err
	case 3: // payload.ack
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Ack)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Ack{msg}
		return true, err
	case 4: // payload.indirect_ping
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IndirectPing)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_IndirectPing{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// payload
	switch x := m.Payload.(type) {
	case *Message_Ping:
		s := proto.Size(x.Ping)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Ack:
		s := proto.Size(x.Ack)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_IndirectPing:
		s := proto.Size(x.IndirectPing)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Ping struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

type Ack struct {
	Payload              string   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}
func (*Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *Ack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ack.Unmarshal(m, b)
}
func (m *Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ack.Marshal(b, m, deterministic)
}
func (m *Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ack.Merge(m, src)
}
func (m *Ack) XXX_Size() int {
	return xxx_messageInfo_Ack.Size(m)
}
func (m *Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_Ack proto.InternalMessageInfo

func (m *Ack) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type IndirectPing struct {
	Target               string   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Nack                 bool     `protobuf:"varint,3,opt,name=nack,proto3" json:"nack,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndirectPing) Reset()         { *m = IndirectPing{} }
func (m *IndirectPing) String() string { return proto.CompactTextString(m) }
func (*IndirectPing) ProtoMessage()    {}
func (*IndirectPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *IndirectPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndirectPing.Unmarshal(m, b)
}
func (m *IndirectPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndirectPing.Marshal(b, m, deterministic)
}
func (m *IndirectPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndirectPing.Merge(m, src)
}
func (m *IndirectPing) XXX_Size() int {
	return xxx_messageInfo_IndirectPing.Size(m)
}
func (m *IndirectPing) XXX_DiscardUnknown() {
	xxx_messageInfo_IndirectPing.DiscardUnknown(m)
}

var xxx_messageInfo_IndirectPing proto.InternalMessageInfo

func (m *IndirectPing) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *IndirectPing) GetNack() bool {
	if m != nil {
		return m.Nack
	}
	return false
}

type PiggyBack struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// status
	// 0 - Unknown, 1 - Alive, 2 - Suspected, 3 - Dead
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PiggyBack) Reset()         { *m = PiggyBack{} }
func (m *PiggyBack) String() string { return proto.CompactTextString(m) }
func (*PiggyBack) ProtoMessage()    {}
func (*PiggyBack) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *PiggyBack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PiggyBack.Unmarshal(m, b)
}
func (m *PiggyBack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PiggyBack.Marshal(b, m, deterministic)
}
func (m *PiggyBack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PiggyBack.Merge(m, src)
}
func (m *PiggyBack) XXX_Size() int {
	return xxx_messageInfo_PiggyBack.Size(m)
}
func (m *PiggyBack) XXX_DiscardUnknown() {
	xxx_messageInfo_PiggyBack.DiscardUnknown(m)
}

var xxx_messageInfo_PiggyBack proto.InternalMessageInfo

func (m *PiggyBack) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PiggyBack) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PiggyBack) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*Ping)(nil), "pb.Ping")
	proto.RegisterType((*Ack)(nil), "pb.Ack")
	proto.RegisterType((*IndirectPing)(nil), "pb.IndirectPing")
	proto.RegisterType((*PiggyBack)(nil), "pb.PiggyBack")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0xbb, 0x6a, 0xc3, 0x30,
	0x14, 0x8d, 0x1f, 0xb1, 0xa3, 0xdb, 0xb8, 0x84, 0x3b, 0x14, 0x43, 0xa1, 0x0d, 0x9e, 0x02, 0x05,
	0x0f, 0xed, 0x50, 0xe8, 0x96, 0x4c, 0xe9, 0x10, 0x28, 0xfa, 0x81, 0x22, 0x5b, 0x42, 0x08, 0xb7,
	0xb6, 0x6a, 0xa9, 0x43, 0x3e, 0xb1, 0x7f, 0x55, 0xa4, 0x28, 0x4e, 0xb6, 0x7b, 0x1e, 0xf7, 0xe8,
	0x48, 0x82, 0xe2, 0x5b, 0x18, 0xc3, 0xa4, 0xa8, 0xf5, 0x38, 0xd8, 0x01, 0x63, 0xdd, 0x54, 0x7f,
	0x11, 0xe4, 0x87, 0x13, 0x8b, 0x2b, 0x48, 0x8c, 0xf8, 0x29, 0xa3, 0x75, 0xb4, 0x21, 0xd4, 0x8d,
	0xf8, 0x00, 0xa9, 0x56, 0xbd, 0x2c, 0xe3, 0x75, 0xb4, 0xb9, 0x79, 0x5e, 0xd4, 0xba, 0xa9, 0x3f,
	0x54, 0x2f, 0xf7, 0x33, 0xea, 0x79, 0xbc, 0x87, 0x84, 0xb5, 0x5d, 0x99, 0x78, 0x39, 0x77, 0xf2,
	0xb6, 0xed, 0xf6, 0x33, 0xea, 0x58, 0x7c, 0x85, 0x42, 0xf5, 0x5c, 0x8d, 0xa2, 0xb5, 0x9f, 0x3e,
	0x25, 0xf5, 0xb6, 0x95, 0xb3, 0xbd, 0x07, 0x21, 0xa4, 0x2d, 0xd5, 0x15, 0xc6, 0x27, 0x20, 0x5a,
	0x49, 0x79, 0xdc, 0xb9, 0xec, 0xb9, 0x5f, 0x2a, 0x4e, 0x47, 0x07, 0x92, 0x5e, 0xf4, 0x1d, 0x81,
	0x5c, 0xb3, 0xe3, 0xd7, 0xc0, 0x78, 0x95, 0x41, 0xea, 0xf6, 0xab, 0x47, 0x48, 0xb6, 0x6d, 0x87,
	0xe5, 0xa4, 0xf8, 0xfe, 0x84, 0x4e, 0xc6, 0x37, 0x58, 0x5e, 0x17, 0xc0, 0x3b, 0xc8, 0x2c, 0x1b,
	0xa5, 0xb0, 0xc1, 0x18, 0x10, 0x22, 0xa4, 0xfd, 0xf9, 0x7e, 0x0b, 0xea, 0xe7, 0xea, 0x00, 0x64,
	0xea, 0x81, 0xb7, 0x10, 0x2b, 0x1e, 0x1e, 0x2c, 0x56, 0xdc, 0x05, 0x19, 0xcb, 0xec, 0xaf, 0xf1,
	0x41, 0x73, 0x1a, 0x90, 0xab, 0xc2, 0x38, 0x1f, 0x85, 0x31, 0x3e, 0x8b, 0xd0, 0x33, 0x6c, 0x32,
	0xff, 0x15, 0x2f, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xd0, 0x74, 0x2b, 0x9b, 0x01, 0x00,
	0x00,
}
