package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module.
	ModuleName = "oracle"
	// ModuleVersion defines the current module version
	ModuleVersion = 1

	// Version defines the current version the IBC oracle module supports
	// TODO: Using our new version for oracle packet (new ics?)
	Version = "oracle-1"

	// StoreKey to be used when creating the KVStore.
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the oracle module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the oracle module
	RouterKey = ModuleName

	// PortID is the default port id that oracle module binds to.
	PortID = ModuleName
)

var (
	// RollingSeedSizeInBytes is the size of rolling block hash for random seed.
	RollingSeedSizeInBytes = 32
	// GlobalStoreKeyPrefix is the prefix for global primitive state variables.
	GlobalStoreKeyPrefix = []byte{0x00}
	// RollingSeedStoreKey is the key that keeps the seed based on the first 8-bit of the most recent 32 block hashes.
	RollingSeedStoreKey = append(GlobalStoreKeyPrefix, []byte("RollingSeed")...)
	// RequestCountStoreKey is the key that keeps the total request count.
	RequestCountStoreKey = append(GlobalStoreKeyPrefix, []byte("RequestCount")...)
	// RequestLastExpiredStoreKey is the key that keeps the ID of the last expired request, or 0 if none.
	RequestLastExpiredStoreKey = append(GlobalStoreKeyPrefix, []byte("RequestLastExpired")...)
	// PendingResolveListStoreKey is the key that keeps the list of pending-resolve requests.
	PendingResolveListStoreKey = append(GlobalStoreKeyPrefix, []byte("PendingList")...)
	// DataSourceCountStoreKey is the key that keeps the total data source count.
	DataSourceCountStoreKey = append(GlobalStoreKeyPrefix, []byte("DataSourceCount")...)
	// OracleScriptCountStoreKey is the key that keeps the total oracle sciprt count.
	OracleScriptCountStoreKey = append(GlobalStoreKeyPrefix, []byte("OracleScriptCount")...)
	// AccumulatedDataProvidersRewardsStoreKey the key that keeps the total allocated rewards to data providers for one day
	AccumulatedDataProvidersRewardsStoreKey = append(GlobalStoreKeyPrefix, []byte("AllocateRewardsToDataProviders")...)
	// AccumulatedPaymentsForDataStoreKey the key that keeps the accumulated payments for data
	AccumulatedPaymentsForDataStoreKey = append(GlobalStoreKeyPrefix, []byte("AccumulatedPaymentsForData")...)
	// OraclePoolStoreKey is the key that keeps the oracle pool
	OraclePoolStoreKey = append(GlobalStoreKeyPrefix, []byte("OraclePool")...) // key for global oracle pool state

	OracleModuleCoinsAccountKey = append(GlobalStoreKeyPrefix, []byte("OracleModuleCoinsAccount")...)

	// RequestStoreKeyPrefix is the prefix for request store.
	RequestStoreKeyPrefix = []byte{0x01}
	// ReportStoreKeyPrefix is the prefix for report store.
	ReportStoreKeyPrefix = []byte{0x02}
	// DataSourceStoreKeyPrefix is the prefix for data source store.
	DataSourceStoreKeyPrefix = []byte{0x03}
	// OracleScriptStoreKeyPrefix is the prefix for oracle script store.
	OracleScriptStoreKeyPrefix = []byte{0x04}
	// ReporterStoreKeyPrefix is the prefix for reporter store.
	ReporterStoreKeyPrefix = []byte{0x05}
	// ValidatorStatusKeyPrefix is the prefix for validator status store.
	ValidatorStatusKeyPrefix     = []byte{0x06}
	DataProviderRewardsKeyPrefix = []byte{0x07}
	// DataRequesterFeesKeyPrefix is the prefix for the data requester address with accumulated fee
	DataRequesterFeesKeyPrefix = []byte{0x08}
	// ResultStoreKeyPrefix is the prefix for request result store.
	ResultStoreKeyPrefix = []byte{0xff}

	// PortKey defines the key to store the port ID in store
	PortKey = []byte{0xf0}
)

// RequestStoreKey returns the key to retrieve a specific request from the store.
func RequestStoreKey(requestID RequestID) []byte {
	return append(RequestStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(requestID))...)
}

// ReportStoreKey returns the key to retrieve all data reports for a request.
func ReportStoreKey(requestID RequestID) []byte {
	return append(ReportStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(requestID))...)
}

// DataSourceStoreKey returns the key to retrieve a specific data source from the store.
func DataSourceStoreKey(dataSourceID DataSourceID) []byte {
	return append(DataSourceStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(dataSourceID))...)
}

// OracleScriptStoreKey returns the key to retrieve a specific oracle script from the store.
func OracleScriptStoreKey(oracleScriptID OracleScriptID) []byte {
	return append(OracleScriptStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(oracleScriptID))...)
}

// ReporterStoreKey returns the key to check whether an address is a reporter of a validator.
func ReporterStoreKey(validatorAddress sdk.ValAddress, reporterAddress sdk.AccAddress) []byte {
	buf := append(ReporterStoreKeyPrefix, []byte(validatorAddress)...)
	buf = append(buf, []byte(reporterAddress)...)
	return buf
}

// ValidatorStatusStoreKey returns the key to a validator's status.
func ValidatorStatusStoreKey(v sdk.ValAddress) []byte {
	return append(ValidatorStatusKeyPrefix, v.Bytes()...)
}

// ResultStoreKey returns the key to a request result in the store.
func ResultStoreKey(requestID RequestID) []byte {
	return append(ResultStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(requestID))...)
}

// ReportsOfValidatorPrefixKey returns the prefix key to get all reports for a request from a validator.
func ReportsOfValidatorPrefixKey(reqID RequestID, val sdk.ValAddress) []byte {
	buf := append(ReportStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(reqID))...)
	buf = append(buf, val.Bytes()...)
	return buf
}

// ReportersOfValidatorPrefixKey returns the prefix key to get all reporters of a validator.
func ReportersOfValidatorPrefixKey(val sdk.ValAddress) []byte {
	return append(ReporterStoreKeyPrefix, val.Bytes()...)
}

func DataProviderRewardsPrefixKey(acc sdk.AccAddress) []byte {
	return append(DataProviderRewardsKeyPrefix, acc.Bytes()...)
}
