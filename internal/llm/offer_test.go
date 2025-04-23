package llm

import (
	"slices"
	"testing"
)

func TestFormatLLMResponse(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "Hyphen list",
			input: "- one\n- two",
			want:  []string{"one", "two"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := formatLLMOffer(tc.input)
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %q, but want %q", got, tc.want)
			}
		})

	}
}
