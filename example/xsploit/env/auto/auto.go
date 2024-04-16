package auto

import (
	"xsploit/env/minute"
	"xsploit/env/second"
)

type Spec struct {
	Minute minute.Spec
	Second second.Spec
}

func Auto() (env Spec) {
	env = Spec{
		Minute: minute.Minute(),
		Second: second.Second(),
	}
	return
}
