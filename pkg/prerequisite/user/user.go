package user

import (
	"fmt"
	"os/user"
	"strconv"

	"github.com/ctrsploit/sploit-spec/pkg/prerequisite"
)

type MustBe struct {
	ExpectedUser uint
	prerequisite.BasePrerequisite
}

var MustBeRoot = MustBe{
	ExpectedUser: 0,
	BasePrerequisite: prerequisite.BasePrerequisite{
		Name: "root",
		Info: "Current user must be root",
	},
}

var MustBeRootToWriteReleaseAgent = MustBe{
	ExpectedUser: MustBeRoot.ExpectedUser,
	BasePrerequisite: prerequisite.BasePrerequisite{
		Name: MustBeRoot.Name,
		Info: "Current user must be root to write release_agent",
	},
}

func (p *MustBe) Check() (satisfied bool, err error) {
	return p.CheckTemplate(func() (bool, error) {
		current, err := user.Current()
		if err != nil {
			p.Err = fmt.Errorf("failed to check [%s], caused by getting current user: %w", p.GetName(), err)
			return p.Satisfied, p.Err
		}
		u, err := strconv.Atoi(current.Uid)
		if err != nil {
			p.Err = fmt.Errorf("failed to check [%s], caused by converting uid: %w", p.GetName(), err)
			return p.Satisfied, p.Err
		}
		p.Satisfied = uint(u) == p.ExpectedUser
		return p.Satisfied, p.Err
	})
}
