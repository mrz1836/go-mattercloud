package main

import (
	"context"
	"log"
	"os"

	"github.com/mrz1836/go-mattercloud"
)

func main() {

	client, err := mattercloud.NewClient(
		os.Getenv("MATTERCLOUD_API_KEY"),
		mattercloud.NetworkMain,
		nil, nil,
	)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var result *mattercloud.BroadcastResponse
	if result, err = client.Broadcast(
		context.Background(),
		"01000000013d1296c350dcd7279da2582e4f01235fe6aacd0a18c4a8672d6e800defb0fedd000000006a473044022013c965d6bf1a383f5c798ff5b61773cfb30227b216cffb54e53e9a6c7eb4be7702202f1d69eb07bb5309dc5628a58c2facab81943324dcd0d9aa4ea39ec3f6ebb35e412103673dffd80561b87825658f74076da805c238e8c47f25b5d804893c335514d074ffffffff02c4090000000000001976a914f1e235c050767381252e6c48961371a2aad9628288acec220000000000001976a91467d93a70ac575e15abb31bc8272a00ab1495d48388ac00000000",
	); err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("success: ", result.Success, "tx id: ", result.Result.TxID)
}
