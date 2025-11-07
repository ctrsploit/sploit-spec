package testcase

import "github.com/ctrsploit/sploit-spec/pkg/risk"

type Testcase struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Level        risk.Level
	Risk         risk.Risk
	Background   string   `json:"background"`
	Source       []string `json:"source"`
	Precondition []string `json:"precondition"`
}
