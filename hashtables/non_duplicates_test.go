package hashtables

import (
	"csgtdsa/internal/assert"
	"testing"
)

func TestFirstNonDuplicate(t *testing.T) {
	input := "minimum"

	tests := []struct {
		name  string
		input string
		want  rune
	}{
		{
			name:  "FirstNonDuplicate",
			input: input,
			want:  'n',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FirstNonDuplicate(tt.input)

			assert.Equal(t, result, tt.want)
		})
	}
}
