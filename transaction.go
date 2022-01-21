package mattercloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Transaction this endpoint retrieves specific transaction
//
// For more information: https://developers.mattercloud.io/#get-transaction
func (c *Client) Transaction(ctx context.Context, tx string) (transaction *Transaction, err error) {

	var resp string

	// GET https://api.mattercloud.io/api/v3/main/tx/<txid>
	if resp, err = c.Request(
		ctx,
		"tx/"+tx,
		http.MethodGet, nil,
	); err != nil {
		return
	}

	// Check for error
	if c.LastRequest.StatusCode != http.StatusOK {
		var apiError APIInternalError
		if err = json.Unmarshal([]byte(resp), &apiError); err != nil {
			return
		}
		err = fmt.Errorf("error: %s", apiError.ErrorMessage)
		return
	}

	// Process the response
	err = json.Unmarshal([]byte(resp), &transaction)
	return
}

// TransactionBatch this endpoint retrieves details for multiple transactions at same time
//
// For more information: https://developers.mattercloud.io/#get-transaction-batch
func (c *Client) TransactionBatch(ctx context.Context, txIDs []string) (transactions []*Transaction, err error) {

	// Check ids
	if len(txIDs) == 0 {
		err = fmt.Errorf("missing tx ids")
		return
	}

	// Marshall into JSON
	var data []byte
	if data, err = json.Marshal(&TransactionList{TxIDs: strings.Join(txIDs, ",")}); err != nil {
		return
	}

	var resp string

	// POST https://api.mattercloud.io/api/v3/main/tx
	if resp, err = c.Request(
		ctx,
		"tx",
		http.MethodPost, data,
	); err != nil {
		return
	}

	// Check for error
	if c.LastRequest.StatusCode != http.StatusOK {
		var apiError APIInternalError
		if err = json.Unmarshal([]byte(resp), &apiError); err != nil {
			return
		}
		err = fmt.Errorf("error: %s", apiError.ErrorMessage)
		return
	}

	// Process the response
	err = json.Unmarshal([]byte(resp), &transactions)
	return
}

// Broadcast this endpoint broadcasts a raw transaction to the network
//
// For more information: https://developers.mattercloud.io/#broadcast-transaction
func (c *Client) Broadcast(ctx context.Context, rawTx string) (response *BroadcastResponse, err error) {

	var resp string
	// POST https://api.mattercloud.io/api/v3/main/tx/send

	resp, err = c.Request(
		ctx,
		"tx/send",
		http.MethodPost, []byte(fmt.Sprintf(`{"rawtx":"%s"}`, rawTx)),
	)
	if err != nil {
		return
	}

	// Check for error
	if c.LastRequest.StatusCode != http.StatusOK {
		var apiError APIInternalError
		if err = json.Unmarshal([]byte(resp), &apiError); err != nil {
			return
		}
		err = fmt.Errorf("error: %s", apiError.ErrorMessage)
		return
	}

	// Process the response
	err = json.Unmarshal([]byte(resp), &response)
	return
}
