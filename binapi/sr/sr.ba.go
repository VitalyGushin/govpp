// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              21.06.0-13~g94dced301
// source: /usr/share/vpp/api/core/sr.api.json

// Package sr contains generated bindings for API file sr.api.
//
// Contents:
//   2 structs
//  22 messages
//
package sr

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/edwarnicke/govpp/binapi/interface_types"
	ip_types "github.com/edwarnicke/govpp/binapi/ip_types"
	sr_types "github.com/edwarnicke/govpp/binapi/sr_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "sr"
	APIVersion = "2.0.0"
	VersionCrc = 0xb17a64be
)

// Srv6SidList defines type 'srv6_sid_list'.
type Srv6SidList struct {
	NumSids uint8                   `binapi:"u8,name=num_sids" json:"num_sids,omitempty"`
	Weight  uint32                  `binapi:"u32,name=weight" json:"weight,omitempty"`
	Sids    [16]ip_types.IP6Address `binapi:"ip6_address[16],name=sids" json:"sids,omitempty"`
}

// Srv6SidListWithSlIndex defines type 'srv6_sid_list_with_sl_index'.
type Srv6SidListWithSlIndex struct {
	NumSids uint8                   `binapi:"u8,name=num_sids" json:"num_sids,omitempty"`
	Weight  uint32                  `binapi:"u32,name=weight" json:"weight,omitempty"`
	SlIndex uint32                  `binapi:"u32,name=sl_index" json:"sl_index,omitempty"`
	Sids    [16]ip_types.IP6Address `binapi:"ip6_address[16],name=sids" json:"sids,omitempty"`
}

// SrLocalsidAddDel defines message 'sr_localsid_add_del'.
type SrLocalsidAddDel struct {
	IsDel     bool                           `binapi:"bool,name=is_del,default=false" json:"is_del,omitempty"`
	Localsid  ip_types.IP6Address            `binapi:"ip6_address,name=localsid" json:"localsid,omitempty"`
	EndPsp    bool                           `binapi:"bool,name=end_psp" json:"end_psp,omitempty"`
	Behavior  sr_types.SrBehavior            `binapi:"sr_behavior,name=behavior" json:"behavior,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index,default=4294967295" json:"sw_if_index,omitempty"`
	VlanIndex uint32                         `binapi:"u32,name=vlan_index" json:"vlan_index,omitempty"`
	FibTable  uint32                         `binapi:"u32,name=fib_table" json:"fib_table,omitempty"`
	NhAddr    ip_types.Address               `binapi:"address,name=nh_addr" json:"nh_addr,omitempty"`
}

func (m *SrLocalsidAddDel) Reset()               { *m = SrLocalsidAddDel{} }
func (*SrLocalsidAddDel) GetMessageName() string { return "sr_localsid_add_del" }
func (*SrLocalsidAddDel) GetCrcString() string   { return "5a36c324" }
func (*SrLocalsidAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrLocalsidAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsDel
	size += 1 * 16 // m.Localsid
	size += 1      // m.EndPsp
	size += 1      // m.Behavior
	size += 4      // m.SwIfIndex
	size += 4      // m.VlanIndex
	size += 4      // m.FibTable
	size += 1      // m.NhAddr.Af
	size += 1 * 16 // m.NhAddr.Un
	return size
}
func (m *SrLocalsidAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsDel)
	buf.EncodeBytes(m.Localsid[:], 16)
	buf.EncodeBool(m.EndPsp)
	buf.EncodeUint8(uint8(m.Behavior))
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint32(m.VlanIndex)
	buf.EncodeUint32(m.FibTable)
	buf.EncodeUint8(uint8(m.NhAddr.Af))
	buf.EncodeBytes(m.NhAddr.Un.XXX_UnionData[:], 16)
	return buf.Bytes(), nil
}
func (m *SrLocalsidAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsDel = buf.DecodeBool()
	copy(m.Localsid[:], buf.DecodeBytes(16))
	m.EndPsp = buf.DecodeBool()
	m.Behavior = sr_types.SrBehavior(buf.DecodeUint8())
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.VlanIndex = buf.DecodeUint32()
	m.FibTable = buf.DecodeUint32()
	m.NhAddr.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.NhAddr.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	return nil
}

