package item

import (
	"github.com/ctrsploit/sploit-spec/pkg/colorful"
	"github.com/ssst0n3/awesome_libs"
)

type Bool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Result      bool   `json:"result"`
}

func (i Bool) IsEmpty() bool {
	return i.Name == "" && i.Description == "" && i.Result == false
}

func (i Bool) Text() string {
	tpl := `{.result}  {.name}	{.description}`
	return awesome_libs.Format(tpl, awesome_libs.Dict{
		"result":      colorful.Bool(colorful.Text{}, i.Result),
		"name":        i.Name,
		"description": getDescription(i.Description),
	})
}

func (i Bool) Colorful() string {
	output := colorful.Colorful{}
	tpl := `{.result}  {.name}	{.description}`
	return awesome_libs.Format(tpl, awesome_libs.Dict{
		"result":      colorful.Bool(output, i.Result),
		"name":        output.Name(i.Name),
		"description": output.Description(getDescription(i.Description)),
	})
}
