package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("got item #%d %s", itemCount, item)
		}
	}()
	return out
}
