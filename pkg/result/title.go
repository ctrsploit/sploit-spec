package result

import (
	"github.com/ssst0n3/awesome_libs"
	"strings"
)

type Title struct {
	Name string `json:"name"`
}

func (t Title) IsEmpty() bool {
	return t.Name == ""
}

func (t Title) Text() string {
	tpl := `{.padding}{.title}{.padding}`
	return awesome_libs.Format(tpl, awesome_libs.Dict{
		"padding": strings.Repeat("=", 11),
		"title":   t.Name,
	})
}

func (t Title) Colorful() string {
	return t.Text()
}
