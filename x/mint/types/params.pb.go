// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mint/params.proto

package types

import (
	fmt "fmt"
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

// Params holds parameters for the mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// maximum annual change in inflation rate
	InflationRateChange github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=inflation_rate_change,json=inflationRateChange,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_rate_change" yaml:"inflation_rate_change"`
	// maximum inflation rate
	InflationMax github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=inflation_max,json=inflationMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_max" yaml:"inflation_max"`
	// minimum inflation rate
	InflationMin github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=inflation_min,json=inflationMin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_min" yaml:"inflation_min"`
	// goal of percent bonded atoms
	GoalBonded github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=goal_bonded,json=goalBonded,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"goal_bonded" yaml:"goal_bonded"`
	// expected blocks per year
	BlocksPerYear uint64 `protobuf:"varint,6,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty" yaml:"blocks_per_year"`
	// max amount to withdraw per time
	MaxWithdrawalPerTime github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,7,rep,name=max_withdrawal_per_time,json=maxWithdrawalPerTime,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"max_withdrawal_per_time" yaml:"max_withdrawal_per_time"`
	// map with smart contracts addresses
	IntegrationAddresses map[string]string `protobuf:"bytes,8,rep,name=integration_addresses,json=integrationAddresses,proto3" json:"integration_addresses,omitempty" yaml:"integration_addresses" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// flag if minting from air
	MintAir bool `protobuf:"varint,9,opt,name=mint_air,json=mintAir,proto3" json:"mint_air,omitempty" yaml:"mint_air"`
	// eligible to withdraw accounts
	EligibleAccountsPool []string `protobuf:"bytes,10,rep,name=eligible_accounts_pool,json=eligibleAccountsPool,proto3" json:"eligible_accounts_pool,omitempty" yaml:"eligible_accounts_pool"`
	// max allowed mint volume
	MaxAllowedMintVolume github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,11,rep,name=max_allowed_mint_volume,json=maxAllowedMintVolume,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"max_allowed_mint_volume" yaml:"max_allowed_mint_volume"`
	// allowed mint denoms
	AllowedMintDenoms []string `protobuf:"bytes,12,rep,name=allowed_mint_denoms,json=allowedMintDenoms,proto3" json:"allowed_mint_denoms,omitempty" yaml:"allowed_mint_denoms"`
	// allowed minter
	AllowedMinter []string `protobuf:"bytes,13,rep,name=allowed_minter,json=allowedMinter,proto3" json:"allowed_minter,omitempty" yaml:"allowed_minter"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d03971f940ff2c, []int{0}
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

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetBlocksPerYear() uint64 {
	if m != nil {
		return m.BlocksPerYear
	}
	return 0
}

func (m *Params) GetMaxWithdrawalPerTime() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.MaxWithdrawalPerTime
	}
	return nil
}

func (m *Params) GetIntegrationAddresses() map[string]string {
	if m != nil {
		return m.IntegrationAddresses
	}
	return nil
}

func (m *Params) GetMintAir() bool {
	if m != nil {
		return m.MintAir
	}
	return false
}

func (m *Params) GetEligibleAccountsPool() []string {
	if m != nil {
		return m.EligibleAccountsPool
	}
	return nil
}

func (m *Params) GetMaxAllowedMintVolume() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.MaxAllowedMintVolume
	}
	return nil
}

func (m *Params) GetAllowedMintDenoms() []string {
	if m != nil {
		return m.AllowedMintDenoms
	}
	return nil
}

func (m *Params) GetAllowedMinter() []string {
	if m != nil {
		return m.AllowedMinter
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "mint.Params")
	proto.RegisterMapType((map[string]string)(nil), "mint.Params.IntegrationAddressesEntry")
}

func init() { proto.RegisterFile("mint/params.proto", fileDescriptor_04d03971f940ff2c) }

var fileDescriptor_04d03971f940ff2c = []byte{
	// 734 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xbf, 0x4f, 0xdb, 0x40,
	0x14, 0x8e, 0x09, 0x3f, 0x92, 0x03, 0x4a, 0x31, 0x81, 0x9a, 0xa8, 0xd8, 0xa9, 0x07, 0x94, 0x05,
	0x5b, 0xb4, 0x4b, 0xc5, 0xd4, 0x98, 0xb4, 0x08, 0xa9, 0xa0, 0xc8, 0xaa, 0x8a, 0xda, 0xc5, 0xba,
	0xd8, 0xaf, 0xe1, 0x14, 0xdb, 0x17, 0x9d, 0x1d, 0x48, 0x86, 0x2e, 0xfd, 0x0b, 0x3a, 0x76, 0x64,
	0xe8, 0xd4, 0xbf, 0x84, 0x91, 0xb1, 0xea, 0xe0, 0x56, 0xb0, 0x74, 0xe8, 0x94, 0xbf, 0xa0, 0xf2,
	0x9d, 0x09, 0x81, 0x06, 0xa9, 0x54, 0x5d, 0x92, 0xbb, 0xef, 0x7d, 0xf7, 0x7d, 0xef, 0xdd, 0x3d,
	0x3f, 0xb4, 0x18, 0x90, 0x30, 0x36, 0x3b, 0x98, 0xe1, 0x20, 0x32, 0x3a, 0x8c, 0xc6, 0x54, 0x9e,
	0x4c, 0xa1, 0x72, 0xa9, 0x45, 0x5b, 0x94, 0x03, 0x66, 0xba, 0x12, 0xb1, 0xb2, 0xea, 0xd2, 0x28,
	0xa0, 0x91, 0xd9, 0xc4, 0x11, 0x98, 0x47, 0x9b, 0x4d, 0x88, 0xf1, 0xa6, 0xe9, 0x52, 0x12, 0x66,
	0xf1, 0x05, 0x2e, 0x97, 0xfe, 0x08, 0x40, 0xff, 0x85, 0xd0, 0x74, 0x83, 0xab, 0xcb, 0x6b, 0x08,
	0xa5, 0x01, 0xc7, 0x83, 0x90, 0x06, 0x8a, 0x54, 0x91, 0xaa, 0x45, 0xbb, 0x98, 0x22, 0xf5, 0x14,
	0x90, 0x3f, 0x48, 0x68, 0x99, 0x84, 0xef, 0x7c, 0x1c, 0x13, 0x1a, 0x3a, 0x0c, 0xc7, 0xe0, 0xb8,
	0x87, 0x38, 0x6c, 0x81, 0x32, 0x91, 0x52, 0xad, 0xfd, 0xd3, 0x44, 0xcb, 0x7d, 0x4b, 0xb4, 0xf5,
	0x16, 0x89, 0x0f, 0xbb, 0x4d, 0xc3, 0xa5, 0x81, 0x99, 0x65, 0x23, 0xfe, 0x36, 0x22, 0xaf, 0x6d,
	0xc6, 0xfd, 0x0e, 0x44, 0x46, 0x1d, 0xdc, 0x41, 0xa2, 0x3d, 0xec, 0xe3, 0xc0, 0xdf, 0xd2, 0xc7,
	0x8a, 0xea, 0xf6, 0xd2, 0x10, 0xb7, 0x71, 0x0c, 0xdb, 0x1c, 0x95, 0xdb, 0x68, 0xfe, 0x8a, 0x1e,
	0xe0, 0x9e, 0x92, 0xe7, 0xde, 0x2f, 0xee, 0xec, 0x5d, 0xba, 0xe9, 0x1d, 0xe0, 0x9e, 0x6e, 0xcf,
	0x0d, 0xf7, 0x7b, 0xb8, 0x77, 0xc3, 0x8c, 0x84, 0xca, 0xe4, 0x7f, 0x33, 0x23, 0xe1, 0x35, 0x33,
	0x12, 0xca, 0x80, 0x66, 0x5b, 0x14, 0xfb, 0x4e, 0x93, 0x86, 0x1e, 0x78, 0xca, 0x14, 0xb7, 0xaa,
	0xdf, 0xd9, 0x4a, 0x16, 0x56, 0x23, 0x52, 0xba, 0x8d, 0xd2, 0x9d, 0xc5, 0x37, 0xb2, 0x85, 0x16,
	0x9a, 0x3e, 0x75, 0xdb, 0x91, 0xd3, 0x01, 0xe6, 0xf4, 0x01, 0x33, 0x65, 0xba, 0x22, 0x55, 0x27,
	0xad, 0xf2, 0x20, 0xd1, 0x56, 0xc4, 0xe1, 0x1b, 0x04, 0xdd, 0x9e, 0x17, 0x48, 0x03, 0xd8, 0x1b,
	0xc0, 0x4c, 0xfe, 0x2c, 0xa1, 0x07, 0x01, 0xee, 0x39, 0xc7, 0x24, 0x3e, 0xf4, 0x18, 0x3e, 0xc6,
	0x3e, 0xe7, 0xc6, 0x24, 0x00, 0x65, 0xa6, 0x92, 0xaf, 0xce, 0x3e, 0x5e, 0x35, 0x44, 0x7a, 0x46,
	0xda, 0x87, 0x46, 0xd6, 0x87, 0xc6, 0x36, 0x25, 0xa1, 0x65, 0xa7, 0x25, 0x0d, 0x12, 0x4d, 0x15,
	0x5e, 0xb7, 0xe8, 0xe8, 0x5f, 0xbe, 0x6b, 0xd5, 0xbf, 0x28, 0x3a, 0x95, 0x8c, 0xec, 0x52, 0x80,
	0x7b, 0x07, 0x43, 0x91, 0x06, 0xb0, 0x57, 0x24, 0x00, 0xf9, 0x7d, 0xda, 0xaf, 0x31, 0xb4, 0x98,
	0xb8, 0x73, 0xec, 0x79, 0x0c, 0xa2, 0x08, 0x22, 0xa5, 0xc0, 0x73, 0x5c, 0x37, 0xf8, 0x67, 0x20,
	0x9a, 0xdf, 0xd8, 0xbd, 0x62, 0xd6, 0x2e, 0x89, 0xcf, 0xc3, 0x98, 0xf5, 0xad, 0xca, 0x68, 0xa7,
	0x8e, 0x91, 0xd3, 0xed, 0x12, 0x19, 0x73, 0x58, 0x36, 0x50, 0x81, 0x7f, 0x4e, 0x98, 0x30, 0xa5,
	0x58, 0x91, 0xaa, 0x05, 0x6b, 0x69, 0x90, 0x68, 0x0b, 0x59, 0xd9, 0x59, 0x44, 0xb7, 0x67, 0xd2,
	0x65, 0x8d, 0x30, 0xf9, 0x00, 0xad, 0x80, 0x4f, 0x5a, 0xa4, 0xe9, 0x83, 0x83, 0x5d, 0x97, 0x76,
	0xc3, 0x38, 0x72, 0x3a, 0x94, 0xfa, 0x0a, 0xaa, 0xe4, 0xab, 0x45, 0xeb, 0xd1, 0x20, 0xd1, 0xd6,
	0xc4, 0xe9, 0xf1, 0x3c, 0xdd, 0x2e, 0x5d, 0x06, 0x6a, 0x19, 0xde, 0xa0, 0xd4, 0x1f, 0x3e, 0x17,
	0xf6, 0x7d, 0x7a, 0x0c, 0x9e, 0xc3, 0xbd, 0x8f, 0xa8, 0xdf, 0x0d, 0x40, 0x99, 0xfd, 0x87, 0xe7,
	0x1a, 0xa3, 0x73, 0xf7, 0xe7, 0xaa, 0x09, 0x91, 0x3d, 0x12, 0xc6, 0xaf, 0xb9, 0x84, 0xbc, 0x8f,
	0x96, 0xae, 0x29, 0xf3, 0x31, 0x14, 0x29, 0x73, 0xbc, 0x78, 0x75, 0x90, 0x68, 0x65, 0x91, 0xc2,
	0x18, 0x92, 0x6e, 0x2f, 0xe2, 0x2b, 0x3d, 0x3e, 0xae, 0x22, 0xf9, 0x19, 0xba, 0x37, 0x4a, 0x05,
	0xa6, 0xcc, 0x73, 0xa9, 0xd5, 0x41, 0xa2, 0x2d, 0xff, 0x29, 0x05, 0x69, 0x9f, 0x8f, 0xa8, 0x00,
	0x2b, 0xef, 0xa0, 0xd5, 0x5b, 0xdb, 0x42, 0xbe, 0x8f, 0xf2, 0x6d, 0xe8, 0x67, 0x63, 0x32, 0x5d,
	0xca, 0x25, 0x34, 0x75, 0x84, 0xfd, 0x6e, 0x36, 0x0f, 0x6d, 0xb1, 0xd9, 0x9a, 0x78, 0x2a, 0x6d,
	0x15, 0x3e, 0x9d, 0x68, 0xb9, 0x9f, 0x27, 0x9a, 0x64, 0xed, 0x9e, 0x9e, 0xab, 0xd2, 0xd9, 0xb9,
	0x2a, 0xfd, 0x38, 0x57, 0xa5, 0x8f, 0x17, 0x6a, 0xee, 0xec, 0x42, 0xcd, 0x7d, 0xbd, 0x50, 0x73,
	0x6f, 0xcd, 0x91, 0xeb, 0xdb, 0x01, 0x5a, 0xb7, 0x36, 0x5e, 0x92, 0x80, 0xc4, 0xe0, 0x99, 0xd4,
	0x23, 0xe1, 0x86, 0x4b, 0x19, 0x98, 0x3d, 0x3e, 0xb9, 0xc5, 0x5d, 0x36, 0xa7, 0xf9, 0x00, 0x7f,
	0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x09, 0x2c, 0xa8, 0xf9, 0x22, 0x06, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.MintDenom != that1.MintDenom {
		return false
	}
	if !this.InflationRateChange.Equal(that1.InflationRateChange) {
		return false
	}
	if !this.InflationMax.Equal(that1.InflationMax) {
		return false
	}
	if !this.InflationMin.Equal(that1.InflationMin) {
		return false
	}
	if !this.GoalBonded.Equal(that1.GoalBonded) {
		return false
	}
	if this.BlocksPerYear != that1.BlocksPerYear {
		return false
	}
	if len(this.MaxWithdrawalPerTime) != len(that1.MaxWithdrawalPerTime) {
		return false
	}
	for i := range this.MaxWithdrawalPerTime {
		if !this.MaxWithdrawalPerTime[i].Equal(&that1.MaxWithdrawalPerTime[i]) {
			return false
		}
	}
	if len(this.IntegrationAddresses) != len(that1.IntegrationAddresses) {
		return false
	}
	for i := range this.IntegrationAddresses {
		if this.IntegrationAddresses[i] != that1.IntegrationAddresses[i] {
			return false
		}
	}
	if this.MintAir != that1.MintAir {
		return false
	}
	if len(this.EligibleAccountsPool) != len(that1.EligibleAccountsPool) {
		return false
	}
	for i := range this.EligibleAccountsPool {
		if this.EligibleAccountsPool[i] != that1.EligibleAccountsPool[i] {
			return false
		}
	}
	if len(this.MaxAllowedMintVolume) != len(that1.MaxAllowedMintVolume) {
		return false
	}
	for i := range this.MaxAllowedMintVolume {
		if !this.MaxAllowedMintVolume[i].Equal(&that1.MaxAllowedMintVolume[i]) {
			return false
		}
	}
	if len(this.AllowedMintDenoms) != len(that1.AllowedMintDenoms) {
		return false
	}
	for i := range this.AllowedMintDenoms {
		if this.AllowedMintDenoms[i] != that1.AllowedMintDenoms[i] {
			return false
		}
	}
	if len(this.AllowedMinter) != len(that1.AllowedMinter) {
		return false
	}
	for i := range this.AllowedMinter {
		if this.AllowedMinter[i] != that1.AllowedMinter[i] {
			return false
		}
	}
	return true
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
	if len(m.AllowedMinter) > 0 {
		for iNdEx := len(m.AllowedMinter) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedMinter[iNdEx])
			copy(dAtA[i:], m.AllowedMinter[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.AllowedMinter[iNdEx])))
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.AllowedMintDenoms) > 0 {
		for iNdEx := len(m.AllowedMintDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedMintDenoms[iNdEx])
			copy(dAtA[i:], m.AllowedMintDenoms[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.AllowedMintDenoms[iNdEx])))
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.MaxAllowedMintVolume) > 0 {
		for iNdEx := len(m.MaxAllowedMintVolume) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxAllowedMintVolume[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.EligibleAccountsPool) > 0 {
		for iNdEx := len(m.EligibleAccountsPool) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.EligibleAccountsPool[iNdEx])
			copy(dAtA[i:], m.EligibleAccountsPool[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.EligibleAccountsPool[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.MintAir {
		i--
		if m.MintAir {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.IntegrationAddresses) > 0 {
		for k := range m.IntegrationAddresses {
			v := m.IntegrationAddresses[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintParams(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintParams(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintParams(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.MaxWithdrawalPerTime) > 0 {
		for iNdEx := len(m.MaxWithdrawalPerTime) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxWithdrawalPerTime[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.BlocksPerYear != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BlocksPerYear))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.GoalBonded.Size()
		i -= size
		if _, err := m.GoalBonded.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.InflationMin.Size()
		i -= size
		if _, err := m.InflationMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.InflationMax.Size()
		i -= size
		if _, err := m.InflationMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.InflationRateChange.Size()
		i -= size
		if _, err := m.InflationRateChange.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
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
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.InflationRateChange.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.InflationMax.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.InflationMin.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.GoalBonded.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.BlocksPerYear != 0 {
		n += 1 + sovParams(uint64(m.BlocksPerYear))
	}
	if len(m.MaxWithdrawalPerTime) > 0 {
		for _, e := range m.MaxWithdrawalPerTime {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.IntegrationAddresses) > 0 {
		for k, v := range m.IntegrationAddresses {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovParams(uint64(len(k))) + 1 + len(v) + sovParams(uint64(len(v)))
			n += mapEntrySize + 1 + sovParams(uint64(mapEntrySize))
		}
	}
	if m.MintAir {
		n += 2
	}
	if len(m.EligibleAccountsPool) > 0 {
		for _, s := range m.EligibleAccountsPool {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.MaxAllowedMintVolume) > 0 {
		for _, e := range m.MaxAllowedMintVolume {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.AllowedMintDenoms) > 0 {
		for _, s := range m.AllowedMintDenoms {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.AllowedMinter) > 0 {
		for _, s := range m.AllowedMinter {
			l = len(s)
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
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
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
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRateChange", wireType)
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
			if err := m.InflationRateChange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMax", wireType)
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
			if err := m.InflationMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMin", wireType)
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
			if err := m.InflationMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoalBonded", wireType)
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
			if err := m.GoalBonded.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerYear", wireType)
			}
			m.BlocksPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxWithdrawalPerTime", wireType)
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
			m.MaxWithdrawalPerTime = append(m.MaxWithdrawalPerTime, types.Coin{})
			if err := m.MaxWithdrawalPerTime[len(m.MaxWithdrawalPerTime)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntegrationAddresses", wireType)
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
			if m.IntegrationAddresses == nil {
				m.IntegrationAddresses = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthParams
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthParams
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthParams
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthParams
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipParams(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthParams
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.IntegrationAddresses[mapkey] = mapvalue
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintAir", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MintAir = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EligibleAccountsPool", wireType)
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
			m.EligibleAccountsPool = append(m.EligibleAccountsPool, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAllowedMintVolume", wireType)
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
			m.MaxAllowedMintVolume = append(m.MaxAllowedMintVolume, types.Coin{})
			if err := m.MaxAllowedMintVolume[len(m.MaxAllowedMintVolume)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedMintDenoms", wireType)
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
			m.AllowedMintDenoms = append(m.AllowedMintDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedMinter", wireType)
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
			m.AllowedMinter = append(m.AllowedMinter, string(dAtA[iNdEx:postIndex]))
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
