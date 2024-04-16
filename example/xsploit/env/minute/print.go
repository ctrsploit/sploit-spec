package minute

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

func Human(machine Spec) (human printer.Interface) {
	human = item.Short{
		Name:        "minute",
		Description: "minute of current time",
		Result:      fmt.Sprintf("%d", machine.Minute),
	}
	return
}

func Print() {
	machine := Minute()
	human := Human(machine)
	u := result.Union{
		Machine: machine,
		Human:   human,
	}
	fmt.Println(printer.Printer.Print(u))
}
