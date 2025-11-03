package prerequisite

import "fmt"

func (p *BasePrerequisite) WrapErr(err error) error {
	return fmt.Errorf("failed to check [%s], caused by %w", p.GetName(), err)
}