// SrLocalsidAddDelReply defines message 'sr_localsid_add_del_reply'.
type SrLocalsidAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrLocalsidAddDelReply) Reset()               { *m = SrLocalsidAddDelReply{} }
func (*SrLocalsidAddDelReply) GetMessageName() string { return "sr_localsid_add_del_reply" }
func (*SrLocalsidAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*SrLocalsidAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrLocalsidAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrLocalsidAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrLocalsidAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrLocalsidsDetails defines message 'sr_localsids_details'.
type SrLocalsidsDetails struct {
	Addr                    ip_types.IP6Address `binapi:"ip6_address,name=addr" json:"addr,omitempty"`
	EndPsp                  bool                `binapi:"bool,name=end_psp" json:"end_psp,omitempty"`
	Behavior                sr_types.SrBehavior `binapi:"sr_behavior,name=behavior" json:"behavior,omitempty"`
	FibTable                uint32              `binapi:"u32,name=fib_table" json:"fib_table,omitempty"`
	VlanIndex               uint32              `binapi:"u32,name=vlan_index" json:"vlan_index,omitempty"`
	XconnectNhAddr          ip_types.Address    `binapi:"address,name=xconnect_nh_addr" json:"xconnect_nh_addr,omitempty"`
	XconnectIfaceOrVrfTable uint32              `binapi:"u32,name=xconnect_iface_or_vrf_table" json:"xconnect_iface_or_vrf_table,omitempty"`
}

func (m *SrLocalsidsDetails) Reset()               { *m = SrLocalsidsDetails{} }
func (*SrLocalsidsDetails) GetMessageName() string { return "sr_localsids_details" }
func (*SrLocalsidsDetails) GetCrcString() string   { return "2e9221b9" }
func (*SrLocalsidsDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrLocalsidsDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 16 // m.Addr
	size += 1      // m.EndPsp
	size += 1      // m.Behavior
	size += 4      // m.FibTable
	size += 4      // m.VlanIndex
	size += 1      // m.XconnectNhAddr.Af
	size += 1 * 16 // m.XconnectNhAddr.Un
	size += 4      // m.XconnectIfaceOrVrfTable
	return size
}
func (m *SrLocalsidsDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.Addr[:], 16)
	buf.EncodeBool(m.EndPsp)
	buf.EncodeUint8(uint8(m.Behavior))
	buf.EncodeUint32(m.FibTable)
	buf.EncodeUint32(m.VlanIndex)
	buf.EncodeUint8(uint8(m.XconnectNhAddr.Af))
	buf.EncodeBytes(m.XconnectNhAddr.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.XconnectIfaceOrVrfTable)
	return buf.Bytes(), nil
}
func (m *SrLocalsidsDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.Addr[:], buf.DecodeBytes(16))
	m.EndPsp = buf.DecodeBool()
	m.Behavior = sr_types.SrBehavior(buf.DecodeUint8())
	m.FibTable = buf.DecodeUint32()
	m.VlanIndex = buf.DecodeUint32()
	m.XconnectNhAddr.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.XconnectNhAddr.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.XconnectIfaceOrVrfTable = buf.DecodeUint32()
	return nil
}

// SrLocalsidsDump defines message 'sr_localsids_dump'.
type SrLocalsidsDump struct{}

func (m *SrLocalsidsDump) Reset()               { *m = SrLocalsidsDump{} }
func (*SrLocalsidsDump) GetMessageName() string { return "sr_localsids_dump" }
func (*SrLocalsidsDump) GetCrcString() string   { return "51077d14" }
func (*SrLocalsidsDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrLocalsidsDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *SrLocalsidsDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *SrLocalsidsDump) Unmarshal(b []byte) error {
	return nil
}

// SrPoliciesDetails defines message 'sr_policies_details'.
type SrPoliciesDetails struct {
	Bsid        ip_types.IP6Address `binapi:"ip6_address,name=bsid" json:"bsid,omitempty"`
	IsSpray     bool                `binapi:"bool,name=is_spray" json:"is_spray,omitempty"`
	IsEncap     bool                `binapi:"bool,name=is_encap" json:"is_encap,omitempty"`
	FibTable    uint32              `binapi:"u32,name=fib_table" json:"fib_table,omitempty"`
	NumSidLists uint8               `binapi:"u8,name=num_sid_lists" json:"-"`
	SidLists    []Srv6SidList       `binapi:"srv6_sid_list[num_sid_lists],name=sid_lists" json:"sid_lists,omitempty"`
}

