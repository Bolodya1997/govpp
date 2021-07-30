// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              21.10-rc0~109-g6f6663f3b
// source: /usr/share/vpp/api/core/udp.api.json

// Package udp contains generated bindings for API file udp.api.
//
// Contents:
//   1 enum
//   2 structs
//   8 messages
//
package udp

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	ip_types "github.com/edwarnicke/govpp/binapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "udp"
	APIVersion = "1.1.0"
	VersionCrc = 0x2d3db5fa
)

// UDPDecapNextProto defines enum 'udp_decap_next_proto'.
type UDPDecapNextProto uint32

const (
	UDP_API_DECAP_PROTO_IP4  UDPDecapNextProto = 1
	UDP_API_DECAP_PROTO_IP6  UDPDecapNextProto = 2
	UDP_API_DECAP_PROTO_MPLS UDPDecapNextProto = 3
)

var (
	UDPDecapNextProto_name = map[uint32]string{
		1: "UDP_API_DECAP_PROTO_IP4",
		2: "UDP_API_DECAP_PROTO_IP6",
		3: "UDP_API_DECAP_PROTO_MPLS",
	}
	UDPDecapNextProto_value = map[string]uint32{
		"UDP_API_DECAP_PROTO_IP4":  1,
		"UDP_API_DECAP_PROTO_IP6":  2,
		"UDP_API_DECAP_PROTO_MPLS": 3,
	}
)

func (x UDPDecapNextProto) String() string {
	s, ok := UDPDecapNextProto_name[uint32(x)]
	if ok {
		return s
	}
	return "UDPDecapNextProto(" + strconv.Itoa(int(x)) + ")"
}

// UDPDecap defines type 'udp_decap'.
type UDPDecap struct {
	IsIP4     uint8             `binapi:"u8,name=is_ip4" json:"is_ip4,omitempty"`
	Port      uint16            `binapi:"u16,name=port" json:"port,omitempty"`
	NextProto UDPDecapNextProto `binapi:"udp_decap_next_proto,name=next_proto" json:"next_proto,omitempty"`
}

// UDPEncap defines type 'udp_encap'.
type UDPEncap struct {
	TableID uint32           `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	SrcPort uint16           `binapi:"u16,name=src_port" json:"src_port,omitempty"`
	DstPort uint16           `binapi:"u16,name=dst_port" json:"dst_port,omitempty"`
	SrcIP   ip_types.Address `binapi:"address,name=src_ip" json:"src_ip,omitempty"`
	DstIP   ip_types.Address `binapi:"address,name=dst_ip" json:"dst_ip,omitempty"`
	ID      uint32           `binapi:"u32,name=id" json:"id,omitempty"`
}

// UDPDecapAddDel defines message 'udp_decap_add_del'.
type UDPDecapAddDel struct {
	IsAdd    bool     `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	UDPDecap UDPDecap `binapi:"udp_decap,name=udp_decap" json:"udp_decap,omitempty"`
}

func (m *UDPDecapAddDel) Reset()               { *m = UDPDecapAddDel{} }
func (*UDPDecapAddDel) GetMessageName() string { return "udp_decap_add_del" }
func (*UDPDecapAddDel) GetCrcString() string   { return "d155c6ac" }
func (*UDPDecapAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *UDPDecapAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 1 // m.UDPDecap.IsIP4
	size += 2 // m.UDPDecap.Port
	size += 4 // m.UDPDecap.NextProto
	return size
}
func (m *UDPDecapAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint8(m.UDPDecap.IsIP4)
	buf.EncodeUint16(m.UDPDecap.Port)
	buf.EncodeUint32(uint32(m.UDPDecap.NextProto))
	return buf.Bytes(), nil
}
func (m *UDPDecapAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.UDPDecap.IsIP4 = buf.DecodeUint8()
	m.UDPDecap.Port = buf.DecodeUint16()
	m.UDPDecap.NextProto = UDPDecapNextProto(buf.DecodeUint32())
	return nil
}

