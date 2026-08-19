package main

import (
	"reflect"
	"testing"
)

func TestExtractUserInput(t *testing.T) {
	type funcValues struct {
		cmd       string
		args      []string
		wantError bool
	}

	testCases := []struct {
		name     string
		input    []string
		expected funcValues
	}{
		{name: "empty input", input: []string{}, expected: funcValues{cmd: "", args: nil, wantError: true}},
		{name: "single word", input: []string{"foo"}, expected: funcValues{cmd: "foo", args: nil, wantError: false}},
		{name: "multi word", input: []string{"foo", "bar", "foo"}, expected: funcValues{cmd: "foo", args: []string{"bar", "foo"}, wantError: false}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualCmd, actualArgs, actualErr := extractUserInput(tc.input)

			hasErr := actualErr != nil
			if hasErr != tc.expected.wantError {
				t.Errorf("wantError = %v, got err = %v", tc.expected.wantError, actualErr)
			}

			if actualCmd != tc.expected.cmd {
				t.Errorf("cmd: expected %q, got %q", tc.expected.cmd, actualCmd)
			}

			if !reflect.DeepEqual(actualArgs, tc.expected.args) {
				t.Errorf("args: expected %v, got %v", tc.expected.args, actualArgs)
			}
		})
	}
}
