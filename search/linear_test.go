package search

import (
	"csgtdsa/internal/assert"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	nums := []int{73, 28, 91, 45, 12, 67, 36, 82, 19, 54, 3, 96, 40, 8, 61, 25, 79, 50, 14, 88}

	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{
			name:   "Target Exists",
			nums:   nums,
			target: 25,
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
			result := Linear(tt.nums, tt.target)

			assert.Equal(t, result, tt.want)
		})
	}
}
