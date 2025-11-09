package vul

import (
	"context"
	"errors"
	"fmt"

	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

type Vulnerabilities []Vulnerability
type Result map[string]printer.Interface

func (vulnerabilities Vulnerabilities) Check(ctx context.Context) error {
	var errs []error
	for _, v := range vulnerabilities {
		_, err := v.CheckSec(ctx)
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
