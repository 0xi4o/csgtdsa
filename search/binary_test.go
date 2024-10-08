package search

import (
	"csgtdsa/internal/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{3, 7, 12, 19, 24, 31, 35, 42, 48, 53, 59, 64, 70, 75, 81, 86, 90, 93, 97, 99}

	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{
			name:   "Target Exists",
			nums:   nums,
			target: 64,
			want:   true,
		},
		{
			name:   "Target Does Not Exist",
			nums:   nums,
			target: 123,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Binary(tt.nums, tt.target)

			assert.Equal(t, result, tt.want)
		})
	}
}
