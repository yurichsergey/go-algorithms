package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"16_twitter/contest"
	"16_twitter/store"
	"16_twitter/twitter"
)

func main() {
	var tweetID string
	flag.StringVar(&tweetID, "tweet", "", "Tweet ID to fetch retweeters for the contest")
	var filePath string
	flag.StringVar(&filePath, "file", "retweeters.json", "File path to save the usernames list")
	var pickCount int
	flag.IntVar(&pickCount, "pick", 0, "Number of winners to pick from the retweeters list")

	var every time.Duration
	flag.DurationVar(&every, "every", 5*time.Minute, "How often to poll")
	var runFor time.Duration
	flag.DurationVar(&runFor, "for", 0, "Total polling time; 0 disables polling")

	flag.Parse()

	if tweetID == "" {
		fmt.Fprintln(os.Stderr, "error: tweet ID is required")
		os.Exit(1)
	}
	if pickCount < 0 {
		fmt.Fprintln(os.Stderr, "error: pick count cannot be negative")
		os.Exit(1)
	}

	if every <= 0 {
		fmt.Fprintln(os.Stderr, "error: polling interval must be greater than zero")
		os.Exit(1)
	}

	if runFor < 0 {
		fmt.Fprintln(os.Stderr, "error: total polling time cannot be negative")
		os.Exit(1)
	}

	fmt.Printf("tweet=%s, file=%s, pick=%d\n", tweetID, filePath, pickCount)
	fmt.Printf("every=%s, for=%s\n", every, runFor)

	usernamesSaved, err := store.Load(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading retweeters from file: %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("Loaded %d retweeters from file: %s\n", len(usernamesSaved), filePath)

	if err := godotenv.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "error loading .env file: %v\n", err)
		os.Exit(1)
	}

	url := os.Getenv("TWITTER_BEARER_URL")
	token := os.Getenv("TWITTER_BEARER_TOKEN")

	if url == "" || token == "" {
		fmt.Fprintln(os.Stderr, "error: TWITTER_BEARER_URL and TWITTER_BEARER_TOKEN must be set in .env file")
		os.Exit(1)
	}

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

	if pickCount > 0 {
		winners, err := contest.Pick(usernamesNew, pickCount)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error picking winners: %v\n", err)
			os.Exit(5)
		}
		fmt.Printf("Picked %d winners:\n", len(winners))
		for _, winner := range winners {
			fmt.Println(winner)
		}
	}
}
