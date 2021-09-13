// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              21.06.0-14~g57387f511
// source: /usr/share/vpp/api/core/pg.api.json

// Package pg contains generated bindings for API file pg.api.
//
// Contents:
//   8 messages
//
package pg

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
	APIFile    = "pg"
	APIVersion = "2.0.0"
	VersionCrc = 0x3a037a0e
)

// PgCapture defines message 'pg_capture'.
type PgCapture struct {
	InterfaceID  interface_types.InterfaceIndex `binapi:"interface_index,name=interface_id" json:"interface_id,omitempty"`
	IsEnabled    bool                           `binapi:"bool,name=is_enabled,default=true" json:"is_enabled,omitempty"`
	Count        uint32                         `binapi:"u32,name=count" json:"count,omitempty"`
	PcapFileName string                         `binapi:"string[],name=pcap_file_name" json:"pcap_file_name,omitempty"`
}

func (m *PgCapture) Reset()               { *m = PgCapture{} }
func (*PgCapture) GetMessageName() string { return "pg_capture" }
func (*PgCapture) GetCrcString() string   { return "3712fb6c" }
func (*PgCapture) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PgCapture) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                       // m.InterfaceID
	size += 1                       // m.IsEnabled
	size += 4                       // m.Count
	size += 4 + len(m.PcapFileName) // m.PcapFileName
	return size
}
func (m *PgCapture) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.InterfaceID))
	buf.EncodeBool(m.IsEnabled)
	buf.EncodeUint32(m.Count)
	buf.EncodeString(m.PcapFileName, 0)
	return buf.Bytes(), nil
}
func (m *PgCapture) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.InterfaceID = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsEnabled = buf.DecodeBool()
	m.Count = buf.DecodeUint32()
	m.PcapFileName = buf.DecodeString(0)
	return nil
}

// PgCaptureReply defines message 'pg_capture_reply'.
type PgCaptureReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PgCaptureReply) Reset()               { *m = PgCaptureReply{} }
func (*PgCaptureReply) GetMessageName() string { return "pg_capture_reply" }
func (*PgCaptureReply) GetCrcString() string   { return "e8d4e804" }
func (*PgCaptureReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PgCaptureReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PgCaptureReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PgCaptureReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// PgCreateInterface defines message 'pg_create_interface'.
type PgCreateInterface struct {
	InterfaceID interface_types.InterfaceIndex `binapi:"interface_index,name=interface_id" json:"interface_id,omitempty"`
	GsoEnabled  bool                           `binapi:"bool,name=gso_enabled" json:"gso_enabled,omitempty"`
	GsoSize     uint32                         `binapi:"u32,name=gso_size" json:"gso_size,omitempty"`
}

func (m *PgCreateInterface) Reset()               { *m = PgCreateInterface{} }
func (*PgCreateInterface) GetMessageName() string { return "pg_create_interface" }
func (*PgCreateInterface) GetCrcString() string   { return "b7c893d7" }
func (*PgCreateInterface) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PgCreateInterface) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.InterfaceID
	size += 1 // m.GsoEnabled
	size += 4 // m.GsoSize
	return size
}
func (m *PgCreateInterface) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.InterfaceID))
	buf.EncodeBool(m.GsoEnabled)
	buf.EncodeUint32(m.GsoSize)
	return buf.Bytes(), nil
}
func (m *PgCreateInterface) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.InterfaceID = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.GsoEnabled = buf.DecodeBool()
	m.GsoSize = buf.DecodeUint32()
	return nil
}

// PgCreateInterfaceReply defines message 'pg_create_interface_reply'.
type PgCreateInterfaceReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *PgCreateInterfaceReply) Reset()               { *m = PgCreateInterfaceReply{} }
func (*PgCreateInterfaceReply) GetMessageName() string { return "pg_create_interface_reply" }
func (*PgCreateInterfaceReply) GetCrcString() string   { return "5383d31f" }
func (*PgCreateInterfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PgCreateInterfaceReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *PgCreateInterfaceReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *PgCreateInterfaceReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// PgEnableDisable defines message 'pg_enable_disable'.
type PgEnableDisable struct {
	IsEnabled  bool   `binapi:"bool,name=is_enabled,default=true" json:"is_enabled,omitempty"`
	StreamName string `binapi:"string[],name=stream_name" json:"stream_name,omitempty"`
}

func (m *PgEnableDisable) Reset()               { *m = PgEnableDisable{} }
func (*PgEnableDisable) GetMessageName() string { return "pg_enable_disable" }
func (*PgEnableDisable) GetCrcString() string   { return "01f94f3a" }
func (*PgEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PgEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1                     // m.IsEnabled
	size += 4 + len(m.StreamName) // m.StreamName
	return size
}
func (m *PgEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsEnabled)
	buf.EncodeString(m.StreamName, 0)
	return buf.Bytes(), nil
}
func (m *PgEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsEnabled = buf.DecodeBool()
	m.StreamName = buf.DecodeString(0)
	return nil
}

