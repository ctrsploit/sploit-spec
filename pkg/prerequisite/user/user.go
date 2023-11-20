package user

import (
	"github.com/ctrsploit/sploit-spec/pkg/prerequisite"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"os/user"
	"strconv"
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

func (p *MustBe) Check() (err error) {
	err = p.BasePrerequisite.Check()
	if err != nil {
		return
	}
	current, err := user.Current()
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	u, err := strconv.Atoi(current.Uid)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	p.Satisfied = uint(u) == p.ExpectedUser
	return
}
