package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/i-am-arvinder-singh/demo/internal/fair_threads/fair"
	seq "github.com/i-am-arvinder-singh/demo/internal/fair_threads/sequential"
	"github.com/i-am-arvinder-singh/demo/internal/fair_threads/unfair"
	pessimisticlockingwithmutexinfloop "github.com/i-am-arvinder-singh/demo/internal/pessimistic-locking-with-mutex/inf_loop"
	pessimisticlockingwithmutex "github.com/i-am-arvinder-singh/demo/internal/pessimistic-locking-with-mutex/simple"
	"github.com/i-am-arvinder-singh/demo/internal/https-protocol"
)

func get_extra_args() int64 {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a mode (fair, unfair, seq) and a number as command-line arguments.")
		os.Exit(1)
	}
	i, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println("Invalid number provided. Please provide a valid number.")
		os.Exit(1)
	}

	return i
}

func main() {
	mode := os.Args[1]

	start := time.Now() // get the current time before execution

	var ans int
	switch mode {
	case "fair":
		i := get_extra_args()
		ans = fair.Find_primes_fair_concurrent(i)
	case "unfair":
		i := get_extra_args()
		ans = unfair.Find_primes_unfair_concurrent(i)
	case "seq":
		i := get_extra_args()
		ans = seq.Find_primes_sequential(i)
	case "mutex_inf":
		i := get_extra_args()
		pessimisticlockingwithmutexinfloop.Pessimistic_locking_with_mutex_inf_loop_impl(i)
	case "mutex_simple":
		i := get_extra_args()
		pessimisticlockingwithmutex.Pessimistic_locking_with_mutex(i)
	case "http_protocol":
		httpsprotocol.Http_protocol()
	default:
		fmt.Println("Invalid mode provided. Please provide a valid mode (fair, unfair, seq).")
		return
	}

	elapsed := time.Since(start) // calculate the time difference

	if mode != "mutex" {
		fmt.Println(ans)
	}
	
	fmt.Printf("Program took %s to execute.\n", elapsed)
}
