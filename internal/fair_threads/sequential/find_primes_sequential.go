package seq

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

func Find_primes_sequential(count int64) int {
	total_primes := 0
	var i int64 = 2
	for ; i <= count; i++ {
		if is_prime(i) {
			total_primes++
		}
	}
	return total_primes
}
