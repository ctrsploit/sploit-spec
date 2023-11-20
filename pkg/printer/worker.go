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

func extractPrinter(field reflect.Value) (printer Printer, drop bool) {
	if p, ok := field.Interface().(item.Bool); ok {
		printer = p
		if !p.Result {
			// drop result after False item.Bool
			drop = true
			return
		}
	}
	if p, ok := field.Interface().(Printer); ok {
		printer = p
		return
	}
	return
}

// extractPrinter extracts Printers from a result struct
// if there's an item.Bool, and it's result is false, drop results after the item.Bool
func extractPrinters(v reflect.Value) (printers []Printer) {
	switch v.Kind() {
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			field := v.Index(i)
			printer, drop := extractPrinter(field)
			if printer != nil {
				printers = append(printers, printer)
				if drop {
					return
				}
			} else {
				printers = append(printers, extractPrinters(field)...)
			}
		}
	case reflect.Struct:
		for i := 0; i < v.Type().NumField(); i++ {
			field := v.Field(i)
			printer, drop := extractPrinter(field)
			if printer != nil {
				printers = append(printers, printer)
				if drop {
					return
				}
			} else {
				printers = append(printers, extractPrinters(field)...)
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
		printers := extractPrinters(reflect.ValueOf(object))
		s = Print(w.PrintFunc, printers...)
	}
	return
}
