package prerequisite

import (
	"fmt"
	"github.com/ctrsploit/sploit-spec/pkg/app"
	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

type Interface interface {
	Check() error
	Output()
	GetSatisfied() bool
}
type Prerequisites []Interface

func (ps Prerequisites) Satisfied() (satisfied bool, err error) {
	satisfied = true
	for _, p := range ps {
		err = p.Check()
		if err != nil {
			return
		}
		p.Output()
		if err != nil {
			return
		}
		if !p.GetSatisfied() {
			satisfied = false
		}
	}
	return
}

type BasePrerequisite struct {
	Name      string
	Info      string
	checked   bool
	Satisfied bool
}

func (p *BasePrerequisite) GetSatisfied() bool {
	return p.Satisfied
}

func (p *BasePrerequisite) Check() (err error) {
	p.checked = true
	return
}

type Result struct {
	Name         result.SubTitle
	Prerequisite item.Bool
}

// Output print prerequisite with colorful; must be used after p.Check().
func (p *BasePrerequisite) Output() {
	if !p.checked {
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
	log.Logger.Debugf("prerequisite\n%s\n", app.Printer.Print(r))
	return
}
