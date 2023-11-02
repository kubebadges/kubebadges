package badges

import (
	"testing"
)

func TestFormatString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "hello-world",
			expected: "hello--world",
		},
		{
			input:    "foo_bar",
			expected: "foo__bar",
		},
		{
			input:    "hello world",
			expected: "hello_world",
		},
		{
			input:    "foo-bar_baz",
			expected: "foo--bar__baz",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := formatString(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %q, but got %q", tc.expected, actual)
			}
		})
	}
}