func (m *SrPoliciesDetails) Reset()               { *m = SrPoliciesDetails{} }
func (*SrPoliciesDetails) GetMessageName() string { return "sr_policies_details" }
func (*SrPoliciesDetails) GetCrcString() string   { return "db6ff2a1" }
func (*SrPoliciesDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrPoliciesDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 16 // m.Bsid
	size += 1      // m.IsSpray
	size += 1      // m.IsEncap
	size += 4      // m.FibTable
	size += 1      // m.NumSidLists
	for j1 := 0; j1 < len(m.SidLists); j1++ {
		var s1 Srv6SidList
		_ = s1
		if j1 < len(m.SidLists) {
			s1 = m.SidLists[j1]
		}
		size += 1 // s1.NumSids
		size += 4 // s1.Weight
		for j2 := 0; j2 < 16; j2++ {
			size += 1 * 16 // s1.Sids[j2]
		}
	}
	return size
}
func (m *SrPoliciesDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.Bsid[:], 16)
	buf.EncodeBool(m.IsSpray)
	buf.EncodeBool(m.IsEncap)
	buf.EncodeUint32(m.FibTable)
	buf.EncodeUint8(uint8(len(m.SidLists)))
	for j0 := 0; j0 < len(m.SidLists); j0++ {
		var v0 Srv6SidList // SidLists
		if j0 < len(m.SidLists) {
			v0 = m.SidLists[j0]
		}
		buf.EncodeUint8(v0.NumSids)
		buf.EncodeUint32(v0.Weight)
		for j1 := 0; j1 < 16; j1++ {
			buf.EncodeBytes(v0.Sids[j1][:], 16)
		}
	}
	return buf.Bytes(), nil
}
func (m *SrPoliciesDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.Bsid[:], buf.DecodeBytes(16))
	m.IsSpray = buf.DecodeBool()
	m.IsEncap = buf.DecodeBool()
	m.FibTable = buf.DecodeUint32()
	m.NumSidLists = buf.DecodeUint8()
	m.SidLists = make([]Srv6SidList, m.NumSidLists)
	for j0 := 0; j0 < len(m.SidLists); j0++ {
		m.SidLists[j0].NumSids = buf.DecodeUint8()
		m.SidLists[j0].Weight = buf.DecodeUint32()
		for j1 := 0; j1 < 16; j1++ {
			copy(m.SidLists[j0].Sids[j1][:], buf.DecodeBytes(16))
		}
	}
	return nil
}

// SrPoliciesDump defines message 'sr_policies_dump'.
type SrPoliciesDump struct{}

func (m *SrPoliciesDump) Reset()               { *m = SrPoliciesDump{} }
func (*SrPoliciesDump) GetMessageName() string { return "sr_policies_dump" }
func (*SrPoliciesDump) GetCrcString() string   { return "51077d14" }
func (*SrPoliciesDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrPoliciesDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *SrPoliciesDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *SrPoliciesDump) Unmarshal(b []byte) error {
	return nil
}

// SrPoliciesWithSlIndexDetails defines message 'sr_policies_with_sl_index_details'.
// InProgress: the message form may change in the future versions
type SrPoliciesWithSlIndexDetails struct {
	Bsid        ip_types.IP6Address      `binapi:"ip6_address,name=bsid" json:"bsid,omitempty"`
	IsSpray     bool                     `binapi:"bool,name=is_spray" json:"is_spray,omitempty"`
	IsEncap     bool                     `binapi:"bool,name=is_encap" json:"is_encap,omitempty"`
	FibTable    uint32                   `binapi:"u32,name=fib_table" json:"fib_table,omitempty"`
	NumSidLists uint8                    `binapi:"u8,name=num_sid_lists" json:"-"`
	SidLists    []Srv6SidListWithSlIndex `binapi:"srv6_sid_list_with_sl_index[num_sid_lists],name=sid_lists" json:"sid_lists,omitempty"`
}

