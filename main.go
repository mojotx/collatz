// Collatz Conjecture demo code in Go
// https://en.wikipedia.org/wiki/Collatz_conjecture

package main

import (
	"fmt"
	"os"
	"strconv"
)

func usage() {
	fmt.Printf("usage: %s <n>\n\nwhere 'n' is an integer", os.Args[0])
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("error parsing %q: %s\n", os.Args[1], err.Error())
		usage()
		os.Exit(1)
	}

	for i := 0; i < 1000; i++ {
		fmt.Printf("%d ", n)
		n = collatz(n)
	}
	fmt.Println("")

}

func collatz(n int64) int64 {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}
