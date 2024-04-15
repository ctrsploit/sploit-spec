package item

import (
	"github.com/ctrsploit/sploit-spec/pkg/colorful"
	"github.com/ssst0n3/awesome_libs"
	"strings"
)

type List struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Result      []string `json:"result_list"`
}

func (i List) IsEmpty() bool {
	return i.Name == "" && i.Description == "" && len(i.Result) == 0
}

func (s List) Text() string {
	tpl := `{.name}:{.eol}{.result_title}{.eol}{.description}{.eol}{.result}`
	return awesome_libs.Format(tpl, awesome_libs.Dict{
		"format":       getFormat(s.Name),
		"eol":          "\n",
		"name":         s.Name,
		"description":  getDescription(s.Description),
		"result_title": "Result List >",
		"result":       strings.Join(s.Result, "\n"),
	})
}

func (s List) Colorful() string {
	output := colorful.Colorful{}
	tpl := `{.name}:{.eol}{.result_title}{.eol}{.description}{.eol}{.result}`
	var colorResult = make([]string, len(s.Result))
	for i, str := range s.Result {
		colorResult[i] = output.Result(str)
	}

	return awesome_libs.Format(tpl, awesome_libs.Dict{
		"format":       getFormat(s.Name),
		"eol":          "\n",
		"name":         output.Name(s.Name),
		"description":  output.Description(getDescription(s.Description)),
		"result_title": output.Result("Result List >"),
		"result":       strings.Join(colorResult, "\n"),
	})

}
