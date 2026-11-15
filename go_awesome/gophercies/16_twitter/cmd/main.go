package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/yurichsergey/go-algorithms/go_awesome/gophercies/16_twitter/contest"
	"github.com/yurichsergey/go-algorithms/go_awesome/gophercies/16_twitter/store"
	"github.com/yurichsergey/go-algorithms/go_awesome/gophercies/16_twitter/twitter"
)

func main() {
	var tweetID string
	flag.StringVar(&tweetID, "tweet", "", "Tweet ID to fetch retweeters for the contest")
	var filePath string
	flag.StringVar(&filePath, "file", "retweeters.json", "File path to save the usernames list")
	var pickCount int
	flag.IntVar(&pickCount, "pick", 0, "Number of winners to pick from the retweeters list")

	flag.Parse()

	if tweetID == "" {
		fmt.Fprintln(os.Stderr, "error: tweet ID is required")
		os.Exit(1)
	}

	fmt.Printf("tweet=%s, file=%s, pick=%d\n", tweetID, filePath, pickCount)

	usernamesSaved, err := store.Load(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading retweeters from file: %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("Loaded %d retweeters from file: %s\n", len(usernamesSaved), filePath)

	// url := "https://api.twitter.com"
	// token := os.Getenv("TWITTER_BEARER_TOKEN")

	url := "http://localhost:8080"
	token := "test-token" // test token for local testing

	client := twitter.NewClient(http.DefaultClient, url, token)
	usernamesNew, err := client.Retweeters(context.Background(), tweetID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error fetching retweeters from Twitter API: %v\n", err)
		os.Exit(3)
	}

	usernamesNew = contest.Merge(usernamesSaved, usernamesNew)

	err = store.Save(filePath, usernamesNew)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error saving retweeters to file: %v\n", err)
		os.Exit(4)
	}

	fmt.Printf("Saved %d retweeters to file: %s\n", len(usernamesNew), filePath)
}
