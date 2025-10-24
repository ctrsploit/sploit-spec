package prerequisite

import (
	"time"

	"github.com/ctrsploit/sploit-spec/pkg/prerequisite"
)

type Time struct {
	n int
	prerequisite.BasePrerequisite
}

func (p *Time) Check() (satisfied bool, err error) {
	if !p.Checked {
		p.Satisfied = time.Now().Second()%p.n == 0
		p.Checked = true
	}
	satisfied = p.Satisfied
	return
}

var (
	EvenTime = Time{
		n: 2,
		BasePrerequisite: prerequisite.BasePrerequisite{
			Name: "2 | Time",
			Info: "time %% 2 == 0",
		},
	}
)
