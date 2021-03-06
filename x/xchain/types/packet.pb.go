// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: xchain/packet.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type XchainPacketData struct {
	// Types that are valid to be assigned to Packet:
	//	*XchainPacketData_NoData
	//	*XchainPacketData_MsgMintRequestPacket
	Packet isXchainPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *XchainPacketData) Reset()         { *m = XchainPacketData{} }
func (m *XchainPacketData) String() string { return proto.CompactTextString(m) }
func (*XchainPacketData) ProtoMessage()    {}
func (*XchainPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7807c5f9069db2f1, []int{0}
}
func (m *XchainPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *XchainPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_XchainPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *XchainPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XchainPacketData.Merge(m, src)
}
func (m *XchainPacketData) XXX_Size() int {
	return m.Size()
}
func (m *XchainPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_XchainPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_XchainPacketData proto.InternalMessageInfo

type isXchainPacketData_Packet interface {
	isXchainPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type XchainPacketData_NoData struct {
	NoData *NoData `protobuf:"bytes,1,opt,name=noData,proto3,oneof" json:"noData,omitempty"`
}
type XchainPacketData_MsgMintRequestPacket struct {
	MsgMintRequestPacket *MsgMintRequestPacketData `protobuf:"bytes,2,opt,name=msgMintRequestPacket,proto3,oneof" json:"msgMintRequestPacket,omitempty"`
}

func (*XchainPacketData_NoData) isXchainPacketData_Packet()               {}
func (*XchainPacketData_MsgMintRequestPacket) isXchainPacketData_Packet() {}

func (m *XchainPacketData) GetPacket() isXchainPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *XchainPacketData) GetNoData() *NoData {
	if x, ok := m.GetPacket().(*XchainPacketData_NoData); ok {
		return x.NoData
	}
	return nil
}

func (m *XchainPacketData) GetMsgMintRequestPacket() *MsgMintRequestPacketData {
	if x, ok := m.GetPacket().(*XchainPacketData_MsgMintRequestPacket); ok {
		return x.MsgMintRequestPacket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*XchainPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*XchainPacketData_NoData)(nil),
		(*XchainPacketData_MsgMintRequestPacket)(nil),
	}
}

type NoData struct {
}

func (m *NoData) Reset()         { *m = NoData{} }
func (m *NoData) String() string { return proto.CompactTextString(m) }
func (*NoData) ProtoMessage()    {}
func (*NoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7807c5f9069db2f1, []int{1}
}
func (m *NoData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoData.Merge(m, src)
}
func (m *NoData) XXX_Size() int {
	return m.Size()
}
func (m *NoData) XXX_DiscardUnknown() {
	xxx_messageInfo_NoData.DiscardUnknown(m)
}

var xxx_messageInfo_NoData proto.InternalMessageInfo

// this line is used by starport scaffolding # ibc/packet/proto/message
// MsgMintRequestPacketData defines a struct for the packet payload
type MsgMintRequestPacketData struct {
	Reciever    string `protobuf:"bytes,1,opt,name=reciever,proto3" json:"reciever,omitempty"`
	Amount      uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee         uint64 `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
	TokenId     uint32 `protobuf:"varint,4,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	SrcChainId  uint32 `protobuf:"varint,5,opt,name=srcChainId,proto3" json:"srcChainId,omitempty"`
	DestChainId uint32 `protobuf:"varint,6,opt,name=destChainId,proto3" json:"destChainId,omitempty"`
}

func (m *MsgMintRequestPacketData) Reset()         { *m = MsgMintRequestPacketData{} }
func (m *MsgMintRequestPacketData) String() string { return proto.CompactTextString(m) }
func (*MsgMintRequestPacketData) ProtoMessage()    {}
func (*MsgMintRequestPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7807c5f9069db2f1, []int{2}
}
func (m *MsgMintRequestPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintRequestPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintRequestPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintRequestPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintRequestPacketData.Merge(m, src)
}
func (m *MsgMintRequestPacketData) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintRequestPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintRequestPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintRequestPacketData proto.InternalMessageInfo

func (m *MsgMintRequestPacketData) GetReciever() string {
	if m != nil {
		return m.Reciever
	}
	return ""
}

func (m *MsgMintRequestPacketData) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *MsgMintRequestPacketData) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *MsgMintRequestPacketData) GetTokenId() uint32 {
	if m != nil {
		return m.TokenId
	}
	return 0
}

func (m *MsgMintRequestPacketData) GetSrcChainId() uint32 {
	if m != nil {
		return m.SrcChainId
	}
	return 0
}

