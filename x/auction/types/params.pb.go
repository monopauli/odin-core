// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auction/params.proto

package types

import (
	fmt "fmt"
	types1 "github.com/GeoDB-Limited/odin-core/x/coinswap/types"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Params is the data structure that keeps the parameters of the auction module.
type Params struct {
	// AuctionStartThreshold is the threshold at which the auction starts
	AuctionStartThreshold github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=auction_start_threshold,json=auctionStartThreshold,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"auction_start_threshold"`
	// ExchangeRate is a rate for buying coins throw the auction
	ExchangeRate types1.Exchange `protobuf:"bytes,2,opt,name=exchange_rate,json=exchangeRate,proto3" json:"exchange_rate"`
	// BlocksAuctionDuration is a duration of the auction
	BlocksAuctionDuration uint64 `protobuf:"varint,3,opt,name=blocks_auction_duration,json=blocksAuctionDuration,proto3" json:"blocks_auction_duration,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_e96a95233ccbd0c2, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAuctionStartThreshold() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.AuctionStartThreshold
	}
	return nil
}

func (m *Params) GetExchangeRate() types1.Exchange {
	if m != nil {
		return m.ExchangeRate
	}
	return types1.Exchange{}
}

func (m *Params) GetBlocksAuctionDuration() uint64 {
	if m != nil {
		return m.BlocksAuctionDuration
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "auction.Params")
}

func init() { proto.RegisterFile("auction/params.proto", fileDescriptor_e96a95233ccbd0c2) }

var fileDescriptor_e96a95233ccbd0c2 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x3f, 0x4f, 0xbb, 0x40,
	0x18, 0xc7, 0xa1, 0x6d, 0xfa, 0x4b, 0xe8, 0xcf, 0x85, 0xb4, 0xe9, 0x9f, 0x81, 0x36, 0x4e, 0x2c,
	0xbd, 0xb3, 0x35, 0x71, 0x30, 0x71, 0x10, 0x6b, 0x5c, 0x3a, 0x18, 0x74, 0x72, 0x21, 0x07, 0x9c,
	0x40, 0x5a, 0x78, 0xc8, 0xdd, 0xa1, 0xed, 0xec, 0xe8, 0xe2, 0x4b, 0x70, 0xf6, 0x95, 0x74, 0xec,
	0xe8, 0xa4, 0xa6, 0x5d, 0x7c, 0x19, 0x06, 0x38, 0x12, 0x9d, 0x78, 0x78, 0x3e, 0xcf, 0xe5, 0xf3,
	0x7c, 0xef, 0xb4, 0x36, 0xc9, 0x3c, 0x11, 0x41, 0x82, 0x53, 0xc2, 0x48, 0xcc, 0x51, 0xca, 0x40,
	0x80, 0xfe, 0x4f, 0x76, 0x07, 0xed, 0x00, 0x02, 0x28, 0x7a, 0x38, 0xaf, 0x4a, 0x3c, 0x30, 0x3c,
	0xe0, 0x31, 0x70, 0xec, 0x12, 0x4e, 0xf1, 0xc3, 0xc4, 0xa5, 0x82, 0x4c, 0xb0, 0x07, 0x51, 0x22,
	0x79, 0x3f, 0x00, 0x08, 0x96, 0x14, 0x17, 0x7f, 0x6e, 0x76, 0x8f, 0x49, 0xb2, 0x96, 0xa8, 0x93,
	0x8f, 0xf1, 0x47, 0x92, 0xfe, 0x11, 0x1e, 0x3e, 0xd7, 0xb4, 0xe6, 0x75, 0xd1, 0xd0, 0x9f, 0x54,
	0xad, 0x2b, 0xf5, 0x0e, 0x17, 0x84, 0x09, 0x47, 0x84, 0x8c, 0xf2, 0x10, 0x96, 0x7e, 0x4f, 0x1d,
	0xd5, 0xcd, 0xd6, 0xb4, 0x8f, 0x4a, 0x3f, 0xca, 0xfd, 0x48, 0xfa, 0xd1, 0x05, 0x44, 0x89, 0x75,
	0xb4, 0xf9, 0x18, 0x2a, 0x6f, 0x9f, 0x43, 0x33, 0x88, 0x44, 0x98, 0xb9, 0xc8, 0x83, 0x18, 0xcb,
	0x65, 0xcb, 0xcf, 0x98, 0xfb, 0x0b, 0x2c, 0xd6, 0x29, 0xe5, 0xc5, 0x01, 0x6e, 0x77, 0xa4, 0xeb,
	0x26, 0x57, 0xdd, 0x56, 0x26, 0xfd, 0x4c, 0x3b, 0xa0, 0x2b, 0x2f, 0x24, 0x49, 0x40, 0x1d, 0x46,
	0x04, 0xed, 0xd5, 0x46, 0xaa, 0xd9, 0x9a, 0xea, 0xa8, 0xda, 0x1f, 0x5d, 0x4a, 0x6c, 0x35, 0x72,
	0xa7, 0xfd, 0xbf, 0x1a, 0xb7, 0x89, 0xa0, 0xfa, 0x89, 0xd6, 0x75, 0x97, 0xe0, 0x2d, 0xb8, 0x53,
	0x45, 0xf1, 0x33, 0x46, 0xf2, 0xa2, 0x57, 0x1f, 0xa9, 0x66, 0xc3, 0xee, 0x94, 0xf8, 0xbc, 0xa4,
	0x33, 0x09, 0x4f, 0x1b, 0xdf, 0xaf, 0x43, 0xc5, 0x9a, 0x6f, 0x76, 0x86, 0xba, 0xdd, 0x19, 0xea,
	0xd7, 0xce, 0x50, 0x5f, 0xf6, 0x86, 0xb2, 0xdd, 0x1b, 0xca, 0xfb, 0xde, 0x50, 0xee, 0xa6, 0xbf,
	0x72, 0x5d, 0x51, 0x98, 0x59, 0xe3, 0x79, 0x14, 0x47, 0x82, 0xfa, 0x18, 0xfc, 0x28, 0x19, 0x7b,
	0xc0, 0x28, 0x5e, 0xe1, 0xea, 0x4d, 0x8b, 0x9c, 0x6e, 0xb3, 0xb8, 0xe2, 0xe3, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x9b, 0x52, 0x0c, 0xb7, 0xeb, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlocksAuctionDuration != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BlocksAuctionDuration))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.ExchangeRate.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.AuctionStartThreshold) > 0 {
		for iNdEx := len(m.AuctionStartThreshold) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AuctionStartThreshold[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AuctionStartThreshold) > 0 {
		for _, e := range m.AuctionStartThreshold {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.ExchangeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.BlocksAuctionDuration != 0 {
		n += 1 + sovParams(uint64(m.BlocksAuctionDuration))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionStartThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuctionStartThreshold = append(m.AuctionStartThreshold, types.Coin{})
			if err := m.AuctionStartThreshold[len(m.AuctionStartThreshold)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExchangeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksAuctionDuration", wireType)
			}
			m.BlocksAuctionDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksAuctionDuration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
