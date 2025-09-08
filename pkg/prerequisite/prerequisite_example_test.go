package prerequisite_test

import (
	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/ctrsploit/sploit-spec/pkg/prerequisite"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

var p *prerequisite.BasePrerequisite

func ExampleBasePrerequisite() {
	p = &prerequisite.BasePrerequisite{
		Name: "CAP_SYS_ADMIN",
		Info: "Container with cap_sys_admin is dangerous",
	}
	_, err := p.Check()
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
}

func ExampleBasePrerequisite_Output_text() {
	ExampleBasePrerequisite()
	printer.Printer = printer.NewWorker(printer.TypeText)
	log.Logger.SetLevel(logrus.DebugLevel)
	p.Output()
}

func ExampleBasePrerequisite_Output_colorful() {
	ExampleBasePrerequisite()
	printer.Printer = printer.NewWorker(printer.TypeColorful)
	log.Logger.SetLevel(logrus.DebugLevel)
	p.Output()
}

func ExampleBasePrerequisite_Output_json() {
	ExampleBasePrerequisite()
	printer.Printer = printer.NewWorker(printer.TypeJson)
	log.Logger.SetLevel(logrus.DebugLevel)
	p.Output()
}
