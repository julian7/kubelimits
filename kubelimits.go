package kubelimits

import (
	"errors"
)

var ErrScan = errors.New("Scan error")

type Setter struct {
	loggers []func(string)
}

func Set(loggers ...func(string)) error {
	s := &Setter{loggers: loggers}
	if err := s.SetMemory(); err != nil {
		return err
	}

	return s.SetCPU()
}

func (s *Setter) log(msg string) {
	for _, logger := range s.loggers {
		logger(msg)
	}
}
