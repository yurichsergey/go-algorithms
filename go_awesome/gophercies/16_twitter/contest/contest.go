package contest

import (
	"errors"
	rand "math/rand/v2"
	"sort"
	"strings"
)

// Merge returns a sorted list of unique, normalized usernames.
func Merge(oldUsers, newUsers []string) []string {
	unique := make(map[string]struct{}, len(oldUsers)+len(newUsers))
	for _, users := range [][]string{oldUsers, newUsers} {
		for _, user := range users {
			normalized := strings.ToLower(strings.TrimPrefix(strings.TrimSpace(user), "@"))
			if normalized != "" {
				unique[normalized] = struct{}{}
			}
		}
	}

	merged := make([]string, 0, len(unique))
	for user := range unique {
		merged = append(merged, user)
	}
	sort.Strings(merged)
	return merged
}

func Pick(users []string, n int) ([]string, error) {
	if n <= 0 || n > len(users) {
		return nil, errors.New("invalid number of users to pick")
	}
	result := make([]string, len(users))
	copy(result, users)
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result[:n], nil
}
