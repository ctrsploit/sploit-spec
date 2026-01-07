package prerequisite_test

import (
	"errors"
	"testing"

	"github.com/ctrsploit/sploit-spec/pkg/prerequisite"
)

// mockSet is a mock object that implements the prerequisite.Set interface for testing purposes.
// It allows us to control its return values and track the number of times its Check() method is called.
type mockSet struct {
	name       string // The name of the mock set.
	satisfied  bool   // The boolean value that the Check() method will return.
	err        error  // The error that the Check() method will return.
	checkCount int    // The number of times the Check() method has been called.
	checked    bool   // The mocked 'checked' state.
}

func (m *mockSet) GetChecked() bool {
	return m.checked
}

func (m *mockSet) Check() (bool, error) {
	m.checkCount++
	m.checked = true
	return m.satisfied, m.err
}

func (m *mockSet) Range() <-chan prerequisite.Set {
	ch := make(chan prerequisite.Set)
	go func() {
		defer close(ch)
		ch <- m
	}()
	return ch
}

func (m *mockSet) GetName() string {
	return m.name
}

// Output is a no-op implementation because we are not testing the specific content
// of log output, only the logic.
func (m *mockSet) Output() {}

// newMock is a helper function to create a mockSet.
func newMock(name string, satisfied bool, err error) *mockSet {
	return &mockSet{name: name, satisfied: satisfied, err: err}
}

// TestSetNot tests the Not logic.
func TestSetNot(t *testing.T) {
	mockErr := errors.New("mock error")

	tests := []struct {
		name          string
		inputSet      prerequisite.Set
		wantSatisfied bool
		wantErr       error
		wantName      string
	}{
		{
			name:          "Input True -> Output False",
			inputSet:      newMock("A", true, nil),
			wantSatisfied: false,
			wantErr:       nil,
			wantName:      "!(A)",
		},
		{
			name:          "Input False -> Output True",
			inputSet:      newMock("B", false, nil),
			wantSatisfied: true,
			wantErr:       nil,
			wantName:      "!(B)",
		},
		{
			name:          "Input with Error",
			inputSet:      newMock("C", false, mockErr),
			wantSatisfied: true, // !false is true
			wantErr:       mockErr,
			wantName:      "!(C)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := prerequisite.Not(tt.inputSet)

			// 1. Test GetName()
			if gotName := s.GetName(); gotName != tt.wantName {
				t.Errorf("GetName() = %v, want %v", gotName, tt.wantName)
			}

			// 2. First call to Check()
			gotSatisfied, gotErr := s.Check()
			if gotSatisfied != tt.wantSatisfied {
				t.Errorf("Check() gotSatisfied = %v, want %v", gotSatisfied, tt.wantSatisfied)
			}
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("Check() gotErr = %v, want %v", gotErr, tt.wantErr)
			}

			// 3. Test memoization
			// Call Check() again, the inner mockSet's Check() should not be executed again.
			s.Check()
			mock := tt.inputSet.(*mockSet)
			if mock.checkCount != 1 {
				t.Errorf("underlying Set.Check() was called %d times, want 1", mock.checkCount)
			}

			// 4. Test Range()
			count := 0
			for range s.Range() {
				count++
			}
			if count != 1 {
				t.Errorf("Range() should yield 1 item, but got %d", count)
			}

			// 5. Test if Output() panics
			// We don't check the specific output, just ensure the program doesn't crash upon calling it.
			s.Output()
		})
	}
}

