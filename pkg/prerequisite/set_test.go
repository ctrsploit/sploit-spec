package prerequisite

import (
	"errors"
	"reflect"
	"testing"
)

// mockSet is a mock implementation of the Set interface for testing purposes.
// It allows us to control the behavior of a prerequisite check for our tests.
type mockSet struct {
	name           string
	checkSatisfied bool
	checkErr       error
	checkCallCount int
}

// Check simulates checking a prerequisite. It returns configured values and counts how many times it was called.
// This is crucial for testing memoization (caching).
func (m *mockSet) Check() (bool, error) {
	m.checkCallCount++
	return m.checkSatisfied, m.checkErr
}

// Range for a mockSet, being a leaf node, returns a channel that yields the mockSet instance itself.
func (m *mockSet) Range() <-chan Set {
	ch := make(chan Set)
	go func() {
		defer close(ch)
		ch <- m
	}()
	return ch
}

// GetName returns the mock's configured name.
func (m *mockSet) GetName() string {
	return m.name
}

// Output is a no-op for the mock, as we are not testing the logging/output functionality.
func (m *mockSet) Output() {
	// No operation needed for testing logic.
}

// TestSetNot covers the functionality of the SetNot struct.
func TestSetNot(t *testing.T) {
	// Test case for the GetName method.
	t.Run("GetName", func(t *testing.T) {
		mockA := &mockSet{name: "A"}
		notA := Not(mockA)
		expected := "!(A)"
		if got := notA.GetName(); got != expected {
			t.Errorf("Not.GetName() = %v, want %v", got, expected)
		}
	})

	// Test cases for the Check method, covering logic and error propagation.
	t.Run("Check", func(t *testing.T) {
		testErr := errors.New("test error")

		testCases := []struct {
			name              string
			innerSatisfied    bool
			innerErr          error
			expectedSatisfied bool
			expectedErr       error
		}{
			{"InnerTrue", true, nil, false, nil},
			{"InnerFalse", false, nil, true, nil},
			{"InnerTrueWithError", true, testErr, false, testErr},
			{"InnerFalseWithError", false, testErr, true, testErr},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				mock := &mockSet{checkSatisfied: tc.innerSatisfied, checkErr: tc.innerErr}
				notSet := Not(mock)
				satisfied, err := notSet.Check()

				if satisfied != tc.expectedSatisfied {
					t.Errorf("Expected satisfied %v, got %v", tc.expectedSatisfied, satisfied)
				}
				if !errors.Is(err, tc.expectedErr) {
					t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
				}
			})
		}
	})

	// Test case for memoization (caching) of the Check result.
	t.Run("CheckMemoization", func(t *testing.T) {
		mock := &mockSet{checkSatisfied: true}
		notSet := Not(mock)

		notSet.Check() // First call
		notSet.Check() // Second call

		if mock.checkCallCount != 1 {
			t.Errorf("Inner Set.Check() should be called once, but was called %d times", mock.checkCallCount)
		}
	})

	// Test case for the Range method.
	t.Run("Range", func(t *testing.T) {
		mockA := &mockSet{name: "A"}
		notA := Not(mockA)

		var results []Set
		for s := range notA.Range() {
			results = append(results, s)
		}

		if len(results) != 1 {
			t.Fatalf("Expected Range to yield 1 item, but got %d", len(results))
		}
		if results[0] != notA {
			t.Errorf("Expected Range to yield the inner set, but it didn't")
		}
	})
}

