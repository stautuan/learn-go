/*
If users send a prime number of text messages, they get a reward.

Write a function that prints all of the prime numbers up to and including
max. It should skip any numbers that are not prime.
*/

package main

import "fmt"

func printPrime(max int) {
    // start checking at 2 because 1 is NOT a prime
    for n := 2; n < max; n++ {
        if n == 2 {
            fmt.Println(n)
        }
        if n % 2 == 0 {
            // skip even numbers because they can't be prime
            continue
        }
        isPrime := true
        // check up to the square root of n
        // anything higher than the square root has no chance of multiplying evenly into n
        for i := 3; i * i <= n; i += 2 {
            if n % i == 0 {
                isPrime = false
                break
            }
        }
        if isPrime {
            fmt.Println(n)
        }
    }
    
}

func test(max unit) {
    fmt.Printf("Primes up to %v:\n", max)
    printPrimes(max)
    fmt.Println("=========================================================")
}

func main() {
    test(10)
    test(20)
    test(30)
}