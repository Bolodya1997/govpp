// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              21.10-rc0~109-g6f6663f3b
// source: /usr/share/vpp/api/plugins/flowprobe.api.json

// Package flowprobe contains generated bindings for API file flowprobe.api.
//
// Contents:
//   2 enums
//   4 messages
//
package flowprobe

import (
	"strconv"

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
	APIFile    = "flowprobe"
	APIVersion = "1.0.0"
	VersionCrc = 0x8da9f1c
)

// FlowprobeRecordFlags defines enum 'flowprobe_record_flags'.
type FlowprobeRecordFlags uint8

const (
	FLOWPROBE_RECORD_FLAG_L2 FlowprobeRecordFlags = 1
	FLOWPROBE_RECORD_FLAG_L3 FlowprobeRecordFlags = 2
	FLOWPROBE_RECORD_FLAG_L4 FlowprobeRecordFlags = 4
)

var (
	FlowprobeRecordFlags_name = map[uint8]string{
		1: "FLOWPROBE_RECORD_FLAG_L2",
		2: "FLOWPROBE_RECORD_FLAG_L3",
		4: "FLOWPROBE_RECORD_FLAG_L4",
	}
	FlowprobeRecordFlags_value = map[string]uint8{
		"FLOWPROBE_RECORD_FLAG_L2": 1,
		"FLOWPROBE_RECORD_FLAG_L3": 2,
		"FLOWPROBE_RECORD_FLAG_L4": 4,
	}
)

func (x FlowprobeRecordFlags) String() string {
	s, ok := FlowprobeRecordFlags_name[uint8(x)]
	if ok {
		return s
	}
	str := func(n uint8) string {
		s, ok := FlowprobeRecordFlags_name[uint8(n)]
		if ok {
			return s
		}
		return "FlowprobeRecordFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint8(0); i <= 8; i++ {
		val := uint8(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint8(x))
	}
	return s
}

// FlowprobeWhichFlags defines enum 'flowprobe_which_flags'.
type FlowprobeWhichFlags uint8

const (
	FLOWPROBE_WHICH_FLAG_IP4 FlowprobeWhichFlags = 1
	FLOWPROBE_WHICH_FLAG_L2  FlowprobeWhichFlags = 2
	FLOWPROBE_WHICH_FLAG_IP6 FlowprobeWhichFlags = 4
)

var (
	FlowprobeWhichFlags_name = map[uint8]string{
		1: "FLOWPROBE_WHICH_FLAG_IP4",
		2: "FLOWPROBE_WHICH_FLAG_L2",
		4: "FLOWPROBE_WHICH_FLAG_IP6",
	}
	FlowprobeWhichFlags_value = map[string]uint8{
		"FLOWPROBE_WHICH_FLAG_IP4": 1,
		"FLOWPROBE_WHICH_FLAG_L2":  2,
		"FLOWPROBE_WHICH_FLAG_IP6": 4,
	}
)

func (x FlowprobeWhichFlags) String() string {
	s, ok := FlowprobeWhichFlags_name[uint8(x)]
	if ok {
		return s
	}
	str := func(n uint8) string {
		s, ok := FlowprobeWhichFlags_name[uint8(n)]
		if ok {
			return s
		}
		return "FlowprobeWhichFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint8(0); i <= 8; i++ {
		val := uint8(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint8(x))
	}
	return s
}

// FlowprobeParams defines message 'flowprobe_params'.
type FlowprobeParams struct {
	RecordFlags  FlowprobeRecordFlags `binapi:"flowprobe_record_flags,name=record_flags" json:"record_flags,omitempty"`
	ActiveTimer  uint32               `binapi:"u32,name=active_timer" json:"active_timer,omitempty"`
	PassiveTimer uint32               `binapi:"u32,name=passive_timer" json:"passive_timer,omitempty"`
}

