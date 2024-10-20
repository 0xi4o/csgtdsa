package sort

import (
	"csgtdsa/internal/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	nums := []int{73, 28, 91, 45, 12, 67, 36, 82, 19, 54, 3, 96, 40, 8, 61, 25, 79, 50, 14, 88}

	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "InsertionSort",
			input: nums,
			want:  []int{3, 8, 12, 14, 19, 25, 28, 36, 40, 45, 50, 54, 61, 67, 73, 79, 82, 88, 91, 96},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Insertion(tt.input)

			assert.CompareSlices(t, result, tt.want)
		})
	}
}
