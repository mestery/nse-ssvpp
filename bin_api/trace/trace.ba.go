// Code generated by GoVPP binapi-generator. DO NOT EDIT.
// source: api/trace.api.json

/*
Package trace is a generated VPP binary API of the 'trace' VPP module.

It is generated from this file:
	trace.api.json

It contains these VPP binary API objects:
	6 messages
	3 services
*/
package trace

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

/* Messages */

// TraceProfileAdd represents the VPP binary API message 'trace_profile_add'.
// Generated from 'trace.api.json', line 4:
//
//            "trace_profile_add",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "trace_type"
//            ],
//            [
//                "u8",
//                "num_elts"
//            ],
//            [
//                "u8",
//                "trace_tsp"
//            ],
//            [
//                "u32",
//                "node_id"
//            ],
//            [
//                "u32",
//                "app_data"
//            ],
//            {
//                "crc": "0xde08aa6d"
//            }
//
type TraceProfileAdd struct {
	TraceType uint8
	NumElts   uint8
	TraceTsp  uint8
	NodeID    uint32
	AppData   uint32
}

func (*TraceProfileAdd) GetMessageName() string {
	return "trace_profile_add"
}
func (*TraceProfileAdd) GetCrcString() string {
	return "de08aa6d"
}
func (*TraceProfileAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewTraceProfileAdd() api.Message {
	return &TraceProfileAdd{}
}

// TraceProfileAddReply represents the VPP binary API message 'trace_profile_add_reply'.
// Generated from 'trace.api.json', line 42:
//
//            "trace_profile_add_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type TraceProfileAddReply struct {
	Retval int32
}

func (*TraceProfileAddReply) GetMessageName() string {
	return "trace_profile_add_reply"
}
func (*TraceProfileAddReply) GetCrcString() string {
	return "e8d4e804"
}
func (*TraceProfileAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewTraceProfileAddReply() api.Message {
	return &TraceProfileAddReply{}
}

// TraceProfileDel represents the VPP binary API message 'trace_profile_del'.
// Generated from 'trace.api.json', line 60:
//
//            "trace_profile_del",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type TraceProfileDel struct{}

func (*TraceProfileDel) GetMessageName() string {
	return "trace_profile_del"
}
func (*TraceProfileDel) GetCrcString() string {
	return "51077d14"
}
func (*TraceProfileDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewTraceProfileDel() api.Message {
	return &TraceProfileDel{}
}

// TraceProfileDelReply represents the VPP binary API message 'trace_profile_del_reply'.
// Generated from 'trace.api.json', line 78:
//
//            "trace_profile_del_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type TraceProfileDelReply struct {
	Retval int32
}

func (*TraceProfileDelReply) GetMessageName() string {
	return "trace_profile_del_reply"
}
func (*TraceProfileDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*TraceProfileDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewTraceProfileDelReply() api.Message {
	return &TraceProfileDelReply{}
}

// TraceProfileShowConfig represents the VPP binary API message 'trace_profile_show_config'.
// Generated from 'trace.api.json', line 96:
//
//            "trace_profile_show_config",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type TraceProfileShowConfig struct{}

func (*TraceProfileShowConfig) GetMessageName() string {
	return "trace_profile_show_config"
}
func (*TraceProfileShowConfig) GetCrcString() string {
	return "51077d14"
}
func (*TraceProfileShowConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewTraceProfileShowConfig() api.Message {
	return &TraceProfileShowConfig{}
}

// TraceProfileShowConfigReply represents the VPP binary API message 'trace_profile_show_config_reply'.
// Generated from 'trace.api.json', line 114:
//
//            "trace_profile_show_config_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            [
//                "u8",
//                "trace_type"
//            ],
//            [
//                "u8",
//                "num_elts"
//            ],
//            [
//                "u8",
//                "trace_tsp"
//            ],
//            [
//                "u32",
//                "node_id"
//            ],
//            [
//                "u32",
//                "app_data"
//            ],
//            {
//                "crc": "0x0f1d374c"
//            }
//
type TraceProfileShowConfigReply struct {
	Retval    int32
	TraceType uint8
	NumElts   uint8
	TraceTsp  uint8
	NodeID    uint32
	AppData   uint32
}

func (*TraceProfileShowConfigReply) GetMessageName() string {
	return "trace_profile_show_config_reply"
}
func (*TraceProfileShowConfigReply) GetCrcString() string {
	return "0f1d374c"
}
func (*TraceProfileShowConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewTraceProfileShowConfigReply() api.Message {
	return &TraceProfileShowConfigReply{}
}

/* Services */

type Services interface {
	TraceProfileAdd(*TraceProfileAdd) (*TraceProfileAddReply, error)
	TraceProfileDel(*TraceProfileDel) (*TraceProfileDelReply, error)
	TraceProfileShowConfig(*TraceProfileShowConfig) (*TraceProfileShowConfigReply, error)
}

func init() {
	api.RegisterMessage((*TraceProfileAdd)(nil), "trace.TraceProfileAdd")
	api.RegisterMessage((*TraceProfileAddReply)(nil), "trace.TraceProfileAddReply")
	api.RegisterMessage((*TraceProfileDel)(nil), "trace.TraceProfileDel")
	api.RegisterMessage((*TraceProfileDelReply)(nil), "trace.TraceProfileDelReply")
	api.RegisterMessage((*TraceProfileShowConfig)(nil), "trace.TraceProfileShowConfig")
	api.RegisterMessage((*TraceProfileShowConfigReply)(nil), "trace.TraceProfileShowConfigReply")
}
