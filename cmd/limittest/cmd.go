package main

import (
	"fmt"

	"github.com/julian7/kubelimits"
)

func main() {
	err := kubelimits.Set(func(s string) { fmt.Printf("log: %s\n", s) })
	if err != nil {
		fmt.Printf(err.Error())
	}
}