func (m *SrPoliciesWithSlIndexDetails) Reset() { *m = SrPoliciesWithSlIndexDetails{} }
func (*SrPoliciesWithSlIndexDetails) GetMessageName() string {
	return "sr_policies_with_sl_index_details"
}
func (*SrPoliciesWithSlIndexDetails) GetCrcString() string { return "ca2e9bc8" }
func (*SrPoliciesWithSlIndexDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrPoliciesWithSlIndexDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 16 // m.Bsid
	size += 1      // m.IsSpray
	size += 1      // m.IsEncap
	size += 4      // m.FibTable
	size += 1      // m.NumSidLists
	for j1 := 0; j1 < len(m.SidLists); j1++ {
		var s1 Srv6SidListWithSlIndex
		_ = s1
		if j1 < len(m.SidLists) {
			s1 = m.SidLists[j1]
		}
		size += 1 // s1.NumSids
		size += 4 // s1.Weight
		size += 4 // s1.SlIndex
		for j2 := 0; j2 < 16; j2++ {
			size += 1 * 16 // s1.Sids[j2]
		}
	}
	return size
}
func (m *SrPoliciesWithSlIndexDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.Bsid[:], 16)
	buf.EncodeBool(m.IsSpray)
	buf.EncodeBool(m.IsEncap)
	buf.EncodeUint32(m.FibTable)
	buf.EncodeUint8(uint8(len(m.SidLists)))
	for j0 := 0; j0 < len(m.SidLists); j0++ {
		var v0 Srv6SidListWithSlIndex // SidLists
		if j0 < len(m.SidLists) {
			v0 = m.SidLists[j0]
		}
		buf.EncodeUint8(v0.NumSids)
		buf.EncodeUint32(v0.Weight)
		buf.EncodeUint32(v0.SlIndex)
		for j1 := 0; j1 < 16; j1++ {
			buf.EncodeBytes(v0.Sids[j1][:], 16)
		}
	}
	return buf.Bytes(), nil
}
func (m *SrPoliciesWithSlIndexDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.Bsid[:], buf.DecodeBytes(16))
	m.IsSpray = buf.DecodeBool()
	m.IsEncap = buf.DecodeBool()
	m.FibTable = buf.DecodeUint32()
	m.NumSidLists = buf.DecodeUint8()
	m.SidLists = make([]Srv6SidListWithSlIndex, m.NumSidLists)
	for j0 := 0; j0 < len(m.SidLists); j0++ {
		m.SidLists[j0].NumSids = buf.DecodeUint8()
		m.SidLists[j0].Weight = buf.DecodeUint32()
		m.SidLists[j0].SlIndex = buf.DecodeUint32()
		for j1 := 0; j1 < 16; j1++ {
			copy(m.SidLists[j0].Sids[j1][:], buf.DecodeBytes(16))
		}
	}
	return nil
}

// SrPoliciesWithSlIndexDump defines message 'sr_policies_with_sl_index_dump'.
// InProgress: the message form may change in the future versions
type SrPoliciesWithSlIndexDump struct{}

func (m *SrPoliciesWithSlIndexDump) Reset()               { *m = SrPoliciesWithSlIndexDump{} }
func (*SrPoliciesWithSlIndexDump) GetMessageName() string { return "sr_policies_with_sl_index_dump" }
func (*SrPoliciesWithSlIndexDump) GetCrcString() string   { return "51077d14" }
func (*SrPoliciesWithSlIndexDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrPoliciesWithSlIndexDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *SrPoliciesWithSlIndexDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *SrPoliciesWithSlIndexDump) Unmarshal(b []byte) error {
	return nil
}

// SrPolicyAdd defines message 'sr_policy_add'.
type SrPolicyAdd struct {
	BsidAddr ip_types.IP6Address `binapi:"ip6_address,name=bsid_addr" json:"bsid_addr,omitempty"`
	Weight   uint32              `binapi:"u32,name=weight" json:"weight,omitempty"`
	IsEncap  bool                `binapi:"bool,name=is_encap" json:"is_encap,omitempty"`
	IsSpray  bool                `binapi:"bool,name=is_spray" json:"is_spray,omitempty"`
	FibTable uint32              `binapi:"u32,name=fib_table" json:"fib_table,omitempty"`
	Sids     Srv6SidList         `binapi:"srv6_sid_list,name=sids" json:"sids,omitempty"`
}

func (m *SrPolicyAdd) Reset()               { *m = SrPolicyAdd{} }
func (*SrPolicyAdd) GetMessageName() string { return "sr_policy_add" }
func (*SrPolicyAdd) GetCrcString() string   { return "44ac92e8" }
func (*SrPolicyAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrPolicyAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 16 // m.BsidAddr
	size += 4      // m.Weight
	size += 1      // m.IsEncap
	size += 1      // m.IsSpray
	size += 4      // m.FibTable
	size += 1      // m.Sids.NumSids
	size += 4      // m.Sids.Weight
	for j2 := 0; j2 < 16; j2++ {
		size += 1 * 16 // m.Sids.Sids[j2]
	}
	return size
}
func (m *SrPolicyAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.BsidAddr[:], 16)
	buf.EncodeUint32(m.Weight)
	buf.EncodeBool(m.IsEncap)
	buf.EncodeBool(m.IsSpray)
	buf.EncodeUint32(m.FibTable)
	buf.EncodeUint8(m.Sids.NumSids)
	buf.EncodeUint32(m.Sids.Weight)
	for j1 := 0; j1 < 16; j1++ {
		buf.EncodeBytes(m.Sids.Sids[j1][:], 16)
	}
	return buf.Bytes(), nil
}
func (m *SrPolicyAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.BsidAddr[:], buf.DecodeBytes(16))
	m.Weight = buf.DecodeUint32()
	m.IsEncap = buf.DecodeBool()
	m.IsSpray = buf.DecodeBool()
	m.FibTable = buf.DecodeUint32()
	m.Sids.NumSids = buf.DecodeUint8()
	m.Sids.Weight = buf.DecodeUint32()
	for j1 := 0; j1 < 16; j1++ {
		copy(m.Sids.Sids[j1][:], buf.DecodeBytes(16))
	}
	return nil
}

