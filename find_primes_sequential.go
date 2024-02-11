package main

func is_prime(i int) bool {
	for j:=2; j*j<=i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func find_primes(count int) int {
	total_primes:=0
	for i:=0; i<count; i++ {
		if is_prime(i) {
			total_primes++
		}
	}
	return total_primes
}
