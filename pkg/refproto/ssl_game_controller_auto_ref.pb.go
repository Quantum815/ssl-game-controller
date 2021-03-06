// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_game_controller_auto_ref.proto

package refproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AutoRefRegistration is the first message that a client must send to the controller to identify itself
type AutoRefRegistration struct {
	// identifier is a unique name of the client
	Identifier *string `protobuf:"bytes,1,req,name=identifier" json:"identifier,omitempty"`
	// signature can optionally be specified to enable secure communication
	Signature            *Signature `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AutoRefRegistration) Reset()         { *m = AutoRefRegistration{} }
func (m *AutoRefRegistration) String() string { return proto.CompactTextString(m) }
func (*AutoRefRegistration) ProtoMessage()    {}
func (*AutoRefRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{0}
}

func (m *AutoRefRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoRefRegistration.Unmarshal(m, b)
}
func (m *AutoRefRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoRefRegistration.Marshal(b, m, deterministic)
}
func (m *AutoRefRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoRefRegistration.Merge(m, src)
}
func (m *AutoRefRegistration) XXX_Size() int {
	return xxx_messageInfo_AutoRefRegistration.Size(m)
}
func (m *AutoRefRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoRefRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_AutoRefRegistration proto.InternalMessageInfo

func (m *AutoRefRegistration) GetIdentifier() string {
	if m != nil && m.Identifier != nil {
		return *m.Identifier
	}
	return ""
}

func (m *AutoRefRegistration) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

// AutoRefToControllerRequest is the wrapper message for all subsequent requests from the client to the controller
type AutoRefToControllerRequest struct {
	// signature can optionally be specified to enable secure communication
	Signature *Signature `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	// game_event is an optional event that the autoRef detected during the game
	GameEvent *GameEvent `protobuf:"bytes,2,opt,name=game_event,json=gameEvent" json:"game_event,omitempty"`
	// auto_ref_message is an optional message that describes the current state or situation of the game/autoRef
	AutoRefMessage       *AutoRefMessage `protobuf:"bytes,3,opt,name=auto_ref_message,json=autoRefMessage" json:"auto_ref_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AutoRefToControllerRequest) Reset()         { *m = AutoRefToControllerRequest{} }
func (m *AutoRefToControllerRequest) String() string { return proto.CompactTextString(m) }
func (*AutoRefToControllerRequest) ProtoMessage()    {}
func (*AutoRefToControllerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{1}
}

func (m *AutoRefToControllerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoRefToControllerRequest.Unmarshal(m, b)
}
func (m *AutoRefToControllerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoRefToControllerRequest.Marshal(b, m, deterministic)
}
func (m *AutoRefToControllerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoRefToControllerRequest.Merge(m, src)
}
func (m *AutoRefToControllerRequest) XXX_Size() int {
	return xxx_messageInfo_AutoRefToControllerRequest.Size(m)
}
func (m *AutoRefToControllerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoRefToControllerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AutoRefToControllerRequest proto.InternalMessageInfo

func (m *AutoRefToControllerRequest) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *AutoRefToControllerRequest) GetGameEvent() *GameEvent {
	if m != nil {
		return m.GameEvent
	}
	return nil
}

func (m *AutoRefToControllerRequest) GetAutoRefMessage() *AutoRefMessage {
	if m != nil {
		return m.AutoRefMessage
	}
	return nil
}

// a message from autoRef, describing the current state or situation
type AutoRefMessage struct {
	// Types that are valid to be assigned to Message:
	//	*AutoRefMessage_Custom
	//	*AutoRefMessage_WaitForBots_
	Message              isAutoRefMessage_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AutoRefMessage) Reset()         { *m = AutoRefMessage{} }
func (m *AutoRefMessage) String() string { return proto.CompactTextString(m) }
func (*AutoRefMessage) ProtoMessage()    {}
func (*AutoRefMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{2}
}

func (m *AutoRefMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoRefMessage.Unmarshal(m, b)
}
func (m *AutoRefMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoRefMessage.Marshal(b, m, deterministic)
}
func (m *AutoRefMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoRefMessage.Merge(m, src)
}
func (m *AutoRefMessage) XXX_Size() int {
	return xxx_messageInfo_AutoRefMessage.Size(m)
}
func (m *AutoRefMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoRefMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AutoRefMessage proto.InternalMessageInfo

type isAutoRefMessage_Message interface {
	isAutoRefMessage_Message()
}

type AutoRefMessage_Custom struct {
	Custom string `protobuf:"bytes,1,opt,name=custom,oneof"`
}

type AutoRefMessage_WaitForBots_ struct {
	WaitForBots *AutoRefMessage_WaitForBots `protobuf:"bytes,2,opt,name=wait_for_bots,json=waitForBots,oneof"`
}

func (*AutoRefMessage_Custom) isAutoRefMessage_Message() {}

func (*AutoRefMessage_WaitForBots_) isAutoRefMessage_Message() {}

func (m *AutoRefMessage) GetMessage() isAutoRefMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *AutoRefMessage) GetCustom() string {
	if x, ok := m.GetMessage().(*AutoRefMessage_Custom); ok {
		return x.Custom
	}
	return ""
}

func (m *AutoRefMessage) GetWaitForBots() *AutoRefMessage_WaitForBots {
	if x, ok := m.GetMessage().(*AutoRefMessage_WaitForBots_); ok {
		return x.WaitForBots
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AutoRefMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AutoRefMessage_OneofMarshaler, _AutoRefMessage_OneofUnmarshaler, _AutoRefMessage_OneofSizer, []interface{}{
		(*AutoRefMessage_Custom)(nil),
		(*AutoRefMessage_WaitForBots_)(nil),
	}
}

func _AutoRefMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AutoRefMessage)
	// message
	switch x := m.Message.(type) {
	case *AutoRefMessage_Custom:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Custom)
	case *AutoRefMessage_WaitForBots_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WaitForBots); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AutoRefMessage.Message has unexpected type %T", x)
	}
	return nil
}

func _AutoRefMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AutoRefMessage)
	switch tag {
	case 1: // message.custom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Message = &AutoRefMessage_Custom{x}
		return true, err
	case 2: // message.wait_for_bots
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AutoRefMessage_WaitForBots)
		err := b.DecodeMessage(msg)
		m.Message = &AutoRefMessage_WaitForBots_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AutoRefMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AutoRefMessage)
	// message
	switch x := m.Message.(type) {
	case *AutoRefMessage_Custom:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Custom)))
		n += len(x.Custom)
	case *AutoRefMessage_WaitForBots_:
		s := proto.Size(x.WaitForBots)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// the bots that is waited for
type AutoRefMessage_WaitForBots struct {
	// the bots that are waited for
	Violators            []*AutoRefMessage_WaitForBots_Violator `protobuf:"bytes,1,rep,name=violators" json:"violators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *AutoRefMessage_WaitForBots) Reset()         { *m = AutoRefMessage_WaitForBots{} }
func (m *AutoRefMessage_WaitForBots) String() string { return proto.CompactTextString(m) }
func (*AutoRefMessage_WaitForBots) ProtoMessage()    {}
func (*AutoRefMessage_WaitForBots) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{2, 0}
}

