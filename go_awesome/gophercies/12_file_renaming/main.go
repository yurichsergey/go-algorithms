package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

type rule struct {
	re     *regexp.Regexp
	rename func(matches []string, total int) string
}

var rules = []rule{
	{
		// birthday_001.txt → birthday (1 of N).txt  (N = actual file count)
		re: regexp.MustCompile(`^birthday_(\d+)\.txt$`),
		rename: func(m []string, total int) string {
			n, err := strconv.Atoi(m[1])
			if err != nil {
				return ""
			}
			return fmt.Sprintf("birthday (%d of %d).txt", n, total)
		},
	},
	{
		// n_008.txt → 008 - n.txt
		re: regexp.MustCompile(`^n_(\d+)\.txt$`),
		rename: func(m []string, _ int) string {
			return fmt.Sprintf("%s - n.txt", m[1])
		},
	},
	{
		// christmas 2016 (3 of 100).txt → christmas 2016 - 003.txt
		re: regexp.MustCompile(`^(christmas \d+) \((\d+) of \d+\)\.txt$`),
		rename: func(m []string, _ int) string {
			n, err := strconv.Atoi(m[2])
			if err != nil {
				return ""
			}
			return fmt.Sprintf("%s - %03d.txt", m[1], n)
		},
	},
}

type candidate struct {
	path    string
	name    string
	ruleIdx int
	matches []string
}

func main() {
	root := "sample"
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	// Pass 1: collect all matching files
	var candidates []candidate
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		name := info.Name()
		for i, r := range rules {
			if m := r.re.FindStringSubmatch(name); m != nil {
				candidates = append(candidates, candidate{path, name, i, m})
				break
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Count matches per rule so rename funcs can use the real total
	counts := make(map[int]int)
	for _, c := range candidates {
		counts[c.ruleIdx]++
	}

	// Pass 2: rename
	for _, c := range candidates {
		newName := rules[c.ruleIdx].rename(c.matches, counts[c.ruleIdx])
		if newName == "" {
			fmt.Printf("skip %s: could not build new name\n", c.name)
			continue
		}
		newPath := filepath.Join(filepath.Dir(c.path), newName)
		if _, err := os.Stat(newPath); err == nil {
			fmt.Printf("skip %s: %s already exists\n", c.name, newName)
			continue
		}
		fmt.Printf("%s → %s\n", c.name, newName)
		if err := os.Rename(c.path, newPath); err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
}