// SrPolicyAddReply defines message 'sr_policy_add_reply'.
type SrPolicyAddReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrPolicyAddReply) Reset()               { *m = SrPolicyAddReply{} }
func (*SrPolicyAddReply) GetMessageName() string { return "sr_policy_add_reply" }
func (*SrPolicyAddReply) GetCrcString() string   { return "e8d4e804" }
func (*SrPolicyAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrPolicyAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrPolicyAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrPolicyAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrPolicyDel defines message 'sr_policy_del'.
type SrPolicyDel struct {
	BsidAddr      ip_types.IP6Address `binapi:"ip6_address,name=bsid_addr" json:"bsid_addr,omitempty"`
	SrPolicyIndex uint32              `binapi:"u32,name=sr_policy_index" json:"sr_policy_index,omitempty"`
}

func (m *SrPolicyDel) Reset()               { *m = SrPolicyDel{} }
func (*SrPolicyDel) GetMessageName() string { return "sr_policy_del" }
func (*SrPolicyDel) GetCrcString() string   { return "cb4d48d5" }
func (*SrPolicyDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrPolicyDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 16 // m.BsidAddr
	size += 4      // m.SrPolicyIndex
	return size
}
func (m *SrPolicyDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.BsidAddr[:], 16)
	buf.EncodeUint32(m.SrPolicyIndex)
	return buf.Bytes(), nil
}
func (m *SrPolicyDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.BsidAddr[:], buf.DecodeBytes(16))
	m.SrPolicyIndex = buf.DecodeUint32()
	return nil
}

// SrPolicyDelReply defines message 'sr_policy_del_reply'.
type SrPolicyDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrPolicyDelReply) Reset()               { *m = SrPolicyDelReply{} }
func (*SrPolicyDelReply) GetMessageName() string { return "sr_policy_del_reply" }
func (*SrPolicyDelReply) GetCrcString() string   { return "e8d4e804" }
func (*SrPolicyDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrPolicyDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrPolicyDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrPolicyDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrPolicyMod defines message 'sr_policy_mod'.
type SrPolicyMod struct {
	BsidAddr      ip_types.IP6Address `binapi:"ip6_address,name=bsid_addr" json:"bsid_addr,omitempty"`
	SrPolicyIndex uint32              `binapi:"u32,name=sr_policy_index" json:"sr_policy_index,omitempty"`
	FibTable      uint32              `binapi:"u32,name=fib_table" json:"fib_table,omitempty"`
	Operation     sr_types.SrPolicyOp `binapi:"sr_policy_op,name=operation" json:"operation,omitempty"`
	SlIndex       uint32              `binapi:"u32,name=sl_index" json:"sl_index,omitempty"`
	Weight        uint32              `binapi:"u32,name=weight" json:"weight,omitempty"`
	Sids          Srv6SidList         `binapi:"srv6_sid_list,name=sids" json:"sids,omitempty"`
}

func (m *SrPolicyMod) Reset()               { *m = SrPolicyMod{} }
func (*SrPolicyMod) GetMessageName() string { return "sr_policy_mod" }
func (*SrPolicyMod) GetCrcString() string   { return "b97bb56e" }
func (*SrPolicyMod) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrPolicyMod) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 16 // m.BsidAddr
	size += 4      // m.SrPolicyIndex
	size += 4      // m.FibTable
	size += 1      // m.Operation
	size += 4      // m.SlIndex
	size += 4      // m.Weight
	size += 1      // m.Sids.NumSids
	size += 4      // m.Sids.Weight
	for j2 := 0; j2 < 16; j2++ {
		size += 1 * 16 // m.Sids.Sids[j2]
	}
	return size
}
func (m *SrPolicyMod) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.BsidAddr[:], 16)
	buf.EncodeUint32(m.SrPolicyIndex)
	buf.EncodeUint32(m.FibTable)
	buf.EncodeUint8(uint8(m.Operation))
	buf.EncodeUint32(m.SlIndex)
	buf.EncodeUint32(m.Weight)
	buf.EncodeUint8(m.Sids.NumSids)
	buf.EncodeUint32(m.Sids.Weight)
	for j1 := 0; j1 < 16; j1++ {
		buf.EncodeBytes(m.Sids.Sids[j1][:], 16)
	}
	return buf.Bytes(), nil
}
func (m *SrPolicyMod) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.BsidAddr[:], buf.DecodeBytes(16))
	m.SrPolicyIndex = buf.DecodeUint32()
	m.FibTable = buf.DecodeUint32()
	m.Operation = sr_types.SrPolicyOp(buf.DecodeUint8())
	m.SlIndex = buf.DecodeUint32()
	m.Weight = buf.DecodeUint32()
	m.Sids.NumSids = buf.DecodeUint8()
	m.Sids.Weight = buf.DecodeUint32()
	for j1 := 0; j1 < 16; j1++ {
		copy(m.Sids.Sids[j1][:], buf.DecodeBytes(16))
	}
	return nil
}