// UDPDecapAddDelReply defines message 'udp_decap_add_del_reply'.
type UDPDecapAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *UDPDecapAddDelReply) Reset()               { *m = UDPDecapAddDelReply{} }
func (*UDPDecapAddDelReply) GetMessageName() string { return "udp_decap_add_del_reply" }
func (*UDPDecapAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*UDPDecapAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *UDPDecapAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *UDPDecapAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *UDPDecapAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// UDPEncapAdd defines message 'udp_encap_add'.
type UDPEncapAdd struct {
	UDPEncap UDPEncap `binapi:"udp_encap,name=udp_encap" json:"udp_encap,omitempty"`
}

func (m *UDPEncapAdd) Reset()               { *m = UDPEncapAdd{} }
func (*UDPEncapAdd) GetMessageName() string { return "udp_encap_add" }
func (*UDPEncapAdd) GetCrcString() string   { return "f74a60b1" }
func (*UDPEncapAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *UDPEncapAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.UDPEncap.TableID
	size += 2      // m.UDPEncap.SrcPort
	size += 2      // m.UDPEncap.DstPort
	size += 1      // m.UDPEncap.SrcIP.Af
	size += 1 * 16 // m.UDPEncap.SrcIP.Un
	size += 1      // m.UDPEncap.DstIP.Af
	size += 1 * 16 // m.UDPEncap.DstIP.Un
	size += 4      // m.UDPEncap.ID
	return size
}
func (m *UDPEncapAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.UDPEncap.TableID)
	buf.EncodeUint16(m.UDPEncap.SrcPort)
	buf.EncodeUint16(m.UDPEncap.DstPort)
	buf.EncodeUint8(uint8(m.UDPEncap.SrcIP.Af))
	buf.EncodeBytes(m.UDPEncap.SrcIP.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.UDPEncap.DstIP.Af))
	buf.EncodeBytes(m.UDPEncap.DstIP.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.UDPEncap.ID)
	return buf.Bytes(), nil
}
func (m *UDPEncapAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.UDPEncap.TableID = buf.DecodeUint32()
	m.UDPEncap.SrcPort = buf.DecodeUint16()
	m.UDPEncap.DstPort = buf.DecodeUint16()
	m.UDPEncap.SrcIP.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.UDPEncap.SrcIP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.UDPEncap.DstIP.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.UDPEncap.DstIP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.UDPEncap.ID = buf.DecodeUint32()
	return nil
}

// UDPEncapAddReply defines message 'udp_encap_add_reply'.
type UDPEncapAddReply struct {
	Retval int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ID     uint32 `binapi:"u32,name=id" json:"id,omitempty"`
}

func (m *UDPEncapAddReply) Reset()               { *m = UDPEncapAddReply{} }
func (*UDPEncapAddReply) GetMessageName() string { return "udp_encap_add_reply" }
func (*UDPEncapAddReply) GetCrcString() string   { return "e2fc8294" }
func (*UDPEncapAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *UDPEncapAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.ID
	return size
}
func (m *UDPEncapAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.ID)
	return buf.Bytes(), nil
}
func (m *UDPEncapAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.ID = buf.DecodeUint32()
	return nil
}

// UDPEncapDel defines message 'udp_encap_del'.
type UDPEncapDel struct {
	ID uint32 `binapi:"u32,name=id" json:"id,omitempty"`
}

func (m *UDPEncapDel) Reset()               { *m = UDPEncapDel{} }
func (*UDPEncapDel) GetMessageName() string { return "udp_encap_del" }
func (*UDPEncapDel) GetCrcString() string   { return "3a91bde5" }
func (*UDPEncapDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *UDPEncapDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.ID
	return size
}
func (m *UDPEncapDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.ID)
	return buf.Bytes(), nil
}
func (m *UDPEncapDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ID = buf.DecodeUint32()
	return nil
}

