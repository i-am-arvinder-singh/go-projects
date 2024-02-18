package pessimisticlockingwithmutexinfloop

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var count int32 = 0
var wg sync.WaitGroup
var mc = 2
var mu sync.Mutex

func inc_count(batch_name string, total_count int32) {
	defer wg.Done()

	time_start := time.Now()
	for {
		mu.Lock()
		if count >= total_count {
			mu.Unlock()
			break
		}
		count++
		mu.Unlock()
	}
	elapsed := time.Since(time_start)
	fmt.Printf("Batch %s took %s to execute.\n", batch_name, elapsed)
}

func Pessimistic_locking_with_mutex_inf_loop_impl(total_count int64) {
	for i := 0; i < mc; i++ {
		wg.Add(1)
		go inc_count(strconv.Itoa(i), int32(total_count))
	}
	wg.Wait()
	fmt.Println("Incrementing count by 1 million: ", count)
}
