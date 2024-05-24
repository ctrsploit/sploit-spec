package vul

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
	"github.com/urfave/cli/v2"
)

type Vulnerabilities []Vulnerability
type Result map[string]printer.Interface

func (vulnerabilities Vulnerabilities) Check(ctx *cli.Context) (err error) {
	for _, v := range vulnerabilities {
		_, err := v.CheckSec(ctx)
		if err != nil {
			continue
		}
	}
	return nil
}

func (vulnerabilities Vulnerabilities) Exploit(ctx *cli.Context) (err error) {
	for _, v := range vulnerabilities {
		_, err := v.Exploit(ctx)
		if err != nil {
			continue
		}
	}
	return nil
}

func (vulnerabilities Vulnerabilities) Output() {
	result := Result{}
	for _, v := range vulnerabilities {
		result[v.GetName()] = item.Resp{
			Name:        v.GetName(),
			Description: v.GetDescription(),
			Result:      v.GetVulnerabilityExists(),
			Response:    v.GetVulnerabilityResponse(),
		}
	}
	fmt.Println(printer.Printer.Print(result))
	return
}
