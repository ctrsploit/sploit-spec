package vul

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/app"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

type Vulnerabilities []Vulnerability

func (vulnerabilities Vulnerabilities) Check() (err error) {
	for _, v := range vulnerabilities {
		_, err = v.CheckSec()
		if err != nil {
			return
		}
	}
	return
}

func (vulnerabilities Vulnerabilities) Output() {
	var result []item.Bool
	for _, v := range vulnerabilities {
		result = append(result, item.Bool{
			Name:        v.GetName(),
			Description: v.GetDescription(),
			Result:      v.GetVulnerabilityExists(),
		})
	}
	fmt.Println(app.Printer.Print(result))
	return
}
