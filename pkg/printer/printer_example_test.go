package printer_test

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

type Result struct {
	Name  result.Title
	RuleA item.Bool  `json:"rule_a"`
	RuleB item.Short `json:"rule_b"`
	RuleC item.Long  `json:"rule_c"`
}

var r = Result{
	Name: result.Title{
		Name: "Example for structured result",
	},
	RuleA: item.Bool{
		Name:        "Rule A",
		Description: "aaaaa",
		Result:      true,
	},
	RuleB: item.Short{
		Name:        "Rule B",
		Description: "bbbbb",
		Result:      "value",
	},
	RuleC: item.Long{
		Name:        "Rule C",
		Description: "ccccc",
		Result:      "word",
	},
}

func (r Result) string(t int) (s string) {
	p := printer.GetPrinter(t)
	switch t {
	case printer.TypeJson:
		return p(r)
	}
	s += printer.Print(p, r.Name, r.RuleA)
	if r.RuleA.Result {
		s += printer.Print(p, r.RuleB, r.RuleC)
	}
	return
}

func (r Result) String() string {
	return r.string(printer.TypeColorful)
}

func ExamplePrint_colorful() {
	fmt.Println(r)
}

func ExamplePrint_text() {
	fmt.Println(r.string(printer.TypeText))
	// Output: ===========Example for structured result===========
	//[Y]  Rule A	# aaaaa
	//Rule B:			value	# bbbbb
	//Rule C	# ccccc
	//word
}

func ExamplePrint_json() {
	fmt.Println(r.string(printer.TypeJson))
	// Output: {"Name":{"name":"Example for structured result"},"rule_a":{"name":"Rule A","description":"aaaaa","result":true},"rule_b":{"name":"Rule B","description":"bbbbb","result":"value"},"rule_c":{"name":"Rule C","description":"ccccc","result":"word"}}
}
