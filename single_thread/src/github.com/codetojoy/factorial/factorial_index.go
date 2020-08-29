
package factorial

import (
    "github.com/codetojoy/config"
)

type FactorialIndex struct {
    config config.Config
    table map[int]Factorial
}

func NewFactorialIndex(config config.Config) *FactorialIndex {
    table := make(map[int]Factorial)
    return &FactorialIndex{config: config, table: table}
}

func (fi *FactorialIndex) Get(n int) Factorial {
    result, ok := fi.table[n]

    if ! ok {
        result = NewFactorial(n, fi.config)
        result.Compute()
        fi.table[n] = result
    }

    return result
}
