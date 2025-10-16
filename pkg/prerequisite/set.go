package prerequisite

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ctrsploit/sploit-spec/pkg/log"
	"github.com/ctrsploit/sploit-spec/pkg/printer"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

type Set interface {
	Check() (satisfied bool, err error)
	Range() <-chan Set
	GetName() string
	Output()
}

type SetNot struct {
	Set       Set
	satisfied bool
	checked   bool
	err       error
}

func Not(set Set) *SetNot {
	return &SetNot{
		Set: set,
	}
}

func (s *SetNot) Check() (bool, error) {
	if s.checked {
		return s.satisfied, s.err
	}
	s.checked = true

	var satisfied bool
	satisfied, s.err = s.Set.Check()
	s.satisfied = !satisfied
	return s.satisfied, s.err
}

// Range only return itself channel
func (s *SetNot) Range() <-chan Set {
	ch := make(chan Set)
	go func() {
		defer close(ch)
		ch <- s
	}()
	return ch
}

func (s *SetNot) GetName() string {
	return fmt.Sprintf("!(%s)", s.Set.GetName())
}

func (s *SetNot) Output() {
	r := Result{
		Name: result.SubTitle{
			Name: fmt.Sprintf("PREREQUISITE %s", s.GetName()),
		},
		Prerequisite: item.Bool{
			Name:        s.GetName(),
			Description: "",
			Result:      s.satisfied,
		},
	}
	log.Logger.Debugf("prerequisite\n%s\n", printer.Printer.Print(r))
}

type SetAnd struct {
	Sets      []Set
	satisfied bool
	checked   bool
	err       error
}

func And(sets ...Set) *SetAnd {
	return &SetAnd{
		Sets: sets,
	}
}

func (s *SetAnd) Check() (bool, error) {
	if s.checked {
		return s.satisfied, s.err
	}
	s.checked = true
	s.satisfied = true
	for _, set := range s.Sets {
		if set == nil {
			continue
		}
		r, err := set.Check()
		if err != nil {
			awesome_error.CheckDebug(fmt.Errorf("[PREREQUISITE %s]:[%s] %w", s.GetName(), set.GetName(), err))
			s.err = errors.Join(s.err, err)
			// Removed 'continue' here to allow boolean evaluation below.
			// continue
		}
		if !r {
			s.satisfied = false
		}
	}
	return s.satisfied, s.err
}

func (s *SetAnd) Range() <-chan Set {
	ch := make(chan Set)
	go func() {
		defer close(ch)
		for _, set := range s.Sets {
			if set == nil {
				continue
			}
			ch <- set
		}
	}()
	return ch
}

func (s *SetAnd) GetName() string {
	var names []string
	for _, set := range s.Sets {
		if set == nil {
			continue
		}
		names = append(names, fmt.Sprintf("(%s)", set.GetName()))
	}
	return strings.Join(names, " && ")
}

func (s *SetAnd) Output() {
	r := Result{
		Name: result.SubTitle{
			Name: fmt.Sprintf("PREREQUISITE %s", s.GetName()),
		},
		Prerequisite: item.Bool{
			Name:        s.GetName(),
			Description: "",
			Result:      s.satisfied,
		},
	}
	log.Logger.Debugf("prerequisite\n%s\n", printer.Printer.Print(r))
	for i := range s.Range() {
		i.Output()
	}
}

type SetOr struct {
	Sets      []Set
	satisfied bool
	checked   bool
	err       error
}

func Or(sets ...Set) *SetOr {
	return &SetOr{
		Sets: sets,
	}
}

func (s *SetOr) Check() (bool, error) {
	if s.checked {
		return s.satisfied, s.err
	}
	s.checked = true
	for _, set := range s.Sets {
		if set == nil {
			continue
		}
		r, err := set.Check()
		if err != nil {
			awesome_error.CheckDebug(fmt.Errorf("[PREREQUISITE %s]:[%s] %w", s.GetName(), set.GetName(), err))
			s.err = errors.Join(s.err, err)
			// Removed 'continue' here to allow boolean evaluation below.
			// continue
		}
		if r {
			s.satisfied = true
		}
	}
	return s.satisfied, s.err
}

func (s *SetOr) Range() <-chan Set {
	ch := make(chan Set)
	go func() {
		defer close(ch)
		for _, set := range s.Sets {
			if set == nil {
				continue
			}
			ch <- set
		}
	}()
	return ch
}

func (s *SetOr) GetName() string {
	var names []string
	for _, set := range s.Sets {
		if set == nil {
			continue
		}
		names = append(names, fmt.Sprintf("(%s)", set.GetName()))
	}
	return strings.Join(names, " || ")
}

func (s *SetOr) Output() {
	r := Result{
		Name: result.SubTitle{
			Name: fmt.Sprintf("PREREQUISITE %s", s.GetName()),
		},
		Prerequisite: item.Bool{
			Name:        s.GetName(),
			Description: "",
			Result:      s.satisfied,
		},
	}
	log.Logger.Debugf("prerequisite\n%s\n", printer.Printer.Print(r))
	for i := range s.Range() {
		i.Output()
	}
}
