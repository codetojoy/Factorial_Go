
package factor

import (
    "errors"
    "fmt"
    "strings"

    "github.com/codetojoy/config"
)

type Factors struct {
    config config.Config
    factors []int
}

func NewFactorsForTesting(config config.Config, factors[] int) Factors {
    return Factors{config: config, factors: factors}
}

func NewFactors(config config.Config) Factors {
    result := Factors{config: config}

    // this seems wrong:
    for i := 0; i <= config.HighestPrime; i++ {
        result.factors = append(result.factors, 0)
    }

    return result
}

func (f *Factors) Multiply(g Factors) (Factors, error) {
    result := NewFactors(f.config)

    // fmt.Printf("factors.Multiply cp 1 :: %d %d %d %d\n", f.max, len(f.factors), len(g.factors), len(result.factors))

    var e error = nil
    if len(f.factors) == len(g.factors) {
        for i := 0; i < len(f.factors); i++ {
            result.factors[i] = f.factors[i] + g.factors[i]
        }
    } else {
        fmt.Printf("TRACER Multiply.else error condition ?? \n")
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
    index := f.config.PrimeIndex.GetIndexForPrime(prime)
    f.factors[index] = value
}

func (f Factors) String() string {
    var result = strings.Builder{}

    if len(f.factors) == 0 {
        result.WriteString("ERROR: NONE")
    }

    for index, value := range f.factors {
        if value > 0 {
            prime := f.config.PrimeIndex.GetPrimeForIndex(index)
            result.WriteString(fmt.Sprintf("%d: %d, ", prime, value))
        }
    }

    return result.String()
}
