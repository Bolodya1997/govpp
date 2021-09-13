// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              21.06.0-14~g57387f511
// source: /usr/share/vpp/api/plugins/builtinurl.api.json

// Package builtinurl contains generated bindings for API file builtinurl.api.
//
// Contents:
//   2 messages
//
package builtinurl

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "builtinurl"
	APIVersion = "1.0.0"
	VersionCrc = 0x25045d63
)

// BuiltinurlEnable defines message 'builtinurl_enable'.
type BuiltinurlEnable struct{}

func (m *BuiltinurlEnable) Reset()               { *m = BuiltinurlEnable{} }
func (*BuiltinurlEnable) GetMessageName() string { return "builtinurl_enable" }
func (*BuiltinurlEnable) GetCrcString() string   { return "51077d14" }
func (*BuiltinurlEnable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *BuiltinurlEnable) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *BuiltinurlEnable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *BuiltinurlEnable) Unmarshal(b []byte) error {
	return nil
}

// BuiltinurlEnableReply defines message 'builtinurl_enable_reply'.
type BuiltinurlEnableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *BuiltinurlEnableReply) Reset()               { *m = BuiltinurlEnableReply{} }
func (*BuiltinurlEnableReply) GetMessageName() string { return "builtinurl_enable_reply" }
func (*BuiltinurlEnableReply) GetCrcString() string   { return "e8d4e804" }
func (*BuiltinurlEnableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *BuiltinurlEnableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *BuiltinurlEnableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *BuiltinurlEnableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_builtinurl_binapi_init() }
func file_builtinurl_binapi_init() {
	api.RegisterMessage((*BuiltinurlEnable)(nil), "builtinurl_enable_51077d14")
	api.RegisterMessage((*BuiltinurlEnableReply)(nil), "builtinurl_enable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*BuiltinurlEnable)(nil),
		(*BuiltinurlEnableReply)(nil),
	}
}
