package item

import (
	"testing"
)

func TestShort(t *testing.T) {
	s := Short{
		Name:        "Name-Test",
		Description: "Description-Test",
		Result:      "Result-Test",
	}

	t.Logf("IsEmpty: %t", s.IsEmpty())
	t.Logf("Colorful: %s", s.Colorful())
	t.Logf("Text: %s", s.Text())
}

func TestList(t *testing.T) {
	s := List{
		Name:        "Name-Test",
		Description: "Description-Test",
		Result: []string{
			"Result-Test: 1",
			"Result-Test: 2",
			"Result-Test: 3",
		},
	}

	t.Logf("IsEmpty: %t", s.IsEmpty())
	t.Logf("Colorful: %s", s.Colorful())
	t.Logf("Text: %s", s.Text())
}
