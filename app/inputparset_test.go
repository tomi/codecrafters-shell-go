package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		input    string
		expected Command
	}{
		{
			input: "echo hello",
			expected: Command{
				Name: "echo",
				Args: []string{"hello"},
			},
		},
		{
			input: "echo 'hello'",
			expected: Command{
				Name: "echo",
				Args: []string{"hello"},
			},
		},
		{
			input: "echo 'hello''world'",
			expected: Command{
				Name: "echo",
				Args: []string{"hello", "world"},
			},
		},
		{
			input: "echo \"hello\"",
			expected: Command{
				Name: "echo",
				Args: []string{"hello"},
			},
		},
		{
			input: "echo \"hello\" \"world\"",
			expected: Command{
				Name: "echo",
				Args: []string{"hello", "world"},
			},
		},
	}

	for _, test := range tests {
		result, err := ParseInput(test.input)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if result.Name != test.expected.Name {
			t.Errorf("expected name %v, got %v", test.expected.Name, result.Name)
		}
		if !reflect.DeepEqual(result.Args, test.expected.Args) {
			t.Errorf("expected args %v, got %v", test.expected.Args, result.Args)
		}
	}
}
