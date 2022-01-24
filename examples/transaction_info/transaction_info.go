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

	var info *mattercloud.Transaction
	if info, err = client.Transaction(
		context.Background(),
		"908c26f8227fa99f1b26f99a19648653a1382fb3b37b03870e9c138894d29b3b",
	); err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("block: " + info.BlockHash)
}
