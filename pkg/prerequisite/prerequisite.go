package prerequisite

import (
	"fmt"

	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

type Interface interface {
	GetExeEnv() int
	Output()
	Check() (bool, error)
}

type BasePrerequisite struct {
	Name      string
	Info      string
	ExeEnv    int
	Checked   bool
	Satisfied bool
	Err       error
}

func (p *BasePrerequisite) GetName() string {
	return p.Name
}

func (p *BasePrerequisite) GetExeEnv() int {
	return p.ExeEnv
}

func (p *BasePrerequisite) GetChecked() bool {
	return p.Checked
}

// Check is a suggested example implementation; override it in derived structs.
func (p *BasePrerequisite) Check() (bool, error) {
	return p.CheckTemplate(func() {
		p.Satisfied = false
		p.Err = nil
	})
}

// CheckTemplate is a template method for Check();
// f is the actual check function. sets p.Satisfied and p.Err in f().
func (p *BasePrerequisite) CheckTemplate(f func()) (bool, error) {
	if p.Checked {
		return p.Satisfied, p.Err
	}
	p.Checked = true
	// reset
	p.Err = nil
	p.Satisfied = false
	// call the actual check function
	f()
	return p.Satisfied, p.Err
}

type Result struct {
	Name         result.SubTitle
	Prerequisite item.Bool
}

// Output print prerequisite with colorful; must be used after p.Check().
func (p *BasePrerequisite) Output() {
	if !p.Checked {
		panic("prerequisite.Interface.Output() must be used after Check()")
	}
	r := Result{
		Name: result.SubTitle{
			Name: fmt.Sprintf("PREREQUISITE %s", p.Name),
		},
		Prerequisite: item.Bool{
			Name:        p.Name,
			Description: p.Info,
			Result:      p.Satisfied,
		},
	}
	log.Logger.Debugf("prerequisite\n%s\n", printer.Printer.Print(r))
	return
}

func (p *BasePrerequisite) Range() <-chan Set {
	ch := make(chan Set)
	go func() {
		defer close(ch)
		ch <- p
	}()
	return ch
}
