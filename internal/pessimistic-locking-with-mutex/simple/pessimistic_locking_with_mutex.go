package pessimisticlockingwithmutex

import (
	"fmt"
	"sync"
	"time"
)

var count int32 = 0
var wg sync.WaitGroup
var mu sync.Mutex

func inc_count() {
	defer wg.Done()
	mu.Lock()
	count++
	mu.Unlock()
}

func Pessimistic_locking_with_mutex(total_count int64) {

	time_start := time.Now()
	for i := 0; i < int(total_count); i++ {
		wg.Add(1)
		go inc_count()
	}
	wg.Wait()
	elapsed := time.Since(time_start)
	fmt.Printf("Incrementing count by 1 million took %s secs %d val", elapsed, count)
}