// SrPolicyModReply defines message 'sr_policy_mod_reply'.
type SrPolicyModReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrPolicyModReply) Reset()               { *m = SrPolicyModReply{} }
func (*SrPolicyModReply) GetMessageName() string { return "sr_policy_mod_reply" }
func (*SrPolicyModReply) GetCrcString() string   { return "e8d4e804" }
func (*SrPolicyModReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrPolicyModReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrPolicyModReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrPolicyModReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrSetEncapHopLimit defines message 'sr_set_encap_hop_limit'.
type SrSetEncapHopLimit struct {
	HopLimit uint8 `binapi:"u8,name=hop_limit" json:"hop_limit,omitempty"`
}

func (m *SrSetEncapHopLimit) Reset()               { *m = SrSetEncapHopLimit{} }
func (*SrSetEncapHopLimit) GetMessageName() string { return "sr_set_encap_hop_limit" }
func (*SrSetEncapHopLimit) GetCrcString() string   { return "aa75d7d0" }
func (*SrSetEncapHopLimit) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrSetEncapHopLimit) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.HopLimit
	return size
}
func (m *SrSetEncapHopLimit) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.HopLimit)
	return buf.Bytes(), nil
}
func (m *SrSetEncapHopLimit) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.HopLimit = buf.DecodeUint8()
	return nil
}

// SrSetEncapHopLimitReply defines message 'sr_set_encap_hop_limit_reply'.
type SrSetEncapHopLimitReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrSetEncapHopLimitReply) Reset()               { *m = SrSetEncapHopLimitReply{} }
func (*SrSetEncapHopLimitReply) GetMessageName() string { return "sr_set_encap_hop_limit_reply" }
func (*SrSetEncapHopLimitReply) GetCrcString() string   { return "e8d4e804" }
func (*SrSetEncapHopLimitReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrSetEncapHopLimitReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrSetEncapHopLimitReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrSetEncapHopLimitReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrSetEncapSource defines message 'sr_set_encap_source'.
type SrSetEncapSource struct {
	EncapsSource ip_types.IP6Address `binapi:"ip6_address,name=encaps_source" json:"encaps_source,omitempty"`
}

func (m *SrSetEncapSource) Reset()               { *m = SrSetEncapSource{} }
func (*SrSetEncapSource) GetMessageName() string { return "sr_set_encap_source" }
func (*SrSetEncapSource) GetCrcString() string   { return "d3bad5e1" }
func (*SrSetEncapSource) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrSetEncapSource) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 16 // m.EncapsSource
	return size
}
func (m *SrSetEncapSource) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.EncapsSource[:], 16)
	return buf.Bytes(), nil
}
func (m *SrSetEncapSource) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	copy(m.EncapsSource[:], buf.DecodeBytes(16))
	return nil
}

