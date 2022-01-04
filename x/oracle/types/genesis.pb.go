// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/v1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the oracle module's genesis state.
type GenesisState struct {
	Params             Params         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	DataSources        []DataSource   `protobuf:"bytes,2,rep,name=data_sources,json=dataSources,proto3" json:"data_sources"`
	OracleScripts      []OracleScript `protobuf:"bytes,3,rep,name=oracle_scripts,json=oracleScripts,proto3" json:"oracle_scripts"`
	OraclePool         OraclePool     `protobuf:"bytes,4,opt,name=oracle_pool,json=oraclePool,proto3" json:"oracle_pool"`
	ModuleCoinsAccount string         `protobuf:"bytes,5,opt,name=module_coins_account,json=moduleCoinsAccount,proto3" json:"module_coins_account,omitempty" yaml:"module_coins_account"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_14b982a0a6345d1d, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetDataSources() []DataSource {
	if m != nil {
		return m.DataSources
	}
	return nil
}

func (m *GenesisState) GetOracleScripts() []OracleScript {
	if m != nil {
		return m.OracleScripts
	}
	return nil
}

func (m *GenesisState) GetOraclePool() OraclePool {
	if m != nil {
		return m.OraclePool
	}
	return OraclePool{}
}

func (m *GenesisState) GetModuleCoinsAccount() string {
	if m != nil {
		return m.ModuleCoinsAccount
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "oracle.v1.GenesisState")
}

func init() { proto.RegisterFile("oracle/v1/genesis.proto", fileDescriptor_14b982a0a6345d1d) }

var fileDescriptor_14b982a0a6345d1d = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x6a, 0xea, 0x40,
	0x14, 0x86, 0x13, 0xf5, 0x0a, 0x4e, 0xbc, 0x17, 0x6e, 0xd0, 0x6b, 0xf0, 0x42, 0x14, 0x57, 0x6e,
	0xcc, 0xa0, 0xdd, 0x95, 0x52, 0x68, 0x2a, 0xb8, 0x68, 0xa1, 0x56, 0x77, 0xdd, 0x84, 0x71, 0x32,
	0xa4, 0x81, 0x24, 0x27, 0x64, 0x26, 0x52, 0xdf, 0xa2, 0x8f, 0xe5, 0xd2, 0x55, 0xe9, 0x4a, 0x8a,
	0xbe, 0x41, 0x9f, 0xa0, 0x38, 0x19, 0xac, 0x60, 0x77, 0xc3, 0xf7, 0xff, 0xe7, 0x3b, 0x03, 0x07,
	0xb5, 0x20, 0x23, 0x34, 0x62, 0x78, 0x39, 0xc4, 0x01, 0x4b, 0x18, 0x0f, 0xb9, 0x93, 0x66, 0x20,
	0xc0, 0xac, 0x15, 0x81, 0xb3, 0x1c, 0xb6, 0x1b, 0x01, 0x04, 0x20, 0x29, 0x3e, 0xbc, 0x8a, 0x42,
	0xfb, 0xdf, 0xf7, 0xa4, 0xaa, 0x9e, 0xf1, 0x94, 0x64, 0x24, 0x56, 0xc2, 0xde, 0x5b, 0x09, 0xd5,
	0x27, 0xc5, 0x8a, 0xb9, 0x20, 0x82, 0x99, 0x18, 0x55, 0x8b, 0x82, 0xa5, 0x77, 0xf5, 0xbe, 0x31,
	0xfa, 0xeb, 0x1c, 0x57, 0x3a, 0x53, 0x19, 0xb8, 0x95, 0xf5, 0xb6, 0xa3, 0xcd, 0x54, 0xcd, 0xbc,
	0x46, 0x75, 0x9f, 0x08, 0xe2, 0x71, 0xc8, 0x33, 0xca, 0xb8, 0x55, 0xea, 0x96, 0xfb, 0xc6, 0xa8,
	0x79, 0x32, 0x36, 0x26, 0x82, 0xcc, 0x65, 0xaa, 0x46, 0x0d, 0xff, 0x48, 0xb8, 0x39, 0x46, 0x7f,
	0x8a, 0xaa, 0xc7, 0x69, 0x16, 0xa6, 0x82, 0x5b, 0x65, 0x69, 0x68, 0x9d, 0x18, 0x1e, 0xe4, 0x6b,
	0x2e, 0x73, 0xe5, 0xf8, 0x0d, 0x27, 0x8c, 0x9b, 0x57, 0xc8, 0x50, 0x96, 0x14, 0x20, 0xb2, 0x2a,
	0xf2, 0xef, 0xcd, 0x33, 0xc5, 0x14, 0x20, 0x52, 0x02, 0x04, 0x47, 0x62, 0x3e, 0xa2, 0x46, 0x0c,
	0x7e, 0x1e, 0x31, 0x8f, 0x42, 0x98, 0x70, 0x8f, 0x50, 0x0a, 0x79, 0x22, 0xac, 0x5f, 0x5d, 0xbd,
	0x5f, 0x73, 0x3b, 0x9f, 0xdb, 0xce, 0xff, 0x15, 0x89, 0xa3, 0xcb, 0xde, 0x4f, 0xad, 0xde, 0xcc,
	0x2c, 0xf0, 0xed, 0x81, 0xde, 0x14, 0xd0, 0xbd, 0x5b, 0xef, 0x6c, 0x7d, 0xb3, 0xb3, 0xf5, 0x8f,
	0x9d, 0xad, 0xbf, 0xee, 0x6d, 0x6d, 0xb3, 0xb7, 0xb5, 0xf7, 0xbd, 0xad, 0x3d, 0x0d, 0x83, 0x50,
	0x3c, 0xe7, 0x0b, 0x87, 0x42, 0x8c, 0x27, 0x0c, 0xc6, 0xee, 0xe0, 0x3e, 0x8c, 0x43, 0xc1, 0x7c,
	0x0c, 0x7e, 0x98, 0x0c, 0x28, 0x64, 0x0c, 0xbf, 0xa8, 0xeb, 0x61, 0xb1, 0x4a, 0x19, 0x5f, 0x54,
	0xe5, 0xb1, 0x2e, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x36, 0xff, 0x3e, 0x42, 0x18, 0x02, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ModuleCoinsAccount) > 0 {
		i -= len(m.ModuleCoinsAccount)
		copy(dAtA[i:], m.ModuleCoinsAccount)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ModuleCoinsAccount)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.OraclePool.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.OracleScripts) > 0 {
		for iNdEx := len(m.OracleScripts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OracleScripts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DataSources) > 0 {
		for iNdEx := len(m.DataSources) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DataSources[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.DataSources) > 0 {
		for _, e := range m.DataSources {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OracleScripts) > 0 {
		for _, e := range m.OracleScripts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.OraclePool.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.ModuleCoinsAccount)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataSources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataSources = append(m.DataSources, DataSource{})
			if err := m.DataSources[len(m.DataSources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleScripts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleScripts = append(m.OracleScripts, OracleScript{})
			if err := m.OracleScripts[len(m.OracleScripts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OraclePool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OraclePool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleCoinsAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleCoinsAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
