package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello, Quiz!")

	csvFileP := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimitP := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	shuffleP := flag.Bool("shuffle", false, "shuffle the quiz in questions")
	flag.Parse()
	csvFile, timeLimit, shuffle := *csvFileP, *timeLimitP, *shuffleP

	fmt.Printf("Time limit: %d seconds\n", timeLimit)

	f, err := os.Open(csvFile)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s\n", err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("Failed to close the CSV file: %s\n", err)
		}
	}(f)

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read the CSV file: %s\n", err)
		return
	}
	fmt.Println("Successfully opened the CSV file. Loaded records:", len(records), "question records.")

	if shuffle {
		fmt.Printf("SHUFFLING QUESTIONS\n")
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}

	fmt.Print("Press 'Enter' to start quiz...")
	_, errStart := bufio.NewReader(os.Stdin).ReadString('\n')
	if errStart != nil {
		return
	}

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	doneCh := make(chan bool)

	score := 0
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for i, record := range records {
			fmt.Printf("Q%d: %s = ?\n", i+1, record[0])

			answer, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Failed to read the answer from %s: %s\n", record, err)
				return
			}

			answer = strings.TrimSpace(answer)
			if answer == record[1] {
				fmt.Printf("Correct! You answered: %s\n", answer)
				score++
			} else {
				fmt.Printf("Wrong answer! You answered: %s\n", answer)
			}
		}
		doneCh <- true
	}()

	select {
	case <-doneCh:
		fmt.Printf("Quiz completed! Your score is %d out of %d\n", score, len(records))
	case <-timer.C:
		fmt.Printf("Time limit exceeded! Your score is %d out of %d\n\n", score, len(records))
	}
}
