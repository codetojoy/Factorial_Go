
package factor

import (
    "errors"
    "fmt"
    "strings"
    "github.com/codetojoy/prime"
)

type Factors struct {
    factors []int
    primeIndex prime.PrimeIndex
    max int
}

func NewFactorsForTesting(max int, factors[] int) Factors {
    return Factors{factors: factors, max: max}
}

func NewFactors(max int, primeIndex prime.PrimeIndex) Factors {
    result := Factors{max: max, primeIndex: primeIndex}

    // this seems wrong:
    for i := 0; i <= max; i++ {
        result.factors = append(result.factors, 0)
    }

    return result
}

func (f *Factors) Multiply(g Factors) (Factors, error) {
    result := NewFactors(f.max, f.primeIndex)

    // fmt.Printf("factors.Multiply cp 1 :: %d %d %d %d\n", f.max, len(f.factors), len(g.factors), len(result.factors))

    var e error = nil
    if len(f.factors) == len(g.factors) {
        for i := 0; i < len(f.factors); i++ {
            result.factors[i] = f.factors[i] + g.factors[i]
        }
    } else {
        e = errors.New(fmt.Sprintf("illegal Factors %d %d", len(f.factors), len(g.factors)))
    }

    return result, e
}

func (f *Factors) Equals(g Factors) bool {
    result := (len(f.factors) == len(g.factors))

    index := 0
    for (result) && (index < len(f.factors)) {
        result = f.factors[index] == g.factors[index]
        index++
    }

    return result
}

func (f *Factors) Put(prime int, value int) {
    index := f.primeIndex.GetIndexForPrime(prime)
    f.factors[index] = value
}

func (f Factors) String() string {
    var result = strings.Builder{}

    if len(f.factors) == 0 {
        result.WriteString("ERROR: NONE")
    }

    for index, value := range f.factors {
        if value > 0 {
            prime := f.primeIndex.GetPrimeForIndex(index)
            result.WriteString(fmt.Sprintf("%d: %d, ", prime, value))
        }
    }

    return result.String()
}
