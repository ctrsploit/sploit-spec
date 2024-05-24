package item

import (
	"github.com/ctrsploit/sploit-spec/pkg/colorful"
	"github.com/ssst0n3/awesome_libs"
)

type Resp struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Result      bool   `json:"result"`
	Response    string `json:"response"`
}

func (i Resp) IsEmpty() bool {
	return i.Name == "" && i.Description == "" && i.Result == false
}

func (i Resp) Text() string {
	tpl := `{.result}  {.name}	{.description}{.eol}{.result_title}{.eol}{.response}`
	return awesome_libs.Format(tpl, awesome_libs.Dict{
		"result":         colorful.Bool(colorful.Text{}, i.Result),
		"eol":            "\n",
		"name":           i.Name,
		"description":    getDescription(i.Description),
		"response_title": "Response >",
		"response":       i.Response,
	})
}

func (i Resp) Colorful() string {
	output := colorful.Colorful{}
	tpl := `{.result}  {.name}	{.description}{.eol}{.result_title}{.eol}{.response}`
	return awesome_libs.Format(tpl, awesome_libs.Dict{
		"result":         colorful.Bool(output, i.Result),
		"eol":            "\n",
		"name":           output.Name(i.Name),
		"description":    output.Description(getDescription(i.Description)),
		"response_title": output.Description("Response >"),
		"response":       output.Description(i.Response),
	})
}
