package maps

type Cache struct {
	data map[string]int
}

func NewCache() *Cache {
	// TODO: initialize the cache
	// Read README.md for the instructions
	mp := make(map[string]int)
	cache := Cache{data: mp}
	return &cache
}

func (c *Cache) Set(key string, value int) {
	c.data[key] = value
}

func (c *Cache) Get(key string) (int, bool) {
	exists, ok := c.data[key]
	if ok {
		return exists, true
	}
	return 0, false
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}

func (c *Cache) Count() int {
	return len(c.data)
}

func (c *Cache) AllKeys() []string {
	if len(c.data) == 0 {
		return []string{}
	}
	keys := make([]string, 0, len(c.data))
	for k := range c.data {
		keys = append(keys, k)
	}
	return keys
}

func (c *Cache) RemoveBelow(limit int) {
	for k, v := range c.data {
		if v < limit {
			delete(c.data, k)
		}
	}
}