func (m *FlowprobeParams) Reset()               { *m = FlowprobeParams{} }
func (*FlowprobeParams) GetMessageName() string { return "flowprobe_params" }
func (*FlowprobeParams) GetCrcString() string   { return "baa46c09" }
func (*FlowprobeParams) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *FlowprobeParams) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.RecordFlags
	size += 4 // m.ActiveTimer
	size += 4 // m.PassiveTimer
	return size
}
func (m *FlowprobeParams) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.RecordFlags))
	buf.EncodeUint32(m.ActiveTimer)
	buf.EncodeUint32(m.PassiveTimer)
	return buf.Bytes(), nil
}
func (m *FlowprobeParams) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.RecordFlags = FlowprobeRecordFlags(buf.DecodeUint8())
	m.ActiveTimer = buf.DecodeUint32()
	m.PassiveTimer = buf.DecodeUint32()
	return nil
}

// FlowprobeParamsReply defines message 'flowprobe_params_reply'.
type FlowprobeParamsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *FlowprobeParamsReply) Reset()               { *m = FlowprobeParamsReply{} }
func (*FlowprobeParamsReply) GetMessageName() string { return "flowprobe_params_reply" }
func (*FlowprobeParamsReply) GetCrcString() string   { return "e8d4e804" }
func (*FlowprobeParamsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *FlowprobeParamsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *FlowprobeParamsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *FlowprobeParamsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// FlowprobeTxInterfaceAddDel defines message 'flowprobe_tx_interface_add_del'.
type FlowprobeTxInterfaceAddDel struct {
	IsAdd     bool                           `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	Which     FlowprobeWhichFlags            `binapi:"flowprobe_which_flags,name=which" json:"which,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *FlowprobeTxInterfaceAddDel) Reset()               { *m = FlowprobeTxInterfaceAddDel{} }
func (*FlowprobeTxInterfaceAddDel) GetMessageName() string { return "flowprobe_tx_interface_add_del" }
func (*FlowprobeTxInterfaceAddDel) GetCrcString() string   { return "b782c976" }
func (*FlowprobeTxInterfaceAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *FlowprobeTxInterfaceAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 1 // m.Which
	size += 4 // m.SwIfIndex
	return size
}
func (m *FlowprobeTxInterfaceAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint8(uint8(m.Which))
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *FlowprobeTxInterfaceAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Which = FlowprobeWhichFlags(buf.DecodeUint8())
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// FlowprobeTxInterfaceAddDelReply defines message 'flowprobe_tx_interface_add_del_reply'.
type FlowprobeTxInterfaceAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *FlowprobeTxInterfaceAddDelReply) Reset() { *m = FlowprobeTxInterfaceAddDelReply{} }
func (*FlowprobeTxInterfaceAddDelReply) GetMessageName() string {
	return "flowprobe_tx_interface_add_del_reply"
}
func (*FlowprobeTxInterfaceAddDelReply) GetCrcString() string { return "e8d4e804" }
func (*FlowprobeTxInterfaceAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *FlowprobeTxInterfaceAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *FlowprobeTxInterfaceAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *FlowprobeTxInterfaceAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_flowprobe_binapi_init() }
func file_flowprobe_binapi_init() {
	api.RegisterMessage((*FlowprobeParams)(nil), "flowprobe_params_baa46c09")
	api.RegisterMessage((*FlowprobeParamsReply)(nil), "flowprobe_params_reply_e8d4e804")
	api.RegisterMessage((*FlowprobeTxInterfaceAddDel)(nil), "flowprobe_tx_interface_add_del_b782c976")
	api.RegisterMessage((*FlowprobeTxInterfaceAddDelReply)(nil), "flowprobe_tx_interface_add_del_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*FlowprobeParams)(nil),
		(*FlowprobeParamsReply)(nil),
		(*FlowprobeTxInterfaceAddDel)(nil),
		(*FlowprobeTxInterfaceAddDelReply)(nil),
	}
}
