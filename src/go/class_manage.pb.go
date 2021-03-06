// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/class_manage.proto

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

type MemberManageReq_ManageType int32

const (
	MemberManageReq_DELETE               MemberManageReq_ManageType = 0
	MemberManageReq_SET_ADMINISTRATOR    MemberManageReq_ManageType = 1
	MemberManageReq_CANCEL_ADMINISTRATOR MemberManageReq_ManageType = 2
)

var MemberManageReq_ManageType_name = map[int32]string{
	0: "DELETE",
	1: "SET_ADMINISTRATOR",
	2: "CANCEL_ADMINISTRATOR",
}

var MemberManageReq_ManageType_value = map[string]int32{
	"DELETE":               0,
	"SET_ADMINISTRATOR":    1,
	"CANCEL_ADMINISTRATOR": 2,
}

func (x MemberManageReq_ManageType) String() string {
	return proto.EnumName(MemberManageReq_ManageType_name, int32(x))
}

func (MemberManageReq_ManageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1c8e358db319373f, []int{8, 0}
}

// 批量获取班级信息(只包括基础的信息，班级名称)
type GetClassesReq struct {
	PageIndex            int64    `protobuf:"varint,1,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageNum              int64    `protobuf:"varint,2,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClassesReq) Reset()         { *m = GetClassesReq{} }
func (m *GetClassesReq) String() string { return proto.CompactTextString(m) }
func (*GetClassesReq) ProtoMessage()    {}
func (*GetClassesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8e358db319373f, []int{0}
}

func (m *GetClassesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClassesReq.Unmarshal(m, b)
}
func (m *GetClassesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClassesReq.Marshal(b, m, deterministic)
}
func (m *GetClassesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClassesReq.Merge(m, src)
}
func (m *GetClassesReq) XXX_Size() int {
	return xxx_messageInfo_GetClassesReq.Size(m)
}
func (m *GetClassesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClassesReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetClassesReq proto.InternalMessageInfo

func (m *GetClassesReq) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *GetClassesReq) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

type GetClassesResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Classes              []*Class `protobuf:"bytes,2,rep,name=classes,proto3" json:"classes,omitempty"`
	PageIndex            int64    `protobuf:"varint,3,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageNum              int64    `protobuf:"varint,4,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	Total                int64    `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClassesResp) Reset()         { *m = GetClassesResp{} }
func (m *GetClassesResp) String() string { return proto.CompactTextString(m) }
func (*GetClassesResp) ProtoMessage()    {}
func (*GetClassesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8e358db319373f, []int{1}
}

func (m *GetClassesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClassesResp.Unmarshal(m, b)
}
func (m *GetClassesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClassesResp.Marshal(b, m, deterministic)
}
func (m *GetClassesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClassesResp.Merge(m, src)
}
func (m *GetClassesResp) XXX_Size() int {
	return xxx_messageInfo_GetClassesResp.Size(m)
}
func (m *GetClassesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClassesResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetClassesResp proto.InternalMessageInfo

func (m *GetClassesResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetClassesResp) GetClasses() []*Class {
	if m != nil {
		return m.Classes
	}
	return nil
}

func (m *GetClassesResp) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *GetClassesResp) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *GetClassesResp) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

// 根据ID获得班级具体信息
type GetClassByIDReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClassByIDReq) Reset()         { *m = GetClassByIDReq{} }
func (m *GetClassByIDReq) String() string { return proto.CompactTextString(m) }
func (*GetClassByIDReq) ProtoMessage()    {}
func (*GetClassByIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8e358db319373f, []int{2}
}

