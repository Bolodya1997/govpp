// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package vrrp contains generated bindings for API file vrrp.api.
//
// Contents:
//   2 enums
//   5 structs
//  17 messages
//
package vrrp

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	ethernet_types "github.com/edwarnicke/govpp/binapi/ethernet_types"
	interface_types "github.com/edwarnicke/govpp/binapi/interface_types"
	ip_types "github.com/edwarnicke/govpp/binapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "vrrp"
	APIVersion = "1.0.1"
	VersionCrc = 0x6a3c71cd
)

// VrrpVrFlags defines enum 'vrrp_vr_flags'.
type VrrpVrFlags uint32

const (
	VRRP_API_VR_PREEMPT VrrpVrFlags = 1
	VRRP_API_VR_ACCEPT  VrrpVrFlags = 2
	VRRP_API_VR_UNICAST VrrpVrFlags = 4
	VRRP_API_VR_IPV6    VrrpVrFlags = 8
)

var (
	VrrpVrFlags_name = map[uint32]string{
		1: "VRRP_API_VR_PREEMPT",
		2: "VRRP_API_VR_ACCEPT",
		4: "VRRP_API_VR_UNICAST",
		8: "VRRP_API_VR_IPV6",
	}
	VrrpVrFlags_value = map[string]uint32{
		"VRRP_API_VR_PREEMPT": 1,
		"VRRP_API_VR_ACCEPT":  2,
		"VRRP_API_VR_UNICAST": 4,
		"VRRP_API_VR_IPV6":    8,
	}
)

func (x VrrpVrFlags) String() string {
	s, ok := VrrpVrFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := VrrpVrFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "VrrpVrFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// VrrpVrState defines enum 'vrrp_vr_state'.
type VrrpVrState uint32

const (
	VRRP_API_VR_STATE_INIT      VrrpVrState = 0
	VRRP_API_VR_STATE_BACKUP    VrrpVrState = 1
	VRRP_API_VR_STATE_MASTER    VrrpVrState = 2
	VRRP_API_VR_STATE_INTF_DOWN VrrpVrState = 3
)

var (
	VrrpVrState_name = map[uint32]string{
		0: "VRRP_API_VR_STATE_INIT",
		1: "VRRP_API_VR_STATE_BACKUP",
		2: "VRRP_API_VR_STATE_MASTER",
		3: "VRRP_API_VR_STATE_INTF_DOWN",
	}
	VrrpVrState_value = map[string]uint32{
		"VRRP_API_VR_STATE_INIT":      0,
		"VRRP_API_VR_STATE_BACKUP":    1,
		"VRRP_API_VR_STATE_MASTER":    2,
		"VRRP_API_VR_STATE_INTF_DOWN": 3,
	}
)

func (x VrrpVrState) String() string {
	s, ok := VrrpVrState_name[uint32(x)]
	if ok {
		return s
	}
	return "VrrpVrState(" + strconv.Itoa(int(x)) + ")"
}

// VrrpVrConf defines type 'vrrp_vr_conf'.
type VrrpVrConf struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	VrID      uint8                          `binapi:"u8,name=vr_id" json:"vr_id,omitempty"`
	Priority  uint8                          `binapi:"u8,name=priority" json:"priority,omitempty"`
	Interval  uint16                         `binapi:"u16,name=interval" json:"interval,omitempty"`
	Flags     VrrpVrFlags                    `binapi:"vrrp_vr_flags,name=flags" json:"flags,omitempty"`
}

// VrrpVrKey defines type 'vrrp_vr_key'.
type VrrpVrKey struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	VrID      uint8                          `binapi:"u8,name=vr_id" json:"vr_id,omitempty"`
	IsIPv6    uint8                          `binapi:"u8,name=is_ipv6" json:"is_ipv6,omitempty"`
}

// VrrpVrRuntime defines type 'vrrp_vr_runtime'.
type VrrpVrRuntime struct {
	State         VrrpVrState               `binapi:"vrrp_vr_state,name=state" json:"state,omitempty"`
	MasterAdvInt  uint16                    `binapi:"u16,name=master_adv_int" json:"master_adv_int,omitempty"`
	Skew          uint16                    `binapi:"u16,name=skew" json:"skew,omitempty"`
	MasterDownInt uint16                    `binapi:"u16,name=master_down_int" json:"master_down_int,omitempty"`
	Mac           ethernet_types.MacAddress `binapi:"mac_address,name=mac" json:"mac,omitempty"`
	Tracking      VrrpVrTracking            `binapi:"vrrp_vr_tracking,name=tracking" json:"tracking,omitempty"`
}

