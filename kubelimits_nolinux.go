//go:build !linux

package kubelimits

func (s *Setter) SetMemory() error {
	return nil
}

func (s *Setter) SetCPU() error {
	return nil
}
