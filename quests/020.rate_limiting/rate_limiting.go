package rate_limiting

import (
	"fmt"
	"time"
)

func PingPong(requests []string) {
	ch := make(chan time.Time, 4)

	// Pre-fill the bucket with four tokens.
	for i := 0; i < 4; i++ {
		ch <- time.Now()
	}

	// Refill the bucket every 100ms.
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	go func() {
		for t := range ticker.C {
			ch <- t
		}
	}()

	// Buffer all incoming requests.
	requestCh := make(chan string, len(requests))

	for _, req := range requests {
		requestCh <- req
	}
	close(requestCh)

	for req := range requestCh {
		// Wait until a token is available.
		<-ch

		response := "Pong"
		if req == "Pong" {
			response = "Ping"
		}

		fmt.Println(response, time.Now().Format("15:04:05.000"))
	}
}
