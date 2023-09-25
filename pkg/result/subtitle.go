package result

import (
	"github.com/ssst0n3/awesome_libs"
)

type SubTitle struct {
	Name string `json:"name"`
}

func (t SubTitle) Text() string {
	tpl := `[{.title}]`
	return awesome_libs.Format(tpl, awesome_libs.Dict{
		"title": t.Name,
	})
}

func (t SubTitle) Colorful() string {
	return t.Text()
}
