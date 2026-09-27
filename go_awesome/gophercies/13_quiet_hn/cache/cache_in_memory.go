package cache

import (
	"13_quiet_hn/reader"
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	stories  []reader.Item
	expireAt time.Time
}

func (c *Cache) GetStories(numStories int) ([]reader.Item, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if time.Now().Before(c.expireAt) {
		return c.stories, nil
	}

	stories, err := reader.GetStories(numStories)
	if err != nil {
		return nil, err
	}

	c.stories = stories
	c.expireAt = time.Now().Add(15 * time.Minute)
	return c.stories, nil
}
