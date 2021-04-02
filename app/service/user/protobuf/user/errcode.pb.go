// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protocol/user/errcode.proto

package user

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

// [10000 - 11000] 用户错误码范围。
type ErrCode int32

const (
	ErrCode_Success           ErrCode = 0
	ErrCode_UserNotFound      ErrCode = 10001
	ErrCode_PasswordIncorrect ErrCode = 10002
	ErrCode_UserNotLogin      ErrCode = 10003
)

var ErrCode_name = map[int32]string{
	0:     "Success",
	10001: "UserNotFound",
	10002: "PasswordIncorrect",
	10003: "UserNotLogin",
}

var ErrCode_value = map[string]int32{
	"Success":           0,
	"UserNotFound":      10001,
	"PasswordIncorrect": 10002,
	"UserNotLogin":      10003,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cbe3c2346557a30c, []int{0}
}

func init() {
	proto.RegisterEnum("user.ErrCode", ErrCode_name, ErrCode_value)
}

func init() { proto.RegisterFile("protocol/user/errcode.proto", fileDescriptor_cbe3c2346557a30c) }

var fileDescriptor_cbe3c2346557a30c = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x4f, 0x2d, 0x2a, 0x4a, 0xce, 0x4f,
	0x49, 0xd5, 0x03, 0x8b, 0x0a, 0xb1, 0x80, 0xc4, 0xb4, 0x42, 0xb8, 0xd8, 0x5d, 0x8b, 0x8a, 0x9c,
	0xf3, 0x53, 0x52, 0x85, 0xb8, 0xb9, 0xd8, 0x83, 0x4b, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0x05, 0x18,
	0x84, 0x04, 0xb9, 0x78, 0x42, 0x8b, 0x53, 0x8b, 0xfc, 0xf2, 0x4b, 0xdc, 0xf2, 0x4b, 0xf3, 0x52,
	0x04, 0x26, 0xfa, 0x09, 0x89, 0x71, 0x09, 0x06, 0x24, 0x16, 0x17, 0x97, 0xe7, 0x17, 0xa5, 0x78,
	0xe6, 0x25, 0xe7, 0x17, 0x15, 0xa5, 0x26, 0x97, 0x08, 0x4c, 0xf2, 0x43, 0x52, 0xea, 0x93, 0x9f,
	0x9e, 0x99, 0x27, 0x30, 0xd9, 0xcf, 0x49, 0xfe, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x8a, 0x17, 0x6c, 0x69, 0x52, 0x69, 0x1a,
	0xd8, 0x29, 0x49, 0x6c, 0x60, 0xae, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x89, 0x77, 0x55, 0xf2,
	0xa2, 0x00, 0x00, 0x00,
}
