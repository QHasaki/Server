// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/class_enter.proto

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

// 搜索班级
type SearchClassesReq struct {
	PageIndex            int64    `protobuf:"varint,1,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageNum              int64    `protobuf:"varint,2,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	Keyword              string   `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchClassesReq) Reset()         { *m = SearchClassesReq{} }
func (m *SearchClassesReq) String() string { return proto.CompactTextString(m) }
func (*SearchClassesReq) ProtoMessage()    {}
func (*SearchClassesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_064de64e14ee3580, []int{0}
}

func (m *SearchClassesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchClassesReq.Unmarshal(m, b)
}
func (m *SearchClassesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchClassesReq.Marshal(b, m, deterministic)
}
func (m *SearchClassesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchClassesReq.Merge(m, src)
}
func (m *SearchClassesReq) XXX_Size() int {
	return xxx_messageInfo_SearchClassesReq.Size(m)
}
func (m *SearchClassesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchClassesReq.DiscardUnknown(m)
}

var xxx_messageInfo_SearchClassesReq proto.InternalMessageInfo

func (m *SearchClassesReq) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *SearchClassesReq) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *SearchClassesReq) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type SearchClassesResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Classes              []*Class `protobuf:"bytes,2,rep,name=classes,proto3" json:"classes,omitempty"`
	PageIndex            int64    `protobuf:"varint,3,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageNum              int64    `protobuf:"varint,4,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	Total                int64    `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchClassesResp) Reset()         { *m = SearchClassesResp{} }
func (m *SearchClassesResp) String() string { return proto.CompactTextString(m) }
func (*SearchClassesResp) ProtoMessage()    {}
func (*SearchClassesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_064de64e14ee3580, []int{1}
}

func (m *SearchClassesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchClassesResp.Unmarshal(m, b)
}
func (m *SearchClassesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchClassesResp.Marshal(b, m, deterministic)
}
func (m *SearchClassesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchClassesResp.Merge(m, src)
}
func (m *SearchClassesResp) XXX_Size() int {
	return xxx_messageInfo_SearchClassesResp.Size(m)
}
func (m *SearchClassesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchClassesResp.DiscardUnknown(m)
}

var xxx_messageInfo_SearchClassesResp proto.InternalMessageInfo

func (m *SearchClassesResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SearchClassesResp) GetClasses() []*Class {
	if m != nil {
		return m.Classes
	}
	return nil
}

func (m *SearchClassesResp) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *SearchClassesResp) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *SearchClassesResp) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

// 申请进入班级
type EnterClassReq struct {
	ClassId              int64    `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnterClassReq) Reset()         { *m = EnterClassReq{} }
func (m *EnterClassReq) String() string { return proto.CompactTextString(m) }
func (*EnterClassReq) ProtoMessage()    {}
func (*EnterClassReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_064de64e14ee3580, []int{2}
}

func (m *EnterClassReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterClassReq.Unmarshal(m, b)
}
func (m *EnterClassReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterClassReq.Marshal(b, m, deterministic)
}
func (m *EnterClassReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterClassReq.Merge(m, src)
}
func (m *EnterClassReq) XXX_Size() int {
	return xxx_messageInfo_EnterClassReq.Size(m)
}
func (m *EnterClassReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterClassReq.DiscardUnknown(m)
}

var xxx_messageInfo_EnterClassReq proto.InternalMessageInfo

func (m *EnterClassReq) GetClassId() int64 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

type EnterClassResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IsSuccess            bool     `protobuf:"varint,2,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnterClassResp) Reset()         { *m = EnterClassResp{} }
func (m *EnterClassResp) String() string { return proto.CompactTextString(m) }
func (*EnterClassResp) ProtoMessage()    {}
func (*EnterClassResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_064de64e14ee3580, []int{3}
}

func (m *EnterClassResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterClassResp.Unmarshal(m, b)
}
func (m *EnterClassResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterClassResp.Marshal(b, m, deterministic)
}
func (m *EnterClassResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterClassResp.Merge(m, src)
}
func (m *EnterClassResp) XXX_Size() int {
	return xxx_messageInfo_EnterClassResp.Size(m)
}
func (m *EnterClassResp) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterClassResp.DiscardUnknown(m)
}

var xxx_messageInfo_EnterClassResp proto.InternalMessageInfo

func (m *EnterClassResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *EnterClassResp) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

