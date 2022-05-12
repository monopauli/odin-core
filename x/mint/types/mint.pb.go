// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mint/mint.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type MintPool struct {
	// treasury pool
	TreasuryPool github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=treasury_pool,json=treasuryPool,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"treasury_pool" yaml:"treasury_pool"`
}

func (m *MintPool) Reset()         { *m = MintPool{} }
func (m *MintPool) String() string { return proto.CompactTextString(m) }
func (*MintPool) ProtoMessage()    {}
func (*MintPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1b9fbb701b2a577, []int{0}
}
func (m *MintPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintPool.Merge(m, src)
}
func (m *MintPool) XXX_Size() int {
	return m.Size()
}
func (m *MintPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MintPool.DiscardUnknown(m)
}

var xxx_messageInfo_MintPool proto.InternalMessageInfo

func (m *MintPool) GetTreasuryPool() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TreasuryPool
	}
	return nil
}

// Minter represents the minting state.
type Minter struct {
	// current annual inflation rate
	Inflation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation"`
	// current annual expected provisions
	AnnualProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=annual_provisions,json=annualProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"annual_provisions" yaml:"annual_provisions"`
	// current mint volume
	CurrentMintVolume github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=current_mint_volume,json=currentMintVolume,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"current_mint_volume" yaml:"current_mint_volume"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1b9fbb701b2a577, []int{1}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

func (m *Minter) GetCurrentMintVolume() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.CurrentMintVolume
	}
	return nil
}

// AllowedDenom represents the allowed minting denom.
type AllowedDenom struct {
	// little version of coin
	LittleDenom string `protobuf:"bytes,1,opt,name=little_denom,json=littleDenom,proto3" json:"little_denom,omitempty" yaml:"little_denom"`
	// large version of coin
	LargeDenom string `protobuf:"bytes,2,opt,name=large_denom,json=largeDenom,proto3" json:"large_denom,omitempty" yaml:"large_denom"`
}

func (m *AllowedDenom) Reset()         { *m = AllowedDenom{} }
func (m *AllowedDenom) String() string { return proto.CompactTextString(m) }
func (*AllowedDenom) ProtoMessage()    {}
func (*AllowedDenom) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1b9fbb701b2a577, []int{2}
}
func (m *AllowedDenom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedDenom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedDenom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedDenom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedDenom.Merge(m, src)
}
func (m *AllowedDenom) XXX_Size() int {
	return m.Size()
}
func (m *AllowedDenom) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedDenom.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedDenom proto.InternalMessageInfo

func (m *AllowedDenom) GetLittleDenom() string {
	if m != nil {
		return m.LittleDenom
	}
	return ""
}

func (m *AllowedDenom) GetLargeDenom() string {
	if m != nil {
		return m.LargeDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*MintPool)(nil), "mint.MintPool")
	proto.RegisterType((*Minter)(nil), "mint.Minter")
	proto.RegisterType((*AllowedDenom)(nil), "mint.AllowedDenom")
}

func init() { proto.RegisterFile("mint/mint.proto", fileDescriptor_e1b9fbb701b2a577) }

var fileDescriptor_e1b9fbb701b2a577 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0x6d, 0x55, 0xb5, 0xd7, 0x20, 0xa8, 0x5b, 0x41, 0x9a, 0xc1, 0xae, 0x3c, 0xa0,
	0x2c, 0xf1, 0xa9, 0x30, 0x20, 0x65, 0x23, 0xcd, 0x40, 0x51, 0x69, 0x22, 0x0b, 0x31, 0xb0, 0x44,
	0x8e, 0x73, 0x35, 0x27, 0xce, 0xf7, 0x2c, 0xdf, 0x39, 0x25, 0xff, 0x01, 0x6c, 0x8c, 0x30, 0xd1,
	0x99, 0xbf, 0xa4, 0x63, 0x47, 0xc4, 0x60, 0x50, 0xb2, 0xb0, 0x92, 0xbf, 0x00, 0xdd, 0x9d, 0x4b,
	0xcb, 0x0f, 0x09, 0xba, 0x24, 0xf7, 0xde, 0xbb, 0xef, 0xe7, 0xbd, 0xf7, 0xb5, 0x8d, 0x6f, 0xa6,
	0x4c, 0x28, 0xa2, 0x7f, 0x82, 0x2c, 0x07, 0x05, 0xce, 0x8a, 0x3e, 0x37, 0xb7, 0x13, 0x48, 0xc0,
	0x24, 0x88, 0x3e, 0xd9, 0x5a, 0xd3, 0x4b, 0x00, 0x12, 0x4e, 0x89, 0x89, 0x46, 0xc5, 0x31, 0x51,
	0x2c, 0xa5, 0x52, 0x45, 0x69, 0x56, 0x5d, 0xd8, 0xf9, 0xfd, 0x42, 0x24, 0xa6, 0x55, 0xc9, 0x8d,
	0x41, 0xa6, 0x20, 0xc9, 0x28, 0x92, 0x94, 0x4c, 0xf6, 0x46, 0x54, 0x45, 0x7b, 0x24, 0x06, 0x26,
	0x6c, 0xdd, 0xff, 0x80, 0xf0, 0xda, 0x13, 0x26, 0xd4, 0x00, 0x80, 0x3b, 0xaf, 0x11, 0xbe, 0xa1,
	0x72, 0x1a, 0xc9, 0x22, 0x9f, 0x0e, 0x33, 0x00, 0xde, 0x40, 0xbb, 0xcb, 0xad, 0x8d, 0x7b, 0x3b,
	0x81, 0xa5, 0x04, 0x9a, 0x12, 0x54, 0x94, 0x60, 0x1f, 0x98, 0xe8, 0x3e, 0x3a, 0x2b, 0xbd, 0xda,
	0xa2, 0xf4, 0xb6, 0xa7, 0x51, 0xca, 0x3b, 0xfe, 0x2f, 0x6a, 0xff, 0xe3, 0x17, 0xaf, 0x95, 0x30,
	0xf5, 0xa2, 0x18, 0x05, 0x31, 0xa4, 0xa4, 0x1a, 0xc5, 0xfe, 0xb5, 0xe5, 0xf8, 0x25, 0x51, 0xd3,
	0x8c, 0x4a, 0x03, 0x92, 0x61, 0xfd, 0x42, 0xab, 0x47, 0xe9, 0xac, 0xbd, 0x3b, 0xf5, 0xd0, 0xb7,
	0x53, 0x0f, 0xf9, 0xdf, 0x97, 0xf0, 0xaa, 0x9e, 0x90, 0xe6, 0xce, 0x21, 0x5e, 0x67, 0xe2, 0x98,
	0x47, 0x8a, 0x81, 0x68, 0xa0, 0x5d, 0xd4, 0x5a, 0xef, 0x06, 0xba, 0xff, 0xe7, 0xd2, 0xbb, 0xfb,
	0x1f, 0x7d, 0x7a, 0x34, 0x0e, 0x2f, 0x01, 0xce, 0x09, 0xde, 0x8c, 0x84, 0x28, 0x22, 0x3e, 0xcc,
	0x72, 0x98, 0x30, 0xc9, 0x40, 0xc8, 0xc6, 0x92, 0xa1, 0x3e, 0xbe, 0x1e, 0x75, 0x51, 0x7a, 0x0d,
	0xbb, 0xff, 0x1f, 0x40, 0x3f, 0xbc, 0x65, 0x73, 0x83, 0x9f, 0x29, 0xe7, 0x3d, 0xc2, 0x5b, 0x71,
	0x91, 0xe7, 0x54, 0xa8, 0xa1, 0x7e, 0xec, 0xc3, 0x09, 0xf0, 0x22, 0xa5, 0x8d, 0xe5, 0x7f, 0x99,
	0x7d, 0x54, 0x99, 0xdd, 0xb4, 0xcd, 0xfe, 0xc2, 0xb8, 0x9e, 0xe5, 0x9b, 0x15, 0x41, 0xdb, 0xfb,
	0xcc, 0xe8, 0x3b, 0x2b, 0xc6, 0xf3, 0x37, 0x08, 0xd7, 0x1f, 0x72, 0x0e, 0x27, 0x74, 0xdc, 0xa3,
	0x02, 0x52, 0xa7, 0x83, 0xeb, 0x9c, 0x29, 0xc5, 0xe9, 0x70, 0xac, 0xe3, 0xca, 0xfc, 0x3b, 0x8b,
	0xd2, 0xdb, 0xb2, 0xb3, 0x5c, 0xad, 0xfa, 0xe1, 0x86, 0x0d, 0xad, 0xf6, 0x01, 0xde, 0xe0, 0x51,
	0x9e, 0x5c, 0x48, 0xad, 0xc3, 0xb7, 0x17, 0xa5, 0xe7, 0x54, 0xd2, 0xcb, 0xa2, 0x1f, 0x62, 0x13,
	0x19, 0xa1, 0x9d, 0xa5, 0x7b, 0x70, 0x36, 0x73, 0xd1, 0xf9, 0xcc, 0x45, 0x5f, 0x67, 0x2e, 0x7a,
	0x3b, 0x77, 0x6b, 0xe7, 0x73, 0xb7, 0xf6, 0x69, 0xee, 0xd6, 0x9e, 0x93, 0x2b, 0x8b, 0xf6, 0x7b,
	0x07, 0x47, 0xed, 0x41, 0xd8, 0x7f, 0xda, 0xdf, 0xef, 0x1f, 0x12, 0x18, 0x33, 0xd1, 0x8e, 0x21,
	0xa7, 0xe4, 0x95, 0xf9, 0xc8, 0xec, 0xd6, 0xa3, 0x55, 0xf3, 0xce, 0xdf, 0xff, 0x11, 0x00, 0x00,
	0xff, 0xff, 0xcd, 0x05, 0xb2, 0x3b, 0x7e, 0x03, 0x00, 0x00,
}

func (this *MintPool) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MintPool)
	if !ok {
		that2, ok := that.(MintPool)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.TreasuryPool) != len(that1.TreasuryPool) {
		return false
	}
	for i := range this.TreasuryPool {
		if !this.TreasuryPool[i].Equal(&that1.TreasuryPool[i]) {
			return false
		}
	}
	return true
}
func (this *Minter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Minter)
	if !ok {
		that2, ok := that.(Minter)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Inflation.Equal(that1.Inflation) {
		return false
	}
	if !this.AnnualProvisions.Equal(that1.AnnualProvisions) {
		return false
	}
	if len(this.CurrentMintVolume) != len(that1.CurrentMintVolume) {
		return false
	}
	for i := range this.CurrentMintVolume {
		if !this.CurrentMintVolume[i].Equal(&that1.CurrentMintVolume[i]) {
			return false
		}
	}
	return true
}
func (this *AllowedDenom) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AllowedDenom)
	if !ok {
		that2, ok := that.(AllowedDenom)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.LittleDenom != that1.LittleDenom {
		return false
	}
	if this.LargeDenom != that1.LargeDenom {
		return false
	}
	return true
}
func (m *MintPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TreasuryPool) > 0 {
		for iNdEx := len(m.TreasuryPool) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TreasuryPool[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMint(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CurrentMintVolume) > 0 {
		for iNdEx := len(m.CurrentMintVolume) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CurrentMintVolume[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMint(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size := m.AnnualProvisions.Size()
		i -= size
		if _, err := m.AnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *AllowedDenom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedDenom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedDenom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LargeDenom) > 0 {
		i -= len(m.LargeDenom)
		copy(dAtA[i:], m.LargeDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.LargeDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.LittleDenom) > 0 {
		i -= len(m.LittleDenom)
		copy(dAtA[i:], m.LittleDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.LittleDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MintPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TreasuryPool) > 0 {
		for _, e := range m.TreasuryPool {
			l = e.Size()
			n += 1 + l + sovMint(uint64(l))
		}
	}
	return n
}

func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.AnnualProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	if len(m.CurrentMintVolume) > 0 {
		for _, e := range m.CurrentMintVolume {
			l = e.Size()
			n += 1 + l + sovMint(uint64(l))
		}
	}
	return n
}

func (m *AllowedDenom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LittleDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = len(m.LargeDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MintPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: MintPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TreasuryPool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TreasuryPool = append(m.TreasuryPool, types.Coin{})
			if err := m.TreasuryPool[len(m.TreasuryPool)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentMintVolume", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentMintVolume = append(m.CurrentMintVolume, types.Coin{})
			if err := m.CurrentMintVolume[len(m.CurrentMintVolume)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *AllowedDenom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: AllowedDenom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedDenom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LittleDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LittleDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LargeDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LargeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