// TestSetAnd covers the functionality of the SetAnd struct.
func TestSetAnd(t *testing.T) {
	// Test case for the GetName method.
	t.Run("GetName", func(t *testing.T) {
		mockA := &mockSet{name: "A"}
		mockB := &mockSet{name: "B"}
		andSet := And(mockA, mockB)
		expected := "(A) && (B)"
		if got := andSet.GetName(); got != expected {
			t.Errorf("And.GetName() = %v, want %v", got, expected)
		}
	})

	// Test cases for the Check method.
	t.Run("Check", func(t *testing.T) {
		testErr1 := errors.New("error 1")
		testErr2 := errors.New("error 2")

		testCases := []struct {
			name              string
			sets              []Set
			expectedSatisfied bool
			expectedErr       error
		}{
			{"TrueAndTrue", []Set{&mockSet{checkSatisfied: true}, &mockSet{checkSatisfied: true}}, true, nil},
			{"TrueAndFalse", []Set{&mockSet{checkSatisfied: true}, &mockSet{checkSatisfied: false}}, false, nil},
			{"FalseAndTrue", []Set{&mockSet{checkSatisfied: false}, &mockSet{checkSatisfied: true}}, false, nil},
			{"FalseAndFalse", []Set{&mockSet{checkSatisfied: false}, &mockSet{checkSatisfied: false}}, false, nil},
			{"Empty", []Set{}, true, nil},
			{"WithNil", []Set{&mockSet{checkSatisfied: true}, nil, &mockSet{checkSatisfied: true}}, true, nil},
			{"TrueAndErrorFalse", []Set{&mockSet{checkSatisfied: true}, &mockSet{checkSatisfied: false, checkErr: testErr1}}, false, testErr1},
			{"ErrorTrueAndTrue", []Set{&mockSet{checkSatisfied: true, checkErr: testErr1}, &mockSet{checkSatisfied: true}}, true, testErr1},
			{"ErrorFalseAndErrorFalse", []Set{&mockSet{checkSatisfied: false, checkErr: testErr1}, &mockSet{checkSatisfied: false, checkErr: testErr2}}, false, errors.Join(testErr1, testErr2)},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				andSet := And(tc.sets...)
				satisfied, err := andSet.Check()

				if satisfied != tc.expectedSatisfied {
					t.Errorf("Expected satisfied %v, got %v", tc.expectedSatisfied, satisfied)
				}
				// Compare error strings because errors.Join creates a new error instance which fails direct comparison.
				if (err == nil && tc.expectedErr != nil) || (err != nil && tc.expectedErr == nil) || (err != nil && tc.expectedErr != nil && err.Error() != tc.expectedErr.Error()) {
					t.Errorf("Expected error '%v', got '%v'", tc.expectedErr, err)
				}
			})
		}
	})

	// Test case for memoization (caching) of the Check result.
	t.Run("CheckMemoization", func(t *testing.T) {
		mockA := &mockSet{checkSatisfied: true}
		mockB := &mockSet{checkSatisfied: true}
		andSet := And(mockA, mockB)

		andSet.Check() // First call
		andSet.Check() // Second call

		if mockA.checkCallCount != 1 {
			t.Errorf("Inner Set A Check() should be called once, but was called %d times", mockA.checkCallCount)
		}
		if mockB.checkCallCount != 1 {
			t.Errorf("Inner Set B Check() should be called once, but was called %d times", mockB.checkCallCount)
		}
	})

	// Test case for the Range method.
	t.Run("Range", func(t *testing.T) {
		mockA := &mockSet{name: "A"}
		mockB := &mockSet{name: "B"}
		// Note: And.Range is not recursive based on the provided implementation. It yields its direct children.
		orSet := Or(mockB)
		andSet := And(mockA, orSet, nil)

		var results []Set
		for s := range andSet.Range() {
			results = append(results, s)
		}

		expected := []Set{mockA, orSet}
		if !reflect.DeepEqual(results, expected) {
			t.Errorf("Expected Range to yield %v, but got %v", expected, results)
		}
	})
}