func (m *GetClassByIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClassByIDReq.Unmarshal(m, b)
}
func (m *GetClassByIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClassByIDReq.Marshal(b, m, deterministic)
}
func (m *GetClassByIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClassByIDReq.Merge(m, src)
}
func (m *GetClassByIDReq) XXX_Size() int {
	return xxx_messageInfo_GetClassByIDReq.Size(m)
}
func (m *GetClassByIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClassByIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetClassByIDReq proto.InternalMessageInfo

func (m *GetClassByIDReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetClassByIDResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Class                *Class   `protobuf:"bytes,2,opt,name=class,proto3" json:"class,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClassByIDResp) Reset()         { *m = GetClassByIDResp{} }
func (m *GetClassByIDResp) String() string { return proto.CompactTextString(m) }
func (*GetClassByIDResp) ProtoMessage()    {}
func (*GetClassByIDResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8e358db319373f, []int{3}
}

func (m *GetClassByIDResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClassByIDResp.Unmarshal(m, b)
}
func (m *GetClassByIDResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClassByIDResp.Marshal(b, m, deterministic)
}
func (m *GetClassByIDResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClassByIDResp.Merge(m, src)
}
func (m *GetClassByIDResp) XXX_Size() int {
	return xxx_messageInfo_GetClassByIDResp.Size(m)
}
func (m *GetClassByIDResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClassByIDResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetClassByIDResp proto.InternalMessageInfo

func (m *GetClassByIDResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetClassByIDResp) GetClass() *Class {
	if m != nil {
		return m.Class
	}
	return nil
}

// 新增班级
type AddClassReq struct {
	Class                *Class   `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddClassReq) Reset()         { *m = AddClassReq{} }
func (m *AddClassReq) String() string { return proto.CompactTextString(m) }
func (*AddClassReq) ProtoMessage()    {}
func (*AddClassReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8e358db319373f, []int{4}
}

func (m *AddClassReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddClassReq.Unmarshal(m, b)
}
func (m *AddClassReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddClassReq.Marshal(b, m, deterministic)
}
func (m *AddClassReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddClassReq.Merge(m, src)
}
func (m *AddClassReq) XXX_Size() int {
	return xxx_messageInfo_AddClassReq.Size(m)
}
func (m *AddClassReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddClassReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddClassReq proto.InternalMessageInfo

func (m *AddClassReq) GetClass() *Class {
	if m != nil {
		return m.Class
	}
	return nil
}

type AddClassResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IsSuccess            bool     `protobuf:"varint,2,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddClassResp) Reset()         { *m = AddClassResp{} }
func (m *AddClassResp) String() string { return proto.CompactTextString(m) }
func (*AddClassResp) ProtoMessage()    {}
func (*AddClassResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8e358db319373f, []int{5}
}

func (m *AddClassResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddClassResp.Unmarshal(m, b)
}
func (m *AddClassResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddClassResp.Marshal(b, m, deterministic)
}
func (m *AddClassResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddClassResp.Merge(m, src)
}
func (m *AddClassResp) XXX_Size() int {
	return xxx_messageInfo_AddClassResp.Size(m)
}
func (m *AddClassResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddClassResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddClassResp proto.InternalMessageInfo

func (m *AddClassResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AddClassResp) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

// 修改班级信息
type EditClassReq struct {
	Class                *Class   `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditClassReq) Reset()         { *m = EditClassReq{} }
func (m *EditClassReq) String() string { return proto.CompactTextString(m) }
func (*EditClassReq) ProtoMessage()    {}
func (*EditClassReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8e358db319373f, []int{6}
}

func (m *EditClassReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditClassReq.Unmarshal(m, b)
}
func (m *EditClassReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditClassReq.Marshal(b, m, deterministic)
}
func (m *EditClassReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditClassReq.Merge(m, src)
}
func (m *EditClassReq) XXX_Size() int {
	return xxx_messageInfo_EditClassReq.Size(m)
}
func (m *EditClassReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EditClassReq.DiscardUnknown(m)
}

var xxx_messageInfo_EditClassReq proto.InternalMessageInfo

func (m *EditClassReq) GetClass() *Class {
	if m != nil {
		return m.Class
	}
	return nil
}

type EditClassResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IsSuccess            bool     `protobuf:"varint,2,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditClassResp) Reset()         { *m = EditClassResp{} }
func (m *EditClassResp) String() string { return proto.CompactTextString(m) }
func (*EditClassResp) ProtoMessage()    {}
func (*EditClassResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8e358db319373f, []int{7}
}

func (m *EditClassResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditClassResp.Unmarshal(m, b)
}
func (m *EditClassResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditClassResp.Marshal(b, m, deterministic)
}
func (m *EditClassResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditClassResp.Merge(m, src)
}
func (m *EditClassResp) XXX_Size() int {
	return xxx_messageInfo_EditClassResp.Size(m)
}
func (m *EditClassResp) XXX_DiscardUnknown() {
	xxx_messageInfo_EditClassResp.DiscardUnknown(m)
}

var xxx_messageInfo_EditClassResp proto.InternalMessageInfo

func (m *EditClassResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *EditClassResp) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

// 小组成员管理
type MemberManageReq struct {
	ManageType           MemberManageReq_ManageType `protobuf:"varint,1,opt,name=manage_type,json=manageType,proto3,enum=protocol.MemberManageReq_ManageType" json:"manage_type,omitempty"`
	ClassId              int64                      `protobuf:"varint,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	MemberId             int64                      `protobuf:"varint,3,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MemberManageReq) Reset()         { *m = MemberManageReq{} }
