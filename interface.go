package mattercloud

import "context"

// ClientInterface is the MatterCloud client interface
type ClientInterface interface {
	Network() NetworkType
	AddressBalance(ctx context.Context, address string) (balance *Balance, err error)
	AddressBalanceBatch(ctx context.Context, addresses []string) (balances []*Balance, err error)
	AddressHistory(ctx context.Context, address string) (history *History, err error)
	AddressHistoryBatch(ctx context.Context, addresses []string) (history *History, err error)
	AddressUtxos(ctx context.Context, address string) (utxos []*UnspentTransaction, err error)
	AddressUtxosBatch(ctx context.Context, addresses []string) (utxos []*UnspentTransaction, err error)
	Broadcast(ctx context.Context, rawTx string) (response *BroadcastResponse, err error)
	Transaction(ctx context.Context, tx string) (transaction *Transaction, err error)
	TransactionBatch(ctx context.Context, txIDs []string) (transactions []*Transaction, err error)
	Request(ctx context.Context, endpoint, method string,
		payload []byte) (response string, err error)
}