// TestSetOr covers the functionality of the SetOr struct.
func TestSetOr(t *testing.T) {
	// Test case for the GetName method.
	t.Run("GetName", func(t *testing.T) {
		mockA := &mockSet{name: "A"}
		mockB := &mockSet{name: "B"}
		orSet := Or(mockA, mockB)
		expected := "(A) || (B)"
		if got := orSet.GetName(); got != expected {
			t.Errorf("Or.GetName() = %v, want %v", got, expected)
		}
	})

	// Test cases for the Check method.
	t.Run("Check", func(t *testing.T) {
		testErr1 := errors.New("error 1")
		testErr2 := errors.New("error 2")

		testCases := []struct {
			name              string
			sets              []Set
			expectedSatisfied bool
			expectedErr       error
		}{
			{"TrueOrTrue", []Set{&mockSet{checkSatisfied: true}, &mockSet{checkSatisfied: true}}, true, nil},
			{"TrueOrFalse", []Set{&mockSet{checkSatisfied: true}, &mockSet{checkSatisfied: false}}, true, nil},
			{"FalseOrTrue", []Set{&mockSet{checkSatisfied: false}, &mockSet{checkSatisfied: true}}, true, nil},
			{"FalseOrFalse", []Set{&mockSet{checkSatisfied: false}, &mockSet{checkSatisfied: false}}, false, nil},
			{"Empty", []Set{}, false, nil},
			{"WithNil", []Set{&mockSet{checkSatisfied: false}, nil, &mockSet{checkSatisfied: true}}, true, nil},
			{"FalseOrErrorTrue", []Set{&mockSet{checkSatisfied: false}, &mockSet{checkSatisfied: true, checkErr: testErr1}}, true, testErr1},
			{"ErrorFalseOrFalse", []Set{&mockSet{checkSatisfied: false, checkErr: testErr1}, &mockSet{checkSatisfied: false}}, false, testErr1},
			{"ErrorFalseOrErrorFalse", []Set{&mockSet{checkSatisfied: false, checkErr: testErr1}, &mockSet{checkSatisfied: false, checkErr: testErr2}}, false, errors.Join(testErr1, testErr2)},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				orSet := Or(tc.sets...)
				satisfied, err := orSet.Check()

				if satisfied != tc.expectedSatisfied {
					t.Errorf("Expected satisfied %v, got %v", tc.expectedSatisfied, satisfied)
				}
				// Compare error strings because errors.Join creates a new error instance.
				if (err == nil && tc.expectedErr != nil) || (err != nil && tc.expectedErr == nil) || (err != nil && tc.expectedErr != nil && err.Error() != tc.expectedErr.Error()) {
					t.Errorf("Expected error '%v', got '%v'", tc.expectedErr, err)
				}
			})
		}
	})

	// Test case for memoization (caching) of the Check result.
	t.Run("CheckMemoization", func(t *testing.T) {
		mockA := &mockSet{checkSatisfied: false}
		mockB := &mockSet{checkSatisfied: false}
		orSet := Or(mockA, mockB)

		orSet.Check() // First call
		orSet.Check() // Second call

		if mockA.checkCallCount != 1 {
			t.Errorf("Inner Set A Check() should be called once, but was called %d times", mockA.checkCallCount)
		}
		if mockB.checkCallCount != 1 {
			t.Errorf("Inner Set B Check() should be called once, but was called %d times", mockB.checkCallCount)
		}
	})

	// Test case for the Range method.
	t.Run("Range", func(t *testing.T) {
		mockA := &mockSet{name: "A"}
		mockB := &mockSet{name: "B"}
		// Note: Or.Range is not recursive based on the provided implementation. It yields its direct children.
		andSet := And(mockB)
		orSet := Or(mockA, andSet, nil)

		var results []Set
		for s := range orSet.Range() {
			results = append(results, s)
		}

		expected := []Set{mockA, andSet}
		if !reflect.DeepEqual(results, expected) {
			t.Errorf("Expected Range to yield %v, but got %v", expected, results)
		}
	})
}

// TestComplexNesting demonstrates the behavior of nested structures.
func TestComplexNesting(t *testing.T) {
	// Represents the expression: ! (A && (B || C))
	// A = true, B = false, C = true
	// B || C -> true
	// A && (B || C) -> true && true -> true
	// ! (A && (B || C)) -> !true -> false
	mockA := &mockSet{name: "A", checkSatisfied: true}
	mockB := &mockSet{name: "B", checkSatisfied: false}
	mockC := &mockSet{name: "C", checkSatisfied: true}

	orBC := Or(mockB, mockC)
	andABC := And(mockA, orBC)
	root := Not(andABC)

	// Test GetName for the complex structure.
	t.Run("GetName", func(t *testing.T) {
		expectedName := "!((A) && ((B) || (C)))"
		if name := root.GetName(); name != expectedName {
			t.Errorf("GetName() = %q, want %q", name, expectedName)
		}
	})

	// Test Check for the complex structure, including memoization.
	t.Run("Check", func(t *testing.T) {
		satisfied, err := root.Check()
		if err != nil {
			t.Fatalf("Check() returned unexpected error: %v", err)
		}
		if satisfied {
			t.Errorf("Check() returned satisfied=true, want false")
		}

		// Verify memoization works through the layers. Each leaf should be checked only once.
		if mockA.checkCallCount != 1 || mockB.checkCallCount != 1 || mockC.checkCallCount != 1 {
			t.Errorf("Expected each mock Check to be called once, got A:%d, B:%d, C:%d",
				mockA.checkCallCount, mockB.checkCallCount, mockC.checkCallCount)
		}
		// A second check on the root should not re-evaluate any children.
		root.Check()
		if mockA.checkCallCount != 1 || mockB.checkCallCount != 1 || mockC.checkCallCount != 1 {
			t.Errorf("Expected each mock Check to still be called once after second root check, got A:%d, B:%d, C:%d",
				mockA.checkCallCount, mockB.checkCallCount, mockC.checkCallCount)
		}
	})

	// Test Range for the complex structure.
	t.Run("Range", func(t *testing.T) {
		// As noted, Not.Range is recursive, but And/Or are not in the provided code.
		// Therefore, Not(And(A, Or(B, C))) will yield the direct children of And(...),
		// which are A and Or(B, C).
		var results []Set
		for s := range root.Range() {
			results = append(results, s)
		}
		expected := []Set{root}
		if !reflect.DeepEqual(results, expected) {
			t.Errorf("Expected Range to yield %v, but got %v", expected, results)
		}
	})
}
