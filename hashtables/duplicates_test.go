package hashtables

import (
	"csgtdsa/internal/assert"
	"testing"
)

func TestFirstDuplicate(t *testing.T) {
	input := "abcdcef"

	tests := []struct {
		name  string
		input string
		want  rune
	}{
		{
			name:  "FirstDuplicate",
			input: input,
			want:  'c',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FirstDuplicate(tt.input)

			assert.Equal(t, result, tt.want)
		})
	}
}