func (m *MemberManageReq) String() string { return proto.CompactTextString(m) }
func (*MemberManageReq) ProtoMessage()    {}
func (*MemberManageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8e358db319373f, []int{8}
}

func (m *MemberManageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberManageReq.Unmarshal(m, b)
}
func (m *MemberManageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberManageReq.Marshal(b, m, deterministic)
}
func (m *MemberManageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberManageReq.Merge(m, src)
}
func (m *MemberManageReq) XXX_Size() int {
	return xxx_messageInfo_MemberManageReq.Size(m)
}
func (m *MemberManageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberManageReq.DiscardUnknown(m)
}

var xxx_messageInfo_MemberManageReq proto.InternalMessageInfo

func (m *MemberManageReq) GetManageType() MemberManageReq_ManageType {
	if m != nil {
		return m.ManageType
	}
	return MemberManageReq_DELETE
}

func (m *MemberManageReq) GetClassId() int64 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *MemberManageReq) GetMemberId() int64 {
	if m != nil {
		return m.MemberId
	}
	return 0
}

type MemberManageResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IsSuccess            bool     `protobuf:"varint,2,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberManageResp) Reset()         { *m = MemberManageResp{} }
func (m *MemberManageResp) String() string { return proto.CompactTextString(m) }
func (*MemberManageResp) ProtoMessage()    {}
func (*MemberManageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8e358db319373f, []int{9}
}

func (m *MemberManageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberManageResp.Unmarshal(m, b)
}
func (m *MemberManageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberManageResp.Marshal(b, m, deterministic)
}
func (m *MemberManageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberManageResp.Merge(m, src)
}
func (m *MemberManageResp) XXX_Size() int {
	return xxx_messageInfo_MemberManageResp.Size(m)
}
func (m *MemberManageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberManageResp.DiscardUnknown(m)
}

var xxx_messageInfo_MemberManageResp proto.InternalMessageInfo

func (m *MemberManageResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *MemberManageResp) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func init() {
	proto.RegisterEnum("protocol.MemberManageReq_ManageType", MemberManageReq_ManageType_name, MemberManageReq_ManageType_value)
	proto.RegisterType((*GetClassesReq)(nil), "protocol.GetClassesReq")
	proto.RegisterType((*GetClassesResp)(nil), "protocol.GetClassesResp")
	proto.RegisterType((*GetClassByIDReq)(nil), "protocol.GetClassByIDReq")
	proto.RegisterType((*GetClassByIDResp)(nil), "protocol.GetClassByIDResp")
	proto.RegisterType((*AddClassReq)(nil), "protocol.AddClassReq")
	proto.RegisterType((*AddClassResp)(nil), "protocol.AddClassResp")
	proto.RegisterType((*EditClassReq)(nil), "protocol.EditClassReq")
	proto.RegisterType((*EditClassResp)(nil), "protocol.EditClassResp")
	proto.RegisterType((*MemberManageReq)(nil), "protocol.MemberManageReq")
	proto.RegisterType((*MemberManageResp)(nil), "protocol.MemberManageResp")
}

func init() { proto.RegisterFile("proto/class_manage.proto", fileDescriptor_1c8e358db319373f) }

var fileDescriptor_1c8e358db319373f = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x61, 0x8b, 0xd3, 0x30,
	0x18, 0xc7, 0x6d, 0xe7, 0x76, 0xdb, 0xd3, 0xdb, 0x56, 0xc3, 0x09, 0x55, 0x39, 0x38, 0x83, 0xc2,
	0x7c, 0x33, 0x61, 0xea, 0x2b, 0x5f, 0xed, 0xb6, 0x22, 0x81, 0xdb, 0x84, 0xac, 0xa0, 0xe8, 0x8b,
	0xd2, 0x6b, 0xc2, 0x28, 0x2c, 0x6b, 0xef, 0x92, 0x82, 0xfb, 0x60, 0x7e, 0x1a, 0xbf, 0x8c, 0x24,
	0x69, 0xed, 0x6d, 0xc8, 0xc1, 0x60, 0xaf, 0x9a, 0xfc, 0xff, 0xcf, 0xf3, 0xcf, 0x2f, 0xe1, 0x29,
	0x04, 0xc5, 0x7d, 0xae, 0xf2, 0xf7, 0xe9, 0x26, 0x91, 0x32, 0x16, 0xc9, 0x36, 0x59, 0xf3, 0xb1,
	0x91, 0x50, 0xd7, 0x7c, 0xd2, 0x7c, 0xf3, 0x12, 0x55, 0x35, 0xb9, 0x10, 0xf9, 0xd6, 0xba, 0xb5,
	0x26, 0x55, 0xa2, 0x4a, 0x69, 0x35, 0x4c, 0xa0, 0xff, 0x85, 0xab, 0x99, 0x8e, 0xe2, 0x92, 0xf2,
	0x3b, 0x74, 0x09, 0x50, 0x24, 0x6b, 0x1e, 0x67, 0x5b, 0xc6, 0x7f, 0x05, 0xce, 0x95, 0x33, 0x6a,
	0xd1, 0x9e, 0x56, 0x88, 0x16, 0xd0, 0x0b, 0xe8, 0x1a, 0x7b, 0x5b, 0x8a, 0xc0, 0x35, 0xe6, 0x99,
	0xde, 0x2f, 0x4b, 0x81, 0x7f, 0x3b, 0x30, 0x78, 0x98, 0x25, 0x0b, 0x34, 0x82, 0x8e, 0x3d, 0xcd,
	0x04, 0x79, 0x13, 0x7f, 0x5c, 0x03, 0x8e, 0x57, 0x46, 0xa7, 0x95, 0x8f, 0xde, 0xc1, 0x59, 0x6a,
	0x1b, 0x03, 0xf7, 0xaa, 0x35, 0xf2, 0x26, 0xc3, 0xa6, 0xd4, 0x24, 0xd2, 0xda, 0x3f, 0x20, 0x6c,
	0x3d, 0x46, 0xf8, 0x74, 0x8f, 0x10, 0x5d, 0x40, 0x5b, 0xe5, 0x2a, 0xd9, 0x04, 0x6d, 0xa3, 0xdb,
	0x0d, 0x7e, 0x0d, 0xc3, 0x1a, 0xfb, 0x7a, 0x47, 0xe6, 0xfa, 0x11, 0x06, 0xe0, 0x66, 0xac, 0xba,
	0xbc, 0x9b, 0x31, 0x9c, 0x82, 0xbf, 0x5f, 0x72, 0xd4, 0xdd, 0xde, 0x42, 0xdb, 0xb0, 0x9b, 0x07,
	0xfb, 0xcf, 0xcd, 0xac, 0x8b, 0x3f, 0x82, 0x37, 0x65, 0xcc, 0x4a, 0xfc, 0xae, 0xe9, 0x72, 0x1e,
	0xed, 0xfa, 0x06, 0xe7, 0x4d, 0xd7, 0x51, 0x58, 0x97, 0x00, 0x99, 0x8c, 0x65, 0x99, 0xa6, 0xbc,
	0x62, 0xeb, 0xd2, 0x5e, 0x26, 0x57, 0x56, 0xc0, 0x9f, 0xe0, 0x3c, 0x64, 0x99, 0x3a, 0x96, 0xe7,
	0x3b, 0xf4, 0x1f, 0xb4, 0x9d, 0x12, 0xe8, 0x8f, 0x03, 0xc3, 0x05, 0x17, 0xb7, 0xfc, 0x7e, 0x61,
	0x66, 0x5e, 0x43, 0x85, 0xe0, 0xd9, 0x1f, 0x20, 0x56, 0xbb, 0x82, 0x9b, 0x13, 0x06, 0x93, 0x37,
	0xcd, 0x09, 0x07, 0xf5, 0x63, 0xbb, 0x8a, 0x76, 0x05, 0xa7, 0x20, 0xfe, 0xad, 0xf5, 0xcc, 0xd8,
	0xbf, 0x29, 0x63, 0xf5, 0x54, 0x9b, 0x3d, 0x61, 0xe8, 0x15, 0xf4, 0x84, 0x09, 0xd1, 0x9e, 0x1d,
	0xb6, 0xae, 0x15, 0x08, 0xc3, 0x04, 0xa0, 0x49, 0x44, 0x00, 0x9d, 0x79, 0x78, 0x13, 0x46, 0xa1,
	0xff, 0x04, 0x3d, 0x87, 0x67, 0xab, 0x30, 0x8a, 0xa7, 0xf3, 0x05, 0x59, 0x92, 0x55, 0x44, 0xa7,
	0xd1, 0x57, 0xea, 0x3b, 0x28, 0x80, 0x8b, 0xd9, 0x74, 0x39, 0x0b, 0x6f, 0x0e, 0x1c, 0x17, 0xff,
	0x04, 0x7f, 0x1f, 0xf6, 0x84, 0x4f, 0x77, 0xdd, 0xff, 0xe1, 0xad, 0xf3, 0xcf, 0x75, 0xf3, 0x6d,
	0xc7, 0xac, 0x3e, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xca, 0x02, 0x3d, 0x49, 0x04, 0x00,
	0x00,
}
