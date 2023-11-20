package prerequisite_test

import (
	"github.com/ctrsploit/sploit-spec/pkg/app"
	"github.com/ctrsploit/sploit-spec/pkg/prerequisite"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

var p *prerequisite.BasePrerequisite

func ExampleBasePrerequisite() {
	p = &prerequisite.BasePrerequisite{
		Name: "CAP_SYS_ADMIN",
		Info: "Container with cap_sys_admin is dangerous",
	}
	err := p.Check()
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
}

func ExampleBasePrerequisite_Output_text() {
	ExampleBasePrerequisite()
	app.Printer = printer.NewWorker(printer.TypeText)
	p.Output()
	// Output: [PREREQUISITE CAP_SYS_ADMIN]
	//[N]  CAP_SYS_ADMIN	# Container with cap_sys_admin is dangerous
}

func ExampleBasePrerequisite_Output_colorful() {
	ExampleBasePrerequisite()
	app.Printer = printer.NewWorker(printer.TypeColorful)
	p.Output()
	// _Output: [PREREQUISITE CAP_SYS_ADMIN]
	//✘  CAP_SYS_ADMIN
}

func ExampleBasePrerequisite_Output_json() {
	ExampleBasePrerequisite()
	app.Printer = printer.NewWorker(printer.TypeJson)
	p.Output()
	// Output: {"Name":{"name":"PREREQUISITE CAP_SYS_ADMIN"},"Prerequisite":{"name":"CAP_SYS_ADMIN","description":"Container with cap_sys_admin is dangerous","result":false}}
}
