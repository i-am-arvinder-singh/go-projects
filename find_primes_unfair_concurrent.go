package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var CONCURRENCY = int(1e1)
var total_primes int32 = 0

func process_batch(batch_name string, wg *sync.WaitGroup, start int, end int) {
	defer wg.Done()

	time_start := time.Now()
	for i:=start; i<=end; i++ {
		if is_prime(i) {
			atomic.AddInt32(&total_primes, 1);
		}
	}
	elapsed := time.Since(time_start)
	fmt.Printf("Batch %s took %s to execute.\n", batch_name, elapsed)
}

func find_primes_unfair_concurrent(count int) int {
	
	batch_size := int(float64(count)/float64(CONCURRENCY))
	var wg sync.WaitGroup
	for i:=1; i<=CONCURRENCY; i++ {
		wg.Add(1)
		go process_batch(fmt.Sprintf("%d", i), &wg, (i-1)*batch_size+1, i*batch_size)
	}
	wg.Wait()
	fmt.Println("Total primes found :", total_primes)
	return int(total_primes)
}


