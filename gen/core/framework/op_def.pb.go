// Code generated by protoc-gen-go.
// source: tensorflow/core/framework/op_def.proto
// DO NOT EDIT!

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Defines an operation. A NodeDef in a GraphDef specifies an Op by
// using the "op" field which should match the name of a OpDef.
type OpDef struct {
	// Op names starting with an underscore are reserved for internal use.
	// Names should be CamelCase and match the regexp "[A-Z][a-zA-Z0-9_]*".
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Description of the input(s).
	InputArg []*OpDef_ArgDef `protobuf:"bytes,2,rep,name=input_arg" json:"input_arg,omitempty"`
	// Description of the output(s).
	OutputArg []*OpDef_ArgDef  `protobuf:"bytes,3,rep,name=output_arg" json:"output_arg,omitempty"`
	Attr      []*OpDef_AttrDef `protobuf:"bytes,4,rep,name=attr" json:"attr,omitempty"`
	// One-line human-readable description of what the Op does.
	Summary string `protobuf:"bytes,5,opt,name=summary" json:"summary,omitempty"`
	// Additional, longer human-readable description of what the Op does.
	Description string `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	// True if the operation is commutative ("op(a,b) == op(b,a)" for all inputs)
	IsCommutative bool `protobuf:"varint,18,opt,name=is_commutative" json:"is_commutative,omitempty"`
	// If is_aggregate is true, then this operation accepts N >= 2
	// inputs and produces 1 output all of the same type.  Should be
	// associative and commutative, and produce output with the same
	// shape as the input.  The optimizer may replace an aggregate op
	// taking input from multiple devices with a tree of aggregate ops
	// that aggregate locally within each device (and possibly within
	// groups of nearby devices) before communicating.
	// TODO(josh11b): Implement that optimization.
	IsAggregate bool `protobuf:"varint,16,opt,name=is_aggregate" json:"is_aggregate,omitempty"`
	// By default Ops may be moved between devices.  Stateful ops should
	// either not be moved, or should only be moved if that state can also
	// be moved (e.g. via some sort of save / restore).
	// Stateful ops are guaranteed to never be optimized away by Common
	// Subexpression Elimination (CSE).
	IsStateful bool `protobuf:"varint,17,opt,name=is_stateful" json:"is_stateful,omitempty"`
	// By default, all inputs to an Op must be initialized Tensors.  Ops
	// that may initialize tensors for the first time should set this
	// field to true, to allow the Op to take an uninitialized Tensor as
	// input.
	AllowsUninitializedInput bool `protobuf:"varint,19,opt,name=allows_uninitialized_input" json:"allows_uninitialized_input,omitempty"`
}

func (m *OpDef) Reset()                    { *m = OpDef{} }
func (m *OpDef) String() string            { return proto.CompactTextString(m) }
func (*OpDef) ProtoMessage()               {}
func (*OpDef) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *OpDef) GetInputArg() []*OpDef_ArgDef {
	if m != nil {
		return m.InputArg
	}
	return nil
}

func (m *OpDef) GetOutputArg() []*OpDef_ArgDef {
	if m != nil {
		return m.OutputArg
	}
	return nil
}

func (m *OpDef) GetAttr() []*OpDef_AttrDef {
	if m != nil {
		return m.Attr
	}
	return nil
}

// For describing inputs and outputs.
type OpDef_ArgDef struct {
	// Name for the input/output.  Should match the regexp "[a-z][a-z0-9_]*".
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Human readable description.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Describes the type of one or more tensors that are accepted/produced
	// by this input/output arg.  The only legal combinations are:
	// * For a single tensor: either the "type" field is set or the
	//   "type_attr" field is set to the name of an attr with type "type".
	// * For a sequence of tensors with the same type: the "number_attr"
	//   field will be set to the name of an attr with type "int", and
	//   either the "type" or "type_attr" field will be set as for
	//   single tensors.
	// * For a sequence of tensors, the "type_list_attr" field will be set
	//   to the name of an attr with type "list(type)".
	Type       DataType `protobuf:"varint,3,opt,name=type,enum=tensorflow.DataType" json:"type,omitempty"`
	TypeAttr   string   `protobuf:"bytes,4,opt,name=type_attr" json:"type_attr,omitempty"`
	NumberAttr string   `protobuf:"bytes,5,opt,name=number_attr" json:"number_attr,omitempty"`
	// If specified, attr must have type "list(type)", and none of
	// type, type_attr, and number_attr may be specified.
	TypeListAttr string `protobuf:"bytes,6,opt,name=type_list_attr" json:"type_list_attr,omitempty"`
	// For inputs: if true, the inputs are required to be refs.
	//   By default, inputs can be either refs or non-refs.
	// For outputs: if true, outputs are refs, otherwise they are not.
	IsRef bool `protobuf:"varint,16,opt,name=is_ref" json:"is_ref,omitempty"`
}

func (m *OpDef_ArgDef) Reset()                    { *m = OpDef_ArgDef{} }
func (m *OpDef_ArgDef) String() string            { return proto.CompactTextString(m) }
func (*OpDef_ArgDef) ProtoMessage()               {}
func (*OpDef_ArgDef) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0, 0} }

// Description of the graph-construction-time configuration of this
// Op.  That is to say, this describes the attr fields that will
// be specified in the NodeDef.
type OpDef_AttrDef struct {
	// A descriptive name for the argument.  May be used, e.g. by the
	// Python client, as a keyword argument name, and so should match
	// the regexp "[a-z][a-z0-9_]+".
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// One of the type names from attr_value.proto ("string", "list(string)",
	// "int", etc.).
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// A reasonable default for this attribute if the user does not supply
	// a value.  If not specified, the user must supply a value.
	DefaultValue *AttrValue `protobuf:"bytes,3,opt,name=default_value" json:"default_value,omitempty"`
	// Human-readable description.
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// For type == "int", this is a minimum value.  For "list(___)"
	// types, this is the minimum length.
	HasMinimum bool  `protobuf:"varint,5,opt,name=has_minimum" json:"has_minimum,omitempty"`
	Minimum    int64 `protobuf:"varint,6,opt,name=minimum" json:"minimum,omitempty"`
	// The set of allowed values.  Has type that is the "list" version
	// of the "type" field above (uses the ..._list, fields of AttrValue).
	// If type == "type" or "list(type)" above, then the type_list field
	// of allowed_values has the set of allowed DataTypes.
	// If type == "string" or "list(string)", then the s_list field has
	// the set of allowed strings.
	AllowedValues *AttrValue `protobuf:"bytes,7,opt,name=allowed_values" json:"allowed_values,omitempty"`
}

func (m *OpDef_AttrDef) Reset()                    { *m = OpDef_AttrDef{} }
func (m *OpDef_AttrDef) String() string            { return proto.CompactTextString(m) }
func (*OpDef_AttrDef) ProtoMessage()               {}
func (*OpDef_AttrDef) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0, 1} }

func (m *OpDef_AttrDef) GetDefaultValue() *AttrValue {
	if m != nil {
		return m.DefaultValue
	}
	return nil
}

func (m *OpDef_AttrDef) GetAllowedValues() *AttrValue {
	if m != nil {
		return m.AllowedValues
	}
	return nil
}

// A collection of OpDefs
type OpList struct {
	Op []*OpDef `protobuf:"bytes,1,rep,name=op" json:"op,omitempty"`
}

func (m *OpList) Reset()                    { *m = OpList{} }
func (m *OpList) String() string            { return proto.CompactTextString(m) }
func (*OpList) ProtoMessage()               {}
func (*OpList) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *OpList) GetOp() []*OpDef {
	if m != nil {
		return m.Op
	}
	return nil
}

func init() {
	proto.RegisterType((*OpDef)(nil), "tensorflow.OpDef")
	proto.RegisterType((*OpDef_ArgDef)(nil), "tensorflow.OpDef.ArgDef")
	proto.RegisterType((*OpDef_AttrDef)(nil), "tensorflow.OpDef.AttrDef")
	proto.RegisterType((*OpList)(nil), "tensorflow.OpList")
}

var fileDescriptor7 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0x5d, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0x35, 0xf9, 0x99, 0x34, 0x37, 0x21, 0x90, 0x69, 0x41, 0x66, 0x24, 0xa4, 0x2a, 0x12,
	0x34, 0x82, 0x92, 0x48, 0x61, 0x05, 0x48, 0x7d, 0x44, 0xea, 0x0b, 0xe2, 0xd5, 0x72, 0x13, 0xcf,
	0x60, 0x31, 0x1e, 0x5b, 0xfe, 0x69, 0x55, 0x56, 0xc0, 0x2e, 0x58, 0x09, 0x7b, 0xe3, 0xda, 0xce,
	0xd0, 0x1f, 0x28, 0x7d, 0x3d, 0xf3, 0xdd, 0xeb, 0x73, 0xce, 0x5c, 0x78, 0xe3, 0x78, 0x6b, 0x95,
	0xa9, 0x1a, 0x75, 0xb5, 0xde, 0x2a, 0xc3, 0xd7, 0x95, 0x61, 0x92, 0x5f, 0x29, 0xf3, 0x6d, 0xad,
	0x34, 0xdd, 0xf1, 0x6a, 0xa5, 0x8d, 0x72, 0xaa, 0x80, 0x1b, 0xae, 0x7c, 0xfb, 0xf0, 0x0c, 0x73,
	0xce, 0xd0, 0x4b, 0xd6, 0x78, 0x9e, 0xe6, 0xca, 0xd7, 0x0f, 0xb3, 0xee, 0x5a, 0x73, 0x9b, 0xb0,
	0xc5, 0x8f, 0x21, 0x0c, 0xcf, 0xf5, 0x19, 0xaf, 0x8a, 0x29, 0x0c, 0x5a, 0x24, 0x48, 0x76, 0x9c,
	0x2d, 0xc7, 0xc5, 0x3b, 0x18, 0x8b, 0x56, 0x7b, 0x47, 0x99, 0xa9, 0x49, 0xef, 0xb8, 0xbf, 0x9c,
	0x6c, 0xc8, 0xea, 0x66, 0xe5, 0x2a, 0xce, 0xac, 0x3e, 0x9a, 0x3a, 0x8c, 0x9e, 0x02, 0x28, 0xef,
	0x3a, 0xba, 0xff, 0x08, 0x7d, 0x02, 0x83, 0xe0, 0x96, 0x0c, 0x22, 0xf7, 0xf2, 0x1f, 0x1c, 0x7e,
	0x0d, 0xe0, 0x53, 0x18, 0x59, 0x2f, 0x25, 0x33, 0xd7, 0x64, 0x18, 0x4d, 0x1d, 0xc2, 0x64, 0xc7,
	0xed, 0xd6, 0x08, 0xed, 0x84, 0x6a, 0x49, 0x1e, 0xc5, 0x17, 0x30, 0x13, 0x96, 0x6e, 0x95, 0x94,
	0xde, 0x31, 0x27, 0x2e, 0x39, 0x29, 0x50, 0x3f, 0x28, 0x8e, 0x60, 0x8a, 0x3a, 0xab, 0x6b, 0xc3,
	0x6b, 0xe6, 0x38, 0x79, 0x16, 0x55, 0x5c, 0x81, 0xaa, 0x45, 0x92, 0x57, 0xbe, 0x21, 0xf3, 0x28,
	0x2e, 0xa0, 0x64, 0x0d, 0xbe, 0x6f, 0xa9, 0x6f, 0x45, 0x2b, 0x9c, 0x60, 0x8d, 0xf8, 0xce, 0x77,
	0x34, 0x36, 0x40, 0x0e, 0x03, 0x53, 0xfe, 0xcc, 0x20, 0xdf, 0x07, 0xb8, 0xdb, 0xd4, 0x3d, 0x53,
	0xbd, 0x28, 0x2e, 0x60, 0x10, 0x5a, 0xc6, 0x2e, 0xb2, 0xe5, 0x6c, 0x73, 0x74, 0x3b, 0xe3, 0x19,
	0x73, 0xec, 0x33, 0x7e, 0x2b, 0xe6, 0x30, 0x0e, 0x0c, 0xdd, 0x97, 0xb1, 0xdf, 0xd5, 0x7a, 0x79,
	0xc1, 0x4d, 0x12, 0x87, 0x5d, 0xc0, 0xc8, 0x35, 0xc2, 0xba, 0xa4, 0xa7, 0xe0, 0x33, 0xc8, 0x31,
	0x8a, 0xe1, 0x55, 0x8a, 0x56, 0xfe, 0xca, 0x60, 0xd4, 0x55, 0x77, 0xd7, 0xe2, 0x74, 0xef, 0x26,
	0x79, 0x3b, 0x85, 0x27, 0x78, 0x5e, 0xcc, 0x37, 0x2e, 0x1d, 0x4c, 0x34, 0x39, 0xd9, 0x3c, 0xbf,
	0x6d, 0x32, 0xec, 0xf9, 0x12, 0x3e, 0xde, 0x8f, 0xf7, 0xc7, 0xe7, 0x57, 0x66, 0xa9, 0xc4, 0xb2,
	0xa4, 0x97, 0xd1, 0xe7, 0x41, 0xf8, 0x5d, 0x9d, 0x10, 0x0c, 0xf6, 0x8b, 0xf7, 0x30, 0x8b, 0xb5,
	0x62, 0x93, 0xf1, 0x21, 0x4b, 0x46, 0xff, 0x79, 0x69, 0x71, 0x02, 0xf9, 0xb9, 0xfe, 0x84, 0x21,
	0x8b, 0x57, 0xd0, 0x53, 0x1a, 0xbd, 0x87, 0xfb, 0x98, 0xff, 0x75, 0x1f, 0x17, 0x79, 0x3c, 0xdd,
	0x0f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xea, 0x01, 0xcd, 0xbd, 0x43, 0x03, 0x00, 0x00,
}