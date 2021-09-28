// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              21.06.0-13~g94dced301
// source: /usr/share/vpp/api/plugins/tracedump.api.json

// Package tracedump contains generated bindings for API file tracedump.api.
//
// Contents:
//   1 enum
//   9 messages
//
package tracedump

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "tracedump"
	APIVersion = "0.1.0"
	VersionCrc = 0x691543d5
)

// TraceFilterFlag defines enum 'trace_filter_flag'.
type TraceFilterFlag uint32

const (
	TRACE_FF_NONE               TraceFilterFlag = 0
	TRACE_FF_INCLUDE_NODE       TraceFilterFlag = 1
	TRACE_FF_EXCLUDE_NODE       TraceFilterFlag = 2
	TRACE_FF_INCLUDE_CLASSIFIER TraceFilterFlag = 3
	TRACE_FF_EXCLUDE_CLASSIFIER TraceFilterFlag = 4
)

var (
	TraceFilterFlag_name = map[uint32]string{
		0: "TRACE_FF_NONE",
		1: "TRACE_FF_INCLUDE_NODE",
		2: "TRACE_FF_EXCLUDE_NODE",
		3: "TRACE_FF_INCLUDE_CLASSIFIER",
		4: "TRACE_FF_EXCLUDE_CLASSIFIER",
	}
	TraceFilterFlag_value = map[string]uint32{
		"TRACE_FF_NONE":               0,
		"TRACE_FF_INCLUDE_NODE":       1,
		"TRACE_FF_EXCLUDE_NODE":       2,
		"TRACE_FF_INCLUDE_CLASSIFIER": 3,
		"TRACE_FF_EXCLUDE_CLASSIFIER": 4,
	}
)

func (x TraceFilterFlag) String() string {
	s, ok := TraceFilterFlag_name[uint32(x)]
	if ok {
		return s
	}
	return "TraceFilterFlag(" + strconv.Itoa(int(x)) + ")"
}

// TraceCapturePackets defines message 'trace_capture_packets'.
// InProgress: the message form may change in the future versions
type TraceCapturePackets struct {
	NodeIndex       uint32 `binapi:"u32,name=node_index" json:"node_index,omitempty"`
	MaxPackets      uint32 `binapi:"u32,name=max_packets" json:"max_packets,omitempty"`
	UseFilter       bool   `binapi:"bool,name=use_filter" json:"use_filter,omitempty"`
	Verbose         bool   `binapi:"bool,name=verbose" json:"verbose,omitempty"`
	PreCaptureClear bool   `binapi:"bool,name=pre_capture_clear" json:"pre_capture_clear,omitempty"`
}

func (m *TraceCapturePackets) Reset()               { *m = TraceCapturePackets{} }
func (*TraceCapturePackets) GetMessageName() string { return "trace_capture_packets" }
func (*TraceCapturePackets) GetCrcString() string   { return "9e791a9b" }
func (*TraceCapturePackets) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceCapturePackets) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.NodeIndex
	size += 4 // m.MaxPackets
	size += 1 // m.UseFilter
	size += 1 // m.Verbose
	size += 1 // m.PreCaptureClear
	return size
}
func (m *TraceCapturePackets) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.NodeIndex)
	buf.EncodeUint32(m.MaxPackets)
	buf.EncodeBool(m.UseFilter)
	buf.EncodeBool(m.Verbose)
	buf.EncodeBool(m.PreCaptureClear)
	return buf.Bytes(), nil
}
func (m *TraceCapturePackets) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.NodeIndex = buf.DecodeUint32()
	m.MaxPackets = buf.DecodeUint32()
	m.UseFilter = buf.DecodeBool()
	m.Verbose = buf.DecodeBool()
	m.PreCaptureClear = buf.DecodeBool()
	return nil
}

// TraceCapturePacketsReply defines message 'trace_capture_packets_reply'.
// InProgress: the message form may change in the future versions
type TraceCapturePacketsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TraceCapturePacketsReply) Reset()               { *m = TraceCapturePacketsReply{} }
func (*TraceCapturePacketsReply) GetMessageName() string { return "trace_capture_packets_reply" }
func (*TraceCapturePacketsReply) GetCrcString() string   { return "e8d4e804" }
func (*TraceCapturePacketsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceCapturePacketsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TraceCapturePacketsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TraceCapturePacketsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// TraceClearCapture defines message 'trace_clear_capture'.
// InProgress: the message form may change in the future versions
type TraceClearCapture struct{}

func (m *TraceClearCapture) Reset()               { *m = TraceClearCapture{} }
func (*TraceClearCapture) GetMessageName() string { return "trace_clear_capture" }
func (*TraceClearCapture) GetCrcString() string   { return "51077d14" }
func (*TraceClearCapture) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceClearCapture) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *TraceClearCapture) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *TraceClearCapture) Unmarshal(b []byte) error {
	return nil
}

