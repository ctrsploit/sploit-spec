package prerequisite

type Set interface {
	GetSatisfied() (satisfied bool, err error)
}

type SetAnd struct {
	Sets []Set
}

func And(sets ...Set) SetAnd {
	return SetAnd{
		sets,
	}
}

func (s SetAnd) GetSatisfied() (satisfied bool, err error) {
	satisfied = true
	for _, set := range s.Sets {
		r, err := set.GetSatisfied()
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

func (s SetOr) GetSatisfied() (satisfied bool, err error) {
	for _, set := range s.Sets {
		r, err := set.GetSatisfied()
		if err != nil {
			return false, err
		}
		if r {
			satisfied = true
		}
	}
	return
}