func (m *MsgMintRequestPacketData) GetDestChainId() uint32 {
	if m != nil {
		return m.DestChainId
	}
	return 0
}

// MsgMintRequestPacketAck defines a struct for the packet acknowledgment
type MsgMintRequestPacketAck struct {
	TxId string `protobuf:"bytes,1,opt,name=txId,proto3" json:"txId,omitempty"`
}

func (m *MsgMintRequestPacketAck) Reset()         { *m = MsgMintRequestPacketAck{} }
func (m *MsgMintRequestPacketAck) String() string { return proto.CompactTextString(m) }
func (*MsgMintRequestPacketAck) ProtoMessage()    {}
func (*MsgMintRequestPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_7807c5f9069db2f1, []int{3}
}
func (m *MsgMintRequestPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintRequestPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintRequestPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintRequestPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintRequestPacketAck.Merge(m, src)
}
func (m *MsgMintRequestPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintRequestPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintRequestPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintRequestPacketAck proto.InternalMessageInfo

func (m *MsgMintRequestPacketAck) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func init() {
	proto.RegisterType((*XchainPacketData)(nil), "rosenlabs.xchain.xchain.XchainPacketData")
	proto.RegisterType((*NoData)(nil), "rosenlabs.xchain.xchain.NoData")
	proto.RegisterType((*MsgMintRequestPacketData)(nil), "rosenlabs.xchain.xchain.MsgMintRequestPacketData")
	proto.RegisterType((*MsgMintRequestPacketAck)(nil), "rosenlabs.xchain.xchain.MsgMintRequestPacketAck")
}

func init() { proto.RegisterFile("xchain/packet.proto", fileDescriptor_7807c5f9069db2f1) }

var fileDescriptor_7807c5f9069db2f1 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x31, 0x4f, 0x32, 0x41,
	0x10, 0xbd, 0xfd, 0xe0, 0x3b, 0x71, 0x88, 0x09, 0x59, 0x8d, 0x6c, 0x2c, 0x56, 0x72, 0x15, 0x89,
	0xe1, 0x88, 0x5a, 0x59, 0x8a, 0x9a, 0x48, 0x81, 0x31, 0x57, 0x19, 0xbb, 0x63, 0x6f, 0x84, 0x0b,
	0xb2, 0x8b, 0xb7, 0x8b, 0xc1, 0x7f, 0xe1, 0x0f, 0xb2, 0xb0, 0xb4, 0xa4, 0xb4, 0x34, 0xf0, 0x47,
	0xcc, 0xed, 0x2d, 0x86, 0x02, 0xaa, 0x7d, 0x33, 0xfb, 0xe6, 0xe5, 0xbd, 0xc9, 0xc0, 0xfe, 0x4c,
	0x0c, 0xe3, 0x54, 0xb6, 0x27, 0xb1, 0x18, 0xa1, 0x09, 0x27, 0x99, 0x32, 0x8a, 0xd6, 0x33, 0xa5,
	0x51, 0x3e, 0xc7, 0x7d, 0x1d, 0x16, 0xdf, 0xee, 0x09, 0x3e, 0x09, 0xd4, 0x1e, 0x2c, 0xbc, 0xb7,
	0xfc, 0xeb, 0xd8, 0xc4, 0xf4, 0x02, 0x7c, 0xa9, 0x72, 0xc4, 0x48, 0x83, 0x34, 0xab, 0x67, 0xc7,
	0xe1, 0x96, 0xf1, 0xf0, 0xce, 0xd2, 0x6e, 0xbd, 0xc8, 0x0d, 0xd0, 0x01, 0x1c, 0x8c, 0xf5, 0xa0,
	0x97, 0x4a, 0x13, 0xe1, 0xcb, 0x14, 0xb5, 0x29, 0x64, 0xd9, 0x3f, 0x2b, 0x74, 0xba, 0x55, 0xa8,
	0xb7, 0x61, 0xc8, 0x49, 0x6f, 0x14, 0xec, 0x54, 0xc0, 0x2f, 0x12, 0x06, 0x15, 0xf0, 0x0b, 0x1b,
	0xc1, 0x07, 0x01, 0xb6, 0x4d, 0x88, 0x1e, 0x41, 0x25, 0x43, 0x91, 0xe2, 0x2b, 0x66, 0x36, 0xd6,
	0x6e, 0xf4, 0x57, 0xd3, 0x43, 0xf0, 0xe3, 0xb1, 0x9a, 0xca, 0xc2, 0x67, 0x39, 0x72, 0x15, 0xad,
	0x41, 0xe9, 0x09, 0x91, 0x95, 0x6c, 0x33, 0x87, 0x94, 0xc1, 0x8e, 0x51, 0x23, 0x94, 0xdd, 0x84,
	0x95, 0x1b, 0xa4, 0xb9, 0x17, 0xad, 0x4a, 0xca, 0x01, 0x74, 0x26, 0xae, 0xf2, 0x34, 0xdd, 0x84,
	0xfd, 0xb7, 0x9f, 0x6b, 0x1d, 0xda, 0x80, 0x6a, 0x82, 0xda, 0xac, 0x08, 0xbe, 0x25, 0xac, 0xb7,
	0x82, 0x16, 0xd4, 0x37, 0xb9, 0xbf, 0x14, 0x23, 0x4a, 0xa1, 0x6c, 0x66, 0xdd, 0xc4, 0x19, 0xb7,
	0xb8, 0x73, 0xf3, 0xb5, 0xe0, 0x64, 0xbe, 0xe0, 0xe4, 0x67, 0xc1, 0xc9, 0xfb, 0x92, 0x7b, 0xf3,
	0x25, 0xf7, 0xbe, 0x97, 0xdc, 0x7b, 0x3c, 0x19, 0xa4, 0x66, 0x38, 0xed, 0x87, 0x42, 0x8d, 0xdb,
	0x76, 0xe1, 0xad, 0x7c, 0xe3, 0x6d, 0x77, 0x18, 0xb3, 0x15, 0x30, 0x6f, 0x13, 0xd4, 0x7d, 0xdf,
	0x5e, 0xc8, 0xf9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xa5, 0xbc, 0x3c, 0x38, 0x02, 0x00,
	0x00,
}

