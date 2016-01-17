// Code generated by protoc-gen-go.
// source: tensorflow/python/training/checkpoint_state.proto
// DO NOT EDIT!

/*
Package tensorflow is a generated protocol buffer package.

It is generated from these files:
	tensorflow/python/training/checkpoint_state.proto
	tensorflow/python/training/saver.proto

It has these top-level messages:
	CheckpointState
	SaverDef
*/
package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Protocol buffer representing the checkpoint state.
//
// TODO(touts): Add other attributes as needed.
type CheckpointState struct {
	// Path to the most-recent model checkpoint.
	ModelCheckpointPath string `protobuf:"bytes,1,opt,name=model_checkpoint_path" json:"model_checkpoint_path,omitempty"`
	// Paths to all not-yet-deleted model checkpoints, sorted from oldest to
	// newest.
	// Note that the value of model_checkpoint_path should be the last item in
	// this list.
	AllModelCheckpointPaths []string `protobuf:"bytes,2,rep,name=all_model_checkpoint_paths" json:"all_model_checkpoint_paths,omitempty"`
}

func (m *CheckpointState) Reset()                    { *m = CheckpointState{} }
func (m *CheckpointState) String() string            { return proto.CompactTextString(m) }
func (*CheckpointState) ProtoMessage()               {}
func (*CheckpointState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*CheckpointState)(nil), "tensorflow.CheckpointState")
}

var fileDescriptor0 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x2c, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0xd7, 0x2f, 0xa8, 0x2c, 0xc9, 0xc8, 0xcf, 0xd3, 0x2f, 0x29,
	0x4a, 0xcc, 0xcc, 0xcb, 0xcc, 0x4b, 0xd7, 0x4f, 0xce, 0x48, 0x4d, 0xce, 0x2e, 0xc8, 0xcf, 0xcc,
	0x2b, 0x89, 0x2f, 0x2e, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42,
	0x68, 0x51, 0x0a, 0xe1, 0xe2, 0x77, 0x86, 0xab, 0x0a, 0x06, 0x29, 0x12, 0x92, 0xe5, 0x12, 0xcd,
	0xcd, 0x4f, 0x49, 0xcd, 0x89, 0x47, 0xd2, 0x5e, 0x90, 0x58, 0x92, 0x21, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x29, 0xa4, 0xc4, 0x25, 0x95, 0x98, 0x93, 0x13, 0x8f, 0x55, 0x49, 0xb1, 0x04, 0x93, 0x02,
	0xb3, 0x06, 0x67, 0x12, 0x1b, 0xd8, 0x22, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x81,
	0x73, 0x93, 0x9d, 0x00, 0x00, 0x00,
}