// TestSetAnd tests the And logic.
func TestSetAnd(t *testing.T) {
	mockErr1 := errors.New("mock error 1")
	mockErr2 := errors.New("mock error 2")

	tests := []struct {
		name           string
		inputSets      []prerequisite.Set
		wantSatisfied  bool
		wantErr        error
		wantName       string
		wantRangeCount int
	}{
		{
			name:           "All True",
			inputSets:      []prerequisite.Set{newMock("A", true, nil), newMock("B", true, nil)},
			wantSatisfied:  true,
			wantErr:        nil,
			wantName:       "(A) && (B)",
			wantRangeCount: 2,
		},
		{
			name:           "One False",
			inputSets:      []prerequisite.Set{newMock("A", true, nil), newMock("B", false, nil)},
			wantSatisfied:  false,
			wantErr:        nil,
			wantName:       "(A) && (B)",
			wantRangeCount: 2,
		},
		{
			name:           "All False",
			inputSets:      []prerequisite.Set{newMock("A", false, nil), newMock("B", false, nil)},
			wantSatisfied:  false,
			wantErr:        nil,
			wantName:       "(A) && (B)",
			wantRangeCount: 2,
		},
		{
			name:           "One Error, others True",
			inputSets:      []prerequisite.Set{newMock("A", true, nil), newMock("B", true, mockErr1)},
			wantSatisfied:  true, // The error does not affect the boolean result calculation.
			wantErr:        mockErr1,
			wantName:       "(A) && (B)",
			wantRangeCount: 2,
		},
		{
			name:           "One Error, one False",
			inputSets:      []prerequisite.Set{newMock("A", false, nil), newMock("B", true, mockErr1)},
			wantSatisfied:  false, // because A is false
			wantErr:        mockErr1,
			wantName:       "(A) && (B)",
			wantRangeCount: 2,
		},
		{
			name:           "Multiple Errors",
			inputSets:      []prerequisite.Set{newMock("A", true, mockErr1), newMock("B", true, mockErr2)},
			wantSatisfied:  true,
			wantErr:        errors.Join(mockErr1, mockErr2),
			wantName:       "(A) && (B)",
			wantRangeCount: 2,
		},
		{
			name:           "With nil set",
			inputSets:      []prerequisite.Set{newMock("A", true, nil), nil, newMock("B", true, nil)},
			wantSatisfied:  true,
			wantErr:        nil,
			wantName:       "(A) && (B)",
			wantRangeCount: 2,
		},
		{
			name:           "Empty sets",
			inputSets:      []prerequisite.Set{},
			wantSatisfied:  true, // The initial value for And is true
			wantErr:        nil,
			wantName:       "",
			wantRangeCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := prerequisite.And(tt.inputSets...)

			// 1. Test GetName()
			if gotName := s.GetName(); gotName != tt.wantName {
				t.Errorf("GetName() = %v, want %v", gotName, tt.wantName)
			}

			// 2. First call to Check()
			gotSatisfied, gotErr := s.Check()
			if gotSatisfied != tt.wantSatisfied {
				t.Errorf("Check() gotSatisfied = %v, want %v", gotSatisfied, tt.wantSatisfied)
			}
			// Compare the combined error
			if (gotErr == nil && tt.wantErr != nil) || (gotErr != nil && tt.wantErr == nil) || (gotErr != nil && tt.wantErr != nil && gotErr.Error() != tt.wantErr.Error()) {
				t.Errorf("Check() gotErr = %v, want %v", gotErr, tt.wantErr)
			}

			// 3. Test memoization
			s.Check()
			for _, set := range tt.inputSets {
				if mock, ok := set.(*mockSet); ok {
					if mock.checkCount != 1 {
						t.Errorf("underlying Set.Check() for %s was called %d times, want 1", mock.name, mock.checkCount)
					}
				}
			}

			// 4. Test Range()
			count := 0
			for range s.Range() {
				count++
			}
			if count != tt.wantRangeCount {
				t.Errorf("Range() should yield %d items, but got %d", tt.wantRangeCount, count)
			}

			// 5. Test if Output() panics
			s.Output()
		})
	}
}