func (m *AutoRefMessage_WaitForBots) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoRefMessage_WaitForBots.Unmarshal(m, b)
}
func (m *AutoRefMessage_WaitForBots) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoRefMessage_WaitForBots.Marshal(b, m, deterministic)
}
func (m *AutoRefMessage_WaitForBots) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoRefMessage_WaitForBots.Merge(m, src)
}
func (m *AutoRefMessage_WaitForBots) XXX_Size() int {
	return xxx_messageInfo_AutoRefMessage_WaitForBots.Size(m)
}
func (m *AutoRefMessage_WaitForBots) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoRefMessage_WaitForBots.DiscardUnknown(m)
}

var xxx_messageInfo_AutoRefMessage_WaitForBots proto.InternalMessageInfo

func (m *AutoRefMessage_WaitForBots) GetViolators() []*AutoRefMessage_WaitForBots_Violator {
	if m != nil {
		return m.Violators
	}
	return nil
}

type AutoRefMessage_WaitForBots_Violator struct {
	// the id of the violator
	BotId *BotId `protobuf:"bytes,1,req,name=bot_id,json=botId" json:"bot_id,omitempty"`
	// the distance to the next valid position
	Distance             *float32 `protobuf:"fixed32,2,req,name=distance" json:"distance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoRefMessage_WaitForBots_Violator) Reset()         { *m = AutoRefMessage_WaitForBots_Violator{} }
func (m *AutoRefMessage_WaitForBots_Violator) String() string { return proto.CompactTextString(m) }
func (*AutoRefMessage_WaitForBots_Violator) ProtoMessage()    {}
func (*AutoRefMessage_WaitForBots_Violator) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{2, 0, 0}
}

func (m *AutoRefMessage_WaitForBots_Violator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoRefMessage_WaitForBots_Violator.Unmarshal(m, b)
}
func (m *AutoRefMessage_WaitForBots_Violator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoRefMessage_WaitForBots_Violator.Marshal(b, m, deterministic)
}
func (m *AutoRefMessage_WaitForBots_Violator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoRefMessage_WaitForBots_Violator.Merge(m, src)
}
func (m *AutoRefMessage_WaitForBots_Violator) XXX_Size() int {
	return xxx_messageInfo_AutoRefMessage_WaitForBots_Violator.Size(m)
}
func (m *AutoRefMessage_WaitForBots_Violator) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoRefMessage_WaitForBots_Violator.DiscardUnknown(m)
}

var xxx_messageInfo_AutoRefMessage_WaitForBots_Violator proto.InternalMessageInfo

func (m *AutoRefMessage_WaitForBots_Violator) GetBotId() *BotId {
	if m != nil {
		return m.BotId
	}
	return nil
}

func (m *AutoRefMessage_WaitForBots_Violator) GetDistance() float32 {
	if m != nil && m.Distance != nil {
		return *m.Distance
	}
	return 0
}

func init() {
	proto.RegisterType((*AutoRefRegistration)(nil), "AutoRefRegistration")
	proto.RegisterType((*AutoRefToControllerRequest)(nil), "AutoRefToControllerRequest")
	proto.RegisterType((*AutoRefMessage)(nil), "AutoRefMessage")
	proto.RegisterType((*AutoRefMessage_WaitForBots)(nil), "AutoRefMessage.WaitForBots")
	proto.RegisterType((*AutoRefMessage_WaitForBots_Violator)(nil), "AutoRefMessage.WaitForBots.Violator")
}

func init() { proto.RegisterFile("ssl_game_controller_auto_ref.proto", fileDescriptor_2fbb3ba3bab9727c) }

var fileDescriptor_2fbb3ba3bab9727c = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0xaf, 0x93, 0x40,
	0x14, 0x7d, 0xf0, 0x62, 0x7d, 0x5c, 0xe2, 0xd3, 0x8c, 0x1b, 0xc4, 0x68, 0x08, 0x71, 0x81, 0x1b,
	0xa2, 0xec, 0xba, 0x2c, 0xa6, 0x5a, 0x17, 0x6e, 0x46, 0xa3, 0xcb, 0xc9, 0x14, 0x06, 0x32, 0x09,
	0x70, 0x75, 0xe6, 0xd2, 0xfe, 0x10, 0x37, 0xfe, 0x08, 0x7f, 0xa4, 0x69, 0xcb, 0x47, 0x6b, 0x5e,
	0xba, 0x82, 0x73, 0xee, 0xb9, 0xf7, 0x9c, 0xb9, 0x17, 0x62, 0x6b, 0x1b, 0x51, 0xcb, 0x56, 0x89,
	0x02, 0x3b, 0x32, 0xd8, 0x34, 0xca, 0x08, 0xd9, 0x13, 0x0a, 0xa3, 0xaa, 0xf4, 0xa7, 0x41, 0xc2,
	0x30, 0x7a, 0x48, 0x53, 0x60, 0xdb, 0x62, 0x37, 0x28, 0x5e, 0x4c, 0x0a, 0xb5, 0x53, 0x1d, 0x89,
	0xec, 0xdd, 0xfb, 0xe5, 0xa9, 0x14, 0x0b, 0x78, 0xbe, 0xea, 0x09, 0xb9, 0xaa, 0xb8, 0xaa, 0xb5,
	0x25, 0x23, 0x49, 0x63, 0xc7, 0x5e, 0x03, 0xe8, 0x52, 0x75, 0xa4, 0x2b, 0xad, 0x4c, 0xe0, 0x44,
	0x6e, 0xe2, 0xf1, 0x33, 0x86, 0x25, 0xe0, 0x59, 0x5d, 0x77, 0x92, 0x7a, 0xa3, 0x02, 0x37, 0x72,
	0x12, 0x3f, 0x83, 0xf4, 0xeb, 0xc8, 0xf0, 0xb9, 0x18, 0xff, 0x75, 0x20, 0x1c, 0x1c, 0xbe, 0xe1,
	0x87, 0x29, 0x20, 0x57, 0xbf, 0x7a, 0x65, 0xe9, 0x72, 0x90, 0x73, 0x65, 0x10, 0x7b, 0x0b, 0x30,
	0x3f, 0x61, 0xf2, 0xfc, 0x24, 0x5b, 0xb5, 0x3e, 0x30, 0xdc, 0xab, 0xc7, 0x5f, 0xb6, 0x84, 0x67,
	0xe3, 0x8e, 0x44, 0xab, 0xac, 0x95, 0xb5, 0x0a, 0x6e, 0x8f, 0x0d, 0x4f, 0xd3, 0x21, 0xcb, 0x97,
	0x13, 0xcd, 0xef, 0xe5, 0x05, 0x8e, 0x7f, 0xbb, 0x70, 0x7f, 0x29, 0x61, 0x01, 0x2c, 0x8a, 0xde,
	0x12, 0xb6, 0xc7, 0x7c, 0xde, 0xe6, 0x86, 0x0f, 0x98, 0xad, 0xe0, 0xc9, 0x5e, 0x6a, 0x12, 0x15,
	0x1a, 0xb1, 0x45, 0xb2, 0x43, 0xaa, 0x97, 0xff, 0x99, 0xa4, 0x3f, 0xa4, 0xa6, 0x8f, 0x68, 0x72,
	0x24, 0xbb, 0xb9, 0xe1, 0xfe, 0x7e, 0x86, 0xe1, 0x1f, 0x07, 0xfc, 0xb3, 0x32, 0xcb, 0xc1, 0xdb,
	0x69, 0x6c, 0x24, 0xa1, 0xb1, 0x81, 0x13, 0xdd, 0x26, 0x7e, 0xf6, 0xe6, 0xca, 0xb8, 0xf4, 0xfb,
	0x20, 0xe6, 0x73, 0x5b, 0xb8, 0x86, 0xbb, 0x91, 0x66, 0xaf, 0x60, 0xb1, 0x45, 0x12, 0xba, 0x3c,
	0x1e, 0xd1, 0xcf, 0x16, 0x69, 0x8e, 0xf4, 0xb9, 0xe4, 0x8f, 0xb6, 0x87, 0x0f, 0x0b, 0xe1, 0xae,
	0xd4, 0x96, 0x64, 0x57, 0x1c, 0xce, 0xe8, 0x26, 0x2e, 0x9f, 0x70, 0xee, 0xc1, 0xe3, 0x61, 0x79,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x94, 0x94, 0x88, 0x87, 0x02, 0x00, 0x00,
}
