// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/jkmathew/antha/api/v1/state.proto

package org_antha_lang_antha_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type State int32

const (
	// Task created (initial state)
	State_CREATED State = 0
	// Task eligible to run
	State_SCHEDULED State = 1
	// Task waiting on external input to run
	State_WAITING State = 2
	// Task running
	State_RUNNING State = 3
	// Task finished running successfully
	State_SUCCEEDED State = 4
	// Task finished running unsuccessfully
	State_FAILED State = 5
)

var State_name = map[int32]string{
	0: "CREATED",
	1: "SCHEDULED",
	2: "WAITING",
	3: "RUNNING",
	4: "SUCCEEDED",
	5: "FAILED",
}
var State_value = map[string]int32{
	"CREATED":   0,
	"SCHEDULED": 1,
	"WAITING":   2,
	"RUNNING":   3,
	"SUCCEEDED": 4,
	"FAILED":    5,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}
func (State) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type Status struct {
	State State `protobuf:"varint,1,opt,name=state,enum=org.antha_lang.antha.v1.State" json:"state,omitempty"`
	// Any message associated with the current state
	Message_ string `protobuf:"bytes,2,opt,name=message_,json=message" json:"message_,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Status) GetState() State {
	if m != nil {
		return m.State
	}
	return State_CREATED
}

func (m *Status) GetMessage_() string {
	if m != nil {
		return m.Message_
	}
	return ""
}

func init() {
	proto.RegisterType((*Status)(nil), "org.antha_lang.antha.v1.Status")
	proto.RegisterEnum("org.antha_lang.antha.v1.State", State_name, State_value)
}

func init() { proto.RegisterFile("github.com/jkmathew/antha/api/v1/state.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcc, 0x2b, 0xc9, 0x48, 0xd4, 0xcd, 0x49, 0xcc,
	0x4b, 0x87, 0x30, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0xcb, 0x0c, 0xf5, 0x8b, 0x4b, 0x12, 0x4b, 0x52,
	0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0xf3, 0x8b, 0xd2, 0xf5, 0xc0, 0xb2, 0xf1, 0x20,
	0x85, 0x10, 0xa6, 0x5e, 0x99, 0xa1, 0x52, 0x24, 0x17, 0x5b, 0x70, 0x49, 0x62, 0x49, 0x69, 0xb1,
	0x90, 0x09, 0x17, 0x2b, 0x58, 0x87, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x9f, 0x91, 0x9c, 0x1e, 0x0e,
	0x2d, 0x7a, 0x20, 0xf5, 0xa9, 0x41, 0x10, 0xc5, 0x42, 0x92, 0x5c, 0x1c, 0xb9, 0xa9, 0xc5, 0xc5,
	0x89, 0xe9, 0xa9, 0xf1, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xec, 0x50, 0xbe, 0x56, 0x04,
	0x17, 0x2b, 0x58, 0xa9, 0x10, 0x37, 0x17, 0xbb, 0x73, 0x90, 0xab, 0x63, 0x88, 0xab, 0x8b, 0x00,
	0x83, 0x10, 0x2f, 0x17, 0x67, 0xb0, 0xb3, 0x87, 0xab, 0x4b, 0xa8, 0x8f, 0xab, 0x8b, 0x00, 0x23,
	0x48, 0x2e, 0xdc, 0xd1, 0x33, 0xc4, 0xd3, 0xcf, 0x5d, 0x80, 0x09, 0xc4, 0x09, 0x0a, 0xf5, 0xf3,
	0x03, 0x71, 0x98, 0xc1, 0x0a, 0x43, 0x9d, 0x9d, 0x5d, 0x5d, 0x5d, 0x5c, 0x5d, 0x04, 0x58, 0x84,
	0xb8, 0xb8, 0xd8, 0xdc, 0x1c, 0x3d, 0x41, 0x9a, 0x58, 0x93, 0xd8, 0xc0, 0x9e, 0x32, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xc9, 0x66, 0x71, 0x12, 0x06, 0x01, 0x00, 0x00,
}
