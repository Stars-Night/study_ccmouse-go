package persist

import (
	"ccmouse-go/C16/crawler2/engine"
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("item saver: got item %#d: %v", itemCount, item)

			err := save(item)
			if err != nil {
				log.Printf("Item Saver: error "+"saveing item %v: %v", item, err)
			}
		}
	}()
	return out
}

func save(item engine.Item) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.99.100:9200"),
		// Must turn off sniff in docker
		elastic.SetSniff(false),
	)

	if err != nil {
		return nil
	}

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().Index("dating_profile").Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())
	if err != nil {
		return nil
	}

	return nil
}
