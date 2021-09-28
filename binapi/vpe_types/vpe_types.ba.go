// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              21.06.0-13~g94dced301
// source: /usr/share/vpp/api/core/vpe_types.api.json

// Package vpe_types contains generated bindings for API file vpe_types.api.
//
// Contents:
//   2 aliases
//   1 enum
//   1 struct
//
package vpe_types

import (
	"strconv"
	"time"

	api "git.fd.io/govpp.git/api"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

// LogLevel defines enum 'log_level'.
type LogLevel uint32

const (
	VPE_API_LOG_LEVEL_EMERG    LogLevel = 0
	VPE_API_LOG_LEVEL_ALERT    LogLevel = 1
	VPE_API_LOG_LEVEL_CRIT     LogLevel = 2
	VPE_API_LOG_LEVEL_ERR      LogLevel = 3
	VPE_API_LOG_LEVEL_WARNING  LogLevel = 4
	VPE_API_LOG_LEVEL_NOTICE   LogLevel = 5
	VPE_API_LOG_LEVEL_INFO     LogLevel = 6
	VPE_API_LOG_LEVEL_DEBUG    LogLevel = 7
	VPE_API_LOG_LEVEL_DISABLED LogLevel = 8
)

var (
	LogLevel_name = map[uint32]string{
		0: "VPE_API_LOG_LEVEL_EMERG",
		1: "VPE_API_LOG_LEVEL_ALERT",
		2: "VPE_API_LOG_LEVEL_CRIT",
		3: "VPE_API_LOG_LEVEL_ERR",
		4: "VPE_API_LOG_LEVEL_WARNING",
		5: "VPE_API_LOG_LEVEL_NOTICE",
		6: "VPE_API_LOG_LEVEL_INFO",
		7: "VPE_API_LOG_LEVEL_DEBUG",
		8: "VPE_API_LOG_LEVEL_DISABLED",
	}
	LogLevel_value = map[string]uint32{
		"VPE_API_LOG_LEVEL_EMERG":    0,
		"VPE_API_LOG_LEVEL_ALERT":    1,
		"VPE_API_LOG_LEVEL_CRIT":     2,
		"VPE_API_LOG_LEVEL_ERR":      3,
		"VPE_API_LOG_LEVEL_WARNING":  4,
		"VPE_API_LOG_LEVEL_NOTICE":   5,
		"VPE_API_LOG_LEVEL_INFO":     6,
		"VPE_API_LOG_LEVEL_DEBUG":    7,
		"VPE_API_LOG_LEVEL_DISABLED": 8,
	}
)

func (x LogLevel) String() string {
	s, ok := LogLevel_name[uint32(x)]
	if ok {
		return s
	}
	return "LogLevel(" + strconv.Itoa(int(x)) + ")"
}

// Timedelta defines alias 'timedelta'.
type Timedelta float64

// Timestamp defines alias 'timestamp'.
type Timestamp float64

func NewTimestamp(t time.Time) Timestamp {
	sec := int64(t.Unix())
	nsec := int32(t.Nanosecond())
	ns := float64(sec) + float64(nsec/1e9)
	return Timestamp(ns)
}
func (x Timestamp) ToTime() time.Time {
	ns := int64(x * 1e9)
	sec := ns / 1e9
	nsec := ns % 1e9
	return time.Unix(sec, nsec)
}
func (x Timestamp) String() string {
	return x.ToTime().String()
}
func (x *Timestamp) MarshalText() ([]byte, error) {
	return []byte(x.ToTime().Format(time.RFC3339Nano)), nil
}
func (x *Timestamp) UnmarshalText(text []byte) error {
	t, err := time.Parse(time.RFC3339Nano, string(text))
	if err != nil {
		return err
	}
	*x = NewTimestamp(t)
	return nil
}

// Version defines type 'version'.
type Version struct {
	Major         uint32 `binapi:"u32,name=major" json:"major,omitempty"`
	Minor         uint32 `binapi:"u32,name=minor" json:"minor,omitempty"`
	Patch         uint32 `binapi:"u32,name=patch" json:"patch,omitempty"`
	PreRelease    []byte `binapi:"u8[17],name=pre_release" json:"pre_release,omitempty"`
	BuildMetadata []byte `binapi:"u8[17],name=build_metadata" json:"build_metadata,omitempty"`
}
