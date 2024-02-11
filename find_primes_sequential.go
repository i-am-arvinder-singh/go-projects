package main

func is_prime(i int) bool {
	if i == 2 {
		return true
	}
	for j:=2; j*j<=i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func find_primes_sequential(count int) int {
	total_primes:=0
	for i:=2; i<=count; i++ {
		if is_prime(i) {
			total_primes++
		}
	}
	return total_primes
}