func (m *XchainPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *XchainPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *XchainPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Packet != nil {
		{
			size := m.Packet.Size()
			i -= size
			if _, err := m.Packet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *XchainPacketData_NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *XchainPacketData_NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoData != nil {
		{
			size, err := m.NoData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *XchainPacketData_MsgMintRequestPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *XchainPacketData_MsgMintRequestPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MsgMintRequestPacket != nil {
		{
			size, err := m.MsgMintRequestPacket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *NoData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgMintRequestPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintRequestPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintRequestPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DestChainId != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.DestChainId))
		i--
		dAtA[i] = 0x30
	}
	if m.SrcChainId != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.SrcChainId))
		i--
		dAtA[i] = 0x28
	}
	if m.TokenId != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.TokenId))
		i--
		dAtA[i] = 0x20
	}
	if m.Fee != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Fee))
		i--
		dAtA[i] = 0x18
	}
	if m.Amount != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Reciever) > 0 {
		i -= len(m.Reciever)
		copy(dAtA[i:], m.Reciever)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Reciever)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMintRequestPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintRequestPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintRequestPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxId) > 0 {
		i -= len(m.TxId)
		copy(dAtA[i:], m.TxId)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.TxId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *XchainPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		n += m.Packet.Size()
	}
	return n
}

func (m *XchainPacketData_NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoData != nil {
		l = m.NoData.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *XchainPacketData_MsgMintRequestPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MsgMintRequestPacket != nil {
		l = m.MsgMintRequestPacket.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgMintRequestPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Reciever)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovPacket(uint64(m.Amount))
	}
	if m.Fee != 0 {
		n += 1 + sovPacket(uint64(m.Fee))
	}
	if m.TokenId != 0 {
		n += 1 + sovPacket(uint64(m.TokenId))
	}
	if m.SrcChainId != 0 {
		n += 1 + sovPacket(uint64(m.SrcChainId))
	}
	if m.DestChainId != 0 {
		n += 1 + sovPacket(uint64(m.DestChainId))
	}
	return n
}

func (m *MsgMintRequestPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxId)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *XchainPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: XchainPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: XchainPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NoData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &XchainPacketData_NoData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgMintRequestPacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MsgMintRequestPacketData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &XchainPacketData_MsgMintRequestPacket{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NoData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NoData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgMintRequestPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgMintRequestPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintRequestPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reciever", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reciever = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			m.Fee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			m.TokenId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcChainId", wireType)
			}
			m.SrcChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SrcChainId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestChainId", wireType)
			}
			m.DestChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DestChainId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgMintRequestPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgMintRequestPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintRequestPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
