//go:build linux && !go1.25

package kubelimits

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"runtime"
)

const cpuMaxFilename = "/sys/fs/cgroup/cpu.max"

func (s *setter) SetCPU() error {
	file, err := os.Open(cpuMaxFilename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			s.log("no cpu limit set")
			return nil
		}
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