// SrSetEncapSourceReply defines message 'sr_set_encap_source_reply'.
type SrSetEncapSourceReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrSetEncapSourceReply) Reset()               { *m = SrSetEncapSourceReply{} }
func (*SrSetEncapSourceReply) GetMessageName() string { return "sr_set_encap_source_reply" }
func (*SrSetEncapSourceReply) GetCrcString() string   { return "e8d4e804" }
func (*SrSetEncapSourceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrSetEncapSourceReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrSetEncapSourceReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrSetEncapSourceReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrSteeringAddDel defines message 'sr_steering_add_del'.
type SrSteeringAddDel struct {
	IsDel         bool                           `binapi:"bool,name=is_del,default=false" json:"is_del,omitempty"`
	BsidAddr      ip_types.IP6Address            `binapi:"ip6_address,name=bsid_addr" json:"bsid_addr,omitempty"`
	SrPolicyIndex uint32                         `binapi:"u32,name=sr_policy_index" json:"sr_policy_index,omitempty"`
	TableID       uint32                         `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	Prefix        ip_types.Prefix                `binapi:"prefix,name=prefix" json:"prefix,omitempty"`
	SwIfIndex     interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	TrafficType   sr_types.SrSteer               `binapi:"sr_steer,name=traffic_type" json:"traffic_type,omitempty"`
}

func (m *SrSteeringAddDel) Reset()               { *m = SrSteeringAddDel{} }
func (*SrSteeringAddDel) GetMessageName() string { return "sr_steering_add_del" }
func (*SrSteeringAddDel) GetCrcString() string   { return "e46b0a0f" }
func (*SrSteeringAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrSteeringAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsDel
	size += 1 * 16 // m.BsidAddr
	size += 4      // m.SrPolicyIndex
	size += 4      // m.TableID
	size += 1      // m.Prefix.Address.Af
	size += 1 * 16 // m.Prefix.Address.Un
	size += 1      // m.Prefix.Len
	size += 4      // m.SwIfIndex
	size += 1      // m.TrafficType
	return size
}
func (m *SrSteeringAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsDel)
	buf.EncodeBytes(m.BsidAddr[:], 16)
	buf.EncodeUint32(m.SrPolicyIndex)
	buf.EncodeUint32(m.TableID)
	buf.EncodeUint8(uint8(m.Prefix.Address.Af))
	buf.EncodeBytes(m.Prefix.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Prefix.Len)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(uint8(m.TrafficType))
	return buf.Bytes(), nil
}
func (m *SrSteeringAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsDel = buf.DecodeBool()
	copy(m.BsidAddr[:], buf.DecodeBytes(16))
	m.SrPolicyIndex = buf.DecodeUint32()
	m.TableID = buf.DecodeUint32()
	m.Prefix.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Prefix.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Prefix.Len = buf.DecodeUint8()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.TrafficType = sr_types.SrSteer(buf.DecodeUint8())
	return nil
}

// SrSteeringAddDelReply defines message 'sr_steering_add_del_reply'.
type SrSteeringAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrSteeringAddDelReply) Reset()               { *m = SrSteeringAddDelReply{} }
func (*SrSteeringAddDelReply) GetMessageName() string { return "sr_steering_add_del_reply" }
func (*SrSteeringAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*SrSteeringAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrSteeringAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrSteeringAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrSteeringAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrSteeringPolDetails defines message 'sr_steering_pol_details'.
type SrSteeringPolDetails struct {
	TrafficType sr_types.SrSteer               `binapi:"sr_steer,name=traffic_type" json:"traffic_type,omitempty"`
	FibTable    uint32                         `binapi:"u32,name=fib_table" json:"fib_table,omitempty"`
	Prefix      ip_types.Prefix                `binapi:"prefix,name=prefix" json:"prefix,omitempty"`
	SwIfIndex   interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Bsid        ip_types.IP6Address            `binapi:"ip6_address,name=bsid" json:"bsid,omitempty"`
}

func (m *SrSteeringPolDetails) Reset()               { *m = SrSteeringPolDetails{} }
func (*SrSteeringPolDetails) GetMessageName() string { return "sr_steering_pol_details" }
func (*SrSteeringPolDetails) GetCrcString() string   { return "d41258c9" }
func (*SrSteeringPolDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrSteeringPolDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.TrafficType
	size += 4      // m.FibTable
	size += 1      // m.Prefix.Address.Af
	size += 1 * 16 // m.Prefix.Address.Un
	size += 1      // m.Prefix.Len
	size += 4      // m.SwIfIndex
	size += 1 * 16 // m.Bsid
	return size
}
func (m *SrSteeringPolDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.TrafficType))
	buf.EncodeUint32(m.FibTable)
	buf.EncodeUint8(uint8(m.Prefix.Address.Af))
	buf.EncodeBytes(m.Prefix.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Prefix.Len)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBytes(m.Bsid[:], 16)
	return buf.Bytes(), nil
}
func (m *SrSteeringPolDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.TrafficType = sr_types.SrSteer(buf.DecodeUint8())
	m.FibTable = buf.DecodeUint32()
	m.Prefix.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Prefix.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Prefix.Len = buf.DecodeUint8()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	copy(m.Bsid[:], buf.DecodeBytes(16))
	return nil
}

// SrSteeringPolDump defines message 'sr_steering_pol_dump'.
type SrSteeringPolDump struct{}

func (m *SrSteeringPolDump) Reset()               { *m = SrSteeringPolDump{} }
func (*SrSteeringPolDump) GetMessageName() string { return "sr_steering_pol_dump" }
func (*SrSteeringPolDump) GetCrcString() string   { return "51077d14" }
func (*SrSteeringPolDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrSteeringPolDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *SrSteeringPolDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *SrSteeringPolDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_sr_binapi_init() }
func file_sr_binapi_init() {
	api.RegisterMessage((*SrLocalsidAddDel)(nil), "sr_localsid_add_del_5a36c324")
	api.RegisterMessage((*SrLocalsidAddDelReply)(nil), "sr_localsid_add_del_reply_e8d4e804")
	api.RegisterMessage((*SrLocalsidsDetails)(nil), "sr_localsids_details_2e9221b9")
	api.RegisterMessage((*SrLocalsidsDump)(nil), "sr_localsids_dump_51077d14")
	api.RegisterMessage((*SrPoliciesDetails)(nil), "sr_policies_details_db6ff2a1")
	api.RegisterMessage((*SrPoliciesDump)(nil), "sr_policies_dump_51077d14")
	api.RegisterMessage((*SrPoliciesWithSlIndexDetails)(nil), "sr_policies_with_sl_index_details_ca2e9bc8")
	api.RegisterMessage((*SrPoliciesWithSlIndexDump)(nil), "sr_policies_with_sl_index_dump_51077d14")
	api.RegisterMessage((*SrPolicyAdd)(nil), "sr_policy_add_44ac92e8")
	api.RegisterMessage((*SrPolicyAddReply)(nil), "sr_policy_add_reply_e8d4e804")
	api.RegisterMessage((*SrPolicyDel)(nil), "sr_policy_del_cb4d48d5")
	api.RegisterMessage((*SrPolicyDelReply)(nil), "sr_policy_del_reply_e8d4e804")
	api.RegisterMessage((*SrPolicyMod)(nil), "sr_policy_mod_b97bb56e")
	api.RegisterMessage((*SrPolicyModReply)(nil), "sr_policy_mod_reply_e8d4e804")
	api.RegisterMessage((*SrSetEncapHopLimit)(nil), "sr_set_encap_hop_limit_aa75d7d0")
	api.RegisterMessage((*SrSetEncapHopLimitReply)(nil), "sr_set_encap_hop_limit_reply_e8d4e804")
	api.RegisterMessage((*SrSetEncapSource)(nil), "sr_set_encap_source_d3bad5e1")
	api.RegisterMessage((*SrSetEncapSourceReply)(nil), "sr_set_encap_source_reply_e8d4e804")
	api.RegisterMessage((*SrSteeringAddDel)(nil), "sr_steering_add_del_e46b0a0f")
	api.RegisterMessage((*SrSteeringAddDelReply)(nil), "sr_steering_add_del_reply_e8d4e804")
	api.RegisterMessage((*SrSteeringPolDetails)(nil), "sr_steering_pol_details_d41258c9")
	api.RegisterMessage((*SrSteeringPolDump)(nil), "sr_steering_pol_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SrLocalsidAddDel)(nil),
		(*SrLocalsidAddDelReply)(nil),
		(*SrLocalsidsDetails)(nil),
		(*SrLocalsidsDump)(nil),
		(*SrPoliciesDetails)(nil),
		(*SrPoliciesDump)(nil),
		(*SrPoliciesWithSlIndexDetails)(nil),
		(*SrPoliciesWithSlIndexDump)(nil),
		(*SrPolicyAdd)(nil),
		(*SrPolicyAddReply)(nil),
		(*SrPolicyDel)(nil),
		(*SrPolicyDelReply)(nil),
		(*SrPolicyMod)(nil),
		(*SrPolicyModReply)(nil),
		(*SrSetEncapHopLimit)(nil),
		(*SrSetEncapHopLimitReply)(nil),
		(*SrSetEncapSource)(nil),
		(*SrSetEncapSourceReply)(nil),
		(*SrSteeringAddDel)(nil),
		(*SrSteeringAddDelReply)(nil),
		(*SrSteeringPolDetails)(nil),
		(*SrSteeringPolDump)(nil),
	}
}
