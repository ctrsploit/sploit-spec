package env

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
	"time"
)

func Second() (result printer.Interface) {
	result = item.Short{
		Name:        "second",
		Description: "second of current time",
		Result:      fmt.Sprintf("%d", time.Now().Second()),
	}
	return result
}
