// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user_manage.proto

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

// 批量获取用户信息
type GetUsersReq struct {
	Role                 int64    `protobuf:"varint,1,opt,name=role,proto3" json:"role,omitempty"`
	PageIndex            int64    `protobuf:"varint,2,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageNum              int64    `protobuf:"varint,3,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersReq) Reset()         { *m = GetUsersReq{} }
func (m *GetUsersReq) String() string { return proto.CompactTextString(m) }
func (*GetUsersReq) ProtoMessage()    {}
func (*GetUsersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_79107de06ed33a41, []int{0}
}

func (m *GetUsersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersReq.Unmarshal(m, b)
}
func (m *GetUsersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersReq.Marshal(b, m, deterministic)
}
func (m *GetUsersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersReq.Merge(m, src)
}
func (m *GetUsersReq) XXX_Size() int {
	return xxx_messageInfo_GetUsersReq.Size(m)
}
func (m *GetUsersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersReq proto.InternalMessageInfo

func (m *GetUsersReq) GetRole() int64 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *GetUsersReq) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *GetUsersReq) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

type GetUsersResp struct {
	Status               *Status     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Users                []*UserInfo `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	PageIndex            int64       `protobuf:"varint,3,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageNum              int64       `protobuf:"varint,4,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	Total                int64       `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetUsersResp) Reset()         { *m = GetUsersResp{} }
func (m *GetUsersResp) String() string { return proto.CompactTextString(m) }
func (*GetUsersResp) ProtoMessage()    {}
func (*GetUsersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_79107de06ed33a41, []int{1}
}

func (m *GetUsersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersResp.Unmarshal(m, b)
}
func (m *GetUsersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersResp.Marshal(b, m, deterministic)
}
func (m *GetUsersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersResp.Merge(m, src)
}
func (m *GetUsersResp) XXX_Size() int {
	return xxx_messageInfo_GetUsersResp.Size(m)
}
func (m *GetUsersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersResp proto.InternalMessageInfo

func (m *GetUsersResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetUsersResp) GetUsers() []*UserInfo {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *GetUsersResp) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *GetUsersResp) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *GetUsersResp) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

// 批量增加用户
type AddUsersReq struct {
	Users                []*UserInfo `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AddUsersReq) Reset()         { *m = AddUsersReq{} }
func (m *AddUsersReq) String() string { return proto.CompactTextString(m) }
func (*AddUsersReq) ProtoMessage()    {}
func (*AddUsersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_79107de06ed33a41, []int{2}
}

