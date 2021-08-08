package client

import (
	"ccmouse-go/C17/crawler/engine"
	"ccmouse-go/C17/crawler_distributed/config"
	"ccmouse-go/C17/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("item saver: got item %#d: %v", itemCount, item)

			// call RPC to save item
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: error "+"saveing item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}
