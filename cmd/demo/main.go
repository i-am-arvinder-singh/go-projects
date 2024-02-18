package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/i-am-arvinder-singh/demo/internal/fair_threads/fair"
	"github.com/i-am-arvinder-singh/demo/internal/fair_threads/sequential"
	"github.com/i-am-arvinder-singh/demo/internal/fair_threads/unfair"
	"github.com/i-am-arvinder-singh/demo/internal/pessimistic-locking-with-mutex/inf_loop"
	"github.com/i-am-arvinder-singh/demo/internal/pessimistic-locking-with-mutex/simple"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a mode (fair, unfair, seq) and a number as command-line arguments.")
		return
	}

	mode := os.Args[1]

	i, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println("Invalid number provided. Please provide a valid number.")
		return
	}

	start := time.Now() // get the current time before execution

	var ans int
	switch mode {
	case "fair":
		ans = fair.Find_primes_fair_concurrent(i)
	case "unfair":
		ans = unfair.Find_primes_unfair_concurrent(i)
	case "seq":
		ans = seq.Find_primes_sequential(i)
	case "mutex_inf":
		pessimisticlockingwithmutexinfloop.Pessimistic_locking_with_mutex_inf_loop_impl(i)
	case "mutex_simple":
		pessimisticlockingwithmutex.Pessimistic_locking_with_mutex(i)
	default:
		fmt.Println("Invalid mode provided. Please provide a valid mode (fair, unfair, seq).")
		return
	}

	elapsed := time.Since(start) // calculate the time difference

	if mode != "mutex" {
		fmt.Println(ans)
	}
	fmt.Printf("The find_primes function took %s to execute.\n", elapsed)
}