// 申请退出班级
type QuitClassReq struct {
	ClassId              int64    `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuitClassReq) Reset()         { *m = QuitClassReq{} }
func (m *QuitClassReq) String() string { return proto.CompactTextString(m) }
func (*QuitClassReq) ProtoMessage()    {}
func (*QuitClassReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_064de64e14ee3580, []int{4}
}

func (m *QuitClassReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuitClassReq.Unmarshal(m, b)
}
func (m *QuitClassReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuitClassReq.Marshal(b, m, deterministic)
}
func (m *QuitClassReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuitClassReq.Merge(m, src)
}
func (m *QuitClassReq) XXX_Size() int {
	return xxx_messageInfo_QuitClassReq.Size(m)
}
func (m *QuitClassReq) XXX_DiscardUnknown() {
	xxx_messageInfo_QuitClassReq.DiscardUnknown(m)
}

var xxx_messageInfo_QuitClassReq proto.InternalMessageInfo

func (m *QuitClassReq) GetClassId() int64 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

type QuitClassResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IsSuccess            bool     `protobuf:"varint,2,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuitClassResp) Reset()         { *m = QuitClassResp{} }
func (m *QuitClassResp) String() string { return proto.CompactTextString(m) }
func (*QuitClassResp) ProtoMessage()    {}
func (*QuitClassResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_064de64e14ee3580, []int{5}
}

func (m *QuitClassResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuitClassResp.Unmarshal(m, b)
}
func (m *QuitClassResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuitClassResp.Marshal(b, m, deterministic)
}
func (m *QuitClassResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuitClassResp.Merge(m, src)
}
func (m *QuitClassResp) XXX_Size() int {
	return xxx_messageInfo_QuitClassResp.Size(m)
}
func (m *QuitClassResp) XXX_DiscardUnknown() {
	xxx_messageInfo_QuitClassResp.DiscardUnknown(m)
}

var xxx_messageInfo_QuitClassResp proto.InternalMessageInfo

func (m *QuitClassResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *QuitClassResp) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

// 获取班级成员
// 包括了查看未批准的进入班级的请求
// 当学生申请进入班级的时候，就会在user——class表中插入一条记录，对于需要进行教师认证的班级，其is-check字段为false
// 这个请求返回user-class表中所有的记录，包括通过的和未通过的，is-check字段在ClassMember中
type GetMemberReq struct {
	ClassId              int64    `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	PageIndex            int64    `protobuf:"varint,2,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageNum              int64    `protobuf:"varint,3,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMemberReq) Reset()         { *m = GetMemberReq{} }
func (m *GetMemberReq) String() string { return proto.CompactTextString(m) }
func (*GetMemberReq) ProtoMessage()    {}
func (*GetMemberReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_064de64e14ee3580, []int{6}
}

func (m *GetMemberReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemberReq.Unmarshal(m, b)
}
func (m *GetMemberReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemberReq.Marshal(b, m, deterministic)
}
func (m *GetMemberReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemberReq.Merge(m, src)
}
func (m *GetMemberReq) XXX_Size() int {
	return xxx_messageInfo_GetMemberReq.Size(m)
}
func (m *GetMemberReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemberReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemberReq proto.InternalMessageInfo

func (m *GetMemberReq) GetClassId() int64 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *GetMemberReq) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *GetMemberReq) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

type GetMemberResp struct {
	Status               *Status        `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Members              []*ClassMember `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	PageIndex            int64          `protobuf:"varint,3,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageNum              int64          `protobuf:"varint,4,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	Total                int64          `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetMemberResp) Reset()         { *m = GetMemberResp{} }
func (m *GetMemberResp) String() string { return proto.CompactTextString(m) }
func (*GetMemberResp) ProtoMessage()    {}
func (*GetMemberResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_064de64e14ee3580, []int{7}
}

func (m *GetMemberResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemberResp.Unmarshal(m, b)
}
func (m *GetMemberResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemberResp.Marshal(b, m, deterministic)
}
func (m *GetMemberResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemberResp.Merge(m, src)
}
func (m *GetMemberResp) XXX_Size() int {
	return xxx_messageInfo_GetMemberResp.Size(m)
}
func (m *GetMemberResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemberResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemberResp proto.InternalMessageInfo

func (m *GetMemberResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetMemberResp) GetMembers() []*ClassMember {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *GetMemberResp) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *GetMemberResp) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *GetMemberResp) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

