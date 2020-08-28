
package main

import (
    "fmt"
    "log"
)

type Factorization struct {
    n int
    max int
    factors Factors
    primeIndex PrimeIndex
    primeIterator PrimeIterator
}

func NewFactorization(n int, max int, primeIndex PrimeIndex) Factorization {
    result := Factorization{n: n, max: max, primeIndex: primeIndex}
    result.factors = NewFactors(max, primeIndex)
    result.primeIterator = NewPrimeIterator(primeIndex)
    return result
}

func GetExponent(n int, i int) int {
    result := 0

    chop := n
    for chop % i == 0 {
        chop /= i
        result++
    }

    return result
}

func (f *Factorization) Multiply(g Factorization) Factorization {
    result := NewFactorization(f.n, f.max, f.primeIndex)
    resultFactors, err := f.factors.Multiply(g.factors)

    if err != nil {
        log.Fatal(err)
    } else {
        result.factors = resultFactors
    }

    return result
}

func (f *Factorization) Factor() {
    if f.primeIndex.IsPrime(f.n) {
        f.factors.Put(f.n, 1)
    } else {
        ceiling := (f.n/2) + 1
        isDone := false
        f.primeIterator.Reset()
        prime := f.primeIterator.Next()
        for ! isDone {
            if (f.n % prime == 0) {
                value := GetExponent(f.n, prime)
                f.factors.Put(prime, value)
            }
            prime = f.primeIterator.Next()

            isDone = (prime > ceiling)
        }
    }
}

func (f Factorization) String() string {
    return fmt.Sprintf("%d :: %s", f.n, f.factors.String())
}

func (f Factorization) GetPureFactorsString() string {
    return f.factors.String()
}
