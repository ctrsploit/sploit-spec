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
}

func (p *BasePrerequisite) GetExeEnv() int {
	return p.ExeEnv
}

func (p *BasePrerequisite) Check() (bool, error) {
	return false, nil
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
