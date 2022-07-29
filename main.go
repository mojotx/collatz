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

	// make sure we passed one command-line argument
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	// convert the argument to an integer, and handle errors
	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("error parsing %q: %s\n", os.Args[1], err.Error())
		usage()
		os.Exit(1)
	}

	// Count the number of steps to reach 1
	var steps int64

	// Iterate until n==1
	for steps = 0; n != 1; steps++ {
		fmt.Printf("%d ", n)
		n = collatz(n)
	}
	fmt.Printf("1\nsteps: %d\n", steps)

}

// collatz conjecture
//         / -
// f(n) = /      n/2 if n === 0 (mod 2)
//         \    3n+1 if n === 1 (mod 2)
//          \-
func collatz(n int64) int64 {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}
