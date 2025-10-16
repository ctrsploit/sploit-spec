package vul

import (
	"errors"
	"fmt"

	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
	"github.com/urfave/cli/v2"
)

type Vulnerabilities []Vulnerability
type Result map[string]printer.Interface

func (vulnerabilities Vulnerabilities) Check(context *cli.Context) error {
	var errs []error
	for _, v := range vulnerabilities {
		_, err := v.CheckSec(context)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func (vulnerabilities Vulnerabilities) Output() {
	result := Result{}
	for _, v := range vulnerabilities {
		result[v.GetName()] = item.Bool{
			Name:        v.GetName(),
			Description: v.GetDescription(),
			Result:      v.GetVulnerabilityExists(),
		}
	}
	fmt.Println(printer.Printer.Print(result))
	return
}
