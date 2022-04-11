package jsol

import (
	"fmt"
	"os"
)

var (
	Must   = must
	Should = should
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func should(err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()
	must(err)
}
