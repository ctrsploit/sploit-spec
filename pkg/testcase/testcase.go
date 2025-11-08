package testcase

import "github.com/ctrsploit/sploit-spec/pkg/risk"

type Testcase struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// Level indicates the risk level of this test (e.g., Low, Medium, High).
	Level risk.Level
	// Risk specifies the type of risk this test is designed to assess,
	// such as injection, misconfiguration, or access control risks.
	Risk risk.Risk
	// Background contains contextual information about the test,
	// such as underlying vulnerability details, attack mechanism,
	// or applicable scenarios.
	Background string `json:"background"`
	// Source lists the reference sources for this test case,
	// for example: official documents, CVE entries, or OWASP materials.
	Source []string `json:"source"`
	// Precondition defines the conditions that must be satisfied
	// before executing this test, like required configurations or data states.
	Precondition []string `json:"precondition"`
	// Steps defines the ordered list of steps to perform during the test.
	// Each Step represents a specific action or verification point.
	Steps []Step `json:"steps"`
	// Expected specifies the expected outcomes after executing the test,
	// used to determine whether it passes or fails.
	Expected []string `json:"expected"`
	// Automated describes how to reproduce the same effects as the manual
	// test steps using automation commands or scripts. It usually includes
	// command-line examples, scripts, or tool configurations that can
	// perform the test automatically.
	Automated string `json:"automated"`
	// Reproduces lists detailed procedures to reproduce the vulnerability
	// or security issue that this test targets.
	Reproduces []Reproduce `json:"reproduces"`
	// Defenses enumerates the defense mechanisms or mitigations
	// that can prevent or reduce the associated risk.
	Defenses []Defense `json:"defenses"`
	// Scores stores scoring or rating information for the test case,
	// used to measure severity or priority of the identified risk.
	Scores []Score `json:"scores"`
	// Examples provides related real-world cases or Capture The Flag (CTF)
	// challenges that demonstrate similar vulnerabilities or exploitation
	// techniques encountered in practice.
	Examples []string `json:"examples"`
}

type Step struct {
	Action  string   `json:"action"`
	Example []string `json:"example"`
}

type Reproduce struct {
	Env     string `json:"env"`
	Startup string `json:"startup"`
	Step    string `json:"step"`
}

type Defense struct {
	Name     string      `json:"name"`
	Content  string      `json:"content"`
	Examples []string    `json:"examples"`
	Type     DefenseType `json:"type"`
}

type DefenseType int

const (
	DefenseTypeUnknown DefenseType = iota
	DefenseTypeWorkaround
	DefenseTypeComplete
	DefenseTypeDetect
)

type Score struct {
	Scenario string   `json:"scenario"`
	Score    float32  `json:"score"`
	Vectors  []Vector `json:"vectors"`
}

type Vector struct {
	Name   string `json:"name"`
	Level  string `json:"level"`
	Reason string `json:"reason"`
}
