package kubelimits

import (
	"errors"
)

var ErrScan = errors.New("Scan error")

type setter struct {
	loggers []func(string)
}

func Set(loggers ...func(string)) error {
	s := &setter{loggers: loggers}
	if err := s.SetMemory(); err != nil {
		return err
	}

	return s.SetCPU()
}

func (s *setter) log(msg string) {
	for _, logger := range s.loggers {
		logger(msg)
	}
}
