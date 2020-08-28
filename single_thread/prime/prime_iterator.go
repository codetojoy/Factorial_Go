
package prime

type PrimeIterator struct {
    index int
    primeIndex PrimeIndex
}

func NewPrimeIterator(primeIndex PrimeIndex) PrimeIterator {
    return PrimeIterator{index: 0, primeIndex: primeIndex}
}

func (pi *PrimeIterator) Next() int {
    result := pi.primeIndex.GetPrimeForIndex(pi.index)
    pi.index++
    return result
}

func (pi *PrimeIterator) Reset() {
    pi.index = 0
}
