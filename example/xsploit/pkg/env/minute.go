package env

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
	"time"
)

func Minute() (result printer.Interface) {
	result = item.Short{
		Name:        "minute",
		Description: "second of current minute",
		Result:      fmt.Sprintf("%d", time.Now().Minute()),
	}
	return
}
