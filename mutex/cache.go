package main

import (
	"time"
	"sync"
	"fmt"
)

// TTL CACHE
type entry struct {
	value     string
	expiresAt time.Time
}

//Cache is a TTL cache that is safe for concurrent use
type Cache struct {
	entries map[string]*entry
	mx sync.RWMutex // read write mutex
	//TODO: protect this for concurrent use!
}

//NewCache constructs a new Cache object
func NewCache() *Cache {
	c := &Cache {
		entries: map[string]*entry{},
	}
	// spin up janitor on separate thread
	go c.janitor()

	return c
}

//Set adds a key/value to the cache
func (c *Cache) Set(key string, value string, timeToLive time.Duration) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.entries[key] = &entry{value, time.Now().Add(timeToLive)}
}

//Get gets the value associated with a key
func (c *Cache) Get(key string) (string, bool) {
	//TODO: implement this
	c.mx.RLock() // read locks
	defer c.mx.RUnlock() // read locks
	entry, found := c.entries[key]
	if found {
		return entry.value, true
	}

	return "", false
}

func (c *Cache) janitor() {
	for {
		time.Sleep(time.Second)
		now := time.Now()
		c.mx.Lock()
		fmt.Println("janitor is running")
		for key, entry := range c.entries {
			if entry.expiresAt.Before(now) {
				fmt.Printf("purging key %s\n", key)
				delete(c.entries, key)
			}
		}
		c.mx.Unlock()
	}
}
