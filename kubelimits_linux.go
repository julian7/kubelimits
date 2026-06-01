//go:build linux

package kubelimits

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
)

const memMaxFilename = "/sys/fs/cgroup/memory.max"

func (s *setter) SetMemory() error {
	file, err := os.Open(memMaxFilename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			s.log("no memory limit set")
			return nil
		}
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data, err := scanInt64(scanner)
	if err != nil {
		return err
	}
	s.log(fmt.Sprintf("setting memory limit to %d", data))
	debug.SetMemoryLimit(data)
	return nil
}

func scanInt64(scanner *bufio.Scanner) (int64, error) {
	if !scanner.Scan() {
		return 0, ErrScan
	}

	data, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		return 0, err
	}

	return data, nil
}
