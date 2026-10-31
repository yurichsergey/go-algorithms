package contest

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		oldUsers []string
		newUsers []string
		want     []string
	}{
		{
			name:     "basic merge",
			oldUsers: []string{"@Alice", "Bob", "charlie"},
			newUsers: []string{"@bob", "dave", "Eve"},
			want:     []string{"alice", "bob", "charlie", "dave", "eve"},
		},
		{
			name:     "with spaces and empty strings",
			oldUsers: []string{"  @Alice  ", "", "Bob"},
			newUsers: []string{"@bob", "  ", "Eve"},
			want:     []string{"alice", "bob", "eve"},
		},
		{
			name:     "all duplicates",
			oldUsers: []string{"@Alice", "@Bob"},
			newUsers: []string{"@alice", "@bob"},
			want:     []string{"alice", "bob"},
		},
		{
			name:     "no users",
			oldUsers: []string{},
			newUsers: []string{},
			want:     []string{},
		},
		{
			name:     "only new users",
			oldUsers: []string{},
			newUsers: []string{"@Alice", "@Bob"},
			want:     []string{"alice", "bob"},
		},
		{
			name:     "only old users",
			oldUsers: []string{"@Alice", "@Bob"},
			newUsers: []string{},
			want:     []string{"alice", "bob"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.oldUsers, tt.newUsers)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