// VrrpVrTrackIf defines type 'vrrp_vr_track_if'.
type VrrpVrTrackIf struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Priority  uint8                          `binapi:"u8,name=priority" json:"priority,omitempty"`
}

// VrrpVrTracking defines type 'vrrp_vr_tracking'.
type VrrpVrTracking struct {
	InterfacesDec uint32 `binapi:"u32,name=interfaces_dec" json:"interfaces_dec,omitempty"`
	Priority      uint8  `binapi:"u8,name=priority" json:"priority,omitempty"`
}

// VrrpVrAddDel defines message 'vrrp_vr_add_del'.
type VrrpVrAddDel struct {
	IsAdd     uint8                          `binapi:"u8,name=is_add" json:"is_add,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	VrID      uint8                          `binapi:"u8,name=vr_id" json:"vr_id,omitempty"`
	Priority  uint8                          `binapi:"u8,name=priority" json:"priority,omitempty"`
	Interval  uint16                         `binapi:"u16,name=interval" json:"interval,omitempty"`
	Flags     VrrpVrFlags                    `binapi:"vrrp_vr_flags,name=flags" json:"flags,omitempty"`
	NAddrs    uint8                          `binapi:"u8,name=n_addrs" json:"-"`
	Addrs     []ip_types.Address             `binapi:"address[n_addrs],name=addrs" json:"addrs,omitempty"`
}

func (m *VrrpVrAddDel) Reset()               { *m = VrrpVrAddDel{} }
func (*VrrpVrAddDel) GetMessageName() string { return "vrrp_vr_add_del" }
func (*VrrpVrAddDel) GetCrcString() string   { return "c5cf15aa" }
func (*VrrpVrAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VrrpVrAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 4 // m.SwIfIndex
	size += 1 // m.VrID
	size += 1 // m.Priority
	size += 2 // m.Interval
	size += 4 // m.Flags
	size += 1 // m.NAddrs
	for j1 := 0; j1 < len(m.Addrs); j1++ {
		var s1 ip_types.Address
		_ = s1
		if j1 < len(m.Addrs) {
			s1 = m.Addrs[j1]
		}
		size += 1      // s1.Af
		size += 1 * 16 // s1.Un
	}
	return size
}
func (m *VrrpVrAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.IsAdd)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(m.VrID)
	buf.EncodeUint8(m.Priority)
	buf.EncodeUint16(m.Interval)
	buf.EncodeUint32(uint32(m.Flags))
	buf.EncodeUint8(uint8(len(m.Addrs)))
	for j0 := 0; j0 < len(m.Addrs); j0++ {
		var v0 ip_types.Address // Addrs
		if j0 < len(m.Addrs) {
			v0 = m.Addrs[j0]
		}
		buf.EncodeUint8(uint8(v0.Af))
		buf.EncodeBytes(v0.Un.XXX_UnionData[:], 16)
	}
	return buf.Bytes(), nil
}
func (m *VrrpVrAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeUint8()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.VrID = buf.DecodeUint8()
	m.Priority = buf.DecodeUint8()
	m.Interval = buf.DecodeUint16()
	m.Flags = VrrpVrFlags(buf.DecodeUint32())
	m.NAddrs = buf.DecodeUint8()
	m.Addrs = make([]ip_types.Address, m.NAddrs)
	for j0 := 0; j0 < len(m.Addrs); j0++ {
		m.Addrs[j0].Af = ip_types.AddressFamily(buf.DecodeUint8())
		copy(m.Addrs[j0].Un.XXX_UnionData[:], buf.DecodeBytes(16))
	}
	return nil
}

// VrrpVrAddDelReply defines message 'vrrp_vr_add_del_reply'.
type VrrpVrAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *VrrpVrAddDelReply) Reset()               { *m = VrrpVrAddDelReply{} }
func (*VrrpVrAddDelReply) GetMessageName() string { return "vrrp_vr_add_del_reply" }
func (*VrrpVrAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*VrrpVrAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VrrpVrAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *VrrpVrAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *VrrpVrAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// VrrpVrDetails defines message 'vrrp_vr_details'.
type VrrpVrDetails struct {
	Config  VrrpVrConf         `binapi:"vrrp_vr_conf,name=config" json:"config,omitempty"`
	Runtime VrrpVrRuntime      `binapi:"vrrp_vr_runtime,name=runtime" json:"runtime,omitempty"`
	NAddrs  uint8              `binapi:"u8,name=n_addrs" json:"-"`
	Addrs   []ip_types.Address `binapi:"address[n_addrs],name=addrs" json:"addrs,omitempty"`
}

func (m *VrrpVrDetails) Reset()               { *m = VrrpVrDetails{} }
func (*VrrpVrDetails) GetMessageName() string { return "vrrp_vr_details" }
func (*VrrpVrDetails) GetCrcString() string   { return "46edcebd" }
func (*VrrpVrDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VrrpVrDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.Config.SwIfIndex
	size += 1     // m.Config.VrID
	size += 1     // m.Config.Priority
	size += 2     // m.Config.Interval
	size += 4     // m.Config.Flags
	size += 4     // m.Runtime.State
	size += 2     // m.Runtime.MasterAdvInt
	size += 2     // m.Runtime.Skew
	size += 2     // m.Runtime.MasterDownInt
	size += 1 * 6 // m.Runtime.Mac
	size += 4     // m.Runtime.Tracking.InterfacesDec
	size += 1     // m.Runtime.Tracking.Priority
	size += 1     // m.NAddrs
	for j1 := 0; j1 < len(m.Addrs); j1++ {
		var s1 ip_types.Address
		_ = s1
		if j1 < len(m.Addrs) {
			s1 = m.Addrs[j1]
		}
		size += 1      // s1.Af
		size += 1 * 16 // s1.Un
	}
	return size
}
func (m *VrrpVrDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Config.SwIfIndex))
	buf.EncodeUint8(m.Config.VrID)
	buf.EncodeUint8(m.Config.Priority)
	buf.EncodeUint16(m.Config.Interval)
	buf.EncodeUint32(uint32(m.Config.Flags))
	buf.EncodeUint32(uint32(m.Runtime.State))
	buf.EncodeUint16(m.Runtime.MasterAdvInt)
	buf.EncodeUint16(m.Runtime.Skew)
	buf.EncodeUint16(m.Runtime.MasterDownInt)
	buf.EncodeBytes(m.Runtime.Mac[:], 6)
	buf.EncodeUint32(m.Runtime.Tracking.InterfacesDec)
	buf.EncodeUint8(m.Runtime.Tracking.Priority)
	buf.EncodeUint8(uint8(len(m.Addrs)))
	for j0 := 0; j0 < len(m.Addrs); j0++ {
		var v0 ip_types.Address // Addrs
		if j0 < len(m.Addrs) {
			v0 = m.Addrs[j0]
		}
		buf.EncodeUint8(uint8(v0.Af))
		buf.EncodeBytes(v0.Un.XXX_UnionData[:], 16)
	}
	return buf.Bytes(), nil
}
func (m *VrrpVrDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Config.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Config.VrID = buf.DecodeUint8()
	m.Config.Priority = buf.DecodeUint8()
	m.Config.Interval = buf.DecodeUint16()
	m.Config.Flags = VrrpVrFlags(buf.DecodeUint32())
	m.Runtime.State = VrrpVrState(buf.DecodeUint32())
	m.Runtime.MasterAdvInt = buf.DecodeUint16()
	m.Runtime.Skew = buf.DecodeUint16()
	m.Runtime.MasterDownInt = buf.DecodeUint16()
	copy(m.Runtime.Mac[:], buf.DecodeBytes(6))
	m.Runtime.Tracking.InterfacesDec = buf.DecodeUint32()
	m.Runtime.Tracking.Priority = buf.DecodeUint8()
	m.NAddrs = buf.DecodeUint8()
	m.Addrs = make([]ip_types.Address, m.NAddrs)
	for j0 := 0; j0 < len(m.Addrs); j0++ {
		m.Addrs[j0].Af = ip_types.AddressFamily(buf.DecodeUint8())
		copy(m.Addrs[j0].Un.XXX_UnionData[:], buf.DecodeBytes(16))
	}
	return nil
}

// VrrpVrDump defines message 'vrrp_vr_dump'.
type VrrpVrDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *VrrpVrDump) Reset()               { *m = VrrpVrDump{} }
func (*VrrpVrDump) GetMessageName() string { return "vrrp_vr_dump" }
func (*VrrpVrDump) GetCrcString() string   { return "f9e6675e" }
func (*VrrpVrDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VrrpVrDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *VrrpVrDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *VrrpVrDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// VrrpVrEvent defines message 'vrrp_vr_event'.
type VrrpVrEvent struct {
	PID      uint32      `binapi:"u32,name=pid" json:"pid,omitempty"`
	Vr       VrrpVrKey   `binapi:"vrrp_vr_key,name=vr" json:"vr,omitempty"`
	OldState VrrpVrState `binapi:"vrrp_vr_state,name=old_state" json:"old_state,omitempty"`
	NewState VrrpVrState `binapi:"vrrp_vr_state,name=new_state" json:"new_state,omitempty"`
}

func (m *VrrpVrEvent) Reset()               { *m = VrrpVrEvent{} }
func (*VrrpVrEvent) GetMessageName() string { return "vrrp_vr_event" }
func (*VrrpVrEvent) GetCrcString() string   { return "c1fea6a5" }
func (*VrrpVrEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

func (m *VrrpVrEvent) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.PID
	size += 4 // m.Vr.SwIfIndex
	size += 1 // m.Vr.VrID
	size += 1 // m.Vr.IsIPv6
	size += 4 // m.OldState
	size += 4 // m.NewState
	return size
}
func (m *VrrpVrEvent) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.PID)
	buf.EncodeUint32(uint32(m.Vr.SwIfIndex))
	buf.EncodeUint8(m.Vr.VrID)
	buf.EncodeUint8(m.Vr.IsIPv6)
	buf.EncodeUint32(uint32(m.OldState))
	buf.EncodeUint32(uint32(m.NewState))
	return buf.Bytes(), nil
}
func (m *VrrpVrEvent) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.PID = buf.DecodeUint32()
	m.Vr.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Vr.VrID = buf.DecodeUint8()
	m.Vr.IsIPv6 = buf.DecodeUint8()
	m.OldState = VrrpVrState(buf.DecodeUint32())
	m.NewState = VrrpVrState(buf.DecodeUint32())
	return nil
}

// VrrpVrPeerDetails defines message 'vrrp_vr_peer_details'.
type VrrpVrPeerDetails struct {
	SwIfIndex  interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	VrID       uint8                          `binapi:"u8,name=vr_id" json:"vr_id,omitempty"`
	IsIPv6     uint8                          `binapi:"u8,name=is_ipv6" json:"is_ipv6,omitempty"`
	NPeerAddrs uint8                          `binapi:"u8,name=n_peer_addrs" json:"-"`
	PeerAddrs  []ip_types.Address             `binapi:"address[n_peer_addrs],name=peer_addrs" json:"peer_addrs,omitempty"`
}

func (m *VrrpVrPeerDetails) Reset()               { *m = VrrpVrPeerDetails{} }
func (*VrrpVrPeerDetails) GetMessageName() string { return "vrrp_vr_peer_details" }
func (*VrrpVrPeerDetails) GetCrcString() string   { return "3d99c108" }
func (*VrrpVrPeerDetails) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VrrpVrPeerDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.VrID
	size += 1 // m.IsIPv6
	size += 1 // m.NPeerAddrs
	for j1 := 0; j1 < len(m.PeerAddrs); j1++ {
		var s1 ip_types.Address
		_ = s1
		if j1 < len(m.PeerAddrs) {
			s1 = m.PeerAddrs[j1]
		}
		size += 1      // s1.Af
		size += 1 * 16 // s1.Un
	}
	return size
}
func (m *VrrpVrPeerDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(m.VrID)
	buf.EncodeUint8(m.IsIPv6)
	buf.EncodeUint8(uint8(len(m.PeerAddrs)))
	for j0 := 0; j0 < len(m.PeerAddrs); j0++ {
		var v0 ip_types.Address // PeerAddrs
		if j0 < len(m.PeerAddrs) {
			v0 = m.PeerAddrs[j0]
		}
		buf.EncodeUint8(uint8(v0.Af))
		buf.EncodeBytes(v0.Un.XXX_UnionData[:], 16)
	}
	return buf.Bytes(), nil
}
func (m *VrrpVrPeerDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.VrID = buf.DecodeUint8()
	m.IsIPv6 = buf.DecodeUint8()
	m.NPeerAddrs = buf.DecodeUint8()
	m.PeerAddrs = make([]ip_types.Address, m.NPeerAddrs)
	for j0 := 0; j0 < len(m.PeerAddrs); j0++ {
		m.PeerAddrs[j0].Af = ip_types.AddressFamily(buf.DecodeUint8())
		copy(m.PeerAddrs[j0].Un.XXX_UnionData[:], buf.DecodeBytes(16))
	}
	return nil
}

// VrrpVrPeerDump defines message 'vrrp_vr_peer_dump'.
type VrrpVrPeerDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsIPv6    uint8                          `binapi:"u8,name=is_ipv6" json:"is_ipv6,omitempty"`
	VrID      uint8                          `binapi:"u8,name=vr_id" json:"vr_id,omitempty"`
}

func (m *VrrpVrPeerDump) Reset()               { *m = VrrpVrPeerDump{} }
func (*VrrpVrPeerDump) GetMessageName() string { return "vrrp_vr_peer_dump" }
func (*VrrpVrPeerDump) GetCrcString() string   { return "6fa3f7c4" }
func (*VrrpVrPeerDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VrrpVrPeerDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.IsIPv6
	size += 1 // m.VrID
	return size
}
func (m *VrrpVrPeerDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(m.IsIPv6)
	buf.EncodeUint8(m.VrID)
	return buf.Bytes(), nil
}
func (m *VrrpVrPeerDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsIPv6 = buf.DecodeUint8()
	m.VrID = buf.DecodeUint8()
	return nil
}

// VrrpVrSetPeers defines message 'vrrp_vr_set_peers'.
type VrrpVrSetPeers struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	VrID      uint8                          `binapi:"u8,name=vr_id" json:"vr_id,omitempty"`
	IsIPv6    uint8                          `binapi:"u8,name=is_ipv6" json:"is_ipv6,omitempty"`
	NAddrs    uint8                          `binapi:"u8,name=n_addrs" json:"-"`
	Addrs     []ip_types.Address             `binapi:"address[n_addrs],name=addrs" json:"addrs,omitempty"`
}

func (m *VrrpVrSetPeers) Reset()               { *m = VrrpVrSetPeers{} }
func (*VrrpVrSetPeers) GetMessageName() string { return "vrrp_vr_set_peers" }
func (*VrrpVrSetPeers) GetCrcString() string   { return "20bec71f" }
func (*VrrpVrSetPeers) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VrrpVrSetPeers) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.VrID
	size += 1 // m.IsIPv6
	size += 1 // m.NAddrs
	for j1 := 0; j1 < len(m.Addrs); j1++ {
		var s1 ip_types.Address
		_ = s1
		if j1 < len(m.Addrs) {
			s1 = m.Addrs[j1]
		}
		size += 1      // s1.Af
		size += 1 * 16 // s1.Un
	}
	return size
}
func (m *VrrpVrSetPeers) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(m.VrID)
	buf.EncodeUint8(m.IsIPv6)
	buf.EncodeUint8(uint8(len(m.Addrs)))
	for j0 := 0; j0 < len(m.Addrs); j0++ {
		var v0 ip_types.Address // Addrs
		if j0 < len(m.Addrs) {
			v0 = m.Addrs[j0]
		}
		buf.EncodeUint8(uint8(v0.Af))
		buf.EncodeBytes(v0.Un.XXX_UnionData[:], 16)
	}
	return buf.Bytes(), nil
}
func (m *VrrpVrSetPeers) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.VrID = buf.DecodeUint8()
	m.IsIPv6 = buf.DecodeUint8()
	m.NAddrs = buf.DecodeUint8()
	m.Addrs = make([]ip_types.Address, m.NAddrs)
	for j0 := 0; j0 < len(m.Addrs); j0++ {
		m.Addrs[j0].Af = ip_types.AddressFamily(buf.DecodeUint8())
		copy(m.Addrs[j0].Un.XXX_UnionData[:], buf.DecodeBytes(16))
	}
	return nil
}

// VrrpVrSetPeersReply defines message 'vrrp_vr_set_peers_reply'.
type VrrpVrSetPeersReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *VrrpVrSetPeersReply) Reset()               { *m = VrrpVrSetPeersReply{} }
func (*VrrpVrSetPeersReply) GetMessageName() string { return "vrrp_vr_set_peers_reply" }
func (*VrrpVrSetPeersReply) GetCrcString() string   { return "e8d4e804" }
func (*VrrpVrSetPeersReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VrrpVrSetPeersReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *VrrpVrSetPeersReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *VrrpVrSetPeersReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// VrrpVrStartStop defines message 'vrrp_vr_start_stop'.
type VrrpVrStartStop struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	VrID      uint8                          `binapi:"u8,name=vr_id" json:"vr_id,omitempty"`
	IsIPv6    uint8                          `binapi:"u8,name=is_ipv6" json:"is_ipv6,omitempty"`
	IsStart   uint8                          `binapi:"u8,name=is_start" json:"is_start,omitempty"`
}

func (m *VrrpVrStartStop) Reset()               { *m = VrrpVrStartStop{} }
func (*VrrpVrStartStop) GetMessageName() string { return "vrrp_vr_start_stop" }
func (*VrrpVrStartStop) GetCrcString() string   { return "0662a3b7" }
func (*VrrpVrStartStop) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VrrpVrStartStop) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.VrID
	size += 1 // m.IsIPv6
	size += 1 // m.IsStart
	return size
}
func (m *VrrpVrStartStop) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(m.VrID)
	buf.EncodeUint8(m.IsIPv6)
	buf.EncodeUint8(m.IsStart)
	return buf.Bytes(), nil
}
func (m *VrrpVrStartStop) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.VrID = buf.DecodeUint8()
	m.IsIPv6 = buf.DecodeUint8()
	m.IsStart = buf.DecodeUint8()
	return nil
}

// VrrpVrStartStopReply defines message 'vrrp_vr_start_stop_reply'.
type VrrpVrStartStopReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *VrrpVrStartStopReply) Reset()               { *m = VrrpVrStartStopReply{} }
func (*VrrpVrStartStopReply) GetMessageName() string { return "vrrp_vr_start_stop_reply" }
func (*VrrpVrStartStopReply) GetCrcString() string   { return "e8d4e804" }
func (*VrrpVrStartStopReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VrrpVrStartStopReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *VrrpVrStartStopReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *VrrpVrStartStopReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// VrrpVrTrackIfAddDel defines message 'vrrp_vr_track_if_add_del'.
type VrrpVrTrackIfAddDel struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsIPv6    uint8                          `binapi:"u8,name=is_ipv6" json:"is_ipv6,omitempty"`
	VrID      uint8                          `binapi:"u8,name=vr_id" json:"vr_id,omitempty"`
	IsAdd     uint8                          `binapi:"u8,name=is_add" json:"is_add,omitempty"`
	NIfs      uint8                          `binapi:"u8,name=n_ifs" json:"-"`
	Ifs       []VrrpVrTrackIf                `binapi:"vrrp_vr_track_if[n_ifs],name=ifs" json:"ifs,omitempty"`
}

func (m *VrrpVrTrackIfAddDel) Reset()               { *m = VrrpVrTrackIfAddDel{} }
func (*VrrpVrTrackIfAddDel) GetMessageName() string { return "vrrp_vr_track_if_add_del" }
func (*VrrpVrTrackIfAddDel) GetCrcString() string   { return "d67df299" }
func (*VrrpVrTrackIfAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VrrpVrTrackIfAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.IsIPv6
	size += 1 // m.VrID
	size += 1 // m.IsAdd
	size += 1 // m.NIfs
	for j1 := 0; j1 < len(m.Ifs); j1++ {
		var s1 VrrpVrTrackIf
		_ = s1
		if j1 < len(m.Ifs) {
			s1 = m.Ifs[j1]
		}
		size += 4 // s1.SwIfIndex
		size += 1 // s1.Priority
	}
	return size
}
func (m *VrrpVrTrackIfAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(m.IsIPv6)
	buf.EncodeUint8(m.VrID)
	buf.EncodeUint8(m.IsAdd)
	buf.EncodeUint8(uint8(len(m.Ifs)))
	for j0 := 0; j0 < len(m.Ifs); j0++ {
		var v0 VrrpVrTrackIf // Ifs
		if j0 < len(m.Ifs) {
			v0 = m.Ifs[j0]
		}
		buf.EncodeUint32(uint32(v0.SwIfIndex))
		buf.EncodeUint8(v0.Priority)
	}
	return buf.Bytes(), nil
}
func (m *VrrpVrTrackIfAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsIPv6 = buf.DecodeUint8()
	m.VrID = buf.DecodeUint8()
	m.IsAdd = buf.DecodeUint8()
	m.NIfs = buf.DecodeUint8()
	m.Ifs = make([]VrrpVrTrackIf, m.NIfs)
	for j0 := 0; j0 < len(m.Ifs); j0++ {
		m.Ifs[j0].SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
		m.Ifs[j0].Priority = buf.DecodeUint8()
	}
	return nil
}

// VrrpVrTrackIfAddDelReply defines message 'vrrp_vr_track_if_add_del_reply'.
type VrrpVrTrackIfAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *VrrpVrTrackIfAddDelReply) Reset()               { *m = VrrpVrTrackIfAddDelReply{} }
func (*VrrpVrTrackIfAddDelReply) GetMessageName() string { return "vrrp_vr_track_if_add_del_reply" }
func (*VrrpVrTrackIfAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*VrrpVrTrackIfAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VrrpVrTrackIfAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *VrrpVrTrackIfAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *VrrpVrTrackIfAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// VrrpVrTrackIfDetails defines message 'vrrp_vr_track_if_details'.
type VrrpVrTrackIfDetails struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	VrID      uint8                          `binapi:"u8,name=vr_id" json:"vr_id,omitempty"`
	IsIPv6    uint8                          `binapi:"u8,name=is_ipv6" json:"is_ipv6,omitempty"`
	NIfs      uint8                          `binapi:"u8,name=n_ifs" json:"-"`
	Ifs       []VrrpVrTrackIf                `binapi:"vrrp_vr_track_if[n_ifs],name=ifs" json:"ifs,omitempty"`
}

func (m *VrrpVrTrackIfDetails) Reset()               { *m = VrrpVrTrackIfDetails{} }
func (*VrrpVrTrackIfDetails) GetMessageName() string { return "vrrp_vr_track_if_details" }
func (*VrrpVrTrackIfDetails) GetCrcString() string   { return "73c36f81" }
func (*VrrpVrTrackIfDetails) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VrrpVrTrackIfDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.VrID
	size += 1 // m.IsIPv6
	size += 1 // m.NIfs
	for j1 := 0; j1 < len(m.Ifs); j1++ {
		var s1 VrrpVrTrackIf
		_ = s1
		if j1 < len(m.Ifs) {
			s1 = m.Ifs[j1]
		}
		size += 4 // s1.SwIfIndex
		size += 1 // s1.Priority
	}
	return size
}
func (m *VrrpVrTrackIfDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(m.VrID)
	buf.EncodeUint8(m.IsIPv6)
	buf.EncodeUint8(uint8(len(m.Ifs)))
	for j0 := 0; j0 < len(m.Ifs); j0++ {
		var v0 VrrpVrTrackIf // Ifs
		if j0 < len(m.Ifs) {
			v0 = m.Ifs[j0]
		}
		buf.EncodeUint32(uint32(v0.SwIfIndex))
		buf.EncodeUint8(v0.Priority)
	}
	return buf.Bytes(), nil
}
func (m *VrrpVrTrackIfDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.VrID = buf.DecodeUint8()
	m.IsIPv6 = buf.DecodeUint8()
	m.NIfs = buf.DecodeUint8()
	m.Ifs = make([]VrrpVrTrackIf, m.NIfs)
	for j0 := 0; j0 < len(m.Ifs); j0++ {
		m.Ifs[j0].SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
		m.Ifs[j0].Priority = buf.DecodeUint8()
	}
	return nil
}

// VrrpVrTrackIfDump defines message 'vrrp_vr_track_if_dump'.
type VrrpVrTrackIfDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsIPv6    uint8                          `binapi:"u8,name=is_ipv6" json:"is_ipv6,omitempty"`
	VrID      uint8                          `binapi:"u8,name=vr_id" json:"vr_id,omitempty"`
	DumpAll   uint8                          `binapi:"u8,name=dump_all" json:"dump_all,omitempty"`
}

func (m *VrrpVrTrackIfDump) Reset()               { *m = VrrpVrTrackIfDump{} }
func (*VrrpVrTrackIfDump) GetMessageName() string { return "vrrp_vr_track_if_dump" }
func (*VrrpVrTrackIfDump) GetCrcString() string   { return "a34dfc6d" }
func (*VrrpVrTrackIfDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VrrpVrTrackIfDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.IsIPv6
	size += 1 // m.VrID
	size += 1 // m.DumpAll
	return size
}
func (m *VrrpVrTrackIfDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(m.IsIPv6)
	buf.EncodeUint8(m.VrID)
	buf.EncodeUint8(m.DumpAll)
	return buf.Bytes(), nil
}
func (m *VrrpVrTrackIfDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsIPv6 = buf.DecodeUint8()
	m.VrID = buf.DecodeUint8()
	m.DumpAll = buf.DecodeUint8()
	return nil
}

// WantVrrpVrEvents defines message 'want_vrrp_vr_events'.
type WantVrrpVrEvents struct {
	EnableDisable bool   `binapi:"bool,name=enable_disable" json:"enable_disable,omitempty"`
	PID           uint32 `binapi:"u32,name=pid" json:"pid,omitempty"`
}

func (m *WantVrrpVrEvents) Reset()               { *m = WantVrrpVrEvents{} }
func (*WantVrrpVrEvents) GetMessageName() string { return "want_vrrp_vr_events" }
func (*WantVrrpVrEvents) GetCrcString() string   { return "c5e2af94" }
func (*WantVrrpVrEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *WantVrrpVrEvents) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.EnableDisable
	size += 4 // m.PID
	return size
}
func (m *WantVrrpVrEvents) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.EnableDisable)
	buf.EncodeUint32(m.PID)
	return buf.Bytes(), nil
}
func (m *WantVrrpVrEvents) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.EnableDisable = buf.DecodeBool()
	m.PID = buf.DecodeUint32()
	return nil
}

// WantVrrpVrEventsReply defines message 'want_vrrp_vr_events_reply'.
type WantVrrpVrEventsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *WantVrrpVrEventsReply) Reset()               { *m = WantVrrpVrEventsReply{} }
func (*WantVrrpVrEventsReply) GetMessageName() string { return "want_vrrp_vr_events_reply" }
func (*WantVrrpVrEventsReply) GetCrcString() string   { return "e8d4e804" }
func (*WantVrrpVrEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *WantVrrpVrEventsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *WantVrrpVrEventsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *WantVrrpVrEventsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_vrrp_binapi_init() }
func file_vrrp_binapi_init() {
	api.RegisterMessage((*VrrpVrAddDel)(nil), "vrrp_vr_add_del_c5cf15aa")
	api.RegisterMessage((*VrrpVrAddDelReply)(nil), "vrrp_vr_add_del_reply_e8d4e804")
	api.RegisterMessage((*VrrpVrDetails)(nil), "vrrp_vr_details_46edcebd")
	api.RegisterMessage((*VrrpVrDump)(nil), "vrrp_vr_dump_f9e6675e")
	api.RegisterMessage((*VrrpVrEvent)(nil), "vrrp_vr_event_c1fea6a5")
	api.RegisterMessage((*VrrpVrPeerDetails)(nil), "vrrp_vr_peer_details_3d99c108")
	api.RegisterMessage((*VrrpVrPeerDump)(nil), "vrrp_vr_peer_dump_6fa3f7c4")
	api.RegisterMessage((*VrrpVrSetPeers)(nil), "vrrp_vr_set_peers_20bec71f")
	api.RegisterMessage((*VrrpVrSetPeersReply)(nil), "vrrp_vr_set_peers_reply_e8d4e804")
	api.RegisterMessage((*VrrpVrStartStop)(nil), "vrrp_vr_start_stop_0662a3b7")
	api.RegisterMessage((*VrrpVrStartStopReply)(nil), "vrrp_vr_start_stop_reply_e8d4e804")
	api.RegisterMessage((*VrrpVrTrackIfAddDel)(nil), "vrrp_vr_track_if_add_del_d67df299")
	api.RegisterMessage((*VrrpVrTrackIfAddDelReply)(nil), "vrrp_vr_track_if_add_del_reply_e8d4e804")
	api.RegisterMessage((*VrrpVrTrackIfDetails)(nil), "vrrp_vr_track_if_details_73c36f81")
	api.RegisterMessage((*VrrpVrTrackIfDump)(nil), "vrrp_vr_track_if_dump_a34dfc6d")
	api.RegisterMessage((*WantVrrpVrEvents)(nil), "want_vrrp_vr_events_c5e2af94")
	api.RegisterMessage((*WantVrrpVrEventsReply)(nil), "want_vrrp_vr_events_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*VrrpVrAddDel)(nil),
		(*VrrpVrAddDelReply)(nil),
		(*VrrpVrDetails)(nil),
		(*VrrpVrDump)(nil),
		(*VrrpVrEvent)(nil),
		(*VrrpVrPeerDetails)(nil),
		(*VrrpVrPeerDump)(nil),
		(*VrrpVrSetPeers)(nil),
		(*VrrpVrSetPeersReply)(nil),
		(*VrrpVrStartStop)(nil),
		(*VrrpVrStartStopReply)(nil),
		(*VrrpVrTrackIfAddDel)(nil),
		(*VrrpVrTrackIfAddDelReply)(nil),
		(*VrrpVrTrackIfDetails)(nil),
		(*VrrpVrTrackIfDump)(nil),
		(*WantVrrpVrEvents)(nil),
		(*WantVrrpVrEventsReply)(nil),
	}
}
