// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/contextgraph/v1alpha1/raw_in_memory.proto

package contextgraph

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "istio.io/istio/mixer/adapter/stackdriver/internal/google.golang.org/genproto/googleapis/graphservice/v0"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

func init() {
	proto.RegisterFile("google/cloud/contextgraph/v1alpha1/raw_in_memory.proto", fileDescriptor_raw_in_memory_b2b45986a6697c05)
}

var fileDescriptor_raw_in_memory_b2b45986a6697c05 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x8f, 0xbd, 0xae, 0xc2, 0x30,
	0x0c, 0x85, 0xb7, 0x3b, 0xdc, 0x91, 0xb1, 0x0b, 0x12, 0x1b, 0x4b, 0x4c, 0x85, 0xc4, 0xc2, 0xc6,
	0x03, 0x20, 0x66, 0x96, 0xca, 0x2d, 0xc6, 0x8d, 0x94, 0xc4, 0x91, 0x9b, 0x16, 0x78, 0x7b, 0xa4,
	0xa4, 0x48, 0x65, 0x61, 0x3c, 0x3f, 0xdf, 0xb1, 0xfc, 0x7f, 0x60, 0x11, 0x76, 0x04, 0x9d, 0x93,
	0xf1, 0x06, 0x9d, 0x84, 0x44, 0xcf, 0xc4, 0x8a, 0xb1, 0x87, 0xa9, 0x46, 0x17, 0x7b, 0xac, 0x41,
	0xf1, 0xd1, 0xd8, 0xd0, 0x78, 0xf2, 0xa2, 0x2f, 0x13, 0x55, 0x92, 0xac, 0x36, 0x85, 0x33, 0x99,
	0x33, 0x4b, 0xce, 0x7c, 0xb8, 0x6a, 0x3b, 0x6f, 0x67, 0x7b, 0x20, 0x9d, 0x6c, 0x47, 0x30, 0xed,
	0xbe, 0x74, 0x99, 0xab, 0xd6, 0x73, 0x35, 0xab, 0x76, 0xbc, 0x43, 0xb2, 0x9e, 0x86, 0x84, 0x3e,
	0x96, 0xc2, 0xe9, 0x72, 0x3d, 0xcf, 0x17, 0x59, 0x1c, 0x06, 0x36, 0xa2, 0x0c, 0x4c, 0x21, 0xa7,
	0x50, 0x22, 0x8c, 0x76, 0xf8, 0xf5, 0xc8, 0x71, 0xe9, 0xb6, 0x7f, 0x19, 0xdd, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xd2, 0x7e, 0xff, 0x78, 0x02, 0x01, 0x00, 0x00,
}