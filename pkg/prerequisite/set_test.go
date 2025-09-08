package prerequisite

import (
	"errors"
	"testing"
)

// mockSet is a mock implementation of the Set interface for testing purposes.
// We can preset the boolean and error it returns when GetSatisfied is called.
type mockSet struct {
	satisfied bool
	err       error
}

// GetSatisfied implements the Set interface, returning the preset satisfied and err values.
func (m mockSet) Check() (bool, error) {
	return m.satisfied, m.err
}

// TestSetAnd tests the And function and the GetSatisfied method of the SetAnd struct.
func TestSetAnd(t *testing.T) {
	// Define some reusable mockSet instances
	trueSet := mockSet{satisfied: true}
	falseSet := mockSet{satisfied: false}
	errorSet := mockSet{err: errors.New("mock error")}

	// Define the test case table
	testCases := []struct {
		name          string // Description of the test case
		input         Set    // The input Set
		expectedSatis bool   // The expected satisfied result
		expectErr     bool   // Whether an error is expected
	}{
		{
			name:          "Empty set should return true",
			input:         And(),
			expectedSatis: true,
			expectErr:     false,
		},
		{
			name:          "Single true set should return true",
			input:         And(trueSet),
			expectedSatis: true,
			expectErr:     false,
		},
		{
			name:          "Single false set should return false",
			input:         And(falseSet),
			expectedSatis: false,
			expectErr:     false,
		},
		{
			name:          "All sets being true should return true",
			input:         And(trueSet, trueSet, trueSet),
			expectedSatis: true,
			expectErr:     false,
		},
		{
			name:          "Any set being false should return false",
			input:         And(trueSet, falseSet, trueSet),
			expectedSatis: false,
			expectErr:     false,
		},
		{
			name:          "All sets being false should return false",
			input:         And(falseSet, falseSet),
			expectedSatis: false,
			expectErr:     false,
		},
		{
			name:          "Should return false and error immediately on encountering an error set",
			input:         And(trueSet, errorSet, falseSet), // The set after errorSet should not be evaluated
			expectedSatis: false,
			expectErr:     true,
		},
		{
			name:          "Nested logic: And(true, Or(true, false)) should return true",
			input:         And(trueSet, Or(trueSet, falseSet)),
			expectedSatis: true,
			expectErr:     false,
		},
		{
			name:          "Nested logic: And(true, Or(false, false)) should return false",
			input:         And(trueSet, Or(falseSet, falseSet)),
			expectedSatis: false,
			expectErr:     false,
		},
	}

	// Iterate through and execute all test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			satisfied, err := tc.input.Check()

			// Check if the error matches the expectation
			if tc.expectErr {
				if err == nil {
					t.Errorf("expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error, but got: %v", err)
				}
			}

			// Check if the boolean result matches the expectation
			if satisfied != tc.expectedSatis {
				t.Errorf("expected satisfied to be %v, but got %v", tc.expectedSatis, satisfied)
			}
		})
	}
}

// TestSetOr tests the Or function and the GetSatisfied method of the SetOr struct.
func TestSetOr(t *testing.T) {
	// Define some reusable mockSet instances
	trueSet := mockSet{satisfied: true}
	falseSet := mockSet{satisfied: false}
	errorSet := mockSet{err: errors.New("mock error")}

	// Define the test case table
	testCases := []struct {
		name          string // Description of the test case
		input         Set    // The input Set
		expectedSatis bool   // The expected satisfied result
		expectErr     bool   // Whether an error is expected
	}{
		{
			name:          "Empty set should return false",
			input:         Or(),
			expectedSatis: false,
			expectErr:     false,
		},
		{
			name:          "Single true set should return true",
			input:         Or(trueSet),
			expectedSatis: true,
			expectErr:     false,
		},
		{
			name:          "Single false set should return false",
			input:         Or(falseSet),
			expectedSatis: false,
			expectErr:     false,
		},
		{
			name:          "Any set being true should return true",
			input:         Or(falseSet, trueSet, falseSet),
			expectedSatis: true,
			expectErr:     false,
		},
		{
			name:          "All sets being true should return true",
			input:         Or(trueSet, trueSet),
			expectedSatis: true,
			expectErr:     false,
		},
		{
			name:          "All sets being false should return false",
			input:         Or(falseSet, falseSet, falseSet),
			expectedSatis: false,
			expectErr:     false,
		},
		{
			name:          "Should return false and error immediately on encountering an error set",
			input:         Or(falseSet, errorSet, trueSet), // The set after errorSet should not be evaluated
			expectedSatis: false,
			expectErr:     true,
		},
		{
			name:          "Nested logic: Or(false, And(true, true)) should return true",
			input:         Or(falseSet, And(trueSet, trueSet)),
			expectedSatis: true,
			expectErr:     false,
		},
		{
			name:          "Nested logic: Or(false, And(true, false)) should return false",
			input:         Or(falseSet, And(trueSet, falseSet)),
			expectedSatis: false,
			expectErr:     false,
		},
	}

	// Iterate through and execute all test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			satisfied, err := tc.input.Check()

			// Check if the error matches the expectation
			if tc.expectErr {
				if err == nil {
					t.Errorf("expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error, but got: %v", err)
				}
			}

			// Check if the boolean result matches the expectation
			if satisfied != tc.expectedSatis {
				t.Errorf("expected satisfied to be %v, but got %v", tc.expectedSatis, satisfied)
			}
		})
	}
}
