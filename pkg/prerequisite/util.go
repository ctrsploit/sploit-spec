package prerequisite

import "fmt"

func (p *BasePrerequisite) WrapErr(err error) error {
	if err == nil {
		return nil
	}

	if p == nil {
		panic("BasePrerequisite.WrapErr called with nil receiver")
	}

	return fmt.Errorf("failed to check [%s]: %w", p.GetName(), err)
}