// PgEnableDisableReply defines message 'pg_enable_disable_reply'.
type PgEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PgEnableDisableReply) Reset()               { *m = PgEnableDisableReply{} }
func (*PgEnableDisableReply) GetMessageName() string { return "pg_enable_disable_reply" }
func (*PgEnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*PgEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PgEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PgEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PgEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// PgInterfaceEnableDisableCoalesce defines message 'pg_interface_enable_disable_coalesce'.
// InProgress: the message form may change in the future versions
type PgInterfaceEnableDisableCoalesce struct {
	SwIfIndex       interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	CoalesceEnabled bool                           `binapi:"bool,name=coalesce_enabled" json:"coalesce_enabled,omitempty"`
}

func (m *PgInterfaceEnableDisableCoalesce) Reset() { *m = PgInterfaceEnableDisableCoalesce{} }
func (*PgInterfaceEnableDisableCoalesce) GetMessageName() string {
	return "pg_interface_enable_disable_coalesce"
}
func (*PgInterfaceEnableDisableCoalesce) GetCrcString() string { return "a2ef99e7" }
func (*PgInterfaceEnableDisableCoalesce) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PgInterfaceEnableDisableCoalesce) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.CoalesceEnabled
	return size
}
func (m *PgInterfaceEnableDisableCoalesce) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.CoalesceEnabled)
	return buf.Bytes(), nil
}
func (m *PgInterfaceEnableDisableCoalesce) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.CoalesceEnabled = buf.DecodeBool()
	return nil
}

// PgInterfaceEnableDisableCoalesceReply defines message 'pg_interface_enable_disable_coalesce_reply'.
// InProgress: the message form may change in the future versions
type PgInterfaceEnableDisableCoalesceReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PgInterfaceEnableDisableCoalesceReply) Reset() { *m = PgInterfaceEnableDisableCoalesceReply{} }
func (*PgInterfaceEnableDisableCoalesceReply) GetMessageName() string {
	return "pg_interface_enable_disable_coalesce_reply"
}
func (*PgInterfaceEnableDisableCoalesceReply) GetCrcString() string { return "e8d4e804" }
func (*PgInterfaceEnableDisableCoalesceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PgInterfaceEnableDisableCoalesceReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PgInterfaceEnableDisableCoalesceReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PgInterfaceEnableDisableCoalesceReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_pg_binapi_init() }
func file_pg_binapi_init() {
	api.RegisterMessage((*PgCapture)(nil), "pg_capture_3712fb6c")
	api.RegisterMessage((*PgCaptureReply)(nil), "pg_capture_reply_e8d4e804")
	api.RegisterMessage((*PgCreateInterface)(nil), "pg_create_interface_b7c893d7")
	api.RegisterMessage((*PgCreateInterfaceReply)(nil), "pg_create_interface_reply_5383d31f")
	api.RegisterMessage((*PgEnableDisable)(nil), "pg_enable_disable_01f94f3a")
	api.RegisterMessage((*PgEnableDisableReply)(nil), "pg_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*PgInterfaceEnableDisableCoalesce)(nil), "pg_interface_enable_disable_coalesce_a2ef99e7")
	api.RegisterMessage((*PgInterfaceEnableDisableCoalesceReply)(nil), "pg_interface_enable_disable_coalesce_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PgCapture)(nil),
		(*PgCaptureReply)(nil),
		(*PgCreateInterface)(nil),
		(*PgCreateInterfaceReply)(nil),
		(*PgEnableDisable)(nil),
		(*PgEnableDisableReply)(nil),
		(*PgInterfaceEnableDisableCoalesce)(nil),
		(*PgInterfaceEnableDisableCoalesceReply)(nil),
	}
}
