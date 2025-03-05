package kubelimits

import (
	"errors"
)

// ErrScan is an error value Set() returns with, when
// cgroup2 limit files cannot be parsed.
var ErrScan = errors.New("Scan error")

type setter struct {
	loggers []func(string)
}

// Set reads cgroup2 cpu.max and memory.max, and sets GOMAXPROCS and runtime
// memory limit accordingly. It is a no-op on non-linux environments.
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