// TestSetOr tests the Or logic.
func TestSetOr(t *testing.T) {
	mockErr1 := errors.New("mock error 1")
	mockErr2 := errors.New("mock error 2")

	tests := []struct {
		name           string
		inputSets      []prerequisite.Set
		wantSatisfied  bool
		wantErr        error
		wantName       string
		wantRangeCount int
	}{
		{
			name:           "One True",
			inputSets:      []prerequisite.Set{newMock("A", false, nil), newMock("B", true, nil)},
			wantSatisfied:  true,
			wantErr:        nil,
			wantName:       "(A) || (B)",
			wantRangeCount: 2,
		},
		{
			name:           "All True",
			inputSets:      []prerequisite.Set{newMock("A", true, nil), newMock("B", true, nil)},
			wantSatisfied:  true,
			wantErr:        nil,
			wantName:       "(A) || (B)",
			wantRangeCount: 2,
		},
		{
			name:           "All False",
			inputSets:      []prerequisite.Set{newMock("A", false, nil), newMock("B", false, nil)},
			wantSatisfied:  false,
			wantErr:        nil,
			wantName:       "(A) || (B)",
			wantRangeCount: 2,
		},
		{
			name:           "One Error, but another is True",
			inputSets:      []prerequisite.Set{newMock("A", true, nil), newMock("B", false, mockErr1)},
			wantSatisfied:  true, // because A is true
			wantErr:        mockErr1,
			wantName:       "(A) || (B)",
			wantRangeCount: 2,
		},
		{
			name:           "One Error, others False",
			inputSets:      []prerequisite.Set{newMock("A", false, nil), newMock("B", false, mockErr1)},
			wantSatisfied:  false, // because there is no true result
			wantErr:        mockErr1,
			wantName:       "(A) || (B)",
			wantRangeCount: 2,
		},
		{
			name:           "Multiple Errors",
			inputSets:      []prerequisite.Set{newMock("A", false, mockErr1), newMock("B", false, mockErr2)},
			wantSatisfied:  false,
			wantErr:        errors.Join(mockErr1, mockErr2),
			wantName:       "(A) || (B)",
			wantRangeCount: 2,
		},
		{
			name:           "With nil set",
			inputSets:      []prerequisite.Set{newMock("A", false, nil), nil, newMock("B", true, nil)},
			wantSatisfied:  true,
			wantErr:        nil,
			wantName:       "(A) || (B)",
			wantRangeCount: 2,
		},
		{
			name:           "Empty sets",
			inputSets:      []prerequisite.Set{},
			wantSatisfied:  false, // The initial value for Or is false
			wantErr:        nil,
			wantName:       "",
			wantRangeCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := prerequisite.Or(tt.inputSets...)

			// 1. Test GetName()
			if gotName := s.GetName(); gotName != tt.wantName {
				t.Errorf("GetName() = %v, want %v", gotName, tt.wantName)
			}

			// 2. First call to Check()
			gotSatisfied, gotErr := s.Check()
			if gotSatisfied != tt.wantSatisfied {
				t.Errorf("Check() gotSatisfied = %v, want %v", gotSatisfied, tt.wantSatisfied)
			}
			if (gotErr == nil && tt.wantErr != nil) || (gotErr != nil && tt.wantErr == nil) || (gotErr != nil && tt.wantErr != nil && gotErr.Error() != tt.wantErr.Error()) {
				t.Errorf("Check() gotErr = %v, want %v", gotErr, tt.wantErr)
			}

			// 3. Test memoization
			s.Check()
			for _, set := range tt.inputSets {
				if mock, ok := set.(*mockSet); ok {
					if mock.checkCount != 1 {
						t.Errorf("underlying Set.Check() for %s was called %d times, want 1", mock.name, mock.checkCount)
					}
				}
			}

			// 4. Test Range()
			count := 0
			for range s.Range() {
				count++
			}
			if count != tt.wantRangeCount {
				t.Errorf("Range() should yield %d items, but got %d", tt.wantRangeCount, count)
			}

			// 5. Test if Output() panics
			s.Output()
		})
	}
}

// TestNestedPrerequisites tests nested combinations of prerequisites.
func TestNestedPrerequisites(t *testing.T) {
	// Expression: (A || B) && !(C)
	// A = false, B = true -> (A || B) = true
	// C = false -> !(C) = true
	// result = true && true = true
	mockA := newMock("A", false, nil)
	mockB := newMock("B", true, nil)
	mockC := newMock("C", false, nil)

	orSet := prerequisite.Or(mockA, mockB)
	notSet := prerequisite.Not(mockC)
	andSet := prerequisite.And(orSet, notSet)

	// 1. Test GetName()
	expectedName := "((A) || (B)) && (!(C))"
	if name := andSet.GetName(); name != expectedName {
		t.Errorf("GetName() got %s, want %s", name, expectedName)
	}

	// 2. Test Check()
	satisfied, err := andSet.Check()
	if err != nil {
		t.Errorf("Check() returned unexpected error: %v", err)
	}
	if !satisfied {
		t.Errorf("Check() got satisfied = false, want true")
	}

	// 3. Test memoization
	// Call Check() again to ensure the underlying mock's Check() is not called again.
	andSet.Check()
	if mockA.checkCount != 1 {
		t.Errorf("mockA.Check() called %d times, want 1", mockA.checkCount)
	}
	if mockB.checkCount != 1 {
		t.Errorf("mockB.Check() called %d times, want 1", mockB.checkCount)
	}
	if mockC.checkCount != 1 {
		t.Errorf("mockC.Check() called %d times, want 1", mockC.checkCount)
	}

	// 4. Test Range()
	count := 0
	// andSet.Range() will yield orSet and notSet
	for range andSet.Range() {
		count++
	}
	if count != 2 {
		t.Errorf("Range() should yield 2 items, but got %d", count)
	}
}
