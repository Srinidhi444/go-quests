package tickers

import (
	"fmt"
	"time"
)

func Ticker() {
	helloTicker := time.NewTicker(500 * time.Millisecond)
	worldTicker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(3 * time.Second)

	defer helloTicker.Stop()
	defer worldTicker.Stop()
	defer timer.Stop()

	for {
		select {
		case <-helloTicker.C:
			fmt.Println("hello")

		case <-worldTicker.C:
			fmt.Println("world")

		case <-timer.C:
			return
		}
	}
}
