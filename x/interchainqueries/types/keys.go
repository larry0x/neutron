package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "interchainqueries"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_interchainqueries"
)

const (
	prefixRegisteredQuery = iota + 1
	prefixRegisteredQueryResult

	prefixSubmittedTx
)

var (
	RegisteredQueryKey       = []byte{prefixRegisteredQuery}
	RegisteredQueryResultKey = []byte{prefixRegisteredQueryResult}

	SubmittedTxKey = []byte{prefixSubmittedTx}

	LastRegisteredQueryIdKey = []byte{0x64}

	LastSubmittedTransactionIDKey = []byte{0x65}
)

func GetRegisteredQueryByIDKey(id uint64) []byte {
	return append(RegisteredQueryKey, sdk.Uint64ToBigEndian(id)...)
}

func GetLastSubmittedTransactionIDForQueryKey(queryID uint64) []byte {
	return append(LastSubmittedTransactionIDKey, sdk.Uint64ToBigEndian(queryID)...)
}

func GetSubmittedTransactionIDForQueryKeyPrefix(queryID uint64) []byte {
	return append(SubmittedTxKey, sdk.Uint64ToBigEndian(queryID)...)
}

func GetSubmittedTransactionIDForQueryKey(queryID uint64, transactionID uint64) []byte {
	return append(GetSubmittedTransactionIDForQueryKeyPrefix(queryID), sdk.Uint64ToBigEndian(transactionID)...)
}

func GetRegisteredQueryResultByIDKey(id uint64) []byte {
	return append(RegisteredQueryResultKey, sdk.Uint64ToBigEndian(id)...)
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
