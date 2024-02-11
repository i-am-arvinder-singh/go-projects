package fair

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var current_num int64 = 1
var total_primes_fair_concurrent int64 = 0
var LIMIT int64 = 0
var CONCURRENCY = int64(5)

func Find_primes_fair_concurrent(count int64) int {
	var wg sync.WaitGroup
	var i int64 = 1
	LIMIT = count
	for ; i <= CONCURRENCY; i++ {
		wg.Add(1)
		go process_batch_fair(fmt.Sprintf("%d", i), &wg)
	}
	wg.Wait()
	fmt.Println("Total primes found :", total_primes_fair_concurrent)
	return int(total_primes_fair_concurrent)
}

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

func process_batch_fair(batch_name string, wg *sync.WaitGroup) {
	defer wg.Done()

	time_start := time.Now()
	// Below commented code is not the correct way to do it
	// ------------
	// for current_num <= LIMIT {
	// 	atomic.AddInt64(&current_num, 1)
	// 	if is_prime(current_num) {
	// 		atomic.AddInt64(&total_primes_fair_concurrent, 1)
	// 	}
	// }
	// ------------
	for {
		num := atomic.AddInt64(&current_num, 1)
		if num > LIMIT {
			break
		}
		if is_prime(num) {
			atomic.AddInt64(&total_primes_fair_concurrent, 1)
		}
	
	}
	elapsed := time.Since(time_start)
	fmt.Printf("Batch %s took %s to execute.\n", batch_name, elapsed)
}
