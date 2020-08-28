
package factorial

import (
    "fmt"
    "github.com/codetojoy/factor"
    "github.com/codetojoy/prime"
)

type Factorial struct {
    n int
    max int
    factorization factor.Factorization
    primeIndex prime.PrimeIndex
}

func NewFactorial(n int, max int, primeIndex prime.PrimeIndex) Factorial {
    return Factorial{n: n, max: max, primeIndex: primeIndex}
}

func (f *Factorial) Compute() {
    accumulation := factor.NewFactorization(2, f.max, f.primeIndex)
    accumulation.Factor()
    for i := 3; i <= f.n; i++ {
        current := factor.NewFactorization(i, f.max, f.primeIndex)
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
