package printer

import (
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
	"reflect"
)

type Worker struct {
	Type      int
	PrintFunc PrintFunc
}

func NewWorker(t int) *Worker {
	return &Worker{
		Type:      t,
		PrintFunc: GetPrinter(t),
	}
}

// extractPrinter extracts Printers from a result struct
// if there's an item.Bool, and it's result is false, drop results after the item.Bool
func extractPrinter(v reflect.Value) (printers []Printer) {
	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.Type().NumField(); i++ {
			field := v.Field(i)
			printer, ok := field.Interface().(item.Bool)
			if ok {
				printers = append(printers, printer)
				if !printer.Result {
					// drop result after False item.Bool
					return
				}
			} else {
				printer, ok := field.Interface().(Printer)
				if ok {
					printers = append(printers, printer)
				} else {
					printers = append(printers, extractPrinter(field)...)
				}
			}
		}
	default:
	}
	return
}

func (w Worker) Print(object interface{}) (s string) {
	switch w.Type {
	case TypeJson:
		s = w.PrintFunc(object)
	default:
		printers := extractPrinter(reflect.ValueOf(object))
		s = Print(w.PrintFunc, printers...)
	}
	return
}
