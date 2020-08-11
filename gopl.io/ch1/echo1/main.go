// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	fmt.Println("Original")
	original()
	fmt.Println("Exercise 1.1")
	ex1()
	fmt.Println("Exercise 1.2")
	ex2()
	fmt.Println("Exercise 1.3")
	ex3()
}

func original() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println()
}

/**
 * Print os.Args[0]
 */
func ex1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println()
}

/**
 * Print index and value of each arg
 */
func ex2() {
	for i, v := range os.Args {
		fmt.Printf("%v %v", i, v)
		fmt.Println()
	}
	fmt.Println()
}

/**
 * Measure time between original and strings.Join
 */
 func ex3() {
	start := time.Now()
	ex1()
	orig_nanosecs := float64(time.Since(start).Nanoseconds())

	start = time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	join_nanosecs := float64(time.Since(start).Nanoseconds())

	fmt.Printf("Orig: %.2fs nanosecs\nJoin: %.2fs", orig_nanosecs, join_nanosecs)
}

//!-
