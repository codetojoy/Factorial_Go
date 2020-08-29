
package prime

import (
    "math/big"
)

type PrimeIndex struct {
    primes []int
    index_map map[int]int
}

func New(index int) PrimeIndex {
    result := PrimeIndex{}
    result.initialize(index)
    return result
}

// TODO: make this a constructor function?
func (p *PrimeIndex) initialize(index int) {
    p.index_map = make(map[int]int)
    num_primes := len(p.primes)

    if num_primes == 0 {
        p.primes = append(p.primes, 2)
        num_primes = len(p.primes)
        p.index_map[2] = 0
    }

    isDone := false
    num_primes = len(p.primes)
    counter := p.primes[num_primes - 1] + 1
    ordinal := num_primes - 1
    for ! isDone {
        // fmt.Printf("TRACER GetIthPrime i: %d c: %d o: %d\n", index, counter, ordinal)
        if big.NewInt(int64(counter)).ProbablyPrime(0) {
            p.primes = append(p.primes, counter)
            ordinal++
            p.index_map[counter] = ordinal
        }

        if ordinal == index {
            isDone = true
        }
        counter++
    }
    num_primes = len(p.primes)
}

func (p *PrimeIndex) IsPrime(prime int) bool {
    return (prime == 2 || p.index_map[prime] > 0)
}

// p[2] = 0
// p[3] = 1
// etc
func (p *PrimeIndex) GetIndexForPrime(prime int) int {
    return p.index_map[prime]
}

// p[0] = 2
// p[1] = 3
// etc
func (p *PrimeIndex) GetPrimeForIndex(index int) int {
    return p.primes[index]
}
