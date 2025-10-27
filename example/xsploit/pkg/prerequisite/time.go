package prerequisite

import (
	"time"

	"github.com/ctrsploit/sploit-spec/pkg/prerequisite"
)

type Time struct {
	n int
	prerequisite.BasePrerequisite
}

func (p *Time) Check() (bool, error) {
	return p.CheckTemplate(func() (bool, error) {
		p.Satisfied = time.Now().Second()%p.n == 0
		return p.Satisfied, p.Err
	})
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
