package unfair

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var CONCURRENCY = int64(5)
var total_primes int64 = 0

func is_prime(i int64) bool {
	if i == 1 {
		return false
	}
	if i == 2 {
		return true
	}
	var j int64 = 2
	for ; j*j <= i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func process_batch(batch_name string, wg *sync.WaitGroup, start int64, end int64) {
	defer wg.Done()

	time_start := time.Now()
	for i := start; i <= end; i++ {
		if is_prime(i) {
			atomic.AddInt64(&total_primes, 1)
		}
	}
	elapsed := time.Since(time_start)
	fmt.Printf("Batch %s took %s to execute.\n", batch_name, elapsed)
}

func Find_primes_unfair_concurrent(count int64) int {

	batch_size := int64(float64(count) / float64(CONCURRENCY))
	var wg sync.WaitGroup
	var i int64 = 1
	for ; i <= CONCURRENCY; i++ {
		wg.Add(1)
		go process_batch(fmt.Sprintf("%d", i), &wg, (i-1)*batch_size+1, i*batch_size)
	}
	wg.Wait()
	fmt.Println("Total primes found :", total_primes)
	return int(total_primes)
}
