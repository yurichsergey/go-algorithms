package contest

import (
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
