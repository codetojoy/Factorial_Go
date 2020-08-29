
package factorial

import (
    "fmt"

    "github.com/codetojoy/config"
    "github.com/codetojoy/factor"
)

type Factorial struct {
    n int
    config config.Config
    factorization factor.Factorization
}

func NewFactorial(n int, config config.Config) Factorial {
    return Factorial{n: n, config: config}
}

func (f *Factorial) Compute() {
    accumulation := factor.NewFactorization(2, f.config)
    accumulation.Factor()
    for i := 3; i <= f.n; i++ {
        current := factor.NewFactorization(i, f.config)
        current.Factor()
        // fmt.Printf("TRACER acc Compute(%d) ... %v\n", f.n, accumulation.GetPureFactorsString())
        // fmt.Printf("TRACER cur Compute(%d) ... %v\n", f.n, current.GetPureFactorsString())
        accumulation = accumulation.Multiply(current)
    }

    f.factorization = accumulation
}

func (f *Factorial) IsProductMatch(b Factorial, c Factorial) (bool, error) {
    return f.factorization.IsProductMatch(b.factorization, c.factorization)
}

func (f Factorial) String() string {
    return fmt.Sprintf("%d! => %s", f.n, f.factorization.String())
}

func (f Factorial) GetPureFactorsString() string {
    return f.factorization.GetPureFactorsString()
}
