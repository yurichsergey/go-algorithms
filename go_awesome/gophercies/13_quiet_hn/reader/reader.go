package reader

import (
	"13_quiet_hn/hn"
	"errors"
	"math"
	"net/url"
	"strings"
	"sync"
	"time"
)

// Item is the same as the hn.Item, but adds the Host field
type Item struct {
	hn.Item
	Host string
}

type TemplateData struct {
	Stories []Item
	Time    time.Duration
}

func GetStories(numStories int) ([]Item, error) {
	var client hn.Client
	ids, err := client.TopItems()
	if err != nil {
		return nil, errors.New("failed to load top stories")
	}
	numToFetch := int(math.Ceil(float64(numStories) * 1.25))
	ids = ids[:numToFetch]

	results := make([]Item, len(ids))

	var wg sync.WaitGroup
	for i, id := range ids {
		wg.Add(1)
		go func(i, id int) {
			defer wg.Done()
			hnItem, err := client.GetItem(id)
			if err == nil {
				results[i] = parseHNItem(hnItem)
			}
		}(i, id)
	}
	wg.Wait()

	var stories []Item
	for _, it := range results {
		if isStoryLink(it) {
			stories = append(stories, it)
			if len(stories) >= numStories {
				break
			}
		}
	}
	return stories, nil
}

func isStoryLink(item Item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn.Item) Item {
	ret := Item{Item: hnItem}
	targetUrl, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(targetUrl.Hostname(), "www.")
	}
	return ret
}
