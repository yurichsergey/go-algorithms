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

func TestPick(t *testing.T) {
	tests := []struct {
		name    string
		users   []string
		n       int
		wantLen int
		wantErr bool
	}{
		{
			name:    "pick 2 from 5",
			users:   []string{"alice", "bob", "charlie", "dave", "eve"},
			n:       2,
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "pick all users",
			users:   []string{"alice", "bob", "charlie"},
			n:       3,
			wantLen: 3,
			wantErr: false,
		},
		{
			name:    "pick more than available",
			users:   []string{"alice", "bob"},
			n:       3,
			wantLen: 0,
			wantErr: true,
		},
		{
			name:	"pick zero users",
			users:   []string{"alice", "bob"},
			n:       0,
			wantLen: 0,
			wantErr: true,
		},
		{
			name:    "pick negative users",
			users:   []string{"alice", "bob"},
			n:       -1,
			wantLen: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Pick(tt.users, tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pick() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("Pick() got length = %v, want length %v", len(got), tt.wantLen)
			}
		})
	}
}