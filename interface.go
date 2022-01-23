package mattercloud

import "context"

// AddressService is the MatterCloud address related requests
type AddressService interface {
	AddressBalance(ctx context.Context, address string) (balance *Balance, err error)
	AddressBalanceBatch(ctx context.Context, addresses []string) (balances []*Balance, err error)
	AddressHistory(ctx context.Context, address string) (history *History, err error)
	AddressHistoryBatch(ctx context.Context, addresses []string) (history *History, err error)
	AddressUtxos(ctx context.Context, address string) (utxos []*UnspentTransaction, err error)
	AddressUtxosBatch(ctx context.Context, addresses []string) (utxos []*UnspentTransaction, err error)
}

// TransactionService is the MatterCloud transaction related requests
type TransactionService interface {
	Broadcast(ctx context.Context, rawTx string) (response *BroadcastResponse, err error)
	Transaction(ctx context.Context, tx string) (transaction *Transaction, err error)
	TransactionBatch(ctx context.Context, txIDs []string) (transactions []*Transaction, err error)
}

// ClientInterface is the MatterCloud client interface
type ClientInterface interface {
	AddressService
	TransactionService
	Network() NetworkType
	Request(ctx context.Context, endpoint, method string, payload []byte) (response string, err error)
}