// 教师批准/拒绝学生进入班级的申请
type ApplyEnterRequestReq struct {
	ClassId              int64    `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsApply              bool     `protobuf:"varint,3,opt,name=is_apply,json=isApply,proto3" json:"is_apply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyEnterRequestReq) Reset()         { *m = ApplyEnterRequestReq{} }
func (m *ApplyEnterRequestReq) String() string { return proto.CompactTextString(m) }
func (*ApplyEnterRequestReq) ProtoMessage()    {}
func (*ApplyEnterRequestReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_064de64e14ee3580, []int{8}
}

func (m *ApplyEnterRequestReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyEnterRequestReq.Unmarshal(m, b)
}
func (m *ApplyEnterRequestReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyEnterRequestReq.Marshal(b, m, deterministic)
}
func (m *ApplyEnterRequestReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyEnterRequestReq.Merge(m, src)
}
func (m *ApplyEnterRequestReq) XXX_Size() int {
	return xxx_messageInfo_ApplyEnterRequestReq.Size(m)
}
func (m *ApplyEnterRequestReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyEnterRequestReq.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyEnterRequestReq proto.InternalMessageInfo

func (m *ApplyEnterRequestReq) GetClassId() int64 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *ApplyEnterRequestReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ApplyEnterRequestReq) GetIsApply() bool {
	if m != nil {
		return m.IsApply
	}
	return false
}

type ApplyEnterRequestResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IsSuccess            bool     `protobuf:"varint,2,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyEnterRequestResp) Reset()         { *m = ApplyEnterRequestResp{} }
func (m *ApplyEnterRequestResp) String() string { return proto.CompactTextString(m) }
func (*ApplyEnterRequestResp) ProtoMessage()    {}
func (*ApplyEnterRequestResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_064de64e14ee3580, []int{9}
}

func (m *ApplyEnterRequestResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyEnterRequestResp.Unmarshal(m, b)
}
func (m *ApplyEnterRequestResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyEnterRequestResp.Marshal(b, m, deterministic)
}
func (m *ApplyEnterRequestResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyEnterRequestResp.Merge(m, src)
}
func (m *ApplyEnterRequestResp) XXX_Size() int {
	return xxx_messageInfo_ApplyEnterRequestResp.Size(m)
}
func (m *ApplyEnterRequestResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyEnterRequestResp.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyEnterRequestResp proto.InternalMessageInfo

func (m *ApplyEnterRequestResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ApplyEnterRequestResp) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func init() {
	proto.RegisterType((*SearchClassesReq)(nil), "protocol.SearchClassesReq")
	proto.RegisterType((*SearchClassesResp)(nil), "protocol.SearchClassesResp")
	proto.RegisterType((*EnterClassReq)(nil), "protocol.EnterClassReq")
	proto.RegisterType((*EnterClassResp)(nil), "protocol.EnterClassResp")
	proto.RegisterType((*QuitClassReq)(nil), "protocol.QuitClassReq")
	proto.RegisterType((*QuitClassResp)(nil), "protocol.QuitClassResp")
	proto.RegisterType((*GetMemberReq)(nil), "protocol.GetMemberReq")
	proto.RegisterType((*GetMemberResp)(nil), "protocol.GetMemberResp")
	proto.RegisterType((*ApplyEnterRequestReq)(nil), "protocol.ApplyEnterRequestReq")
	proto.RegisterType((*ApplyEnterRequestResp)(nil), "protocol.ApplyEnterRequestResp")
}

func init() { proto.RegisterFile("proto/class_enter.proto", fileDescriptor_064de64e14ee3580) }

var fileDescriptor_064de64e14ee3580 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x95, 0x71, 0xc1, 0x66, 0x80, 0x16, 0x5c, 0x10, 0x6e, 0x25, 0x24, 0xe4, 0x93, 0xe9, 0x01,
	0x24, 0x7a, 0xec, 0xa9, 0xad, 0xaa, 0x8a, 0x43, 0x2b, 0x75, 0xb9, 0xb4, 0xbd, 0x38, 0xc6, 0x9e,
	0x24, 0x56, 0xfc, 0x85, 0x77, 0xad, 0x84, 0x7f, 0x16, 0xe5, 0xd7, 0x45, 0xbb, 0xeb, 0x55, 0xc0,
	0x8a, 0x48, 0x90, 0x38, 0xd9, 0xf3, 0xe6, 0xe3, 0xbd, 0x7d, 0x33, 0x30, 0xce, 0x8b, 0x8c, 0x65,
	0x8b, 0x20, 0xf6, 0x29, 0xf5, 0x30, 0x65, 0x58, 0xcc, 0x05, 0x62, 0x99, 0xe2, 0x13, 0x64, 0xf1,
	0x47, 0xab, 0x2a, 0xc9, 0x92, 0x24, 0x4b, 0x65, 0x56, 0x61, 0x94, 0xf9, 0xac, 0xa4, 0x12, 0x73,
	0x2e, 0xa1, 0xbf, 0x46, 0xbf, 0x08, 0xae, 0xbf, 0xf3, 0x61, 0x48, 0x09, 0x6e, 0xad, 0x09, 0x40,
	0xee, 0x5f, 0xa1, 0x17, 0xa5, 0x21, 0xde, 0xd9, 0xda, 0x54, 0x73, 0x75, 0xd2, 0xe6, 0xc8, 0x8a,
	0x03, 0xd6, 0x07, 0x30, 0x45, 0x3a, 0x2d, 0x13, 0xbb, 0x21, 0x92, 0x06, 0x8f, 0x7f, 0x97, 0x89,
	0x65, 0x83, 0x71, 0x83, 0xbb, 0xdb, 0xac, 0x08, 0x6d, 0x7d, 0xaa, 0xb9, 0x6d, 0xa2, 0x42, 0xe7,
	0x5e, 0x83, 0x41, 0x8d, 0x88, 0xe6, 0x96, 0x0b, 0x2d, 0xa9, 0x46, 0xb0, 0x74, 0x96, 0xfd, 0xb9,
	0x7a, 0xc0, 0x7c, 0x2d, 0x70, 0x52, 0xe5, 0xad, 0x19, 0x18, 0x81, 0x6c, 0xb4, 0x1b, 0x53, 0xdd,
	0xed, 0x2c, 0xdf, 0x3d, 0x95, 0x8a, 0x89, 0x44, 0xe5, 0x6b, 0xf2, 0xf5, 0x63, 0xf2, 0xdf, 0x1c,
	0xca, 0x1f, 0x42, 0x93, 0x65, 0xcc, 0x8f, 0xed, 0xa6, 0xc0, 0x65, 0xe0, 0x7c, 0x82, 0xde, 0x0f,
	0xee, 0xb1, 0xa4, 0xc1, 0x2d, 0x9f, 0x20, 0xad, 0x8f, 0xc2, 0xca, 0x1d, 0xc9, 0xbd, 0x0a, 0x9d,
	0x7f, 0xf0, 0x76, 0xbf, 0xf6, 0xa4, 0x27, 0x4e, 0x00, 0x22, 0xea, 0xd1, 0x32, 0x08, 0x90, 0x52,
	0xe1, 0xac, 0x49, 0xda, 0x11, 0x5d, 0x4b, 0xc0, 0x99, 0x41, 0xf7, 0x4f, 0x19, 0xb1, 0xd7, 0xa8,
	0xf8, 0x0b, 0xbd, 0xbd, 0xd2, 0x73, 0x8a, 0x08, 0xa0, 0xfb, 0x13, 0xd9, 0x2f, 0x4c, 0x36, 0x58,
	0x1c, 0x17, 0x51, 0x5b, 0x43, 0xe3, 0xd8, 0x1a, 0xf4, 0x83, 0x35, 0x38, 0x0f, 0x1a, 0xf4, 0xf6,
	0x58, 0x4e, 0xd2, 0xbf, 0x00, 0x23, 0x11, 0x7d, 0xea, 0x4e, 0x46, 0xb5, 0x3b, 0xa9, 0xa6, 0xaa,
	0xaa, 0xb3, 0x5f, 0x0b, 0xc2, 0xf0, 0x6b, 0x9e, 0xc7, 0x3b, 0x71, 0x06, 0x04, 0xb7, 0x25, 0x52,
	0xf6, 0x82, 0x53, 0x63, 0x30, 0x4a, 0x8a, 0x05, 0xcf, 0x48, 0x9b, 0x5a, 0x3c, 0x5c, 0x85, 0xbc,
	0x27, 0xa2, 0x9e, 0xcf, 0xc7, 0x09, 0x65, 0x26, 0x31, 0x22, 0x2a, 0xa6, 0x3b, 0x17, 0x30, 0x7a,
	0x86, 0xe6, 0x8c, 0xab, 0xfe, 0xf6, 0xfe, 0xff, 0x40, 0x75, 0x7e, 0x51, 0x3f, 0x9b, 0x96, 0xf8,
	0xfb, 0xfc, 0x18, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x95, 0xd5, 0xf6, 0x82, 0x04, 0x00, 0x00,
}
