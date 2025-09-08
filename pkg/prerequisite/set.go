package prerequisite

type Set interface {
	Check() (satisfied bool, err error)
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
			return false, err
		}
		if !r {
			satisfied = false
		}
	}
	return
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
			return false, err
		}
		if r {
			satisfied = true
		}
	}
	return
}
