package worker_pool

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

const WORKER_COUNT = 5

func worker(id int, inputs <-chan string, outputs chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for input := range inputs {
		time.Sleep(time.Second) // optional simulated work

		upper := strings.ToUpper(input)
		outputs <- upper + "-" + strconv.Itoa(id)
	}
}

func WorkerPool(strs []string) []string {
	inputs := make(chan string, len(strs))
	outputs := make(chan string, len(strs))

	var wg sync.WaitGroup

	wg.Add(WORKER_COUNT)

	for i := 1; i <= WORKER_COUNT; i++ {
		go worker(i, inputs, outputs, &wg)
	}

	for _, str := range strs {
		inputs <- str
	}

	close(inputs)

	wg.Wait()
	close(outputs)

	results := make([]string, 0, len(strs))

	for result := range outputs {
		results = append(results, result)
	}

	return results
}