// TraceClearCaptureReply defines message 'trace_clear_capture_reply'.
// InProgress: the message form may change in the future versions
type TraceClearCaptureReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TraceClearCaptureReply) Reset()               { *m = TraceClearCaptureReply{} }
func (*TraceClearCaptureReply) GetMessageName() string { return "trace_clear_capture_reply" }
func (*TraceClearCaptureReply) GetCrcString() string   { return "e8d4e804" }
func (*TraceClearCaptureReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceClearCaptureReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TraceClearCaptureReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TraceClearCaptureReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// TraceDetails defines message 'trace_details'.
// InProgress: the message form may change in the future versions
type TraceDetails struct {
	ThreadID       uint32 `binapi:"u32,name=thread_id" json:"thread_id,omitempty"`
	Position       uint32 `binapi:"u32,name=position" json:"position,omitempty"`
	MoreThisThread uint8  `binapi:"u8,name=more_this_thread" json:"more_this_thread,omitempty"`
	MoreThreads    uint8  `binapi:"u8,name=more_threads" json:"more_threads,omitempty"`
	Done           uint8  `binapi:"u8,name=done" json:"done,omitempty"`
	PacketNumber   uint32 `binapi:"u32,name=packet_number" json:"packet_number,omitempty"`
	TraceData      string `binapi:"string[],name=trace_data" json:"trace_data,omitempty"`
}

func (m *TraceDetails) Reset()               { *m = TraceDetails{} }
func (*TraceDetails) GetMessageName() string { return "trace_details" }
func (*TraceDetails) GetCrcString() string   { return "1553e9eb" }
func (*TraceDetails) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                    // m.ThreadID
	size += 4                    // m.Position
	size += 1                    // m.MoreThisThread
	size += 1                    // m.MoreThreads
	size += 1                    // m.Done
	size += 4                    // m.PacketNumber
	size += 4 + len(m.TraceData) // m.TraceData
	return size
}
func (m *TraceDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.ThreadID)
	buf.EncodeUint32(m.Position)
	buf.EncodeUint8(m.MoreThisThread)
	buf.EncodeUint8(m.MoreThreads)
	buf.EncodeUint8(m.Done)
	buf.EncodeUint32(m.PacketNumber)
	buf.EncodeString(m.TraceData, 0)
	return buf.Bytes(), nil
}
func (m *TraceDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ThreadID = buf.DecodeUint32()
	m.Position = buf.DecodeUint32()
	m.MoreThisThread = buf.DecodeUint8()
	m.MoreThreads = buf.DecodeUint8()
	m.Done = buf.DecodeUint8()
	m.PacketNumber = buf.DecodeUint32()
	m.TraceData = buf.DecodeString(0)
	return nil
}

// TraceDump defines message 'trace_dump'.
// InProgress: the message form may change in the future versions
type TraceDump struct {
	ClearCache uint8  `binapi:"u8,name=clear_cache" json:"clear_cache,omitempty"`
	ThreadID   uint32 `binapi:"u32,name=thread_id" json:"thread_id,omitempty"`
	Position   uint32 `binapi:"u32,name=position" json:"position,omitempty"`
	MaxRecords uint32 `binapi:"u32,name=max_records" json:"max_records,omitempty"`
}

func (m *TraceDump) Reset()               { *m = TraceDump{} }
func (*TraceDump) GetMessageName() string { return "trace_dump" }
func (*TraceDump) GetCrcString() string   { return "c7d6681f" }
func (*TraceDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.ClearCache
	size += 4 // m.ThreadID
	size += 4 // m.Position
	size += 4 // m.MaxRecords
	return size
}
func (m *TraceDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.ClearCache)
	buf.EncodeUint32(m.ThreadID)
	buf.EncodeUint32(m.Position)
	buf.EncodeUint32(m.MaxRecords)
	return buf.Bytes(), nil
}
func (m *TraceDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ClearCache = buf.DecodeUint8()
	m.ThreadID = buf.DecodeUint32()
	m.Position = buf.DecodeUint32()
	m.MaxRecords = buf.DecodeUint32()
	return nil
}

// TraceDumpReply defines message 'trace_dump_reply'.
// InProgress: the message form may change in the future versions
type TraceDumpReply struct {
	Retval         int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	LastThreadID   uint32 `binapi:"u32,name=last_thread_id" json:"last_thread_id,omitempty"`
	LastPosition   uint32 `binapi:"u32,name=last_position" json:"last_position,omitempty"`
	MoreThisThread uint8  `binapi:"u8,name=more_this_thread" json:"more_this_thread,omitempty"`
	MoreThreads    uint8  `binapi:"u8,name=more_threads" json:"more_threads,omitempty"`
	FlushOnly      uint8  `binapi:"u8,name=flush_only" json:"flush_only,omitempty"`
	Done           uint8  `binapi:"u8,name=done" json:"done,omitempty"`
}

