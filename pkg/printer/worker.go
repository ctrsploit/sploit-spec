package printer

import (
	"github.com/ctrsploit/sploit-spec/pkg/result"
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

func extractPrinter(field reflect.Value, dropAfterFalse bool) (printer Interface, drop bool) {
	if p, ok := field.Interface().(item.Bool); ok {
		printer = p
		if !p.Result {
			// drop result after False item.Bool
			drop = true
		}
	}
	if p, ok := field.Interface().(Interface); ok {
		printer = p
	}
	drop = drop && dropAfterFalse
	return
}

// extractPrinter extracts Printers from a result struct
// if there's an item.Bool, and it's result is false, drop results after the item.Bool
func extractPrinters(v reflect.Value, dropAfterFalse bool) (printers []Interface) {
	printer, _ := extractPrinter(v, dropAfterFalse)
	if printer != nil {
		printers = append(printers, printer)
		return
	}
	switch v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			field := v.MapIndex(key)
			printer, drop := extractPrinter(field, dropAfterFalse)
			if printer != nil {
				printers = append(printers, printer)
				if drop {
					return
				}
			} else {
				printers = append(printers, extractPrinters(field, dropAfterFalse)...)
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			field := v.Index(i)
			printer, drop := extractPrinter(field, dropAfterFalse)
			if printer != nil {
				printers = append(printers, printer)
				if drop {
					return
				}
			} else {
				printers = append(printers, extractPrinters(field, dropAfterFalse)...)
			}
		}
	case reflect.Struct:
		for i := 0; i < v.Type().NumField(); i++ {
			field := v.Field(i)
			printer, drop := extractPrinter(field, dropAfterFalse)
			if printer != nil {
				printers = append(printers, printer)
				if drop {
					return
				}
			} else {
				printers = append(printers, extractPrinters(field, dropAfterFalse)...)
			}
		}
	default:
	}
	return
}

func (w Worker) print(object interface{}, dropAfterFalse bool) (s string) {
	switch w.Type {
	case TypeJson:
		if _, ok := object.(result.Union); ok {
			object = object.(result.Union).Machine
		}
		s = w.PrintFunc(object)
	default:
		if _, ok := object.(result.Union); ok {
			object = object.(result.Union).Human
		}
		printers := extractPrinters(reflect.ValueOf(object), dropAfterFalse)
		s = Print(w.PrintFunc, printers...)
	}
	return
}

func (w Worker) PrintDropAfterFalse(object interface{}) (s string) {
	return w.print(object, true)
}

func (w Worker) Print(object interface{}) (s string) {
	return w.print(object, false)
}

// Printer is the default worker, default equal to Text, will be overwritten if --colorful/--json is set
var Printer = NewWorker(TypeText)
