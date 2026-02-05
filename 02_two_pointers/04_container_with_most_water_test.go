package _2_two_pointers

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "Example 1: Provided mixed heights",
			heights: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:    49, // (8-1) * min(8, 7) = 7 * 7 = 49
		},
		{
			name:    "Example 2: All same heights",
			heights: []int{2, 2, 2},
			want:    4, // (2-0) * 2 = 4
		},
		{
			name:    "Minimum length array",
			heights: []int{1, 1},
			want:    1,
		},
		{
			name:    "Increasing heights",
			heights: []int{1, 2, 3, 4, 5},
			want:    6, // indices 1 and 4: (4-1) * 2 = 6 or indices 2 and 4: (4-2) * 3 = 6
		},
		{
			name:    "Decreasing heights",
			heights: []int{5, 4, 3, 2, 1},
			want:    6,
		},
		{
			name:    "Zero height bars",
			heights: []int{0, 10, 0, 10, 0},
			want:    20, // indices 1 and 3: (3-1) * 10 = 20
		},
		{
			name:    "Large distance with small height vs small distance with large height",
			heights: []int{1, 100, 100, 1},
			want:    100, // (2-1) * 100 = 100
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.heights); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMaxArea allows you to compare the performance of different implementations
func BenchmarkMaxArea(b *testing.B) {
	heights := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		heights[i] = i % 100
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxArea(heights)
	}
}
