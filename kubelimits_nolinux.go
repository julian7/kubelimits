//go:build !linux

package kubelimits

func (s *setter) SetMemory() error {
	return nil
}

func (s *setter) SetCPU() error {
	return nil
}
