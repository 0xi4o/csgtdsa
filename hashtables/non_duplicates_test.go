package hashtables

import (
	"csgtdsa/internal/assert"
	"testing"
)

func TestNonDuplicates(t *testing.T) {
	input := "minimum"

	tests := []struct {
		name  string
		input string
		want  rune
	}{
		{
			name:  "NonDuplicates",
			input: input,
			want:  'n',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NonDuplicates(tt.input)

			assert.CompareSlices(t, result, tt.want)
		})
	}
}
