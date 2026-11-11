package mutexes

import "sync"

// TODO:Implement producer-consumer problem using mutexes
// Read README.md for the instructions
type Counter struct {
	mu    sync.Mutex
	items int
}

func (c *Counter) Produce(amount int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items += amount
}

func (c *Counter) Consume(amount int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items -= amount
}

func (c *Counter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.items
}
