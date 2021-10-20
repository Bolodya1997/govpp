// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package l2e contains generated bindings for API file l2e.api.
//
// Contents:
//   2 messages
//
package l2e

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/edwarnicke/govpp/binapi/interface_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "l2e"
	APIVersion = "1.0.0"
	VersionCrc = 0x6e8abdfb
)

// L2Emulation defines message 'l2_emulation'.
// InProgress: the message form may change in the future versions
type L2Emulation struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Enable    bool                           `binapi:"bool,name=enable" json:"enable,omitempty"`
}

func (m *L2Emulation) Reset()               { *m = L2Emulation{} }
func (*L2Emulation) GetMessageName() string { return "l2_emulation" }
func (*L2Emulation) GetCrcString() string   { return "ae6cfcfb" }
func (*L2Emulation) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *L2Emulation) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.Enable
	return size
}
func (m *L2Emulation) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.Enable)
	return buf.Bytes(), nil
}
func (m *L2Emulation) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Enable = buf.DecodeBool()
	return nil
}

// L2EmulationReply defines message 'l2_emulation_reply'.
// InProgress: the message form may change in the future versions
type L2EmulationReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *L2EmulationReply) Reset()               { *m = L2EmulationReply{} }
func (*L2EmulationReply) GetMessageName() string { return "l2_emulation_reply" }
func (*L2EmulationReply) GetCrcString() string   { return "e8d4e804" }
func (*L2EmulationReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *L2EmulationReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *L2EmulationReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *L2EmulationReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_l2e_binapi_init() }
func file_l2e_binapi_init() {
	api.RegisterMessage((*L2Emulation)(nil), "l2_emulation_ae6cfcfb")
	api.RegisterMessage((*L2EmulationReply)(nil), "l2_emulation_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*L2Emulation)(nil),
		(*L2EmulationReply)(nil),
	}
}
