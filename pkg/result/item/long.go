package item

import (
	"github.com/ctrsploit/sploit-spec/pkg/colorful"
	"github.com/ssst0n3/awesome_libs"
)

type Long struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Result      string `json:"result"`
}

func (l Long) Text() string {
	tpl := `{.name}	{.description}
{.result}`
	return awesome_libs.Format(tpl, awesome_libs.Dict{
		"name":        l.Name,
		"result":      l.Result,
		"description": getDescription(l.Description),
	})
}

func (l Long) Colorful() string {
	output := colorful.Colorful{}
	tpl := `{.name}	{.description}
{.result}`
	return awesome_libs.Format(tpl, awesome_libs.Dict{
		"name":        output.Name(l.Name),
		"result":      output.Result(l.Result),
		"description": output.Description(getDescription(l.Description)),
	})
}
