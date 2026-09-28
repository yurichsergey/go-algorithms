package reader

import (
	"13_quiet_hn/hn"
	"errors"
	"math"
	"net/url"
	"strings"
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

type result struct {
	idx  int
	Item Item
}
type job struct {
	idx int
	id  int
}

func GetStories(numStories int) ([]Item, error) {
	var client hn.Client
	ids, err := client.TopItems()
	if err != nil {
		return nil, errors.New("failed to load top stories")
	}
	numToFetch := int(math.Ceil(float64(numStories) * 1.25))
	ids = ids[:numToFetch]

	jobs := make(chan job, len(ids))       // this buffered pipe is empty now, holds up to len(ids)
	results := make(chan result, len(ids)) // this buffered pipe is empty now, holds up to len(ids)

	for w := 0; w < 5; w++ {
		go func() {
			for job := range jobs { // worker pulls jobs from the pipe until a channel is empty+closed
				hnItem, err := client.GetItem(job.id)
				if err == nil {
					results <- result{job.idx, parseHNItem(hnItem)} // pushes a result into the result pipe
				}
			}
		}()
	}

	for i, id := range ids {
		jobs <- job{i, id} // pushes a job into the jobs pipe after initializing goroutines
	}
	close(jobs) // closing the jobs pipe, after filling, but before it will be closed

	items := make([]Item, len(ids))
	for i := 0; i < len(ids); i++ {
		r := <-results // pull the one result out - blocks until one is ready
		items[r.idx] = r.Item
	}

	var stories []Item
	for _, it := range items {
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
