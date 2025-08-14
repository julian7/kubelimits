//go:build linux && go1.25

package kubelimits

func (s *setter) SetCPU() error {
	return nil
}
