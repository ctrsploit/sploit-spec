package item

import (
	"github.com/ctrsploit/sploit-spec/pkg/colorful"
	"github.com/ssst0n3/awesome_libs"
)

type Short struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Result      string `json:"result"`
}

func getFormat(name string) string {
	// name + ':'
	if len(name) < 7 {
		return "\t\t\t"
	}
	return "\t"
}

func (s Short) Text() string {
	tpl := `{.name}:{.format}{.result}	{.description}`
	return awesome_libs.Format(tpl, awesome_libs.Dict{
		"format":      getFormat(s.Name),
		"name":        s.Name,
		"result":      s.Result,
		"description": getDescription(s.Description),
	})
}

func (s Short) Colorful() string {
	output := colorful.Colorful{}
	tpl := `{.name}:{.format}{.result}	{.description}`
	return awesome_libs.Format(tpl, awesome_libs.Dict{
		"format":      getFormat(s.Name),
		"name":        output.Name(s.Name),
		"description": output.Description(getDescription(s.Description)),
		"result":      output.Result(s.Result),
	})
}
