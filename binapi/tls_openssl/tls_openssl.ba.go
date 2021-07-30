// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              21.10-rc0~109-g6f6663f3b
// source: /usr/share/vpp/api/plugins/tls_openssl.api.json

// Package tls_openssl contains generated bindings for API file tls_openssl.api.
//
// Contents:
//   2 messages
//
package tls_openssl

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
	APIFile    = "tls_openssl"
	APIVersion = "2.0.0"
	VersionCrc = 0x7386fbcd
)

// TLSOpensslSetEngine defines message 'tls_openssl_set_engine'.
type TLSOpensslSetEngine struct {
	AsyncEnable uint32 `binapi:"u32,name=async_enable" json:"async_enable,omitempty"`
	Engine      []byte `binapi:"u8[64],name=engine" json:"engine,omitempty"`
	Algorithm   []byte `binapi:"u8[64],name=algorithm" json:"algorithm,omitempty"`
	Ciphers     []byte `binapi:"u8[64],name=ciphers" json:"ciphers,omitempty"`
}

func (m *TLSOpensslSetEngine) Reset()               { *m = TLSOpensslSetEngine{} }
func (*TLSOpensslSetEngine) GetMessageName() string { return "tls_openssl_set_engine" }
func (*TLSOpensslSetEngine) GetCrcString() string   { return "e34d95c1" }
func (*TLSOpensslSetEngine) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TLSOpensslSetEngine) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.AsyncEnable
	size += 1 * 64 // m.Engine
	size += 1 * 64 // m.Algorithm
	size += 1 * 64 // m.Ciphers
	return size
}
func (m *TLSOpensslSetEngine) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.AsyncEnable)
	buf.EncodeBytes(m.Engine, 64)
	buf.EncodeBytes(m.Algorithm, 64)
	buf.EncodeBytes(m.Ciphers, 64)
	return buf.Bytes(), nil
}
func (m *TLSOpensslSetEngine) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.AsyncEnable = buf.DecodeUint32()
	m.Engine = make([]byte, 64)
	copy(m.Engine, buf.DecodeBytes(len(m.Engine)))
	m.Algorithm = make([]byte, 64)
	copy(m.Algorithm, buf.DecodeBytes(len(m.Algorithm)))
	m.Ciphers = make([]byte, 64)
	copy(m.Ciphers, buf.DecodeBytes(len(m.Ciphers)))
	return nil
}

// TLSOpensslSetEngineReply defines message 'tls_openssl_set_engine_reply'.
type TLSOpensslSetEngineReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TLSOpensslSetEngineReply) Reset()               { *m = TLSOpensslSetEngineReply{} }
func (*TLSOpensslSetEngineReply) GetMessageName() string { return "tls_openssl_set_engine_reply" }
func (*TLSOpensslSetEngineReply) GetCrcString() string   { return "e8d4e804" }
func (*TLSOpensslSetEngineReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TLSOpensslSetEngineReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TLSOpensslSetEngineReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TLSOpensslSetEngineReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_tls_openssl_binapi_init() }
func file_tls_openssl_binapi_init() {
	api.RegisterMessage((*TLSOpensslSetEngine)(nil), "tls_openssl_set_engine_e34d95c1")
	api.RegisterMessage((*TLSOpensslSetEngineReply)(nil), "tls_openssl_set_engine_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*TLSOpensslSetEngine)(nil),
		(*TLSOpensslSetEngineReply)(nil),
	}
}
