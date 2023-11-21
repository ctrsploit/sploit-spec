package env

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
)

type Result map[string]printer.Interface

func Auto() {
	s := Second()
	m := Minute()
	result := Result{
		"second": s,
		"minute": m,
	}
	fmt.Println(printer.Printer.Print(result))
}
