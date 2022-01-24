package mattercloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// AddressBalance this endpoint retrieves balance for a specific address.
//
// For more information: https://developers.mattercloud.io/#get-balance
func (c *Client) AddressBalance(ctx context.Context, address string) (balance *Balance, err error) {

	var resp string

	// GET https://api.mattercloud.io/api/v3/main/address/<address>/balance
	if resp, err = c.Request(
		ctx,
		"address/"+address+"/balance",
		http.MethodGet, nil,
	); err != nil {
		return
	}

	// Check for error
	if c.LastRequest.StatusCode != http.StatusOK {
		return nil, processError(resp)
	}

	// Process the response
	err = json.Unmarshal([]byte(resp), &balance)
	return
}

// AddressBalanceBatch this endpoint retrieves balances for multiple addresses at same time
//
// For more information: https://developers.mattercloud.io/#get-balance-batch
func (c *Client) AddressBalanceBatch(ctx context.Context, addresses []string) (balances []*Balance, err error) {

	// Check addresses
	if len(addresses) == 0 {
		err = fmt.Errorf("missing addresses")
		return
	}

	// Marshall into JSON
	var data []byte
	if data, err = json.Marshal(
		&AddressList{Addrs: strings.Join(addresses, ",")},
	); err != nil {
		return
	}

	var resp string

	// POST https://api.mattercloud.io/api/v3/main/address/balance
	if resp, err = c.Request(
		ctx,
		"address/balance",
		http.MethodPost, data,
	); err != nil {
		return
	}

	// Check for error
	if c.LastRequest.StatusCode != http.StatusOK {
		return nil, processError(resp)
	}

	// Process the response
	err = json.Unmarshal([]byte(resp), &balances)
	return
}

// AddressUtxos this endpoint retrieves utxos for a specific address
//
// For more information: https://developers.mattercloud.io/#get-utxos
func (c *Client) AddressUtxos(ctx context.Context, address string) (utxos []*UnspentTransaction, err error) {

	var resp string

	// GET https://api.mattercloud.io/api/v3/main/address/<address>/utxo
	if resp, err = c.Request(
		ctx,
		"address/"+address+"/utxo",
		http.MethodGet, nil,
	); err != nil {
		return
	}

	// Check for error
	if c.LastRequest.StatusCode != http.StatusOK {
		return nil, processError(resp)
	}

	// Process the response
	err = json.Unmarshal([]byte(resp), &utxos)
	return
}

// AddressUtxosBatch this endpoint retrieves utxos for multiple addresses
//
// For more information: https://developers.mattercloud.io/#get-utxos-batch
func (c *Client) AddressUtxosBatch(ctx context.Context, addresses []string) (utxos []*UnspentTransaction, err error) {

	// Check addresses
	if len(addresses) == 0 {
		err = fmt.Errorf("missing addresses")
		return
	}

	// Marshall into JSON
	var data []byte
	if data, err = json.Marshal(
		&AddressList{Addrs: strings.Join(addresses, ",")},
	); err != nil {
		return
	}

	var resp string

	// POST https://api.mattercloud.io/api/v3/main/address/utxo
	if resp, err = c.Request(
		ctx,
		"address/utxo",
		http.MethodPost, data,
	); err != nil {
		return
	}

	// Check for error
	if c.LastRequest.StatusCode != http.StatusOK {
		return nil, processError(resp)
	}

	// Process the response
	err = json.Unmarshal([]byte(resp), &utxos)
	return
}

// AddressHistory this endpoint retrieves history for a specific address
//
// For more information: https://developers.mattercloud.io/#get-history
func (c *Client) AddressHistory(ctx context.Context, address string) (history *History, err error) {

	var resp string

	// GET https://api.mattercloud.io/api/v3/main/address/<address>/history
	if resp, err = c.Request(
		ctx,
		"address/"+address+"/history",
		http.MethodGet, nil,
	); err != nil {
		return
	}

	// Check for error
	if c.LastRequest.StatusCode != http.StatusOK {
		return nil, processError(resp)
	}

	// Process the response
	err = json.Unmarshal([]byte(resp), &history)
	return
}

// AddressHistoryBatch this endpoint retrieves history for multiple addresses
//
// For more information: https://developers.mattercloud.io/#get-history-batch
func (c *Client) AddressHistoryBatch(ctx context.Context, addresses []string) (history *History, err error) {

	// Check addresses
	if len(addresses) == 0 {
		err = fmt.Errorf("missing addresses")
		return
	}

	// Marshall into JSON
	var data []byte
	if data, err = json.Marshal(
		&AddressList{Addrs: strings.Join(addresses, ",")},
	); err != nil {
		return
	}

	var resp string

	// POST https://api.mattercloud.io/api/v3/main/address/history
	if resp, err = c.Request(
		ctx,
		"address/history",
		http.MethodPost, data,
	); err != nil {
		return
	}

	// Check for error
	if c.LastRequest.StatusCode != http.StatusOK {
		return nil, processError(resp)
	}

	// Process the response
	err = json.Unmarshal([]byte(resp), &history)
	return
}
