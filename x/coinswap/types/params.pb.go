// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: coinswap/params.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Exchange struct {
	From           string                                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To             string                                 `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	RateMultiplier github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=rate_multiplier,json=rateMultiplier,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate_multiplier"`
}

func (m *Exchange) Reset()         { *m = Exchange{} }
func (m *Exchange) String() string { return proto.CompactTextString(m) }
func (*Exchange) ProtoMessage()    {}
func (*Exchange) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba08dc77a0b23efc, []int{0}
}
func (m *Exchange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Exchange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Exchange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Exchange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Exchange.Merge(m, src)
}
func (m *Exchange) XXX_Size() int {
	return m.Size()
}
func (m *Exchange) XXX_DiscardUnknown() {
	xxx_messageInfo_Exchange.DiscardUnknown(m)
}

var xxx_messageInfo_Exchange proto.InternalMessageInfo

func (m *Exchange) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Exchange) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type Params struct {
	ExchangeRates []Exchange `protobuf:"bytes,1,rep,name=exchange_rates,json=exchangeRates,proto3" json:"exchange_rates"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba08dc77a0b23efc, []int{1}
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

func (m *Params) GetExchangeRates() []Exchange {
	if m != nil {
		return m.ExchangeRates
	}
	return nil
}

func init() {
	proto.RegisterType((*Exchange)(nil), "coinswap.Exchange")
	proto.RegisterType((*Params)(nil), "coinswap.Params")
}

func init() { proto.RegisterFile("coinswap/params.proto", fileDescriptor_ba08dc77a0b23efc) }

var fileDescriptor_ba08dc77a0b23efc = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x31, 0x4f, 0xc2, 0x40,
	0x14, 0xc7, 0x7b, 0x40, 0x08, 0xd6, 0x88, 0x49, 0xa3, 0x49, 0x65, 0xb8, 0x12, 0x06, 0xc3, 0x42,
	0x2f, 0xc8, 0xe6, 0x62, 0x82, 0x38, 0x98, 0x88, 0x25, 0xd5, 0xc4, 0xc4, 0x85, 0x5c, 0xcb, 0x51,
	0x1a, 0x69, 0x5f, 0xd3, 0x3b, 0x14, 0x3e, 0x81, 0xab, 0xa3, 0x23, 0x1f, 0x87, 0x91, 0xd1, 0x38,
	0x10, 0x03, 0x8b, 0x1f, 0xc3, 0xdc, 0x41, 0x89, 0xd3, 0xbd, 0xfb, 0xff, 0xff, 0xf7, 0xee, 0x77,
	0xef, 0xf4, 0x53, 0x1f, 0xc2, 0x98, 0xbf, 0xd1, 0x84, 0x24, 0x34, 0xa5, 0x11, 0xb7, 0x93, 0x14,
	0x04, 0x18, 0xa5, 0x4c, 0xae, 0x9c, 0x04, 0x10, 0x80, 0x12, 0x89, 0xac, 0xb6, 0x7e, 0x05, 0xfb,
	0xc0, 0x23, 0xe0, 0xc4, 0xa3, 0x9c, 0x91, 0xd7, 0xa6, 0xc7, 0x04, 0x6d, 0x12, 0x79, 0x66, 0xe7,
	0x9f, 0x05, 0x00, 0xc1, 0x98, 0x11, 0xb5, 0xf3, 0x26, 0x43, 0x42, 0xe3, 0xd9, 0xd6, 0xaa, 0xbd,
	0x23, 0xbd, 0x74, 0x33, 0xf5, 0x47, 0x34, 0x0e, 0x98, 0x61, 0xe8, 0x85, 0x61, 0x0a, 0x91, 0x89,
	0xaa, 0xa8, 0x7e, 0xe0, 0xaa, 0xda, 0x28, 0xeb, 0x39, 0x01, 0x66, 0x4e, 0x29, 0x39, 0x01, 0xc6,
	0x93, 0x7e, 0x9c, 0x52, 0xc1, 0xfa, 0xd1, 0x64, 0x2c, 0xc2, 0x64, 0x1c, 0xb2, 0xd4, 0xcc, 0x4b,
	0xb3, 0x6d, 0x2f, 0x56, 0x96, 0xf6, 0xbd, 0xb2, 0xce, 0x83, 0x50, 0x8c, 0x26, 0x9e, 0xed, 0x43,
	0x44, 0x76, 0x5c, 0xdb, 0xa5, 0xc1, 0x07, 0x2f, 0x44, 0xcc, 0x12, 0xc6, 0xed, 0x0e, 0xf3, 0xdd,
	0xb2, 0x6c, 0xd3, 0xdd, 0x77, 0xa9, 0x3d, 0xe8, 0xc5, 0x9e, 0x7a, 0xb4, 0x71, 0xa5, 0x97, 0xd9,
	0x0e, 0xa9, 0x2f, 0x43, 0xdc, 0x44, 0xd5, 0x7c, 0xfd, 0xf0, 0xc2, 0xb0, 0xb3, 0x39, 0xd8, 0x19,
	0x72, 0xbb, 0x20, 0x6f, 0x75, 0x8f, 0xb2, 0xbc, 0x2b, 0xe3, 0x97, 0xa5, 0xcf, 0xb9, 0x85, 0x7e,
	0xe7, 0x96, 0xd6, 0xee, 0x2e, 0xd6, 0x18, 0x2d, 0xd7, 0x18, 0xfd, 0xac, 0x31, 0xfa, 0xd8, 0x60,
	0x6d, 0xb9, 0xc1, 0xda, 0xd7, 0x06, 0x6b, 0xcf, 0xad, 0x7f, 0x98, 0x4e, 0xe7, 0xf6, 0xbe, 0xd1,
	0x73, 0x9d, 0x47, 0xe7, 0xda, 0xb9, 0x23, 0x30, 0x08, 0xe3, 0x86, 0x0f, 0x29, 0x23, 0x53, 0xb2,
	0xff, 0x0f, 0xc5, 0xed, 0x15, 0xd5, 0xd0, 0x5a, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x44, 0x59,
	0xa8, 0x5d, 0xa8, 0x01, 0x00, 0x00,
}

func (m *Exchange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Exchange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Exchange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RateMultiplier.Size()
		i -= size
		if _, err := m.RateMultiplier.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintParams(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintParams(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if len(m.ExchangeRates) > 0 {
		for iNdEx := len(m.ExchangeRates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExchangeRates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *Exchange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.RateMultiplier.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ExchangeRates) > 0 {
		for _, e := range m.ExchangeRates {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Exchange) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Exchange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Exchange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateMultiplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RateMultiplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeRates", wireType)
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
			m.ExchangeRates = append(m.ExchangeRates, Exchange{})
			if err := m.ExchangeRates[len(m.ExchangeRates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
