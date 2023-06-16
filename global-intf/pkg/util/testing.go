package util

import "testing"

// TestCase represents a test case for a method
type TestCase[Input any, Output any] struct {
	Name     string
	Input    *Input // only accept pointer through generic.
	Expected *Output
}

// DoTestCases runs the test cases
func DoTestCases[Input, Output any](t *testing.T, tcs []*TestCase[Input, Output], fn func(t *testing.T, tc *TestCase[Input, Output])) {
	// INFO: register DoTestCases as the helper function of testing.
	t.Helper()

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			fn(t, tc)
		})
	}
}
