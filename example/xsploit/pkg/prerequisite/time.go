package prerequisite

import (
	"github.com/ctrsploit/sploit-spec/pkg/prerequisite"
	"time"
)

type Time struct {
	n int
	prerequisite.BasePrerequisite
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

func (p *Time) Check() (err error) {
	err = p.BasePrerequisite.Check()
	if err != nil {
		return
	}
	p.Satisfied = time.Now().Second()%p.n == 0
	return
}
