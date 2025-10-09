package prerequisite

import "github.com/ssst0n3/awesome_libs/awesome_error"

type Set interface {
	Check() (satisfied bool, err error)
	Range() <-chan Set
}

type SetAnd struct {
	Sets []Set
}

func And(sets ...Set) SetAnd {
	return SetAnd{
		sets,
	}
}

func (s SetAnd) Check() (satisfied bool, err error) {
	satisfied = true
	for _, set := range s.Sets {
		if set == nil {
			continue
		}
		r, err := set.Check()
		if err != nil {
			awesome_error.CheckWarning(err)
			continue
		}
		if !r {
			satisfied = false
		}
	}
	return
}

func (s SetAnd) Range() <-chan Set {
	ch := make(chan Set)
	go func() {
		defer close(ch)
		for _, set := range s.Sets {
			if set == nil {
				continue
			}
			for leaf := range set.Range() {
				ch <- leaf
			}
		}
	}()
	return ch
}

type SetOr struct {
	Sets []Set
}

func Or(sets ...Set) SetOr {
	return SetOr{
		sets,
	}
}

func (s SetOr) Check() (satisfied bool, err error) {
	for _, set := range s.Sets {
		if set == nil {
			continue
		}
		r, err := set.Check()
		if err != nil {
			awesome_error.CheckWarning(err)
			continue
		}
		if r {
			satisfied = true
		}
	}
	return
}

func (s SetOr) Range() <-chan Set {
	ch := make(chan Set)
	go func() {
		defer close(ch)
		for _, set := range s.Sets {
			if set == nil {
				continue
			}
			for leaf := range set.Range() {
				ch <- leaf
			}
		}
	}()
	return ch
}
