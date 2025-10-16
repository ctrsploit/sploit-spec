package user

import (
	"os/user"
	"strconv"

	"github.com/ctrsploit/sploit-spec/pkg/prerequisite"
	"github.com/ssst0n3/awesome_libs/awesome_error"
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
	if !p.Checked {
		current, err := user.Current()
		if err != nil {
			awesome_error.CheckErr(err)
			return false, err
		}
		u, err := strconv.Atoi(current.Uid)
		if err != nil {
			awesome_error.CheckErr(err)
			return false, err
		}
		p.Satisfied = uint(u) == p.ExpectedUser
		p.Checked = true
	}
	satisfied = p.Satisfied
	return
}
