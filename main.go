package main

import (
    "fmt"
    "time"
)

func main() {
    i := int(1e8)

    start := time.Now() // get the current time before execution

    ans := find_primes_unfair_concurrent(i)

    elapsed := time.Since(start) // calculate the time difference

    fmt.Println(ans)
    fmt.Printf("The find_primes function took %s to execute.\n", elapsed)
}