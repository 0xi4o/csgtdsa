package hashtables

import (
	"csgtdsa/internal/assert"
	"testing"
)

func TestIntersection(t *testing.T) {
	nums1 := []int{73, 28, 91, 45, 12, 67, 36, 82, 19, 54, 3, 96, 40, 8, 61, 25, 79, 50, 14, 88}
	nums2 := []int{73, 28, 91, 45, 11, 67, 36, 82, 19, 54, 4, 96, 40, 8, 61, 25, 79, 51, 14, 88}

	tests := []struct {
		name   string
		input1 []int
		input2 []int
		want   []int
	}{
		{
			name:   "Intersection",
			input1: nums1,
			input2: nums2,
			want:   []int{73, 28, 91, 45, 67, 36, 82, 19, 54, 96, 40, 8, 61, 25, 79, 14, 88},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Intersection(tt.input1, tt.input2)

			assert.CompareSlices(t, result, tt.want)
		})
	}
}
