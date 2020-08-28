
package main

import (
    "fmt"
)

type Factorial struct {
    n int
    max int
    factorization Factorization
    primeIndex PrimeIndex
}

func NewFactorial(n int, max int, primeIndex PrimeIndex) Factorial {
    return Factorial{n: n, max: max, primeIndex: primeIndex}
}

func (f *Factorial) Compute() {
    accumulation := NewFactorization(2, f.max, f.primeIndex)
    accumulation.Factor()
    for i := 3; i <= f.n; i++ {
        current := NewFactorization(i, f.max, f.primeIndex)
        current.Factor()
        fmt.Printf("TRACER acc Compute(%d) ... %v\n", f.n, accumulation.GetPureFactorsString())
        fmt.Printf("TRACER cur Compute(%d) ... %v\n", f.n, current.GetPureFactorsString())
        accumulation = accumulation.Multiply(current)
    }

    f.factorization = accumulation
}

func (f Factorial) String() string {
    return fmt.Sprintf("%d! => %s", f.n, f.factorization.String())
}

func (f Factorial) GetPureFactorsString() string {
    return f.factorization.GetPureFactorsString()
}
