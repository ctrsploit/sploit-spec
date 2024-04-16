package auto

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"xsploit/env/minute"
	"xsploit/env/second"
)

type Result map[string]printer.Interface

func Human(machine Spec) (human Result) {
	human = Result{
		"second": second.Human(machine.Second),
		"minute": minute.Human(machine.Minute),
	}
	return
}

func Print() {
	machine := Auto()
	human := Human(machine)
	u := result.Union{
		Machine: machine,
		Human:   human,
	}
	fmt.Println(printer.Printer.Print(u))
}
