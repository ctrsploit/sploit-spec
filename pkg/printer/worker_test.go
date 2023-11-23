package printer

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type Nested struct {
	RuleA item.Short `json:"rule_a"`
}

type Result struct {
	NotPrinter string `json:"-"`
	Name       result.Title
	Nested     Nested
	Array      []item.Short
	RuleC      item.Bool `json:"rule_c"`
	RuleD      item.Long `json:"rule_d"`
}

var r = Result{
	NotPrinter: "not a printer",
	Name: result.Title{
		Name: "Example for structured result",
	},
	Nested: Nested{
		RuleA: item.Short{
			Name:        "Rule A",
			Description: "aaaaa",
			Result:      "value",
		},
	},
	Array: []item.Short{
		{
			Name:        "b1",
			Description: "b1",
			Result:      "b1",
		},
		{
			Name:        "b2",
			Description: "b2",
			Result:      "b2",
		},
	},
	RuleC: item.Bool{
		Name:        "Rule C",
		Description: "ccccc",
		Result:      false,
	},
	RuleD: item.Long{
		Name:        "Rule D",
		Description: "ddddd",
		Result:      "word",
	},
}

func Test_extractPrinter(t *testing.T) {
	t.Run("pass printer after false item.Bool", func(t *testing.T) {
		printers := extractPrinters(reflect.ValueOf(r), true)
		expect := []Interface{
			result.Title{Name: "Example for structured result"},
			item.Short{
				Name:        "Rule A",
				Description: "aaaaa",
				Result:      "value",
			},
			item.Short{
				Name:        "b1",
				Description: "b1",
				Result:      "b1",
			},
			item.Short{
				Name:        "b2",
				Description: "b2",
				Result:      "b2",
			},
			item.Bool{
				Name:        "Rule C",
				Description: "ccccc",
				Result:      false,
			},
		}
		assert.Equal(t, expect, printers)
	})
	t.Run("slice", func(t *testing.T) {
		r := []item.Bool{
			{
				Name:        "a",
				Description: "a",
				Result:      false,
			},
			{
				Name:        "b",
				Description: "b",
				Result:      true,
			},
		}
		expect := []Interface{item.Bool{
			Name:        "a",
			Description: "a",
			Result:      false,
		}}
		printers := extractPrinters(reflect.ValueOf(r), true)
		assert.Equal(t, expect, printers)
	})
	t.Run("map", func(t *testing.T) {
		r := map[string]item.Bool{
			"a": {
				Name:        "a",
				Description: "a",
				Result:      false,
			},
			"b": {
				Name:        "b",
				Description: "b",
				Result:      true,
			},
		}
		expect := []Interface{item.Bool{
			Name:        "a",
			Description: "a",
			Result:      false,
		}}
		printers := extractPrinters(reflect.ValueOf(r), true)
		assert.Equal(t, expect, printers)
	})
}

func TestWorker_Print(t *testing.T) {
	{
		printer := NewWorker(TypeText)
		s := printer.Print(r)
		expect := `===========Example for structured result===========
Rule A:			value	# aaaaa
b1:			b1	# b1
b2:			b2	# b2
[N]  Rule C	# ccccc
Rule D	# ddddd
word
`
		assert.Equal(t, expect, s)
	}
	{
		printer := NewWorker(TypeJson)
		s := printer.Print(r)
		expect := `{"Name":{"name":"Example for structured result"},"Nested":{"rule_a":{"name":"Rule A","description":"aaaaa","result":"value"}},"Array":[{"name":"b1","description":"b1","result":"b1"},{"name":"b2","description":"b2","result":"b2"}],"rule_c":{"name":"Rule C","description":"ccccc","result":false},"rule_d":{"name":"Rule D","description":"ddddd","result":"word"}}`
		assert.Equal(t, expect, s)
	}
	{
		printer := NewWorker(TypeColorful)
		s := printer.Print(r)
		fmt.Println(s)
	}
}
