
package factor

import (
    "fmt"
    "log"
    "github.com/codetojoy/prime"
)

type Factorization struct {
    n int
    max int
    factors Factors
    primeIndex prime.PrimeIndex
    primeIterator prime.PrimeIterator
}

const (
    UNKNOWN_VALUE = -1
)

func NewFactorization(n int, max int, primeIndex prime.PrimeIndex) Factorization {
    result := Factorization{n: n, max: max, primeIndex: primeIndex}
    result.factors = NewFactors(max, primeIndex)
    result.primeIterator = prime.NewPrimeIterator(primeIndex)
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

    // TODO: this should bubble up the error where all are handled in one place
    if err != nil {
        log.Fatal(err)
    } else {
        result.factors = resultFactors
        result.n = UNKNOWN_VALUE
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