func (m *TraceDumpReply) Reset()               { *m = TraceDumpReply{} }
func (*TraceDumpReply) GetMessageName() string { return "trace_dump_reply" }
func (*TraceDumpReply) GetCrcString() string   { return "e0e87f9d" }
func (*TraceDumpReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceDumpReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.LastThreadID
	size += 4 // m.LastPosition
	size += 1 // m.MoreThisThread
	size += 1 // m.MoreThreads
	size += 1 // m.FlushOnly
	size += 1 // m.Done
	return size
}
func (m *TraceDumpReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.LastThreadID)
	buf.EncodeUint32(m.LastPosition)
	buf.EncodeUint8(m.MoreThisThread)
	buf.EncodeUint8(m.MoreThreads)
	buf.EncodeUint8(m.FlushOnly)
	buf.EncodeUint8(m.Done)
	return buf.Bytes(), nil
}
func (m *TraceDumpReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.LastThreadID = buf.DecodeUint32()
	m.LastPosition = buf.DecodeUint32()
	m.MoreThisThread = buf.DecodeUint8()
	m.MoreThreads = buf.DecodeUint8()
	m.FlushOnly = buf.DecodeUint8()
	m.Done = buf.DecodeUint8()
	return nil
}

// TraceSetFilters defines message 'trace_set_filters'.
// InProgress: the message form may change in the future versions
type TraceSetFilters struct {
	Flag                 TraceFilterFlag `binapi:"trace_filter_flag,name=flag" json:"flag,omitempty"`
	Count                uint32          `binapi:"u32,name=count" json:"count,omitempty"`
	NodeIndex            uint32          `binapi:"u32,name=node_index,default=4294967295" json:"node_index,omitempty"`
	ClassifierTableIndex uint32          `binapi:"u32,name=classifier_table_index,default=4294967295" json:"classifier_table_index,omitempty"`
}

func (m *TraceSetFilters) Reset()               { *m = TraceSetFilters{} }
func (*TraceSetFilters) GetMessageName() string { return "trace_set_filters" }
func (*TraceSetFilters) GetCrcString() string   { return "f522b44a" }
func (*TraceSetFilters) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceSetFilters) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Flag
	size += 4 // m.Count
	size += 4 // m.NodeIndex
	size += 4 // m.ClassifierTableIndex
	return size
}
func (m *TraceSetFilters) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Flag))
	buf.EncodeUint32(m.Count)
	buf.EncodeUint32(m.NodeIndex)
	buf.EncodeUint32(m.ClassifierTableIndex)
	return buf.Bytes(), nil
}
func (m *TraceSetFilters) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Flag = TraceFilterFlag(buf.DecodeUint32())
	m.Count = buf.DecodeUint32()
	m.NodeIndex = buf.DecodeUint32()
	m.ClassifierTableIndex = buf.DecodeUint32()
	return nil
}

// TraceSetFiltersReply defines message 'trace_set_filters_reply'.
// InProgress: the message form may change in the future versions
type TraceSetFiltersReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TraceSetFiltersReply) Reset()               { *m = TraceSetFiltersReply{} }
func (*TraceSetFiltersReply) GetMessageName() string { return "trace_set_filters_reply" }
func (*TraceSetFiltersReply) GetCrcString() string   { return "e8d4e804" }
func (*TraceSetFiltersReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceSetFiltersReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TraceSetFiltersReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TraceSetFiltersReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_tracedump_binapi_init() }
func file_tracedump_binapi_init() {
	api.RegisterMessage((*TraceCapturePackets)(nil), "trace_capture_packets_9e791a9b")
	api.RegisterMessage((*TraceCapturePacketsReply)(nil), "trace_capture_packets_reply_e8d4e804")
	api.RegisterMessage((*TraceClearCapture)(nil), "trace_clear_capture_51077d14")
	api.RegisterMessage((*TraceClearCaptureReply)(nil), "trace_clear_capture_reply_e8d4e804")
	api.RegisterMessage((*TraceDetails)(nil), "trace_details_1553e9eb")
	api.RegisterMessage((*TraceDump)(nil), "trace_dump_c7d6681f")
	api.RegisterMessage((*TraceDumpReply)(nil), "trace_dump_reply_e0e87f9d")
	api.RegisterMessage((*TraceSetFilters)(nil), "trace_set_filters_f522b44a")
	api.RegisterMessage((*TraceSetFiltersReply)(nil), "trace_set_filters_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*TraceCapturePackets)(nil),
		(*TraceCapturePacketsReply)(nil),
		(*TraceClearCapture)(nil),
		(*TraceClearCaptureReply)(nil),
		(*TraceDetails)(nil),
		(*TraceDump)(nil),
		(*TraceDumpReply)(nil),
		(*TraceSetFilters)(nil),
		(*TraceSetFiltersReply)(nil),
	}
}