// UDPEncapDelReply defines message 'udp_encap_del_reply'.
type UDPEncapDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *UDPEncapDelReply) Reset()               { *m = UDPEncapDelReply{} }
func (*UDPEncapDelReply) GetMessageName() string { return "udp_encap_del_reply" }
func (*UDPEncapDelReply) GetCrcString() string   { return "e8d4e804" }
func (*UDPEncapDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *UDPEncapDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *UDPEncapDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *UDPEncapDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// UDPEncapDetails defines message 'udp_encap_details'.
type UDPEncapDetails struct {
	UDPEncap UDPEncap `binapi:"udp_encap,name=udp_encap" json:"udp_encap,omitempty"`
}

func (m *UDPEncapDetails) Reset()               { *m = UDPEncapDetails{} }
func (*UDPEncapDetails) GetMessageName() string { return "udp_encap_details" }
func (*UDPEncapDetails) GetCrcString() string   { return "8cfb9c76" }
func (*UDPEncapDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *UDPEncapDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.UDPEncap.TableID
	size += 2      // m.UDPEncap.SrcPort
	size += 2      // m.UDPEncap.DstPort
	size += 1      // m.UDPEncap.SrcIP.Af
	size += 1 * 16 // m.UDPEncap.SrcIP.Un
	size += 1      // m.UDPEncap.DstIP.Af
	size += 1 * 16 // m.UDPEncap.DstIP.Un
	size += 4      // m.UDPEncap.ID
	return size
}
func (m *UDPEncapDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.UDPEncap.TableID)
	buf.EncodeUint16(m.UDPEncap.SrcPort)
	buf.EncodeUint16(m.UDPEncap.DstPort)
	buf.EncodeUint8(uint8(m.UDPEncap.SrcIP.Af))
	buf.EncodeBytes(m.UDPEncap.SrcIP.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.UDPEncap.DstIP.Af))
	buf.EncodeBytes(m.UDPEncap.DstIP.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.UDPEncap.ID)
	return buf.Bytes(), nil
}
func (m *UDPEncapDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.UDPEncap.TableID = buf.DecodeUint32()
	m.UDPEncap.SrcPort = buf.DecodeUint16()
	m.UDPEncap.DstPort = buf.DecodeUint16()
	m.UDPEncap.SrcIP.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.UDPEncap.SrcIP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.UDPEncap.DstIP.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.UDPEncap.DstIP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.UDPEncap.ID = buf.DecodeUint32()
	return nil
}

// UDPEncapDump defines message 'udp_encap_dump'.
type UDPEncapDump struct{}

func (m *UDPEncapDump) Reset()               { *m = UDPEncapDump{} }
func (*UDPEncapDump) GetMessageName() string { return "udp_encap_dump" }
func (*UDPEncapDump) GetCrcString() string   { return "51077d14" }
func (*UDPEncapDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *UDPEncapDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *UDPEncapDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *UDPEncapDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_udp_binapi_init() }
func file_udp_binapi_init() {
	api.RegisterMessage((*UDPDecapAddDel)(nil), "udp_decap_add_del_d155c6ac")
	api.RegisterMessage((*UDPDecapAddDelReply)(nil), "udp_decap_add_del_reply_e8d4e804")
	api.RegisterMessage((*UDPEncapAdd)(nil), "udp_encap_add_f74a60b1")
	api.RegisterMessage((*UDPEncapAddReply)(nil), "udp_encap_add_reply_e2fc8294")
	api.RegisterMessage((*UDPEncapDel)(nil), "udp_encap_del_3a91bde5")
	api.RegisterMessage((*UDPEncapDelReply)(nil), "udp_encap_del_reply_e8d4e804")
	api.RegisterMessage((*UDPEncapDetails)(nil), "udp_encap_details_8cfb9c76")
	api.RegisterMessage((*UDPEncapDump)(nil), "udp_encap_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*UDPDecapAddDel)(nil),
		(*UDPDecapAddDelReply)(nil),
		(*UDPEncapAdd)(nil),
		(*UDPEncapAddReply)(nil),
		(*UDPEncapDel)(nil),
		(*UDPEncapDelReply)(nil),
		(*UDPEncapDetails)(nil),
		(*UDPEncapDump)(nil),
	}
}
