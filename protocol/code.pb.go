// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/code.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Code int32

const (
	Code_ok Code = 0
)

var Code_name = map[int32]string{
	0: "ok",
}

var Code_value = map[string]int32{
	"ok": 0,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}

func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4153a92670452c72, []int{0}
}

func init() {
	proto.RegisterEnum("protocol.Code", Code_name, Code_value)
}

func init() { proto.RegisterFile("proto/code.proto", fileDescriptor_4153a92670452c72) }

var fileDescriptor_4153a92670452c72 = []byte{
	// 75 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0x4f, 0x49, 0xd5, 0x03, 0x33, 0x85, 0x38, 0xc0, 0x54, 0x72, 0x7e, 0x8e,
	0x16, 0x1f, 0x17, 0x8b, 0x73, 0x7e, 0x4a, 0xaa, 0x10, 0x1b, 0x17, 0x53, 0x7e, 0xb6, 0x00, 0x83,
	0x13, 0x6f, 0x14, 0x77, 0x7a, 0xbe, 0x35, 0x4c, 0x3a, 0x89, 0x0d, 0xcc, 0x32, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xc7, 0xbf, 0x09, 0xb2, 0x43, 0x00, 0x00, 0x00,
}