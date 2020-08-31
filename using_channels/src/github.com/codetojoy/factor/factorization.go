
package factor

import (
    "fmt"
    "log"

    "github.com/codetojoy/config"
    "github.com/codetojoy/prime"
)

type Factorization struct {
    n int
    config config.Config
    factors Factors
    primeIterator prime.PrimeIterator
}

const (
    UNKNOWN_VALUE = -1
)

func NewFactorization(n int, config config.Config) Factorization {
    result := Factorization{n: n, config: config}
    result.factors = NewFactors(config)
    result.primeIterator = prime.NewPrimeIterator(config.PrimeIndex)
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
    result := NewFactorization(f.n, f.config)
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
    if f.config.PrimeIndex.IsPrime(f.n) {
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

func (f *Factorization) IsProductMatch(b Factorization, c Factorization) (bool, error) {
    return f.factors.IsProductMatch(b.factors, c.factors)
}

func (f Factorization) String() string {
    return fmt.Sprintf("%d :: %s", f.n, f.factors.String())
}

func (f Factorization) GetPureFactorsString() string {
    return f.factors.String()
}
