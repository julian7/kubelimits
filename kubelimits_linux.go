//go:build linux

package kubelimits

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
)

const cpuMaxFilename = "/sys/fs/cgroup/cpu.max"
const memMaxFilename = "/sys/fs/cgroup/memory.max"

func (s *setter) SetMemory() error {
	file, err := os.Open(memMaxFilename)
	if err != nil {
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

func (s *setter) SetCPU() error {
	file, err := os.Open(cpuMaxFilename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	max, err := scanInt64(scanner)
	if err != nil {
		return err
	}

	period, err := scanInt64(scanner)
	if err != nil {
		return err
	}

	maxprocs := math.Max(1, math.Floor(float64(max)/float64(period)))

	s.log(fmt.Sprintf("setting maxprocs to %d", int(maxprocs)))
	runtime.GOMAXPROCS(int(maxprocs))

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