func (m *AddUsersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUsersReq.Unmarshal(m, b)
}
func (m *AddUsersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUsersReq.Marshal(b, m, deterministic)
}
func (m *AddUsersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUsersReq.Merge(m, src)
}
func (m *AddUsersReq) XXX_Size() int {
	return xxx_messageInfo_AddUsersReq.Size(m)
}
func (m *AddUsersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUsersReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddUsersReq proto.InternalMessageInfo

func (m *AddUsersReq) GetUsers() []*UserInfo {
	if m != nil {
		return m.Users
	}
	return nil
}

type AddUsersResp struct {
	Status               *Status     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Succeed              []*UserInfo `protobuf:"bytes,2,rep,name=succeed,proto3" json:"succeed,omitempty"`
	Fail                 []*UserInfo `protobuf:"bytes,3,rep,name=fail,proto3" json:"fail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AddUsersResp) Reset()         { *m = AddUsersResp{} }
func (m *AddUsersResp) String() string { return proto.CompactTextString(m) }
func (*AddUsersResp) ProtoMessage()    {}
func (*AddUsersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_79107de06ed33a41, []int{3}
}

func (m *AddUsersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUsersResp.Unmarshal(m, b)
}
func (m *AddUsersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUsersResp.Marshal(b, m, deterministic)
}
func (m *AddUsersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUsersResp.Merge(m, src)
}
func (m *AddUsersResp) XXX_Size() int {
	return xxx_messageInfo_AddUsersResp.Size(m)
}
func (m *AddUsersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUsersResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddUsersResp proto.InternalMessageInfo

func (m *AddUsersResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AddUsersResp) GetSucceed() []*UserInfo {
	if m != nil {
		return m.Succeed
	}
	return nil
}

func (m *AddUsersResp) GetFail() []*UserInfo {
	if m != nil {
		return m.Fail
	}
	return nil
}

// 批量修改用户
type UpdateUsersReq struct {
	Users                []*UserInfo `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateUsersReq) Reset()         { *m = UpdateUsersReq{} }
func (m *UpdateUsersReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUsersReq) ProtoMessage()    {}
func (*UpdateUsersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_79107de06ed33a41, []int{4}
}

func (m *UpdateUsersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUsersReq.Unmarshal(m, b)
}
func (m *UpdateUsersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUsersReq.Marshal(b, m, deterministic)
}
func (m *UpdateUsersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUsersReq.Merge(m, src)
}
func (m *UpdateUsersReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUsersReq.Size(m)
}
func (m *UpdateUsersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUsersReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUsersReq proto.InternalMessageInfo

func (m *UpdateUsersReq) GetUsers() []*UserInfo {
	if m != nil {
		return m.Users
	}
	return nil
}

type UpdateUsersResp struct {
	Status               *Status     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Succeed              []*UserInfo `protobuf:"bytes,2,rep,name=succeed,proto3" json:"succeed,omitempty"`
	Fail                 []*UserInfo `protobuf:"bytes,3,rep,name=fail,proto3" json:"fail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateUsersResp) Reset()         { *m = UpdateUsersResp{} }
func (m *UpdateUsersResp) String() string { return proto.CompactTextString(m) }
func (*UpdateUsersResp) ProtoMessage()    {}
func (*UpdateUsersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_79107de06ed33a41, []int{5}
}

func (m *UpdateUsersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUsersResp.Unmarshal(m, b)
}
func (m *UpdateUsersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUsersResp.Marshal(b, m, deterministic)
}
func (m *UpdateUsersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUsersResp.Merge(m, src)
}
func (m *UpdateUsersResp) XXX_Size() int {
	return xxx_messageInfo_UpdateUsersResp.Size(m)
}
func (m *UpdateUsersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUsersResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUsersResp proto.InternalMessageInfo

func (m *UpdateUsersResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *UpdateUsersResp) GetSucceed() []*UserInfo {
	if m != nil {
		return m.Succeed
	}
	return nil
}

func (m *UpdateUsersResp) GetFail() []*UserInfo {
	if m != nil {
		return m.Fail
	}
	return nil
}

// 批量删除用户
type DelUsersReq struct {
	UsersId              []int64  `protobuf:"varint,1,rep,packed,name=users_id,json=usersId,proto3" json:"users_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelUsersReq) Reset()         { *m = DelUsersReq{} }
func (m *DelUsersReq) String() string { return proto.CompactTextString(m) }
func (*DelUsersReq) ProtoMessage()    {}
func (*DelUsersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_79107de06ed33a41, []int{6}
}

func (m *DelUsersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelUsersReq.Unmarshal(m, b)
}
func (m *DelUsersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelUsersReq.Marshal(b, m, deterministic)
}
func (m *DelUsersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelUsersReq.Merge(m, src)
}
func (m *DelUsersReq) XXX_Size() int {
	return xxx_messageInfo_DelUsersReq.Size(m)
}
func (m *DelUsersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelUsersReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelUsersReq proto.InternalMessageInfo

func (m *DelUsersReq) GetUsersId() []int64 {
	if m != nil {
		return m.UsersId
	}
	return nil
}

type DelUsersResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Succeed              []int64  `protobuf:"varint,2,rep,packed,name=succeed,proto3" json:"succeed,omitempty"`
	Fail                 []int64  `protobuf:"varint,3,rep,packed,name=fail,proto3" json:"fail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelUsersResp) Reset()         { *m = DelUsersResp{} }
func (m *DelUsersResp) String() string { return proto.CompactTextString(m) }
func (*DelUsersResp) ProtoMessage()    {}
func (*DelUsersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_79107de06ed33a41, []int{7}
}

func (m *DelUsersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelUsersResp.Unmarshal(m, b)
}
func (m *DelUsersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelUsersResp.Marshal(b, m, deterministic)
}
func (m *DelUsersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelUsersResp.Merge(m, src)
}
func (m *DelUsersResp) XXX_Size() int {
	return xxx_messageInfo_DelUsersResp.Size(m)
}
func (m *DelUsersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DelUsersResp.DiscardUnknown(m)
}

var xxx_messageInfo_DelUsersResp proto.InternalMessageInfo

func (m *DelUsersResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DelUsersResp) GetSucceed() []int64 {
	if m != nil {
		return m.Succeed
	}
	return nil
}

func (m *DelUsersResp) GetFail() []int64 {
	if m != nil {
		return m.Fail
	}
	return nil
}

// 获得做题记录
type GetSubmitRecordReq struct {
	PageIndex            int64    `protobuf:"varint,3,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageNum              int64    `protobuf:"varint,4,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	Total                int64    `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSubmitRecordReq) Reset()         { *m = GetSubmitRecordReq{} }
func (m *GetSubmitRecordReq) String() string { return proto.CompactTextString(m) }
func (*GetSubmitRecordReq) ProtoMessage()    {}
func (*GetSubmitRecordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_79107de06ed33a41, []int{8}
}

func (m *GetSubmitRecordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSubmitRecordReq.Unmarshal(m, b)
}
func (m *GetSubmitRecordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSubmitRecordReq.Marshal(b, m, deterministic)
}
func (m *GetSubmitRecordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSubmitRecordReq.Merge(m, src)
}
func (m *GetSubmitRecordReq) XXX_Size() int {
	return xxx_messageInfo_GetSubmitRecordReq.Size(m)
}
func (m *GetSubmitRecordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSubmitRecordReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetSubmitRecordReq proto.InternalMessageInfo

func (m *GetSubmitRecordReq) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *GetSubmitRecordReq) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *GetSubmitRecordReq) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type GetSubmitRecordResp struct {
	Status               *Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	SubmitRecords        []*SubmitRecord `protobuf:"bytes,2,rep,name=submit_records,json=submitRecords,proto3" json:"submit_records,omitempty"`
	PageIndex            int64           `protobuf:"varint,3,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageNum              int64           `protobuf:"varint,4,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	Total                int64           `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetSubmitRecordResp) Reset()         { *m = GetSubmitRecordResp{} }
func (m *GetSubmitRecordResp) String() string { return proto.CompactTextString(m) }
func (*GetSubmitRecordResp) ProtoMessage()    {}
func (*GetSubmitRecordResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_79107de06ed33a41, []int{9}
}

func (m *GetSubmitRecordResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSubmitRecordResp.Unmarshal(m, b)
}
func (m *GetSubmitRecordResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSubmitRecordResp.Marshal(b, m, deterministic)
}
func (m *GetSubmitRecordResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSubmitRecordResp.Merge(m, src)
}
func (m *GetSubmitRecordResp) XXX_Size() int {
	return xxx_messageInfo_GetSubmitRecordResp.Size(m)
}
func (m *GetSubmitRecordResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSubmitRecordResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetSubmitRecordResp proto.InternalMessageInfo

func (m *GetSubmitRecordResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetSubmitRecordResp) GetSubmitRecords() []*SubmitRecord {
	if m != nil {
		return m.SubmitRecords
	}
	return nil
}

func (m *GetSubmitRecordResp) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *GetSubmitRecordResp) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *GetSubmitRecordResp) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*GetUsersReq)(nil), "protocol.GetUsersReq")
	proto.RegisterType((*GetUsersResp)(nil), "protocol.GetUsersResp")
	proto.RegisterType((*AddUsersReq)(nil), "protocol.AddUsersReq")
	proto.RegisterType((*AddUsersResp)(nil), "protocol.AddUsersResp")
	proto.RegisterType((*UpdateUsersReq)(nil), "protocol.UpdateUsersReq")
	proto.RegisterType((*UpdateUsersResp)(nil), "protocol.UpdateUsersResp")
	proto.RegisterType((*DelUsersReq)(nil), "protocol.DelUsersReq")
	proto.RegisterType((*DelUsersResp)(nil), "protocol.DelUsersResp")
	proto.RegisterType((*GetSubmitRecordReq)(nil), "protocol.GetSubmitRecordReq")
	proto.RegisterType((*GetSubmitRecordResp)(nil), "protocol.GetSubmitRecordResp")
}

func init() { proto.RegisterFile("proto/user_manage.proto", fileDescriptor_79107de06ed33a41) }

var fileDescriptor_79107de06ed33a41 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0x4f, 0x4b, 0xe3, 0x40,
	0x18, 0xc6, 0x49, 0xa7, 0xff, 0xf6, 0x4d, 0xb7, 0xbb, 0x3b, 0x5d, 0x76, 0xd3, 0x85, 0x85, 0x92,
	0x83, 0xe4, 0x20, 0x15, 0xea, 0x41, 0x50, 0x3c, 0x28, 0x42, 0xe9, 0xc5, 0x43, 0x4a, 0x2f, 0x7a,
	0x08, 0x69, 0x66, 0x5a, 0x02, 0x49, 0x26, 0x66, 0x26, 0xe0, 0xa7, 0x10, 0xfc, 0x2a, 0x7e, 0x10,
	0x3f, 0x93, 0xe4, 0x4d, 0xd3, 0xa4, 0x8a, 0x95, 0x4a, 0x0f, 0x9e, 0x3a, 0xef, 0xf3, 0x3e, 0xef,
	0xbc, 0xbf, 0xa7, 0x64, 0xe0, 0x6f, 0x9c, 0x08, 0x25, 0x8e, 0x52, 0xc9, 0x13, 0x27, 0x74, 0x23,
	0x77, 0xc9, 0x87, 0xa8, 0xd0, 0x36, 0xfe, 0x78, 0x22, 0xf8, 0x47, 0x73, 0x8b, 0x27, 0xc2, 0x50,
	0x44, 0x79, 0xb7, 0xd0, 0xa4, 0x72, 0x55, 0x2a, 0x73, 0xcd, 0xbc, 0x05, 0x7d, 0xcc, 0xd5, 0x4c,
	0xf2, 0x44, 0xda, 0xfc, 0x8e, 0x52, 0xa8, 0x27, 0x22, 0xe0, 0x86, 0x36, 0xd0, 0x2c, 0x62, 0xe3,
	0x99, 0xfe, 0x07, 0x88, 0xdd, 0x25, 0x77, 0xfc, 0x88, 0xf1, 0x7b, 0xa3, 0x86, 0x9d, 0x6f, 0x99,
	0x32, 0xc9, 0x04, 0xda, 0x87, 0x36, 0xb6, 0xa3, 0x34, 0x34, 0x08, 0x36, 0x5b, 0x59, 0x7d, 0x9d,
	0x86, 0xe6, 0x93, 0x06, 0x9d, 0xf2, 0x76, 0x19, 0x53, 0x0b, 0x9a, 0xf9, 0x76, 0x5c, 0xa0, 0x8f,
	0x7e, 0x0e, 0x0b, 0xe0, 0xe1, 0x14, 0x75, 0x7b, 0xd5, 0xa7, 0x16, 0x34, 0xb2, 0x78, 0xd2, 0xa8,
	0x0d, 0x88, 0xa5, 0x8f, 0x68, 0x69, 0xcc, 0x6e, 0x9b, 0x44, 0x0b, 0x61, 0xe7, 0x86, 0x57, 0x78,
	0x64, 0x1b, 0x5e, 0x7d, 0x03, 0x8f, 0xfe, 0x86, 0x86, 0x12, 0xca, 0x0d, 0x8c, 0x06, 0xea, 0x79,
	0x61, 0x9e, 0x80, 0x7e, 0xc1, 0xd8, 0xfa, 0x1f, 0x59, 0x83, 0x68, 0x1f, 0x80, 0x98, 0x0f, 0x1a,
	0x74, 0xca, 0xc9, 0x9d, 0xd2, 0x1e, 0x42, 0x4b, 0xa6, 0x9e, 0xc7, 0x39, 0xdb, 0x92, 0xb7, 0xb0,
	0xd0, 0x03, 0xa8, 0x2f, 0x5c, 0x3f, 0x30, 0xc8, 0xbb, 0x56, 0xec, 0x9b, 0xa7, 0xd0, 0x9d, 0xc5,
	0xcc, 0x55, 0xfc, 0x13, 0x61, 0x1e, 0x35, 0xf8, 0xb1, 0x31, 0xfc, 0x05, 0xf2, 0x58, 0xa0, 0x5f,
	0xf1, 0x60, 0x1d, 0xa6, 0x0f, 0x6d, 0x64, 0x75, 0x7c, 0x86, 0x79, 0x88, 0xdd, 0xc2, 0x7a, 0xc2,
	0xcc, 0x05, 0x74, 0x4a, 0xe7, 0x4e, 0xe4, 0xc6, 0x26, 0x39, 0x29, 0x29, 0x69, 0x85, 0x92, 0xac,
	0x88, 0x18, 0xd0, 0x31, 0x57, 0xd3, 0x74, 0x1e, 0xfa, 0xca, 0xe6, 0x9e, 0x48, 0x58, 0x06, 0xb6,
	0xef, 0x2f, 0xf2, 0x59, 0x83, 0xde, 0x9b, 0x35, 0x3b, 0xa5, 0x3a, 0x87, 0xae, 0xc4, 0x69, 0x27,
	0xc1, 0xf1, 0xe2, 0x59, 0xfd, 0xa9, 0x4c, 0x54, 0x6f, 0xff, 0x2e, 0x2b, 0xd5, 0xde, 0x9f, 0xd8,
	0x65, 0xef, 0xe6, 0x57, 0xb1, 0xf7, 0xac, 0x38, 0xcc, 0x9b, 0x78, 0x3a, 0x7e, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xcd, 0x1c, 0x32, 0x15, 0xdd, 0x04, 0x00, 0x00,
}
