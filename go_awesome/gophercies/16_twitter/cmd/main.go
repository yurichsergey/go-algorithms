package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
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

	var usernames []string
	if runFor == 0 {
		updated, err := updateRetweeters(context.Background(), filePath, tweetID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "poll: error updating retweeters: %v\n", err)
			os.Exit(1)
		}
		usernames = updated
	} else {
		signalCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()

		runCtx, cancel := context.WithTimeout(signalCtx, runFor)
		defer cancel()

		ticker := time.NewTicker(every)
		defer ticker.Stop()

		poll := func(ctx context.Context) {
			updated, err := updateRetweeters(ctx, filePath, tweetID)
			if err != nil {
				fmt.Fprintf(os.Stderr, "poll: error updating retweeters: %v\n", err)
				return // retry on the next tick
			}
			usernames = updated
		}

		poll(runCtx) // don't wait for the first tick, poll immediately

		pollingLoop:
			for {
				select {
				case <-runCtx.Done():
					// If the context is done, we exit the loop; distinguish between timeout and interrupt Ctrl+C
					fmt.Println("Polling finished.")
					break pollingLoop
				case <-ticker.C:
					poll(runCtx)
				}
			}

		finalCtx, finalCancel := context.WithTimeout(context.Background(), 15*time.Second)
		poll(finalCtx) // final poll to ensure we have the latest data
		finalCancel()
	}

	if pickCount > 0 {
		winners, err := contest.Pick(usernames, pickCount)
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

func updateRetweeters(ctx context.Context, filePath string, tweetID string) (usernames []string, err error) {
	usernamesSaved, err := store.Load(filePath)
	if err != nil {
		return nil, fmt.Errorf("error loading retweeters from file: %v\n", err)
	}

	fmt.Printf("Loaded %d retweeters from file: %s\n", len(usernamesSaved), filePath)

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v\n", err)
	}

	url := os.Getenv("TWITTER_BEARER_URL")
	token := os.Getenv("TWITTER_BEARER_TOKEN")

	if url == "" || token == "" {
		return nil, fmt.Errorf("error: TWITTER_BEARER_URL and TWITTER_BEARER_TOKEN must be set in .env file")
	}

	client := twitter.NewClient(http.DefaultClient, url, token)
	usernamesNew, err := client.Retweeters(ctx, tweetID)
	if err != nil {
		return nil, fmt.Errorf("error fetching retweeters from Twitter API: %v\n", err)
	}

	usernamesNew = contest.Merge(usernamesSaved, usernamesNew)

	err = store.Save(filePath, usernamesNew)
	if err != nil {
		return nil, fmt.Errorf("error saving retweeters to file: %w", err)
	}

	fmt.Printf("Saved %d retweeters to file: %s\n", len(usernamesNew), filePath)
	return usernamesNew, nil
}
