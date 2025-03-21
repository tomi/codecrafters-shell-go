package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Command
	}{
		{
			name:  "single word",
			input: "echo hello",
			expected: Command{
				Name: "echo",
				Args: []string{"hello"},
			},
		},
		{
			name:  "single word with singlequotes",
			input: "echo 'hello'",
			expected: Command{
				Name: "echo",
				Args: []string{"hello"},
			},
		},
		{
			name:  "single word with merged quotes",
			input: "echo 'hello''world'",
			expected: Command{
				Name: "echo",
				Args: []string{"helloworld"},
			},
		},
		{
			name:  "single word with doublequotes",
			input: "echo \"hello\"",
			expected: Command{
				Name: "echo",
				Args: []string{"hello"},
			},
		},
		{
			name:  "multiple words with doublequotes",
			input: "echo \"hello\" \"world\"",
			expected: Command{
				Name: "echo",
				Args: []string{"hello", "world"},
			},
		},
		{
			name:  "single arg with multiple spaces",
			input: "echo 'shell       hello'",
			expected: Command{
				Name: "echo",
				Args: []string{"shell       hello"},
			},
		},
		{
			name:  "singlequote in doublequotes",
			input: "echo \"hello's world\"",
			expected: Command{
				Name: "echo",
				Args: []string{"hello's world"},
			},
		},
	}

	for _, test := range tests {
		result, err := ParseInput(test.input)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if result.Name != test.expected.Name {
			t.Errorf("%s: expected name %v, got %v", test.name, test.expected.Name, result.Name)
		}
		if !reflect.DeepEqual(result.Args, test.expected.Args) {
			t.Errorf("%s: expected args %v, got %v", test.name, test.expected.Args, result.Args)
		}
	}
}
