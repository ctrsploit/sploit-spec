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
	RuleB      item.Bool `json:"rule_b"`
	RuleC      item.Long `json:"rule_c"`
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
	RuleB: item.Bool{
		Name:        "Rule B",
		Description: "bbbbb",
		Result:      false,
	},
	RuleC: item.Long{
		Name:        "Rule C",
		Description: "ccccc",
		Result:      "word",
	},
}

func Test_extractPrinter(t *testing.T) {
	printers := extractPrinter(reflect.ValueOf(r))
	expect := []Printer{
		result.Title{Name: "Example for structured result"},
		item.Short{
			Name:        "Rule A",
			Description: "aaaaa",
			Result:      "value",
		},
		item.Bool{
			Name:        "Rule B",
			Description: "bbbbb",
			Result:      false,
		},
	}
	assert.Equal(t, expect, printers)
}

func TestWorker_Print(t *testing.T) {
	{
		printer := NewWorker(TypeText)
		s := printer.Print(r)
		expect := `===========Example for structured result===========
Rule A:			value	# aaaaa
[N]  Rule B	# bbbbb
`
		assert.Equal(t, expect, s)
	}
	{
		printer := NewWorker(TypeJson)
		s := printer.Print(r)
		expect := `{"Name":{"name":"Example for structured result"},"Nested":{"rule_a":{"name":"Rule A","description":"aaaaa","result":"value"}},"rule_b":{"name":"Rule B","description":"bbbbb","result":false},"rule_c":{"name":"Rule C","description":"ccccc","result":"word"}}`
		assert.Equal(t, expect, s)
	}
	{
		printer := NewWorker(TypeColorful)
		s := printer.Print(r)
		fmt.Println(s)
	}
}
