
package config

import (
    "github.com/codetojoy/prime"
)

type Config struct {
    Max int
    HighestPrime int
    PrimeIndex prime.PrimeIndex
}

func getHighestPrime(n int, primeIndex prime.PrimeIndex) int {
    primeIterator := prime.NewPrimeIterator(primeIndex)
    primeValue := primeIterator.Next()
    result := primeValue
    isDone := false

    for ! isDone {
        if (primeValue > n) {
            isDone = true
        } else {
            result = primeValue
            primeValue = primeIterator.Next()
        }
    }

    return result
}

func New(max int) Config {
    primeIndex := prime.New(max)
    highestPrime := getHighestPrime(max, primeIndex)

    return Config{Max: max, HighestPrime: highestPrime, PrimeIndex: primeIndex}
}
