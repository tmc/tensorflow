// Code generated by protoc-gen-go.
// source: tensorflow/core/framework/device_attributes.proto
// DO NOT EDIT!

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// BusAdjacency identifies the ability of a device to participate in
// maximally efficient DMA operations within the local context of a
// process.
//
// This is currently ignored.
type BusAdjacency int32

const (
	BusAdjacency_BUS_0               BusAdjacency = 0
	BusAdjacency_BUS_1               BusAdjacency = 1
	BusAdjacency_BUS_ANY             BusAdjacency = 2
	BusAdjacency_BUS_NUM_ADJACENCIES BusAdjacency = 3
)

var BusAdjacency_name = map[int32]string{
	0: "BUS_0",
	1: "BUS_1",
	2: "BUS_ANY",
	3: "BUS_NUM_ADJACENCIES",
}
var BusAdjacency_value = map[string]int32{
	"BUS_0":               0,
	"BUS_1":               1,
	"BUS_ANY":             2,
	"BUS_NUM_ADJACENCIES": 3,
}

func (x BusAdjacency) String() string {
	return proto.EnumName(BusAdjacency_name, int32(x))
}
func (BusAdjacency) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type DeviceAttributes struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// String representation of device_type.
	DeviceType string `protobuf:"bytes,2,opt,name=device_type" json:"device_type,omitempty"`
	// Memory capacity of device in bytes.
	MemoryLimit  int64        `protobuf:"varint,4,opt,name=memory_limit" json:"memory_limit,omitempty"`
	BusAdjacency BusAdjacency `protobuf:"varint,5,opt,name=bus_adjacency,enum=tensorflow.BusAdjacency" json:"bus_adjacency,omitempty"`
	// A device is assigned a global unique number each time it is
	// initialized. "incarnation" should never be 0.
	Incarnation uint64 `protobuf:"fixed64,6,opt,name=incarnation" json:"incarnation,omitempty"`
	// String representation of the physical device that this device maps to.
	PhysicalDeviceDesc string `protobuf:"bytes,7,opt,name=physical_device_desc" json:"physical_device_desc,omitempty"`
}

func (m *DeviceAttributes) Reset()                    { *m = DeviceAttributes{} }
func (m *DeviceAttributes) String() string            { return proto.CompactTextString(m) }
func (*DeviceAttributes) ProtoMessage()               {}
func (*DeviceAttributes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func init() {
	proto.RegisterType((*DeviceAttributes)(nil), "tensorflow.DeviceAttributes")
	proto.RegisterEnum("tensorflow.BusAdjacency", BusAdjacency_name, BusAdjacency_value)
}

var fileDescriptor3 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x49, 0x7f, 0x52, 0xf5, 0x36, 0xa0, 0xc8, 0xad, 0x84, 0x07, 0x06, 0xc4, 0x84, 0x18,
	0x12, 0x0a, 0x4f, 0xe0, 0xfe, 0x0c, 0x54, 0x22, 0x4b, 0xd5, 0x81, 0xc9, 0x72, 0x1c, 0x57, 0x18,
	0x12, 0x3b, 0xb2, 0x1d, 0xaa, 0x3c, 0x15, 0xaf, 0x48, 0x12, 0xb5, 0x54, 0xdd, 0xce, 0x3d, 0xf7,
	0xdc, 0xef, 0x48, 0x17, 0xe6, 0x4e, 0x28, 0xab, 0xcd, 0x3e, 0xd7, 0x87, 0x98, 0x6b, 0x23, 0xe2,
	0xbd, 0x61, 0x85, 0x38, 0x68, 0xf3, 0x1d, 0x67, 0xe2, 0x47, 0x72, 0x41, 0x99, 0x73, 0x46, 0xa6,
	0x95, 0x13, 0x36, 0x2a, 0x8d, 0x76, 0x1a, 0xc1, 0xf9, 0xe4, 0xe1, 0xd7, 0x83, 0x70, 0xd5, 0xe5,
	0xc8, 0x7f, 0x0c, 0x05, 0x30, 0x50, 0x0d, 0x04, 0x7b, 0xf7, 0xde, 0xe3, 0x18, 0x4d, 0x61, 0x72,
	0x24, 0xb9, 0xba, 0x14, 0xb8, 0xd7, 0x99, 0x33, 0x08, 0x0a, 0x51, 0x68, 0x53, 0xd3, 0x5c, 0x16,
	0xd2, 0xe1, 0x41, 0xe3, 0xf6, 0x51, 0x0c, 0xd7, 0x69, 0x65, 0x29, 0xcb, 0xbe, 0x18, 0x17, 0x8a,
	0xd7, 0x78, 0xd8, 0xd8, 0x37, 0x2f, 0x38, 0x3a, 0x37, 0x46, 0x8b, 0xca, 0x92, 0xd3, 0xbe, 0x65,
	0x4b, 0xc5, 0x99, 0x51, 0xcc, 0x49, 0xad, 0xb0, 0xdf, 0xc4, 0x7d, 0x74, 0x07, 0xb3, 0xf2, 0xb3,
	0xb6, 0x92, 0xb3, 0x9c, 0x1e, 0x9b, 0x33, 0x61, 0x39, 0x1e, 0xb5, 0xcd, 0x4f, 0x1b, 0x08, 0x2e,
	0x10, 0x63, 0x18, 0x2e, 0x76, 0x5b, 0xfa, 0x1c, 0x5e, 0x9d, 0xe4, 0x3c, 0xf4, 0xd0, 0x04, 0x46,
	0xad, 0x24, 0xc9, 0x47, 0xd8, 0x43, 0xb7, 0x30, 0x6d, 0x87, 0x64, 0xf7, 0x4e, 0xc9, 0x6a, 0x43,
	0x96, 0xeb, 0x64, 0xf9, 0xb6, 0xde, 0x86, 0xfd, 0xd4, 0xef, 0x1e, 0xf2, 0xfa, 0x17, 0x00, 0x00,
	0xff, 0xff, 0xd7, 0xfb, 0x86, 0x52, 0x45, 0x01, 0x00, 0x00,
}